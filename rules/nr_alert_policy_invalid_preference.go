package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

type NrAlertPolicyInvalidPreferenceRule struct {
	tflint.DefaultRule

	resourceType    string
	attributeName   string
	preferenceTypes map[string]bool
}

// NewNrAlertPolicyInvalidPreferenceRule returns a new rule
func NewNrAlertPolicyInvalidPreferenceRule() *NrAlertPolicyInvalidPreferenceRule {
	return &NrAlertPolicyInvalidPreferenceRule{
		resourceType:  "newrelic_alert_policy",
		attributeName: "incident_preference",
		preferenceTypes: map[string]bool{
			"PER_POLICY":               true,
			"PER_CONDITION":            true,
			"PER_CONDITION_AND_TARGET": true,
		},
	}
}

// Name returns the rule name
func (r *NrAlertPolicyInvalidPreferenceRule) Name() string {
	return "nr_alert_policy_invalid_preference"
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
	return project.ReferenceLink(r.Name())
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
			if !r.preferenceTypes[incidentPreference] {
				return runner.EmitIssue(
					r,
					fmt.Sprintf("'%s' is invalid value for incident_preference", incidentPreference),
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
