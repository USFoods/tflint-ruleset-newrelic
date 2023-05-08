package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrSyntheticsScriptMonitorInvalidAggregationWindowRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				type = "SCRIPT_API"
				aggregation_window = 10
			}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsScriptMonitorInvalidAggregationWindowRule(),
					Message: "'10' is invalid aggregation_window",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 26},
						End:      hcl.Pos{Line: 5, Column: 28},
					},
				},
			},
		},
		{
			Name: "no issue found",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				type = "SCRIPT_API"
				aggregation_window = 60
			}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrSyntheticsScriptMonitorInvalidAggregationWindowRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
