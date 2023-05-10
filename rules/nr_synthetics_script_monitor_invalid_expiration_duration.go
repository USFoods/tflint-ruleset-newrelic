package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// NrSyntheticsScriptMonitorInvalidExpirationDurationRule checks whether newrelic_synthetics_script_monitor has valid expiration_duration
type NrSyntheticsScriptMonitorInvalidExpirationDurationRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewNrSyntheticsScriptMonitorInvalidExpirationDurationRule returns a new rule
func NewNrSyntheticsScriptMonitorInvalidExpirationDurationRule() *NrSyntheticsScriptMonitorInvalidExpirationDurationRule {
	return &NrSyntheticsScriptMonitorInvalidExpirationDurationRule{
		resourceType:  "newrelic_synthetics_script_monitor",
		attributeName: "expiration_duration",
	}
}

// Name returns the rule name
func (r *NrSyntheticsScriptMonitorInvalidExpirationDurationRule) Name() string {
	return "newrelic_synthetics_script_monitor_invalid_expiration_duration"
}

// Enabled returns whether the rule is enabled by default
func (r *NrSyntheticsScriptMonitorInvalidExpirationDurationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrSyntheticsScriptMonitorInvalidExpirationDurationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrSyntheticsScriptMonitorInvalidExpirationDurationRule) Link() string {
	return ""
}

// Check checks whether newrelic_synthetics_script_monitor has valid expiration_duration
func (r *NrSyntheticsScriptMonitorInvalidExpirationDurationRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(expirationDuration int) error {
			if expirationDuration < 30 || expirationDuration > 172800 {
				runner.EmitIssue(
					r,
					fmt.Sprintf("'%d' is invalid expiration_duration", expirationDuration),
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
