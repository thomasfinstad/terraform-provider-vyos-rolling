---
page_title: "vyos_qos_policy_fair_queue Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)
  ⯯
  Service Policy definitions
  ⯯
  Stochastic Fairness Queueing
---

# vyos_qos_policy_fair_queue (Resource)
<center>

Quality of Service (QoS)
⯯
Service Policy definitions
⯯
**Stochastic Fairness Queueing**


</center>

## Schema

### Required

- `fair_queue_id` (String) Stochastic Fairness Queueing

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `hash_interval` (Number) Interval in seconds for queue algorithm perturbation

    &emsp;|Format  &emsp;|Description                                                         |
    |----------|----------------------------------------------------------------------|
    &emsp;|0       &emsp;|No perturbation                                                     |
    &emsp;|1-127   &emsp;|Interval in seconds for queue algorithm perturbation (advised: 10)  |
- `queue_limit` (Number) Upper limit of the SFQ

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|1-127   &emsp;|Queue size in packets  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
