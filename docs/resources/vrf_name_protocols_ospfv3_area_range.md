---
page_title: "vyos_vrf_name_protocols_ospfv3_area_range Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  Routing protocol parameters
  ⯯
  Open Shortest Path First (OSPF) for IPv6
  ⯯
  OSPFv3 Area
  ⯯
  Specify IPv6 prefix (border routers only)
---

# vyos_vrf_name_protocols_ospfv3_area_range (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Open Shortest Path First (OSPF) for IPv6
⯯
OSPFv3 Area
⯯
**Specify IPv6 prefix (border routers only)**


</center>

## Schema

### Required

- `area_id` (String) OSPFv3 Area

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|u32     &emsp;|Area ID as a decimal value   |
    &emsp;|ipv4    &emsp;|Area ID in IP address forma  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `range_id` (String) Specify IPv6 prefix (border routers only)

    &emsp;|Format   &emsp;|Description                                |
    |-----------|---------------------------------------------|
    &emsp;|ipv6net  &emsp;|Specify IPv6 prefix (border routers only)  |

### Optional

- `advertise` (Boolean) Advertise this range
- `not_advertise` (Boolean) Do not advertise this range

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
