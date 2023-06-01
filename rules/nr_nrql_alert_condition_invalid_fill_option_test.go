package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlertConditionInvalidFillOption(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
  name = "My NRQL Alert Condition"
  type = "static"
  fill_option = "TODO"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlertConditionInvalidFillOptionRule(),
					Message: "'TODO' is an invalid value for fill_option",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 17},
						End:      hcl.Pos{Line: 5, Column: 23},
					},
				},
			},
		},
		{
			Name: "no issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
  name = "My NRQL Alert Condition"
  type = "static"
  fill_option = "none"
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found null",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
  name = "My NRQL Alert Condition"
  type = "static"
  fill_option = null
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlAlertConditionInvalidFillOptionRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
