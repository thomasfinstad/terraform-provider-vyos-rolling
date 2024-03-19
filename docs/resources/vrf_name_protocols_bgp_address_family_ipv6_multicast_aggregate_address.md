---
page_title: "vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_aggregate_address Resource - terraform-provider-vyos"
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
  BGP address-family parameters
  ⯯
  Multicast IPv6 BGP settings
  ⯯
  BGP aggregate network/prefix
---

# vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_aggregate_address (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Border Gateway Protocol (BGP)
⯯
BGP address-family parameters
⯯
Multicast IPv6 BGP settings
⯯
**BGP aggregate network/prefix**


</center>

## Schema

### Required

- `aggregate_address_id` (String) BGP aggregate network/prefix

    &emsp;|Format   &emsp;|Description                   |
    |-----------|--------------------------------|
    &emsp;|ipv6net  &emsp;|BGP aggregate network/prefix  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Optional

- `as_set` (Boolean) Generate AS-set path information for this aggregate address
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `summary_only` (Boolean) Announce the aggregate summary network only

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
