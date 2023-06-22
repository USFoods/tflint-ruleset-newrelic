package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"

	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// TODO: Write the rule's description here
// NrNrqlAlertConditionInvalidFillOptionRule checks ...
type NrNrqlAlertConditionInvalidFillOptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	fillOptions   map[string]bool
}

// NewNrNrqlAlertConditionInvalidFillOptionRule returns new rule with default attributes
func NewNrNrqlAlertConditionInvalidFillOptionRule() *NrNrqlAlertConditionInvalidFillOptionRule {
	return &NrNrqlAlertConditionInvalidFillOptionRule{
		// TODO: Write resource type and attribute name here
		resourceType:  "newrelic_nrql_alert_condition",
		attributeName: "fill_option",
		fillOptions: map[string]bool{
			"none":       true,
			"last_value": true,
			"static":     true,
		},
	}
}

// Name returns the rule name
func (r *NrNrqlAlertConditionInvalidFillOptionRule) Name() string {
	return "nr_nrql_alert_condition_invalid_fill_option"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlertConditionInvalidFillOptionRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlertConditionInvalidFillOptionRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlertConditionInvalidFillOptionRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *NrNrqlAlertConditionInvalidFillOptionRule) Check(runner tflint.Runner) error {
	// TODO: Write the implementation here. See this documentation for what tflint.Runner can do.
	//       https://pkg.go.dev/github.com/terraform-linters/tflint-plugin-sdk/tflint#Runner

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

		var fillOption string
		if err := gocty.FromCtyValue(attrCty, &fillOption); err != nil {
			return err
		}

		if !r.fillOptions[fillOption] {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("'%s' is an invalid value for fill_option", fillOption),
				attribute.Expr.Range(),
			)
		}
	}

	return nil
}
