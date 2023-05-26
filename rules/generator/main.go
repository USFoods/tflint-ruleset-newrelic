package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var heading = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

type metadata struct {
	RuleName   string
	RuleNameCC string
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("Rule name? (e.g. aws_instance_invalid_type): ")
	ruleName, err := buf.ReadString('\n')
	if err != nil {
		panic(err)
	}
	ruleName = strings.Trim(ruleName, "\n")

	meta := &metadata{RuleNameCC: ToCamel(ruleName), RuleName: ruleName}

	GenerateFileWithLogs(fmt.Sprintf("rules/%s.go", ruleName), "rules/rule.go.tmpl", meta)
	GenerateFileWithLogs(fmt.Sprintf("rules/%s_test.go", ruleName), "rules/rule_test.go.tmpl", meta)
	GenerateFileWithLogs(fmt.Sprintf("docs/rules/%s.md", ruleName), "docs/rules/rule.md.tmpl", meta)
}

func ToCamel(str string) string {
	exceptions := map[string]string{}
	parts := strings.Split(str, "_")
	replaced := make([]string, len(parts))
	for i, s := range parts {
		conv, ok := exceptions[s]
		if ok {
			replaced[i] = conv
		} else {
			replaced[i] = s
		}
	}
	str = strings.Join(replaced, "_")

	return heading.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

// GenerateFile generates a new file from the passed template and metadata
func GenerateFile(fileName string, tmplName string, meta interface{}) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles(tmplName))
	err = tmpl.Execute(file, meta)
	if err != nil {
		panic(err)
	}
}

// GenerateFileWithLogs generates a new file from the passed template and metadata
// The difference from GenerateFile function is to output logs
func GenerateFileWithLogs(fileName string, tmplName string, meta interface{}) {
	GenerateFile(fileName, tmplName, meta)
	fmt.Printf("Created: %s\n", fileName)
}
