---
page_title: "vyos_qos_policy_priority_queue Resource - vyos"

subcategory: "Qos"

description: |- 
  Quality of Service (QoS)⯯Service Policy definitions⯯Priority queuing based policy
---

# vyos_qos_policy_priority_queue (Resource)
<center>

Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
**Priority queuing based policy**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `default` (Attributes) Default policy (see [below for nested schema](#nestedatt--default))
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `priority_queue` (String) Priority queuing based policy

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Policy name  |


<a id="nestedatt--default"></a>
### Nested Schema for `default`

Optional:

- `codel_quantum` (Number) Deficit in the fair queuing algorithm

    |Format     &emsp;|Description                        |
    |-------------|-------------------------------------|
    |0-1048576  &emsp;|Number of bytes used as &#39;deficit&#39;  |
- `flows` (Number) Number of flows into which the incoming packets are classified

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |1-65536  &emsp;|Number of flows  |
- `interval` (Number) Interval used to measure the delay

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |u32     &emsp;|Interval in milliseconds  |
- `queue_limit` (Number) Maximum queue size

    |Format        &emsp;|Description            |
    |----------------|-------------------------|
    |1-4294967295  &emsp;|Queue size in packets  |
- `queue_type` (String) Queue type for default traffic

    |Format         &emsp;|Description                   |
    |-----------------|--------------------------------|
    |drop-tail      &emsp;|First-In-First-Out (FIFO)     |
    |fair-queue     &emsp;|Stochastic Fair Queue (SFQ)   |
    |fq-codel       &emsp;|Fair Queue Codel              |
    |priority       &emsp;|Priority queuing              |
    |random-detect  &emsp;|Random Early Detection (RED)  |
- `target` (Number) Acceptable minimum standing/persistent queue delay

    |Format  &emsp;|Description                  |
    |----------|-------------------------------|
    |u32     &emsp;|Queue delay in milliseconds  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
