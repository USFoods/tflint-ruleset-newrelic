package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
)

// NrNrqlAlerConditionInvalidAggregationDelayRule checks whether newrelic_synthetics_monitor has valid aggregation_delay
type NrNrqlAlerConditionInvalidAggregationDelayRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewNrNrqlAlerConditionInvalidAggregationDelayRule returns a new rule
func NewNrNrqlAlerConditionInvalidAggregationDelayRule() *NrNrqlAlerConditionInvalidAggregationDelayRule {
	return &NrNrqlAlerConditionInvalidAggregationDelayRule{
		resourceType: "newrelic_nrql_alert_condition",
	}
}

// Name returns the rule name
func (r *NrNrqlAlerConditionInvalidAggregationDelayRule) Name() string {
	return "nr_synthetics_script_monitor_invalid_aggregation_delay"
}

// Enabled returns whether the rule is enabled by default
func (r *NrNrqlAlerConditionInvalidAggregationDelayRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NrNrqlAlerConditionInvalidAggregationDelayRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NrNrqlAlerConditionInvalidAggregationDelayRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether newrelic_nrql_alert_condition has valid aggregation_delay
func (r *NrNrqlAlerConditionInvalidAggregationDelayRule) Check(runner tflint.Runner) error {
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
		attr, ok := resource.Body.Attributes["aggregation_delay"]

		if !ok {
			continue
		}

		var aggregationDelay int
		err := runner.EvaluateExpr(attr.Expr, &aggregationDelay, nil)

		if err != nil {
			return err
		}

		attr, ok = resource.Body.Attributes["aggregation_method"]

		if !ok {
			continue
		}

		var aggregationMethod string
		err = runner.EvaluateExpr(attr.Expr, &aggregationMethod, nil)

		if err != nil {
			return err
		}

		if aggregationMethod == "event_timer" {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("aggregation_delay invalid attribute with aggregation_method '%s'", aggregationMethod),
				attr.Expr.Range(),
			)
		}

		if aggregationMethod == "event_flow" {
			if aggregationDelay > 1200 {
				return runner.EmitIssue(
					r,
					fmt.Sprintf("'%d' invalid value for aggregation_delay with aggregation_method '%s'", aggregationDelay, aggregationMethod),
					attr.Expr.Range(),
				)
			}
		}

		if aggregationMethod == "cadence" {
			if aggregationDelay > 3600 {
				return runner.EmitIssue(
					r,
					fmt.Sprintf("'%d' invalid value for aggregation_delay with aggregation_method '%s'", aggregationDelay, aggregationMethod),
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
