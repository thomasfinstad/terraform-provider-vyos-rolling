---
page_title: "vyos_vrf_name_protocols_bgp_neighbor_local_role Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  Routing protocol parameters
  ⯯
  Border Gateway Protocol (BGP)
  ⯯
  BGP neighbor
  ⯯
  Local role for BGP neighbor (RFC9234)
---

# vyos_vrf_name_protocols_bgp_neighbor_local_role (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Border Gateway Protocol (BGP)
⯯
BGP neighbor
⯯
**Local role for BGP neighbor (RFC9234)**


</center>

## Schema

### Required

- `local_role_id` (String) Local role for BGP neighbor (RFC9234)

    &emsp;|Format     &emsp;|Description             |
    |-------------|--------------------------|
    &emsp;|customer   &emsp;|Using Transit           |
    &emsp;|peer       &emsp;|Public/Private Peering  |
    &emsp;|provider   &emsp;|Providing Transit       |
    &emsp;|rs-client  &emsp;|RS Client               |
    &emsp;|rs-server  &emsp;|Route Server            |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `neighbor_id` (String) BGP neighbor

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|ipv4    &emsp;|BGP neighbor IP address    |
    &emsp;|ipv6    &emsp;|BGP neighbor IPv6 address  |
    &emsp;|txt     &emsp;|Interface name             |

### Optional

- `strict` (Boolean) Neighbor must send this exact capability, otherwise a role missmatch notification will be sent

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
