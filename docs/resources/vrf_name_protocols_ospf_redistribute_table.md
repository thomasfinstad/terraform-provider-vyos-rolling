---
page_title: "vyos_vrf_name_protocols_ospf_redistribute_table Resource - terraform-provider-vyos"
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
  Redistribute information from another routing protocol
  ⯯
  Redistribute non-main Kernel Routing Table
---

# vyos_vrf_name_protocols_ospf_redistribute_table (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Open Shortest Path First (OSPF)
⯯
Redistribute information from another routing protocol
⯯
**Redistribute non-main Kernel Routing Table**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `table_id` (Number) Redistribute non-main Kernel Routing Table

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|1-200   &emsp;|Policy route table number  |

### Optional

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
