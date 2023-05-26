package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrNrqlAlerConditionInvalidAggregationWindowRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  type = "SCRIPT_API"
  aggregation_window = 10
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrNrqlAlerConditionInvalidAggregationWindowRule(),
					Message: "'10' is invalid aggregation_window",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 24},
						End:      hcl.Pos{Line: 5, Column: 26},
					},
				},
			},
		},
		{
			Name: "no issue found",
			Content: `
resource "newrelic_nrql_alert_condition" "monitor" {
  name = "My Monitor"
  type = "SCRIPT_API"
  aggregation_window = 60
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewNrNrqlAlerConditionInvalidAggregationWindowRule()

	for _, test := range tests {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, test.Expected, runner.Issues)
	}
}
