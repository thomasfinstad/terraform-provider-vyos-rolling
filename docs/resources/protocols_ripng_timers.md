---
page_title: "vyos_protocols_ripng_timers Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Routing Information Protocol (RIPng) parameters⯯RIPng timer values
---

# vyos_protocols_ripng_timers (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Routing Information Protocol (RIPng) parameters  
⯯  
**RIPng timer values**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `garbage_collection` (Number) Garbage collection timer

    &emsp;|Format        &emsp;|Description             |
    |----------------|--------------------------|
    &emsp;|5-2147483647  &emsp;|Garbage colletion time  |
- `timeout` (Number) Routing information timeout timer

    &emsp;|Format        &emsp;|Description                        |
    |----------------|-------------------------------------|
    &emsp;|5-2147483647  &emsp;|Routing information timeout timer  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `update` (Number) Routing table update timer

    &emsp;|Format        &emsp;|Description                            |
    |----------------|-----------------------------------------|
    &emsp;|5-2147483647  &emsp;|Routing table update timer in seconds  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
