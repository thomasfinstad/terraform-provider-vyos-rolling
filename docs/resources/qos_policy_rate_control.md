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

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `bandwidth` (String) Available bandwidth for this policy

    |Format        &emsp;|Description                         |
    |----------------|--------------------------------------|
    |&lt;number&gt;      &emsp;|Bits per second                     |
    |&lt;number&gt;bit   &emsp;|Bits per second                     |
    |&lt;number&gt;kbit  &emsp;|Kilobits per second                 |
    |&lt;number&gt;mbit  &emsp;|Megabits per second                 |
    |&lt;number&gt;gbit  &emsp;|Gigabits per second                 |
    |&lt;number&gt;tbit  &emsp;|Terabits per second                 |
    |&lt;number&gt;%%    &emsp;|Percentage of interface link speed  |
- `burst` (String) Burst size for this class

    |Format            &emsp;|Description                             |
    |--------------------|------------------------------------------|
    |&lt;number&gt;          &emsp;|Bytes                                   |
    |&lt;number&gt;&lt;suffix&gt;  &emsp;|Bytes with scaling suffix (kb, mb, gb)  |
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `latency` (String) Maximum latency

    |Format    &emsp;|Description           |
    |------------|------------------------|
    |&lt;number&gt;  &emsp;|Time in milliseconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `rate_control` (String) Rate limiting policy (Token Bucket Filter)

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Policy name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
