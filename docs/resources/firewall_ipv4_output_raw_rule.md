---
page_title: "vyos_firewall_ipv4_output_raw_rule Resource - vyos"

subcategory: "Firewall"

description: |- 
  Firewall⯯IPv4 firewall⯯IPv4 output firewall⯯IPv4 firewall output raw⯯IPv4 Firewall output raw rule number
---

# vyos_firewall_ipv4_output_raw_rule (Resource)
<center>

Firewall  
⯯  
IPv4 firewall  
⯯  
IPv4 output firewall  
⯯  
IPv4 firewall output raw  
⯯  
**IPv4 Firewall output raw rule number**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `rule_id` (Number) IPv4 Firewall output raw rule number

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
    &emsp;|notrack   &emsp;|Ignore connection tracking                                                     |
- `add_address_to_group` (Attributes) Add ip address to dynamic address-group (see [below for nested schema](#nestedatt--add_address_to_group))
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
- `icmp` (Attributes) ICMP type and code information (see [below for nested schema](#nestedatt--icmp))
- `ipsec` (Attributes) Outbound IPsec packets (see [below for nested schema](#nestedatt--ipsec))
- `limit` (Attributes) Rate limit using a token bucket filter (see [below for nested schema](#nestedatt--limit))
- `log` (Boolean) Log packets hitting this rule
- `log_options` (Attributes) Log options (see [below for nested schema](#nestedatt--log_options))
- `outbound_interface` (Attributes) Match outbound-interface (see [below for nested schema](#nestedatt--outbound_interface))
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
- `tcp` (Attributes) TCP options to match (see [below for nested schema](#nestedatt--tcp))
- `time` (Attributes) Time to match rule (see [below for nested schema](#nestedatt--time))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `ttl` (Attributes) Time to live limit (see [below for nested schema](#nestedatt--ttl))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--add_address_to_group&#34;&gt;&lt;/a&gt;
### Nested Schema for `add_address_to_group`

Optional:

- `destination_address` (Attributes) Add destination ip addresses to dynamic address-group (see [below for nested schema](#nestedatt--add_address_to_group--destination_address))
- `source_address` (Attributes) Add source ip addresses to dynamic address-group (see [below for nested schema](#nestedatt--add_address_to_group--source_address))

&lt;a id=&#34;nestedatt--add_address_to_group--destination_address&#34;&gt;&lt;/a&gt;
### Nested Schema for `add_address_to_group.destination_address`

Optional:

- `address_group` (String) Dynamic address-group
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

- `address_group` (String) Dynamic address-group
- `timeout` (String) Set timeout

    &emsp;|Format     &emsp;|Description               |
    |-------------|----------------------------|
    &emsp;|&lt;number&gt;s  &emsp;|Timeout value in seconds  |
    &emsp;|&lt;number&gt;m  &emsp;|Timeout value in minutes  |
    &emsp;|&lt;number&gt;h  &emsp;|Timeout value in hours    |
    &emsp;|&lt;number&gt;d  &emsp;|Timeout value in days     |



&lt;a id=&#34;nestedatt--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination`

Optional:

- `address` (String) IP address, subnet, or range

    &emsp;|Format      &emsp;|Description                                    |
    |--------------|-------------------------------------------------|
    &emsp;|ipv4        &emsp;|IPv4 address to match                          |
    &emsp;|ipv4net     &emsp;|IPv4 prefix to match                           |
    &emsp;|ipv4range   &emsp;|IPv4 address range to match                    |
    &emsp;|!ipv4       &emsp;|Match everything except the specified address  |
    &emsp;|!ipv4net    &emsp;|Match everything except the specified prefix   |
    &emsp;|!ipv4range  &emsp;|Match everything except the specified range    |
- `address_mask` (String) IP mask

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|ipv4    &emsp;|IPv4 mask to apply  |
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
- `mac_group` (String) Group of MAC addresses
- `network_group` (String) Group of networks
- `port_group` (String) Group of ports



&lt;a id=&#34;nestedatt--fragment&#34;&gt;&lt;/a&gt;
### Nested Schema for `fragment`

Optional:

- `match_frag` (Boolean) Second and further fragments of fragmented packets
- `match_non_frag` (Boolean) Head fragments or unfragmented packets


&lt;a id=&#34;nestedatt--icmp&#34;&gt;&lt;/a&gt;
### Nested Schema for `icmp`

Optional:

- `code` (Number) ICMP code

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|0-255   &emsp;|ICMP code (0-255)  |
- `type` (Number) ICMP type

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|0-255   &emsp;|ICMP type (0-255)  |
- `type_name` (String) ICMP type-name

    &emsp;|Format                   &emsp;|Description                           |
    |---------------------------|----------------------------------------|
    &emsp;|echo-reply               &emsp;|ICMP type 0: echo-reply               |
    &emsp;|destination-unreachable  &emsp;|ICMP type 3: destination-unreachable  |
    &emsp;|source-quench            &emsp;|ICMP type 4: source-quench            |
    &emsp;|redirect                 &emsp;|ICMP type 5: redirect                 |
    &emsp;|echo-request             &emsp;|ICMP type 8: echo-request             |
    &emsp;|router-advertisement     &emsp;|ICMP type 9: router-advertisement     |
    &emsp;|router-solicitation      &emsp;|ICMP type 10: router-solicitation     |
    &emsp;|time-exceeded            &emsp;|ICMP type 11: time-exceeded           |
    &emsp;|parameter-problem        &emsp;|ICMP type 12: parameter-problem       |
    &emsp;|timestamp-request        &emsp;|ICMP type 13: timestamp-request       |
    &emsp;|timestamp-reply          &emsp;|ICMP type 14: timestamp-reply         |
    &emsp;|info-request             &emsp;|ICMP type 15: info-request            |
    &emsp;|info-reply               &emsp;|ICMP type 16: info-reply              |
    &emsp;|address-mask-request     &emsp;|ICMP type 17: address-mask-request    |
    &emsp;|address-mask-reply       &emsp;|ICMP type 18: address-mask-reply      |


&lt;a id=&#34;nestedatt--ipsec&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipsec`

Optional:

- `match_ipsec_out` (Boolean) Outbound traffic to be IPsec encapsulated
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
    &emsp;|ipv4        &emsp;|IPv4 address to match                          |
    &emsp;|ipv4net     &emsp;|IPv4 prefix to match                           |
    &emsp;|ipv4range   &emsp;|IPv4 address range to match                    |
    &emsp;|!ipv4       &emsp;|Match everything except the specified address  |
    &emsp;|!ipv4net    &emsp;|Match everything except the specified prefix   |
    &emsp;|!ipv4range  &emsp;|Match everything except the specified range    |
- `address_mask` (String) IP mask

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|ipv4    &emsp;|IPv4 mask to apply  |
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

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--ttl&#34;&gt;&lt;/a&gt;
### Nested Schema for `ttl`

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
