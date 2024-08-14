---
page_title: "vyos_protocols_mpls_ldp_targeted_neighbor_ipv6 Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Multiprotocol Label Switching (MPLS)⯯Label Distribution Protocol (LDP)⯯Targeted LDP neighbor/session parameters⯯Targeted IPv6 neighbor/session parameters
---

# vyos_protocols_mpls_ldp_targeted_neighbor_ipv6 (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Multiprotocol Label Switching (MPLS)  
⯯  
Label Distribution Protocol (LDP)  
⯯  
Targeted LDP neighbor/session parameters  
⯯  
**Targeted IPv6 neighbor/session parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `address` (List of String) Neighbor/session address

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|ipv6    &emsp;|Neighbor/session address  |
- `enable` (Boolean) Accept and respond to targeted hellos
- `hello_holdtime` (Number) Hello hold time

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Time in seconds  |
- `hello_interval` (Number) Hello interval

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Time in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  