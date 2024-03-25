---
page_title: "vyos_vrf_name_protocols_bgp_neighbor Resource - vyos"
subcategory: "vrf"
description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Border Gateway Protocol (BGP)⯯BGP neighbor
---

# vyos_vrf_name_protocols_bgp_neighbor (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Border Gateway Protocol (BGP)  
⯯  
**BGP neighbor**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `neighbor_id` (String) BGP neighbor

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|ipv4    &emsp;|BGP neighbor IP address    |
    &emsp;|ipv6    &emsp;|BGP neighbor IPv6 address  |
    &emsp;|txt     &emsp;|Interface name             |

### Optional

- `address_family` (Attributes) Address-family parameters (see [below for nested schema](#nestedatt--address_family))
- `advertisement_interval` (Number) Minimum interval for sending routing updates

    &emsp;|Format  &emsp;|Description                        |
    |----------|-------------------------------------|
    &emsp;|0-600   &emsp;|Advertisement interval in seconds  |
- `bfd` (Attributes) Enable Bidirectional Forwarding Detection (BFD) support (see [below for nested schema](#nestedatt--bfd))
- `capability` (Attributes) Advertise capabilities to this peer-group (see [below for nested schema](#nestedatt--capability))
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `disable_capability_negotiation` (Boolean) Disable capability negotiation with this neighbor
- `disable_connected_check` (Boolean) Disable check to see if eBGP peer address is a connected route
- `ebgp_multihop` (Number) Allow this EBGP neighbor to not be on a directly connected network

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|1-255   &emsp;|Number of hops  |
- `enforce_first_as` (Boolean) Ensure the first AS in the AS path matches the peer AS
- `graceful_restart` (String) BGP graceful restart functionality

    &emsp;|Format          &emsp;|Description                                            |
    |------------------|---------------------------------------------------------|
    &emsp;|enable          &emsp;|Enable BGP graceful restart at peer level              |
    &emsp;|disable         &emsp;|Disable BGP graceful restart at peer level             |
    &emsp;|restart-helper  &emsp;|Enable BGP graceful restart helper only functionality  |
- `interface` (Attributes) Interface parameters (see [below for nested schema](#nestedatt--interface))
- `override_capability` (Boolean) Ignore capability negotiation with specified neighbor
- `passive` (Boolean) Do not initiate a session with this neighbor
- `password` (String) BGP MD5 password
- `path_attribute` (Attributes) Manipulate path attributes from incoming UPDATE messages (see [below for nested schema](#nestedatt--path_attribute))
- `peer_group` (String) Peer group for this peer

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|Peer-group name  |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `remote_as` (String) Neighbor BGP AS number

    &emsp;|Format        &emsp;|Description                         |
    |----------------|--------------------------------------|
    &emsp;|1-4294967294  &emsp;|Neighbor AS number                  |
    &emsp;|external      &emsp;|Any AS different from the local AS  |
    &emsp;|internal      &emsp;|Neighbor AS number                  |
- `shutdown` (Boolean) Administratively shutdown this neighbor
- `solo` (Boolean) Do not send back prefixes learned from the neighbor
- `strict_capability_match` (Boolean) Enable strict capability negotiation
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `timers` (Attributes) Neighbor timers (see [below for nested schema](#nestedatt--timers))
- `ttl_security` (Attributes) Ttl security mechanism (see [below for nested schema](#nestedatt--ttl_security))
- `update_source` (String) Source IP of routing updates

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address of route source  |
    &emsp;|ipv6    &emsp;|IPv6 address of route source  |
    &emsp;|txt     &emsp;|Interface as route source     |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--address_family&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family`

Optional:

- `ipv4_flowspec` (Attributes) IPv4 Flow Specification BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv4_flowspec))
- `ipv4_labeled_unicast` (Attributes) IPv4 Labeled Unicast BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast))
- `ipv4_multicast` (Attributes) IPv4 Multicast BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv4_multicast))
- `ipv4_unicast` (Attributes) IPv4 BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv4_unicast))
- `ipv4_vpn` (Attributes) IPv4 VPN BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv4_vpn))
- `ipv6_flowspec` (Attributes) IPv6 Flow Specification BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv6_flowspec))
- `ipv6_labeled_unicast` (Attributes) IPv6 Labeled Unicast BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast))
- `ipv6_multicast` (Attributes) IPv6 Multicast BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv6_multicast))
- `ipv6_unicast` (Attributes) IPv6 BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv6_unicast))
- `ipv6_vpn` (Attributes) IPv6 VPN BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv6_vpn))
- `l2vpn_evpn` (Attributes) L2VPN EVPN BGP settings (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn))

&lt;a id=&#34;nestedatt--address_family--ipv4_flowspec&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_flowspec`

Optional:

- `filter_list` (Attributes) as-path-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_flowspec--filter_list))
- `prefix_list` (Attributes) IPv4-Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_flowspec--prefix_list))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_flowspec--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv4_flowspec--soft_reconfiguration))

&lt;a id=&#34;nestedatt--address_family--ipv4_flowspec--filter_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_flowspec.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


&lt;a id=&#34;nestedatt--address_family--ipv4_flowspec--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_flowspec.prefix_list`

Optional:

- `export` (String) IPv4-Prefix-list to filter outgoing route updates to this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |
- `import` (String) IPv4-Prefix-list to filter incoming route updates from this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |


&lt;a id=&#34;nestedatt--address_family--ipv4_flowspec--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_flowspec.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_flowspec--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_flowspec.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast`

Optional:

- `addpath_tx_all` (Boolean) Use addpath to advertise all paths to a neighbor
- `addpath_tx_per_as` (Boolean) Use addpath to advertise the bestpath per each neighboring AS
- `allowas_in` (Attributes) Accept route that contains the local-as in the as-path (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--allowas_in))
- `as_override` (Boolean) Override ASN in outbound updates to configured neighbor local-as
- `attribute_unchanged` (Attributes) BGP attributes are sent unchanged (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--attribute_unchanged))
- `capability` (Attributes) Advertise capabilities to this neighbor (IPv4) (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--capability))
- `conditionally_advertise` (Attributes) Use route-map to conditionally advertise routes (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--conditionally_advertise))
- `default_originate` (Attributes) Originate default route to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--default_originate))
- `disable_send_community` (Attributes) Disable sending community attributes to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--disable_send_community))
- `distribute_list` (Attributes) Access-list to filter route updates to/from this peer-group (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--distribute_list))
- `filter_list` (Attributes) as-path-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--filter_list))
- `maximum_prefix` (Number) Maximum number of prefixes to accept from this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--nexthop_self))
- `prefix_list` (Attributes) IPv4-Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `weight` (Number) Default weight for routes from this peer

    &emsp;|Format   &emsp;|Description     |
    |-----------|------------------|
    &emsp;|1-65535  &emsp;|Default weight  |

&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--allowas_in&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-10    &emsp;|Number of times AS is allowed in path  |


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--attribute_unchanged&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--capability&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.capability`

Optional:

- `orf` (Attributes) Advertise ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--capability--orf))

&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--capability--orf&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.capability.orf`

Optional:

- `prefix_list` (Attributes) Advertise prefix-list ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--capability--orf--prefix_list))

&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--capability--orf--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.capability.orf.prefix_list`

Optional:

- `receive` (Boolean) Capability to receive the ORF
- `send` (Boolean) Capability to send the ORF




&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--conditionally_advertise&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--default_originate&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.default_originate`

Optional:

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--disable_send_community&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--distribute_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    &emsp;|Format   &emsp;|Description                                                      |
    |-----------|-------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    &emsp;|Format   &emsp;|Description                                                        |
    |-----------|---------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter incoming route updates from this peer-group  |


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--filter_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--nexthop_self&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.prefix_list`

Optional:

- `export` (String) IPv4-Prefix-list to filter outgoing route updates to this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |
- `import` (String) IPv4-Prefix-list to filter incoming route updates from this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--remove_private_as&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_labeled_unicast--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_labeled_unicast.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



&lt;a id=&#34;nestedatt--address_family--ipv4_multicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast`

Optional:

- `addpath_tx_all` (Boolean) Use addpath to advertise all paths to a neighbor
- `addpath_tx_per_as` (Boolean) Use addpath to advertise the bestpath per each neighboring AS
- `allowas_in` (Attributes) Accept route that contains the local-as in the as-path (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--allowas_in))
- `as_override` (Boolean) Override ASN in outbound updates to configured neighbor local-as
- `attribute_unchanged` (Attributes) BGP attributes are sent unchanged (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--attribute_unchanged))
- `capability` (Attributes) Advertise capabilities to this neighbor (IPv4) (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--capability))
- `conditionally_advertise` (Attributes) Use route-map to conditionally advertise routes (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--conditionally_advertise))
- `default_originate` (Attributes) Originate default route to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--default_originate))
- `disable_send_community` (Attributes) Disable sending community attributes to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--disable_send_community))
- `distribute_list` (Attributes) Access-list to filter route updates to/from this peer-group (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--distribute_list))
- `filter_list` (Attributes) as-path-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--filter_list))
- `maximum_prefix` (Number) Maximum number of prefixes to accept from this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--nexthop_self))
- `prefix_list` (Attributes) IPv4-Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `weight` (Number) Default weight for routes from this peer

    &emsp;|Format   &emsp;|Description     |
    |-----------|------------------|
    &emsp;|1-65535  &emsp;|Default weight  |

&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--allowas_in&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-10    &emsp;|Number of times AS is allowed in path  |


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--attribute_unchanged&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--capability&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.capability`

Optional:

- `orf` (Attributes) Advertise ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--capability--orf))

&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--capability--orf&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.capability.orf`

Optional:

- `prefix_list` (Attributes) Advertise prefix-list ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_multicast--capability--orf--prefix_list))

&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--capability--orf--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.capability.orf.prefix_list`

Optional:

- `receive` (Boolean) Capability to receive the ORF
- `send` (Boolean) Capability to send the ORF




&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--conditionally_advertise&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--default_originate&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.default_originate`

Optional:

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--disable_send_community&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--distribute_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    &emsp;|Format   &emsp;|Description                                                      |
    |-----------|-------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    &emsp;|Format   &emsp;|Description                                                        |
    |-----------|---------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter incoming route updates from this peer-group  |


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--filter_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--nexthop_self&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.prefix_list`

Optional:

- `export` (String) IPv4-Prefix-list to filter outgoing route updates to this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |
- `import` (String) IPv4-Prefix-list to filter incoming route updates from this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--remove_private_as&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_multicast--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_multicast.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



&lt;a id=&#34;nestedatt--address_family--ipv4_unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast`

Optional:

- `addpath_tx_all` (Boolean) Use addpath to advertise all paths to a neighbor
- `addpath_tx_per_as` (Boolean) Use addpath to advertise the bestpath per each neighboring AS
- `allowas_in` (Attributes) Accept route that contains the local-as in the as-path (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--allowas_in))
- `as_override` (Boolean) Override ASN in outbound updates to configured neighbor local-as
- `attribute_unchanged` (Attributes) BGP attributes are sent unchanged (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--attribute_unchanged))
- `capability` (Attributes) Advertise capabilities to this neighbor (IPv4) (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--capability))
- `conditionally_advertise` (Attributes) Use route-map to conditionally advertise routes (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--conditionally_advertise))
- `default_originate` (Attributes) Originate default route to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--default_originate))
- `disable_send_community` (Attributes) Disable sending community attributes to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--disable_send_community))
- `distribute_list` (Attributes) Access-list to filter route updates to/from this peer-group (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--distribute_list))
- `filter_list` (Attributes) as-path-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--filter_list))
- `maximum_prefix` (Number) Maximum number of prefixes to accept from this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--nexthop_self))
- `prefix_list` (Attributes) IPv4-Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `weight` (Number) Default weight for routes from this peer

    &emsp;|Format   &emsp;|Description     |
    |-----------|------------------|
    &emsp;|1-65535  &emsp;|Default weight  |

&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--allowas_in&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-10    &emsp;|Number of times AS is allowed in path  |


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--attribute_unchanged&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--capability&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.capability`

Optional:

- `orf` (Attributes) Advertise ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--capability--orf))

&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--capability--orf&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.capability.orf`

Optional:

- `prefix_list` (Attributes) Advertise prefix-list ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--capability--orf--prefix_list))

&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--capability--orf--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.capability.orf.prefix_list`

Optional:

- `receive` (Boolean) Capability to receive the ORF
- `send` (Boolean) Capability to send the ORF




&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--conditionally_advertise&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--default_originate&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.default_originate`

Optional:

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--disable_send_community&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--distribute_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    &emsp;|Format   &emsp;|Description                                                      |
    |-----------|-------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    &emsp;|Format   &emsp;|Description                                                        |
    |-----------|---------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter incoming route updates from this peer-group  |


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--filter_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--nexthop_self&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.prefix_list`

Optional:

- `export` (String) IPv4-Prefix-list to filter outgoing route updates to this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |
- `import` (String) IPv4-Prefix-list to filter incoming route updates from this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--remove_private_as&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_unicast--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_unicast.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



&lt;a id=&#34;nestedatt--address_family--ipv4_vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn`

Optional:

- `addpath_tx_all` (Boolean) Use addpath to advertise all paths to a neighbor
- `addpath_tx_per_as` (Boolean) Use addpath to advertise the bestpath per each neighboring AS
- `allowas_in` (Attributes) Accept route that contains the local-as in the as-path (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--allowas_in))
- `as_override` (Boolean) Override ASN in outbound updates to configured neighbor local-as
- `attribute_unchanged` (Attributes) BGP attributes are sent unchanged (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--attribute_unchanged))
- `conditionally_advertise` (Attributes) Use route-map to conditionally advertise routes (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--conditionally_advertise))
- `disable_send_community` (Attributes) Disable sending community attributes to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--disable_send_community))
- `distribute_list` (Attributes) Access-list to filter route updates to/from this peer-group (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--distribute_list))
- `filter_list` (Attributes) as-path-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--filter_list))
- `maximum_prefix` (Number) Maximum number of prefixes to accept from this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--nexthop_self))
- `prefix_list` (Attributes) IPv4-Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `weight` (Number) Default weight for routes from this peer

    &emsp;|Format   &emsp;|Description     |
    |-----------|------------------|
    &emsp;|1-65535  &emsp;|Default weight  |

&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--allowas_in&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-10    &emsp;|Number of times AS is allowed in path  |


&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--attribute_unchanged&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--conditionally_advertise&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--disable_send_community&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--distribute_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    &emsp;|Format   &emsp;|Description                                                      |
    |-----------|-------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    &emsp;|Format   &emsp;|Description                                                        |
    |-----------|---------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter incoming route updates from this peer-group  |


&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--filter_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--nexthop_self&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.prefix_list`

Optional:

- `export` (String) IPv4-Prefix-list to filter outgoing route updates to this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |
- `import` (String) IPv4-Prefix-list to filter incoming route updates from this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |


&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--remove_private_as&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv4_vpn--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv4_vpn.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



&lt;a id=&#34;nestedatt--address_family--ipv6_flowspec&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_flowspec`

Optional:

- `filter_list` (Attributes) as-path-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_flowspec--filter_list))
- `prefix_list` (Attributes) Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_flowspec--prefix_list))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_flowspec--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv6_flowspec--soft_reconfiguration))

&lt;a id=&#34;nestedatt--address_family--ipv6_flowspec--filter_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_flowspec.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_flowspec--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_flowspec.prefix_list`

Optional:

- `export` (String) Prefix-list to filter outgoing route updates to this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |
- `import` (String) Prefix-list to filter incoming route updates from this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |


&lt;a id=&#34;nestedatt--address_family--ipv6_flowspec--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_flowspec.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_flowspec--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_flowspec.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast`

Optional:

- `addpath_tx_all` (Boolean) Use addpath to advertise all paths to a neighbor
- `addpath_tx_per_as` (Boolean) Use addpath to advertise the bestpath per each neighboring AS
- `allowas_in` (Attributes) Accept route that contains the local-as in the as-path (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--allowas_in))
- `as_override` (Boolean) Override ASN in outbound updates to configured neighbor local-as
- `attribute_unchanged` (Attributes) BGP attributes are sent unchanged (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--attribute_unchanged))
- `capability` (Attributes) Advertise capabilities to this neighbor (IPv6) (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--capability))
- `conditionally_advertise` (Attributes) Use route-map to conditionally advertise routes (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--conditionally_advertise))
- `default_originate` (Attributes) Originate default route to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--default_originate))
- `disable_send_community` (Attributes) Disable sending community attributes to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--disable_send_community))
- `distribute_list` (Attributes) Access-list to filter route updates to/from this peer-group (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--distribute_list))
- `filter_list` (Attributes) as-path-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--filter_list))
- `maximum_prefix` (Number) Maximum number of prefixes to accept from this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `nexthop_local` (Attributes) Nexthop attributes (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--nexthop_local))
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--nexthop_self))
- `prefix_list` (Attributes) Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `weight` (Number) Default weight for routes from this peer

    &emsp;|Format   &emsp;|Description     |
    |-----------|------------------|
    &emsp;|1-65535  &emsp;|Default weight  |

&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--allowas_in&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-10    &emsp;|Number of times AS is allowed in path  |


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--attribute_unchanged&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--capability&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.capability`

Optional:

- `orf` (Attributes) Advertise ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--capability--orf))

&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--capability--orf&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.capability.orf`

Optional:

- `prefix_list` (Attributes) Advertise prefix-list ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--capability--orf--prefix_list))

&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--capability--orf--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.capability.orf.prefix_list`

Optional:

- `receive` (Boolean) Capability to receive the ORF
- `send` (Boolean) Capability to send the ORF




&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--conditionally_advertise&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--default_originate&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.default_originate`

Optional:

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--disable_send_community&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--distribute_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    &emsp;|Format   &emsp;|Description                                                      |
    |-----------|-------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    &emsp;|Format   &emsp;|Description                                                        |
    |-----------|---------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter incoming route updates from this peer-group  |


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--filter_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--nexthop_local&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.nexthop_local`

Optional:

- `unchanged` (Boolean) Leave link-local nexthop unchanged for this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--nexthop_self&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.prefix_list`

Optional:

- `export` (String) Prefix-list to filter outgoing route updates to this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |
- `import` (String) Prefix-list to filter incoming route updates from this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--remove_private_as&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_labeled_unicast--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_labeled_unicast.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



&lt;a id=&#34;nestedatt--address_family--ipv6_multicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast`

Optional:

- `addpath_tx_all` (Boolean) Use addpath to advertise all paths to a neighbor
- `addpath_tx_per_as` (Boolean) Use addpath to advertise the bestpath per each neighboring AS
- `allowas_in` (Attributes) Accept route that contains the local-as in the as-path (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--allowas_in))
- `as_override` (Boolean) Override ASN in outbound updates to configured neighbor local-as
- `attribute_unchanged` (Attributes) BGP attributes are sent unchanged (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--attribute_unchanged))
- `conditionally_advertise` (Attributes) Use route-map to conditionally advertise routes (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--conditionally_advertise))
- `default_originate` (Attributes) Originate default route to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--default_originate))
- `disable_send_community` (Attributes) Disable sending community attributes to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--disable_send_community))
- `distribute_list` (Attributes) Access-list to filter route updates to/from this peer-group (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--distribute_list))
- `filter_list` (Attributes) as-path-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--filter_list))
- `maximum_prefix` (Number) Maximum number of prefixes to accept from this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `nexthop_local` (Attributes) Nexthop attributes (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--nexthop_local))
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--nexthop_self))
- `prefix_list` (Attributes) Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv6_multicast--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `weight` (Number) Default weight for routes from this peer

    &emsp;|Format   &emsp;|Description     |
    |-----------|------------------|
    &emsp;|1-65535  &emsp;|Default weight  |

&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--allowas_in&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-10    &emsp;|Number of times AS is allowed in path  |


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--attribute_unchanged&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--conditionally_advertise&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--default_originate&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.default_originate`

Optional:

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--disable_send_community&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--distribute_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    &emsp;|Format   &emsp;|Description                                                      |
    |-----------|-------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    &emsp;|Format   &emsp;|Description                                                        |
    |-----------|---------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter incoming route updates from this peer-group  |


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--filter_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--nexthop_local&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.nexthop_local`

Optional:

- `unchanged` (Boolean) Leave link-local nexthop unchanged for this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--nexthop_self&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.prefix_list`

Optional:

- `export` (String) Prefix-list to filter outgoing route updates to this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |
- `import` (String) Prefix-list to filter incoming route updates from this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--remove_private_as&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_multicast--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_multicast.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



&lt;a id=&#34;nestedatt--address_family--ipv6_unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast`

Optional:

- `addpath_tx_all` (Boolean) Use addpath to advertise all paths to a neighbor
- `addpath_tx_per_as` (Boolean) Use addpath to advertise the bestpath per each neighboring AS
- `allowas_in` (Attributes) Accept route that contains the local-as in the as-path (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--allowas_in))
- `as_override` (Boolean) Override ASN in outbound updates to configured neighbor local-as
- `attribute_unchanged` (Attributes) BGP attributes are sent unchanged (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--attribute_unchanged))
- `capability` (Attributes) Advertise capabilities to this neighbor (IPv6) (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--capability))
- `conditionally_advertise` (Attributes) Use route-map to conditionally advertise routes (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--conditionally_advertise))
- `default_originate` (Attributes) Originate default route to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--default_originate))
- `disable_send_community` (Attributes) Disable sending community attributes to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--disable_send_community))
- `distribute_list` (Attributes) Access-list to filter route updates to/from this peer-group (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--distribute_list))
- `filter_list` (Attributes) as-path-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--filter_list))
- `maximum_prefix` (Number) Maximum number of prefixes to accept from this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `nexthop_local` (Attributes) Nexthop attributes (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--nexthop_local))
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--nexthop_self))
- `prefix_list` (Attributes) Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `weight` (Number) Default weight for routes from this peer

    &emsp;|Format   &emsp;|Description     |
    |-----------|------------------|
    &emsp;|1-65535  &emsp;|Default weight  |

&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--allowas_in&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-10    &emsp;|Number of times AS is allowed in path  |


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--attribute_unchanged&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--capability&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.capability`

Optional:

- `orf` (Attributes) Advertise ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--capability--orf))

&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--capability--orf&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.capability.orf`

Optional:

- `prefix_list` (Attributes) Advertise prefix-list ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--capability--orf--prefix_list))

&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--capability--orf--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.capability.orf.prefix_list`

Optional:

- `receive` (Boolean) Capability to receive the ORF
- `send` (Boolean) Capability to send the ORF




&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--conditionally_advertise&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--default_originate&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.default_originate`

Optional:

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--disable_send_community&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--distribute_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    &emsp;|Format   &emsp;|Description                                                      |
    |-----------|-------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    &emsp;|Format   &emsp;|Description                                                        |
    |-----------|---------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter incoming route updates from this peer-group  |


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--filter_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--nexthop_local&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.nexthop_local`

Optional:

- `unchanged` (Boolean) Leave link-local nexthop unchanged for this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--nexthop_self&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.prefix_list`

Optional:

- `export` (String) Prefix-list to filter outgoing route updates to this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |
- `import` (String) Prefix-list to filter incoming route updates from this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--remove_private_as&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_unicast--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_unicast.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



&lt;a id=&#34;nestedatt--address_family--ipv6_vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn`

Optional:

- `addpath_tx_all` (Boolean) Use addpath to advertise all paths to a neighbor
- `addpath_tx_per_as` (Boolean) Use addpath to advertise the bestpath per each neighboring AS
- `allowas_in` (Attributes) Accept route that contains the local-as in the as-path (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--allowas_in))
- `as_override` (Boolean) Override ASN in outbound updates to configured neighbor local-as
- `attribute_unchanged` (Attributes) BGP attributes are sent unchanged (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--attribute_unchanged))
- `conditionally_advertise` (Attributes) Use route-map to conditionally advertise routes (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--conditionally_advertise))
- `disable_send_community` (Attributes) Disable sending community attributes to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--disable_send_community))
- `distribute_list` (Attributes) Access-list to filter route updates to/from this peer-group (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--distribute_list))
- `filter_list` (Attributes) as-path-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--filter_list))
- `maximum_prefix` (Number) Maximum number of prefixes to accept from this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    &emsp;|Format        &emsp;|Description   |
    |----------------|----------------|
    &emsp;|1-4294967295  &emsp;|Prefix limit  |
- `nexthop_local` (Attributes) Nexthop attributes (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--nexthop_local))
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--nexthop_self))
- `prefix_list` (Attributes) Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `weight` (Number) Default weight for routes from this peer

    &emsp;|Format   &emsp;|Description     |
    |-----------|------------------|
    &emsp;|1-65535  &emsp;|Default weight  |

&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--allowas_in&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-10    &emsp;|Number of times AS is allowed in path  |


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--attribute_unchanged&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--conditionally_advertise&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--disable_send_community&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--distribute_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    &emsp;|Format   &emsp;|Description                                                      |
    |-----------|-------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    &emsp;|Format   &emsp;|Description                                                        |
    |-----------|---------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Access-list to filter incoming route updates from this peer-group  |


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--filter_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--nexthop_local&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.nexthop_local`

Optional:

- `unchanged` (Boolean) Leave link-local nexthop unchanged for this peer


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--nexthop_self&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--prefix_list&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.prefix_list`

Optional:

- `export` (String) Prefix-list to filter outgoing route updates to this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |
- `import` (String) Prefix-list to filter incoming route updates from this peer

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--remove_private_as&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--ipv6_vpn--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.ipv6_vpn.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



&lt;a id=&#34;nestedatt--address_family--l2vpn_evpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.l2vpn_evpn`

Optional:

- `allowas_in` (Attributes) Accept route that contains the local-as in the as-path (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn--allowas_in))
- `attribute_unchanged` (Attributes) BGP attributes are sent unchanged (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn--attribute_unchanged))
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn--nexthop_self))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn--soft_reconfiguration))

&lt;a id=&#34;nestedatt--address_family--l2vpn_evpn--allowas_in&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.l2vpn_evpn.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|1-10    &emsp;|Number of times AS is allowed in path  |


&lt;a id=&#34;nestedatt--address_family--l2vpn_evpn--attribute_unchanged&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.l2vpn_evpn.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


&lt;a id=&#34;nestedatt--address_family--l2vpn_evpn--nexthop_self&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.l2vpn_evpn.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


&lt;a id=&#34;nestedatt--address_family--l2vpn_evpn--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.l2vpn_evpn.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--address_family--l2vpn_evpn--soft_reconfiguration&#34;&gt;&lt;/a&gt;
### Nested Schema for `address_family.l2vpn_evpn.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration




&lt;a id=&#34;nestedatt--bfd&#34;&gt;&lt;/a&gt;
### Nested Schema for `bfd`

Optional:

- `check_control_plane_failure` (Boolean) Allow to write CBIT independence in BFD outgoing packets and read both C-BIT value of BFD and lookup BGP peer status
- `profile` (String) Use settings from BFD profile

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|BFD profile name  |


&lt;a id=&#34;nestedatt--capability&#34;&gt;&lt;/a&gt;
### Nested Schema for `capability`

Optional:

- `dynamic` (Boolean) Advertise dynamic capability to this neighbor
- `extended_nexthop` (Boolean) Advertise extended-nexthop capability to this neighbor
- `software_version` (Boolean) Advertise Software Version capability to the peer


&lt;a id=&#34;nestedatt--interface&#34;&gt;&lt;/a&gt;
### Nested Schema for `interface`

Optional:

- `peer_group` (String) Peer group for this peer

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|Peer-group name  |
- `remote_as` (String) Neighbor BGP AS number

    &emsp;|Format        &emsp;|Description                         |
    |----------------|--------------------------------------|
    &emsp;|1-4294967294  &emsp;|Neighbor AS number                  |
    &emsp;|external      &emsp;|Any AS different from the local AS  |
    &emsp;|internal      &emsp;|Neighbor AS number                  |
- `source_interface` (String) Interface used to establish connection

    &emsp;|Format     &emsp;|Description     |
    |-------------|------------------|
    &emsp;|interface  &emsp;|Interface name  |
- `v6only` (Attributes) Enable BGP with v6 link-local only (see [below for nested schema](#nestedatt--interface--v6only))

&lt;a id=&#34;nestedatt--interface--v6only&#34;&gt;&lt;/a&gt;
### Nested Schema for `interface.v6only`

Optional:

- `peer_group` (String) Peer group for this peer

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|Peer-group name  |
- `remote_as` (String) Neighbor BGP AS number

    &emsp;|Format        &emsp;|Description                         |
    |----------------|--------------------------------------|
    &emsp;|1-4294967294  &emsp;|Neighbor AS number                  |
    &emsp;|external      &emsp;|Any AS different from the local AS  |
    &emsp;|internal      &emsp;|Neighbor AS number                  |



&lt;a id=&#34;nestedatt--path_attribute&#34;&gt;&lt;/a&gt;
### Nested Schema for `path_attribute`

Optional:

- `discard` (List of Number) Drop specified attributes from incoming UPDATE messages

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|1-255   &emsp;|Attribute number  |
- `treat_as_withdraw` (Number) Treat-as-withdraw any incoming BGP UPDATE messages that contain the specified attribute

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|1-255   &emsp;|Attribute number  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--timers&#34;&gt;&lt;/a&gt;
### Nested Schema for `timers`

Optional:

- `connect` (String) BGP connect timer for this neighbor

    &emsp;|Format   &emsp;|Description               |
    |-----------|----------------------------|
    &emsp;|1-65535  &emsp;|Connect timer in seconds  |
    &emsp;|0        &emsp;|Disable connect timer     |
- `holdtime` (String) Hold timer

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|1-65535  &emsp;|Hold timer in seconds  |
    &emsp;|0        &emsp;|Disable hold timer     |
- `keepalive` (Number) BGP keepalive interval for this neighbor

    &emsp;|Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    &emsp;|1-65535  &emsp;|Keepalive interval in seconds  |


&lt;a id=&#34;nestedatt--ttl_security&#34;&gt;&lt;/a&gt;
### Nested Schema for `ttl_security`

Optional:

- `hops` (Number) Number of the maximum number of hops to the BGP peer

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|1-254   &emsp;|Number of hops  |  &emsp;|
