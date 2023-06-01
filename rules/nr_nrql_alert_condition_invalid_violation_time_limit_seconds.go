package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty"

	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// TODO: Write the rule's description here
// NrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule checks ...
type NrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	min           int
	max           int
}

// NewNrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule returns new rule with default attributes
func NewNrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule() *NrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule {
	return &NrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule{
		// TODO: Write resource type and attribute name here
		resourceType:  "newrelic_nrql_alert_condition",
		attributeName: "violation_time_limit_seconds",
		min:           300,
		max:           2592000,
	}
}

// Name returns the rule name
func (r *NrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule) Name() string {
	return "nr_nrql_alert_condition_invalid_violation_time_limit_seconds"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *NrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule) Check(runner tflint.Runner) error {
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

		var timeLimit int
		if err := runner.EvaluateExpr(attribute.Expr, &timeLimit, nil); err != nil {
			return err
		}

		if timeLimit < r.min || timeLimit > r.max {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("'%d' is an invalid value for violation_time_limit_seconds", timeLimit),
				attribute.Expr.Range(),
			)
		}
	}

	return nil
}
