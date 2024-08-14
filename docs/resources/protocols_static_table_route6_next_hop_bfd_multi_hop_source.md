---
page_title: "vyos_protocols_static_table_route6_next_hop_bfd_multi_hop_source Resource - vyos"

subcategory: "Protocols"

description: |- 
  Routing protocols⯯Static Routing⯯Policy route table number⯯Static IPv6 route⯯IPv6 gateway address⯯BFD monitoring⯯Use BFD multi hop session⯯Use source for BFD session
---

# vyos_protocols_static_table_route6_next_hop_bfd_multi_hop_source (Resource)
<center>

Routing protocols  
⯯  
Static Routing  
⯯  
Policy route table number  
⯯  
Static IPv6 route  
⯯  
IPv6 gateway address  
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

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `profile` (String) Use settings from BFD profile

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |txt     &emsp;|BFD profile name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

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
- `source` (String) Use source for BFD session

    |Format  &emsp;|Description          |
    |----------|-----------------------|
    |ipv4    &emsp;|IPv4 source address  |
    |ipv6    &emsp;|IPv6 source address  |
- `table` (Number) Policy route table number

    |Format  &emsp;|Description                |
    |----------|-----------------------------|
    |1-200   &emsp;|Policy route table number  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
