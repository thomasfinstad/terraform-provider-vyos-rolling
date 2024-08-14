---
page_title: "vyos_protocols_bgp_address_family_ipv6_labeled_unicast_network Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Border Gateway Protocol (BGP)⯯BGP address-family parameters⯯Labeled Unicast IPv6 BGP settings⯯Import BGP network/prefix into labeled unicast IPv6 RIB
---

# vyos_protocols_bgp_address_family_ipv6_labeled_unicast_network (Resource)
<center>

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP address-family parameters  
⯯  
Labeled Unicast IPv6 BGP settings  
⯯  
**Import BGP network/prefix into labeled unicast IPv6 RIB**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `backdoor` (Boolean) Use BGP network/prefix as a backdoor route
- `route_map` (String) Specify route-map name to use

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Route map name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `network` (String) Import BGP network/prefix into labeled unicast IPv6 RIB

    |Format   &emsp;|Description                              |
    |-----------|-------------------------------------------|
    |ipv6net  &emsp;|Labeled Unicast IPv6 BGP network/prefix  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
