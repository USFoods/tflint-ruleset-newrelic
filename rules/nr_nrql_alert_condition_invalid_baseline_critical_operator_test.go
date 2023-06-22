package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlertConditionInvalidBaselineCriticalOperator(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "alert_conditon" {
	name = "My Alert Condition"
	enabled = true
	type = "baseline"
	critical {
		operator = "equals"
		threshold = 0.5
		threshold_duration = 900
		threshold_occurences = "ALL"
	}
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlertConditionInvalidBaselineCriticalOperatorRule(),
					Message: "'equals' is an invalid value for critical operator with type 'baseline'",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 7, Column: 14},
						End:      hcl.Pos{Line: 7, Column: 22},
					},
				},
			},
		},
		{
			Name: "no issue found",
			Content: `
			resource "newrelic_nrql_alert_condition" "alert_conditon" {
				name = "My Alert Condition"
				enabled = true
				type = "static"
				critical {
					operator = "above"
					threshold = 0.5
					threshold_duration = 900
					threshold_occurences = "ALL"
				}
			}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlAlertConditionInvalidBaselineCriticalOperatorRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
