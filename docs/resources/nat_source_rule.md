---
page_title: "vyos_nat_source_rule Resource - terraform-provider-vyos"
subcategory: "nat"
description: |-
  Network Address Translation (NAT) parameters⯯Source NAT settings⯯Rule number for NAT
---

# vyos_nat_source_rule (Resource)
<center>

Network Address Translation (NAT) parameters  
⯯  
Source NAT settings  
⯯  
**Rule number for NAT**


</center>

## Schema

### Required

- `rule_id` (Number) Rule number for NAT

    &emsp;|Format    &emsp;|Description         |
    |------------|----------------------|
    &emsp;|1-999999  &emsp;|Number of NAT rule  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `destination` (Attributes) NAT destination parameters (see [below for nested schema](#nestedatt--destination))
- `disable` (Boolean) Disable instance
- `exclude` (Boolean) Exclude packets matching this rule from NAT
- `load_balance` (Attributes) Apply NAT load balance (see [below for nested schema](#nestedatt--load_balance))
- `log` (Boolean) Log packets hitting this rule
- `outbound_interface` (Attributes) Match outbound-interface (see [below for nested schema](#nestedatt--outbound_interface))
- `packet_type` (String) Packet type

    &emsp;|Format     &emsp;|Description                                      |
    |-------------|---------------------------------------------------|
    &emsp;|broadcast  &emsp;|Match broadcast packet type                      |
    &emsp;|host       &emsp;|Match host packet type, addressed to local host  |
    &emsp;|multicast  &emsp;|Match multicast packet type                      |
    &emsp;|other      &emsp;|Match packet addressed to another host           |
- `protocol` (String) Protocol to NAT

    &emsp;|Format           &emsp;|Description                                   |
    |-------------------|------------------------------------------------|
    &emsp;|all              &emsp;|All IP protocols                              |
    &emsp;|ip               &emsp;|Internet Protocol, pseudo protocol number     |
    &emsp;|hopopt           &emsp;|IPv6 Hop-by-Hop Option [RFC1883]              |
    &emsp;|icmp             &emsp;|internet control message protocol             |
    &emsp;|igmp             &emsp;|Internet Group Management                     |
    &emsp;|ggp              &emsp;|gateway-gateway protocol                      |
    &emsp;|ipencap          &emsp;|IP encapsulated in IP (officially IP)         |
    &emsp;|st               &emsp;|ST datagram mode                              |
    &emsp;|tcp              &emsp;|transmission control protocol                 |
    &emsp;|egp              &emsp;|exterior gateway protocol                     |
    &emsp;|igp              &emsp;|any private interior gateway (Cisco)          |
    &emsp;|pup              &emsp;|PARC universal packet protocol                |
    &emsp;|udp              &emsp;|user datagram protocol                        |
    &emsp;|tcp_udp          &emsp;|Both TCP and UDP                              |
    &emsp;|hmp              &emsp;|host monitoring protocol                      |
    &emsp;|xns-idp          &emsp;|Xerox NS IDP                                  |
    &emsp;|rdp              &emsp;|&#34;reliable datagram&#34; protocol                  |
    &emsp;|iso-tp4          &emsp;|ISO Transport Protocol class 4 [RFC905]       |
    &emsp;|dccp             &emsp;|Datagram Congestion Control Prot. [RFC4340]   |
    &emsp;|xtp              &emsp;|Xpress Transfer Protocol                      |
    &emsp;|ddp              &emsp;|Datagram Delivery Protocol                    |
    &emsp;|idpr-cmtp        &emsp;|IDPR Control Message Transport                |
    &emsp;|Ipv6             &emsp;|Internet Protocol, version 6                  |
    &emsp;|ipv6-route       &emsp;|Routing Header for IPv6                       |
    &emsp;|ipv6-frag        &emsp;|Fragment Header for IPv6                      |
    &emsp;|idrp             &emsp;|Inter-Domain Routing Protocol                 |
    &emsp;|rsvp             &emsp;|Reservation Protocol                          |
    &emsp;|gre              &emsp;|General Routing Encapsulation                 |
    &emsp;|esp              &emsp;|Encap Security Payload [RFC2406]              |
    &emsp;|ah               &emsp;|Authentication Header [RFC2402]               |
    &emsp;|skip             &emsp;|SKIP                                          |
    &emsp;|ipv6-icmp        &emsp;|ICMP for IPv6                                 |
    &emsp;|ipv6-nonxt       &emsp;|No Next Header for IPv6                       |
    &emsp;|ipv6-opts        &emsp;|Destination Options for IPv6                  |
    &emsp;|rspf             &emsp;|Radio Shortest Path First (officially CPHB)   |
    &emsp;|vmtp             &emsp;|Versatile Message Transport                   |
    &emsp;|eigrp            &emsp;|Enhanced Interior Routing Protocol (Cisco)    |
    &emsp;|ospf             &emsp;|Open Shortest Path First IGP                  |
    &emsp;|ax.25            &emsp;|AX.25 frames                                  |
    &emsp;|ipip             &emsp;|IP-within-IP Encapsulation Protocol           |
    &emsp;|etherip          &emsp;|Ethernet-within-IP Encapsulation [RFC3378]    |
    &emsp;|encap            &emsp;|Yet Another IP encapsulation [RFC1241]        |
    &emsp;|99               &emsp;|Any private encryption scheme                 |
    &emsp;|pim              &emsp;|Protocol Independent Multicast                |
    &emsp;|ipcomp           &emsp;|IP Payload Compression Protocol               |
    &emsp;|vrrp             &emsp;|Virtual Router Redundancy Protocol [RFC5798]  |
    &emsp;|l2tp             &emsp;|Layer Two Tunneling Protocol [RFC2661]        |
    &emsp;|isis             &emsp;|IS-IS over IPv4                               |
    &emsp;|sctp             &emsp;|Stream Control Transmission Protocol          |
    &emsp;|fc               &emsp;|Fibre Channel                                 |
    &emsp;|mobility-header  &emsp;|Mobility Support for IPv6 [RFC3775]           |
    &emsp;|udplite          &emsp;|UDP-Lite [RFC3828]                            |
    &emsp;|mpls-in-ip       &emsp;|MPLS-in-IP [RFC4023]                          |
    &emsp;|manet            &emsp;|MANET Protocols [RFC5498]                     |
    &emsp;|hip              &emsp;|Host Identity Protocol                        |
    &emsp;|shim6            &emsp;|Shim6 Protocol                                |
    &emsp;|wesp             &emsp;|Wrapped Encapsulating Security Payload        |
    &emsp;|rohc             &emsp;|Robust Header Compression                     |
    &emsp;|0-255            &emsp;|IP protocol number                            |
- `source` (Attributes) NAT source parameters (see [below for nested schema](#nestedatt--source))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `translation` (Attributes) Outside NAT IP (source NAT only) (see [below for nested schema](#nestedatt--translation))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

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
- `group` (Attributes) Group (see [below for nested schema](#nestedatt--destination--group))
- `port` (String) Port number

    &emsp;|Format     &emsp;|Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt        &emsp;|Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    &emsp;|1-65535    &emsp;|Numeric IP port                                                                                                                                                          |
    &emsp;|start-end  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    &emsp;|           &emsp;|\n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using &#39;!&#39;.\nFor example: &#39;!22,telnet,http,123,1001-1005&#39;  |

&lt;a id=&#34;nestedatt--destination--group&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination.group`

Optional:

- `address_group` (String) Group of addresses
- `domain_group` (String) Group of domains
- `mac_group` (String) Group of MAC addresses
- `network_group` (String) Group of networks
- `port_group` (String) Group of ports



&lt;a id=&#34;nestedatt--load_balance&#34;&gt;&lt;/a&gt;
### Nested Schema for `load_balance`

Optional:

- `hash` (List of String) Define the parameters of the packet header to apply the hashing

    &emsp;|Format               &emsp;|Description                                               |
    |-----------------------|------------------------------------------------------------|
    &emsp;|source-address       &emsp;|Use source IP address for hashing                         |
    &emsp;|destination-address  &emsp;|Use destination IP address for hashing                    |
    &emsp;|source-port          &emsp;|Use source port for hashing                               |
    &emsp;|destination-port     &emsp;|Use destination port for hashing                          |
    &emsp;|random               &emsp;|Do not use information from ip header. Use random value.  |


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
- `group` (Attributes) Group (see [below for nested schema](#nestedatt--source--group))
- `port` (String) Port number

    &emsp;|Format     &emsp;|Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt        &emsp;|Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    &emsp;|1-65535    &emsp;|Numeric IP port                                                                                                                                                          |
    &emsp;|start-end  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    &emsp;|           &emsp;|\n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using &#39;!&#39;.\nFor example: &#39;!22,telnet,http,123,1001-1005&#39;  |

&lt;a id=&#34;nestedatt--source--group&#34;&gt;&lt;/a&gt;
### Nested Schema for `source.group`

Optional:

- `address_group` (String) Group of addresses
- `domain_group` (String) Group of domains
- `mac_group` (String) Group of MAC addresses
- `network_group` (String) Group of networks
- `port_group` (String) Group of ports



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--translation&#34;&gt;&lt;/a&gt;
### Nested Schema for `translation`

Optional:

- `address` (String) IP address, subnet, or range

    &emsp;|Format      &emsp;|Description                                       |
    |--------------|----------------------------------------------------|
    &emsp;|ipv4        &emsp;|IPv4 address to match                             |
    &emsp;|ipv4net     &emsp;|IPv4 prefix to match                              |
    &emsp;|ipv4range   &emsp;|IPv4 address range to match                       |
    &emsp;|masquerade  &emsp;|NAT to the primary address of outbound-interface  |
- `options` (Attributes) Translation options (see [below for nested schema](#nestedatt--translation--options))
- `port` (String) Port number

    &emsp;|Format   &emsp;|Description                            |
    |-----------|-----------------------------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port                        |
    &emsp;|range    &emsp;|Numbered port range (e.g., 1001-1005)  |

&lt;a id=&#34;nestedatt--translation--options&#34;&gt;&lt;/a&gt;
### Nested Schema for `translation.options`

Optional:

- `address_mapping` (String) Address mapping options

    &emsp;|Format      &emsp;|Description                                                                |
    |--------------|-----------------------------------------------------------------------------|
    &emsp;|persistent  &emsp;|Gives a client the same source or destination-address for each connection  |
    &emsp;|random      &emsp;|Random source or destination address allocation for each connection        |
- `port_mapping` (String) Port mapping options

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|random        &emsp;|Randomize source port mapping    |
    &emsp;|fully-random  &emsp;|Full port randomization          |
    &emsp;|none          &emsp;|Do not apply port randomization  |  &emsp;|
