---
page_title: "vyos_vrf_name_protocols_static_route_interface Resource - terraform-provider-vyos"
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
  Next-hop IPv4 router interface
---

# vyos_vrf_name_protocols_static_route_interface (Resource)
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
**Next-hop IPv4 router interface**


</center>

## Schema

### Required

- `interface_id` (String) Next-hop IPv4 router interface

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Gateway interface name  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `route_id` (String) Static IPv4 route

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv4net  &emsp;|IPv4 static route  |

### Optional

- `disable` (Boolean) Disable instance
- `distance` (Number) Distance for this route

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Distance for this route  |
- `vrf` (String) VRF to leak route

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Name of VRF to leak to  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
