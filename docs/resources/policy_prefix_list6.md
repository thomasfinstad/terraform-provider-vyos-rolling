---
page_title: "vyos_policy_prefix_list6 Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy
  ⯯
  IPv6 prefix-list filter
---

# vyos_policy_prefix_list6 (Resource)
<center>

Routing policy
⯯
**IPv6 prefix-list filter**


</center>

## Schema

### Required

- `prefix_list6_id` (String) IPv6 prefix-list filter

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
