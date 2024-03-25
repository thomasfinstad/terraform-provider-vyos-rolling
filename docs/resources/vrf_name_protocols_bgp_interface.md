---
page_title: "vyos_vrf_name_protocols_bgp_interface Resource - vyos"
subcategory: "vrf"
description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Border Gateway Protocol (BGP)⯯Configure interface related parameters, e.g. MPLS
---

# vyos_vrf_name_protocols_bgp_interface (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Border Gateway Protocol (BGP)  
⯯  
**Configure interface related parameters, e.g. MPLS**


</center>

## Schema

### Required

- `interface_id` (String) Configure interface related parameters, e.g. MPLS

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Optional

- `mpls` (Attributes) MPLS options (see [below for nested schema](#nestedatt--mpls))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--mpls&#34;&gt;&lt;/a&gt;
### Nested Schema for `mpls`

Optional:

- `forwarding` (Boolean) Enable MPLS forwarding for eBGP directly connected peers


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
