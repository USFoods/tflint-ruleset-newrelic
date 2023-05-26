package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// NrSyntheticsScriptMonitorInvalidTypeRule checks whether newrelic_synthetics_script_monitor has valid type
type NrSyntheticsScriptMonitorInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	monitorTypes  map[string]bool
}

// NewNrSyntheticsScriptMonitorInvalidTypeRule returns a new rule
func NewNrSyntheticsScriptMonitorInvalidTypeRule() *NrSyntheticsScriptMonitorInvalidTypeRule {
	return &NrSyntheticsScriptMonitorInvalidTypeRule{
		resourceType:  "newrelic_synthetics_script_monitor",
		attributeName: "type",
		monitorTypes: map[string]bool{
			"SCRIPT_API":     true,
			"SCRIPT_BROWSER": true,
		},
	}
}

// Name returns the rule name
func (r *NrSyntheticsScriptMonitorInvalidTypeRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *NrSyntheticsScriptMonitorInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrSyntheticsScriptMonitorInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrSyntheticsScriptMonitorInvalidTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_synthetics_script_monitor has valid type
func (r *NrSyntheticsScriptMonitorInvalidTypeRule) Check(runner tflint.Runner) error {
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
			if !r.monitorTypes[monitorType] {
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
