---
page_title: "vyos_qos_policy_network_emulator Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)
  ⯯
  Service Policy definitions
  ⯯
  Network emulator policy
---

# vyos_qos_policy_network_emulator (Resource)
<center>

Quality of Service (QoS)
⯯
Service Policy definitions
⯯
**Network emulator policy**


</center>

## Schema

### Required

- `network_emulator_id` (String) Network emulator policy

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
- `corruption` (String) Introducing error in a random position for chosen percent of packets

    &emsp;|Format    &emsp;|Description                     |
    |------------|----------------------------------|
    &emsp;|&lt;number&gt;  &emsp;|Percentage of packets affected  |
- `delay` (String) Adds delay to packets outgoing to chosen network interface

    &emsp;|Format    &emsp;|Description           |
    |------------|------------------------|
    &emsp;|&lt;number&gt;  &emsp;|Time in milliseconds  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `duplicate` (String) Cosen percent of packets is duplicated before queuing them

    &emsp;|Format    &emsp;|Description                     |
    |------------|----------------------------------|
    &emsp;|&lt;number&gt;  &emsp;|Percentage of packets affected  |
- `loss` (String) Add independent loss probability to the packets outgoing to chosen network interface

    &emsp;|Format    &emsp;|Description                     |
    |------------|----------------------------------|
    &emsp;|&lt;number&gt;  &emsp;|Percentage of packets affected  |
- `queue_limit` (Number) Maximum queue size

    &emsp;|Format        &emsp;|Description            |
    |----------------|-------------------------|
    &emsp;|1-4294967295  &emsp;|Queue size in packets  |
- `reordering` (String) Emulated packet reordering percentage

    &emsp;|Format    &emsp;|Description                     |
    |------------|----------------------------------|
    &emsp;|&lt;number&gt;  &emsp;|Percentage of packets affected  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
