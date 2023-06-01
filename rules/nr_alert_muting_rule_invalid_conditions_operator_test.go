package rules

import (
	"fmt"
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrAlertMutingRuleInvalidConditionsOperator(t *testing.T) {
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
					Rule:    NewNrAlertMutingRuleInvalidConditionsOperatorRule(),
					Message: "'ALL' is invalid value for conditions operator",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 8, Column: 18},
						End:      hcl.Pos{Line: 8, Column: 23},
					},
				},
			},
		},
	}

	operatorTypes := []string{
		"ANY",
		"CONTAINS",
		"ENDS_WITH",
		"EQUALS",
		"IN",
		"IS_BLANK",
		"IS_NOT_BLANK",
		"NOT_CONTAINS",
		"NOT_ENDS_WITH",
		"NOT_EQUALS",
		"NOT_IN",
		"NOT_STARTS_WITH",
		"STARTS_WITH",
	}

	for _, operator := range operatorTypes {
		cases = append(cases, struct {
			Name     string
			Content  string
			Expected helper.Issues
		}{
			Name: fmt.Sprintf("no issue found '%s'", operator),
			Content: `
resource "newrelic_alert_muting_rule" "muting_rule" {
  name = "My Muting Rule"
  enabled = true
  condition {
    conditions {
      attribute = "product"
      operator = "` + operator + `"
      values = ["APM"]
    }
  }
}`,
			Expected: helper.Issues{},
		})
	}

	rule := NewNrAlertMutingRuleInvalidConditionsOperatorRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
