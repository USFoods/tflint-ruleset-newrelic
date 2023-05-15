package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlertConditionInvalidTypeRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
	name = "My Condition"
	type = "BASIC"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlertConditionInvalidTypeRule(),
					Message: "'BASIC' is invalid condition type",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 9},
						End:      hcl.Pos{Line: 4, Column: 16},
					},
				},
			},
		},
		{
			Name: "no issue found static",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
	name = "My Condition"
	type = "static"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found baseline",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
	name = "My Condition"
	type = "baseline"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlAlertConditionInvalidTypeRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
