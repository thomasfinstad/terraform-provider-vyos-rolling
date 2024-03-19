---
page_title: "vyos_firewall_group_mac_group Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Firewall group
  ⯯
  Firewall mac-group
---

# vyos_firewall_group_mac_group (Resource)
<center>

Firewall
⯯
Firewall group
⯯
**Firewall mac-group**


</center>

## Schema

### Required

- `mac_group_id` (String) Firewall mac-group

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `include` (List of String) Include another mac-group
- `mac_address` (List of String) Mac-group member

    &emsp;|Format   &emsp;|Description           |
    |-----------|------------------------|
    &emsp;|macaddr  &emsp;|MAC address to match  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
