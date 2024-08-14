---
page_title: "vyos_protocols_bgp_parameters_tcp_keepalive Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Border Gateway Protocol (BGP)⯯BGP parameters⯯TCP keepalive parameters
---

# vyos_protocols_bgp_parameters_tcp_keepalive (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP parameters  
⯯  
**TCP keepalive parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `idle` (Number) TCP keepalive idle time

    |Format   &emsp;|Description           |
    |-----------|------------------------|
    |1-65535  &emsp;|Idle time in seconds  |
- `interval` (Number) TCP keepalive interval

    |Format   &emsp;|Description          |
    |-----------|-----------------------|
    |1-65535  &emsp;|Interval in seconds  |
- `probes` (Number) TCP keepalive maximum probes

    |Format  &emsp;|Description     |
    |----------|------------------|
    |1-30    &emsp;|Maximum probes  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
