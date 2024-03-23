---
page_title: "vyos_vrf_name_protocols_static_route6 Resource - terraform-provider-vyos"
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
  Static IPv6 route
---

# vyos_vrf_name_protocols_static_route6 (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Static Routing
⯯
**Static IPv6 route**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `route6_id` (String) Static IPv6 route

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv6net  &emsp;|IPv6 static route  |

### Optional

- `blackhole` (Attributes) Silently discard pkts when matched (see [below for nested schema](#nestedatt--blackhole))
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `reject` (Attributes) Emit an ICMP unreachable when matched (see [below for nested schema](#nestedatt--reject))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--blackhole&#34;&gt;&lt;/a&gt;
### Nested Schema for `blackhole`

Optional:

- `distance` (Number) Distance for this route

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Distance for this route  |
- `tag` (Number) Tag value for this route

    &emsp;|Format        &emsp;|Description               |
    |----------------|----------------------------|
    &emsp;|1-4294967295  &emsp;|Tag value for this route  |


&lt;a id=&#34;nestedatt--reject&#34;&gt;&lt;/a&gt;
### Nested Schema for `reject`

Optional:

- `distance` (Number) Distance for this route

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Distance for this route  |
- `tag` (Number) Tag value for this route

    &emsp;|Format        &emsp;|Description               |
    |----------------|----------------------------|
    &emsp;|1-4294967295  &emsp;|Tag value for this route  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
