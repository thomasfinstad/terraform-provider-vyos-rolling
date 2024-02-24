---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_dhcp_relay Resource - vyos"
subcategory: "service"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>service</i>

  <br>
  &darr;
  <br>
  <b>
  Host Configuration Protocol (DHCP) relay agent
  </b>
  </div>
---

# vyos_service_dhcp_relay (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
<b>
Host Configuration Protocol (DHCP) relay agent
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `interface` (List of String) Interface Name to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |
- `listen_interface` (List of String) Interface for DHCP Relay Agent to listen for requests

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |
- `server` (List of String) DHCP server address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  DHCP server IPv4 address  |
- `upstream_interface` (List of String) Interface for DHCP Relay Agent forward requests out

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).