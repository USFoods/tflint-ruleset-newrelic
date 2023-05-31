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
// NrNrqlAlertConditionInvalidBaselineDirectionRule checks ...
type NrNrqlAlertConditionInvalidBaselineDirectionRule struct {
	tflint.DefaultRule

	resourceType       string
	attributeName      string
	baselineDirections map[string]bool
}

// NewNrNrqlAlertConditionInvalidBaselineDirectionRule returns new rule with default attributes
func NewNrNrqlAlertConditionInvalidBaselineDirectionRule() *NrNrqlAlertConditionInvalidBaselineDirectionRule {
	return &NrNrqlAlertConditionInvalidBaselineDirectionRule{
		// TODO: Write resource type and attribute name here
		resourceType:  "newrelic_nrql_alert_condition",
		attributeName: "baseline_direction",
		baselineDirections: map[string]bool{
			"lower_only":      true,
			"upper_only":      true,
			"upper_and_lower": true,
		},
	}
}

// Name returns the rule name
func (r *NrNrqlAlertConditionInvalidBaselineDirectionRule) Name() string {
	return "nr_nrql_alert_condition_invalid_baseline_direction"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlertConditionInvalidBaselineDirectionRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlertConditionInvalidBaselineDirectionRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlertConditionInvalidBaselineDirectionRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *NrNrqlAlertConditionInvalidBaselineDirectionRule) Check(runner tflint.Runner) error {
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
		directionAttr, directionExists := resource.Body.Attributes[r.attributeName]

		if !directionExists {
			continue
		}

		var directionCty cty.Value
		if err := runner.EvaluateExpr(directionAttr.Expr, &directionCty, nil); err != nil {
			return err
		}

		if directionCty.IsNull() || !directionCty.IsKnown() {
			continue
		}

		var baselineDirection string
		if err := gocty.FromCtyValue(directionCty, &baselineDirection); err != nil {
			return err
		}

		if !r.baselineDirections[strings.ToLower(baselineDirection)] {
			runner.EmitIssue(
				r,
				fmt.Sprintf("'%s' is invalid value for baseline_direction", baselineDirection),
				directionAttr.Expr.Range(),
			)
		}
	}

	return nil
}
