---
page_title: "vyos_vrf_name_protocols_ospf_access_list Resource - terraform-provider-vyos"
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
  Access list to filter networks in routing updates
---

# vyos_vrf_name_protocols_ospf_access_list (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Open Shortest Path First (OSPF)
⯯
**Access list to filter networks in routing updates**


</center>

## Schema

### Required

- `access_list_id` (Number) Access list to filter networks in routing updates

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|u32     &emsp;|Access-list number  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Optional

- `export` (List of String) Filter for outgoing routing update

    &emsp;|Format     &emsp;|Description              |
    |-------------|---------------------------|
    &emsp;|bgp        &emsp;|Filter BGP routes        |
    &emsp;|connected  &emsp;|Filter connected routes  |
    &emsp;|isis       &emsp;|Filter IS-IS routes      |
    &emsp;|kernel     &emsp;|Filter Kernel routes     |
    &emsp;|rip        &emsp;|Filter RIP routes        |
    &emsp;|static     &emsp;|Filter static routes     |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
