---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_bgp_address_family_ipv6_unicast_route_map_vpn Resource - vyos"
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
  BGP address-family parameters

  <br>
  &darr;
  <br>
  IPv6 BGP settings

  <br>
  &darr;
  <br>
  Route-map to filter route updates to/from this peer

  <br>
  &darr;
  <br>
  <b>
  Between current address-family and VPN
  </b>
  </div>
---

# vyos_protocols_bgp_address_family_ipv6_unicast_route_map_vpn (Resource)

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
BGP address-family parameters

<br>
&darr;
<br>
IPv6 BGP settings

<br>
&darr;
<br>
Route-map to filter route updates to/from this peer

<br>
&darr;
<br>
<b>
Between current address-family and VPN
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `export` (String) Route-map to filter outgoing route updates

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |
- `import` (String) Route-map to filter incoming route updates

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Route map name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).