---
page_title: "vyos_vrf_name_protocols_ospfv3_area_range Resource - vyos"
subcategory: "vrf"
description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Open Shortest Path First (OSPF) for IPv6⯯OSPFv3 Area⯯Specify IPv6 prefix (border routers only)
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
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
