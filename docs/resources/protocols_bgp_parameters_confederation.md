---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_bgp_parameters_confederation Resource - vyos"
subcategory: "protocols"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>protocols</i>

  <br>
  &darr;
  <br>
  Border Gateway Protocol (BGP)

  <br>
  &darr;
  <br>
  BGP parameters

  <br>
  &darr;
  <br>
  <b>
  AS confederation parameters
  </b>
  </div>
---

# vyos_protocols_bgp_parameters_confederation (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Border Gateway Protocol (BGP)

<br>
&darr;
<br>
BGP parameters

<br>
&darr;
<br>
<b>
AS confederation parameters
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `identifier` (Number) Confederation AS identifier

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967294  &emsp; |  Confederation AS id  |
- `peers` (List of Number) Peer ASs in the BGP confederation

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-4294967294  &emsp; |  Peer AS number  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).