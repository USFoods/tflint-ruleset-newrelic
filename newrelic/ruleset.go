package newrelic

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// RuleSet is the custom ruleset for the New Relic provider
type RuleSet struct {
	tflint.BuiltinRuleSet

	PresetRules map[string][]tflint.Rule
}

func (r *RuleSet) Check(rr tflint.Runner) error {
	runner := NewRunner(rr)

	for _, rule := range r.Rules {
		if err := rule.Check(runner); err != nil {
			return fmt.Errorf("failed to check `%s` rule: %s", rule.Name(), err)
		}
	}

	return nil
}
