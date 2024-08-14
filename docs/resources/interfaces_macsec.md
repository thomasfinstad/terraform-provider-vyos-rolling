---
page_title: "vyos_interfaces_macsec Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯MACsec Interface (802.1ae)
---

# vyos_interfaces_macsec (Resource)
<center>

*interfaces*  
⯯  
**MACsec Interface (802.1ae)**


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
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
- `ipv6` (Attributes) IPv6 routing parameters (see [below for nested schema](#nestedatt--ipv6))
- `mirror` (Attributes) Mirror ingress/egress packets (see [below for nested schema](#nestedatt--mirror))
- `mtu` (Number) Maximum Transmission Unit (MTU)

    &emsp;|Format    &emsp;|Description                        |
    |------------|-------------------------------------|
    &emsp;|68-16000  &emsp;|Maximum Transmission Unit in byte  |
- `redirect` (String) Redirect incoming packet to destination

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|txt     &emsp;|Destination interface name  |
- `security` (Attributes) Security/Encryption Settings (see [below for nested schema](#nestedatt--security))
- `source_interface` (String) Physical interface the traffic will go through

    &emsp;|Format     &emsp;|Description                                     |
    |-------------|--------------------------------------------------|
    &emsp;|interface  &emsp;|Physical interface used for traffic forwarding  |
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

- `macsec` (String) MACsec Interface (802.1ae)

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|macsecN  &emsp;|MACsec interface name  |


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


&lt;a id=&#34;nestedatt--security&#34;&gt;&lt;/a&gt;
### Nested Schema for `security`

Optional:

- `cipher` (String) Cipher suite used

    &emsp;|Format       &emsp;|Description                                         |
    |---------------|------------------------------------------------------|
    &emsp;|gcm-aes-128  &emsp;|Galois/Counter Mode of AES cipher with 128-bit key  |
    &emsp;|gcm-aes-256  &emsp;|Galois/Counter Mode of AES cipher with 256-bit key  |
- `encrypt` (Boolean) Enable optional MACsec encryption
- `mka` (Attributes) MACsec Key Agreement protocol (MKA) (see [below for nested schema](#nestedatt--security--mka))
- `replay_window` (Number) IEEE 802.1X/MACsec replay protection window

    &emsp;|Format        &emsp;|Description                                 |
    |----------------|----------------------------------------------|
    &emsp;|0             &emsp;|No replay window, strict check              |
    &emsp;|1-4294967295  &emsp;|Number of packets that could be misordered  |
- `static` (Attributes) Use static keys for MACsec [static Secure Authentication Key (SAK) mode] (see [below for nested schema](#nestedatt--security--static))

&lt;a id=&#34;nestedatt--security--mka&#34;&gt;&lt;/a&gt;
### Nested Schema for `security.mka`

Optional:

- `cak` (String) Secure Connectivity Association Key

    &emsp;|Format  &emsp;|Description                                                                                                                   |
    |----------|--------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt     &emsp;|16-byte (128-bit) hex-string (32 hex-digits) for gcm-aes-128 or 32-byte (256-bit) hex-string (64 hex-digits) for gcm-aes-256  |
- `ckn` (String) Secure Connectivity Association Key Name

    &emsp;|Format  &emsp;|Description                                             |
    |----------|----------------------------------------------------------|
    &emsp;|txt     &emsp;|1..32-bytes (8..256 bit) hex-string (2..64 hex-digits)  |
- `priority` (Number) Priority of MACsec Key Agreement protocol (MKA) actor

    &emsp;|Format  &emsp;|Description                                   |
    |----------|------------------------------------------------|
    &emsp;|0-255   &emsp;|MACsec Key Agreement protocol (MKA) priority  |


&lt;a id=&#34;nestedatt--security--static&#34;&gt;&lt;/a&gt;
### Nested Schema for `security.static`

Optional:

- `key` (String) MACsec static key

    &emsp;|Format  &emsp;|Description                                                                                                                   |
    |----------|--------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt     &emsp;|16-byte (128-bit) hex-string (32 hex-digits) for gcm-aes-128 or 32-byte (256-bit) hex-string (64 hex-digits) for gcm-aes-256  |



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  