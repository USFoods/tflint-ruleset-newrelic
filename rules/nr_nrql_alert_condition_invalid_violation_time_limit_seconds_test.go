package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlertConditionInvalidViolationTimeLimitSeconds(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found min",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
  name = "My NRQL Alert Condition"
  type = "static"
  violation_time_limit_seconds = 120
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule(),
					Message: "'120' is an invalid value for violation_time_limit_seconds",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 34},
						End:      hcl.Pos{Line: 5, Column: 37},
					},
				},
			},
		},
		{
			Name: "issue found max",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
  name = "My NRQL Alert Condition"
  type = "static"
  violation_time_limit_seconds = 2592001
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule(),
					Message: "'2592001' is an invalid value for violation_time_limit_seconds",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 34},
						End:      hcl.Pos{Line: 5, Column: 41},
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
  violation_time_limit_seconds = 1800
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
