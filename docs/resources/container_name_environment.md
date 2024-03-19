---
page_title: "vyos_container_name_environment Resource - terraform-provider-vyos"
subcategory: "container"
description: |-
  Container applications
  ⯯
  Container name
  ⯯
  Add custom environment variables
---

# vyos_container_name_environment (Resource)
<center>

Container applications
⯯
Container name
⯯
**Add custom environment variables**


</center>

## Schema

### Required

- `environment_id` (String) Add custom environment variables
- `name_id` (String) Container name

### Optional

- `value` (String) Set environment option value

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|txt     &emsp;|Set environment option value  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
