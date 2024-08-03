---
page_title: "vyos_interfaces_wireless Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Wireless (WiFi/WLAN) Network Interface
---

# vyos_interfaces_wireless (Resource)
<center>

*interfaces*  
⯯  
**Wireless (WiFi/WLAN) Network Interface**


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
- `bssid` (String) Basic Service Set Identifier (BSSID) - currently station mode only

    &emsp;|Format   &emsp;|Description          |
    |-----------|-----------------------|
    &emsp;|macaddr  &emsp;|BSSID (MAC) address  |
- `capabilities` (Attributes) HT and VHT capabilities for your card (see [below for nested schema](#nestedatt--capabilities))
- `channel` (String) Wireless radio channel

    &emsp;|Format  &emsp;|Description                        |
    |----------|-------------------------------------|
    &emsp;|0       &emsp;|Automatic Channel Selection (ACS)  |
    &emsp;|1-14    &emsp;|2.4Ghz (802.11 b/g/n) Channel      |
    &emsp;|34-173  &emsp;|5Ghz (802.11 a/h/j/n/ac) Channel   |
    &emsp;|1-233   &emsp;|6Ghz (802.11 ax) Channel           |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `dhcp_options` (Attributes) DHCP client settings/options (see [below for nested schema](#nestedatt--dhcp_options))
- `dhcpv6_options` (Attributes) DHCPv6 client settings/options (see [below for nested schema](#nestedatt--dhcpv6_options))
- `disable` (Boolean) Administratively disable interface
- `disable_broadcast_ssid` (Boolean) Disable broadcast of SSID from access-point
- `disable_link_detect` (Boolean) Ignore link state changes
- `enable_bf_protection` (Boolean) Beacon Protection: management frame protection for Beacon frames, requires Management Frame Protection (MFP)
- `expunge_failing_stations` (Boolean) Disassociate stations based on excessive transmission failures
- `hw_id` (String) Associate Ethernet Interface with given Media Access Control (MAC) address

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|macaddr  &emsp;|Hardware (MAC) address  |
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
- `ipv6` (Attributes) IPv6 routing parameters (see [below for nested schema](#nestedatt--ipv6))
- `isolate_stations` (Boolean) Isolate stations on the AP so they cannot see each other
- `mac` (String) Media Access Control (MAC) address

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|macaddr  &emsp;|Hardware (MAC) address  |
- `max_stations` (Number) Maximum number of wireless radio stations. Excess stations will be rejected upon authentication request.

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|1-2007  &emsp;|Number of allowed stations  |
- `mgmt_frame_protection` (String) Management Frame Protection (MFP) according to IEEE 802.11w

    &emsp;|Format    &emsp;|Description                        |
    |------------|-------------------------------------|
    &emsp;|disabled  &emsp;|no MFP                             |
    &emsp;|optional  &emsp;|MFP optional                       |
    &emsp;|required  &emsp;|MFP enforced (mandatory for WPA3)  |
- `mirror` (Attributes) Mirror ingress/egress packets (see [below for nested schema](#nestedatt--mirror))
- `mode` (String) Wireless radio mode

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|a       &emsp;|802.11a - 54 Mbits/sec        |
    &emsp;|b       &emsp;|802.11b - 11 Mbits/sec        |
    &emsp;|g       &emsp;|802.11g - 54 Mbits/sec        |
    &emsp;|n       &emsp;|802.11n - 600 Mbits/sec       |
    &emsp;|ac      &emsp;|802.11ac - 1300 Mbits/sec     |
    &emsp;|ax      &emsp;|802.11ax (6GHz only for now)  |
- `per_client_thread` (Boolean) Process traffic from each client in a dedicated thread
- `physical_device` (String) Wireless physical device
- `redirect` (String) Redirect incoming packet to destination

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|txt     &emsp;|Destination interface name  |
- `reduce_transmit_power` (Number) Transmission power reduction in dBm

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|0-255   &emsp;|TX power reduction in dBm  |
- `security` (Attributes) Wireless security settings (see [below for nested schema](#nestedatt--security))
- `ssid` (String) Wireless access-point service set identifier (SSID)
- `stationary_ap` (Boolean) Stationary AP config indicates that the AP doesn&#39;t move.
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `type` (String) Wireless device type for this interface

    &emsp;|Format        &emsp;|Description                                             |
    |----------------|----------------------------------------------------------|
    &emsp;|access-point  &emsp;|Access-point forwards packets between other nodes       |
    &emsp;|station       &emsp;|Connects to another access point                        |
    &emsp;|monitor       &emsp;|Passively monitor all packets on the frequency/channel  |
- `vrf` (String) VRF instance name

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `wireless` (String) Wireless (WiFi/WLAN) Network Interface

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|wlanN   &emsp;|Wireless (WiFi/WLAN) interface name  |


&lt;a id=&#34;nestedatt--capabilities&#34;&gt;&lt;/a&gt;
### Nested Schema for `capabilities`

Optional:

- `he` (Attributes) High Efficiency (HE) settings (see [below for nested schema](#nestedatt--capabilities--he))
- `ht` (Attributes) High Throughput (HT) settings (see [below for nested schema](#nestedatt--capabilities--ht))
- `require_he` (Boolean) Require stations to support HE PHY
- `require_ht` (Boolean) Require stations to support HT PHY
- `require_vht` (Boolean) Require stations to support VHT PHY
- `vht` (Attributes) Very High Throughput (VHT) settings (see [below for nested schema](#nestedatt--capabilities--vht))

&lt;a id=&#34;nestedatt--capabilities--he&#34;&gt;&lt;/a&gt;
### Nested Schema for `capabilities.he`

Optional:

- `antenna_pattern_fixed` (Boolean) Tell the AP that antenna positions are fixed and will not change during the lifetime of an association
- `beamform` (Attributes) HE beamforming capabilities (see [below for nested schema](#nestedatt--capabilities--he--beamform))
- `bss_color` (String) BSS coloring helps to prevent channel jamming when multiple APs use the same channels
- `center_channel_freq` (Attributes) HE operating channel center frequency (see [below for nested schema](#nestedatt--capabilities--he--center_channel_freq))
- `channel_set_width` (String) HE operating channel width

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|131     &emsp;|20 MHz channel width     |
    &emsp;|132     &emsp;|40 MHz channel width     |
    &emsp;|133     &emsp;|80 MHz channel width     |
    &emsp;|134     &emsp;|160 MHz channel width    |
    &emsp;|135     &emsp;|80+80 MHz channel width  |

&lt;a id=&#34;nestedatt--capabilities--he--beamform&#34;&gt;&lt;/a&gt;
### Nested Schema for `capabilities.he.beamform`

Optional:

- `multi_user_beamformer` (Boolean) Support for operation as multi user beamformer
- `single_user_beamformee` (Boolean) Support for operation as single user beamformee
- `single_user_beamformer` (Boolean) Support for operation as single user beamformer


&lt;a id=&#34;nestedatt--capabilities--he--center_channel_freq&#34;&gt;&lt;/a&gt;
### Nested Schema for `capabilities.he.center_channel_freq`

Optional:

- `freq_1` (Number) HE operating channel center frequency - center freq 1 (for use with 80, 80+80 and 160 modes)

    &emsp;|Format  &emsp;|Description                                                                                                     |
    |----------|------------------------------------------------------------------------------------------------------------------|
    &emsp;|1-233   &emsp;|6Ghz (802.11 ax) center channel index (use 3 (at 40MHz), 7 (at 80MHz) or 15 (at 160MHz) for primary channel 1)  |
- `freq_2` (Number) HE operating channel center frequency - center freq 2 (for use with the 80+80 mode)

    &emsp;|Format  &emsp;|Description                                                                         |
    |----------|--------------------------------------------------------------------------------------|
    &emsp;|1-233   &emsp;|6Ghz (802.11 ax) center channel index (use 23 (at 80MHz) for secondary channel 17)  |



&lt;a id=&#34;nestedatt--capabilities--ht&#34;&gt;&lt;/a&gt;
### Nested Schema for `capabilities.ht`

Optional:

- `_40mhz_incapable` (Boolean) 40MHz intolerance, use 20MHz only!
- `auto_powersave` (Boolean) Enable WMM-PS unscheduled automatic power save delivery [U-APSD]
- `channel_set_width` (List of String) Supported channel set width

    &emsp;|Format  &emsp;|Description                                                                                      |
    |----------|---------------------------------------------------------------------------------------------------|
    &emsp;|ht20    &emsp;|Supported channel set width both 20 MHz only                                                     |
    &emsp;|ht40+   &emsp;|Supported channel set width both 20 MHz and 40 MHz with secondary channel above primary channel  |
    &emsp;|ht40-   &emsp;|Supported channel set width both 20 MHz and 40 MHz with secondary channel below primary channel  |
- `delayed_block_ack` (Boolean) Enable HT-delayed block ack
- `dsss_cck_40` (Boolean) Enable DSSS_CCK-40
- `greenfield` (Boolean) Enable HT-greenfield
- `ldpc` (Boolean) Enable LDPC coding capability
- `lsig_protection` (Boolean) Enable L-SIG TXOP protection capability
- `max_amsdu` (String) Set maximum A-MSDU length

    &emsp;|Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    &emsp;|3839    &emsp;|Set maximum A-MSDU length to 3839 octets  |
    &emsp;|7935    &emsp;|Set maximum A-MSDU length to 7935 octets  |
- `short_gi` (List of String) Short GI capabilities

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|20      &emsp;|Short GI for 20 MHz  |
    &emsp;|40      &emsp;|Short GI for 40 MHz  |
- `smps` (String) Spatial Multiplexing Power Save (SMPS) settings

    &emsp;|Format   &emsp;|Description                                   |
    |-----------|------------------------------------------------|
    &emsp;|static   &emsp;|STATIC Spatial Multiplexing (SM) Power Save   |
    &emsp;|dynamic  &emsp;|DYNAMIC Spatial Multiplexing (SM) Power Save  |
- `stbc` (Attributes) Support for sending and receiving PPDU using STBC (Space Time Block Coding) (see [below for nested schema](#nestedatt--capabilities--ht--stbc))

&lt;a id=&#34;nestedatt--capabilities--ht--stbc&#34;&gt;&lt;/a&gt;
### Nested Schema for `capabilities.ht.stbc`

Optional:

- `rx` (String) Enable receiving PPDU using STBC (Space Time Block Coding)

    &emsp;|Format  &emsp;|Description                                     |
    |----------|--------------------------------------------------|
    &emsp;|[1-3]+  &emsp;|Number of spacial streams that can use RX STBC  |
- `tx` (Boolean) Enable sending PPDU using STBC (Space Time Block Coding)



&lt;a id=&#34;nestedatt--capabilities--vht&#34;&gt;&lt;/a&gt;
### Nested Schema for `capabilities.vht`

Optional:

- `antenna_count` (Number) Number of antennas on this card

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|1-8     &emsp;|Number of antennas for this card  |
- `antenna_pattern_fixed` (Boolean) Set if antenna pattern does not change during the lifetime of an association
- `beamform` (List of String) VHT beamforming capabilities

    &emsp;|Format                  &emsp;|Description                                      |
    |--------------------------|---------------------------------------------------|
    &emsp;|single-user-beamformer  &emsp;|Support for operation as single user beamformer  |
    &emsp;|single-user-beamformee  &emsp;|Support for operation as single user beamformee  |
    &emsp;|multi-user-beamformer   &emsp;|Support for operation as multi user beamformer   |
    &emsp;|multi-user-beamformee   &emsp;|Support for operation as multi user beamformee   |
- `center_channel_freq` (Attributes) VHT operating channel center frequency (see [below for nested schema](#nestedatt--capabilities--vht--center_channel_freq))
- `channel_set_width` (String) VHT operating Channel width

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|0       &emsp;|20 or 40 MHz channel width  |
    &emsp;|1       &emsp;|80 MHz channel width        |
    &emsp;|2       &emsp;|160 MHz channel width       |
    &emsp;|3       &emsp;|80+80 MHz channel width     |
- `ldpc` (Boolean) Enable LDPC (Low Density Parity Check) coding capability
- `link_adaptation` (String) VHT link adaptation capabilities

    &emsp;|Format       &emsp;|Description                                                                 |
    |---------------|------------------------------------------------------------------------------|
    &emsp;|unsolicited  &emsp;|Station provides only unsolicited VHT MFB                                   |
    &emsp;|both         &emsp;|Station can provide VHT MFB in response to VHT MRQ and unsolicited VHT MFB  |
- `max_mpdu` (String) Increase Maximum MPDU length to 7991 or 11454 octets (otherwise: 3895 octets)

    &emsp;|Format  &emsp;|Description                                  |
    |----------|-----------------------------------------------|
    &emsp;|7991    &emsp;|ncrease Maximum MPDU length to 7991 octets   |
    &emsp;|11454   &emsp;|ncrease Maximum MPDU length to 11454 octets  |
- `max_mpdu_exp` (Number) Set the maximum length of A-MPDU pre-EOF padding that the station can receive

    &emsp;|Format  &emsp;|Description                                                         |
    |----------|----------------------------------------------------------------------|
    &emsp;|0-7     &emsp;|Maximum length of A-MPDU pre-EOF padding = 2 pow(13 + x) -1 octets  |
- `short_gi` (List of String) Short GI capabilities

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|80      &emsp;|Short GI for 80 MHz   |
    &emsp;|160     &emsp;|Short GI for 160 MHz  |
- `stbc` (Attributes) Support for sending and receiving PPDU using STBC (Space Time Block Coding) (see [below for nested schema](#nestedatt--capabilities--vht--stbc))
- `tx_powersave` (Boolean) Enable VHT TXOP Power Save Mode
- `vht_cf` (Boolean) Station supports receiving VHT variant HT Control field

&lt;a id=&#34;nestedatt--capabilities--vht--center_channel_freq&#34;&gt;&lt;/a&gt;
### Nested Schema for `capabilities.vht.center_channel_freq`

Optional:

- `freq_1` (Number) VHT operating channel center frequency - center freq 1 (for use with 80, 80+80 and 160 modes)

    &emsp;|Format  &emsp;|Description                                                                          |
    |----------|---------------------------------------------------------------------------------------|
    &emsp;|34-173  &emsp;|5Ghz (802.11 a/h/j/n/ac) center channel index (use 42 for primary 80MHz channel 36)  |
- `freq_2` (Number) VHT operating channel center frequency - center freq 2 (for use with the 80+80 mode)

    &emsp;|Format  &emsp;|Description                                                                    |
    |----------|---------------------------------------------------------------------------------|
    &emsp;|34-173  &emsp;|5Ghz (802.11 ac) center channel index (use 58 for secondary 80MHz channel 52)  |


&lt;a id=&#34;nestedatt--capabilities--vht--stbc&#34;&gt;&lt;/a&gt;
### Nested Schema for `capabilities.vht.stbc`

Optional:

- `rx` (String) Enable receiving PPDU using STBC (Space Time Block Coding)

    &emsp;|Format  &emsp;|Description                                     |
    |----------|--------------------------------------------------|
    &emsp;|[1-4]+  &emsp;|Number of spacial streams that can use RX STBC  |
- `tx` (Boolean) Enable sending PPDU using STBC (Space Time Block Coding)




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

- `station_address` (Attributes) Station MAC address based authentication (see [below for nested schema](#nestedatt--security--station_address))
- `wep` (Attributes) Wired Equivalent Privacy (WEP) parameters (see [below for nested schema](#nestedatt--security--wep))
- `wpa` (Attributes) Wifi Protected Access (WPA) parameters (see [below for nested schema](#nestedatt--security--wpa))

&lt;a id=&#34;nestedatt--security--station_address&#34;&gt;&lt;/a&gt;
### Nested Schema for `security.station_address`

Optional:

- `accept` (Attributes) Accept station MAC address (see [below for nested schema](#nestedatt--security--station_address--accept))
- `deny` (Attributes) Deny station MAC address (see [below for nested schema](#nestedatt--security--station_address--deny))
- `mode` (String) Select security operation mode

    &emsp;|Format  &emsp;|Description                                   |
    |----------|------------------------------------------------|
    &emsp;|accept  &emsp;|Accept all clients unless found in deny list  |
    &emsp;|deny    &emsp;|Deny all clients unless found in accept list  |

&lt;a id=&#34;nestedatt--security--station_address--accept&#34;&gt;&lt;/a&gt;
### Nested Schema for `security.station_address.accept`

Optional:

- `mac` (List of String) Media Access Control (MAC) address

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|macaddr  &emsp;|Hardware (MAC) address  |


&lt;a id=&#34;nestedatt--security--station_address--deny&#34;&gt;&lt;/a&gt;
### Nested Schema for `security.station_address.deny`

Optional:

- `mac` (List of String) Media Access Control (MAC) address

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|macaddr  &emsp;|Hardware (MAC) address  |



&lt;a id=&#34;nestedatt--security--wep&#34;&gt;&lt;/a&gt;
### Nested Schema for `security.wep`

Optional:

- `key` (List of String) WEP encryption key

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|txt     &emsp;|Wired Equivalent Privacy key  |


&lt;a id=&#34;nestedatt--security--wpa&#34;&gt;&lt;/a&gt;
### Nested Schema for `security.wpa`

Optional:

- `cipher` (List of String) Cipher suite for WPA unicast packets

    &emsp;|Format    &emsp;|Description                                                                                 |
    |------------|----------------------------------------------------------------------------------------------|
    &emsp;|GCMP-256  &emsp;|AES in Galois/counter mode with 256-bit key                                                 |
    &emsp;|GCMP      &emsp;|AES in Galois/counter mode with 128-bit key                                                 |
    &emsp;|CCMP-256  &emsp;|AES in Counter mode with CBC-MAC with 256-bit key                                           |
    &emsp;|CCMP      &emsp;|AES in Counter mode with CBC-MAC [RFC 3610, IEEE 802.11i/D7.0] (supported on all WPA2 APs)  |
    &emsp;|TKIP      &emsp;|Temporal Key Integrity Protocol [IEEE 802.11i/D7.0]                                         |
- `group_cipher` (String) Cipher suite for WPA multicast and broadcast packets

    &emsp;|Format    &emsp;|Description                                                                                 |
    |------------|----------------------------------------------------------------------------------------------|
    &emsp;|GCMP-256  &emsp;|AES in Galois/counter mode with 256-bit key                                                 |
    &emsp;|GCMP      &emsp;|AES in Galois/counter mode with 128-bit key                                                 |
    &emsp;|CCMP-256  &emsp;|AES in Counter mode with CBC-MAC with 256-bit key                                           |
    &emsp;|CCMP      &emsp;|AES in Counter mode with CBC-MAC [RFC 3610, IEEE 802.11i/D7.0] (supported on all WPA2 APs)  |
    &emsp;|TKIP      &emsp;|Temporal Key Integrity Protocol [IEEE 802.11i/D7.0]                                         |
- `group_mgmt_cipher` (String) Group management cipher suite. All the stations connecting to the BSS will also need to support the selected cipher
- `mode` (String) WPA mode

    &emsp;|Format    &emsp;|Description                                                                        |
    |------------|-------------------------------------------------------------------------------------|
    &emsp;|wpa       &emsp;|WPA (IEEE 802.11i/D3.0)                                                            |
    &emsp;|wpa2      &emsp;|WPA2 (full IEEE 802.11i/RSN)                                                       |
    &emsp;|wpa+wpa2  &emsp;|Allow both WPA and WPA2                                                            |
    &emsp;|wpa3      &emsp;|WPA3 (required for 802.11ax, you must also set mgmt-frame-protection as required)  |
- `passphrase` (String) WPA passphrase. If you are using special characters in the WPA passphrase then single quotes are required.

    &emsp;|Format  &emsp;|Description                                                                                                                |
    |----------|-----------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt     &emsp;|Passphrase of at least 8 but not more than 63 printable characters for WPA-Personal and any passphrase for WPA-Enterprise  |
- `radius` (Attributes) RADIUS based user authentication (see [below for nested schema](#nestedatt--security--wpa--radius))
- `username` (String) Username used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Username     |

&lt;a id=&#34;nestedatt--security--wpa--radius&#34;&gt;&lt;/a&gt;
### Nested Schema for `security.wpa.radius`

Optional:

- `source_address` (String) IPv4 source address used to initiate connection

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|ipv4    &emsp;|IPv4 source address  |




&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
