# TFLint Ruleset for terraform-provider-newrelic
[![Build Status](https://github.com/usfoods/tflint-ruleset-newrelic/workflows/build/badge.svg?branch=main)](https://github.com/usfoods/tflint-ruleset-newrelic/actions)

TFLint ruleset plugin for Terraform New Relic Provider

This ruleset focus on possible errors and best practices about New Relic resources. Many rules are enabled by default and warn against code that might fail when running `terraform apply`, or clearly unrecommened.

## Requirements

- TFLint v0.40+
- Go v1.20

## Installation

TODO: This template repository does not contain release binaries, so this installation will not work. Please rewrite for your repository. See the "Building the plugin" section to get this template ruleset working.

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "template" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/terraform-linters/tflint-ruleset-template"

  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----
  mQINBGCqS2YBEADJ7gHktSV5NgUe08hD/uWWPwY07d5WZ1+F9I9SoiK/mtcNGz4P
  JLrYAIUTMBvrxk3I+kuwhp7MCk7CD/tRVkPRIklONgtKsp8jCke7FB3PuFlP/ptL
  SlbaXx53FCZSOzCJo9puZajVWydoGfnZi5apddd11Zw1FuJma3YElHZ1A1D2YvrF
  ...
  KEY
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|aws_instance_example_type|Example rule for accessing and evaluating top-level attributes|ERROR|✔||
|aws_s3_bucket_example_lifecycle_rule|Example rule for accessing top-level/nested blocks and attributes under the blocks|ERROR|✔||
|google_compute_ssl_policy|Example rule with a custom rule config|WARNING|✔||
|terraform_backend_type|Example rule for accessing other than resources|ERROR|✔||

## Building the plugin

Clone the repository locally and run the following command:

```bash
$ make
```

You can easily install the built plugin with the following:

```bash
$ make install
```

You can run the built plugin like the following:

```bash
$ cat << EOS > .tflint.hcl
plugin "template" {
  enabled = true
}
EOS
$ tflint
```
