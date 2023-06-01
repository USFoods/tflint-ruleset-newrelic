# nr_nrql_alert_condition_invalid_aggregation_delay_cadence

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
  type               = "static"
  name               = "foo"
  description        = "Alert when transactions are taking too long"
  runbook_url        = "https://www.example.com"
  enabled            = var.enabled
  aggregation_window = 60
  aggregation_method = "cadence"
  aggregation_delay  = 3900 // invalid value!

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

Error: '3900' is invalid value for aggregation_delay with aggregation_method 'cadence' (nr_nrql_alert_condition_invalid_aggregation_delay_cadence)

  on main.tf line 29:
  29:   aggregation_method = "cadence"

```

## Why

// TODO: Write why you should follow the rule. This section is also a place to explain the value of the rule

## How To Fix

// TODO: Write how to fix it to avoid the problem
