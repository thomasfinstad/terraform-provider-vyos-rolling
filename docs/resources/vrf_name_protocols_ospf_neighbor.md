---
page_title: "vyos_vrf_name_protocols_ospf_neighbor Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  Routing protocol parameters
  ⯯
  Open Shortest Path First (OSPF)
  ⯯
  Specify neighbor router
---

# vyos_vrf_name_protocols_ospf_neighbor (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Open Shortest Path First (OSPF)
⯯
**Specify neighbor router**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `neighbor_id` (String) Specify neighbor router

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|ipv4    &emsp;|Neighbor IP address  |

### Optional

- `poll_interval` (Number) Dead neighbor polling interval

    &emsp;|Format   &emsp;|Description                                     |
    |-----------|--------------------------------------------------|
    &emsp;|1-65535  &emsp;|Seconds between dead neighbor polling interval  |
- `priority` (Number) Neighbor priority in seconds

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|0-255   &emsp;|Neighbor priority  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
