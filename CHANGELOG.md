CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

---

## 3.6.0 (2021-03-16)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 3.5.0 (2021-03-09)
* Upgrade to v2.3.0 of the KeyCloak Terraform Provider

## 3.4.1 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

## 3.4.0 (2021-02-01)
* Upgrade to pulumi-terraform-bridge v2.18.1

## 3.3.0 (2021-01-27)
* Upgrade to v2.2.0 of the KeyCloak Terraform Provider

## 3.2.0 (2021-01-12)
* Upgrade to v2.1.0 of the KeyCloak Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to pulumi v2.16.2

## 3.1.1 (2020-11-24)
* Upgrade to pulumi-terraform-bridge v2.13.2  
  * This adds support for import specific examples in documentation

## 3.1.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 3.0.0 (2020-09-30)
* Upgrade to pulumi-terraform-bridge v2.8.0
* Upgrade to Pulumi v2.10.0
* Upgrade to v2.0.0 of the KeyCloak Terraform Provider

## 2.5.0 (2020-08-31)
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python

## 2.4.0 (2020-07-22)
* Upgrade to v1.20.0 of the KeyCloak Terraform Provider

## 2.3.2 (2020-06-09)
* GitHub action build process improvements

## 2.3.1 (2020-06-09
* Switch to GitHub actions for releases

## 2.3.0 (2020-06-09)
* Upgrade to v1.19.0 of the KeyCloak Terraform Provider

## 2.2.2 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.2.1 (2020-05-12)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.2.0 (2020-04-28)
* Regenerate datasource examples to be async
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.1.0 (2020-04-20)
* Upgrade to v1.18.0 of the KeyCloak Terraform Provider

## 2.0.0 (2020-04-17)
* Upgrade to pulumi-terraform-bridge v2.0.0
* Upgrade to Pulumi v2.0.0

## 1.4.0 (2020-04-02)
* Upgrade to pulumi-terraform-bridge v1.8.4
* Upgrade to Pulumi v1.13.1
* Change layout to support Go modules

## 1.3.0 (2020-03-21)
* Upgrade to v1.17.1 of the KeyCloak Terraform Provider
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.2.0 (2020-03-10)
* Fix the casing of the Go, NodeJS and Python sub packages
* Upgrade to v1.17.0 of the KeyCloak Terraform Provider

## v1.1.0 (2020-02-25)
* Upgrade to v1.16.0 of the KeyCloak Terraform Provider

## 1.0.0 (2020-02-10)
* Initial release of the provider against v1.15.0 of the KeyCloak Terraform Provider
