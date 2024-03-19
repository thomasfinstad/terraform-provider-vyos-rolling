---
page_title: "vyos_qos_policy_random_detect_precedence Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)
  ⯯
  Service Policy definitions
  ⯯
  Weighted Random Early Detect policy
  ⯯
  IP precedence
---

# vyos_qos_policy_random_detect_precedence (Resource)
<center>

Quality of Service (QoS)
⯯
Service Policy definitions
⯯
Weighted Random Early Detect policy
⯯
**IP precedence**


</center>

## Schema

### Required

- `precedence_id` (Number) IP precedence

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|0-7     &emsp;|IP precedence value  |
- `random_detect_id` (String) Weighted Random Early Detect policy

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |

### Optional

- `average_packet` (Number) Average packet size (bytes)

    &emsp;|Format    &emsp;|Description                   |
    |------------|--------------------------------|
    &emsp;|16-10240  &emsp;|Average packet size in bytes  |
- `mark_probability` (String) Mark probability for this precedence

    &emsp;|Format    &emsp;|Description          |
    |------------|-----------------------|
    &emsp;|&lt;number&gt;  &emsp;|Numeric value (1/N)  |
- `maximum_threshold` (Number) Maximum threshold for random detection

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|0-4096  &emsp;|Maximum Threshold in packets  |
- `minimum_threshold` (Number) Minimum  threshold for random detection

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|0-4096  &emsp;|Maximum Threshold in packets  |
- `queue_limit` (Number) Maximum queue size

    &emsp;|Format        &emsp;|Description            |
    |----------------|-------------------------|
    &emsp;|1-4294967295  &emsp;|Queue size in packets  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
