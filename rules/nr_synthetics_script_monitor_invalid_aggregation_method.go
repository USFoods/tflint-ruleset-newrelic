package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// NrSyntheticsScriptMonitorInvalidAggregationMethodRule checks whether newrelic_synthetics_script_monitor has valid aggregation_method
type NrSyntheticsScriptMonitorInvalidAggregationMethodRule struct {
	tflint.DefaultRule

	resourceType       string
	attributeName      string
	aggregationMethods map[string]bool
}

// NewNrSyntheticsScriptMonitorInvalidAggregationMethodRule returns a new rule
func NewNrSyntheticsScriptMonitorInvalidAggregationMethodRule() *NrSyntheticsScriptMonitorInvalidAggregationMethodRule {
	return &NrSyntheticsScriptMonitorInvalidAggregationMethodRule{
		resourceType:  "newrelic_synthetics_script_monitor",
		attributeName: "aggregation_method",
		aggregationMethods: map[string]bool{
			"cadence":     true,
			"event_flow":  true,
			"event_timer": true,
		},
	}
}

// Name returns the rule name
func (r *NrSyntheticsScriptMonitorInvalidAggregationMethodRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_aggregation_method"
}

// Enabled returns whether the rule is enabled by default
func (r *NrSyntheticsScriptMonitorInvalidAggregationMethodRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrSyntheticsScriptMonitorInvalidAggregationMethodRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrSyntheticsScriptMonitorInvalidAggregationMethodRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_synthetics_script_monitor has valid aggregation_method
func (r *NrSyntheticsScriptMonitorInvalidAggregationMethodRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]

		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(aggregationMethod string) error {
			if !r.aggregationMethods[aggregationMethod] {
				return runner.EmitIssue(
					r,
					fmt.Sprintf("'%s' is invalid aggregation method", aggregationMethod),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)

		if err != nil {
			return err
		}
	}

	return nil
}
