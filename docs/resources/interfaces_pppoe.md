---
page_title: "vyos_interfaces_pppoe Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Point-to-Point Protocol over Ethernet (PPPoE) Interface
---

# vyos_interfaces_pppoe (Resource)
<center>

*interfaces*  
⯯  
**Point-to-Point Protocol over Ethernet (PPPoE) Interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `access_concentrator` (String) Access concentrator name
- `authentication` (Attributes) Authentication settings (see [below for nested schema](#nestedatt--authentication))
- `connect_on_demand` (Boolean) Establishment connection automatically when traffic is sent
- `default_route_distance` (Number) Distance for installed default route

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|1-255   &emsp;|Distance for the default route from DHCP server  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `dhcpv6_options` (Attributes) DHCPv6 client settings/options (see [below for nested schema](#nestedatt--dhcpv6_options))
- `disable` (Boolean) Administratively disable interface
- `holdoff` (Number) Delay before re-dial to the access concentrator when PPP session terminated by peer (in seconds)

    &emsp;|Format   &emsp;|Description              |
    |-----------|---------------------------|
    &emsp;|0-86400  &emsp;|Holdoff time in seconds  |
- `host_uniq` (String) PPPoE RFC2516 host-uniq tag

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|txt     &emsp;|Host-uniq tag as byte string in HEX  |
- `idle_timeout` (Number) Delay before disconnecting idle session (in seconds)

    &emsp;|Format   &emsp;|Description              |
    |-----------|---------------------------|
    &emsp;|0-86400  &emsp;|Idle timeout in seconds  |
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
- `ipv6` (Attributes) IPv6 routing parameters (see [below for nested schema](#nestedatt--ipv6))
- `local_address` (String) IPv4 address of local end of the PPPoE link

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|ipv4    &emsp;|Address of local end of the PPPoE link  |
- `mirror` (Attributes) Mirror ingress/egress packets (see [below for nested schema](#nestedatt--mirror))
- `mru` (Number) Maximum Receive Unit (MRU) (default: MTU value)

    &emsp;|Format     &emsp;|Description                   |
    |-------------|--------------------------------|
    &emsp;|128-16384  &emsp;|Maximum Receive Unit in byte  |
- `mtu` (Number) Maximum Transmission Unit (MTU)

    &emsp;|Format   &emsp;|Description                        |
    |-----------|-------------------------------------|
    &emsp;|68-1500  &emsp;|Maximum Transmission Unit in byte  |
- `no_default_route` (Boolean) Do not install default route to system
- `no_peer_dns` (Boolean) Do not use DNS servers provided by the peer
- `redirect` (String) Redirect incoming packet to destination

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|txt     &emsp;|Destination interface name  |
- `remote_address` (String) IPv4 address of remote end of the PPPoE link

    &emsp;|Format  &emsp;|Description                              |
    |----------|-------------------------------------------|
    &emsp;|ipv4    &emsp;|Address of remote end of the PPPoE link  |
- `service_name` (String) Service name, only connect to access concentrators advertising this
- `source_interface` (String) Interface used to establish connection

    &emsp;|Format     &emsp;|Description     |
    |-------------|------------------|
    &emsp;|interface  &emsp;|Interface name  |
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

- `pppoe` (String) Point-to-Point Protocol over Ethernet (PPPoE) Interface

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|pppoeN  &emsp;|PPPoE dialer interface name  |


&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `password` (String) Password used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Password     |
- `username` (String) Username used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Username     |


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


&lt;a id=&#34;nestedatt--ip&#34;&gt;&lt;/a&gt;
### Nested Schema for `ip`

Optional:

- `adjust_mss` (String) Adjust TCP MSS value

    &emsp;|Format             &emsp;|Description                                     |
    |---------------------|--------------------------------------------------|
    &emsp;|clamp-mss-to-pmtu  &emsp;|Automatically sets the MSS to the proper value  |
    &emsp;|536-65535          &emsp;|TCP Maximum segment size in bytes               |
- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `source_validation` (String) Source validation by reversed path (RFC3704)

    &emsp;|Format   &emsp;|Description                                                  |
    |-----------|---------------------------------------------------------------|
    &emsp;|strict   &emsp;|Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    &emsp;|loose    &emsp;|Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    &emsp;|disable  &emsp;|No source validation                                         |


&lt;a id=&#34;nestedatt--ipv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6`

Optional:

- `address` (Attributes) IPv6 address configuration modes (see [below for nested schema](#nestedatt--ipv6--address))
- `adjust_mss` (String) Adjust TCP MSS value

    &emsp;|Format             &emsp;|Description                                     |
    |---------------------|--------------------------------------------------|
    &emsp;|clamp-mss-to-pmtu  &emsp;|Automatically sets the MSS to the proper value  |
    &emsp;|536-65535          &emsp;|TCP Maximum segment size in bytes               |
- `disable_forwarding` (Boolean) Disable IP forwarding on this interface

&lt;a id=&#34;nestedatt--ipv6--address&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6.address`

Optional:

- `autoconf` (Boolean) Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)



&lt;a id=&#34;nestedatt--mirror&#34;&gt;&lt;/a&gt;
### Nested Schema for `mirror`

Optional:

- `egress` (String) Mirror egress traffic to destination interface

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|txt     &emsp;|Destination interface name  |
- `ingress` (String) Mirror ingress traffic to destination interface

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|txt     &emsp;|Destination interface name  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
