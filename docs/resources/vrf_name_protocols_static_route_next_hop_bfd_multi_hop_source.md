---
page_title: "vyos_vrf_name_protocols_static_route_next_hop_bfd_multi_hop_source Resource - terraform-provider-vyos"
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
  Static IPv4 route
  ⯯
  Next-hop IPv4 router address
  ⯯
  BFD monitoring
  ⯯
  Use BFD multi hop session
  ⯯
  Use source for BFD session
---

# vyos_vrf_name_protocols_static_route_next_hop_bfd_multi_hop_source (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Static Routing
⯯
Static IPv4 route
⯯
Next-hop IPv4 router address
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
- `next_hop_id` (String) Next-hop IPv4 router address

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|ipv4    &emsp;|Next-hop router address  |
- `route_id` (String) Static IPv4 route

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv4net  &emsp;|IPv4 static route  |
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
