---
page_title: "vyos_protocols_ospf_timers_throttle_spf Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Open Shortest Path First (OSPF)⯯Adjust routing timers⯯Throttling adaptive timers⯯OSPF SPF timers
---

# vyos_protocols_ospf_timers_throttle_spf (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Open Shortest Path First (OSPF)  
⯯  
Adjust routing timers  
⯯  
Throttling adaptive timers  
⯯  
**OSPF SPF timers**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `delay` (Number) Delay from the first change received to SPF calculation

    |Format    &emsp;|Description            |
    |------------|-------------------------|
    |0-600000  &emsp;|Delay in milliseconds  |
- `initial_holdtime` (Number) Initial hold time between consecutive SPF calculations

    |Format    &emsp;|Description                        |
    |------------|-------------------------------------|
    |0-600000  &emsp;|Initial hold time in milliseconds  |
- `max_holdtime` (Number) Maximum hold time

    |Format    &emsp;|Description                    |
    |------------|---------------------------------|
    |0-600000  &emsp;|Max hold time in milliseconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
