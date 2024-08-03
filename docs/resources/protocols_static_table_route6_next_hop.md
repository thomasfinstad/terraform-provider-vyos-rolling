---
page_title: "vyos_protocols_static_table_route6_next_hop Resource - vyos"

subcategory: "Protocols"

description: |- 
  Routing protocols⯯Static Routing⯯Policy route table number⯯Static IPv6 route⯯IPv6 gateway address
---

# vyos_protocols_static_table_route6_next_hop (Resource)
<center>

Routing protocols  
⯯  
Static Routing  
⯯  
Policy route table number  
⯯  
Static IPv6 route  
⯯  
**IPv6 gateway address**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `bfd` (Attributes) BFD monitoring (see [below for nested schema](#nestedatt--bfd))
- `disable` (Boolean) Disable instance
- `distance` (Number) Distance for this route

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Distance for this route  |
- `interface` (String) Gateway interface name

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Gateway interface name  |
- `segments` (String) SRv6 segments

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Segs (SIDs)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF to leak route

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Name of VRF to leak to  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `next_hop` (String) IPv6 gateway address

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|ipv6    &emsp;|Next-hop IPv6 router  |
- `route6` (String) Static IPv6 route

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv6net  &emsp;|IPv6 static route  |
- `table` (Number) Policy route table number

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|1-200   &emsp;|Policy route table number  |


&lt;a id=&#34;nestedatt--bfd&#34;&gt;&lt;/a&gt;
### Nested Schema for `bfd`

Optional:

- `multi_hop` (Attributes) Use BFD multi hop session (see [below for nested schema](#nestedatt--bfd--multi_hop))
- `profile` (String) Use settings from BFD profile

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|BFD profile name  |

&lt;a id=&#34;nestedatt--bfd--multi_hop&#34;&gt;&lt;/a&gt;
### Nested Schema for `bfd.multi_hop`



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
