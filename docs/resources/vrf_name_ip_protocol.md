---
page_title: "vyos_vrf_name_ip_protocol Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  IPv4 routing parameters
  ⯯
  Filter routing info exchanged between routing protocol and zebra
---

# vyos_vrf_name_ip_protocol (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
IPv4 routing parameters
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
    &emsp;|eigrp      &emsp;|Enhanced Interior Gateway Routing Protocol           |
    &emsp;|isis       &emsp;|Intermediate System to Intermediate System           |
    &emsp;|kernel     &emsp;|Kernel routes (not installed via the zebra RIB)      |
    &emsp;|ospf       &emsp;|Open Shortest Path First (OSPFv2)                    |
    &emsp;|rip        &emsp;|Routing Information Protocol                         |
    &emsp;|static     &emsp;|Statically configured routes                         |

### Optional

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
