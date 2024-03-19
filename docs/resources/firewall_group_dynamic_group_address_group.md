---
page_title: "vyos_firewall_group_dynamic_group_address_group Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Firewall group
  ⯯
  Firewall dynamic group
  ⯯
  Firewall dynamic address group
---

# vyos_firewall_group_dynamic_group_address_group (Resource)
<center>

Firewall
⯯
Firewall group
⯯
Firewall dynamic group
⯯
**Firewall dynamic address group**


</center>

## Schema

### Required

- `address_group_id` (String) Firewall dynamic address group

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
