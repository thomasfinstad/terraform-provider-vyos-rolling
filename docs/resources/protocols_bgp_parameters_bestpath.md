---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_bgp_parameters_bestpath Resource - vyos"
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
  Default bestpath selection mechanism
  </b>
  </div>
---

# vyos_protocols_bgp_parameters_bestpath (Resource)

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
Default bestpath selection mechanism
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `bandwidth` (String) Link Bandwidth attribute

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  default-weight-for-missing  &emsp; |  Assign low default weight (1) to paths not having link bandwidth  |
    |  ignore  &emsp; |  Ignore link bandwidth (do regular ECMP, not weighted)  |
    |  skip-missing  &emsp; |  Ignore paths without link bandwidth for ECMP (if other paths have it)  |
- `compare_routerid` (Boolean) Compare the router-id for identical EBGP paths

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).