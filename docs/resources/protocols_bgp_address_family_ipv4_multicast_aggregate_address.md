---
page_title: "vyos_protocols_bgp_address_family_ipv4_multicast_aggregate_address Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Border Gateway Protocol (BGP)⯯BGP address-family parameters⯯Multicast IPv4 BGP settings⯯BGP aggregate network/prefix
---

# vyos_protocols_bgp_address_family_ipv4_multicast_aggregate_address (Resource)
<center>

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP address-family parameters  
⯯  
Multicast IPv4 BGP settings  
⯯  
**BGP aggregate network/prefix**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `as_set` (Boolean) Generate AS-set path information for this aggregate address
- `route_map` (String) Specify route-map name to use

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Route map name  |
- `summary_only` (Boolean) Announce the aggregate summary network only
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `aggregate_address` (String) BGP aggregate network/prefix

    |Format   &emsp;|Description                   |
    |-----------|--------------------------------|
    |ipv4net  &emsp;|BGP aggregate network/prefix  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
