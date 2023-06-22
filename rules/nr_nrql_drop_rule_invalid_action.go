package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// TODO: Write the rule's description here
// NrNrqlDropRuleInvalidActionRule checks ...
type NrNrqlDropRuleInvalidActionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	actions       map[string]bool
}

// NewNrNrqlDropRuleInvalidActionRule returns new rule with default attributes
func NewNrNrqlDropRuleInvalidActionRule() *NrNrqlDropRuleInvalidActionRule {
	return &NrNrqlDropRuleInvalidActionRule{
		// TODO: Write resource type and attribute name here
		resourceType:  "newrelic_nrql_drop_rule",
		attributeName: "action",
		actions: map[string]bool{
			"drop_data":                              true,
			"drop_attributes":                        true,
			"drop_attributes_from_metric_aggregates": true,
		},
	}
}

// Name returns the rule name
func (r *NrNrqlDropRuleInvalidActionRule) Name() string {
	return "nr_nrql_drop_rule_invalid_action"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlDropRuleInvalidActionRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *NrNrqlDropRuleInvalidActionRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlDropRuleInvalidActionRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *NrNrqlDropRuleInvalidActionRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(actionType string) error {
			if !r.actions[actionType] {
				return runner.EmitIssue(
					r,
					fmt.Sprintf("'%s' is invalid value for action", actionType),
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
