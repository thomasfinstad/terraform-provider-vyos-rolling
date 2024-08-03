---
page_title: "vyos_system_conntrack_tcp Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯Connection Tracking Engine Options⯯TCP options
---

# vyos_system_conntrack_tcp (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
Connection Tracking Engine Options  
⯯  
**TCP options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `half_open_connections` (Number) Maximum number of TCP half-open connections

    &emsp;|Format        &emsp;|Description                            |
    |----------------|-----------------------------------------|
    &emsp;|1-2147483647  &emsp;|Generic connection timeout in seconds  |
- `loose` (String) Policy to track previously established connections

    &emsp;|Format   &emsp;|Description                                                  |
    |-----------|---------------------------------------------------------------|
    &emsp;|enable   &emsp;|Allow tracking of previously established connections         |
    &emsp;|disable  &emsp;|Do not allow tracking of previously established connections  |
- `max_retrans` (Number) Maximum number of packets that can be retransmitted without received an ACK

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-255   &emsp;|Number of packets to be retransmitted  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
