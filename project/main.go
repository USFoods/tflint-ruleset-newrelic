package project

import "fmt"

// Version is ruleset version
const Version string = "0.3.2"

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/usfoods/tflint-ruleset-newrelic/blob/v%s/docs/rules/%s.md", Version, name)
}
