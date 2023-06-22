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
		NewNrNrqlAlertConditionInvalidBaselineCriticalOperatorRule(),
		NewNrNrqlAlertConditionInvalidBaselineDirectionRule(),
		NewNrNrqlAlertConditionInvalidEvaluationDelayRule(),
		NewNrNrqlAlerConditionInvalidExpirationDurationRule(),
		NewNrNrqlAlertConditionInvalidFillOptionRule(),
		NewNrNrqlAlerConditionInvalidSlidyByRule(),
		NewNrNrqlAlertConditionInvalidStaticCriticalOperatorRule(),
		NewNrNrqlAlertConditionInvalidTypeRule(),
		NewNrNrqlAlertConditionInvalidViolationTimeLimitSecondsRule(),
		NewNrNrqlDropRuleInvalidActionRule(),
		NewNrSyntheticsMonitorInvalidPeriodRule(),
		NewNrSyntheticsMonitorInvalidTypeRule(),
		NewNrSyntheticsScriptMonitorInvalidPeriodRule(),
		NewNrSyntheticsScriptMonitorInvalidTypeRule(),
	},
}
