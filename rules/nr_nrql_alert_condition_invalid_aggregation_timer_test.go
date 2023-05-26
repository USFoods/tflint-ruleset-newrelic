package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlerConditionInvalidAggregationTimerRule(t *testing.T) {
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
					Rule:    NewNrNrqlAlerConditionInvalidAggregationTimerRule(),
					Message: "'3601' is invalid aggregation_timer",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 24},
						End:      hcl.Pos{Line: 4, Column: 37},
					},
				},
			},
		},
		{
			Name: "issue found event_flow",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  aggregation_method = "event_flow"
  aggregation_timer = 60
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlerConditionInvalidAggregationTimerRule(),
					Message: "aggregation_timer invalid for aggregation_method 'event_flow'",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 24},
						End:      hcl.Pos{Line: 4, Column: 36},
					},
				},
			},
		},
		{
			Name: "issue found cadence",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  aggregation_method = "cadence"
  aggregation_timer = 60
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlerConditionInvalidAggregationTimerRule(),
					Message: "aggregation_timer invalid for aggregation_method 'cadence'",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 24},
						End:      hcl.Pos{Line: 4, Column: 33},
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

	rule := NewNrNrqlAlerConditionInvalidAggregationTimerRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
