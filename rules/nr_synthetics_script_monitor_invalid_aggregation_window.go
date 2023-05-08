package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// NrSyntheticsScriptMonitorInvalidAggregationWindowRule checks whether newrelic_synthetics_script_monitor has valid aggregation_window
type NrSyntheticsScriptMonitorInvalidAggregationWindowRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewNrSyntheticsScriptMonitorInvalidAggregationWindowRule returns a new rule
func NewNrSyntheticsScriptMonitorInvalidAggregationWindowRule() *NrSyntheticsScriptMonitorInvalidAggregationWindowRule {
	return &NrSyntheticsScriptMonitorInvalidAggregationWindowRule{
		resourceType:  "newrelic_synthetics_script_monitor",
		attributeName: "aggregation_window",
	}
}

// Name returns the rule name
func (r *NrSyntheticsScriptMonitorInvalidAggregationWindowRule) Name() string {
	return "newrelic_synthetics_script_monitor_invalid_aggregation_window"
}

// Enabled returns whether the rule is enabled by default
func (r *NrSyntheticsScriptMonitorInvalidAggregationWindowRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrSyntheticsScriptMonitorInvalidAggregationWindowRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrSyntheticsScriptMonitorInvalidAggregationWindowRule) Link() string {
	return ""
}

// Check checks whether newrelic_synthetics_script_monitor has valid aggregation_window
func (r *NrSyntheticsScriptMonitorInvalidAggregationWindowRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(aggregationWindow int) error {
			if aggregationWindow < 30 || aggregationWindow > 900 {
				runner.EmitIssue(
					r,
					fmt.Sprintf("'%d' is invalid aggregation_window", aggregationWindow),
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
