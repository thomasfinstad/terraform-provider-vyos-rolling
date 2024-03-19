---
page_title: "vyos_vrf_name_protocols_bgp_listen_range Resource - terraform-provider-vyos"
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
  Listen for and accept BGP dynamic neighbors from range
  ⯯
  BGP dynamic neighbors listen range
---

# vyos_vrf_name_protocols_bgp_listen_range (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Border Gateway Protocol (BGP)
⯯
Listen for and accept BGP dynamic neighbors from range
⯯
**BGP dynamic neighbors listen range**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `range_id` (String) BGP dynamic neighbors listen range

    &emsp;|Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 dynamic neighbors listen range  |
    &emsp;|ipv6net  &emsp;|IPv6 dynamic neighbors listen range  |

### Optional

- `peer_group` (String) Peer group for this peer

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|Peer-group name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
