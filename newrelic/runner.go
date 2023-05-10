package newrelic

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// Runner is a custom runner that provides helper functions for this ruleset.
type Runner struct {
	tflint.Runner
}

// NewRunner returns a new custom runner.
func NewRunner(runner tflint.Runner) *Runner {
	return &Runner{Runner: runner}
}
