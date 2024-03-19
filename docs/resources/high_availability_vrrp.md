---
page_title: "vyos_high_availability_vrrp Resource - terraform-provider-vyos"
subcategory: "high"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  High availability settings
  ⯯
  Virtual Router Redundancy Protocol settings
---

# vyos_high_availability_vrrp (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

	High availability settings
⯯
**Virtual Router Redundancy Protocol settings**


</center>

## Schema

### Optional

- `snmp` (Boolean) Enable SNMP

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
