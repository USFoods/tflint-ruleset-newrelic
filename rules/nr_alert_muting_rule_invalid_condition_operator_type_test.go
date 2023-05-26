package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrAlertMutingRuleInvalidConditionOperatorType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "newrelic_alert_muting_rule" "muting_rule" {
	  name = "My Muting Rule"
	  enabled = true
	  condition {
		conditions {
			attribute = "product"
			operator = "ALL"
			values = ["APM"]
		}
		operator = "AND"
	  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewNrAlertMutingRuleInvalidConditionOperatorTypeRule(),
					Message: "'ALL' is invalid condition operator type",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 8, Column: 15},
						End:      hcl.Pos{Line: 8, Column: 20},
					},
				},
			},
		},
	}

	rule := NewNrAlertMutingRuleInvalidConditionOperatorTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
