---
page_title: "vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_node_protecting_index Resource - terraform-provider-vyos"
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
  Local loop free alternate options
  ⯯
  Configure tiebreaker for multiple backups
  ⯯
  Prefer node protecting backup path
  ⯯
  Set preference order among tiebreakers
---

# vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_node_protecting_index (Resource)
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
Local loop free alternate options
⯯
Configure tiebreaker for multiple backups
⯯
Prefer node protecting backup path
⯯
**Set preference order among tiebreakers**


</center>

## Schema

### Required

- `index_id` (Number) Set preference order among tiebreakers

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|The index integer value  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Optional

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
