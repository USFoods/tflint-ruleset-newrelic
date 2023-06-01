package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/usfoods/tflint-ruleset-newrelic/newrelic"
	"github.com/usfoods/tflint-ruleset-newrelic/project"
	"github.com/usfoods/tflint-ruleset-newrelic/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &newrelic.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "newrelic",
				Version: project.Version,
			},
			PresetRules: rules.PresetRules,
		},
	})
}
