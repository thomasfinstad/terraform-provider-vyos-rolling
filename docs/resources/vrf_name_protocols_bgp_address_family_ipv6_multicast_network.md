---
page_title: "vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_network Resource - terraform-provider-vyos"
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
  Import BGP network/prefix into multicast IPv6 RIB
---

# vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_network (Resource)
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
**Import BGP network/prefix into multicast IPv6 RIB**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `network_id` (String) Import BGP network/prefix into multicast IPv6 RIB

    &emsp;|Format   &emsp;|Description                        |
    |-----------|-------------------------------------|
    &emsp;|ipv6net  &emsp;|Multicast IPv6 BGP network/prefix  |

### Optional

- `path_limit` (Number) AS-path hopcount limit

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|0-255   &emsp;|AS path hop count limit  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
