---
page_title: "vyos_protocols_mpls_ldp_import_ipv6_import_filter Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Multiprotocol Label Switching (MPLS)⯯Label Distribution Protocol (LDP)⯯Import parameters⯯IPv6 parameters⯯Forwarding equivalence class import filter
---

# vyos_protocols_mpls_ldp_import_ipv6_import_filter (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Multiprotocol Label Switching (MPLS)  
⯯  
Label Distribution Protocol (LDP)  
⯯  
Import parameters  
⯯  
IPv6 parameters  
⯯  
**Forwarding equivalence class import filter**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `filter_access_list6` (Number) Access-list6 number to apply FEC filtering

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |1-2699  &emsp;|Access list number  |
- `neighbor_access_list6` (Number) Access-list6 number for IPv6 neighbor selection to apply filtering

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |1-2699  &emsp;|Access list number  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
