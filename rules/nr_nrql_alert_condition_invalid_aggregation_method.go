package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"

	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// NrNrqlAlerConditionInvalidAggregationMethodRule checks whether newrelic_nrql_alert_condition has valid aggregation_method
type NrNrqlAlerConditionInvalidAggregationMethodRule struct {
	tflint.DefaultRule

	resourceType       string
	attributeName      string
	aggregationMethods map[string]bool
}

// NewNrNrqlAlerConditionInvalidAggregationMethodRule returns a new rule
func NewNrNrqlAlerConditionInvalidAggregationMethodRule() *NrNrqlAlerConditionInvalidAggregationMethodRule {
	return &NrNrqlAlerConditionInvalidAggregationMethodRule{
		resourceType:  "newrelic_nrql_alert_condition",
		attributeName: "aggregation_method",
		aggregationMethods: map[string]bool{
			"cadence":     true,
			"event_flow":  true,
			"event_timer": true,
		},
	}
}

// Name returns the rule name
func (r *NrNrqlAlerConditionInvalidAggregationMethodRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_aggregation_method"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlerConditionInvalidAggregationMethodRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlerConditionInvalidAggregationMethodRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlerConditionInvalidAggregationMethodRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_nrql_alert_condition has valid aggregation_method
func (r *NrNrqlAlerConditionInvalidAggregationMethodRule) Check(runner tflint.Runner) error {
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

		var methodCty cty.Value
		if err := runner.EvaluateExpr(attribute.Expr, &methodCty, nil); err != nil {
			return err
		}

		if methodCty.IsNull() || !methodCty.IsWhollyKnown() {
			continue
		}

		var aggregationMethod string
		if err := gocty.FromCtyValue(methodCty, &aggregationMethod); err != nil {
			return err
		}

		if !r.aggregationMethods[aggregationMethod] {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("'%s' is invalid aggregation method", aggregationMethod),
				attribute.Expr.Range(),
			)
		}
	}

	return nil
}
