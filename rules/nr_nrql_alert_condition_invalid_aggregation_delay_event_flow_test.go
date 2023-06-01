package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlertConditionInvalidAggregationDelayEventFlow(t *testing.T) {
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
  aggregation_method = "event_flow"
  aggregation_delay = 1800
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlertConditionInvalidAggregationDelayEventFlowRule(),
					Message: "'1800' is invalid value for aggregation_delay with aggregation_method 'event_flow'",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 24},
						End:      hcl.Pos{Line: 4, Column: 36},
					},
				},
			},
		},
		{
			Name: "no issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  aggregation_method = "event_flow"
  aggregation_delay = 600
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found null",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  aggregation_method = "event_flow"
  aggregation_delay = null
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlAlertConditionInvalidAggregationDelayEventFlowRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
