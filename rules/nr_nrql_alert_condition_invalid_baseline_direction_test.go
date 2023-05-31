package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlertConditionInvalidBaselineDirection(t *testing.T) {
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
  type = "baseline"
  baseline_direction = "JUST_UPPER"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlertConditionInvalidBaselineDirectionRule(),
					Message: "'JUST_UPPER' is invalid value for baseline_direction",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 24},
						End:      hcl.Pos{Line: 5, Column: 36},
					},
				},
			},
		},
		{
			Name: "no issue found upper_only",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
  name = "My NRQL Alert Condition"
  type = "baseline"
  baseline_direction = "upper_only"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found lower_only",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
  name = "My NRQL Alert Condition"
  type = "baseline"
  baseline_direction = "lower_only"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found upper_and_lower",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
  name = "My NRQL Alert Condition"
  type = "baseline"
  baseline_direction = "upper_and_lower"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found null",
			Content: `
resource "newrelic_nrql_alert_condition" "condition" {
  name = "My NRQL Alert Condition"
  type = "baseline"
  baseline_direction = null
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlAlertConditionInvalidBaselineDirectionRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
