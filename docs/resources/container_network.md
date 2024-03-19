---
page_title: "vyos_container_network Resource - terraform-provider-vyos"
subcategory: "container"
description: |-
  Container applications
  ⯯
  Network name
---

# vyos_container_network (Resource)
<center>

Container applications
⯯
**Network name**


</center>

## Schema

### Required

- `network_id` (String) Network name

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `prefix` (List of String) Prefix which allocated to that network

    &emsp;|Format   &emsp;|Description          |
    |-----------|-----------------------|
    &emsp;|ipv4net  &emsp;|IPv4 network prefix  |
    &emsp;|ipv6net  &emsp;|IPv6 network prefix  |
- `vrf` (String) VRF instance name

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
