---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_vpn_openconnect_network_settings Resource - vyos"
subcategory: "vpn"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>vpn</i>

  <br>
  &darr;
  <br>
  SSL VPN OpenConnect, AnyConnect compatible server

  <br>
  &darr;
  <br>
  <b>
  Network settings
  </b>
  </div>
---

# vyos_vpn_openconnect_network_settings (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>vpn</i>

<br>
&darr;
<br>
SSL VPN OpenConnect, AnyConnect compatible server

<br>
&darr;
<br>
<b>
Network settings
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `name_server` (List of String) Domain Name Servers (DNS) addresses

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  Domain Name Server (DNS) IPv4 address  |
    |  ipv6  &emsp; |  Domain Name Server (DNS) IPv6 address  |
- `push_route` (List of String) Route to be pushed to the client

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IPv4 network and prefix length  |
    |  ipv6net  &emsp; |  IPv6 network and prefix length  |
- `split_dns` (List of String) Domains over which the provided DNS should be used

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Client prefix length  |
- `tunnel_all_dns` (String) If the tunnel-all-dns option is set to yes, tunnel all DNS queries via the VPN. This is the default when a default route is set.

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  yes  &emsp; |  Enable tunneling of all DNS traffic  |
    |  no  &emsp; |  Disable tunneling of all DNS traffic  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).