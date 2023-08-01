---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_nhrp_tunnel_dynamic_map Resource - vyos"
subcategory: ""
description: |-
  <div style="text-align: center">
  <i>protocols</i>

  <br>
  &darr;
  <br>
  Next Hop Resolution Protocol (NHRP) parameters

  <br>
  &darr;
  <br>
  Tunnel for NHRP

  <br>
  &darr;
  <br>
  <b>
  Set an HUB tunnel address
  </b>
  </div>
---

# vyos_protocols_nhrp_tunnel_dynamic_map (Resource)

<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Next Hop Resolution Protocol (NHRP) parameters

<br>
&darr;
<br>
Tunnel for NHRP

<br>
&darr;
<br>
<b>
Set an HUB tunnel address
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dynamic_map_id` (String) Set an HUB tunnel address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  Set the IP address and prefix length  |
- `tunnel_id` (String) Tunnel for NHRP

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  tunN  &emsp; |  NHRP tunnel name  |

### Optional

- `nbma_domain_name` (String) Set HUB fqdn (nbma-address - fqdn)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  <fqdn>  &emsp; |  Set the external HUB fqdn  |