---
page_title: "vyos_qos_policy_rate_control Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)
  ⯯
  Service Policy definitions
  ⯯
  Rate limiting policy (Token Bucket Filter)
---

# vyos_qos_policy_rate_control (Resource)
<center>

Quality of Service (QoS)
⯯
Service Policy definitions
⯯
**Rate limiting policy (Token Bucket Filter)**


</center>

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

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
