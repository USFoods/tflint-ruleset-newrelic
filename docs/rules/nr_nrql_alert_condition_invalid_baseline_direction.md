# nr_nrql_alert_condition_invalid_baseline_direction

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
  baseline_direction = "JUST_UPPER" // invalid value!

  nrql {
    query = "SELECT average(duration) FROM Transaction where appName = 'Your App'"
  }

  critical {
    operator              = "above"
    threshold             = 5.5
    threshold_duration    = 300
    threshold_occurrences = "ALL"
  }
}
```

```bash
$ tflint

Error: 'JUST_UPPER' is invalid value for baseline_direction (nr_nrql_alert_condition_invalid_baseline_direction)

  on main.tf line 31:
  31:   baseline_direction = "JUST_UPPER" // invalid value!

```

## Why

// TODO: Write why you should follow the rule. This section is also a place to explain the value of the rule

## How To Fix

// TODO: Write how to fix it to avoid the problem
