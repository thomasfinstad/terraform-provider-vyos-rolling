---
page_title: "vyos_protocols_pim Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Protocol Independent Multicast (PIM) and IGMP
---

# vyos_protocols_pim (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
**Protocol Independent Multicast (PIM) and IGMP**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `join_prune_interval` (Number) Join prune send interval

    &emsp;|Format   &emsp;|Description          |
    |-----------|-----------------------|
    &emsp;|1-65535  &emsp;|Interval in seconds  |
- `keep_alive_timer` (Number) Keep alive Timer

    &emsp;|Format   &emsp;|Description                  |
    |-----------|-------------------------------|
    &emsp;|1-65535  &emsp;|Keep alive Timer in seconds  |
- `no_v6_secondary` (Boolean) Disable IPv6 secondary address in hello packets
- `packets` (Number) Packets to process at once

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|1-255   &emsp;|Number of packets  |
- `register_suppress_time` (Number) Register suppress timer

    &emsp;|Format   &emsp;|Description       |
    |-----------|--------------------|
    &emsp;|1-65535  &emsp;|Timer in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  