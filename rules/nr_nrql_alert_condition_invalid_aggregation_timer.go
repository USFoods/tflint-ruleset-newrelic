package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// NrNrqlAlerConditionInvalidAggregationTimerRule checks whether newrelic_nrql_alert_condition has valid aggregation timer
type NrNrqlAlerConditionInvalidAggregationTimerRule struct {
	tflint.DefaultRule

	resourceType string
	min          int
	max          int
}

// NewNrNrqlAlerConditionInvalidAggregationTimerRule returns a new rule
func NewNrNrqlAlerConditionInvalidAggregationTimerRule() *NrNrqlAlerConditionInvalidAggregationTimerRule {
	return &NrNrqlAlerConditionInvalidAggregationTimerRule{
		resourceType: "newrelic_nrql_alert_condition",
		min:          0,
		max:          3600,
	}
}

// Name returns the rule name
func (r *NrNrqlAlerConditionInvalidAggregationTimerRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_aggregation_timer"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlerConditionInvalidAggregationTimerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlerConditionInvalidAggregationTimerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlerConditionInvalidAggregationTimerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_nrql_alert_condition has valid aggregation timer
func (r *NrNrqlAlerConditionInvalidAggregationTimerRule) Check(runner tflint.Runner) error {
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
