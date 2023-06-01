package rules

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

// NrNrqlAlerConditionInvalidAggregationTimerMethodRule checks whether newrelic_nrql_alert_condition has valid aggregation timer
type NrNrqlAlerConditionInvalidAggregationTimerMethodRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewNrNrqlAlerConditionInvalidAggregationTimerMethodRule returns a new rule
func NewNrNrqlAlerConditionInvalidAggregationTimerMethodRule() *NrNrqlAlerConditionInvalidAggregationTimerMethodRule {
	return &NrNrqlAlerConditionInvalidAggregationTimerMethodRule{
		resourceType: "newrelic_nrql_alert_condition",
	}
}

// Name returns the rule name
func (r *NrNrqlAlerConditionInvalidAggregationTimerMethodRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_aggregation_timer_method"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlerConditionInvalidAggregationTimerMethodRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlerConditionInvalidAggregationTimerMethodRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlerConditionInvalidAggregationTimerMethodRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_nrql_alert_condition has valid aggregation timer
func (r *NrNrqlAlerConditionInvalidAggregationTimerMethodRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "aggregation_method"},
			{Name: "aggregation_timer"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		timerAttr, timerExists := resource.Body.Attributes["aggregation_timer"]
		methodAttr, methodExists := resource.Body.Attributes["aggregation_method"]

		if !timerExists {
			continue
		}

		var timerCty cty.Value
		if err := runner.EvaluateExpr(timerAttr.Expr, &timerCty, nil); err != nil {
			return err
		}

		if timerCty.IsNull() || !timerCty.IsKnown() {
			continue
		}

		// Default aggregation_method is event_flow
		var aggregationMethod string = "event_flow"
		// Default range is empty
		var methodRange = hcl.Range{
			Filename: timerAttr.Expr.Range().Filename,
			Start:    hcl.Pos{Line: 0, Column: 0, Byte: 0},
			End:      hcl.Pos{Line: 0, Column: 0, Byte: 0},
		}

		if methodExists {
			var methodCty cty.Value
			if err := runner.EvaluateExpr(methodAttr.Expr, &methodCty, nil); err != nil {
				return err
			}

			if !methodCty.IsNull() || methodCty.IsKnown() {
				if err := gocty.FromCtyValue(methodCty, &aggregationMethod); err != nil {
					return err
				}

				methodRange = methodAttr.Expr.Range()
			}
		}

		if aggregationMethod != "event_timer" {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("aggregation_timer is invalid attribute for aggregation_method '%s'", aggregationMethod),
				methodRange,
			)
		}
	}

	return nil
}
