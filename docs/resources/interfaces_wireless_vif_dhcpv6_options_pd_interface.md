---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_interfaces_wireless_vif_dhcpv6_options_pd_interface Resource - vyos"
subcategory: ""
description: |-
  <div style="text-align: center">
  <i>interfaces</i>

  <br>
  &darr;
  <br>
  Wireless (WiFi/WLAN) Network Interface

  <br>
  &darr;
  <br>
  Virtual Local Area Network (VLAN) ID

  <br>
  &darr;
  <br>
  DHCPv6 client settings/options

  <br>
  &darr;
  <br>
  DHCPv6 prefix delegation interface statement

  <br>
  &darr;
  <br>
  <b>
  Delegate IPv6 prefix from provider to this interface
  </b>
  </div>
---

# vyos_interfaces_wireless_vif_dhcpv6_options_pd_interface (Resource)

<div style="text-align: center">
<i>interfaces</i>

<br>
&darr;
<br>
Wireless (WiFi/WLAN) Network Interface

<br>
&darr;
<br>
Virtual Local Area Network (VLAN) ID

<br>
&darr;
<br>
DHCPv6 client settings/options

<br>
&darr;
<br>
DHCPv6 prefix delegation interface statement

<br>
&darr;
<br>
<b>
Delegate IPv6 prefix from provider to this interface
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_id` (String) Delegate IPv6 prefix from provider to this interface
- `pd_id` (String) DHCPv6 prefix delegation interface statement

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  instance number  &emsp; |  Prefix delegation instance (>= 0)  |
- `vif_id` (String) Virtual Local Area Network (VLAN) ID

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-4094  &emsp; |  Virtual Local Area Network (VLAN) ID  |
- `wireless_id` (String) Wireless (WiFi/WLAN) Network Interface

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  wlanN  &emsp; |  Wireless (WiFi/WLAN) interface name  |

### Optional

- `address` (String) Local interface address assigned to interface (default: EUI-64)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  >0  &emsp; |  Used to form IPv6 interface address  |
- `sla_id` (Number) Interface site-Level aggregator (SLA)

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-65535  &emsp; |  Decimal integer which fits in the length of SLA IDs  |