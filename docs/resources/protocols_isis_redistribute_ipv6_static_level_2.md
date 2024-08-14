---
page_title: "vyos_protocols_isis_redistribute_ipv6_static_level_2 Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Intermediate System to Intermediate System (IS-IS)⯯Redistribute information from another routing protocol⯯Redistribute IPv6 routes⯯Redistribute static routes into IS-IS⯯Redistribute into level-2
---

# vyos_protocols_isis_redistribute_ipv6_static_level_2 (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
Redistribute information from another routing protocol  
⯯  
Redistribute IPv6 routes  
⯯  
Redistribute static routes into IS-IS  
⯯  
**Redistribute into level-2**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  