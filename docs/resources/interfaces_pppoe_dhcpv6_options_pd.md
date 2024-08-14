---
page_title: "vyos_interfaces_pppoe_dhcpv6_options_pd Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Point-to-Point Protocol over Ethernet (PPPoE) Interface⯯DHCPv6 client settings/options⯯DHCPv6 prefix delegation interface statement
---

# vyos_interfaces_pppoe_dhcpv6_options_pd (Resource)
<center>

*interfaces*  
⯯  
Point-to-Point Protocol over Ethernet (PPPoE) Interface  
⯯  
DHCPv6 client settings/options  
⯯  
**DHCPv6 prefix delegation interface statement**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `length` (Number) Request IPv6 prefix length from peer

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |32-64   &emsp;|Length of delegated prefix  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `pd` (String) DHCPv6 prefix delegation interface statement

    |Format           &emsp;|Description                        |
    |-------------------|-------------------------------------|
    |instance number  &emsp;|Prefix delegation instance (&gt;= 0)  |
- `pppoe` (String) Point-to-Point Protocol over Ethernet (PPPoE) Interface

    |Format  &emsp;|Description                  |
    |----------|-------------------------------|
    |pppoeN  &emsp;|PPPoE dialer interface name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
