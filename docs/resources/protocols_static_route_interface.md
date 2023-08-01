---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_static_route_interface Resource - vyos"
subcategory: ""
description: |-
  <div style="text-align: center">
  Routing protocols

  <br>
  &darr;
  <br>
  Static Routing

  <br>
  &darr;
  <br>
  Static IPv4 route

  <br>
  &darr;
  <br>
  <b>
  Next-hop IPv4 router interface
  </b>
  </div>
---

# vyos_protocols_static_route_interface (Resource)

<div style="text-align: center">
Routing protocols

<br>
&darr;
<br>
Static Routing

<br>
&darr;
<br>
Static IPv4 route

<br>
&darr;
<br>
<b>
Next-hop IPv4 router interface
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_id` (String) Next-hop IPv4 router interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Gateway interface name  |
- `route_id` (String) Static IPv4 route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 static route  |

### Optional

- `disable` (Boolean) Disable instance
- `distance` (Number) Distance for this route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Distance for this route  |
- `vrf` (String) VRF to leak route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of VRF to leak to  |