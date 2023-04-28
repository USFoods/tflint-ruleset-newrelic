package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"golang.org/x/exp/slices"
)

// NrSyntheticsMonitorInvalidTypeRule checks whether newrelic_synthetics_monitor has valid type
type NrSyntheticsMonitorInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	monitorTypes  []string
}

// NewNrSyntheticsMonitorInvalidTypeRule returns a new rule
func NewNrSyntheticsMonitorInvalidTypeRule() *NrSyntheticsMonitorInvalidTypeRule {
	return &NrSyntheticsMonitorInvalidTypeRule{
		resourceType:  "newrelic_synthetics_monitor",
		attributeName: "type",
		monitorTypes: []string{
			"SIMPLE",
			"BROWSER",
		},
	}
}

// Name returns the rule name
func (r *NrSyntheticsMonitorInvalidTypeRule) Name() string {
	return "newrelic_synthetics_monitor_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *NrSyntheticsMonitorInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrSyntheticsMonitorInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrSyntheticsMonitorInvalidTypeRule) Link() string {
	return ""
}

// Check checks whether newrelic_synthetics_monitor has valid type
func (r *NrSyntheticsMonitorInvalidTypeRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(monitorType string) error {
			if !slices.Contains(r.monitorTypes, monitorType) {
				return runner.EmitIssue(
					r,
					fmt.Sprintf("'%s' is invalid monitor type", monitorType),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)

		if err != nil {
			return err
		}
	}

	return nil
}
