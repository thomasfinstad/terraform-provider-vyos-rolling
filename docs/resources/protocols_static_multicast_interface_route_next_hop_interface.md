---
page_title: "vyos_protocols_static_multicast_interface_route_next_hop_interface Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯static⯯Multicast static route⯯Multicast interface based route⯯Next-hop interface
---

# vyos_protocols_static_multicast_interface_route_next_hop_interface (Resource)
<center>

*protocols*  
⯯  
*static*  
⯯  
Multicast static route  
⯯  
Multicast interface based route  
⯯  
**Next-hop interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `distance` (Number) Distance value for this route

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Distance for this route  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface_route` (String) Multicast interface based route

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|ipv4net  &emsp;|Network      |
- `next_hop_interface` (String) Next-hop interface


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
