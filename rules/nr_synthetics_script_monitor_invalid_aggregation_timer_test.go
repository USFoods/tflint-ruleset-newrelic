package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrSyntheticsScriptMonitorInvalidAggregationTimerRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found event_timer",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				aggregation_method = "event_timer"
				aggregation_timer = 3601
			}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsScriptMonitorInvalidAggregationTimerRule(),
					Message: "'3601' is invalid aggregation_timer",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 25},
						End:      hcl.Pos{Line: 5, Column: 29},
					},
				},
			},
		},
		{
			Name: "issue found event_flow",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				aggregation_method = "event_flow"
				aggregation_timer = 60
			}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsScriptMonitorInvalidAggregationTimerRule(),
					Message: "aggregation_timer invalid for aggregation_method 'event_flow'",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 25},
						End:      hcl.Pos{Line: 5, Column: 27},
					},
				},
			},
		},
		{
			Name: "issue found cadence",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				aggregation_method = "cadence"
				aggregation_timer = 60
			}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsScriptMonitorInvalidAggregationTimerRule(),
					Message: "aggregation_timer invalid for aggregation_method 'cadence'",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 25},
						End:      hcl.Pos{Line: 5, Column: 27},
					},
				},
			},
		},
	}

	rule := NewNrSyntheticsScriptMonitorInvalidAggregationTimerRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
