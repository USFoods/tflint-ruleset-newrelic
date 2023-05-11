package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// NrSyntheticsScriptMonitorInvalidAggregationDelayRule checks whether newrelic_synthetics_monitor has valid aggregation_delay
type NrSyntheticsScriptMonitorInvalidAggregationDelayRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewNrSyntheticsScriptMonitorInvalidAggregationDelayRule returns a new rule
func NewNrSyntheticsScriptMonitorInvalidAggregationDelayRule() *NrSyntheticsScriptMonitorInvalidAggregationDelayRule {
	return &NrSyntheticsScriptMonitorInvalidAggregationDelayRule{
		resourceType: "newrelic_synthetics_script_monitor",
	}
}

// Name returns the rule name
func (r *NrSyntheticsScriptMonitorInvalidAggregationDelayRule) Name() string {
	return "newrelic_synthetics_script_monitor_invalid_aggregation_delay"
}

// Enabled returns whether the rule is enabled by default
func (r *NrSyntheticsScriptMonitorInvalidAggregationDelayRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrSyntheticsScriptMonitorInvalidAggregationDelayRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrSyntheticsScriptMonitorInvalidAggregationDelayRule) Link() string {
	return ""
}

// Check checks whether newrelic_synthetics_script_monitor has valid aggregation_delay
func (r *NrSyntheticsScriptMonitorInvalidAggregationDelayRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "aggregation_method"},
			{Name: "aggregation_delay"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if attr, exists := resource.Body.Attributes["aggregation_method"]; exists {
			err := runner.EvaluateExpr(attr.Expr, func(aggregationMethod string) error {

				if attr, exists := resource.Body.Attributes["aggregation_delay"]; exists {
					err := runner.EvaluateExpr(attr.Expr, func(aggregationDelay int) error {

						if aggregationMethod == "event_timer" {
							return runner.EmitIssue(
								r,
								fmt.Sprintf("aggregation_delay invalid for aggregation_method '%s'", aggregationMethod),
								attr.Expr.Range(),
							)
						}

						if aggregationMethod == "event_flow" {
							if aggregationDelay > 1200 {
								return runner.EmitIssue(
									r,
									fmt.Sprintf("'%d' invalid aggregation_delay for aggregation_method '%s'", aggregationDelay, aggregationMethod),
									attr.Expr.Range(),
								)
							}
						}

						if aggregationMethod == "cadence" {
							if aggregationDelay > 3600 {
								return runner.EmitIssue(
									r,
									fmt.Sprintf("'%d' invalid aggregation_delay for aggregation_method '%s'", aggregationDelay, aggregationMethod),
									attr.Expr.Range(),
								)
							}
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
