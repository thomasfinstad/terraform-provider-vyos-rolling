---
page_title: "vyos_interfaces_pseudo_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Pseudo Ethernet Interface (Macvlan)⯯QinQ TAG-S Virtual Local Area Network (VLAN) ID⯯QinQ TAG-C Virtual Local Area Network (VLAN) ID⯯DHCPv6 client settings/options⯯DHCPv6 prefix delegation interface statement⯯Delegate IPv6 prefix from provider to this interface
---

# vyos_interfaces_pseudo_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface (Resource)
<center>

*interfaces*  
⯯  
Pseudo Ethernet Interface (Macvlan)  
⯯  
QinQ TAG-S Virtual Local Area Network (VLAN) ID  
⯯  
QinQ TAG-C Virtual Local Area Network (VLAN) ID  
⯯  
DHCPv6 client settings/options  
⯯  
DHCPv6 prefix delegation interface statement  
⯯  
**Delegate IPv6 prefix from provider to this interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (String) Local interface address assigned to interface (default: EUI-64)

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|&gt;0      &emsp;|Used to form IPv6 interface address  |
- `sla_id` (Number) Interface site-Level aggregator (SLA)

    &emsp;|Format   &emsp;|Description                                          |
    |-----------|-------------------------------------------------------|
    &emsp;|0-65535  &emsp;|Decimal integer which fits in the length of SLA IDs  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) Delegate IPv6 prefix from provider to this interface
- `pd` (String) DHCPv6 prefix delegation interface statement

    &emsp;|Format           &emsp;|Description                        |
    |-------------------|-------------------------------------|
    &emsp;|instance number  &emsp;|Prefix delegation instance (&gt;= 0)  |
- `pseudo_ethernet` (String) Pseudo Ethernet Interface (Macvlan)

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|pethN   &emsp;|Pseudo Ethernet interface name  |
- `vif_c` (String) QinQ TAG-C Virtual Local Area Network (VLAN) ID
- `vif_s` (Number) QinQ TAG-S Virtual Local Area Network (VLAN) ID

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|0-4094  &emsp;|QinQ Virtual Local Area Network (VLAN) ID  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
