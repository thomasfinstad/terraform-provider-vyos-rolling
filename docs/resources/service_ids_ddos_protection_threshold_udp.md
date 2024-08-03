---
page_title: "vyos_service_ids_ddos_protection_threshold_udp Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Intrusion Detection System⯯FastNetMon detection and protection parameters⯯Attack limits thresholds⯯UDP threshold
---

# vyos_service_ids_ddos_protection_threshold_udp (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Intrusion Detection System  
⯯  
FastNetMon detection and protection parameters  
⯯  
Attack limits thresholds  
⯯  
**UDP threshold**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `fps` (Number) Flows per second

    &emsp;|Format        &emsp;|Description       |
    |----------------|--------------------|
    &emsp;|0-4294967294  &emsp;|Flows per second  |
- `mbps` (Number) Megabits per second

    &emsp;|Format        &emsp;|Description          |
    |----------------|-----------------------|
    &emsp;|0-4294967294  &emsp;|Megabits per second  |
- `pps` (Number) Packets per second

    &emsp;|Format        &emsp;|Description         |
    |----------------|----------------------|
    &emsp;|0-4294967294  &emsp;|Packets per second  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
