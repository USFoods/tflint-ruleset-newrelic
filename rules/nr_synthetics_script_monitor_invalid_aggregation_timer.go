package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// NrSyntheticsScriptMonitorInvalidAggregationTimerRule checks whether newrelic_synthetics_script_monitor has valid aggregation timer
type NrSyntheticsScriptMonitorInvalidAggregationTimerRule struct {
	tflint.DefaultRule

	resourceType string
	min          int
	max          int
}

// NewNrSyntheticsScriptMonitorInvalidAggregationTimerRule returns a new rule
func NewNrSyntheticsScriptMonitorInvalidAggregationTimerRule() *NrSyntheticsScriptMonitorInvalidAggregationTimerRule {
	return &NrSyntheticsScriptMonitorInvalidAggregationTimerRule{
		resourceType: "newrelic_synthetics_script_monitor",
		min:          0,
		max:          3600,
	}
}

// Name returns the rule name
func (r *NrSyntheticsScriptMonitorInvalidAggregationTimerRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_aggregation_timer"
}

// Enabled returns whether the rule is enabled by default
func (r *NrSyntheticsScriptMonitorInvalidAggregationTimerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrSyntheticsScriptMonitorInvalidAggregationTimerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrSyntheticsScriptMonitorInvalidAggregationTimerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_synthetics_script_monitor has valid aggregation timer
func (r *NrSyntheticsScriptMonitorInvalidAggregationTimerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "aggregation_method"},
			{Name: "aggregation_timer"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attr, ok := resource.Body.Attributes["aggregation_timer"]

		if !ok {
			continue
		}

		var aggregationTimer int
		err := runner.EvaluateExpr(attr.Expr, &aggregationTimer, nil)

		if err != nil {
			return err
		}

		attr, ok = resource.Body.Attributes["aggregation_method"]

		if !ok {
			continue
		}

		var aggregationMethod string
		err = runner.EvaluateExpr(attr.Expr, &aggregationMethod, nil)

		if err != nil {
			return err
		}

		if aggregationMethod != "event_timer" {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("aggregation_timer invalid for aggregation_method '%s'", aggregationMethod),
				attr.Expr.Range(),
			)
		}

		if aggregationTimer < r.min || aggregationTimer > r.max {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("'%d' is invalid aggregation_timer", aggregationTimer),
				attr.Expr.Range(),
			)
		}
	}

	return nil
}
