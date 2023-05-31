package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

// NrNrqlAlerConditionInvalidSlidyByRule checks whether newrelic_nrql_alert_condition has valid slidy_by
type NrNrqlAlerConditionInvalidSlidyByRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewNrNrqlAlerConditionInvalidSlidyByRule returns a new rule
func NewNrNrqlAlerConditionInvalidSlidyByRule() *NrNrqlAlerConditionInvalidSlidyByRule {
	return &NrNrqlAlerConditionInvalidSlidyByRule{
		resourceType: "newrelic_nrql_alert_condition",
	}
}

// Name returns the rule name
func (r *NrNrqlAlerConditionInvalidSlidyByRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_slidy_by"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlerConditionInvalidSlidyByRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlerConditionInvalidSlidyByRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlerConditionInvalidSlidyByRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_nrql_alert_condition has valid slidy_by
func (r *NrNrqlAlerConditionInvalidSlidyByRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "aggregation_window"},
			{Name: "slide_by"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		slideByAttr, slideByExists := resource.Body.Attributes["slide_by"]
		windowsAttr, windowExists := resource.Body.Attributes["aggregation_window"]

		if !slideByExists || !windowExists {
			continue
		}

		var slideByCty cty.Value

		if err := runner.EvaluateExpr(slideByAttr.Expr, &slideByCty, nil); err != nil {
			return err
		}

		if slideByCty.IsNull() || !slideByCty.IsKnown() {
			continue
		}

		var slideBy int
		if err := gocty.FromCtyValue(slideByCty, &slideBy); err != nil {
			return err
		}

		var windowCty cty.Value
		if err := runner.EvaluateExpr(windowsAttr.Expr, &windowCty, nil); err != nil {
			return err
		}

		// Default window for aggregation_window is 60
		aggregationWindow := 60
		if !windowCty.IsNull() && windowCty.IsKnown() {
			if err := gocty.FromCtyValue(windowCty, &aggregationWindow); err != nil {
				return err
			}
		}

		// slide_by must be less than aggregation_window
		if slideBy > aggregationWindow || aggregationWindow%slideBy != 0 {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("'%d' is invalid value for slide_by", slideBy),
				slideByAttr.Expr.Range(),
			)
		}
	}

	return nil
}
