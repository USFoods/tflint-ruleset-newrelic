package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// NrSyntheticsMonitorInvalidAttributesRule checks whether newrelic_synthetics_monitor has valid attributes
type NrSyntheticsMonitorInvalidAttributesRule struct {
	tflint.DefaultRule
}

// NewNrSyntheticsMonitorInvalidAttributesRule returns a new rule
func NewNrSyntheticsMonitorInvalidAttributesRule() *NrSyntheticsMonitorInvalidAttributesRule {
	return &NrSyntheticsMonitorInvalidAttributesRule{}
}

// Name returns the rule name
func (r *NrSyntheticsMonitorInvalidAttributesRule) Name() string {
	return "newrelic_synthetics_monitor_invalid_attributes"
}

// Enabled returns whether the rule is enabled by default
func (r *NrSyntheticsMonitorInvalidAttributesRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrSyntheticsMonitorInvalidAttributesRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrSyntheticsMonitorInvalidAttributesRule) Link() string {
	return ""
}

// Check checks whether newrelic_synthetics_monitor has valid attributes
func (r *NrSyntheticsMonitorInvalidAttributesRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("newrelic_synthetics_monitor", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "type"},
			{Name: "treat_redirect_as_failure"},
			{Name: "bypass_head_request"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		typeAttribute, exists := resource.Body.Attributes["type"]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(typeAttribute.Expr, func(typeValue string) bool {
			fmt.Print(typeValue)

			return true
		}, nil)

		if err != nil {
			return err
		}
	}

	return nil
}
