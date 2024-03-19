---
page_title: "vyos_high_availability_vrrp_group_address Resource - terraform-provider-vyos"
subcategory: "high"
description: |-
  High availability settings
  ⯯
  Virtual Router Redundancy Protocol settings
  ⯯
  VRRP group
  ⯯
  Virtual IP address
---

# vyos_high_availability_vrrp_group_address (Resource)
<center>

High availability settings
⯯
Virtual Router Redundancy Protocol settings
⯯
VRRP group
⯯
**Virtual IP address**


</center>

## Schema

### Required

- `address_id` (String) Virtual IP address

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 address and prefix length  |
    &emsp;|ipv6net  &emsp;|IPv6 address and prefix length  |
    &emsp;|ipv4     &emsp;|IPv4 address                    |
    &emsp;|ipv6     &emsp;|IPv6 address                    |
- `group_id` (String) VRRP group

### Optional

- `interface` (String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
