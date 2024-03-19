---
page_title: "vyos_container_registry Resource - terraform-provider-vyos"
subcategory: "container"
description: |-
  Container applications
  ⯯
  Registry Name
---

# vyos_container_registry (Resource)
<center>

Container applications
⯯
**Registry Name**


</center>

## Schema

### Required

- `registry_id` (String) Registry Name

### Optional

- `authentication` (Attributes) Authentication settings (see [below for nested schema](#nestedatt--authentication))
- `disable` (Boolean) Disable instance

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `password` (String) Password used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Password     |
- `username` (String) Username used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Username     |  &emsp;|
