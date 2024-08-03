---
page_title: "vyos_protocols_bgp_neighbor_local_as Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Border Gateway Protocol (BGP)⯯BGP neighbor⯯Specify alternate ASN for this BGP process
---

# vyos_protocols_bgp_neighbor_local_as (Resource)
<center>

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP neighbor  
⯯  
**Specify alternate ASN for this BGP process**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `no_prepend` (Attributes) Disable prepending local-as from/to updates for eBGP peers (see [below for nested schema](#nestedatt--no_prepend))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `local_as` (Number) Specify alternate ASN for this BGP process

    &emsp;|Format        &emsp;|Description                     |
    |----------------|----------------------------------|
    &emsp;|1-4294967294  &emsp;|Autonomous System Number (ASN)  |
- `neighbor` (String) BGP neighbor

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|ipv4    &emsp;|BGP neighbor IP address    |
    &emsp;|ipv6    &emsp;|BGP neighbor IPv6 address  |
    &emsp;|txt     &emsp;|Interface name             |


&lt;a id=&#34;nestedatt--no_prepend&#34;&gt;&lt;/a&gt;
### Nested Schema for `no_prepend`

Optional:

- `replace_as` (Boolean) Prepend only local-as from/to updates for eBGP peers


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
