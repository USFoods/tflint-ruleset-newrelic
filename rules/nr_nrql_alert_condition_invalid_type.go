package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"

	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// NrNrqlAlertConditionInvalidTypeRule checks whether newrelic_nrql_alert_condition has valid type
type NrNrqlAlertConditionInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType   string
	attributeName  string
	conditionTypes map[string]bool
}

// NewNrNrqlAlertConditionInvalidTypeRule returns a new rule
func NewNrNrqlAlertConditionInvalidTypeRule() *NrNrqlAlertConditionInvalidTypeRule {
	return &NrNrqlAlertConditionInvalidTypeRule{
		resourceType:  "newrelic_nrql_alert_condition",
		attributeName: "type",
		conditionTypes: map[string]bool{
			"static":   true,
			"baseline": true,
		},
	}
}

// Name returns the rule name
func (r *NrNrqlAlertConditionInvalidTypeRule) Name() string {
	return "nr_nrql_alert_condition_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlertConditionInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlertConditionInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlertConditionInvalidTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_nrql_alert_condition has valid type
func (r *NrNrqlAlertConditionInvalidTypeRule) Check(runner tflint.Runner) error {
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

		var monitorType string
		if err := gocty.FromCtyValue(attrCty, &monitorType); err != nil {
			return err
		}

		if !r.conditionTypes[monitorType] {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("'%s' is invalid value for type", monitorType),
				attribute.Expr.Range(),
			)
		}
	}

	return nil
}
