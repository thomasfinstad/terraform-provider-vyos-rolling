---
page_title: "vyos_vrf_name_protocols_static_route_next_hop_bfd_multi_hop_source Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Static Routing⯯Static IPv4 route⯯Next-hop IPv4 router address⯯BFD monitoring⯯Use BFD multi hop session⯯Use source for BFD session
---

# vyos_vrf_name_protocols_static_route_next_hop_bfd_multi_hop_source (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Static Routing  
⯯  
Static IPv4 route  
⯯  
Next-hop IPv4 router address  
⯯  
BFD monitoring  
⯯  
Use BFD multi hop session  
⯯  
**Use source for BFD session**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `next_hop_id` (String) Next-hop IPv4 router address

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|ipv4    &emsp;|Next-hop router address  |
- `route_id` (String) Static IPv4 route

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv4net  &emsp;|IPv4 static route  |
- `source_id` (String) Use source for BFD session

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|ipv4    &emsp;|IPv4 source address  |
    &emsp;|ipv6    &emsp;|IPv6 source address  |

### Optional

- `profile` (String) Use settings from BFD profile

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|BFD profile name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
