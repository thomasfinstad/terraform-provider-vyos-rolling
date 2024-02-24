---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_load_balancing_wan Resource - vyos"
subcategory: "load"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  Configure load-balancing

  <br>
  &darr;
  <br>
  <b>
  Configure Wide Area Network (WAN) load-balancing
  </b>
  </div>
---

# vyos_load_balancing_wan (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
Configure load-balancing

<br>
&darr;
<br>
<b>
Configure Wide Area Network (WAN) load-balancing
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `disable_source_nat` (Boolean) Disable source NAT rules from being configured for WAN load balancing
- `enable_local_traffic` (Boolean) Enable WAN load balancing for locally sourced traffic
- `flush_connections` (Boolean) Flush connection tracking tables on connection state change
- `hook` (String) Script to be executed on interface status change

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Script in /config/scripts  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).