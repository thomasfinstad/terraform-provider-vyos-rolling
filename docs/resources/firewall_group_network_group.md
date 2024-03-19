---
page_title: "vyos_firewall_group_network_group Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Firewall group
  ⯯
  Firewall network-group
---

# vyos_firewall_group_network_group (Resource)
<center>

Firewall
⯯
Firewall group
⯯
**Firewall network-group**


</center>

## Schema

### Required

- `network_group_id` (String) Firewall network-group

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `include` (List of String) Include another network-group
- `network` (List of String) Network-group member

    &emsp;|Format   &emsp;|Description           |
    |-----------|------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 Subnet to match  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
