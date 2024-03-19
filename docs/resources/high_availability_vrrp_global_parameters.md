---
page_title: "vyos_high_availability_vrrp_global_parameters Resource - terraform-provider-vyos"
subcategory: "high"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  High availability settings
  ⯯
  Virtual Router Redundancy Protocol settings
  ⯯
  VRRP global parameters
---

# vyos_high_availability_vrrp_global_parameters (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

	High availability settings
⯯
Virtual Router Redundancy Protocol settings
⯯
**VRRP global parameters**


</center>

## Schema

### Optional

- `startup_delay` (Number) Time VRRP startup process (in seconds)

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|1-600   &emsp;|Interval in seconds  |
- `version` (String) Default VRRP version to use, IPv6 always uses VRRP version 3

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|2       &emsp;|VRRP version 2  |
    &emsp;|3       &emsp;|VRRP version 3  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
