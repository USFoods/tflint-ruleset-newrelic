---
name: Bug report
about: Report a bug
title: ''
labels: ''
assignees: ''

---

Hi there,

Thank you for opening an issue. In order to better assist you with your issue, we kindly ask to follow the template format and instructions. Please note that we try to keep the issue tracker reserved for **bug reports** and **feature requests** only. General usage questions submitted as issues will be closed.

## Include the following

> :warning: **Important:** Failure to include the following, such as omitting the Terraform configuration in question, may delay resolving the issue.

- [ ] Your New Relic `provider` [configuration](#terraform-configuration) (sensitive details redacted)
- [ ] Your TFLint [configuration](#tflint-configuration)
- [ ] A list of [affected resources](#affected-resources) and/or data sources
- [ ] The [configuration](#terraform-configuration) of the related resources (i.e. from the list mentioned above)
- [ ] Description of the [current behavior](#actual-behavior) (the bug)
- [ ] Description of the [expected behavior](#expected-behavior)
- [ ] Any related [log output](#debug-output)

### Terraform Version

Run `terraform -v` to show the version. If you are not running the latest version of Terraform, please upgrade because your issue may have already been fixed.

### Affected Resource(s)

Please list the resources as a list, for example:

- `newrelic_alert_policy`
- `newrelic_nrql_alert_condition`

### Terraform Configuration

> Please include your `provider` configuration (sensitive details redacted) as well as the configuration of the resources and/or data sources related to the bug report.

```hcl
# Paste your Terraform configurations here - for large Terraform configs,
# please provide a link to a GitHub Gist containing the complete configuration
```

### TFLint Configuration

> Please include your TFLint configuration

```hcl
# Paste your TFLint configuration here - for large TFLint configs, please
# provide a link to a GitHub Gist containing the complete configuration
```

### Actual Behavior

What actually happened?

### Expected Behavior

What should have happened?

### Steps to Reproduce

Please list the steps required to reproduce the issue, for example:

1. `tflint --recursive --module`

### Debug Output

Please provide a link to a GitHub Gist containing the complete debug output. Please do NOT paste the debug output in the issue; just paste a link to the Gist.

### Panic Output

If TFLint produced a panic, please provide a link to a GitHub Gist containing the output.

### Important Factoids

Are there anything atypical about your configuration that we should know?

### References

Are there any other GitHub issues (open or closed) or Pull Requests that should be linked here? For example:

- GH-1234
