---
page_title: "vyos_firewall_flowtable Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Flowtable
---

# vyos_firewall_flowtable (Resource)
<center>

Firewall
⯯
**Flowtable**


</center>

## Schema

### Required

- `flowtable_id` (String) Flowtable

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `interface` (List of String) Interfaces to use this flowtable
- `offload` (String) Offloading method

    &emsp;|Format    &emsp;|Description       |
    |------------|--------------------|
    &emsp;|hardware  &emsp;|Hardware offload  |
    &emsp;|software  &emsp;|Software offload  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
