---
page_title: "vyos_vrf_name_protocols_isis_fast_reroute_lfa_remote_prefix_list Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  Routing protocol parameters
  ⯯
  Intermediate System to Intermediate System (IS-IS)
  ⯯
  IS-IS fast reroute configuration
  ⯯
  Loop free alternate functionality
  ⯯
  Remote loop free alternate options
  ⯯
  Filter PQ node router ID based on prefix list
---

# vyos_vrf_name_protocols_isis_fast_reroute_lfa_remote_prefix_list (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Intermediate System to Intermediate System (IS-IS)
⯯
IS-IS fast reroute configuration
⯯
Loop free alternate functionality
⯯
Remote loop free alternate options
⯯
**Filter PQ node router ID based on prefix list**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `prefix_list_id` (String) Filter PQ node router ID based on prefix list

    &emsp;|Format  &emsp;|Description                    |
    |----------|---------------------------------|
    &emsp;|txt     &emsp;|Name of IPv4/IPv6 prefix-list  |

### Optional

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
