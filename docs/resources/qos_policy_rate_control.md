---
page_title: "vyos_qos_policy_rate_control Resource - vyos"

subcategory: "Qos"

description: |- 
  Quality of Service (QoS)⯯Service Policy definitions⯯Rate limiting policy (Token Bucket Filter)
---

# vyos_qos_policy_rate_control (Resource)
<center>

Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
**Rate limiting policy (Token Bucket Filter)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `rate_control_id` (String) Rate limiting policy (Token Bucket Filter)

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |

### Optional

- `bandwidth` (String) Available bandwidth for this policy

    &emsp;|Format        &emsp;|Description                         |
    |----------------|--------------------------------------|
    &emsp;|&lt;number&gt;      &emsp;|Bits per second                     |
    &emsp;|&lt;number&gt;bit   &emsp;|Bits per second                     |
    &emsp;|&lt;number&gt;kbit  &emsp;|Kilobits per second                 |
    &emsp;|&lt;number&gt;mbit  &emsp;|Megabits per second                 |
    &emsp;|&lt;number&gt;gbit  &emsp;|Gigabits per second                 |
    &emsp;|&lt;number&gt;tbit  &emsp;|Terabits per second                 |
    &emsp;|&lt;number&gt;%%    &emsp;|Percentage of interface link speed  |
- `burst` (String) Burst size for this class

    &emsp;|Format            &emsp;|Description                             |
    |--------------------|------------------------------------------|
    &emsp;|&lt;number&gt;          &emsp;|Bytes                                   |
    &emsp;|&lt;number&gt;&lt;suffix&gt;  &emsp;|Bytes with scaling suffix (kb, mb, gb)  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `latency` (String) Maximum latency

    &emsp;|Format    &emsp;|Description           |
    |------------|------------------------|
    &emsp;|&lt;number&gt;  &emsp;|Time in milliseconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
