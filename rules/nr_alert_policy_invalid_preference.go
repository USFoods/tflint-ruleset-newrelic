package rules

import (
	"fmt"

	"golang.org/x/exp/slices"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type NrAlertPolicyInvalidPreferenceRule struct {
	tflint.DefaultRule

	resourceType    string
	attributeName   string
	preferenceTypes []string
}

// NewNrAlertPolicyInvalidPreferenceRule returns a new rule
func NewNrAlertPolicyInvalidPreferenceRule() *NrAlertPolicyInvalidPreferenceRule {
	return &NrAlertPolicyInvalidPreferenceRule{
		resourceType:    "newrelic_alert_policy",
		attributeName:   "incident_preference",
		preferenceTypes: []string{"PER_POLICY", "PER_CONDITION", "PER_CONDITION_AND_TARGET"},
	}
}

// Name returns the rule name
func (r *NrAlertPolicyInvalidPreferenceRule) Name() string {
	return "newrelic_alert_policy_invalid_preference"
}

// Enabled returns whether the rule is enabled by default
func (r *NrAlertPolicyInvalidPreferenceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrAlertPolicyInvalidPreferenceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrAlertPolicyInvalidPreferenceRule) Link() string {
	return ""
}

// Check checks whether newrelic_alert_policy has valid incident_preference
func (r *NrAlertPolicyInvalidPreferenceRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(incidentPreference string) error {
			if !slices.Contains(r.preferenceTypes, incidentPreference) {
				runner.EmitIssue(
					r,
					fmt.Sprintf("'%s' is invalid incident preference", incidentPreference),
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
