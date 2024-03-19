---
page_title: "vyos_firewall_group_ipv6_network_group Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Firewall group
  ⯯
  Firewall ipv6-network-group
---

# vyos_firewall_group_ipv6_network_group (Resource)
<center>

Firewall
⯯
Firewall group
⯯
**Firewall ipv6-network-group**


</center>

## Schema

### Required

- `ipv6_network_group_id` (String) Firewall ipv6-network-group

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `include` (List of String) Include another ipv6-network-group
- `network` (List of String) Network-group member

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|ipv6net  &emsp;|IPv6 address to match  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
