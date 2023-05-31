package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

// TODO: Write the rule's description here
// NrNrqlAlertConditionInvalidAggregationDelayEventFlowRule checks ...
type NrNrqlAlertConditionInvalidAggregationDelayEventFlowRule struct {
	tflint.DefaultRule

	resourceType string
	max          int
}

// NewNrNrqlAlertConditionInvalidAggregationDelayEventFlowRule returns new rule with default attributes
func NewNrNrqlAlertConditionInvalidAggregationDelayEventFlowRule() *NrNrqlAlertConditionInvalidAggregationDelayEventFlowRule {
	return &NrNrqlAlertConditionInvalidAggregationDelayEventFlowRule{
		// TODO: Write resource type and attribute name here
		resourceType: "newrelic_nrql_alert_condition",
		max:          1200,
	}
}

// Name returns the rule name
func (r *NrNrqlAlertConditionInvalidAggregationDelayEventFlowRule) Name() string {
	return "nr_nrql_alert_condition_invalid_aggregation_delay_event_flow"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlertConditionInvalidAggregationDelayEventFlowRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlertConditionInvalidAggregationDelayEventFlowRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlertConditionInvalidAggregationDelayEventFlowRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *NrNrqlAlertConditionInvalidAggregationDelayEventFlowRule) Check(runner tflint.Runner) error {
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

		var delayCty cty.Value
		if err := runner.EvaluateExpr(delayAttr.Expr, &delayCty, nil); err != nil {
			return err
		}

		if delayCty.IsNull() || !delayCty.IsKnown() {
			continue
		}

		var aggregationDelay int
		if err := gocty.FromCtyValue(delayCty, &aggregationDelay); err != nil {
			return err
		}

		var aggregationMethod string
		if err := runner.EvaluateExpr(methodAttr.Expr, &aggregationMethod, nil); err != nil {
			return err
		}

		if aggregationMethod == "event_flow" {
			if aggregationDelay > r.max {
				return runner.EmitIssue(
					r,
					fmt.Sprintf("'%d' is invalid value for aggregation_delay with aggregation_method '%s'", aggregationDelay, aggregationMethod),
					methodAttr.Expr.Range(),
				)
			}
		}

	}

	return nil
}
