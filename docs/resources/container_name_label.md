---
page_title: "vyos_container_name_label Resource - terraform-provider-vyos"
subcategory: "container"
description: |-
  Container applications
  ⯯
  Container name
  ⯯
  Add label variables
---

# vyos_container_name_label (Resource)
<center>

Container applications
⯯
Container name
⯯
**Add label variables**


</center>

## Schema

### Required

- `label_id` (String) Add label variables
- `name_id` (String) Container name

### Optional

- `value` (String) Set label option value

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Set label option value  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
