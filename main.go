package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/usfoods/tflint-ruleset-newrelic/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "newrelic",
			Version: "0.4.0",
			Rules: []tflint.Rule{
				rules.NewNrAlertMutingRuleInvalidConditionsAttributeRule(),
				rules.NewNrAlertMutingRuleInvalidConditionsOperatorRule(),
				rules.NewNrAlertPolicyInvalidPreferenceRule(),
				rules.NewNrNrqlAlertConditionInvalidAggregationDelayCadenceRule(),
				rules.NewNrNrqlAlertConditionInvalidAggregationDelayEventFlowRule(),
				rules.NewNrNrqlAlertConditionInvalidAggregationDelayEventTimerRule(),
				rules.NewNrNrqlAlerConditionInvalidAggregationMethodRule(),
				rules.NewNrNrqlAlerConditionInvalidAggregationTimerMethodRule(),
				rules.NewNrNrqlAlerConditionInvalidAggregationTimerValueRule(),
				rules.NewNrNrqlAlerConditionInvalidAggregationWindowRule(),
				rules.NewNrNrqlAlertConditionInvalidBaselineDirectionRule(),
				rules.NewNrNrqlAlerConditionInvalidExpirationDurationRule(),
				rules.NewNrNrqlAlerConditionInvalidSlidyByRule(),
				rules.NewNrNrqlAlertConditionInvalidTypeRule(),
				rules.NewNrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule(),
				rules.NewNrSyntheticsMonitorInvalidPeriodRule(),
				rules.NewNrSyntheticsMonitorInvalidTypeRule(),
				rules.NewNrSyntheticsScriptMonitorInvalidTypeRule(),
			},
		},
	})
}
