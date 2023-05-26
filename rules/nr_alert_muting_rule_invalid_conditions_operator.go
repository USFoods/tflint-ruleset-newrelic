package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// TODO: Write the rule's description here
// NrAlertMutingRuleInvalidConditionsOperatorRule checks ...
type NrAlertMutingRuleInvalidConditionsOperatorRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	operatorTypes map[string]bool
}

// NewNrAlertMutingRuleInvalidConditionsOperatorRule returns new rule with default attributes
func NewNrAlertMutingRuleInvalidConditionsOperatorRule() *NrAlertMutingRuleInvalidConditionsOperatorRule {
	return &NrAlertMutingRuleInvalidConditionsOperatorRule{
		// TODO: Write resource type and attribute name here
		resourceType:  "newrelic_alert_muting_rule",
		attributeName: "operator",
		operatorTypes: map[string]bool{
			"ANY":             true,
			"CONTAINS":        true,
			"ENDS_WITH":       true,
			"EQUALS":          true,
			"IN":              true,
			"IS_BLANK":        true,
			"IS_NOT_BLANK":    true,
			"NOT_CONTAINS":    true,
			"NOT_ENDS_WITH":   true,
			"NOT_EQUALS":      true,
			"NOT_IN":          true,
			"NOT_STARTS_WITH": true,
			"STARTS_WITH":     true,
		},
	}
}

// Name returns the rule name
func (r *NrAlertMutingRuleInvalidConditionsOperatorRule) Name() string {
	return "nr_alert_muting_rule_invalid_conditions_operator"
}

// Enabled returns whether the rule is enabled by default
func (r *NrAlertMutingRuleInvalidConditionsOperatorRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *NrAlertMutingRuleInvalidConditionsOperatorRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrAlertMutingRuleInvalidConditionsOperatorRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *NrAlertMutingRuleInvalidConditionsOperatorRule) Check(runner tflint.Runner) error {
	// TODO: Write the implementation here. See this documentation for what tflint.Runner can do.
	//       https://pkg.go.dev/github.com/terraform-linters/tflint-plugin-sdk/tflint#Runner

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "condition",
				Body: &hclext.BodySchema{
					Blocks: []hclext.BlockSchema{
						{
							Type: "conditions",
							Body: &hclext.BodySchema{
								Attributes: []hclext.AttributeSchema{
									{Name: r.attributeName},
								},
							},
						},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		for _, condition := range resource.Body.Blocks {
			for _, conditions := range condition.Body.Blocks {
				attribute, exists := conditions.Body.Attributes[r.attributeName]

				if !exists {
					continue
				}

				err := runner.EvaluateExpr(attribute.Expr, func(operatorType string) error {
					if !r.operatorTypes[operatorType] {
						return runner.EmitIssue(
							r,
							fmt.Sprintf("'%s' is invalid value for conditions operator", operatorType),
							attribute.Expr.Range(),
						)
					}
					return nil
				}, nil)

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
