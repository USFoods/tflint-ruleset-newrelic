package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlerConditionInvalidExpirationRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found max",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  expiration_duration = 172801
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlerConditionInvalidExpirationDurationRule(),
					Message: "'172801' is invalid value for expiration_duration",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 25},
						End:      hcl.Pos{Line: 4, Column: 31},
					},
				},
			},
		},
		{
			Name: "issue found min",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  expiration_duration = 10
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlerConditionInvalidExpirationDurationRule(),
					Message: "'10' is invalid value for expiration_duration",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 25},
						End:      hcl.Pos{Line: 4, Column: 27},
					},
				},
			},
		},
		{
			Name: "no issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  expiration_duration = 900
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  expiration_duration = null
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlAlerConditionInvalidExpirationDurationRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
