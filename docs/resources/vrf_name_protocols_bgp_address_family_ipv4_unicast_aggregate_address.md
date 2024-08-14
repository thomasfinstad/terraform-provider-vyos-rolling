---
page_title: "vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_aggregate_address Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Border Gateway Protocol (BGP)⯯BGP address-family parameters⯯IPv4 BGP settings⯯BGP aggregate network
---

# vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_aggregate_address (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP address-family parameters  
⯯  
IPv4 BGP settings  
⯯  
**BGP aggregate network**


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

- `aggregate_address` (String) BGP aggregate network

    |Format   &emsp;|Description            |
    |-----------|-------------------------|
    |ipv4net  &emsp;|BGP aggregate network  |
- `name` (String) Virtual Routing and Forwarding instance

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
