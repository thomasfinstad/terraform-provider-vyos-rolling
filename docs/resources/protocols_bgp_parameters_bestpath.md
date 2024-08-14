---
page_title: "vyos_protocols_bgp_parameters_bestpath Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Border Gateway Protocol (BGP)⯯BGP parameters⯯Default bestpath selection mechanism
---

# vyos_protocols_bgp_parameters_bestpath (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP parameters  
⯯  
**Default bestpath selection mechanism**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `bandwidth` (String) Link Bandwidth attribute

    |Format                      &emsp;|Description                                                            |
    |------------------------------|-------------------------------------------------------------------------|
    |default-weight-for-missing  &emsp;|Assign low default weight (1) to paths not having link bandwidth       |
    |ignore                      &emsp;|Ignore link bandwidth (do regular ECMP, not weighted)                  |
    |skip-missing                &emsp;|Ignore paths without link bandwidth for ECMP (if other paths have it)  |
- `compare_routerid` (Boolean) Compare the router-id for identical EBGP paths
- `med` (List of String) MED attribute comparison parameters

    |Format            &emsp;|Description                                              |
    |--------------------|-----------------------------------------------------------|
    |confed            &emsp;|Compare MEDs among confederation paths                   |
    |missing-as-worst  &emsp;|Treat missing route as a MED as the least preferred one  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
