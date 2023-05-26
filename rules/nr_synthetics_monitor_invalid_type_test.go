package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrSyntheticsMonitorInvalidTypeRule(t *testing.T) {
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
  type = "BASIC"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsMonitorInvalidTypeRule(),
					Message: "'BASIC' is invalid monitor type",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 10},
						End:      hcl.Pos{Line: 4, Column: 17},
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
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrSyntheticsMonitorInvalidTypeRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
