package rules

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

var PresetRules = map[string][]tflint.Rule{
	"all": {
		NewNrAlertMutingRuleInvalidConditionsAttributeRule(),
		NewNrAlertMutingRuleInvalidConditionsOperatorRule(),
		NewNrAlertPolicyInvalidPreferenceRule(),
		NewNrNrqlAlertConditionInvalidAggregationDelayCadenceRule(),
		NewNrNrqlAlertConditionInvalidAggregationDelayEventFlowRule(),
		NewNrNrqlAlertConditionInvalidAggregationDelayEventTimerRule(),
		NewNrNrqlAlerConditionInvalidAggregationMethodRule(),
		NewNrNrqlAlerConditionInvalidAggregationTimerMethodRule(),
		NewNrNrqlAlerConditionInvalidAggregationTimerValueRule(),
		NewNrNrqlAlerConditionInvalidAggregationWindowRule(),
		NewNrNrqlAlertConditionInvalidBaselineDirectionRule(),
		NewNrNrqlAlertConditionInvalidEvaluationDelayRule(),
		NewNrNrqlAlerConditionInvalidExpirationDurationRule(),
		NewNrNrqlAlertConditionInvalidFillOptionRule(),
		NewNrNrqlAlerConditionInvalidSlidyByRule(),
		NewNrNrqlAlertConditionInvalidTypeRule(),
		NewNrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule(),
		NewNrSyntheticsMonitorInvalidPeriodRule(),
		NewNrSyntheticsMonitorInvalidTypeRule(),
		NewNrSyntheticsScriptMonitorInvalidTypeRule(),
	},
}
