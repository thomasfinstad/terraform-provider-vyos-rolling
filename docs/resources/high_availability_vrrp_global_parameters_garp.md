---
page_title: "vyos_high_availability_vrrp_global_parameters_garp Resource - terraform-provider-vyos"
subcategory: "high"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  High availability settings
  ⯯Virtual Router Redundancy Protocol settings⯯VRRP global parameters⯯Gratuitous ARP parameters
---

# vyos_high_availability_vrrp_global_parameters_garp (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

	High availability settings
⯯
Virtual Router Redundancy Protocol settings
⯯
VRRP global parameters
⯯
**Gratuitous ARP parameters**


</center>

## Schema

### Optional

- `interval` (String) Interval between Gratuitous ARP

    &emsp;|Format        &emsp;|Description                                   |
    |----------------|------------------------------------------------|
    &emsp;|&lt;0.000-1000&gt;  &emsp;|Interval in seconds, resolution microseconds  |
- `master_delay` (Number) Delay for second set of gratuitous ARPs after transition to master

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|1-1000  &emsp;|Delay in seconds  |
- `master_refresh` (Number) Minimum time interval for refreshing gratuitous ARPs while beeing master

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|0       &emsp;|No refresh           |
    &emsp;|1-255   &emsp;|Interval in seconds  |
- `master_refresh_repeat` (Number) Number of gratuitous ARP messages to send at a time while beeing master

    &emsp;|Format  &emsp;|Description                        |
    |----------|-------------------------------------|
    &emsp;|1-255   &emsp;|Number of gratuitous ARP messages  |
- `master_repeat` (Number) Number of gratuitous ARP messages to send at a time after transition to master

    &emsp;|Format  &emsp;|Description                        |
    |----------|-------------------------------------|
    &emsp;|1-255   &emsp;|Number of gratuitous ARP messages  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
