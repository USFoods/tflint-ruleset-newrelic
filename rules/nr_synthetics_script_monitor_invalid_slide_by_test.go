package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrSyntheticsScriptMonitorInvalidSlideByRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found greater than",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				aggregation_window = 60
				slide_by = 120
			}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsScriptMonitorInvalidSlidyByRule(),
					Message: "slide_by is greater than aggregation_window",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 16},
						End:      hcl.Pos{Line: 5, Column: 19},
					},
				},
			},
		},
		{
			Name: "issue found not a factor",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				aggregation_window = 120
				slide_by = 45
			}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsScriptMonitorInvalidSlidyByRule(),
					Message: "slide_by is not a factor of aggregation_window",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 16},
						End:      hcl.Pos{Line: 5, Column: 18},
					},
				},
			},
		},
		{
			Name: "no issue found missing aggregation_window",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				slide_by = 60
			}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found missing slide_by",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				aggregation_window = 60
			}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no issue found",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				aggregation_window = 60
				slide_by = 30
			}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrSyntheticsScriptMonitorInvalidSlidyByRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
