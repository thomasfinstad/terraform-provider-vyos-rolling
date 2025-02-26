---
page_title: "vyos_qos_policy_round_robin_class Resource - vyos"

subcategory: "Qos"

description: |-
  Quality of Service (QoS)⯯Service Policy definitions⯯Deficit Round Robin Scheduler⯯Class ID
---

# vyos_qos_policy_round_robin_class (Resource)
<center>


Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
Deficit Round Robin Scheduler  
⯯  
**Class ID**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_qos_policy_round_robin_class (Resource)](#vyos_qos_policy_round_robin_class-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [codel_quantum](#codel_quantum)
      - [description](#description)
      - [flows](#flows)
      - [interval](#interval)
      - [match_group](#match_group)
      - [quantum](#quantum)
      - [queue_limit](#queue_limit)
      - [queue_type](#queue_type)
      - [target](#target)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### codel_quantum
- `codel_quantum` (Number) Deficit in the fair queuing algorithm

    |  Format     &emsp;|  Description                        |
    |-------------|-------------------------------------|
    |  0-1048576  &emsp;|  Number of bytes used as &#39;deficit&#39;  |
#### description
- `description` (String) Description

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  txt     &emsp;|  Description  |
#### flows
- `flows` (Number) Number of flows into which the incoming packets are classified

    |  Format   &emsp;|  Description      |
    |-----------|-------------------|
    |  1-65536  &emsp;|  Number of flows  |
#### interval
- `interval` (Number) Interval used to measure the delay

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  u32     &emsp;|  Interval in milliseconds  |
#### match_group
- `match_group` (List of String) Filter group for QoS policy

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  txt     &emsp;|  Match group name  |
#### quantum
- `quantum` (Number) Packet scheduling quantum

    |  Format        &emsp;|  Description                        |
    |----------------|-------------------------------------|
    |  1-4294967295  &emsp;|  Packet scheduling quantum (bytes)  |
#### queue_limit
- `queue_limit` (Number) Maximum queue size

    |  Format        &emsp;|  Description            |
    |----------------|-------------------------|
    |  1-4294967295  &emsp;|  Queue size in packets  |
#### queue_type
- `queue_type` (String) Queue type for default traffic

    |  Format         &emsp;|  Description                   |
    |-----------------|--------------------------------|
    |  drop-tail      &emsp;|  First-In-First-Out (FIFO)     |
    |  fair-queue     &emsp;|  Stochastic Fair Queue (SFQ)   |
    |  fq-codel       &emsp;|  Fair Queue Codel              |
    |  priority       &emsp;|  Priority queuing              |
    |  random-detect  &emsp;|  Random Early Detection (RED)  |
#### target
- `target` (Number) Acceptable minimum standing/persistent queue delay

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  u32     &emsp;|  Queue delay in milliseconds  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `class` (Number) Class ID

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  1-4095  &emsp;|  Class Identifier  |
- `round_robin` (String) Deficit Round Robin Scheduler

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  txt     &emsp;|  Policy name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_qos_policy_round_robin_class.example "qos__policy__round-robin__<round-robin>__class__<class>"
```
