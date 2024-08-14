---
page_title: "vyos_interfaces_vxlan Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Virtual Extensible LAN (VXLAN) Interface
---

# vyos_interfaces_vxlan (Resource)
<center>

*interfaces*  
⯯  
**Virtual Extensible LAN (VXLAN) Interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (List of String) IP address

    |Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    |ipv4net  &emsp;|IPv4 address and prefix length  |
    |ipv6net  &emsp;|IPv6 address and prefix length  |
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `disable` (Boolean) Administratively disable interface
- `gpe` (Boolean) Enable Generic Protocol extension (VXLAN-GPE)
- `group` (String) Multicast group address for VXLAN interface

    |Format  &emsp;|Description                   |
    |----------|--------------------------------|
    |ipv4    &emsp;|Multicast IPv4 group address  |
    |ipv6    &emsp;|Multicast IPv6 group address  |
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
- `ipv6` (Attributes) IPv6 routing parameters (see [below for nested schema](#nestedatt--ipv6))
- `mac` (String) Media Access Control (MAC) address

    |Format   &emsp;|Description             |
    |-----------|--------------------------|
    |macaddr  &emsp;|Hardware (MAC) address  |
- `mirror` (Attributes) Mirror ingress/egress packets (see [below for nested schema](#nestedatt--mirror))
- `mtu` (Number) Maximum Transmission Unit (MTU)

    |Format      &emsp;|Description                        |
    |--------------|-------------------------------------|
    |1200-16000  &emsp;|Maximum Transmission Unit in byte  |
- `parameters` (Attributes) VXLAN tunnel parameters (see [below for nested schema](#nestedatt--parameters))
- `port` (Number) Port number used by connection

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |1-65535  &emsp;|Numeric IP port  |
- `redirect` (String) Redirect incoming packet to destination

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |txt     &emsp;|Destination interface name  |
- `remote` (List of String) Tunnel remote address

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |ipv4    &emsp;|Tunnel remote IPv4 address  |
    |ipv6    &emsp;|Tunnel remote IPv6 address  |
- `source_address` (String) Source IP address used to initiate connection

    |Format  &emsp;|Description          |
    |----------|-----------------------|
    |ipv4    &emsp;|IPv4 source address  |
    |ipv6    &emsp;|IPv6 source address  |
- `source_interface` (String) Interface used to establish connection

    |Format     &emsp;|Description     |
    |-------------|------------------|
    |interface  &emsp;|Interface name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vni` (Number) Virtual Network Identifier

    |Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    |0-16777214  &emsp;|VXLAN virtual network identifier  |
- `vrf` (String) VRF instance name

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `vxlan` (String) Virtual Extensible LAN (VXLAN) Interface

    |Format  &emsp;|Description           |
    |----------|------------------------|
    |vxlanN  &emsp;|VXLAN interface name  |


<a id="nestedatt--ip"></a>
### Nested Schema for `ip`

Optional:

- `adjust_mss` (String) Adjust TCP MSS value

    |Format             &emsp;|Description                                     |
    |---------------------|--------------------------------------------------|
    |clamp-mss-to-pmtu  &emsp;|Automatically sets the MSS to the proper value  |
    |536-65535          &emsp;|TCP Maximum segment size in bytes               |
- `arp_cache_timeout` (Number) ARP cache entry timeout in seconds

    |Format   &emsp;|Description                        |
    |-----------|-------------------------------------|
    |1-86400  &emsp;|ARP cache entry timout in seconds  |
- `disable_arp_filter` (Boolean) Disable ARP filter on this interface
- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `enable_arp_accept` (Boolean) Enable ARP accept on this interface
- `enable_arp_announce` (Boolean) Enable ARP announce on this interface
- `enable_arp_ignore` (Boolean) Enable ARP ignore on this interface
- `enable_directed_broadcast` (Boolean) Enable directed broadcast forwarding on this interface
- `enable_proxy_arp` (Boolean) Enable proxy-arp on this interface
- `proxy_arp_pvlan` (Boolean) Enable private VLAN proxy ARP on this interface
- `source_validation` (String) Source validation by reversed path (RFC3704)

    |Format   &emsp;|Description                                                  |
    |-----------|---------------------------------------------------------------|
    |strict   &emsp;|Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    |loose    &emsp;|Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    |disable  &emsp;|No source validation                                         |


<a id="nestedatt--ipv6"></a>
### Nested Schema for `ipv6`

Optional:

- `accept_dad` (String) Accept Duplicate Address Detection

    |Format  &emsp;|Description                                                                |
    |----------|-----------------------------------------------------------------------------|
    |0       &emsp;|Disable DAD                                                                |
    |1       &emsp;|Enable DAD                                                                 |
    |2       &emsp;|Enable DAD - disable IPv6 if MAC-based duplicate link-local address found  |
- `address` (Attributes) IPv6 address configuration modes (see [below for nested schema](#nestedatt--ipv6--address))
- `adjust_mss` (String) Adjust TCP MSS value

    |Format             &emsp;|Description                                     |
    |---------------------|--------------------------------------------------|
    |clamp-mss-to-pmtu  &emsp;|Automatically sets the MSS to the proper value  |
    |536-65535          &emsp;|TCP Maximum segment size in bytes               |
- `base_reachable_time` (Number) Base reachable time in seconds

    |Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    |1-86400  &emsp;|Base reachable time in seconds  |
- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `dup_addr_detect_transmits` (Number) Number of NS messages to send while performing DAD

    |Format  &emsp;|Description                                         |
    |----------|------------------------------------------------------|
    |0       &emsp;|Disable Duplicate Address Dectection (DAD)          |
    |1-n     &emsp;|Number of NS messages to send while performing DAD  |
- `source_validation` (String) Source validation by reversed path (RFC3704)

    |Format   &emsp;|Description                                                  |
    |-----------|---------------------------------------------------------------|
    |strict   &emsp;|Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    |loose    &emsp;|Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    |disable  &emsp;|No source validation                                         |

<a id="nestedatt--ipv6--address"></a>
### Nested Schema for `ipv6.address`

Optional:

- `autoconf` (Boolean) Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)
- `eui64` (List of String) Prefix for IPv6 address with MAC-based EUI-64

    |Format                &emsp;|Description       |
    |------------------------|--------------------|
    |&lt;h:h:h:h:h:h:h:h/64&gt;  &emsp;|IPv6 /64 network  |
- `no_default_link_local` (Boolean) Remove the default link-local address from the interface



<a id="nestedatt--mirror"></a>
### Nested Schema for `mirror`

Optional:

- `egress` (String) Mirror egress traffic to destination interface

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |txt     &emsp;|Destination interface name  |
- `ingress` (String) Mirror ingress traffic to destination interface

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |txt     &emsp;|Destination interface name  |


<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `external` (Boolean) Use external control plane
- `ip` (Attributes) IPv4 specific tunnel parameters (see [below for nested schema](#nestedatt--parameters--ip))
- `ipv6` (Attributes) IPv6 specific tunnel parameters (see [below for nested schema](#nestedatt--parameters--ipv6))
- `neighbor_suppress` (Boolean) Enable neighbor discovery (ARP and ND) suppression
- `nolearning` (Boolean) Do not add unknown addresses into forwarding database
- `vni_filter` (Boolean) Enable VNI filter support

<a id="nestedatt--parameters--ip"></a>
### Nested Schema for `parameters.ip`

Optional:

- `df` (String) Usage of the DF (don&#39;t Fragment) bit in outgoing packets

    |Format   &emsp;|Description                           |
    |-----------|----------------------------------------|
    |set      &emsp;|Always set DF (don&#39;t fragment) bit    |
    |unset    &emsp;|Always unset DF (don&#39;t fragment) bit  |
    |inherit  &emsp;|Copy from the original IP header      |
- `tos` (Number) Specifies TOS value to use in outgoing packets

    |Format  &emsp;|Description            |
    |----------|-------------------------|
    |0-99    &emsp;|Type of Service (TOS)  |
- `ttl` (Number) Specifies TTL value to use in outgoing packets

    |Format  &emsp;|Description                                   |
    |----------|------------------------------------------------|
    |0       &emsp;|Inherit - copy value from original IP header  |
    |1-255   &emsp;|Time to Live                                  |


<a id="nestedatt--parameters--ipv6"></a>
### Nested Schema for `parameters.ipv6`

Optional:

- `flowlabel` (String) Specifies the flow label to use in outgoing packets

    |Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    |inherit       &emsp;|Copy field from original header  |
    |0x0-0x0fffff  &emsp;|Tunnel key, or hex value         |



<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
