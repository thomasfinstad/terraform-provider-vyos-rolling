---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_interfaces_ethernet_vif_s_vif_c_dhcpv6_options_pd Resource - vyos"
subcategory: ""
description: |-
  <div style="text-align: center">
  Network interfaces

  <br>
  &darr;
  <br>
  Ethernet Interface

  <br>
  &darr;
  <br>
  QinQ TAG-S Virtual Local Area Network (VLAN) ID

  <br>
  &darr;
  <br>
  QinQ TAG-C Virtual Local Area Network (VLAN) ID

  <br>
  &darr;
  <br>
  DHCPv6 client settings/options

  <br>
  &darr;
  <br>
  <b>
  DHCPv6 prefix delegation interface statement
  </b>
  </div>
---

# vyos_interfaces_ethernet_vif_s_vif_c_dhcpv6_options_pd (Resource)

<div style="text-align: center">
Network interfaces

<br>
&darr;
<br>
Ethernet Interface

<br>
&darr;
<br>
QinQ TAG-S Virtual Local Area Network (VLAN) ID

<br>
&darr;
<br>
QinQ TAG-C Virtual Local Area Network (VLAN) ID

<br>
&darr;
<br>
DHCPv6 client settings/options

<br>
&darr;
<br>
<b>
DHCPv6 prefix delegation interface statement
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ethernet_id` (String) Ethernet Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ethN  &emsp; |  Ethernet interface name  |
- `pd_id` (String) DHCPv6 prefix delegation interface statement

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  instance number  &emsp; |  Prefix delegation instance (>= 0)  |
- `vif_c_id` (String) QinQ TAG-C Virtual Local Area Network (VLAN) ID
- `vif_s_id` (String) QinQ TAG-S Virtual Local Area Network (VLAN) ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4094  &emsp; |  QinQ Virtual Local Area Network (VLAN) ID  |

### Optional

- `length` (Number) Request IPv6 prefix length from peer

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 32-64  &emsp; |  Length of delegated prefix  |