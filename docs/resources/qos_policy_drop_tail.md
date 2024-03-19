---
page_title: "vyos_qos_policy_drop_tail Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)
  ⯯
  Service Policy definitions
  ⯯
  Packet limited First In, First Out queue
---

# vyos_qos_policy_drop_tail (Resource)
<center>

Quality of Service (QoS)
⯯
Service Policy definitions
⯯
**Packet limited First In, First Out queue**


</center>

## Schema

### Required

- `drop_tail_id` (String) Packet limited First In, First Out queue

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `queue_limit` (Number) Maximum queue size

    &emsp;|Format        &emsp;|Description            |
    |----------------|-------------------------|
    &emsp;|1-4294967295  &emsp;|Queue size in packets  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
