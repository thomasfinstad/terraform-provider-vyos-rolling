---
page_title: "vyos_container_name_network Resource - terraform-provider-vyos"
subcategory: "container"
description: |-
  Container applications
  ⯯
  Container name
  ⯯
  Attach user defined network to container
---

# vyos_container_name_network (Resource)
<center>

Container applications
⯯
Container name
⯯
**Attach user defined network to container**


</center>

## Schema

### Required

- `name_id` (String) Container name
- `network_id` (String) Attach user defined network to container

### Optional

- `address` (List of String) Assign static IP address to container

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|ipv4    &emsp;|IPv4 address  |
    &emsp;|ipv6    &emsp;|IPv6 address  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
