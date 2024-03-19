---
page_title: "vyos_firewall_group_domain_group Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Firewall group
  ⯯
  Firewall domain-group
---

# vyos_firewall_group_domain_group (Resource)
<center>

Firewall
⯯
Firewall group
⯯
**Firewall domain-group**


</center>

## Schema

### Required

- `domain_group_id` (String) Firewall domain-group

### Optional

- `address` (List of String) Domain-group member

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|txt     &emsp;|Domain address to match  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
