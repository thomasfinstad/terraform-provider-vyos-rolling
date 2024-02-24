---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_ospf Resource - vyos"
subcategory: "protocols"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>protocols</i>

  <br>
  &darr;
  <br>
  <b>
  Open Shortest Path First (OSPF)
  </b>
  </div>
---

# vyos_protocols_ospf (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
<b>
Open Shortest Path First (OSPF)
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `default_metric` (Number) Metric of redistributed routes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-16777214  &emsp; |  Metric of redistributed routes  |
- `maximum_paths` (Number) Maximum multiple paths (ECMP)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-64  &emsp; |  Maximum multiple paths (ECMP)  |
- `passive_interface` (String) Suppress routing updates on an interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  default  &emsp; |  Default to suppress routing updates on all interfaces  |
- `route_map` (String) Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).