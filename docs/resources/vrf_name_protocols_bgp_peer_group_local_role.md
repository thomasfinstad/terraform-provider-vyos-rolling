---
page_title: "vyos_vrf_name_protocols_bgp_peer_group_local_role Resource - vyos"
subcategory: "vrf"
description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Border Gateway Protocol (BGP)⯯Name of peer-group⯯Local role for BGP neighbor (RFC9234)
---

# vyos_vrf_name_protocols_bgp_peer_group_local_role (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Border Gateway Protocol (BGP)  
⯯  
Name of peer-group  
⯯  
**Local role for BGP neighbor (RFC9234)**


</center>

## Schema

### Required

- `local_role_id` (String) Local role for BGP neighbor (RFC9234)

    &emsp;|Format     &emsp;|Description             |
    |-------------|--------------------------|
    &emsp;|customer   &emsp;|Using Transit           |
    &emsp;|peer       &emsp;|Public/Private Peering  |
    &emsp;|provider   &emsp;|Providing Transit       |
    &emsp;|rs-client  &emsp;|RS Client               |
    &emsp;|rs-server  &emsp;|Route Server            |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `peer_group_id` (String) Name of peer-group

### Optional

- `strict` (Boolean) Neighbor must send this exact capability, otherwise a role missmatch notification will be sent
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
