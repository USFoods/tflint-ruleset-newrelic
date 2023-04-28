package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"golang.org/x/exp/slices"
)

// NrSyntheticsMonitorInvalidPeriodRule checks whether newrelic_synthetics_monitor has valid period
type NrSyntheticsMonitorInvalidPeriodRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	periodTypes   []string
}

// NewNrSyntheticsMonitorInvalidPeriodRule returns a new rule
func NewNrSyntheticsMonitorInvalidPeriodRule() *NrSyntheticsMonitorInvalidPeriodRule {
	return &NrSyntheticsMonitorInvalidPeriodRule{
		resourceType:  "newrelic_synthetics_monitor",
		attributeName: "period",
		periodTypes: []string{
			"EVERY_MINUTE",
			"EVERY_5_MINUTES",
			"EVERY_10_MINUTES",
			"EVERY_15_MINUTES",
			"EVERY_30_MINUTES",
			"EVERY_HOUR",
			"EVERY_6_HOURS",
			"EVERY_12_HOURS",
			"EVERY_DAY",
		},
	}
}

// Name returns the rule name
func (r *NrSyntheticsMonitorInvalidPeriodRule) Name() string {
	return "newrelic_synthetics_monitor_invalid_period"
}

// Enabled returns whether the rule is enabled by default
func (r *NrSyntheticsMonitorInvalidPeriodRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrSyntheticsMonitorInvalidPeriodRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrSyntheticsMonitorInvalidPeriodRule) Link() string {
	return ""
}

// Check checks whether newrelic_synthetics_monitor has valid period
func (r *NrSyntheticsMonitorInvalidPeriodRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(period string) error {
			if !slices.Contains(r.periodTypes, period) {
				runner.EmitIssue(
					r,
					fmt.Sprintf("'%s' is invalid period", period),
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
