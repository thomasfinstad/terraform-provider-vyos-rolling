---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_pppoe_server_limits Resource - vyos"
subcategory: "service"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>service</i>

  <br>
  &darr;
  <br>
  Point to Point over Ethernet (PPPoE) Server

  <br>
  &darr;
  <br>
  <b>
  Limits the connection rate from a single source
  </b>
  </div>
---

# vyos_service_pppoe_server_limits (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
Point to Point over Ethernet (PPPoE) Server

<br>
&darr;
<br>
<b>
Limits the connection rate from a single source
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `burst` (String) Burst count
- `connection_limit` (String) Acceptable rate of connections (e.g. 1/min, 60/sec)
- `timeout` (String) Timeout in seconds

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).