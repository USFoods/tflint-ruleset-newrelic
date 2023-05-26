package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
	"github.com/zclconf/go-cty/cty"
)

// TODO: Write the rule's description here
// NrNrqlAlertConditionInvalidAggregationDelayEventTimerRule checks ...
type NrNrqlAlertConditionInvalidAggregationDelayEventTimerRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewNrNrqlAlertConditionInvalidAggregationDelayEventTimerRule returns new rule with default attributes
func NewNrNrqlAlertConditionInvalidAggregationDelayEventTimerRule() *NrNrqlAlertConditionInvalidAggregationDelayEventTimerRule {
	return &NrNrqlAlertConditionInvalidAggregationDelayEventTimerRule{
		// TODO: Write resource type and attribute name here
		resourceType: "newrelic_nrql_alert_condition",
	}
}

// Name returns the rule name
func (r *NrNrqlAlertConditionInvalidAggregationDelayEventTimerRule) Name() string {
	return "nr_nrql_alert_condition_invalid_aggregation_delay_event_timer"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlertConditionInvalidAggregationDelayEventTimerRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlertConditionInvalidAggregationDelayEventTimerRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlertConditionInvalidAggregationDelayEventTimerRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *NrNrqlAlertConditionInvalidAggregationDelayEventTimerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "aggregation_method"},
			{Name: "aggregation_delay"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		delayAttr, delayExists := resource.Body.Attributes["aggregation_delay"]
		methodAttr, methodExists := resource.Body.Attributes["aggregation_method"]

		if !delayExists || !methodExists {
			continue
		}

		var aggregationDelay cty.Value
		if err := runner.EvaluateExpr(delayAttr.Expr, &aggregationDelay, nil); err != nil {
			return err
		}

		if aggregationDelay.IsNull() {
			continue
		}

		var aggregationMethod string
		if err := runner.EvaluateExpr(methodAttr.Expr, &aggregationMethod, nil); err != nil {
			return err
		}

		if aggregationMethod == "event_timer" {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("aggregation_delay invalid attribute with aggregation_method '%s'", aggregationMethod),
				methodAttr.Expr.Range(),
			)
		}
	}

	return nil
}
