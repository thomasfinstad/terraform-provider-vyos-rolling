---
page_title: "vyos_qos_policy_priority_queue_class Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)
  ⯯
  Service Policy definitions
  ⯯
  Priority queuing based policy
  ⯯
  Class Handle
---

# vyos_qos_policy_priority_queue_class (Resource)
<center>

Quality of Service (QoS)
⯯
Service Policy definitions
⯯
Priority queuing based policy
⯯
**Class Handle**


</center>

## Schema

### Required

- `class_id` (Number) Class Handle

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1-7     &emsp;|Priority     |
- `priority_queue_id` (String) Priority queuing based policy

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
- `queue_limit` (Number) Maximum queue size

    &emsp;|Format        &emsp;|Description            |
    |----------------|-------------------------|
    &emsp;|1-4294967295  &emsp;|Queue size in packets  |
- `queue_type` (String) Queue type for default traffic

    &emsp;|Format         &emsp;|Description                   |
    |-----------------|--------------------------------|
    &emsp;|drop-tail      &emsp;|First-In-First-Out (FIFO)     |
    &emsp;|fair-queue     &emsp;|Stochastic Fair Queue (SFQ)   |
    &emsp;|fq-codel       &emsp;|Fair Queue Codel              |
    &emsp;|priority       &emsp;|Priority queuing              |
    &emsp;|random-detect  &emsp;|Random Early Detection (RED)  |
- `target` (Number) Acceptable minimum standing/persistent queue delay

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|u32     &emsp;|Queue delay in milliseconds  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
