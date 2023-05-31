package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

// NrNrqlAlerConditionInvalidAggregationWindowRule checks whether newrelic_nrql_alert_condition has valid aggregation_window
type NrNrqlAlerConditionInvalidAggregationWindowRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	min           int
	max           int
}

// NewNrNrqlAlerConditionInvalidAggregationWindowRule returns a new rule
func NewNrNrqlAlerConditionInvalidAggregationWindowRule() *NrNrqlAlerConditionInvalidAggregationWindowRule {
	return &NrNrqlAlerConditionInvalidAggregationWindowRule{
		resourceType:  "newrelic_nrql_alert_condition",
		attributeName: "aggregation_window",
		min:           30,
		max:           900,
	}
}

// Name returns the rule name
func (r *NrNrqlAlerConditionInvalidAggregationWindowRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_aggregation_window"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlerConditionInvalidAggregationWindowRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlerConditionInvalidAggregationWindowRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlerConditionInvalidAggregationWindowRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_nrql_alert_condition has valid aggregation_window
func (r *NrNrqlAlerConditionInvalidAggregationWindowRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]

		if !exists {
			continue
		}

		var attrCty cty.Value
		if err := runner.EvaluateExpr(attribute.Expr, &attrCty, nil); err != nil {
			return err
		}

		if attrCty.IsNull() || !attrCty.IsKnown() {
			continue
		}

		var aggregationWindow int
		if err := gocty.FromCtyValue(attrCty, &aggregationWindow); err != nil {
			return err
		}

		if aggregationWindow < r.min || aggregationWindow > r.max {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("'%d' is invalid value for aggregation_window", aggregationWindow),
				attribute.Expr.Range(),
			)
		}
	}

	return nil
}
