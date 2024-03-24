---
page_title: "vyos_firewall_ipv6_name_rule Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall⯯IPv6 firewall⯯IPv6 custom firewall⯯IPv6 Firewall custom rule number
---

# vyos_firewall_ipv6_name_rule (Resource)
<center>

Firewall  
⯯  
IPv6 firewall  
⯯  
IPv6 custom firewall  
⯯  
**IPv6 Firewall custom rule number**


</center>

## Schema

### Required

- `name_id` (String) IPv6 custom firewall
- `rule_id` (Number) IPv6 Firewall custom rule number

    &emsp;|Format    &emsp;|Description                    |
    |------------|---------------------------------|
    &emsp;|1-999999  &emsp;|Number for this firewall rule  |

### Optional

- `action` (String) Rule action

    &emsp;|Format    &emsp;|Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    &emsp;|accept    &emsp;|Accept matching entries                                                        |
    &emsp;|continue  &emsp;|Continue parsing next rule                                                     |
    &emsp;|jump      &emsp;|Jump to another chain                                                          |
    &emsp;|reject    &emsp;|Reject matching entries                                                        |
    &emsp;|return    &emsp;|Return from the current chain and continue at the next rule of the last chain  |
    &emsp;|drop      &emsp;|Drop matching entries                                                          |
    &emsp;|queue     &emsp;|Enqueue packet to userspace                                                    |
    &emsp;|offload   &emsp;|Offload packet via flowtable                                                   |
    &emsp;|synproxy  &emsp;|Synproxy connections                                                           |
- `add_address_to_group` (Attributes) Add ipv6 address to dynamic ipv6-address-group (see [below for nested schema](#nestedatt--add_address_to_group))
- `connection_mark` (List of Number) Connection mark

    &emsp;|Format        &emsp;|Description               |
    |----------------|----------------------------|
    &emsp;|0-2147483647  &emsp;|Connection-mark to match  |
- `connection_status` (Attributes) Connection status (see [below for nested schema](#nestedatt--connection_status))
- `conntrack_helper` (List of String) Match related traffic from conntrack helpers

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|ftp     &emsp;|Related traffic from FTP helper     |
    &emsp;|h323    &emsp;|Related traffic from H.323 helper   |
    &emsp;|pptp    &emsp;|Related traffic from PPTP helper    |
    &emsp;|nfs     &emsp;|Related traffic from NFS helper     |
    &emsp;|rtsp    &emsp;|Related traffic from RTSP helper    |
    &emsp;|sip     &emsp;|Related traffic from SIP helper     |
    &emsp;|tftp    &emsp;|Related traffic from TFTP helper    |
    &emsp;|sqlnet  &emsp;|Related traffic from SQLNet helper  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `destination` (Attributes) Destination parameters (see [below for nested schema](#nestedatt--destination))
- `disable` (Boolean) Disable instance
- `dscp` (List of String) DSCP value

    &emsp;|Format       &emsp;|Description          |
    |---------------|-----------------------|
    &emsp;|0-63         &emsp;|DSCP value to match  |
    &emsp;|&lt;start-end&gt;  &emsp;|DSCP range to match  |
- `dscp_exclude` (List of String) DSCP value not to match

    &emsp;|Format       &emsp;|Description              |
    |---------------|---------------------------|
    &emsp;|0-63         &emsp;|DSCP value not to match  |
    &emsp;|&lt;start-end&gt;  &emsp;|DSCP range not to match  |
- `fragment` (Attributes) IP fragment match (see [below for nested schema](#nestedatt--fragment))
- `hop_limit` (Attributes) Hop limit (see [below for nested schema](#nestedatt--hop_limit))
- `icmpv6` (Attributes) ICMPv6 type and code information (see [below for nested schema](#nestedatt--icmpv6))
- `inbound_interface` (Attributes) Match inbound-interface (see [below for nested schema](#nestedatt--inbound_interface))
- `ipsec` (Attributes) Inbound IPsec packets (see [below for nested schema](#nestedatt--ipsec))
- `jump_target` (String) Set jump target. Action jump must be defined to use this setting
- `limit` (Attributes) Rate limit using a token bucket filter (see [below for nested schema](#nestedatt--limit))
- `log` (Boolean) Log packets hitting this rule
- `log_options` (Attributes) Log options (see [below for nested schema](#nestedatt--log_options))
- `mark` (String) Firewall mark

    &emsp;|Format         &emsp;|Description                            |
    |-----------------|-----------------------------------------|
    &emsp;|0-2147483647   &emsp;|Firewall mark to match                 |
    &emsp;|!0-2147483647  &emsp;|Inverted Firewall mark to match        |
    &emsp;|&lt;start-end&gt;    &emsp;|Firewall mark range to match           |
    &emsp;|!&lt;start-end&gt;   &emsp;|Firewall mark inverted range to match  |
- `offload_target` (String) Set flowtable offload target. Action offload must be defined to use this setting
- `outbound_interface` (Attributes) Match outbound-interface (see [below for nested schema](#nestedatt--outbound_interface))
- `packet_length` (List of String) Payload size in bytes, including header and data to match

    &emsp;|Format       &emsp;|Description                   |
    |---------------|--------------------------------|
    &emsp;|1-65535      &emsp;|Packet length to match        |
    &emsp;|&lt;start-end&gt;  &emsp;|Packet length range to match  |
- `packet_length_exclude` (List of String) Payload size in bytes, including header and data not to match

    &emsp;|Format       &emsp;|Description                       |
    |---------------|------------------------------------|
    &emsp;|1-65535      &emsp;|Packet length not to match        |
    &emsp;|&lt;start-end&gt;  &emsp;|Packet length range not to match  |
- `packet_type` (String) Packet type

    &emsp;|Format     &emsp;|Description                                      |
    |-------------|---------------------------------------------------|
    &emsp;|broadcast  &emsp;|Match broadcast packet type                      |
    &emsp;|host       &emsp;|Match host packet type, addressed to local host  |
    &emsp;|multicast  &emsp;|Match multicast packet type                      |
    &emsp;|other      &emsp;|Match packet addressed to another host           |
- `protocol` (String) Protocol to match (protocol name, number, or &#34;all&#34;)

    &emsp;|Format       &emsp;|Description         |
    |---------------|----------------------|
    &emsp;|all          &emsp;|All IP protocols    |
    &emsp;|tcp_udp      &emsp;|Both TCP and UDP    |
    &emsp;|0-255        &emsp;|IP protocol number  |
    &emsp;|&lt;protocol&gt;   &emsp;|IP protocol name    |
    &emsp;|!&lt;protocol&gt;  &emsp;|IP protocol name    |
- `queue` (Number) Queue target to use. Action queue must be defined to use this setting

    &emsp;|Format   &emsp;|Description   |
    |-----------|----------------|
    &emsp;|0-65535  &emsp;|Queue target  |
- `queue_options` (List of String) Options used for queue target. Action queue must be defined to use this setting

    &emsp;|Format  &emsp;|Description                                                      |
    |----------|-------------------------------------------------------------------|
    &emsp;|bypass  &emsp;|Let packets go through if userspace application cannot back off  |
    &emsp;|fanout  &emsp;|Distribute packets between several queues                        |
- `recent` (Attributes) Parameters for matching recently seen sources (see [below for nested schema](#nestedatt--recent))
- `source` (Attributes) Source parameters (see [below for nested schema](#nestedatt--source))
- `state` (List of String) Session state

    &emsp;|Format       &emsp;|Description        |
    |---------------|---------------------|
    &emsp;|established  &emsp;|Established state  |
    &emsp;|invalid      &emsp;|Invalid state      |
    &emsp;|new          &emsp;|New state          |
    &emsp;|related      &emsp;|Related state      |
- `synproxy` (Attributes) Synproxy options (see [below for nested schema](#nestedatt--synproxy))
- `tcp` (Attributes) TCP options to match (see [below for nested schema](#nestedatt--tcp))
- `time` (Attributes) Time to match rule (see [below for nested schema](#nestedatt--time))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--add_address_to_group&#34;&gt;&lt;/a&gt;
### Nested Schema for `add_address_to_group`

Optional:

- `destination_address` (Attributes) Add destination ipv6 addresses to dynamic ipv6-address-group (see [below for nested schema](#nestedatt--add_address_to_group--destination_address))
- `source_address` (Attributes) Add source ipv6 addresses to dynamic ipv6-address-group (see [below for nested schema](#nestedatt--add_address_to_group--source_address))

&lt;a id=&#34;nestedatt--add_address_to_group--destination_address&#34;&gt;&lt;/a&gt;
### Nested Schema for `add_address_to_group.destination_address`

Optional:

- `address_group` (String) Dynamic ipv6-address-group
- `timeout` (String) Set timeout

    &emsp;|Format     &emsp;|Description               |
    |-------------|----------------------------|
    &emsp;|&lt;number&gt;s  &emsp;|Timeout value in seconds  |
    &emsp;|&lt;number&gt;m  &emsp;|Timeout value in minutes  |
    &emsp;|&lt;number&gt;h  &emsp;|Timeout value in hours    |
    &emsp;|&lt;number&gt;d  &emsp;|Timeout value in days     |


&lt;a id=&#34;nestedatt--add_address_to_group--source_address&#34;&gt;&lt;/a&gt;
### Nested Schema for `add_address_to_group.source_address`

Optional:

- `address_group` (String) Dynamic ipv6-address-group
- `timeout` (String) Set timeout

    &emsp;|Format     &emsp;|Description               |
    |-------------|----------------------------|
    &emsp;|&lt;number&gt;s  &emsp;|Timeout value in seconds  |
    &emsp;|&lt;number&gt;m  &emsp;|Timeout value in minutes  |
    &emsp;|&lt;number&gt;h  &emsp;|Timeout value in hours    |
    &emsp;|&lt;number&gt;d  &emsp;|Timeout value in days     |



&lt;a id=&#34;nestedatt--connection_status&#34;&gt;&lt;/a&gt;
### Nested Schema for `connection_status`

Optional:

- `nat` (String) NAT connection status

    &emsp;|Format       &emsp;|Description                                            |
    |---------------|---------------------------------------------------------|
    &emsp;|destination  &emsp;|Match connections that are subject to destination NAT  |
    &emsp;|source       &emsp;|Match connections that are subject to source NAT       |


&lt;a id=&#34;nestedatt--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination`

Optional:

- `address` (String) IP address, subnet, or range

    &emsp;|Format      &emsp;|Description                                    |
    |--------------|-------------------------------------------------|
    &emsp;|ipv6        &emsp;|IP address to match                            |
    &emsp;|ipv6net     &emsp;|Subnet to match                                |
    &emsp;|ipv6range   &emsp;|IP range to match                              |
    &emsp;|!ipv6       &emsp;|Match everything except the specified address  |
    &emsp;|!ipv6net    &emsp;|Match everything except the specified prefix   |
    &emsp;|!ipv6range  &emsp;|Match everything except the specified range    |
- `address_mask` (String) IP mask

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|ipv6    &emsp;|IP mask to apply  |
- `fqdn` (String) Fully qualified domain name

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|&lt;fqdn&gt;  &emsp;|Fully qualified domain name  |
- `geoip` (Attributes) GeoIP options - Data provided by DB-IP.com (see [below for nested schema](#nestedatt--destination--geoip))
- `group` (Attributes) Group (see [below for nested schema](#nestedatt--destination--group))
- `mac_address` (String) MAC address

    &emsp;|Format    &emsp;|Description                                        |
    |------------|-----------------------------------------------------|
    &emsp;|macaddr   &emsp;|MAC address to match                               |
    &emsp;|!macaddr  &emsp;|Match everything except the specified MAC address  |
- `port` (String) Port

    &emsp;|Format       &emsp;|Description                                                                                                               |
    |---------------|----------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt          &emsp;|Named port (any name in /etc/services, e.g., http)                                                                        |
    &emsp;|1-65535      &emsp;|Numbered port                                                                                                             |
    &emsp;|&lt;start-end&gt;  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                      |
    &emsp;|             &emsp;|\n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: &#39;telnet,http,123,1001-1005&#39;  |

&lt;a id=&#34;nestedatt--destination--geoip&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination.geoip`

Optional:

- `country_code` (List of String) GeoIP country code

    &emsp;|Format     &emsp;|Description                  |
    |-------------|-------------------------------|
    &emsp;|&lt;country&gt;  &emsp;|Country code (2 characters)  |
- `inverse_match` (Boolean) Inverse match of country-codes


&lt;a id=&#34;nestedatt--destination--group&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination.group`

Optional:

- `address_group` (String) Group of addresses
- `domain_group` (String) Group of domains
- `dynamic_address_group` (String) Group of dynamic ipv6 addresses
- `mac_group` (String) Group of MAC addresses
- `network_group` (String) Group of networks
- `port_group` (String) Group of ports



&lt;a id=&#34;nestedatt--fragment&#34;&gt;&lt;/a&gt;
### Nested Schema for `fragment`

Optional:

- `match_frag` (Boolean) Second and further fragments of fragmented packets
- `match_non_frag` (Boolean) Head fragments or unfragmented packets


&lt;a id=&#34;nestedatt--hop_limit&#34;&gt;&lt;/a&gt;
### Nested Schema for `hop_limit`

Optional:

- `eq` (Number) Match on equal value

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|0-255   &emsp;|Equal to value  |
- `gt` (Number) Match on greater then value

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|0-255   &emsp;|Greater then value  |
- `lt` (Number) Match on less then value

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|0-255   &emsp;|Less then value  |


&lt;a id=&#34;nestedatt--icmpv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `icmpv6`

Optional:

- `code` (Number) ICMPv6 code

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|0-255   &emsp;|ICMPv6 code (0-255)  |
- `type` (Number) ICMPv6 type

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|0-255   &emsp;|ICMPv6 type (0-255)  |
- `type_name` (String) ICMPv6 type-name

    &emsp;|Format                   &emsp;|Description                              |
    |---------------------------|-------------------------------------------|
    &emsp;|destination-unreachable  &emsp;|ICMPv6 type 1: destination-unreachable   |
    &emsp;|packet-too-big           &emsp;|ICMPv6 type 2: packet-too-big            |
    &emsp;|time-exceeded            &emsp;|ICMPv6 type 3: time-exceeded             |
    &emsp;|echo-request             &emsp;|ICMPv6 type 128: echo-request            |
    &emsp;|echo-reply               &emsp;|ICMPv6 type 129: echo-reply              |
    &emsp;|mld-listener-query       &emsp;|ICMPv6 type 130: mld-listener-query      |
    &emsp;|mld-listener-report      &emsp;|ICMPv6 type 131: mld-listener-report     |
    &emsp;|mld-listener-reduction   &emsp;|ICMPv6 type 132: mld-listener-reduction  |
    &emsp;|nd-router-solicit        &emsp;|ICMPv6 type 133: nd-router-solicit       |
    &emsp;|nd-router-advert         &emsp;|ICMPv6 type 134: nd-router-advert        |
    &emsp;|nd-neighbor-solicit      &emsp;|ICMPv6 type 135: nd-neighbor-solicit     |
    &emsp;|nd-neighbor-advert       &emsp;|ICMPv6 type 136: nd-neighbor-advert      |
    &emsp;|nd-redirect              &emsp;|ICMPv6 type 137: nd-redirect             |
    &emsp;|parameter-problem        &emsp;|ICMPv6 type 4: parameter-problem         |
    &emsp;|router-renumbering       &emsp;|ICMPv6 type 138: router-renumbering      |
    &emsp;|ind-neighbor-solicit     &emsp;|ICMPv6 type 141: ind-neighbor-solicit    |
    &emsp;|ind-neighbor-advert      &emsp;|ICMPv6 type 142: ind-neighbor-advert     |
    &emsp;|mld2-listener-report     &emsp;|ICMPv6 type 143: mld2-listener-report    |


&lt;a id=&#34;nestedatt--inbound_interface&#34;&gt;&lt;/a&gt;
### Nested Schema for `inbound_interface`

Optional:

- `group` (String) Match interface-group

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|txt     &emsp;|Interface-group name to match           |
    &emsp;|!txt    &emsp;|Inverted interface-group name to match  |
- `name` (String) Match interface

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|txt     &emsp;|Interface name                    |
    &emsp;|txt&amp;    &emsp;|Interface name with wildcard      |
    &emsp;|!txt    &emsp;|Inverted interface name to match  |


&lt;a id=&#34;nestedatt--ipsec&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipsec`

Optional:

- `match_ipsec` (Boolean) Inbound IPsec packets
- `match_none` (Boolean) Inbound non-IPsec packets


&lt;a id=&#34;nestedatt--limit&#34;&gt;&lt;/a&gt;
### Nested Schema for `limit`

Optional:

- `burst` (Number) Maximum number of packets to allow in excess of rate

    &emsp;|Format        &emsp;|Description                                           |
    |----------------|--------------------------------------------------------|
    &emsp;|0-4294967295  &emsp;|Maximum number of packets to allow in excess of rate  |
- `rate` (String) Maximum average matching rate

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|txt     &emsp;|integer/unit (Example: 5/minute)  |


&lt;a id=&#34;nestedatt--log_options&#34;&gt;&lt;/a&gt;
### Nested Schema for `log_options`

Optional:

- `group` (Number) Set log group

    &emsp;|Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    &emsp;|0-65535  &emsp;|Log group to send messages to  |
- `level` (String) Set log-level

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|emerg   &emsp;|Emerg log level     |
    &emsp;|alert   &emsp;|Alert log level     |
    &emsp;|crit    &emsp;|Critical log level  |
    &emsp;|err     &emsp;|Error log level     |
    &emsp;|warn    &emsp;|Warning log level   |
    &emsp;|notice  &emsp;|Notice log level    |
    &emsp;|info    &emsp;|Info log level      |
    &emsp;|debug   &emsp;|Debug log level     |
- `queue_threshold` (Number) Number of packets to queue inside the kernel before sending them to userspace

    &emsp;|Format   &emsp;|Description                                                                    |
    |-----------|---------------------------------------------------------------------------------|
    &emsp;|0-65535  &emsp;|Number of packets to queue inside the kernel before sending them to userspace  |
- `snapshot_length` (Number) Length of packet payload to include in netlink message

    &emsp;|Format  &emsp;|Description                                             |
    |----------|----------------------------------------------------------|
    &emsp;|0-9000  &emsp;|Length of packet payload to include in netlink message  |


&lt;a id=&#34;nestedatt--outbound_interface&#34;&gt;&lt;/a&gt;
### Nested Schema for `outbound_interface`

Optional:

- `group` (String) Match interface-group

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|txt     &emsp;|Interface-group name to match           |
    &emsp;|!txt    &emsp;|Inverted interface-group name to match  |
- `name` (String) Match interface

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|txt     &emsp;|Interface name                    |
    &emsp;|txt&amp;    &emsp;|Interface name with wildcard      |
    &emsp;|!txt    &emsp;|Inverted interface name to match  |


&lt;a id=&#34;nestedatt--recent&#34;&gt;&lt;/a&gt;
### Nested Schema for `recent`

Optional:

- `count` (Number) Source addresses seen more than N times

    &emsp;|Format  &emsp;|Description                              |
    |----------|-------------------------------------------|
    &emsp;|1-255   &emsp;|Source addresses seen more than N times  |
- `time` (String) Source addresses seen in the last second/minute/hour

    &emsp;|Format  &emsp;|Description                                           |
    |----------|--------------------------------------------------------|
    &emsp;|second  &emsp;|Source addresses seen COUNT times in the last second  |
    &emsp;|minute  &emsp;|Source addresses seen COUNT times in the last minute  |
    &emsp;|hour    &emsp;|Source addresses seen COUNT times in the last hour    |


&lt;a id=&#34;nestedatt--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `source`

Optional:

- `address` (String) IP address, subnet, or range

    &emsp;|Format      &emsp;|Description                                    |
    |--------------|-------------------------------------------------|
    &emsp;|ipv6        &emsp;|IP address to match                            |
    &emsp;|ipv6net     &emsp;|Subnet to match                                |
    &emsp;|ipv6range   &emsp;|IP range to match                              |
    &emsp;|!ipv6       &emsp;|Match everything except the specified address  |
    &emsp;|!ipv6net    &emsp;|Match everything except the specified prefix   |
    &emsp;|!ipv6range  &emsp;|Match everything except the specified range    |
- `address_mask` (String) IP mask

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|ipv6    &emsp;|IP mask to apply  |
- `fqdn` (String) Fully qualified domain name

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|&lt;fqdn&gt;  &emsp;|Fully qualified domain name  |
- `geoip` (Attributes) GeoIP options - Data provided by DB-IP.com (see [below for nested schema](#nestedatt--source--geoip))
- `group` (Attributes) Group (see [below for nested schema](#nestedatt--source--group))
- `mac_address` (String) MAC address

    &emsp;|Format    &emsp;|Description                                        |
    |------------|-----------------------------------------------------|
    &emsp;|macaddr   &emsp;|MAC address to match                               |
    &emsp;|!macaddr  &emsp;|Match everything except the specified MAC address  |
- `port` (String) Port

    &emsp;|Format       &emsp;|Description                                                                                                               |
    |---------------|----------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt          &emsp;|Named port (any name in /etc/services, e.g., http)                                                                        |
    &emsp;|1-65535      &emsp;|Numbered port                                                                                                             |
    &emsp;|&lt;start-end&gt;  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                      |
    &emsp;|             &emsp;|\n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: &#39;telnet,http,123,1001-1005&#39;  |

&lt;a id=&#34;nestedatt--source--geoip&#34;&gt;&lt;/a&gt;
### Nested Schema for `source.geoip`

Optional:

- `country_code` (List of String) GeoIP country code

    &emsp;|Format     &emsp;|Description                  |
    |-------------|-------------------------------|
    &emsp;|&lt;country&gt;  &emsp;|Country code (2 characters)  |
- `inverse_match` (Boolean) Inverse match of country-codes


&lt;a id=&#34;nestedatt--source--group&#34;&gt;&lt;/a&gt;
### Nested Schema for `source.group`

Optional:

- `address_group` (String) Group of addresses
- `domain_group` (String) Group of domains
- `dynamic_address_group` (String) Group of dynamic ipv6 addresses
- `mac_group` (String) Group of MAC addresses
- `network_group` (String) Group of networks
- `port_group` (String) Group of ports



&lt;a id=&#34;nestedatt--synproxy&#34;&gt;&lt;/a&gt;
### Nested Schema for `synproxy`

Optional:

- `tcp` (Attributes) TCP synproxy options (see [below for nested schema](#nestedatt--synproxy--tcp))

&lt;a id=&#34;nestedatt--synproxy--tcp&#34;&gt;&lt;/a&gt;
### Nested Schema for `synproxy.tcp`

Optional:

- `mss` (Number) TCP Maximum segment size

    &emsp;|Format     &emsp;|Description                                    |
    |-------------|-------------------------------------------------|
    &emsp;|501-65535  &emsp;|Maximum segment size for synproxy connections  |
- `window_scale` (Number) TCP window scale for synproxy connections

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|1-14    &emsp;|TCP window scale  |



&lt;a id=&#34;nestedatt--tcp&#34;&gt;&lt;/a&gt;
### Nested Schema for `tcp`

Optional:

- `flags` (Attributes) TCP flags to match (see [below for nested schema](#nestedatt--tcp--flags))
- `mss` (String) Maximum segment size (MSS)

    &emsp;|Format       &emsp;|Description                           |
    |---------------|----------------------------------------|
    &emsp;|1-16384      &emsp;|Maximum segment size                  |
    &emsp;|&lt;min&gt;-&lt;max&gt;  &emsp;|TCP MSS range (use &#39;-&#39; as delimiter)  |

&lt;a id=&#34;nestedatt--tcp--flags&#34;&gt;&lt;/a&gt;
### Nested Schema for `tcp.flags`

Optional:

- `ack` (Boolean) Acknowledge flag
- `cwr` (Boolean) Congestion Window Reduced flag
- `ecn` (Boolean) Explicit Congestion Notification flag
- `fin` (Boolean) Finish flag
- `not` (Attributes) Match flags not set (see [below for nested schema](#nestedatt--tcp--flags--not))
- `psh` (Boolean) Push flag
- `rst` (Boolean) Reset flag
- `syn` (Boolean) Synchronise flag
- `urg` (Boolean) Urgent flag

&lt;a id=&#34;nestedatt--tcp--flags--not&#34;&gt;&lt;/a&gt;
### Nested Schema for `tcp.flags.not`

Optional:

- `ack` (Boolean) Acknowledge flag
- `cwr` (Boolean) Congestion Window Reduced flag
- `ecn` (Boolean) Explicit Congestion Notification flag
- `fin` (Boolean) Finish flag
- `psh` (Boolean) Push flag
- `rst` (Boolean) Reset flag
- `syn` (Boolean) Synchronise flag
- `urg` (Boolean) Urgent flag




&lt;a id=&#34;nestedatt--time&#34;&gt;&lt;/a&gt;
### Nested Schema for `time`

Optional:

- `startdate` (String) Date to start matching rule

    &emsp;|Format  &emsp;|Description                                       |
    |----------|----------------------------------------------------|
    &emsp;|txt     &emsp;|Enter date using following notation - YYYY-MM-DD  |
- `starttime` (String) Time of day to start matching rule

    &emsp;|Format  &emsp;|Description                                         |
    |----------|------------------------------------------------------|
    &emsp;|txt     &emsp;|Enter time using using 24 hour notation - hh:mm:ss  |
- `stopdate` (String) Date to stop matching rule

    &emsp;|Format  &emsp;|Description                                       |
    |----------|----------------------------------------------------|
    &emsp;|txt     &emsp;|Enter date using following notation - YYYY-MM-DD  |
- `stoptime` (String) Time of day to stop matching rule

    &emsp;|Format  &emsp;|Description                                         |
    |----------|------------------------------------------------------|
    &emsp;|txt     &emsp;|Enter time using using 24 hour notation - hh:mm:ss  |
- `weekdays` (String) Comma separated weekdays to match rule on

    &emsp;|Format  &emsp;|Description                                                                    |
    |----------|---------------------------------------------------------------------------------|
    &emsp;|txt     &emsp;|Name of day (Monday, Tuesday, Wednesday, Thursdays, Friday, Saturday, Sunday)  |
    &emsp;|0-6     &emsp;|Day number (0 = Sunday ... 6 = Saturday)                                       |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
