package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrSyntheticsMonitorInvalidPeriodRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "newrelic_synthetics_monitor" "monitor" {
	name = "My Monitor"
	type = "SIMPLE"
	period = 300
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsMonitorInvalidPeriodRule(),
					Message: "'300' is invalid period",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 11},
						End:      hcl.Pos{Line: 5, Column: 14},
					},
				},
			},
		},
		{
			Name: "no issue found",
			Content: `
resource "newrelic_synthetics_monitor" "monitor" {
	  name = "My Monitor"
	  type = "SIMPLE"
	  period = "EVERY_5_MINUTES"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrSyntheticsMonitorInvalidPeriodRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
