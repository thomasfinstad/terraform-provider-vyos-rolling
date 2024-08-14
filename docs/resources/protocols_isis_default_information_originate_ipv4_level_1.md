---
page_title: "vyos_protocols_isis_default_information_originate_ipv4_level_1 Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Intermediate System to Intermediate System (IS-IS)⯯Control distribution of default information⯯Distribute a default route⯯Distribute default route for IPv4⯯Distribute default route into level-1
---

# vyos_protocols_isis_default_information_originate_ipv4_level_1 (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
Control distribution of default information  
⯯  
Distribute a default route  
⯯  
Distribute default route for IPv4  
⯯  
**Distribute default route into level-1**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `always` (Boolean) Always advertise default route
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