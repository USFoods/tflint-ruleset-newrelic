package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrSyntheticsMonitorInvalidAttributesRule(t *testing.T) {
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
	type = "BROWSER"
	bypass_head_request = true
	treat_redirect_as_failure = true
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrSyntheticsMonitorInvalidAttributesRule(),
					Message: "'bypass_head_request' is invalid attribute for 'BROWSER' monitor",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 3},
						End:      hcl.Pos{Line: 5, Column: 23},
					},
				},
				{
					Rule:    NewNrSyntheticsMonitorInvalidAttributesRule(),
					Message: "'treat_redirect_as_failure' is invalid attribute for 'BROWSER' monitor",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 6, Column: 3},
						End:      hcl.Pos{Line: 6, Column: 31},
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
	bypass_head_request = true
	treat_redirect_as_failure = true
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrSyntheticsMonitorInvalidAttributesRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
