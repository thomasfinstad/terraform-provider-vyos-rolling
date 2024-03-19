---
page_title: "vyos_vrf_name_protocols_static_route Resource - terraform-provider-vyos"
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
  Static IPv4 route
---

# vyos_vrf_name_protocols_static_route (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Static Routing
⯯
**Static IPv4 route**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `route_id` (String) Static IPv4 route

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv4net  &emsp;|IPv4 static route  |

### Optional

- `blackhole` (Attributes) Silently discard pkts when matched (see [below for nested schema](#nestedatt--blackhole))
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `dhcp_interface` (String) DHCP interface supplying next-hop IP address

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|txt     &emsp;|DHCP interface name  |
- `reject` (Attributes) Emit an ICMP unreachable when matched (see [below for nested schema](#nestedatt--reject))

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
    &emsp;|1-4294967295  &emsp;|Tag value for this route  |  &emsp;|
