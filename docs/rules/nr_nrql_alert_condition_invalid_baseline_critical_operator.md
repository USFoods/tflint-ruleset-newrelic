# nr_nrql_alert_condition_invalid_baseline_critical_operator

// TODO: Write the rule's description here

## Example

```hcl
resource "newrelic_alert_policy" "foo" {
  name                = "example"
  incident_preference = "PER_POLICY"
}

resource "newrelic_nrql_alert_condition" "foo" {
  account_id         = var.account_id
  policy_id          = newrelic_alert_policy.foo.id
  type               = "baseline"
  name               = "foo"
  description        = "Alert when transactions are taking too long"
  runbook_url        = "https://www.example.com"
  enabled            = var.enabled
  aggregation_window = 60 
  aggregation_method = "event_flow" 
  aggregation_delay = 30
  baseline_direction = "UPPER_ONLY" 

  nrql {
    query = "SELECT average(duration) FROM Transaction where appName = 'Your App'"
  }

  critical {
    operator              = "equals" // invalid value!
    threshold             = 5.5
    threshold_duration    = 300
    threshold_occurrences = "ALL"
  }
}
```

```bash
$ tflint

Error: 'equals' is an invalid value for critical operator with type 'baseline' (nr_nrql_alert_condition_invalid_baseline_critical_operator)

  on main.tf line 38:
  38:     operator              = "equals" // invalid value!

Reference: https://github.com/usfoods/tflint-ruleset-newrelic/blob/v0.4.0/docs/rules/nr_nrql_alert_condition_invalid_baseline_critical_operator.md

```

## Why

// TODO: Write why you should follow the rule. This section is also a place to explain the value of the rule

## How To Fix

// TODO: Write how to fix it to avoid the problem
