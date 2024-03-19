---
page_title: "vyos_vrf_name_protocols_bgp_address_family_ipv6_labeled_unicast_network Resource - terraform-provider-vyos"
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
  Labeled Unicast IPv6 BGP settings
  ⯯
  Import BGP network/prefix into labeled unicast IPv6 RIB
---

# vyos_vrf_name_protocols_bgp_address_family_ipv6_labeled_unicast_network (Resource)
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
Labeled Unicast IPv6 BGP settings
⯯
**Import BGP network/prefix into labeled unicast IPv6 RIB**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `network_id` (String) Import BGP network/prefix into labeled unicast IPv6 RIB

    &emsp;|Format   &emsp;|Description                              |
    |-----------|-------------------------------------------|
    &emsp;|ipv6net  &emsp;|Labeled Unicast IPv6 BGP network/prefix  |

### Optional

- `backdoor` (Boolean) Use BGP network/prefix as a backdoor route
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
