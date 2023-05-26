# nr_alert_muting_rule_invalid_conditions_operator

// TODO: Write the rule's description here

## Example

```hcl
resource "newrelic_alert_muting_rule" "foo" {
    name = "Example Muting Rule"
    enabled = true
    description = "muting rule test."
    condition {
        conditions {
            attribute   = "product" 
            operator    = "ALL" // invalid value!
            values      = ["APM"]
        }
        operator = "AND"
    }
}
```

```bash
$ tflint

Error: 'ALL' is invalid value for conditions operator (nr_alert_muting_rule_invalid_conditions_operator)

  on main.tf line 63:
  63:             operator    = "ALL" // invalid value!

```

## Why

// TODO: Write why you should follow the rule. This section is also a place to explain the value of the rule

## How To Fix

// TODO: Write how to fix it to avoid the problem
