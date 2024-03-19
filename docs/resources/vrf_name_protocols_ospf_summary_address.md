---
page_title: "vyos_vrf_name_protocols_ospf_summary_address Resource - terraform-provider-vyos"
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
  External summary address
---

# vyos_vrf_name_protocols_ospf_summary_address (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Open Shortest Path First (OSPF)
⯯
**External summary address**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `summary_address_id` (String) External summary address

    &emsp;|Format   &emsp;|Description                                  |
    |-----------|-----------------------------------------------|
    &emsp;|ipv4net  &emsp;|OSPF area number in dotted decimal notation  |

### Optional

- `no_advertise` (Boolean) Don not advertise summary route
- `tag` (Number) Router tag

    &emsp;|Format        &emsp;|Description       |
    |----------------|--------------------|
    &emsp;|1-4294967295  &emsp;|Router tag value  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
