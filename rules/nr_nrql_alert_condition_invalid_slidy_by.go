package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
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

		attr, ok := resource.Body.Attributes["slide_by"]

		if !ok {
			continue
		}

		var slideBy int
		err := runner.EvaluateExpr(attr.Expr, &slideBy, nil)

		if err != nil {
			return err
		}

		attr, ok = resource.Body.Attributes["aggregation_window"]

		if !ok {
			continue
		}

		var aggregationWindow int
		err = runner.EvaluateExpr(attr.Expr, &aggregationWindow, nil)

		if err != nil {
			return err
		}

		attr, exists := resource.Body.Attributes["aggregation_window"]

		if !exists {
			continue
		}

		// slide_by must be less than aggregation_window
		if slideBy > aggregationWindow {
			return runner.EmitIssue(
				r,
				"slide_by is greater than aggregation_window",
				attr.Expr.Range(),
			)
		}

		// slide_by must be a factor of aggregation_window
		if aggregationWindow%slideBy != 0 {
			return runner.EmitIssue(
				r,
				"slide_by is not a factor of aggregation_window",
				attr.Expr.Range(),
			)
		}
	}

	return nil
}
