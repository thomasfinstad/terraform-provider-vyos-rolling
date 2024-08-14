---
page_title: "vyos_protocols_bgp_interface Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Border Gateway Protocol (BGP)⯯Configure interface related parameters, e.g. MPLS
---

# vyos_protocols_bgp_interface (Resource)
<center>

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
**Configure interface related parameters, e.g. MPLS**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `mpls` (Attributes) MPLS options (see [below for nested schema](#nestedatt--mpls))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Configure interface related parameters, e.g. MPLS

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |


<a id="nestedatt--mpls"></a>
### Nested Schema for `mpls`

Optional:

- `forwarding` (Boolean) Enable MPLS forwarding for eBGP directly connected peers


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
