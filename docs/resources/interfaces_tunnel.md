---
page_title: "vyos_interfaces_tunnel Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Tunnel interface
---

# vyos_interfaces_tunnel (Resource)
<center>

*interfaces*  
⯯  
**Tunnel interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `_6rd_prefix` (String) 6rd network prefix

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv6    &emsp;|IPv6 address and prefix length  |
- `_6rd_relay_prefix` (String) 6rd relay prefix

    &emsp;|Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 prefix of interface for 6rd  |
- `address` (List of String) IP address

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 address and prefix length  |
    &emsp;|ipv6net  &emsp;|IPv6 address and prefix length  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `disable` (Boolean) Administratively disable interface
- `disable_link_detect` (Boolean) Ignore link state changes
- `enable_multicast` (Boolean) Enable multicast operation over tunnel
- `encapsulation` (String) Encapsulation of this tunnel interface

    &emsp;|Format     &emsp;|Description                                           |
    |-------------|--------------------------------------------------------|
    &emsp;|erspan     &emsp;|Encapsulated Remote Switched Port Analyzer            |
    &emsp;|gre        &emsp;|Generic Routing Encapsulation (network layer)         |
    &emsp;|gretap     &emsp;|Generic Routing Encapsulation (datalink layer)        |
    &emsp;|ip6erspan  &emsp;|Encapsulated Remote Switched Port Analyzer over IPv6  |
    &emsp;|ip6gre     &emsp;|GRE over IPv6 (network layer)                         |
    &emsp;|ip6gretap  &emsp;|GRE over IPv6 (datalink layer)                        |
    &emsp;|ip6ip6     &emsp;|IPv6 in IPv6 encapsulation                            |
    &emsp;|ipip       &emsp;|IPv4 in IPv4 encapsulation                            |
    &emsp;|ipip6      &emsp;|IPv4 in IP6 encapsulation                             |
    &emsp;|sit        &emsp;|Simple Internet Transition (IPv6 in IPv4)             |
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
- `ipv6` (Attributes) IPv6 routing parameters (see [below for nested schema](#nestedatt--ipv6))
- `mirror` (Attributes) Mirror ingress/egress packets (see [below for nested schema](#nestedatt--mirror))
- `mtu` (Number) Maximum Transmission Unit (MTU)

    &emsp;|Format    &emsp;|Description                        |
    |------------|-------------------------------------|
    &emsp;|68-16000  &emsp;|Maximum Transmission Unit in byte  |
- `parameters` (Attributes) Tunnel parameters (see [below for nested schema](#nestedatt--parameters))
- `redirect` (String) Redirect incoming packet to destination

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|txt     &emsp;|Destination interface name  |
- `remote` (String) Tunnel remote address

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|ipv4    &emsp;|Tunnel remote IPv4 address  |
    &emsp;|ipv6    &emsp;|Tunnel remote IPv6 address  |
- `source_address` (String) Source IP address used to initiate connection

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|ipv4    &emsp;|IPv4 source address  |
    &emsp;|ipv6    &emsp;|IPv6 source address  |
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

- `tunnel` (String) Tunnel interface

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|tunN    &emsp;|Tunnel interface name  |


&lt;a id=&#34;nestedatt--ip&#34;&gt;&lt;/a&gt;
### Nested Schema for `ip`

Optional:

- `adjust_mss` (String) Adjust TCP MSS value

    &emsp;|Format             &emsp;|Description                                     |
    |---------------------|--------------------------------------------------|
    &emsp;|clamp-mss-to-pmtu  &emsp;|Automatically sets the MSS to the proper value  |
    &emsp;|536-65535          &emsp;|TCP Maximum segment size in bytes               |
- `arp_cache_timeout` (Number) ARP cache entry timeout in seconds

    &emsp;|Format   &emsp;|Description                        |
    |-----------|-------------------------------------|
    &emsp;|1-86400  &emsp;|ARP cache entry timout in seconds  |
- `disable_arp_filter` (Boolean) Disable ARP filter on this interface
- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `enable_arp_accept` (Boolean) Enable ARP accept on this interface
- `enable_arp_announce` (Boolean) Enable ARP announce on this interface
- `enable_arp_ignore` (Boolean) Enable ARP ignore on this interface
- `enable_directed_broadcast` (Boolean) Enable directed broadcast forwarding on this interface
- `enable_proxy_arp` (Boolean) Enable proxy-arp on this interface
- `proxy_arp_pvlan` (Boolean) Enable private VLAN proxy ARP on this interface
- `source_validation` (String) Source validation by reversed path (RFC3704)

    &emsp;|Format   &emsp;|Description                                                  |
    |-----------|---------------------------------------------------------------|
    &emsp;|strict   &emsp;|Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    &emsp;|loose    &emsp;|Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    &emsp;|disable  &emsp;|No source validation                                         |


&lt;a id=&#34;nestedatt--ipv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6`

Optional:

- `accept_dad` (String) Accept Duplicate Address Detection

    &emsp;|Format  &emsp;|Description                                                                |
    |----------|-----------------------------------------------------------------------------|
    &emsp;|0       &emsp;|Disable DAD                                                                |
    &emsp;|1       &emsp;|Enable DAD                                                                 |
    &emsp;|2       &emsp;|Enable DAD - disable IPv6 if MAC-based duplicate link-local address found  |
- `address` (Attributes) IPv6 address configuration modes (see [below for nested schema](#nestedatt--ipv6--address))
- `adjust_mss` (String) Adjust TCP MSS value

    &emsp;|Format             &emsp;|Description                                     |
    |---------------------|--------------------------------------------------|
    &emsp;|clamp-mss-to-pmtu  &emsp;|Automatically sets the MSS to the proper value  |
    &emsp;|536-65535          &emsp;|TCP Maximum segment size in bytes               |
- `base_reachable_time` (Number) Base reachable time in seconds

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|1-86400  &emsp;|Base reachable time in seconds  |
- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `dup_addr_detect_transmits` (Number) Number of NS messages to send while performing DAD

    &emsp;|Format  &emsp;|Description                                         |
    |----------|------------------------------------------------------|
    &emsp;|0       &emsp;|Disable Duplicate Address Dectection (DAD)          |
    &emsp;|1-n     &emsp;|Number of NS messages to send while performing DAD  |
- `source_validation` (String) Source validation by reversed path (RFC3704)

    &emsp;|Format   &emsp;|Description                                                  |
    |-----------|---------------------------------------------------------------|
    &emsp;|strict   &emsp;|Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    &emsp;|loose    &emsp;|Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    &emsp;|disable  &emsp;|No source validation                                         |

&lt;a id=&#34;nestedatt--ipv6--address&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6.address`

Optional:

- `autoconf` (Boolean) Enable acquisition of IPv6 address using stateless autoconfig (SLAAC)
- `eui64` (List of String) Prefix for IPv6 address with MAC-based EUI-64

    &emsp;|Format                &emsp;|Description       |
    |------------------------|--------------------|
    &emsp;|&lt;h:h:h:h:h:h:h:h/64&gt;  &emsp;|IPv6 /64 network  |
- `no_default_link_local` (Boolean) Remove the default link-local address from the interface



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


&lt;a id=&#34;nestedatt--parameters&#34;&gt;&lt;/a&gt;
### Nested Schema for `parameters`

Optional:

- `erspan` (Attributes) ERSPAN tunnel parameters (see [below for nested schema](#nestedatt--parameters--erspan))
- `ip` (Attributes) IPv4-specific tunnel parameters (see [below for nested schema](#nestedatt--parameters--ip))
- `ipv6` (Attributes) IPv6-specific tunnel parameters (see [below for nested schema](#nestedatt--parameters--ipv6))

&lt;a id=&#34;nestedatt--parameters--erspan&#34;&gt;&lt;/a&gt;
### Nested Schema for `parameters.erspan`

Optional:

- `direction` (String) Mirrored traffic direction

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|ingress  &emsp;|Mirror ingress traffic  |
    &emsp;|egress   &emsp;|Mirror egress traffic   |
- `hw_id` (Number) Unique identifier of an ERSPAN engine within a system

    &emsp;|Format     &emsp;|Description                            |
    |-------------|-----------------------------------------|
    &emsp;|0-1048575  &emsp;|Unique identifier of an ERSPAN engine  |
- `index` (Number) ERSPAN version 1 index field

    &emsp;|Format  &emsp;|Description                                                       |
    |----------|--------------------------------------------------------------------|
    &emsp;|0-63    &emsp;|Platform-depedent field for specifying port number and direction  |
- `version` (String) Protocol version

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|1       &emsp;|ERSPAN Type II   |
    &emsp;|2       &emsp;|ERSPAN Type III  |


&lt;a id=&#34;nestedatt--parameters--ip&#34;&gt;&lt;/a&gt;
### Nested Schema for `parameters.ip`

Optional:

- `ignore_df` (Boolean) Ignore the DF (don&#39;t fragment) bit
- `key` (Number) Tunnel key (only GRE tunnels)

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|u32     &emsp;|Tunnel key   |
- `no_pmtu_discovery` (Boolean) Disable path MTU discovery
- `tos` (Number) Specifies TOS value to use in outgoing packets

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|0-99    &emsp;|Type of Service (TOS)  |
- `ttl` (Number) Specifies TTL value to use in outgoing packets

    &emsp;|Format  &emsp;|Description                                   |
    |----------|------------------------------------------------|
    &emsp;|0       &emsp;|Inherit - copy value from original IP header  |
    &emsp;|1-255   &emsp;|Time to Live                                  |


&lt;a id=&#34;nestedatt--parameters--ipv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `parameters.ipv6`

Optional:

- `encaplimit` (String) Set fixed encapsulation limit

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|0-255   &emsp;|Encapsulation limit          |
    &emsp;|none    &emsp;|Disable encapsulation limit  |
- `flowlabel` (String) Specifies the flow label to use in outgoing packets

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|inherit       &emsp;|Copy field from original header  |
    &emsp;|0x0-0x0fffff  &emsp;|Tunnel key, or hex value         |
- `hoplimit` (Number) Hoplimit

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|0-255   &emsp;|Hop limit    |
- `tclass` (String) Traffic class (Tclass)

    &emsp;|Format        &emsp;|Description                            |
    |----------------|-----------------------------------------|
    &emsp;|0x0-0x0fffff  &emsp;|Traffic class, &#39;inherit&#39; or hex value  |



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
