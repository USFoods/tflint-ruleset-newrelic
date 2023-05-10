package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// NrSyntheticsScriptMonitorInvalidSlidyByRule checks whether newrelic_synthetics_script_monitor has valid slidy_by
type NrSyntheticsScriptMonitorInvalidSlidyByRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewNrSyntheticsScriptMonitorInvalidSlidyByRule returns a new rule
func NewNrSyntheticsScriptMonitorInvalidSlidyByRule() *NrSyntheticsScriptMonitorInvalidSlidyByRule {
	return &NrSyntheticsScriptMonitorInvalidSlidyByRule{
		resourceType: "newrelic_synthetics_script_monitor",
	}
}

// Name returns the rule name
func (r *NrSyntheticsScriptMonitorInvalidSlidyByRule) Name() string {
	return "newrelic_synthetics_script_monitor_invalid_slidy_by"
}

// Enabled returns whether the rule is enabled by default
func (r *NrSyntheticsScriptMonitorInvalidSlidyByRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrSyntheticsScriptMonitorInvalidSlidyByRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrSyntheticsScriptMonitorInvalidSlidyByRule) Link() string {
	return ""
}

// Check checks whether newrelic_synthetics_script_monitor has valid slidy_by
func (r *NrSyntheticsScriptMonitorInvalidSlidyByRule) Check(runner tflint.Runner) error {
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
		attr, exists := resource.Body.Attributes["aggregation_window"]

		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attr.Expr, func(aggregationWindow int) error {
			attr, exists := resource.Body.Attributes["slide_by"]

			if !exists {
				return nil
			}

			err := runner.EvaluateExpr(attr.Expr, func(slideBy int) error {
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

				return nil
			}, nil)

			if err != nil {
				return err
			}

			return nil
		}, nil)

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}