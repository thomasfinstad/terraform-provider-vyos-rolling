---
page_title: "vyos_vrf_name_protocols_ospf_area_range Resource - terraform-provider-vyos"
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
  OSPF area settings
  ⯯
  Summarize routes matching a prefix (border routers only)
---

# vyos_vrf_name_protocols_ospf_area_range (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Open Shortest Path First (OSPF)
⯯
OSPF area settings
⯯
**Summarize routes matching a prefix (border routers only)**


</center>

## Schema

### Required

- `area_id` (String) OSPF area settings

    &emsp;|Format  &emsp;|Description                                  |
    |----------|-----------------------------------------------|
    &emsp;|u32     &emsp;|OSPF area number in decimal notation         |
    &emsp;|ipv4    &emsp;|OSPF area number in dotted decimal notation  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `range_id` (String) Summarize routes matching a prefix (border routers only)

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv4net  &emsp;|Area range prefix  |

### Optional

- `cost` (Number) Metric for this range

    &emsp;|Format      &emsp;|Description            |
    |--------------|-------------------------|
    &emsp;|0-16777215  &emsp;|Metric for this range  |
- `not_advertise` (Boolean) Do not advertise this range
- `substitute` (String) Advertise area range as another prefix

    &emsp;|Format   &emsp;|Description                             |
    |-----------|------------------------------------------|
    &emsp;|ipv4net  &emsp;|Advertise area range as another prefix  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
