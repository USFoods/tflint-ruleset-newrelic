# nr_synthetics_monitor_invalid_period

// TODO: Write the rule's description here

## Example

```hcl
resource "newrelic_synthetics_monitor" "monitor" {
  status           = "ENABLED"
  name             = "monitor"
  period           = "60" // invalid value
  uri              = "https://www.one.newrelic.com"
  type             = "SIMPLE"
  locations_public = ["AP_SOUTH_1"]

  custom_header {
    name  = "some_name"
    value = "some_value"
  }

  treat_redirect_as_failure = true
  validation_string         = "success"
  bypass_head_request       = true
  verify_ssl                = true

  tag {
    key    = "some_key"
    values = ["some_value"]
  }
}
```

```bash
$ tflint

Error: '60' is invalid period (nr_synthetics_monitor_invalid_period)

  on main.tf line 18:
  18:   period           = "60" // invalid value

Reference: https://github.com/usfoods/tflint-ruleset-newrelic/blob/v0.4.0/docs/rules/nr_synthetics_monitor_invalid_period.md

```

## Why

// TODO: Write why you should follow the rule. This section is also a place to explain the value of the rule

## How To Fix

// TODO: Write how to fix it to avoid the problem
