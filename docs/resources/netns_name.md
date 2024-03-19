---
page_title: "vyos_netns_name Resource - terraform-provider-vyos"
subcategory: "netns"
description: |-
  Network namespace
  ⯯
  Network namespace name
---

# vyos_netns_name (Resource)
<center>

Network namespace
⯯
**Network namespace name**


</center>

## Schema

### Required

- `name_id` (String) Network namespace name

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
