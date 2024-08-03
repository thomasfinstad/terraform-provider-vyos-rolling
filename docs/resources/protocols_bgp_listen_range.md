---
page_title: "vyos_protocols_bgp_listen_range Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Border Gateway Protocol (BGP)⯯Listen for and accept BGP dynamic neighbors from range⯯BGP dynamic neighbors listen range
---

# vyos_protocols_bgp_listen_range (Resource)
<center>

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
Listen for and accept BGP dynamic neighbors from range  
⯯  
**BGP dynamic neighbors listen range**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `peer_group` (String) Peer group for this peer

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|Peer-group name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `range` (String) BGP dynamic neighbors listen range

    &emsp;|Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 dynamic neighbors listen range  |
    &emsp;|ipv6net  &emsp;|IPv6 dynamic neighbors listen range  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
