---
page_title: "vyos_policy_access_list Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy
  ⯯
  IP access-list filter
---

# vyos_policy_access_list (Resource)
<center>

Routing policy
⯯
**IP access-list filter**


</center>

## Schema

### Required

- `access_list_id` (Number) IP access-list filter

    &emsp;|Format     &emsp;|Description                               |
    |-------------|--------------------------------------------|
    &emsp;|1-99       &emsp;|IP standard access list                   |
    &emsp;|100-199    &emsp;|IP extended access list                   |
    &emsp;|1300-1999  &emsp;|IP standard access list (expanded range)  |
    &emsp;|2000-2699  &emsp;|IP extended access list (expanded range)  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
