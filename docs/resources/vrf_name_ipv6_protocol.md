---
page_title: "vyos_vrf_name_ipv6_protocol Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯IPv6 routing parameters⯯Filter routing info exchanged between routing protocol and zebra
---

# vyos_vrf_name_ipv6_protocol (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
IPv6 routing parameters  
⯯  
**Filter routing info exchanged between routing protocol and zebra**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `protocol_id` (String) Filter routing info exchanged between routing protocol and zebra

    &emsp;|Format     &emsp;|Description                                          |
    |-------------|-------------------------------------------------------|
    &emsp;|any        &emsp;|Any of the above protocols                           |
    &emsp;|babel      &emsp;|Babel routing protocol                               |
    &emsp;|bgp        &emsp;|Border Gateway Protocol                              |
    &emsp;|connected  &emsp;|Connected routes (directly attached subnet or host)  |
    &emsp;|isis       &emsp;|Intermediate System to Intermediate System           |
    &emsp;|kernel     &emsp;|Kernel routes (not installed via the zebra RIB)      |
    &emsp;|ospfv3     &emsp;|Open Shortest Path First (OSPFv3)                    |
    &emsp;|ripng      &emsp;|Routing Information Protocol next-generation         |
    &emsp;|static     &emsp;|Statically configured routes                         |

### Optional

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
