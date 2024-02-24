---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_ospf_default_information_originate Resource - vyos"
subcategory: "protocols"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>protocols</i>

  <br>
  &darr;
  <br>
  Open Shortest Path First (OSPF)

  <br>
  &darr;
  <br>
  Default route advertisment settings

  <br>
  &darr;
  <br>
  <b>
  Distribute a default route
  </b>
  </div>
---

# vyos_protocols_ospf_default_information_originate (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Open Shortest Path First (OSPF)

<br>
&darr;
<br>
Default route advertisment settings

<br>
&darr;
<br>
<b>
Distribute a default route
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `always` (Boolean) Always advertise a default route
- `metric` (Number) OSPF default metric

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-16777214  &emsp; |  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2  &emsp; |  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).