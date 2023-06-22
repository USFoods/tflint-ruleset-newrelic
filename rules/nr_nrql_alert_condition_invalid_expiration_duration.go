package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"

	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// NrNrqlAlerConditionInvalidExpirationDurationRule checks whether newrelic_nrql_alert_condition has valid expiration_duration
type NrNrqlAlerConditionInvalidExpirationDurationRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	min           int
	max           int
}

// NewNrNrqlAlerConditionInvalidExpirationDurationRule returns a new rule
func NewNrNrqlAlerConditionInvalidExpirationDurationRule() *NrNrqlAlerConditionInvalidExpirationDurationRule {
	return &NrNrqlAlerConditionInvalidExpirationDurationRule{
		resourceType:  "newrelic_nrql_alert_condition",
		attributeName: "expiration_duration",
		min:           30,
		max:           172800,
	}
}

// Name returns the rule name
func (r *NrNrqlAlerConditionInvalidExpirationDurationRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_expiration_duration"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlerConditionInvalidExpirationDurationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlerConditionInvalidExpirationDurationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlerConditionInvalidExpirationDurationRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_nrql_alert_condition has valid expiration_duration
func (r *NrNrqlAlerConditionInvalidExpirationDurationRule) Check(runner tflint.Runner) error {
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

		var attrCty cty.Value
		if err := runner.EvaluateExpr(attribute.Expr, &attrCty, nil); err != nil {
			return err
		}

		if attrCty.IsNull() || !attrCty.IsWhollyKnown() {
			continue
		}

		var expirationDuration int
		if err := gocty.FromCtyValue(attrCty, &expirationDuration); err != nil {
			return err
		}

		if expirationDuration < r.min || expirationDuration > r.max {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("'%d' is invalid value for expiration_duration", expirationDuration),
				attribute.Expr.Range(),
			)
		}
	}

	return nil
}
