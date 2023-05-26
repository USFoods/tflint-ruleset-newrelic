package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrSyntheticsScriptMonitorInvalidAggregationDelayRule(t *testing.T) {
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
  aggregation_delay = 60
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsScriptMonitorInvalidAggregationDelayRule(),
					Message: "aggregation_delay invalid for aggregation_method 'event_timer'",
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
resource "newrelic_synthetics_script_monitor" "monitor" {
  name = "My Monitor"
  aggregation_method = "event_flow"
  aggregation_delay = 1500
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsScriptMonitorInvalidAggregationDelayRule(),
					Message: "'1500' invalid aggregation_delay for aggregation_method 'event_flow'",
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
resource "newrelic_synthetics_script_monitor" "monitor" {
  name = "My Monitor"
  aggregation_method = "cadence"
  aggregation_delay = 3900
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsScriptMonitorInvalidAggregationDelayRule(),
					Message: "'3900' invalid aggregation_delay for aggregation_method 'cadence'",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 24},
						End:      hcl.Pos{Line: 4, Column: 33},
					},
				},
			},
		},
		{
			Name: "no issue found event_flow",
			Content: `
resource "newrelic_synthetics_script_monitor" "monitor" {
  name = "My Monitor"
  aggregation_method = "event_flow"
  aggregation_delay = 600
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found cadence",
			Content: `
resource "newrelic_synthetics_script_monitor" "monitor" {
  name = "My Monitor"
  aggregation_method = "cadence"
  aggregation_delay = 3600
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found aggregation_delay not set",
			Content: `
resource "newrelic_synthetics_script_monitor" "monitor" {
  name = "My Monitor"
  aggregation_method = "event_timer"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrSyntheticsScriptMonitorInvalidAggregationDelayRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
