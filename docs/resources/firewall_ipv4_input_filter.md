---
page_title: "vyos_firewall_ipv4_input_filter Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Firewall
  ⯯
  IPv4 firewall
  ⯯
  IPv4 input firewall
  ⯯
  IPv4 firewall input filter
---

# vyos_firewall_ipv4_input_filter (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

	Firewall
⯯
IPv4 firewall
⯯
IPv4 input firewall
⯯
**IPv4 firewall input filter**


</center>

## Schema

### Optional

- `default_action` (String) Default-action for rule-set

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|drop    &emsp;|Drop if no prior rules are hit    |
    &emsp;|accept  &emsp;|Accept if no prior rules are hit  |
- `default_log` (Boolean) Log packets hitting default-action
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
