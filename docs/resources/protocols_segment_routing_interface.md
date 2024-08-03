---
page_title: "vyos_protocols_segment_routing_interface Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Segment Routing⯯Interface specific Segment Routing options
---

# vyos_protocols_segment_routing_interface (Resource)
<center>

*protocols*  
⯯  
Segment Routing  
⯯  
**Interface specific Segment Routing options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `srv6` (Attributes) Accept SR-enabled IPv6 packets on this interface (see [below for nested schema](#nestedatt--srv6))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface specific Segment Routing options

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |


&lt;a id=&#34;nestedatt--srv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `srv6`

Optional:

- `hmac` (String) Define HMAC policy for ingress SR-enabled packets on this interface

    &emsp;|Format  &emsp;|Description                                              |
    |----------|-----------------------------------------------------------|
    &emsp;|accept  &emsp;|Accept packets without HMAC, validate packets with HMAC  |
    &emsp;|drop    &emsp;|Drop packets without HMAC, validate packets with HMAC    |
    &emsp;|ignore  &emsp;|Ignore HMAC field.                                       |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
