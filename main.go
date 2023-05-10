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
			Version: "0.3.0",
			Rules: []tflint.Rule{
				rules.NewNrAlertPolicyInvalidPreferenceRule(),
				rules.NewNrNrqlAlertConditionInvalidTypeRule(),
				rules.NewNrSyntheticsMonitorInvalidPeriodRule(),
				rules.NewNrSyntheticsMonitorInvalidTypeRule(),
				rules.NewNrSyntheticsScriptMonitorInvalidAggregationDelayRule(),
				rules.NewNrSyntheticsScriptMonitorInvalidAggregationMethodRule(),
				rules.NewNrSyntheticsScriptMonitorInvalidAggregationTimerRule(),
				rules.NewNrSyntheticsScriptMonitorInvalidAggregationWindowRule(),
				rules.NewNrSyntheticsScriptMonitorInvalidExpirationDurationRule(),
				rules.NewNrSyntheticsScriptMonitorInvalidPeriodRule(),
				rules.NewNrSyntheticsScriptMonitorInvalidSlidyByRule(),
				rules.NewNrSyntheticsScriptMonitorInvalidTypeRule(),
			},
		},
	})
}
