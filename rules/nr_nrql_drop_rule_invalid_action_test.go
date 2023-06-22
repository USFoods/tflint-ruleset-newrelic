package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlDropRuleInvalidAction(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "newrelic_nrql_drop_rule" "rule" {
	account_id = "123456"
	description = "Drop rule"
	action = "discard"
	nrql = "SELECT * FROM Log"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlDropRuleInvalidActionRule(),
					Message: "'discard' is invalid value for action",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 11},
						End:      hcl.Pos{Line: 5, Column: 20},
					},
				},
			},
		},
		{
			Name: "no issue found drop_data",
			Content: `
resource "newrelic_nrql_drop_rule" "rule" {
	account_id = "123456"
	description = "Drop rule"
	action = "drop_data"
	nrql = "SELECT * FROM Log"
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found drop_attributes",
			Content: `
resource "newrelic_nrql_drop_rule" "rule" {
	account_id = "123456"
	description = "Drop rule"
	action = "drop_attributes"
	nrql = "SELECT * FROM Log"
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found drop_attributes_from_metric_aggregates",
			Content: `
resource "newrelic_nrql_drop_rule" "rule" {
	account_id = "123456"
	description = "Drop rule"
	action = "drop_attributes_from_metric_aggregates"
	nrql = "SELECT * FROM Log"
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlDropRuleInvalidActionRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
