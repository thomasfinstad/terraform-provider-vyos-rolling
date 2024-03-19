---
page_title: "vyos_vrf_name_protocols_static_route6_next_hop_bfd_multi_hop_source Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  Routing protocol parameters
  ⯯
  Static Routing
  ⯯
  Static IPv6 route
  ⯯
  IPv6 gateway address
  ⯯
  BFD monitoring
  ⯯
  Use BFD multi hop session
  ⯯
  Use source for BFD session
---

# vyos_vrf_name_protocols_static_route6_next_hop_bfd_multi_hop_source (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Static Routing
⯯
Static IPv6 route
⯯
IPv6 gateway address
⯯
BFD monitoring
⯯
Use BFD multi hop session
⯯
**Use source for BFD session**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `next_hop_id` (String) IPv6 gateway address

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|ipv6    &emsp;|Next-hop IPv6 router  |
- `route6_id` (String) Static IPv6 route

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv6net  &emsp;|IPv6 static route  |
- `source_id` (String) Use source for BFD session

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|ipv4    &emsp;|IPv4 source address  |
    &emsp;|ipv6    &emsp;|IPv6 source address  |

### Optional

- `profile` (String) Use settings from BFD profile

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|BFD profile name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
