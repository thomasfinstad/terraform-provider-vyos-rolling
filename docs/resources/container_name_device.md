---
page_title: "vyos_container_name_device Resource - terraform-provider-vyos"
subcategory: "container"
description: |-
  Container applications
  ⯯
  Container name
  ⯯
  Add a host device to the container
---

# vyos_container_name_device (Resource)
<center>

Container applications
⯯
Container name
⯯
**Add a host device to the container**


</center>

## Schema

### Required

- `device_id` (String) Add a host device to the container
- `name_id` (String) Container name

### Optional

- `destination` (String) Destination container device (Example: &#34;/dev/x&#34;)

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|txt     &emsp;|Destination container device  |
- `source` (String) Source device (Example: &#34;/dev/x&#34;)

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|txt     &emsp;|Source device  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
