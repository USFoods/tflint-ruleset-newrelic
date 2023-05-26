package rules

import (
	"fmt"
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestNrAlertMutingRuleInvalidConditionsAttribute(t *testing.T) {
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
      attribute = "type"
	  operator = "EQUALS"
	  values = ["APM"]
	}
	operator = "AND"
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewNrAlertMutingRuleInvalidConditionsAttributeRule(),
					Message: "'type' is invalid value for conditions attribute",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 7, Column: 19},
						End:      hcl.Pos{Line: 7, Column: 25},
					},
				},
			},
		},
	}

	attributeTypes := []string{
		"accountId",
		"conditionId",
		"conditionName",
		"conditionRunbookUrl",
		"conditionType",
		"entity.guid",
		"nrqlEventType",
		"nrqlQuery",
		"policyId",
		"policyName",
		"product",
		"targetId",
		"targetName",
	}

	for _, attribute := range attributeTypes {
		cases = append(cases, struct {
			Name     string
			Content  string
			Expected helper.Issues
		}{
			Name: fmt.Sprintf("no issue found '%s'", attribute),
			Content: `
resource "newrelic_alert_muting_rule" "muting_rule" {
  name = "My Muting Rule"
  enabled = true
  condition {
    conditions {
      attribute = "` + attribute + `"
      operator = "EQUALS"
      values = ["APM"]
    }
  }
}`,
			Expected: helper.Issues{},
		})
	}

	rule := NewNrAlertMutingRuleInvalidConditionsAttributeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
