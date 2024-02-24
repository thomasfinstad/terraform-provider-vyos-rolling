---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_vpn_sstp_client_ip_pool Resource - vyos"
subcategory: "vpn"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>vpn</i>

  <br>
  &darr;
  <br>
  Secure Socket Tunneling Protocol (SSTP) server

  <br>
  &darr;
  <br>
  <b>
  Client IP pools and gateway setting
  </b>
  </div>
---

# vyos_vpn_sstp_client_ip_pool (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>vpn</i>

<br>
&darr;
<br>
Secure Socket Tunneling Protocol (SSTP) server

<br>
&darr;
<br>
<b>
Client IP pools and gateway setting
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `subnet` (List of String) Client IP subnet (CIDR notation)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 address and prefix length  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).