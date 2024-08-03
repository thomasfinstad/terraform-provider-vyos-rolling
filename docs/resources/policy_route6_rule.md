---
page_title: "vyos_policy_route6_rule Resource - vyos"

subcategory: "Policy"

description: |- 
  policy⯯Policy route rule set name for IPv6⯯Policy rule number
---

# vyos_policy_route6_rule (Resource)
<center>

*policy*  
⯯  
Policy route rule set name for IPv6  
⯯  
**Policy rule number**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `action` (String) Rule action

    &emsp;|Format  &emsp;|Description                                                                    |
    |----------|---------------------------------------------------------------------------------|
    &emsp;|accept  &emsp;|Accept matching entries                                                        |
    &emsp;|reject  &emsp;|Reject matching entries                                                        |
    &emsp;|return  &emsp;|Return from the current chain and continue at the next rule of the last chain  |
    &emsp;|drop    &emsp;|Drop matching entries                                                          |
- `connection_mark` (List of Number) Connection mark

    &emsp;|Format        &emsp;|Description               |
    |----------------|----------------------------|
    &emsp;|0-2147483647  &emsp;|Connection-mark to match  |
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
- `ipsec` (Attributes) IPsec encapsulated packets (see [below for nested schema](#nestedatt--ipsec))
- `limit` (Attributes) Rate limit using a token bucket filter (see [below for nested schema](#nestedatt--limit))
- `log` (Boolean) Log packets hitting this rule
- `mark` (String) Firewall mark

    &emsp;|Format         &emsp;|Description                            |
    |-----------------|-----------------------------------------|
    &emsp;|0-2147483647   &emsp;|Firewall mark to match                 |
    &emsp;|!0-2147483647  &emsp;|Inverted Firewall mark to match        |
    &emsp;|&lt;start-end&gt;    &emsp;|Firewall mark range to match           |
    &emsp;|!&lt;start-end&gt;   &emsp;|Firewall mark inverted range to match  |
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
    &emsp;|!&lt;protocol&gt;  &emsp;|IP protocol number  |
- `recent` (Attributes) Parameters for matching recently seen sources (see [below for nested schema](#nestedatt--recent))
- `set` (Attributes) Packet modifications (see [below for nested schema](#nestedatt--set))
- `source` (Attributes) Source parameters (see [below for nested schema](#nestedatt--source))
- `state` (List of String) Session state

    &emsp;|Format       &emsp;|Description        |
    |---------------|---------------------|
    &emsp;|established  &emsp;|Established state  |
    &emsp;|invalid      &emsp;|Invalid state      |
    &emsp;|new          &emsp;|New state          |
    &emsp;|related      &emsp;|Related state      |
- `tcp` (Attributes) TCP options to match (see [below for nested schema](#nestedatt--tcp))
- `time` (Attributes) Time to match rule (see [below for nested schema](#nestedatt--time))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `route6` (String) Policy route rule set name for IPv6
- `rule` (Number) Policy rule number

    &emsp;|Format    &emsp;|Description            |
    |------------|-------------------------|
    &emsp;|1-999999  &emsp;|Number of policy rule  |


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
- `group` (Attributes) Group (see [below for nested schema](#nestedatt--destination--group))
- `port` (String) Port

    &emsp;|Format       &emsp;|Description                                                                                                               |
    |---------------|----------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt          &emsp;|Named port (any name in /etc/services, e.g., http)                                                                        |
    &emsp;|1-65535      &emsp;|Numbered port                                                                                                             |
    &emsp;|&lt;start-end&gt;  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                      |
    &emsp;|             &emsp;|\n\n  Multiple destination ports can be specified as a comma-separated list.\n  For example: &#39;telnet,http,123,1001-1005&#39;  |

&lt;a id=&#34;nestedatt--destination--group&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination.group`

Optional:

- `address_group` (String) Group of addresses
- `domain_group` (String) Group of domains
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

- `type` (String) ICMP type-name

    &emsp;|Format                      &emsp;|Description          |
    |------------------------------|-----------------------|
    &emsp;|any                         &emsp;|Any ICMP type/code   |
    &emsp;|echo-reply                  &emsp;|ICMP type/code name  |
    &emsp;|pong                        &emsp;|ICMP type/code name  |
    &emsp;|destination-unreachable     &emsp;|ICMP type/code name  |
    &emsp;|network-unreachable         &emsp;|ICMP type/code name  |
    &emsp;|host-unreachable            &emsp;|ICMP type/code name  |
    &emsp;|protocol-unreachable        &emsp;|ICMP type/code name  |
    &emsp;|port-unreachable            &emsp;|ICMP type/code name  |
    &emsp;|fragmentation-needed        &emsp;|ICMP type/code name  |
    &emsp;|source-route-failed         &emsp;|ICMP type/code name  |
    &emsp;|network-unknown             &emsp;|ICMP type/code name  |
    &emsp;|host-unknown                &emsp;|ICMP type/code name  |
    &emsp;|network-prohibited          &emsp;|ICMP type/code name  |
    &emsp;|host-prohibited             &emsp;|ICMP type/code name  |
    &emsp;|TOS-network-unreachable     &emsp;|ICMP type/code name  |
    &emsp;|TOS-host-unreachable        &emsp;|ICMP type/code name  |
    &emsp;|communication-prohibited    &emsp;|ICMP type/code name  |
    &emsp;|host-precedence-violation   &emsp;|ICMP type/code name  |
    &emsp;|precedence-cutoff           &emsp;|ICMP type/code name  |
    &emsp;|source-quench               &emsp;|ICMP type/code name  |
    &emsp;|redirect                    &emsp;|ICMP type/code name  |
    &emsp;|network-redirect            &emsp;|ICMP type/code name  |
    &emsp;|host-redirect               &emsp;|ICMP type/code name  |
    &emsp;|TOS-network-redirect        &emsp;|ICMP type/code name  |
    &emsp;|TOS host-redirect           &emsp;|ICMP type/code name  |
    &emsp;|echo-request                &emsp;|ICMP type/code name  |
    &emsp;|ping                        &emsp;|ICMP type/code name  |
    &emsp;|router-advertisement        &emsp;|ICMP type/code name  |
    &emsp;|router-solicitation         &emsp;|ICMP type/code name  |
    &emsp;|time-exceeded               &emsp;|ICMP type/code name  |
    &emsp;|ttl-exceeded                &emsp;|ICMP type/code name  |
    &emsp;|ttl-zero-during-transit     &emsp;|ICMP type/code name  |
    &emsp;|ttl-zero-during-reassembly  &emsp;|ICMP type/code name  |
    &emsp;|parameter-problem           &emsp;|ICMP type/code name  |
    &emsp;|ip-header-bad               &emsp;|ICMP type/code name  |
    &emsp;|required-option-missing     &emsp;|ICMP type/code name  |
    &emsp;|timestamp-request           &emsp;|ICMP type/code name  |
    &emsp;|timestamp-reply             &emsp;|ICMP type/code name  |
    &emsp;|address-mask-request        &emsp;|ICMP type/code name  |
    &emsp;|address-mask-reply          &emsp;|ICMP type/code name  |
    &emsp;|packet-too-big              &emsp;|ICMP type/code name  |


&lt;a id=&#34;nestedatt--ipsec&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipsec`

Optional:

- `match_ipsec_in` (Boolean) Inbound traffic that was IPsec encapsulated
- `match_ipsec_out` (Boolean) Outbound traffic to be IPsec encapsulated
- `match_none_in` (Boolean) Inbound traffic that was not IPsec encapsulated
- `match_none_out` (Boolean) Outbound traffic that will not be IPsec encapsulated


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


&lt;a id=&#34;nestedatt--recent&#34;&gt;&lt;/a&gt;
### Nested Schema for `recent`

Optional:

- `count` (Number) Source addresses seen more than N times

    &emsp;|Format  &emsp;|Description                              |
    |----------|-------------------------------------------|
    &emsp;|1-255   &emsp;|Source addresses seen more than N times  |
- `time` (Number) Source addresses seen in the last N seconds

    &emsp;|Format        &emsp;|Description                                  |
    |----------------|-----------------------------------------------|
    &emsp;|0-4294967295  &emsp;|Source addresses seen in the last N seconds  |


&lt;a id=&#34;nestedatt--set&#34;&gt;&lt;/a&gt;
### Nested Schema for `set`

Optional:

- `connection_mark` (Number) Connection marking

    &emsp;|Format        &emsp;|Description         |
    |----------------|----------------------|
    &emsp;|0-2147483647  &emsp;|Connection marking  |
- `dscp` (Number) Packet Differentiated Services Codepoint (DSCP)

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|0-63    &emsp;|DSCP number  |
- `mark` (Number) Packet marking

    &emsp;|Format        &emsp;|Description     |
    |----------------|------------------|
    &emsp;|1-2147483647  &emsp;|Packet marking  |
- `table` (String) Routing table to forward packet with

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|1-200   &emsp;|Table number  |
    &emsp;|main    &emsp;|Main table    |
- `tcp_mss` (Number) TCP Maximum Segment Size

    &emsp;|Format    &emsp;|Description                   |
    |------------|--------------------------------|
    &emsp;|500-1460  &emsp;|Explicitly set TCP MSS value  |
- `vrf` (String) VRF to forward packet with

    &emsp;|Format   &emsp;|Description                      |
    |-----------|-----------------------------------|
    &emsp;|txt      &emsp;|VRF instance name                |
    &emsp;|default  &emsp;|Forward into default global VRF  |


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

&lt;a id=&#34;nestedatt--source--group&#34;&gt;&lt;/a&gt;
### Nested Schema for `source.group`

Optional:

- `address_group` (String) Group of addresses
- `domain_group` (String) Group of domains
- `mac_group` (String) Group of MAC addresses
- `network_group` (String) Group of networks
- `port_group` (String) Group of ports



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

- `monthdays` (String) Monthdays to match rule on
- `startdate` (String) Date to start matching rule
- `starttime` (String) Time of day to start matching rule
- `stopdate` (String) Date to stop matching rule
- `stoptime` (String) Time of day to stop matching rule
- `utc` (Boolean) Interpret times for startdate, stopdate, starttime and stoptime to be UTC
- `weekdays` (String) Weekdays to match rule on


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
