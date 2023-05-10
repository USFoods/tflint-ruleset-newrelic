package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrSyntheticsScriptMonitorInvalidExpirationRule(t *testing.T) {
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
				expiration_duration = 0
			}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsScriptMonitorInvalidExpirationDurationRule(),
					Message: "'0' is invalid expiration_duration",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 27},
						End:      hcl.Pos{Line: 4, Column: 28},
					},
				},
			},
		},
		{
			Name: "no issue found",
			Content: `
			resource "newrelic_synthetics_script_monitor" "monitor" {
				name = "My Monitor"
				expiration_duration = 900
			}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrSyntheticsScriptMonitorInvalidExpirationDurationRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
