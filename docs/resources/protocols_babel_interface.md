---
page_title: "vyos_protocols_babel_interface Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Babel Routing Protocol⯯Interface name
---

# vyos_protocols_babel_interface (Resource)
<center>

*protocols*  
⯯  
Babel Routing Protocol  
⯯  
**Interface name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `channel` (String) Channel number for diversity routing

    &emsp;|Format           &emsp;|Description                                                                                                         |
    |-------------------|----------------------------------------------------------------------------------------------------------------------|
    &emsp;|1-254            &emsp;|Interfaces with a channel number interfere with interfering interfaces and interfaces with the same channel number  |
    &emsp;|interfering      &emsp;|Interfering interfaces are assumed to interfere with all other channels except non-interfering channels             |
    &emsp;|non-interfering  &emsp;|Non-interfering interfaces only interfere with themselves                                                           |
- `enable_timestamps` (Boolean) Enable timestamps with each Hello and IHU message in order to compute RTT values
- `hello_interval` (Number) Time between scheduled hellos

    &emsp;|Format     &emsp;|Description   |
    |-------------|----------------|
    &emsp;|20-655340  &emsp;|Milliseconds  |
- `max_rtt_penalty` (Number) Maximum additional cost due to RTT

    &emsp;|Format   &emsp;|Description                                            |
    |-----------|---------------------------------------------------------|
    &emsp;|0-65535  &emsp;|Milliseconds (0 to disable the use of RTT-based cost)  |
- `rtt_decay` (Number) Decay factor for exponential moving average of RTT samples

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|1-256   &emsp;|Decay factor, in units of 1/256  |
- `rtt_max` (Number) Maximum RTT

    &emsp;|Format   &emsp;|Description   |
    |-----------|----------------|
    &emsp;|1-65535  &emsp;|Milliseconds  |
- `rtt_min` (Number) Minimum RTT

    &emsp;|Format   &emsp;|Description   |
    |-----------|----------------|
    &emsp;|1-65535  &emsp;|Milliseconds  |
- `rxcost` (Number) Base receive cost for this interface

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|1-65534  &emsp;|Base receive cost  |
- `split_horizon` (String) Split horizon parameters

    &emsp;|Format   &emsp;|Description                                                     |
    |-----------|------------------------------------------------------------------|
    &emsp;|default  &emsp;|Enable on wired interfaces, and disable on wireless interfaces  |
    &emsp;|enable   &emsp;|Enable split horizon processing                                 |
    &emsp;|disable  &emsp;|Disable split horizon processing                                |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `type` (String) Interface type

    &emsp;|Format    &emsp;|Description                          |
    |------------|---------------------------------------|
    &emsp;|auto      &emsp;|Automatically detect interface type  |
    &emsp;|wired     &emsp;|Wired interface                      |
    &emsp;|wireless  &emsp;|Wireless interface                   |
- `update_interval` (Number) Time between scheduled updates

    &emsp;|Format     &emsp;|Description   |
    |-------------|----------------|
    &emsp;|20-655340  &emsp;|Milliseconds  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface name

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
