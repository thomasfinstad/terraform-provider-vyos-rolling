---
page_title: "vyos_qos_interface Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)
  ⯯
  Interface to apply QoS policy
---

# vyos_qos_interface (Resource)
<center>

Quality of Service (QoS)
⯯
**Interface to apply QoS policy**


</center>

## Schema

### Required

- `interface_id` (String) Interface to apply QoS policy

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |

### Optional

- `egress` (String) Interface egress traffic policy

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|QoS policy to use  |
- `ingress` (String) Interface ingress traffic policy

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|QoS policy to use  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
