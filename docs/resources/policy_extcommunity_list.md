---
page_title: "vyos_policy_extcommunity_list Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy
  ⯯
  Add a BGP extended community list entry
---

# vyos_policy_extcommunity_list (Resource)
<center>

Routing policy
⯯
**Add a BGP extended community list entry**


</center>

## Schema

### Required

- `extcommunity_list_id` (String) Add a BGP extended community list entry

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|txt     &emsp;|BGP extended community-list name  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
