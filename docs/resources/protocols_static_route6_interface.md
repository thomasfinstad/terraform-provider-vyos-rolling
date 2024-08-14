---
page_title: "vyos_protocols_static_route6_interface Resource - vyos"

subcategory: "Protocols"

description: |- 
  Routing protocols⯯Static Routing⯯Static IPv6 route⯯IPv6 gateway interface name
---

# vyos_protocols_static_route6_interface (Resource)
<center>

Routing protocols  
⯯  
Static Routing  
⯯  
Static IPv6 route  
⯯  
**IPv6 gateway interface name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `disable` (Boolean) Disable instance
- `distance` (Number) Distance for this route

    |Format  &emsp;|Description              |
    |----------|---------------------------|
    |1-255   &emsp;|Distance for this route  |
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

- `interface` (String) IPv6 gateway interface name

    |Format  &emsp;|Description             |
    |----------|--------------------------|
    |txt     &emsp;|Gateway interface name  |
- `route6` (String) Static IPv6 route

    |Format   &emsp;|Description        |
    |-----------|---------------------|
    |ipv6net  &emsp;|IPv6 static route  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
