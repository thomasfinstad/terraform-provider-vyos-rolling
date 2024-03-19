---
page_title: "vyos_vrf Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  **Virtual Routing and Forwarding**
---

# vyos_vrf (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

	**Virtual Routing and Forwarding**


</center>

## Schema

### Optional

- `bind_to_all` (Boolean) Enable binding services to all VRFs

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
