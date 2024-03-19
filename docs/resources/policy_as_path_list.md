---
page_title: "vyos_policy_as_path_list Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy
  ⯯
  Add a BGP autonomous system path filter
---

# vyos_policy_as_path_list (Resource)
<center>

Routing policy
⯯
**Add a BGP autonomous system path filter**


</center>

## Schema

### Required

- `as_path_list_id` (String) Add a BGP autonomous system path filter

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|AS path list name  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
