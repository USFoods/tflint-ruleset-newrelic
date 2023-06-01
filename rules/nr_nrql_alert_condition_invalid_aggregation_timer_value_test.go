package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlerConditionInvalidAggregationTimerValueRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found event_timer",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  aggregation_method = "event_timer"
  aggregation_timer = 3601
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlerConditionInvalidAggregationTimerValueRule(),
					Message: "'3601' is invalid value for aggregation_timer",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 24},
						End:      hcl.Pos{Line: 4, Column: 37},
					},
				},
			},
		},
		{
			Name: "no issue found valid aggregation_timer",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  aggregation_method = "event_timer"
  aggregation_timer = 3600
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found aggregation_timer not set",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  aggregation_method = "event_flow"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlAlerConditionInvalidAggregationTimerValueRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
