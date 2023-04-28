package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrAlertPolicyInvalidPreferenceRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "newrelic_alert_policy" "policy" {
	incident_preference = "PER_ISSUE"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrAlertPolicyInvalidPreferenceRule(),
					Message: "'PER_ISSUE' is invalid incident preference",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 24},
						End:      hcl.Pos{Line: 3, Column: 35},
					},
				},
			},
		},
		{
			Name: "no issue found",
			Content: `
resource "newrelic_alert_policy" "policy" {
	incident_preference = "PER_POLICY"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrAlertPolicyInvalidPreferenceRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
