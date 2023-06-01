package rules

import (
	"fmt"
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlertConditionInvalidEvaluationDelay(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  evaluation_delay = 1800
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlertConditionInvalidEvaluationDelayRule(),
					Message: fmt.Sprintf("'1800' is an invalid value for evaluation_delay"),
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 22},
						End:      hcl.Pos{Line: 4, Column: 26},
					},
				},
			},
		},
		{
			Name: "no issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  evaluation_delay = 600
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found null",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  evaluation_delay = null
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlAlertConditionInvalidEvaluationDelayRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
