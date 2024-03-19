---
page_title: "vyos_vrf_name_protocols_ospf_interface_authentication_md5_key_id Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  Routing protocol parameters
  ⯯
  Open Shortest Path First (OSPF)
  ⯯
  Interface configuration
  ⯯
  Authentication
  ⯯
  MD5 key id
  ⯯
  MD5 key id
---

# vyos_vrf_name_protocols_ospf_interface_authentication_md5_key_id (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Open Shortest Path First (OSPF)
⯯
Interface configuration
⯯
Authentication
⯯
MD5 key id
⯯
**MD5 key id**


</center>

## Schema

### Required

- `interface_id` (String) Interface configuration

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `key_id_id` (Number) MD5 key id

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1-255   &emsp;|MD5 key id   |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Optional

- `md5_key` (String) MD5 authentication type

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|MD5 Key (16 characters or less)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
