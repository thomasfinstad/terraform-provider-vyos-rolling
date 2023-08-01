---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_vrf_name_protocols_static_route6_next_hop Resource - vyos"
subcategory: ""
description: |-
  <div style="text-align: center">
  Virtual Routing and Forwarding

  <br>
  &darr;
  <br>
  Virtual Routing and Forwarding instance

  <br>
  &darr;
  <br>
  Routing protocol parameters

  <br>
  &darr;
  <br>
  Static Routing

  <br>
  &darr;
  <br>
  Static IPv6 route

  <br>
  &darr;
  <br>
  <b>
  IPv6 gateway address
  </b>
  </div>
---

# vyos_vrf_name_protocols_static_route6_next_hop (Resource)

<div style="text-align: center">
Virtual Routing and Forwarding

<br>
&darr;
<br>
Virtual Routing and Forwarding instance

<br>
&darr;
<br>
Routing protocol parameters

<br>
&darr;
<br>
Static Routing

<br>
&darr;
<br>
Static IPv6 route

<br>
&darr;
<br>
<b>
IPv6 gateway address
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  VRF instance name  |
- `next_hop_id` (String) IPv6 gateway address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6  &emsp; |  Next-hop IPv6 router  |
- `route6_id` (String) Static IPv6 route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6net  &emsp; |  IPv6 static route  |

### Optional

- `disable` (Boolean) Disable instance
- `distance` (Number) Distance for this route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-255  &emsp; |  Distance for this route  |
- `interface` (String) Gateway interface name

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Gateway interface name  |
- `vrf` (String) VRF to leak route

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of VRF to leak to  |