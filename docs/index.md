---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos Provider"
subcategory: ""
description: |-
  !> This is for the rolling release of VyOS, it will automatically update when the API schemas change
  -> This provider's version number MIGHT follow <MAJOR>.<MINOR>.<VYOS ROLLING RELEASE DATE>, so Version 1.3 of this provider, built with the API schemas for VyOS rolling release built on 27th of November 1970 would be have the version number 1.3.19701127.This allows for locking to a major version, or even a spessific release of rolling VyOS. This versioning scheme is not final and might change.
  Use Terraform to configure your VyOS instances via API calls.
  Requirements
  To use this provider you must enable the HTTP(S) API on the target instances. See VyOS documentation https://docs.vyos.io/en/latest/configuration/service/https.html for more information.
---

# vyos Provider

!> This is for the rolling release of VyOS, it will automatically update when the API schemas change

-> This provider's version number MIGHT follow `<MAJOR>.<MINOR>.<VYOS ROLLING RELEASE DATE>`, so Version `1.3` of this provider, built with the API schemas for VyOS rolling release built on 27th of November 1970 would be have the version number `1.3.19701127`.This allows for locking to a major version, or even a spessific release of rolling VyOS. This versioning scheme is not final and might change.

Use Terraform to configure your VyOS instances via API calls.

## Requirements
To use this provider you must enable the HTTP(S) API on the target instances. See [VyOS documentation](https://docs.vyos.io/en/latest/configuration/service/https.html) for more information.

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Example Usage

```terraform
# TODO Improve provider example
#  milestone: 3

terraform {
  required_providers {
    vyos = {
      source = "thomasfinstad/vyos-rolling"
    }
  }
}

variable "vyos_host" {
  type    = string
  default = "https://vyos.local"
}

variable "vyos_key" {
  type    = string
  default = "one two three four five"
}

provider "vyos" {
  endpoint = var.vyos_host
  api_key  = var.vyos_key

  certificate = {
    disable_verify = true
  }

  default_timeouts = 2
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_key` (String) VyOS API Key
- `endpoint` (String) VyOS API Endpoint

### Optional

- `certificate` (Attributes) (see [below for nested schema](#nestedatt--certificate))
- `default_timeouts` (Number) Default Create/Read/Update/Destroy timeouts in minutes, can be overridden on a per resource basis. If not configured, defaults to 15.
- `ignore_child_resource_on_delete` (Boolean) Disables the check to see if the resource has any child resources.This can be useful when only a parent resource is configured via terraform.This has no effect on global resources.

  !> **WARNING:** This is extremely destructive and will delete everything below the destroyed resource.
- `ignore_missing_parent_resource_on_create` (Boolean) Disables the check to see if the required parent resource exists on the target machine.This can be helpful when encountering a bug with the provider.
- `overwrite_existing_resources_on_create` (Boolean) Disables the check to see if the resource already exists on the target machine, resulting in possibly overwriting configs without notice.This can be helpful when trying to avoid and change many resources at once.

<a id="nestedatt--certificate"></a>
### Nested Schema for `certificate`

Optional:

- `disable_verify` (Boolean) Disable remote certificate verification, useful for selfsigned certs.
