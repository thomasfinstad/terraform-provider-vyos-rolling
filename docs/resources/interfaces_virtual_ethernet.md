---
page_title: "vyos_interfaces_virtual_ethernet Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Virtual Ethernet (veth) Interface
---

# vyos_interfaces_virtual_ethernet (Resource)
<center>

*interfaces*  
⯯  
**Virtual Ethernet (veth) Interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (List of String) IP address

    &emsp;|Format   &emsp;|Description                                   |
    |-----------|------------------------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 address and prefix length                |
    &emsp;|ipv6net  &emsp;|IPv6 address and prefix length                |
    &emsp;|dhcp     &emsp;|Dynamic Host Configuration Protocol           |
    &emsp;|dhcpv6   &emsp;|Dynamic Host Configuration Protocol for IPv6  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `dhcp_options` (Attributes) DHCP client settings/options (see [below for nested schema](#nestedatt--dhcp_options))
- `dhcpv6_options` (Attributes) DHCPv6 client settings/options (see [below for nested schema](#nestedatt--dhcpv6_options))
- `disable` (Boolean) Administratively disable interface
- `netns` (String) Network namespace name

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Network namespace name  |
- `peer_name` (String) Virtual ethernet peer interface name

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Name of peer interface  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF instance name

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `virtual_ethernet` (String) Virtual Ethernet (veth) Interface

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|vethN   &emsp;|Virtual Ethernet interface name  |


&lt;a id=&#34;nestedatt--dhcp_options&#34;&gt;&lt;/a&gt;
### Nested Schema for `dhcp_options`

Optional:

- `client_id` (String) Identifier used by client to identify itself to the DHCP server

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|txt     &emsp;|DHCP option string  |
- `default_route_distance` (Number) Distance for installed default route

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|1-255   &emsp;|Distance for the default route from DHCP server  |
- `host_name` (String) Override system host-name sent to DHCP server
- `mtu` (Boolean) Use MTU value from DHCP server - ignore interface setting
- `no_default_route` (Boolean) Do not install default route to system
- `reject` (List of String) IP addresses or subnets from which to reject DHCP leases

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|ipv4     &emsp;|IPv4 address to match  |
    &emsp;|ipv4net  &emsp;|IPv4 prefix to match   |
- `user_class` (String) Identify to the DHCP server, user configurable option

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|txt     &emsp;|DHCP option string  |
- `vendor_class_id` (String) Identify the vendor client type to the DHCP server

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|txt     &emsp;|DHCP option string  |


&lt;a id=&#34;nestedatt--dhcpv6_options&#34;&gt;&lt;/a&gt;
### Nested Schema for `dhcpv6_options`

Optional:

- `duid` (String) DHCP unique identifier (DUID) to be sent by client

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|duid    &emsp;|DHCP unique identifier  |
- `no_release` (Boolean) Do not send a release message on client exit
- `parameters_only` (Boolean) Acquire only config parameters, no address
- `rapid_commit` (Boolean) Wait for immediate reply instead of advertisements
- `temporary` (Boolean) IPv6 temporary address


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
