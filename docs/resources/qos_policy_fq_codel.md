---
page_title: "vyos_qos_policy_fq_codel Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)⯯Service Policy definitions⯯Fair Queuing (FQ) with Controlled Delay (CoDel)
---

# vyos_qos_policy_fq_codel (Resource)
<center>

Quality of Service (QoS)
⯯
Service Policy definitions
⯯
**Fair Queuing (FQ) with Controlled Delay (CoDel)**


</center>

## Schema

### Required

- `fq_codel_id` (String) Fair Queuing (FQ) with Controlled Delay (CoDel)

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |

### Optional

- `codel_quantum` (Number) Deficit in the fair queuing algorithm

    &emsp;|Format     &emsp;|Description                        |
    |-------------|-------------------------------------|
    &emsp;|0-1048576  &emsp;|Number of bytes used as &#39;deficit&#39;  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `flows` (Number) Number of flows into which the incoming packets are classified

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65536  &emsp;|Number of flows  |
- `interval` (Number) Interval used to measure the delay

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|u32     &emsp;|Interval in milliseconds  |
- `queue_limit` (Number) Upper limit of the queue

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|2-10999  &emsp;|Queue size in packets  |
- `target` (Number) Acceptable minimum standing/persistent queue delay

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|u32     &emsp;|Queue delay in milliseconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
