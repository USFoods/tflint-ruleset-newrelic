package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// TODO: Write the rule's description here
// NrAlertMutingRuleInvalidConditionsAttributeRule checks ...
type NrAlertMutingRuleInvalidConditionsAttributeRule struct {
	tflint.DefaultRule

	resourceType   string
	attributeName  string
	attributeTypes map[string]bool
}

// NewNrAlertMutingRuleInvalidConditionsAttributeRule returns new rule with default attributes
func NewNrAlertMutingRuleInvalidConditionsAttributeRule() *NrAlertMutingRuleInvalidConditionsAttributeRule {
	return &NrAlertMutingRuleInvalidConditionsAttributeRule{
		// TODO: Write resource type and attribute name here
		resourceType:  "newrelic_alert_muting_rule",
		attributeName: "attribute",
		attributeTypes: map[string]bool{
			"accountId":           true,
			"conditionId":         true,
			"conditionName":       true,
			"conditionRunbookUrl": true,
			"conditionType":       true,
			"entity.guid":         true,
			"nrqlEventType":       true,
			"nrqlQuery":           true,
			"policyId":            true,
			"policyName":          true,
			"product":             true,
			"targetId":            true,
			"targetName":          true,
		},
	}
}

// Name returns the rule name
func (r *NrAlertMutingRuleInvalidConditionsAttributeRule) Name() string {
	return "nr_alert_muting_rule_invalid_conditions_attribute"
}

// Enabled returns whether the rule is enabled by default
func (r *NrAlertMutingRuleInvalidConditionsAttributeRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *NrAlertMutingRuleInvalidConditionsAttributeRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrAlertMutingRuleInvalidConditionsAttributeRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *NrAlertMutingRuleInvalidConditionsAttributeRule) Check(runner tflint.Runner) error {
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

				err := runner.EvaluateExpr(attribute.Expr, func(attributeType string) error {
					if !r.attributeTypes[attributeType] && !strings.HasPrefix(attributeType, "tags.") {
						return runner.EmitIssue(
							r,
							fmt.Sprintf("'%s' is invalid value for conditions attribute", attributeType),
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
