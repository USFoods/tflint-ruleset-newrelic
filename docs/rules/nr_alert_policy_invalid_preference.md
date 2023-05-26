# nr_alert_policy_invalid_preference

// TODO: Write the rule's description here

## Example

```hcl
resource "newrelic_alert_policy" "foo" {
  name = "example"
  incident_preference = "PER_ISSUE" // invalid value!
}
```

```bash
$ tflint

Error: 'PER_ISSUE' is invalid value for incident_preference (nr_alert_policy_invalid_preference)

  on main.tf line 17:
  17:   incident_preference = "PER_ISSUE" // invalid value!

```

## Why

// TODO: Write why you should follow the rule. This section is also a place to explain the value of the rule

## How To Fix

// TODO: Write how to fix it to avoid the problem
