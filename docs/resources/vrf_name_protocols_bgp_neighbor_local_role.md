---
page_title: "vyos_vrf_name_protocols_bgp_neighbor_local_role Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Border Gateway Protocol (BGP)⯯BGP neighbor⯯Local role for BGP neighbor (RFC9234)
---

# vyos_vrf_name_protocols_bgp_neighbor_local_role (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP neighbor  
⯯  
**Local role for BGP neighbor (RFC9234)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `strict` (Boolean) Neighbor must send this exact capability, otherwise a role missmatch notification will be sent
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `local_role` (String) Local role for BGP neighbor (RFC9234)

    &emsp;|Format     &emsp;|Description             |
    |-------------|--------------------------|
    &emsp;|customer   &emsp;|Using Transit           |
    &emsp;|peer       &emsp;|Public/Private Peering  |
    &emsp;|provider   &emsp;|Providing Transit       |
    &emsp;|rs-client  &emsp;|RS Client               |
    &emsp;|rs-server  &emsp;|Route Server            |
- `name` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `neighbor` (String) BGP neighbor

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|ipv4    &emsp;|BGP neighbor IP address    |
    &emsp;|ipv6    &emsp;|BGP neighbor IPv6 address  |
    &emsp;|txt     &emsp;|Interface name             |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
