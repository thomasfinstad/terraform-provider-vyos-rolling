---
page_title: "vyos_firewall_group_ipv6_address_group Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Firewall group
  ⯯
  Firewall ipv6-address-group
---

# vyos_firewall_group_ipv6_address_group (Resource)
<center>

Firewall
⯯
Firewall group
⯯
**Firewall ipv6-address-group**


</center>

## Schema

### Required

- `ipv6_address_group_id` (String) Firewall ipv6-address-group

### Optional

- `address` (List of String) Address-group member

    &emsp;|Format     &emsp;|Description                                  |
    |-------------|-----------------------------------------------|
    &emsp;|ipv6       &emsp;|IPv6 address to match                        |
    &emsp;|ipv6range  &emsp;|IPv6 range to match (e.g. 2002::1-2002::ff)  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `include` (List of String) Include another ipv6-address-group

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
