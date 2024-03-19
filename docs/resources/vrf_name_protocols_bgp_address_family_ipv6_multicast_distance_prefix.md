---
page_title: "vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_distance_prefix Resource - terraform-provider-vyos"
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
  Administrative distances for BGP routes
  ⯯
  Administrative distance for a specific BGP prefix
---

# vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_distance_prefix (Resource)
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
Administrative distances for BGP routes
⯯
**Administrative distance for a specific BGP prefix**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `prefix_id` (String) Administrative distance for a specific BGP prefix

    &emsp;|Format   &emsp;|Description                                        |
    |-----------|-----------------------------------------------------|
    &emsp;|ipv6net  &emsp;|Administrative distance for a specific BGP prefix  |

### Optional

- `distance` (Number) Administrative distance for prefix

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|1-255   &emsp;|Administrative distance for external BGP routes  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
