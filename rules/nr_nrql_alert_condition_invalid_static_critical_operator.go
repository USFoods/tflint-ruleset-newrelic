package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

// TODO: Write the rule's description here
// NrNrqlAlertConditionInvalidStaticCriticalOperatorRule checks ...
type NrNrqlAlertConditionInvalidStaticCriticalOperatorRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	operatorTypes map[string]bool
}

// NewNrNrqlAlertConditionInvalidStaticCriticalOperatorRule returns new rule with default attributes
func NewNrNrqlAlertConditionInvalidStaticCriticalOperatorRule() *NrNrqlAlertConditionInvalidStaticCriticalOperatorRule {
	return &NrNrqlAlertConditionInvalidStaticCriticalOperatorRule{
		// TODO: Write resource type and attribute name here
		resourceType:  "newrelic_nrql_alert_condition",
		attributeName: "operator",
		operatorTypes: map[string]bool{
			"above":           true,
			"above_or_equals": true,
			"below":           true,
			"below_or_equals": true,
			"equals":          true,
			"not_equals":      true,
		},
	}
}

// Name returns the rule name
func (r *NrNrqlAlertConditionInvalidStaticCriticalOperatorRule) Name() string {
	return "nr_nrql_alert_condition_invalid_static_critical_operator"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlertConditionInvalidStaticCriticalOperatorRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlertConditionInvalidStaticCriticalOperatorRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlertConditionInvalidStaticCriticalOperatorRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *NrNrqlAlertConditionInvalidStaticCriticalOperatorRule) Check(runner tflint.Runner) error {
	// TODO: Write the implementation here. See this documentation for what tflint.Runner can do.
	//       https://pkg.go.dev/github.com/terraform-linters/tflint-plugin-sdk/tflint#Runner

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "type"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "critical",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: r.attributeName},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		typeAttr, typeExists := resource.Body.Attributes["type"]

		if !typeExists {
			continue
		}

		var typeCty cty.Value
		err := runner.EvaluateExpr(typeAttr.Expr, &typeCty, nil)

		if typeCty.IsNull() || !typeCty.IsWhollyKnown() {
			continue
		}

		var typeValue string
		if err := gocty.FromCtyValue(typeCty, &typeValue); err != nil {
			return err
		}

		if strings.ToLower(typeValue) == "static" {
			for _, block := range resource.Body.Blocks {
				operatorAttr, operatorExists := block.Body.Attributes[r.attributeName]

				if !operatorExists {
					continue
				}

				var operatorCty cty.Value
				if err := runner.EvaluateExpr(operatorAttr.Expr, &operatorCty, nil); err != nil {
					return err
				}

				if operatorCty.IsNull() || !operatorCty.IsWhollyKnown() {
					continue
				}

				var operatorValue string
				if err := gocty.FromCtyValue(operatorCty, &operatorValue); err != nil {
					return err
				}

				if !r.operatorTypes[strings.ToLower(operatorValue)] {
					return runner.EmitIssue(
						r,
						fmt.Sprintf("'%s' is an invalid value for critical operator with type '%s'", operatorValue, typeValue),
						operatorAttr.Expr.Range(),
					)
				}

			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
