package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// NrNrqlAlerConditionInvalidPeriodRule checks whether newrelic_nrql_alert_condition has valid period
type NrNrqlAlerConditionInvalidPeriodRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	periods       map[string]bool
}

// NewNrNrqlAlerConditionInvalidPeriodRule returns a new rule
func NewNrNrqlAlerConditionInvalidPeriodRule() *NrNrqlAlerConditionInvalidPeriodRule {
	return &NrNrqlAlerConditionInvalidPeriodRule{
		resourceType:  "newrelic_nrql_alert_condition",
		attributeName: "period",
		periods: map[string]bool{
			"EVERY_MINUTE":     true,
			"EVERY_5_MINUTES":  true,
			"EVERY_10_MINUTES": true,
			"EVERY_15_MINUTES": true,
			"EVERY_30_MINUTES": true,
			"EVERY_HOUR":       true,
			"EVERY_6_HOURS":    true,
			"EVERY_12_HOURS":   true,
			"EVERY_DAY":        true,
		},
	}
}

// Name returns the rule name
func (r *NrNrqlAlerConditionInvalidPeriodRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_period"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlerConditionInvalidPeriodRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlerConditionInvalidPeriodRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlerConditionInvalidPeriodRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_nrql_alert_condition has valid period
func (r *NrNrqlAlerConditionInvalidPeriodRule) Check(runner tflint.Runner) error {
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
			if !r.periods[period] {
				return runner.EmitIssue(
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
