---
page_title: "vyos_interfaces_bridge_dhcpv6_options_pd_interface Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Bridge Interface⯯DHCPv6 client settings/options⯯DHCPv6 prefix delegation interface statement⯯Delegate IPv6 prefix from provider to this interface
---

# vyos_interfaces_bridge_dhcpv6_options_pd_interface (Resource)
<center>

*interfaces*  
⯯  
Bridge Interface  
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

    |Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    |&gt;0      &emsp;|Used to form IPv6 interface address  |
- `sla_id` (Number) Interface site-Level aggregator (SLA)

    |Format   &emsp;|Description                                          |
    |-----------|-------------------------------------------------------|
    |0-65535  &emsp;|Decimal integer which fits in the length of SLA IDs  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `bridge` (String) Bridge Interface

    |Format  &emsp;|Description            |
    |----------|-------------------------|
    |brN     &emsp;|Bridge interface name  |
- `interface` (String) Delegate IPv6 prefix from provider to this interface
- `pd` (String) DHCPv6 prefix delegation interface statement

    |Format           &emsp;|Description                        |
    |-------------------|-------------------------------------|
    |instance number  &emsp;|Prefix delegation instance (&gt;= 0)  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
