package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
	"github.com/zclconf/go-cty/cty"
)

// TODO: Write the rule's description here
// NrNrqlAlertConditionInvalidEvaluationDelayRule checks ...
type NrNrqlAlertConditionInvalidEvaluationDelayRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewNrNrqlAlertConditionInvalidEvaluationDelayRule returns new rule with default attributes
func NewNrNrqlAlertConditionInvalidEvaluationDelayRule() *NrNrqlAlertConditionInvalidEvaluationDelayRule {
	return &NrNrqlAlertConditionInvalidEvaluationDelayRule{
		// TODO: Write resource type and attribute name here
		resourceType:  "newrelic_nrql_alert_condition",
		attributeName: "evaluation_delay",
		max:           1200,
	}
}

// Name returns the rule name
func (r *NrNrqlAlertConditionInvalidEvaluationDelayRule) Name() string {
	return "nr_nrql_alert_condition_invalid_evaluation_delay"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlertConditionInvalidEvaluationDelayRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlertConditionInvalidEvaluationDelayRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlertConditionInvalidEvaluationDelayRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *NrNrqlAlertConditionInvalidEvaluationDelayRule) Check(runner tflint.Runner) error {
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

		if attrCty.IsNull() || !attrCty.IsKnown() {
			continue
		}

		var evaluationDelay int
		if err := runner.EvaluateExpr(attribute.Expr, &evaluationDelay, nil); err != nil {
			return err
		}

		if evaluationDelay > r.max {
			runner.EmitIssue(
				r,
				fmt.Sprintf("'%d' is an invalid value for evaluation_delay", evaluationDelay),
				attribute.Expr.Range(),
			)
		}
	}

	return nil
}
