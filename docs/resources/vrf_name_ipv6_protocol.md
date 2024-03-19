---
page_title: "vyos_vrf_name_ipv6_protocol Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  IPv6 routing parameters
  ⯯
  Filter routing info exchanged between routing protocol and zebra
---

# vyos_vrf_name_ipv6_protocol (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
IPv6 routing parameters
⯯
**Filter routing info exchanged between routing protocol and zebra**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `protocol_id` (String) Filter routing info exchanged between routing protocol and zebra

    &emsp;|Format     &emsp;|Description                                          |
    |-------------|-------------------------------------------------------|
    &emsp;|any        &emsp;|Any of the above protocols                           |
    &emsp;|babel      &emsp;|Babel routing protocol                               |
    &emsp;|bgp        &emsp;|Border Gateway Protocol                              |
    &emsp;|connected  &emsp;|Connected routes (directly attached subnet or host)  |
    &emsp;|isis       &emsp;|Intermediate System to Intermediate System           |
    &emsp;|kernel     &emsp;|Kernel routes (not installed via the zebra RIB)      |
    &emsp;|ospfv3     &emsp;|Open Shortest Path First (OSPFv3)                    |
    &emsp;|ripng      &emsp;|Routing Information Protocol next-generation         |
    &emsp;|static     &emsp;|Statically configured routes                         |

### Optional

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
