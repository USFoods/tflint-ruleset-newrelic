# nr_synthetics_script_monitor_invalid_type

// TODO: Write the rule's description here

## Example

```hcl
resource "newrelic_synthetics_script_monitor" "monitor" {
  status               = "ENABLED"
  name                 = "script_monitor"
  type                 = "SCRIPT_MONITOR" // invalid value
  locations_public     = ["US_WEST_1", "US_WEST_2"]
  period               = "EVERY_HOUR" 

  script               = "console.log('it works!')"

  script_language      = "JAVASCRIPT"
  runtime_type         = "NODE_API"
  runtime_type_version = "16.10"

  tag {
    key    = "some_key"
    values = ["some_value"]
  }
}
```

```bash
$ tflint

Error: 'SCRIPT_MONITOR' is invalid monitor type (nr_synthetics_script_monitor_invalid_type)

  on main.tf line 18:
  18:   type                 = "SCRIPT_MONITOR" // invalid value

Reference: https://github.com/usfoods/tflint-ruleset-newrelic/blob/v0.4.0/docs/rules/nr_synthetics_script_monitor_invalid_type.md

```

## Why

// TODO: Write why you should follow the rule. This section is also a place to explain the value of the rule

## How To Fix

// TODO: Write how to fix it to avoid the problem
