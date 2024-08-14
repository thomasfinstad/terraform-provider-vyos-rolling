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

    |Format  &emsp;|Description              |
    |----------|---------------------------|
    |1-255   &emsp;|Distance for this route  |
- `interface` (String) Gateway interface name

    |Format  &emsp;|Description             |
    |----------|--------------------------|
    |txt     &emsp;|Gateway interface name  |
- `segments` (String) SRv6 segments

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Segs (SIDs)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF to leak route

    |Format  &emsp;|Description             |
    |----------|--------------------------|
    |txt     &emsp;|Name of VRF to leak to  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `next_hop` (String) IPv6 gateway address

    |Format  &emsp;|Description           |
    |----------|------------------------|
    |ipv6    &emsp;|Next-hop IPv6 router  |
- `route6` (String) Static IPv6 route

    |Format   &emsp;|Description        |
    |-----------|---------------------|
    |ipv6net  &emsp;|IPv6 static route  |
- `table` (Number) Policy route table number

    |Format  &emsp;|Description                |
    |----------|-----------------------------|
    |1-200   &emsp;|Policy route table number  |


<a id="nestedatt--bfd"></a>
### Nested Schema for `bfd`

Optional:

- `multi_hop` (Attributes) Use BFD multi hop session (see [below for nested schema](#nestedatt--bfd--multi_hop))
- `profile` (String) Use settings from BFD profile

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |txt     &emsp;|BFD profile name  |

<a id="nestedatt--bfd--multi_hop"></a>
### Nested Schema for `bfd.multi_hop`



<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
