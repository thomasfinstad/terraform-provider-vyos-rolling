---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_vrf Resource - vyos"
subcategory: "vrf.md"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <b>
  Virtual Routing and Forwarding
  </b>
  </div>
---

# vyos_vrf (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<b>
Virtual Routing and Forwarding
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `bind_to_all` (Boolean) Enable binding services to all VRFs

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).