---
page_title: "vyos_policy_prefix_list Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy
  ⯯
  IP prefix-list filter
---

# vyos_policy_prefix_list (Resource)
<center>

Routing policy
⯯
**IP prefix-list filter**


</center>

## Schema

### Required

- `prefix_list_id` (String) IP prefix-list filter

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
