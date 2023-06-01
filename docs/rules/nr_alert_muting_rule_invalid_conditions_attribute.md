# nr_alert_muting_rule_invalid_conditions_attribute

// TODO: Write the rule's description here

## Example

```hcl
resource "newrelic_alert_muting_rule" "foo" {
    name = "Example Muting Rule"
    enabled = true
    description = "muting rule test."
    condition {
        conditions {
            attribute   = "type" // invalid value!
            operator    = "EQUALS"
            values      = ["APM"]
        }
        operator = "AND"
    }
}
```

```bash
$ tflint

Error: 'type' is invalid value for conditions attribute (nr_alert_muting_rule_invalid_conditions_attribute)

  on main.tf line 62:
  62:             attribute   = "type" // invalid value!

```

## Why

// TODO: Write why you should follow the rule. This section is also a place to explain the value of the rule

## How To Fix

// TODO: Write how to fix it to avoid the problem
