---
page_title: "vyos_policy_route_map_rule Resource - vyos"

subcategory: "Policy"

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

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `action` (String) Action to take on entries matching this rule

    |Format  &emsp;|Description              |
    |----------|---------------------------|
    |permit  &emsp;|Permit matching entries  |
    |deny    &emsp;|Deny matching entries    |
- `call` (String) Call another route-map on match

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Route map name  |
- `continue` (Number) Jump to a different rule in this route-map on a match

    |Format   &emsp;|Description  |
    |-----------|---------------|
    |1-65535  &emsp;|Rule number  |
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `match` (Attributes) Route parameters to match (see [below for nested schema](#nestedatt--match))
- `on_match` (Attributes) Exit policy on matches (see [below for nested schema](#nestedatt--on_match))
- `set` (Attributes) Route parameters (see [below for nested schema](#nestedatt--set))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `route_map` (String) IP route-map

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Route map name  |
- `rule` (Number) Rule for this route-map

    |Format   &emsp;|Description            |
    |-----------|-------------------------|
    |1-65535  &emsp;|Route-map rule number  |


<a id="nestedatt--match"></a>
### Nested Schema for `match`

Optional:

- `as_path` (String) BGP as-path-list to match
- `community` (Attributes) BGP community-list to match (see [below for nested schema](#nestedatt--match--community))
- `evpn` (Attributes) Ethernet Virtual Private Network (see [below for nested schema](#nestedatt--match--evpn))
- `extcommunity` (String) BGP extended community to match
- `interface` (String) Interface to use

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |
- `ip` (Attributes) IP prefix parameters to match (see [below for nested schema](#nestedatt--match--ip))
- `ipv6` (Attributes) IPv6 prefix parameters to match (see [below for nested schema](#nestedatt--match--ipv6))
- `large_community` (Attributes) Match BGP large communities (see [below for nested schema](#nestedatt--match--large_community))
- `local_preference` (Number) Local Preference

    |Format        &emsp;|Description       |
    |----------------|--------------------|
    |0-4294967295  &emsp;|Local Preference  |
- `metric` (Number) Metric of route to match

    |Format   &emsp;|Description   |
    |-----------|----------------|
    |1-65535  &emsp;|Route metric  |
- `origin` (String) BGP origin code to match

    |Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    |egp         &emsp;|Exterior gateway protocol origin  |
    |igp         &emsp;|Interior gateway protocol origin  |
    |incomplete  &emsp;|Incomplete origin                 |
- `peer` (String) Peer address to match

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |ipv4    &emsp;|Peer IP address    |
    |ipv6    &emsp;|Peer IPv6 address  |
- `protocol` (String) Match protocol via which the route was learnt

    |Format     &emsp;|Description                                                  |
    |-------------|---------------------------------------------------------------|
    |babel      &emsp;|Babel routing protocol (Babel)                               |
    |bgp        &emsp;|Border Gateway Protocol (BGP)                                |
    |connected  &emsp;|Connected routes (directly attached subnet or host)          |
    |isis       &emsp;|Intermediate System to Intermediate System (IS-IS)           |
    |kernel     &emsp;|Kernel routes                                                |
    |ospf       &emsp;|Open Shortest Path First (OSPFv2)                            |
    |ospfv3     &emsp;|Open Shortest Path First (IPv6) (OSPFv3)                     |
    |rip        &emsp;|Routing Information Protocol (RIP)                           |
    |ripng      &emsp;|Routing Information Protocol next-generation (IPv6) (RIPng)  |
    |static     &emsp;|Statically configured routes                                 |
    |table      &emsp;|Non-main Kernel Routing Table                                |
    |vnc        &emsp;|Virtual Network Control (VNC)                                |
- `rpki` (String) Match RPKI validation result

    |Format    &emsp;|Description             |
    |------------|--------------------------|
    |invalid   &emsp;|Match invalid entries   |
    |notfound  &emsp;|Match notfound entries  |
    |valid     &emsp;|Match valid entries     |
- `tag` (Number) Route tag value

    |Format   &emsp;|Description  |
    |-----------|---------------|
    |1-65535  &emsp;|Route tag    |

<a id="nestedatt--match--community"></a>
### Nested Schema for `match.community`

Optional:

- `community_list` (String) BGP community-list to match
- `exact_match` (Boolean) Community-list to exactly match


<a id="nestedatt--match--evpn"></a>
### Nested Schema for `match.evpn`

Optional:

- `default_route` (Boolean) Default EVPN type-5 route
- `rd` (String) Route Distinguisher

    |Format                   &emsp;|Description                                   |
    |---------------------------|------------------------------------------------|
    |ASN:NN_OR_IP-ADDRESS:NN  &emsp;|Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
- `route_type` (String) Match route-type

    |Format     &emsp;|Description   |
    |-------------|----------------|
    |macip      &emsp;|mac-ip route  |
    |multicast  &emsp;|IMET route    |
    |prefix     &emsp;|Prefix route  |
- `vni` (Number) Virtual Network Identifier

    |Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    |0-16777214  &emsp;|VXLAN virtual network identifier  |


<a id="nestedatt--match--ip"></a>
### Nested Schema for `match.ip`

Optional:

- `address` (Attributes) IP address of route to match (see [below for nested schema](#nestedatt--match--ip--address))
- `nexthop` (Attributes) IP next-hop of route to match (see [below for nested schema](#nestedatt--match--ip--nexthop))
- `route_source` (Attributes) Match advertising source address of route (see [below for nested schema](#nestedatt--match--ip--route_source))

<a id="nestedatt--match--ip--address"></a>
### Nested Schema for `match.ip.address`

Optional:

- `access_list` (Number) IP access-list to match

    |Format     &emsp;|Description                               |
    |-------------|--------------------------------------------|
    |1-99       &emsp;|IP standard access list                   |
    |100-199    &emsp;|IP extended access list                   |
    |1300-1999  &emsp;|IP standard access list (expanded range)  |
    |2000-2699  &emsp;|IP extended access list (expanded range)  |
- `prefix_len` (Number) IP prefix-length to match (can be used for kernel routes only)

    |Format  &emsp;|Description    |
    |----------|-----------------|
    |0-32    &emsp;|Prefix length  |
- `prefix_list` (String) IP prefix-list to match


<a id="nestedatt--match--ip--nexthop"></a>
### Nested Schema for `match.ip.nexthop`

Optional:

- `access_list` (Number) IP access-list to match

    |Format     &emsp;|Description                               |
    |-------------|--------------------------------------------|
    |1-99       &emsp;|IP standard access list                   |
    |100-199    &emsp;|IP extended access list                   |
    |1300-1999  &emsp;|IP standard access list (expanded range)  |
    |2000-2699  &emsp;|IP extended access list (expanded range)  |
- `address` (String) IP address to match

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |ipv4    &emsp;|Nexthop IP address  |
- `prefix_len` (Number) IP prefix-length to match

    |Format  &emsp;|Description    |
    |----------|-----------------|
    |0-32    &emsp;|Prefix length  |
- `prefix_list` (String) IP prefix-list to match
- `type` (String) Match type

    |Format     &emsp;|Description  |
    |-------------|---------------|
    |blackhole  &emsp;|Blackhole    |


<a id="nestedatt--match--ip--route_source"></a>
### Nested Schema for `match.ip.route_source`

Optional:

- `access_list` (Number) IP access-list to match

    |Format     &emsp;|Description                               |
    |-------------|--------------------------------------------|
    |1-99       &emsp;|IP standard access list                   |
    |100-199    &emsp;|IP extended access list                   |
    |1300-1999  &emsp;|IP standard access list (expanded range)  |
    |2000-2699  &emsp;|IP extended access list (expanded range)  |
- `prefix_list` (String) IP prefix-list to match



<a id="nestedatt--match--ipv6"></a>
### Nested Schema for `match.ipv6`

Optional:

- `address` (Attributes) IPv6 address of route to match (see [below for nested schema](#nestedatt--match--ipv6--address))
- `nexthop` (Attributes) IPv6 next-hop of route to match (see [below for nested schema](#nestedatt--match--ipv6--nexthop))

<a id="nestedatt--match--ipv6--address"></a>
### Nested Schema for `match.ipv6.address`

Optional:

- `access_list` (String) IPv6 access-list to match

    |Format  &emsp;|Description            |
    |----------|-------------------------|
    |txt     &emsp;|IPV6 access list name  |
- `prefix_len` (Number) IPv6 prefix-length to match (can be used for kernel routes only)

    |Format  &emsp;|Description    |
    |----------|-----------------|
    |0-128   &emsp;|Prefix length  |
- `prefix_list` (String) IPv6 prefix-list to match


<a id="nestedatt--match--ipv6--nexthop"></a>
### Nested Schema for `match.ipv6.nexthop`

Optional:

- `access_list` (String) IPv6 access-list to match

    |Format  &emsp;|Description            |
    |----------|-------------------------|
    |txt     &emsp;|IPV6 access list name  |
- `address` (String) IPv6 address of next-hop

    |Format  &emsp;|Description           |
    |----------|------------------------|
    |ipv6    &emsp;|Nexthop IPv6 address  |
- `prefix_list` (String) IPv6 prefix-list to match
- `type` (String) Match type

    |Format     &emsp;|Description  |
    |-------------|---------------|
    |blackhole  &emsp;|Blackhole    |



<a id="nestedatt--match--large_community"></a>
### Nested Schema for `match.large_community`

Optional:

- `large_community_list` (String) BGP large-community-list to match



<a id="nestedatt--on_match"></a>
### Nested Schema for `on_match`

Optional:

- `goto` (Number) Rule number to goto on match

    |Format   &emsp;|Description  |
    |-----------|---------------|
    |1-65535  &emsp;|Rule number  |
- `next` (Boolean) Next sequence number to goto on match


<a id="nestedatt--set"></a>
### Nested Schema for `set`

Optional:

- `aggregator` (Attributes) BGP aggregator attribute (see [below for nested schema](#nestedatt--set--aggregator))
- `as_path` (Attributes) Transform BGP AS_PATH attribute (see [below for nested schema](#nestedatt--set--as_path))
- `atomic_aggregate` (Boolean) BGP atomic aggregate attribute
- `community` (Attributes) BGP community attribute (see [below for nested schema](#nestedatt--set--community))
- `distance` (Number) Locally significant administrative distance

    |Format  &emsp;|Description     |
    |----------|------------------|
    |0-255   &emsp;|Distance value  |
- `evpn` (Attributes) Ethernet Virtual Private Network (see [below for nested schema](#nestedatt--set--evpn))
- `extcommunity` (Attributes) BGP extended community attribute (see [below for nested schema](#nestedatt--set--extcommunity))
- `ip_next_hop` (String) Nexthop IP address

    |Format        &emsp;|Description                                             |
    |----------------|----------------------------------------------------------|
    |ipv4          &emsp;|IP address                                              |
    |unchanged     &emsp;|Set the BGP nexthop address as unchanged                |
    |peer-address  &emsp;|Set the BGP nexthop address to the address of the peer  |
- `ipv6_next_hop` (Attributes) Nexthop IPv6 address (see [below for nested schema](#nestedatt--set--ipv6_next_hop))
- `l3vpn_nexthop` (Attributes) Next hop Information (see [below for nested schema](#nestedatt--set--l3vpn_nexthop))
- `large_community` (Attributes) BGP large community attribute (see [below for nested schema](#nestedatt--set--large_community))
- `local_preference` (Number) BGP local preference attribute

    |Format        &emsp;|Description             |
    |----------------|--------------------------|
    |0-4294967295  &emsp;|Local preference value  |
- `metric` (String) Destination routing protocol metric

    |Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    |&lt;+/-metric&gt;   &emsp;|Add or subtract metric           |
    |0-4294967295  &emsp;|Metric value                     |
    |&lt;+/-rtt&gt;      &emsp;|Add or subtract round trip time  |
    |&lt;rtt&gt;         &emsp;|Round trip time                  |
- `metric_type` (String) Open Shortest Path First (OSPF) external metric-type

    |Format  &emsp;|Description                  |
    |----------|-------------------------------|
    |type-1  &emsp;|OSPF external type 1 metric  |
    |type-2  &emsp;|OSPF external type 2 metric  |
- `origin` (String) Border Gateway Protocl (BGP) origin code

    |Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    |igp         &emsp;|Interior gateway protocol origin  |
    |egp         &emsp;|Exterior gateway protocol origin  |
    |incomplete  &emsp;|Incomplete origin                 |
- `originator_id` (String) BGP originator ID attribute

    |Format  &emsp;|Description           |
    |----------|------------------------|
    |ipv4    &emsp;|Orignator IP address  |
- `src` (String) Source address for route

    |Format  &emsp;|Description   |
    |----------|----------------|
    |ipv4    &emsp;|IPv4 address  |
    |ipv6    &emsp;|IPv6 address  |
- `table` (Number) Set prefixes to table

    |Format        &emsp;|Description  |
    |----------------|---------------|
    |1-4294967295  &emsp;|Table value  |
- `tag` (Number) Route tag value

    |Format   &emsp;|Description  |
    |-----------|---------------|
    |1-65535  &emsp;|Route tag    |
- `weight` (Number) BGP weight attribute

    |Format        &emsp;|Description  |
    |----------------|---------------|
    |0-4294967295  &emsp;|BGP weight   |

<a id="nestedatt--set--aggregator"></a>
### Nested Schema for `set.aggregator`

Optional:

- `as` (Number) AS number of an aggregation

    |Format        &emsp;|Description  |
    |----------------|---------------|
    |1-4294967295  &emsp;|Rule number  |
- `ip` (String) IP address of an aggregation

    |Format  &emsp;|Description  |
    |----------|---------------|
    |ipv4    &emsp;|IP address   |


<a id="nestedatt--set--as_path"></a>
### Nested Schema for `set.as_path`

Optional:

- `exclude` (String) Remove/exclude from the as-path attribute

    |Format        &emsp;|Description                              |
    |----------------|-------------------------------------------|
    |1-4294967295  &emsp;|AS number                                |
    |all           &emsp;|Exclude all AS numbers from the as-path  |
- `prepend` (Number) Prepend to the as-path

    |Format        &emsp;|Description  |
    |----------------|---------------|
    |1-4294967295  &emsp;|AS number    |
- `prepend_last_as` (Number) Use the last AS-number in the as-path

    |Format  &emsp;|Description                |
    |----------|-----------------------------|
    |1-10    &emsp;|Number of times to insert  |


<a id="nestedatt--set--community"></a>
### Nested Schema for `set.community`

Optional:

- `add` (List of String) Add communities to a prefix

    |Format                      &emsp;|Description                                                         |
    |------------------------------|----------------------------------------------------------------------|
    |&lt;AS:VAL&gt;                    &emsp;|Community number in &lt;0-65535:0-65535&gt; format                        |
    |local-as                    &emsp;|Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03         |
    |no-advertise                &emsp;|Well-known communities value NO_ADVERTISE 0xFFFFFF02                |
    |no-export                   &emsp;|Well-known communities value NO_EXPORT 0xFFFFFF01                   |
    |internet                    &emsp;|Well-known communities value 0                                      |
    |graceful-shutdown           &emsp;|Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000           |
    |accept-own                  &emsp;|Well-known communities value ACCEPT_OWN 0xFFFF0001                  |
    |route-filter-translated-v4  &emsp;|Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002  |
    |route-filter-v4             &emsp;|Well-known communities value ROUTE_FILTER_v4 0xFFFF0003             |
    |route-filter-translated-v6  &emsp;|Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004  |
    |route-filter-v6             &emsp;|Well-known communities value ROUTE_FILTER_v6 0xFFFF0005             |
    |llgr-stale                  &emsp;|Well-known communities value LLGR_STALE 0xFFFF0006                  |
    |no-llgr                     &emsp;|Well-known communities value NO_LLGR 0xFFFF0007                     |
    |accept-own-nexthop          &emsp;|Well-known communities value accept-own-nexthop 0xFFFF0008          |
    |blackhole                   &emsp;|Well-known communities value BLACKHOLE 0xFFFF029A                   |
    |no-peer                     &emsp;|Well-known communities value NOPEER 0xFFFFFF04                      |
- `delete` (String) Remove communities defined in a list from a prefix

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Community-list  |
- `none` (Boolean) Completely remove communities attribute from a prefix
- `replace` (List of String) Set communities for a prefix

    |Format                      &emsp;|Description                                                         |
    |------------------------------|----------------------------------------------------------------------|
    |&lt;AS:VAL&gt;                    &emsp;|Community number in &lt;0-65535:0-65535&gt; format                        |
    |local-as                    &emsp;|Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03         |
    |no-advertise                &emsp;|Well-known communities value NO_ADVERTISE 0xFFFFFF02                |
    |no-export                   &emsp;|Well-known communities value NO_EXPORT 0xFFFFFF01                   |
    |internet                    &emsp;|Well-known communities value 0                                      |
    |graceful-shutdown           &emsp;|Well-known communities value GRACEFUL_SHUTDOWN 0xFFFF0000           |
    |accept-own                  &emsp;|Well-known communities value ACCEPT_OWN 0xFFFF0001                  |
    |route-filter-translated-v4  &emsp;|Well-known communities value ROUTE_FILTER_TRANSLATED_v4 0xFFFF0002  |
    |route-filter-v4             &emsp;|Well-known communities value ROUTE_FILTER_v4 0xFFFF0003             |
    |route-filter-translated-v6  &emsp;|Well-known communities value ROUTE_FILTER_TRANSLATED_v6 0xFFFF0004  |
    |route-filter-v6             &emsp;|Well-known communities value ROUTE_FILTER_v6 0xFFFF0005             |
    |llgr-stale                  &emsp;|Well-known communities value LLGR_STALE 0xFFFF0006                  |
    |no-llgr                     &emsp;|Well-known communities value NO_LLGR 0xFFFF0007                     |
    |accept-own-nexthop          &emsp;|Well-known communities value accept-own-nexthop 0xFFFF0008          |
    |blackhole                   &emsp;|Well-known communities value BLACKHOLE 0xFFFF029A                   |
    |no-peer                     &emsp;|Well-known communities value NOPEER 0xFFFFFF04                      |


<a id="nestedatt--set--evpn"></a>
### Nested Schema for `set.evpn`

Optional:

- `gateway` (Attributes) Set gateway IP for prefix advertisement route (see [below for nested schema](#nestedatt--set--evpn--gateway))

<a id="nestedatt--set--evpn--gateway"></a>
### Nested Schema for `set.evpn.gateway`

Optional:

- `ipv4` (String) Set gateway IPv4 address

    |Format  &emsp;|Description           |
    |----------|------------------------|
    |ipv4    &emsp;|Gateway IPv4 address  |
- `ipv6` (String) Set gateway IPv6 address

    |Format  &emsp;|Description           |
    |----------|------------------------|
    |ipv6    &emsp;|Gateway IPv6 address  |



<a id="nestedatt--set--extcommunity"></a>
### Nested Schema for `set.extcommunity`

Optional:

- `bandwidth` (String) Bandwidth value in Mbps

    |Format          &emsp;|Description                                                                  |
    |------------------|-------------------------------------------------------------------------------|
    |1-25600         &emsp;|Bandwidth value in Mbps                                                      |
    |cumulative      &emsp;|Cumulative bandwidth of all multipaths (outbound-only)                       |
    |num-multipaths  &emsp;|Internally computed bandwidth based on number of multipaths (outbound-only)  |
- `bandwidth_non_transitive` (Boolean) The link bandwidth extended community is encoded as non-transitive
- `none` (Boolean) Completely remove communities attribute from a prefix
- `rt` (List of String) Set route target value

    |Format  &emsp;|Description                                                         |
    |----------|----------------------------------------------------------------------|
    |ASN:NN  &emsp;|based on autonomous system number in format &lt;0-65535:0-4294967295&gt;  |
    |IP:NN   &emsp;|Based on a router-id IP address in format &lt;IP:0-65535&gt;              |
- `soo` (List of String) Set Site of Origin value

    |Format  &emsp;|Description                                                         |
    |----------|----------------------------------------------------------------------|
    |ASN:NN  &emsp;|based on autonomous system number in format &lt;0-65535:0-4294967295&gt;  |
    |IP:NN   &emsp;|Based on a router-id IP address in format &lt;IP:0-65535&gt;              |


<a id="nestedatt--set--ipv6_next_hop"></a>
### Nested Schema for `set.ipv6_next_hop`

Optional:

- `global` (String) Nexthop IPv6 global address

    |Format  &emsp;|Description                     |
    |----------|----------------------------------|
    |ipv6    &emsp;|IPv6 address and prefix length  |
- `local` (String) Nexthop IPv6 local address

    |Format  &emsp;|Description                     |
    |----------|----------------------------------|
    |ipv6    &emsp;|IPv6 address and prefix length  |
- `peer_address` (Boolean) Use peer address (for BGP only)
- `prefer_global` (Boolean) Prefer global address as the nexthop


<a id="nestedatt--set--l3vpn_nexthop"></a>
### Nested Schema for `set.l3vpn_nexthop`

Optional:

- `encapsulation` (Attributes) Encapsulation options (for BGP only) (see [below for nested schema](#nestedatt--set--l3vpn_nexthop--encapsulation))

<a id="nestedatt--set--l3vpn_nexthop--encapsulation"></a>
### Nested Schema for `set.l3vpn_nexthop.encapsulation`

Optional:

- `gre` (Boolean) Accept L3VPN traffic over GRE encapsulation



<a id="nestedatt--set--large_community"></a>
### Nested Schema for `set.large_community`

Optional:

- `add` (List of String) Add large communities to a prefix ;

    |Format          &emsp;|Description                                                   |
    |------------------|----------------------------------------------------------------|
    |&lt;GA:LDP1:LDP2&gt;  &emsp;|Community in format &lt;0-4294967295:0-4294967295:0-4294967295&gt;  |
- `delete` (String) Remove communities defined in a list from a prefix

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Community-list  |
- `none` (Boolean) Completely remove communities attribute from a prefix
- `replace` (List of String) Set large communities for a prefix

    |Format          &emsp;|Description                                                   |
    |------------------|----------------------------------------------------------------|
    |&lt;GA:LDP1:LDP2&gt;  &emsp;|Community in format &lt;0-4294967295:0-4294967295:0-4294967295&gt;  |



<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
