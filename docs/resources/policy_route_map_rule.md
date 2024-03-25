---
page_title: "vyos_policy_route_map_rule Resource - vyos"
subcategory: "policy"
description: |- 
  Routing policy⯯IP route-map⯯Rule for this route-map
---

# vyos_policy_route_map_rule (Resource)
<center>

Routing policy  
⯯  
IP route-map  
⯯  
**Rule for this route-map**


</center>

## Schema

### Required

- `route_map_id` (String) IP route-map

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `rule_id` (Number) Rule for this route-map

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|1-65535  &emsp;|Route-map rule number  |

### Optional

- `action` (String) Action to take on entries matching this rule

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|permit  &emsp;|Permit matching entries  |
    &emsp;|deny    &emsp;|Deny matching entries    |
- `call` (String) Call another route-map on match

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `continue` (Number) Jump to a different rule in this route-map on a match

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|1-65535  &emsp;|Rule number  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `match` (Attributes) Route parameters to match (see [below for nested schema](#nestedatt--match))
- `on_match` (Attributes) Exit policy on matches (see [below for nested schema](#nestedatt--on_match))
- `set` (Attributes) Route parameters (see [below for nested schema](#nestedatt--set))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--match&#34;&gt;&lt;/a&gt;
### Nested Schema for `match`

Optional:

- `as_path` (String) BGP as-path-list to match
- `community` (Attributes) BGP community-list to match (see [below for nested schema](#nestedatt--match--community))
- `evpn` (Attributes) Ethernet Virtual Private Network (see [below for nested schema](#nestedatt--match--evpn))
- `extcommunity` (String) BGP extended community to match
- `interface` (String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `ip` (Attributes) IP prefix parameters to match (see [below for nested schema](#nestedatt--match--ip))
- `ipv6` (Attributes) IPv6 prefix parameters to match (see [below for nested schema](#nestedatt--match--ipv6))
- `large_community` (Attributes) Match BGP large communities (see [below for nested schema](#nestedatt--match--large_community))
- `local_preference` (Number) Local Preference

    &emsp;|Format        &emsp;|Description       |
    |----------------|--------------------|
    &emsp;|0-4294967295  &emsp;|Local Preference  |
- `metric` (Number) Metric of route to match

    &emsp;|Format   &emsp;|Description   |
    |-----------|----------------|
    &emsp;|1-65535  &emsp;|Route metric  |
- `origin` (String) BGP origin code to match

    &emsp;|Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    &emsp;|egp         &emsp;|Exterior gateway protocol origin  |
    &emsp;|igp         &emsp;|Interior gateway protocol origin  |
    &emsp;|incomplete  &emsp;|Incomplete origin                 |
- `peer` (String) Peer address to match

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|ipv4    &emsp;|Peer IP address    |
    &emsp;|ipv6    &emsp;|Peer IPv6 address  |
- `protocol` (String) Match protocol via which the route was learnt

    &emsp;|Format     &emsp;|Description                                                  |
    |-------------|---------------------------------------------------------------|
    &emsp;|babel      &emsp;|Babel routing protocol (Babel)                               |
    &emsp;|bgp        &emsp;|Border Gateway Protocol (BGP)                                |
    &emsp;|connected  &emsp;|Connected routes (directly attached subnet or host)          |
    &emsp;|isis       &emsp;|Intermediate System to Intermediate System (IS-IS)           |
    &emsp;|kernel     &emsp;|Kernel routes                                                |
    &emsp;|ospf       &emsp;|Open Shortest Path First (OSPFv2)                            |
    &emsp;|ospfv3     &emsp;|Open Shortest Path First (IPv6) (OSPFv3)                     |
    &emsp;|rip        &emsp;|Routing Information Protocol (RIP)                           |
    &emsp;|ripng      &emsp;|Routing Information Protocol next-generation (IPv6) (RIPng)  |
    &emsp;|static     &emsp;|Statically configured routes                                 |
    &emsp;|table      &emsp;|Non-main Kernel Routing Table                                |
    &emsp;|vnc        &emsp;|Virtual Network Control (VNC)                                |
- `rpki` (String) Match RPKI validation result

    &emsp;|Format    &emsp;|Description             |
    |------------|--------------------------|
    &emsp;|invalid   &emsp;|Match invalid entries   |
    &emsp;|notfound  &emsp;|Match notfound entries  |
    &emsp;|valid     &emsp;|Match valid entries     |
- `tag` (Number) Route tag value

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|1-65535  &emsp;|Route tag    |

&lt;a id=&#34;nestedatt--match--community&#34;&gt;&lt;/a&gt;
### Nested Schema for `match.community`

Optional:

- `community_list` (String) BGP community-list to match
- `exact_match` (Boolean) Community-list to exactly match


&lt;a id=&#34;nestedatt--match--evpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `match.evpn`

Optional:

- `default_route` (Boolean) Default EVPN type-5 route
- `rd` (String) Route Distinguisher

    &emsp;|Format                   &emsp;|Description                                   |
    |---------------------------|------------------------------------------------|
    &emsp;|ASN:NN_OR_IP-ADDRESS:NN  &emsp;|Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
- `route_type` (String) Match route-type

    &emsp;|Format     &emsp;|Description   |
    |-------------|----------------|
    &emsp;|macip      &emsp;|mac-ip route  |
    &emsp;|multicast  &emsp;|IMET route    |
    &emsp;|prefix     &emsp;|Prefix route  |
- `vni` (Number) Virtual Network Identifier

    &emsp;|Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    &emsp;|0-16777214  &emsp;|VXLAN virtual network identifier  |


&lt;a id=&#34;nestedatt--match--ip&#34;&gt;&lt;/a&gt;
### Nested Schema for `match.ip`

Optional:

- `address` (Attributes) IP address of route to match (see [below for nested schema](#nestedatt--match--ip--address))
- `nexthop` (Attributes) IP next-hop of route to match (see [below for nested schema](#nestedatt--match--ip--nexthop))
- `route_source` (Attributes) Match advertising source address of route (see [below for nested schema](#nestedatt--match--ip--route_source))

&lt;a id=&#34;nestedatt--match--ip--address&#34;&gt;&lt;/a&gt;
### Nested Schema for `match.ip.address`

Optional:

- `access_list` (Number) IP access-list to match

    &emsp;|Format     &emsp;|Description                               |
    |-------------|--------------------------------------------|
    &emsp;|1-99       &emsp;|IP standard access list                   |
    &emsp;|100-199    &emsp;|IP extended access list                   |
    &emsp;|1300-1999  &emsp;|IP standard access list (expanded range)  |
    &emsp;|2000-2699  &emsp;|IP extended access list (expanded range)  |
- `prefix_len` (Number) IP prefix-length to match (can be used for kernel routes only)

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|0-32    &emsp;|Prefix length  |
- `prefix_list` (String) IP prefix-list to match


&lt;a id=&#34;nestedatt--match--ip--nexthop&#34;&gt;&lt;/a&gt;
### Nested Schema for `match.ip.nexthop`

Optional:

- `access_list` (Number) IP access-list to match

    &emsp;|Format     &emsp;|Description                               |
    |-------------|--------------------------------------------|
    &emsp;|1-99       &emsp;|IP standard access list                   |
    &emsp;|100-199    &emsp;|IP extended access list                   |
    &emsp;|1300-1999  &emsp;|IP standard access list (expanded range)  |
    &emsp;|2000-2699  &emsp;|IP extended access list (expanded range)  |
- `address` (String) IP address to match

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|ipv4    &emsp;|Nexthop IP address  |
- `prefix_len` (Number) IP prefix-length to match

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|0-32    &emsp;|Prefix length  |
- `prefix_list` (String) IP prefix-list to match
- `type` (String) Match type

    &emsp;|Format     &emsp;|Description  |
    |-------------|---------------|
    &emsp;|blackhole  &emsp;|Blackhole    |


&lt;a id=&#34;nestedatt--match--ip--route_source&#34;&gt;&lt;/a&gt;
### Nested Schema for `match.ip.route_source`

Optional:

- `access_list` (Number) IP access-list to match

    &emsp;|Format     &emsp;|Description                               |
    |-------------|--------------------------------------------|
    &emsp;|1-99       &emsp;|IP standard access list                   |
    &emsp;|100-199    &emsp;|IP extended access list                   |
    &emsp;|1300-1999  &emsp;|IP standard access list (expanded range)  |
    &emsp;|2000-2699  &emsp;|IP extended access list (expanded range)  |
- `prefix_list` (String) IP prefix-list to match



&lt;a id=&#34;nestedatt--match--ipv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `match.ipv6`

Optional:

- `address` (Attributes) IPv6 address of route to match (see [below for nested schema](#nestedatt--match--ipv6--address))
- `nexthop` (Attributes) IPv6 next-hop of route to match (see [below for nested schema](#nestedatt--match--ipv6--nexthop))

&lt;a id=&#34;nestedatt--match--ipv6--address&#34;&gt;&lt;/a&gt;
### Nested Schema for `match.ipv6.address`

Optional:

- `access_list` (String) IPv6 access-list to match

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|txt     &emsp;|IPV6 access list name  |
- `prefix_len` (Number) IPv6 prefix-length to match (can be used for kernel routes only)

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|0-128   &emsp;|Prefix length  |
- `prefix_list` (String) IPv6 prefix-list to match


&lt;a id=&#34;nestedatt--match--ipv6--nexthop&#34;&gt;&lt;/a&gt;
### Nested Schema for `match.ipv6.nexthop`

Optional:

- `access_list` (String) IPv6 access-list to match

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|txt     &emsp;|IPV6 access list name  |
- `address` (String) IPv6 address of next-hop

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|ipv6    &emsp;|Nexthop IPv6 address  |
- `prefix_list` (String) IPv6 prefix-list to match
- `type` (String) Match type

    &emsp;|Format     &emsp;|Description  |
    |-------------|---------------|
    &emsp;|blackhole  &emsp;|Blackhole    |



&lt;a id=&#34;nestedatt--match--large_community&#34;&gt;&lt;/a&gt;
### Nested Schema for `match.large_community`

Optional:

- `large_community_list` (String) BGP large-community-list to match



&lt;a id=&#34;nestedatt--on_match&#34;&gt;&lt;/a&gt;
### Nested Schema for `on_match`

Optional:

- `goto` (Number) Rule number to goto on match

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|1-65535  &emsp;|Rule number  |
- `next` (Boolean) Next sequence number to goto on match


&lt;a id=&#34;nestedatt--set&#34;&gt;&lt;/a&gt;
### Nested Schema for `set`

Optional:

- `aggregator` (Attributes) BGP aggregator attribute (see [below for nested schema](#nestedatt--set--aggregator))
- `as_path` (Attributes) Transform BGP AS_PATH attribute (see [below for nested schema](#nestedatt--set--as_path))
- `atomic_aggregate` (Boolean) BGP atomic aggregate attribute
- `community` (Attributes) BGP community attribute (see [below for nested schema](#nestedatt--set--community))
- `distance` (Number) Locally significant administrative distance

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|0-255   &emsp;|Distance value  |
- `evpn` (Attributes) Ethernet Virtual Private Network (see [below for nested schema](#nestedatt--set--evpn))
- `extcommunity` (Attributes) BGP extended community attribute (see [below for nested schema](#nestedatt--set--extcommunity))
- `ip_next_hop` (String) Nexthop IP address

    &emsp;|Format        &emsp;|Description                                             |
    |----------------|----------------------------------------------------------|
    &emsp;|ipv4          &emsp;|IP address                                              |
    &emsp;|unchanged     &emsp;|Set the BGP nexthop address as unchanged                |
    &emsp;|peer-address  &emsp;|Set the BGP nexthop address to the address of the peer  |
- `ipv6_next_hop` (Attributes) Nexthop IPv6 address (see [below for nested schema](#nestedatt--set--ipv6_next_hop))
- `l3vpn_nexthop` (Attributes) Next hop Information (see [below for nested schema](#nestedatt--set--l3vpn_nexthop))
- `large_community` (Attributes) BGP large community attribute (see [below for nested schema](#nestedatt--set--large_community))
- `local_preference` (Number) BGP local preference attribute

    &emsp;|Format        &emsp;|Description             |
    |----------------|--------------------------|
    &emsp;|0-4294967295  &emsp;|Local preference value  |
- `metric` (String) Destination routing protocol metric

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|&lt;+/-metric&gt;   &emsp;|Add or subtract metric           |
    &emsp;|0-4294967295  &emsp;|Metric value                     |
    &emsp;|&lt;+/-rtt&gt;      &emsp;|Add or subtract round trip time  |
    &emsp;|&lt;rtt&gt;         &emsp;|Round trip time                  |
- `metric_type` (String) Open Shortest Path First (OSPF) external metric-type

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|type-1  &emsp;|OSPF external type 1 metric  |
    &emsp;|type-2  &emsp;|OSPF external type 2 metric  |
- `origin` (String) Border Gateway Protocl (BGP) origin code

    &emsp;|Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    &emsp;|igp         &emsp;|Interior gateway protocol origin  |
    &emsp;|egp         &emsp;|Exterior gateway protocol origin  |
    &emsp;|incomplete  &emsp;|Incomplete origin                 |
- `originator_id` (String) BGP originator ID attribute

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|ipv4    &emsp;|Orignator IP address  |
- `src` (String) Source address for route

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|ipv4    &emsp;|IPv4 address  |
    &emsp;|ipv6    &emsp;|IPv6 address  |
- `table` (Number) Set prefixes to table

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1-200   &emsp;|Table value  |
- `tag` (Number) Route tag value

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|1-65535  &emsp;|Route tag    |
- `weight` (Number) BGP weight attribute

    &emsp;|Format        &emsp;|Description  |
    |----------------|---------------|
    &emsp;|0-4294967295  &emsp;|BGP weight   |

&lt;a id=&#34;nestedatt--set--aggregator&#34;&gt;&lt;/a&gt;
### Nested Schema for `set.aggregator`

Optional:

- `as` (Number) AS number of an aggregation

    &emsp;|Format        &emsp;|Description  |
    |----------------|---------------|
    &emsp;|1-4294967295  &emsp;|Rule number  |
- `ip` (String) IP address of an aggregation

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|ipv4    &emsp;|IP address   |


&lt;a id=&#34;nestedatt--set--as_path&#34;&gt;&lt;/a&gt;
### Nested Schema for `set.as_path`

Optional:

- `exclude` (String) Remove/exclude from the as-path attribute

    &emsp;|Format        &emsp;|Description                              |
    |----------------|-------------------------------------------|
    &emsp;|1-4294967295  &emsp;|AS number                                |
    &emsp;|all           &emsp;|Exclude all AS numbers from the as-path  |
- `prepend` (Number) Prepend to the as-path

    &emsp;|Format        &emsp;|Description  |
    |----------------|---------------|
    &emsp;|1-4294967295  &emsp;|AS number    |
- `prepend_last_as` (Number) Use the last AS-number in the as-path

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|1-10    &emsp;|Number of times to insert  |


&lt;a id=&#34;nestedatt--set--community&#34;&gt;&lt;/a&gt;
### Nested Schema for `set.community`

Optional:

- `add` (List of String) Add communities to a prefix

    &emsp;|Format                      &emsp;|Description                                                         |
    |------------------------------|----------------------------------------------------------------------|
    &emsp;|&lt;AS:VAL&gt;                    &emsp;|Community number in &lt;0-65535:0-65535&gt; format                        |
    &emsp;|local-as                    &emsp;|Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03         |
    &emsp;|no-advertise                &emsp;|Well-known communities value NO_ADVERTISE 0xFFFFFF02                |
    &emsp;|no-export                   &emsp;|Well-known communities value NO_EXPORT 0xFFFFFF01                   |
    &emsp;|internet                    &emsp;|Well-known communities value 0                                      |
    &emsp;|graceful-shutdown           &emsp;|Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000           |
    &emsp;|accept-own                  &emsp;|Well-known communities value ACCEPT_OWN 0xFFFF0001                  |
    &emsp;|route-filter-translated-v4  &emsp;|Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002  |
    &emsp;|route-filter-v4             &emsp;|Well-known communities value ROUTE_FILTER_v4 0xFFFF0003             |
    &emsp;|route-filter-translated-v6  &emsp;|Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004  |
    &emsp;|route-filter-v6             &emsp;|Well-known communities value ROUTE_FILTER_v6 0xFFFF0005             |
    &emsp;|llgr-stale                  &emsp;|Well-known communities value LLGR_STALE 0xFFFF0006                  |
    &emsp;|no-llgr                     &emsp;|Well-known communities value NO_LLGR 0xFFFF0007                     |
    &emsp;|accept-own-nexthop          &emsp;|Well-known communities value accept-own-nexthop 0xFFFF0008          |
    &emsp;|blackhole                   &emsp;|Well-known communities value BLACKHOLE 0xFFFF029A                   |
    &emsp;|no-peer                     &emsp;|Well-known communities value NOPEER 0xFFFFFF04                      |
- `delete` (String) Remove communities defined in a list from a prefix

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Community-list  |
- `none` (Boolean) Completely remove communities attribute from a prefix
- `replace` (List of String) Set communities for a prefix

    &emsp;|Format                      &emsp;|Description                                                         |
    |------------------------------|----------------------------------------------------------------------|
    &emsp;|&lt;AS:VAL&gt;                    &emsp;|Community number in &lt;0-65535:0-65535&gt; format                        |
    &emsp;|local-as                    &emsp;|Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03         |
    &emsp;|no-advertise                &emsp;|Well-known communities value NO_ADVERTISE 0xFFFFFF02                |
    &emsp;|no-export                   &emsp;|Well-known communities value NO_EXPORT 0xFFFFFF01                   |
    &emsp;|internet                    &emsp;|Well-known communities value 0                                      |
    &emsp;|graceful-shutdown           &emsp;|Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000           |
    &emsp;|accept-own                  &emsp;|Well-known communities value ACCEPT_OWN 0xFFFF0001                  |
    &emsp;|route-filter-translated-v4  &emsp;|Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002  |
    &emsp;|route-filter-v4             &emsp;|Well-known communities value ROUTE_FILTER_v4 0xFFFF0003             |
    &emsp;|route-filter-translated-v6  &emsp;|Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004  |
    &emsp;|route-filter-v6             &emsp;|Well-known communities value ROUTE_FILTER_v6 0xFFFF0005             |
    &emsp;|llgr-stale                  &emsp;|Well-known communities value LLGR_STALE 0xFFFF0006                  |
    &emsp;|no-llgr                     &emsp;|Well-known communities value NO_LLGR 0xFFFF0007                     |
    &emsp;|accept-own-nexthop          &emsp;|Well-known communities value accept-own-nexthop 0xFFFF0008          |
    &emsp;|blackhole                   &emsp;|Well-known communities value BLACKHOLE 0xFFFF029A                   |
    &emsp;|no-peer                     &emsp;|Well-known communities value NOPEER 0xFFFFFF04                      |


&lt;a id=&#34;nestedatt--set--evpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `set.evpn`

Optional:

- `gateway` (Attributes) Set gateway IP for prefix advertisement route (see [below for nested schema](#nestedatt--set--evpn--gateway))

&lt;a id=&#34;nestedatt--set--evpn--gateway&#34;&gt;&lt;/a&gt;
### Nested Schema for `set.evpn.gateway`

Optional:

- `ipv4` (String) Set gateway IPv4 address

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|ipv4    &emsp;|Gateway IPv4 address  |
- `ipv6` (String) Set gateway IPv6 address

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|ipv6    &emsp;|Gateway IPv6 address  |



&lt;a id=&#34;nestedatt--set--extcommunity&#34;&gt;&lt;/a&gt;
### Nested Schema for `set.extcommunity`

Optional:

- `bandwidth` (String) Bandwidth value in Mbps

    &emsp;|Format          &emsp;|Description                                                                  |
    |------------------|-------------------------------------------------------------------------------|
    &emsp;|1-25600         &emsp;|Bandwidth value in Mbps                                                      |
    &emsp;|cumulative      &emsp;|Cumulative bandwidth of all multipaths (outbound-only)                       |
    &emsp;|num-multipaths  &emsp;|Internally computed bandwidth based on number of multipaths (outbound-only)  |
- `bandwidth_non_transitive` (Boolean) The link bandwidth extended community is encoded as non-transitive
- `none` (Boolean) Completely remove communities attribute from a prefix
- `rt` (List of String) Set route target value

    &emsp;|Format  &emsp;|Description                                                         |
    |----------|----------------------------------------------------------------------|
    &emsp;|ASN:NN  &emsp;|based on autonomous system number in format &lt;0-65535:0-4294967295&gt;  |
    &emsp;|IP:NN   &emsp;|Based on a router-id IP address in format &lt;IP:0-65535&gt;              |
- `soo` (List of String) Set Site of Origin value

    &emsp;|Format  &emsp;|Description                                                         |
    |----------|----------------------------------------------------------------------|
    &emsp;|ASN:NN  &emsp;|based on autonomous system number in format &lt;0-65535:0-4294967295&gt;  |
    &emsp;|IP:NN   &emsp;|Based on a router-id IP address in format &lt;IP:0-65535&gt;              |


&lt;a id=&#34;nestedatt--set--ipv6_next_hop&#34;&gt;&lt;/a&gt;
### Nested Schema for `set.ipv6_next_hop`

Optional:

- `global` (String) Nexthop IPv6 global address

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv6    &emsp;|IPv6 address and prefix length  |
- `local` (String) Nexthop IPv6 local address

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv6    &emsp;|IPv6 address and prefix length  |
- `peer_address` (Boolean) Use peer address (for BGP only)
- `prefer_global` (Boolean) Prefer global address as the nexthop


&lt;a id=&#34;nestedatt--set--l3vpn_nexthop&#34;&gt;&lt;/a&gt;
### Nested Schema for `set.l3vpn_nexthop`

Optional:

- `encapsulation` (Attributes) Encapsulation options (for BGP only) (see [below for nested schema](#nestedatt--set--l3vpn_nexthop--encapsulation))

&lt;a id=&#34;nestedatt--set--l3vpn_nexthop--encapsulation&#34;&gt;&lt;/a&gt;
### Nested Schema for `set.l3vpn_nexthop.encapsulation`

Optional:

- `gre` (Boolean) Accept L3VPN traffic over GRE encapsulation



&lt;a id=&#34;nestedatt--set--large_community&#34;&gt;&lt;/a&gt;
### Nested Schema for `set.large_community`

Optional:

- `add` (List of String) Add large communities to a prefix ;

    &emsp;|Format          &emsp;|Description                                                   |
    |------------------|----------------------------------------------------------------|
    &emsp;|&lt;GA:LDP1:LDP2&gt;  &emsp;|Community in format &lt;0-4294967295:0-4294967295:0-4294967295&gt;  |
- `delete` (String) Remove communities defined in a list from a prefix

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Community-list  |
- `none` (Boolean) Completely remove communities attribute from a prefix
- `replace` (List of String) Set large communities for a prefix

    &emsp;|Format          &emsp;|Description                                                   |
    |------------------|----------------------------------------------------------------|
    &emsp;|&lt;GA:LDP1:LDP2&gt;  &emsp;|Community in format &lt;0-4294967295:0-4294967295:0-4294967295&gt;  |



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
