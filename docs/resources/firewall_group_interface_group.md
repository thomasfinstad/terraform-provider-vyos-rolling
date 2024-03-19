---
page_title: "vyos_firewall_group_interface_group Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Firewall group
  ⯯
  Firewall interface-group
---

# vyos_firewall_group_interface_group (Resource)
<center>

Firewall
⯯
Firewall group
⯯
**Firewall interface-group**


</center>

## Schema

### Required

- `interface_group_id` (String) Firewall interface-group

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `include` (List of String) Include another interface-group
- `interface` (List of String) Interface-group member

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
