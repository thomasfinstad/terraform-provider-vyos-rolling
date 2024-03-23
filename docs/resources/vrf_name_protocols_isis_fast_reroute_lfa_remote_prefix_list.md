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
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
