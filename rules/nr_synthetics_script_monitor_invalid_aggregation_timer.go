package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// NrSyntheticsScriptMonitorInvalidAggregationTimerRule checks whether newrelic_synthetics_script_monitor has valid aggregation timer
type NrSyntheticsScriptMonitorInvalidAggregationTimerRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewNrSyntheticsScriptMonitorInvalidAggregationTimerRule returns a new rule
func NewNrSyntheticsScriptMonitorInvalidAggregationTimerRule() *NrSyntheticsScriptMonitorInvalidAggregationTimerRule {
	return &NrSyntheticsScriptMonitorInvalidAggregationTimerRule{
		resourceType: "newrelic_synthetics_script_monitor",
	}
}

// Name returns the rule name
func (r *NrSyntheticsScriptMonitorInvalidAggregationTimerRule) Name() string {
	return "newrelic_synthetics_script_monitor_invalid_aggregation_timer"
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
	return ""
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
		if attr, exists := resource.Body.Attributes["aggregation_method"]; exists {
			err := runner.EvaluateExpr(attr.Expr, func(aggregationMethod string) error {

				if attr, exists := resource.Body.Attributes["aggregation_timer"]; exists {
					err := runner.EvaluateExpr(attr.Expr, func(aggregationTimer int) error {

						if aggregationMethod != "event_timer" {
							runner.EmitIssue(
								r,
								fmt.Sprintf("aggregation_timer invalid for aggregation_method '%s'", aggregationMethod),
								attr.Expr.Range(),
							)
						}

						if aggregationTimer < 0 || aggregationTimer > 3600 {
							runner.EmitIssue(
								r,
								fmt.Sprintf("'%d' is invalid aggregation_timer", aggregationTimer),
								attr.Expr.Range(),
							)
						}

						return nil
					}, nil)

					if err != nil {
						return err
					}
				}

				return nil

			}, nil)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
