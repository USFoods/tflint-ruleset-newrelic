# nr_nrql_drop_rule_invalid_action

// TODO: Write the rule's description here

## Example

```hcl
resource "newrelic_nrql_drop_rule" "foo" {
  account_id  = 12345
  description = "Drops all data for MyCustomEvent that comes from the LoadGeneratingApp in the dev environment, because there is too much and we don’t look at it."
  action      = "discard" // invalid value
  nrql        = "SELECT * FROM MyCustomEvent WHERE appName='LoadGeneratingApp' AND environment='development'"
}
```

```bash
$ tflint

Error: 'discard' is invalid value for action (nr_nrql_drop_rule_invalid_action)

  on main.tf line 18:
  18:   action      = "discard" // invalid value

Reference: https://github.com/usfoods/tflint-ruleset-newrelic/blob/v0.4.0/docs/rules/nr_nrql_drop_rule_invalid_action.md

```

## Why

// TODO: Write why you should follow the rule. This section is also a place to explain the value of the rule

## How To Fix

// TODO: Write how to fix it to avoid the problem
