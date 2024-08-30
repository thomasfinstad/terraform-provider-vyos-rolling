---
page_title: "vyos_protocols_bgp_peer_group Resource - vyos"

subcategory: "Protocols"

description: |-
  protocols⯯Border Gateway Protocol (BGP)⯯Name of peer-group
---

# vyos_protocols_bgp_peer_group (Resource)
<center>

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
**Name of peer-group**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_protocols_bgp_peer_group (Resource)](#vyos_protocols_bgp_peer_group-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [address_family](#address_family)
      - [bfd](#bfd)
      - [capability](#capability)
      - [description](#description)
      - [disable_capability_negotiation](#disable_capability_negotiation)
      - [disable_connected_check](#disable_connected_check)
      - [ebgp_multihop](#ebgp_multihop)
      - [graceful_restart](#graceful_restart)
      - [override_capability](#override_capability)
      - [passive](#passive)
      - [password](#password)
      - [path_attribute](#path_attribute)
      - [port](#port)
      - [remote_as](#remote_as)
      - [shutdown](#shutdown)
      - [timeouts](#timeouts)
      - [ttl_security](#ttl_security)
      - [update_source](#update_source)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `address_family`](#nested-schema-for-address_family)
    - [Nested Schema for `address_family.ipv4_labeled_unicast`](#nested-schema-for-address_familyipv4_labeled_unicast)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.allowas_in`](#nested-schema-for-address_familyipv4_labeled_unicastallowas_in)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.attribute_unchanged`](#nested-schema-for-address_familyipv4_labeled_unicastattribute_unchanged)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.capability`](#nested-schema-for-address_familyipv4_labeled_unicastcapability)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.capability.orf`](#nested-schema-for-address_familyipv4_labeled_unicastcapabilityorf)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.capability.orf.prefix_list`](#nested-schema-for-address_familyipv4_labeled_unicastcapabilityorfprefix_list)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.conditionally_advertise`](#nested-schema-for-address_familyipv4_labeled_unicastconditionally_advertise)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.default_originate`](#nested-schema-for-address_familyipv4_labeled_unicastdefault_originate)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.disable_send_community`](#nested-schema-for-address_familyipv4_labeled_unicastdisable_send_community)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.distribute_list`](#nested-schema-for-address_familyipv4_labeled_unicastdistribute_list)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.filter_list`](#nested-schema-for-address_familyipv4_labeled_unicastfilter_list)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.nexthop_self`](#nested-schema-for-address_familyipv4_labeled_unicastnexthop_self)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.prefix_list`](#nested-schema-for-address_familyipv4_labeled_unicastprefix_list)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.remove_private_as`](#nested-schema-for-address_familyipv4_labeled_unicastremove_private_as)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.route_map`](#nested-schema-for-address_familyipv4_labeled_unicastroute_map)
    - [Nested Schema for `address_family.ipv4_labeled_unicast.soft_reconfiguration`](#nested-schema-for-address_familyipv4_labeled_unicastsoft_reconfiguration)
    - [Nested Schema for `address_family.ipv4_unicast`](#nested-schema-for-address_familyipv4_unicast)
    - [Nested Schema for `address_family.ipv4_unicast.allowas_in`](#nested-schema-for-address_familyipv4_unicastallowas_in)
    - [Nested Schema for `address_family.ipv4_unicast.attribute_unchanged`](#nested-schema-for-address_familyipv4_unicastattribute_unchanged)
    - [Nested Schema for `address_family.ipv4_unicast.capability`](#nested-schema-for-address_familyipv4_unicastcapability)
    - [Nested Schema for `address_family.ipv4_unicast.capability.orf`](#nested-schema-for-address_familyipv4_unicastcapabilityorf)
    - [Nested Schema for `address_family.ipv4_unicast.capability.orf.prefix_list`](#nested-schema-for-address_familyipv4_unicastcapabilityorfprefix_list)
    - [Nested Schema for `address_family.ipv4_unicast.conditionally_advertise`](#nested-schema-for-address_familyipv4_unicastconditionally_advertise)
    - [Nested Schema for `address_family.ipv4_unicast.default_originate`](#nested-schema-for-address_familyipv4_unicastdefault_originate)
    - [Nested Schema for `address_family.ipv4_unicast.disable_send_community`](#nested-schema-for-address_familyipv4_unicastdisable_send_community)
    - [Nested Schema for `address_family.ipv4_unicast.distribute_list`](#nested-schema-for-address_familyipv4_unicastdistribute_list)
    - [Nested Schema for `address_family.ipv4_unicast.filter_list`](#nested-schema-for-address_familyipv4_unicastfilter_list)
    - [Nested Schema for `address_family.ipv4_unicast.nexthop_self`](#nested-schema-for-address_familyipv4_unicastnexthop_self)
    - [Nested Schema for `address_family.ipv4_unicast.prefix_list`](#nested-schema-for-address_familyipv4_unicastprefix_list)
    - [Nested Schema for `address_family.ipv4_unicast.remove_private_as`](#nested-schema-for-address_familyipv4_unicastremove_private_as)
    - [Nested Schema for `address_family.ipv4_unicast.route_map`](#nested-schema-for-address_familyipv4_unicastroute_map)
    - [Nested Schema for `address_family.ipv4_unicast.soft_reconfiguration`](#nested-schema-for-address_familyipv4_unicastsoft_reconfiguration)
    - [Nested Schema for `address_family.ipv4_vpn`](#nested-schema-for-address_familyipv4_vpn)
    - [Nested Schema for `address_family.ipv4_vpn.allowas_in`](#nested-schema-for-address_familyipv4_vpnallowas_in)
    - [Nested Schema for `address_family.ipv4_vpn.attribute_unchanged`](#nested-schema-for-address_familyipv4_vpnattribute_unchanged)
    - [Nested Schema for `address_family.ipv4_vpn.conditionally_advertise`](#nested-schema-for-address_familyipv4_vpnconditionally_advertise)
    - [Nested Schema for `address_family.ipv4_vpn.disable_send_community`](#nested-schema-for-address_familyipv4_vpndisable_send_community)
    - [Nested Schema for `address_family.ipv4_vpn.distribute_list`](#nested-schema-for-address_familyipv4_vpndistribute_list)
    - [Nested Schema for `address_family.ipv4_vpn.filter_list`](#nested-schema-for-address_familyipv4_vpnfilter_list)
    - [Nested Schema for `address_family.ipv4_vpn.nexthop_self`](#nested-schema-for-address_familyipv4_vpnnexthop_self)
    - [Nested Schema for `address_family.ipv4_vpn.prefix_list`](#nested-schema-for-address_familyipv4_vpnprefix_list)
    - [Nested Schema for `address_family.ipv4_vpn.remove_private_as`](#nested-schema-for-address_familyipv4_vpnremove_private_as)
    - [Nested Schema for `address_family.ipv4_vpn.route_map`](#nested-schema-for-address_familyipv4_vpnroute_map)
    - [Nested Schema for `address_family.ipv4_vpn.soft_reconfiguration`](#nested-schema-for-address_familyipv4_vpnsoft_reconfiguration)
    - [Nested Schema for `address_family.ipv6_labeled_unicast`](#nested-schema-for-address_familyipv6_labeled_unicast)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.allowas_in`](#nested-schema-for-address_familyipv6_labeled_unicastallowas_in)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.attribute_unchanged`](#nested-schema-for-address_familyipv6_labeled_unicastattribute_unchanged)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.capability`](#nested-schema-for-address_familyipv6_labeled_unicastcapability)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.capability.orf`](#nested-schema-for-address_familyipv6_labeled_unicastcapabilityorf)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.capability.orf.prefix_list`](#nested-schema-for-address_familyipv6_labeled_unicastcapabilityorfprefix_list)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.conditionally_advertise`](#nested-schema-for-address_familyipv6_labeled_unicastconditionally_advertise)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.default_originate`](#nested-schema-for-address_familyipv6_labeled_unicastdefault_originate)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.disable_send_community`](#nested-schema-for-address_familyipv6_labeled_unicastdisable_send_community)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.distribute_list`](#nested-schema-for-address_familyipv6_labeled_unicastdistribute_list)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.filter_list`](#nested-schema-for-address_familyipv6_labeled_unicastfilter_list)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.nexthop_local`](#nested-schema-for-address_familyipv6_labeled_unicastnexthop_local)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.nexthop_self`](#nested-schema-for-address_familyipv6_labeled_unicastnexthop_self)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.prefix_list`](#nested-schema-for-address_familyipv6_labeled_unicastprefix_list)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.remove_private_as`](#nested-schema-for-address_familyipv6_labeled_unicastremove_private_as)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.route_map`](#nested-schema-for-address_familyipv6_labeled_unicastroute_map)
    - [Nested Schema for `address_family.ipv6_labeled_unicast.soft_reconfiguration`](#nested-schema-for-address_familyipv6_labeled_unicastsoft_reconfiguration)
    - [Nested Schema for `address_family.ipv6_unicast`](#nested-schema-for-address_familyipv6_unicast)
    - [Nested Schema for `address_family.ipv6_unicast.allowas_in`](#nested-schema-for-address_familyipv6_unicastallowas_in)
    - [Nested Schema for `address_family.ipv6_unicast.attribute_unchanged`](#nested-schema-for-address_familyipv6_unicastattribute_unchanged)
    - [Nested Schema for `address_family.ipv6_unicast.capability`](#nested-schema-for-address_familyipv6_unicastcapability)
    - [Nested Schema for `address_family.ipv6_unicast.capability.orf`](#nested-schema-for-address_familyipv6_unicastcapabilityorf)
    - [Nested Schema for `address_family.ipv6_unicast.capability.orf.prefix_list`](#nested-schema-for-address_familyipv6_unicastcapabilityorfprefix_list)
    - [Nested Schema for `address_family.ipv6_unicast.conditionally_advertise`](#nested-schema-for-address_familyipv6_unicastconditionally_advertise)
    - [Nested Schema for `address_family.ipv6_unicast.default_originate`](#nested-schema-for-address_familyipv6_unicastdefault_originate)
    - [Nested Schema for `address_family.ipv6_unicast.disable_send_community`](#nested-schema-for-address_familyipv6_unicastdisable_send_community)
    - [Nested Schema for `address_family.ipv6_unicast.distribute_list`](#nested-schema-for-address_familyipv6_unicastdistribute_list)
    - [Nested Schema for `address_family.ipv6_unicast.filter_list`](#nested-schema-for-address_familyipv6_unicastfilter_list)
    - [Nested Schema for `address_family.ipv6_unicast.nexthop_local`](#nested-schema-for-address_familyipv6_unicastnexthop_local)
    - [Nested Schema for `address_family.ipv6_unicast.nexthop_self`](#nested-schema-for-address_familyipv6_unicastnexthop_self)
    - [Nested Schema for `address_family.ipv6_unicast.prefix_list`](#nested-schema-for-address_familyipv6_unicastprefix_list)
    - [Nested Schema for `address_family.ipv6_unicast.remove_private_as`](#nested-schema-for-address_familyipv6_unicastremove_private_as)
    - [Nested Schema for `address_family.ipv6_unicast.route_map`](#nested-schema-for-address_familyipv6_unicastroute_map)
    - [Nested Schema for `address_family.ipv6_unicast.soft_reconfiguration`](#nested-schema-for-address_familyipv6_unicastsoft_reconfiguration)
    - [Nested Schema for `address_family.ipv6_vpn`](#nested-schema-for-address_familyipv6_vpn)
    - [Nested Schema for `address_family.ipv6_vpn.allowas_in`](#nested-schema-for-address_familyipv6_vpnallowas_in)
    - [Nested Schema for `address_family.ipv6_vpn.attribute_unchanged`](#nested-schema-for-address_familyipv6_vpnattribute_unchanged)
    - [Nested Schema for `address_family.ipv6_vpn.conditionally_advertise`](#nested-schema-for-address_familyipv6_vpnconditionally_advertise)
    - [Nested Schema for `address_family.ipv6_vpn.disable_send_community`](#nested-schema-for-address_familyipv6_vpndisable_send_community)
    - [Nested Schema for `address_family.ipv6_vpn.distribute_list`](#nested-schema-for-address_familyipv6_vpndistribute_list)
    - [Nested Schema for `address_family.ipv6_vpn.filter_list`](#nested-schema-for-address_familyipv6_vpnfilter_list)
    - [Nested Schema for `address_family.ipv6_vpn.nexthop_local`](#nested-schema-for-address_familyipv6_vpnnexthop_local)
    - [Nested Schema for `address_family.ipv6_vpn.nexthop_self`](#nested-schema-for-address_familyipv6_vpnnexthop_self)
    - [Nested Schema for `address_family.ipv6_vpn.prefix_list`](#nested-schema-for-address_familyipv6_vpnprefix_list)
    - [Nested Schema for `address_family.ipv6_vpn.remove_private_as`](#nested-schema-for-address_familyipv6_vpnremove_private_as)
    - [Nested Schema for `address_family.ipv6_vpn.route_map`](#nested-schema-for-address_familyipv6_vpnroute_map)
    - [Nested Schema for `address_family.ipv6_vpn.soft_reconfiguration`](#nested-schema-for-address_familyipv6_vpnsoft_reconfiguration)
    - [Nested Schema for `address_family.l2vpn_evpn`](#nested-schema-for-address_familyl2vpn_evpn)
    - [Nested Schema for `address_family.l2vpn_evpn.allowas_in`](#nested-schema-for-address_familyl2vpn_evpnallowas_in)
    - [Nested Schema for `address_family.l2vpn_evpn.attribute_unchanged`](#nested-schema-for-address_familyl2vpn_evpnattribute_unchanged)
    - [Nested Schema for `address_family.l2vpn_evpn.nexthop_self`](#nested-schema-for-address_familyl2vpn_evpnnexthop_self)
    - [Nested Schema for `address_family.l2vpn_evpn.route_map`](#nested-schema-for-address_familyl2vpn_evpnroute_map)
    - [Nested Schema for `address_family.l2vpn_evpn.soft_reconfiguration`](#nested-schema-for-address_familyl2vpn_evpnsoft_reconfiguration)
    - [Nested Schema for `bfd`](#nested-schema-for-bfd)
    - [Nested Schema for `capability`](#nested-schema-for-capability)
    - [Nested Schema for `path_attribute`](#nested-schema-for-path_attribute)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
    - [Nested Schema for `ttl_security`](#nested-schema-for-ttl_security)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### address_family
- `address_family` (Attributes) Address-family parameters (see [below for nested schema](#nestedatt--address_family))
#### bfd
- `bfd` (Attributes) Enable Bidirectional Forwarding Detection (BFD) support (see [below for nested schema](#nestedatt--bfd))
#### capability
- `capability` (Attributes) Advertise capabilities to this peer-group (see [below for nested schema](#nestedatt--capability))
#### description
- `description` (String) Description

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  txt     &emsp;|  Description  |
#### disable_capability_negotiation
- `disable_capability_negotiation` (Boolean) Disable capability negotiation with this neighbor
#### disable_connected_check
- `disable_connected_check` (Boolean) Allow peerings between eBGP peer using loopback/dummy address
#### ebgp_multihop
- `ebgp_multihop` (Number) Allow this EBGP neighbor to not be on a directly connected network

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  1-255   &emsp;|  Number of hops  |
#### graceful_restart
- `graceful_restart` (String) BGP graceful restart functionality

    |  Format          &emsp;|  Description                                            |
    |------------------|---------------------------------------------------------|
    |  enable          &emsp;|  Enable BGP graceful restart at peer level              |
    |  disable         &emsp;|  Disable BGP graceful restart at peer level             |
    |  restart-helper  &emsp;|  Enable BGP graceful restart helper only functionality  |
#### override_capability
- `override_capability` (Boolean) Ignore capability negotiation with specified neighbor
#### passive
- `passive` (Boolean) Do not initiate a session with this neighbor
#### password
- `password` (String) BGP MD5 password
#### path_attribute
- `path_attribute` (Attributes) Manipulate path attributes from incoming UPDATE messages (see [below for nested schema](#nestedatt--path_attribute))
#### port
- `port` (Number) Port number used by connection

    |  Format   &emsp;|  Description      |
    |-----------|-------------------|
    |  1-65535  &emsp;|  Numeric IP port  |
#### remote_as
- `remote_as` (String) Neighbor BGP AS number

    |  Format        &emsp;|  Description                         |
    |----------------|--------------------------------------|
    |  1-4294967294  &emsp;|  Neighbor AS number                  |
    |  external      &emsp;|  Any AS different from the local AS  |
    |  internal      &emsp;|  Neighbor AS number                  |
#### shutdown
- `shutdown` (Boolean) Administratively shutdown this neighbor
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
#### ttl_security
- `ttl_security` (Attributes) Ttl security mechanism (see [below for nested schema](#nestedatt--ttl_security))
#### update_source
- `update_source` (String) Source IP of routing updates

    |  Format  &emsp;|  Description                   |
    |----------|--------------------------------|
    |  ipv4    &emsp;|  IPv4 address of route source  |
    |  ipv6    &emsp;|  IPv6 address of route source  |
    |  txt     &emsp;|  Interface as route source     |

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `peer_group` (String) Name of peer-group


<a id="nestedatt--address_family"></a>
### Nested Schema for `address_family`

Optional:

- `ipv4_labeled_unicast` (Attributes) IPv4 Labeled Unicast BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast))
- `ipv4_unicast` (Attributes) IPv4 BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv4_unicast))
- `ipv4_vpn` (Attributes) IPv4 VPN BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv4_vpn))
- `ipv6_labeled_unicast` (Attributes) IPv6 Labeled Unicast BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast))
- `ipv6_unicast` (Attributes) IPv6 BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv6_unicast))
- `ipv6_vpn` (Attributes) IPv6 VPN BGP neighbor parameters (see [below for nested schema](#nestedatt--address_family--ipv6_vpn))
- `l2vpn_evpn` (Attributes) L2VPN EVPN BGP settings (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn))

<a id="nestedatt--address_family--ipv4_labeled_unicast"></a>
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

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--nexthop_self))
- `prefix_list` (Attributes) IPv4-Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `weight` (Number) Default weight for routes from this peer

    |  Format   &emsp;|  Description     |
    |-----------|------------------|
    |  1-65535  &emsp;|  Default weight  |

<a id="nestedatt--address_family--ipv4_labeled_unicast--allowas_in"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    |  Format  &emsp;|  Description                            |
    |----------|-----------------------------------------|
    |  1-10    &emsp;|  Number of times AS is allowed in path  |


<a id="nestedatt--address_family--ipv4_labeled_unicast--attribute_unchanged"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


<a id="nestedatt--address_family--ipv4_labeled_unicast--capability"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.capability`

Optional:

- `orf` (Attributes) Advertise ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--capability--orf))

<a id="nestedatt--address_family--ipv4_labeled_unicast--capability--orf"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.capability.orf`

Optional:

- `prefix_list` (Attributes) Advertise prefix-list ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_labeled_unicast--capability--orf--prefix_list))

<a id="nestedatt--address_family--ipv4_labeled_unicast--capability--orf--prefix_list"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.capability.orf.prefix_list`

Optional:

- `receive` (Boolean) Capability to receive the ORF
- `send` (Boolean) Capability to send the ORF




<a id="nestedatt--address_family--ipv4_labeled_unicast--conditionally_advertise"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv4_labeled_unicast--default_originate"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.default_originate`

Optional:

- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv4_labeled_unicast--disable_send_community"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


<a id="nestedatt--address_family--ipv4_labeled_unicast--distribute_list"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    |  Format   &emsp;|  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    |  Format   &emsp;|  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter incoming route updates from this peer-group  |


<a id="nestedatt--address_family--ipv4_labeled_unicast--filter_list"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


<a id="nestedatt--address_family--ipv4_labeled_unicast--nexthop_self"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


<a id="nestedatt--address_family--ipv4_labeled_unicast--prefix_list"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.prefix_list`

Optional:

- `export` (String) IPv4-Prefix-list to filter outgoing route updates to this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv4 prefix-list  |
- `import` (String) IPv4-Prefix-list to filter incoming route updates from this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv4 prefix-list  |


<a id="nestedatt--address_family--ipv4_labeled_unicast--remove_private_as"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


<a id="nestedatt--address_family--ipv4_labeled_unicast--route_map"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `import` (String) Route-map to filter incoming route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv4_labeled_unicast--soft_reconfiguration"></a>
### Nested Schema for `address_family.ipv4_labeled_unicast.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



<a id="nestedatt--address_family--ipv4_unicast"></a>
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

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--nexthop_self))
- `prefix_list` (Attributes) IPv4-Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `weight` (Number) Default weight for routes from this peer

    |  Format   &emsp;|  Description     |
    |-----------|------------------|
    |  1-65535  &emsp;|  Default weight  |

<a id="nestedatt--address_family--ipv4_unicast--allowas_in"></a>
### Nested Schema for `address_family.ipv4_unicast.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    |  Format  &emsp;|  Description                            |
    |----------|-----------------------------------------|
    |  1-10    &emsp;|  Number of times AS is allowed in path  |


<a id="nestedatt--address_family--ipv4_unicast--attribute_unchanged"></a>
### Nested Schema for `address_family.ipv4_unicast.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


<a id="nestedatt--address_family--ipv4_unicast--capability"></a>
### Nested Schema for `address_family.ipv4_unicast.capability`

Optional:

- `orf` (Attributes) Advertise ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--capability--orf))

<a id="nestedatt--address_family--ipv4_unicast--capability--orf"></a>
### Nested Schema for `address_family.ipv4_unicast.capability.orf`

Optional:

- `prefix_list` (Attributes) Advertise prefix-list ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv4_unicast--capability--orf--prefix_list))

<a id="nestedatt--address_family--ipv4_unicast--capability--orf--prefix_list"></a>
### Nested Schema for `address_family.ipv4_unicast.capability.orf.prefix_list`

Optional:

- `receive` (Boolean) Capability to receive the ORF
- `send` (Boolean) Capability to send the ORF




<a id="nestedatt--address_family--ipv4_unicast--conditionally_advertise"></a>
### Nested Schema for `address_family.ipv4_unicast.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv4_unicast--default_originate"></a>
### Nested Schema for `address_family.ipv4_unicast.default_originate`

Optional:

- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv4_unicast--disable_send_community"></a>
### Nested Schema for `address_family.ipv4_unicast.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


<a id="nestedatt--address_family--ipv4_unicast--distribute_list"></a>
### Nested Schema for `address_family.ipv4_unicast.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    |  Format   &emsp;|  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    |  Format   &emsp;|  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter incoming route updates from this peer-group  |


<a id="nestedatt--address_family--ipv4_unicast--filter_list"></a>
### Nested Schema for `address_family.ipv4_unicast.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


<a id="nestedatt--address_family--ipv4_unicast--nexthop_self"></a>
### Nested Schema for `address_family.ipv4_unicast.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


<a id="nestedatt--address_family--ipv4_unicast--prefix_list"></a>
### Nested Schema for `address_family.ipv4_unicast.prefix_list`

Optional:

- `export` (String) IPv4-Prefix-list to filter outgoing route updates to this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv4 prefix-list  |
- `import` (String) IPv4-Prefix-list to filter incoming route updates from this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv4 prefix-list  |


<a id="nestedatt--address_family--ipv4_unicast--remove_private_as"></a>
### Nested Schema for `address_family.ipv4_unicast.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


<a id="nestedatt--address_family--ipv4_unicast--route_map"></a>
### Nested Schema for `address_family.ipv4_unicast.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `import` (String) Route-map to filter incoming route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv4_unicast--soft_reconfiguration"></a>
### Nested Schema for `address_family.ipv4_unicast.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



<a id="nestedatt--address_family--ipv4_vpn"></a>
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

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--nexthop_self))
- `prefix_list` (Attributes) IPv4-Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv4_vpn--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `weight` (Number) Default weight for routes from this peer

    |  Format   &emsp;|  Description     |
    |-----------|------------------|
    |  1-65535  &emsp;|  Default weight  |

<a id="nestedatt--address_family--ipv4_vpn--allowas_in"></a>
### Nested Schema for `address_family.ipv4_vpn.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    |  Format  &emsp;|  Description                            |
    |----------|-----------------------------------------|
    |  1-10    &emsp;|  Number of times AS is allowed in path  |


<a id="nestedatt--address_family--ipv4_vpn--attribute_unchanged"></a>
### Nested Schema for `address_family.ipv4_vpn.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


<a id="nestedatt--address_family--ipv4_vpn--conditionally_advertise"></a>
### Nested Schema for `address_family.ipv4_vpn.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv4_vpn--disable_send_community"></a>
### Nested Schema for `address_family.ipv4_vpn.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


<a id="nestedatt--address_family--ipv4_vpn--distribute_list"></a>
### Nested Schema for `address_family.ipv4_vpn.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    |  Format   &emsp;|  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    |  Format   &emsp;|  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter incoming route updates from this peer-group  |


<a id="nestedatt--address_family--ipv4_vpn--filter_list"></a>
### Nested Schema for `address_family.ipv4_vpn.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


<a id="nestedatt--address_family--ipv4_vpn--nexthop_self"></a>
### Nested Schema for `address_family.ipv4_vpn.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


<a id="nestedatt--address_family--ipv4_vpn--prefix_list"></a>
### Nested Schema for `address_family.ipv4_vpn.prefix_list`

Optional:

- `export` (String) IPv4-Prefix-list to filter outgoing route updates to this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv4 prefix-list  |
- `import` (String) IPv4-Prefix-list to filter incoming route updates from this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv4 prefix-list  |


<a id="nestedatt--address_family--ipv4_vpn--remove_private_as"></a>
### Nested Schema for `address_family.ipv4_vpn.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


<a id="nestedatt--address_family--ipv4_vpn--route_map"></a>
### Nested Schema for `address_family.ipv4_vpn.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `import` (String) Route-map to filter incoming route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv4_vpn--soft_reconfiguration"></a>
### Nested Schema for `address_family.ipv4_vpn.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



<a id="nestedatt--address_family--ipv6_labeled_unicast"></a>
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

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `nexthop_local` (Attributes) Nexthop attributes (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--nexthop_local))
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--nexthop_self))
- `prefix_list` (Attributes) Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `weight` (Number) Default weight for routes from this peer

    |  Format   &emsp;|  Description     |
    |-----------|------------------|
    |  1-65535  &emsp;|  Default weight  |

<a id="nestedatt--address_family--ipv6_labeled_unicast--allowas_in"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    |  Format  &emsp;|  Description                            |
    |----------|-----------------------------------------|
    |  1-10    &emsp;|  Number of times AS is allowed in path  |


<a id="nestedatt--address_family--ipv6_labeled_unicast--attribute_unchanged"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


<a id="nestedatt--address_family--ipv6_labeled_unicast--capability"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.capability`

Optional:

- `orf` (Attributes) Advertise ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--capability--orf))

<a id="nestedatt--address_family--ipv6_labeled_unicast--capability--orf"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.capability.orf`

Optional:

- `prefix_list` (Attributes) Advertise prefix-list ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_labeled_unicast--capability--orf--prefix_list))

<a id="nestedatt--address_family--ipv6_labeled_unicast--capability--orf--prefix_list"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.capability.orf.prefix_list`

Optional:

- `receive` (Boolean) Capability to receive the ORF
- `send` (Boolean) Capability to send the ORF




<a id="nestedatt--address_family--ipv6_labeled_unicast--conditionally_advertise"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv6_labeled_unicast--default_originate"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.default_originate`

Optional:

- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv6_labeled_unicast--disable_send_community"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


<a id="nestedatt--address_family--ipv6_labeled_unicast--distribute_list"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    |  Format   &emsp;|  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    |  Format   &emsp;|  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter incoming route updates from this peer-group  |


<a id="nestedatt--address_family--ipv6_labeled_unicast--filter_list"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


<a id="nestedatt--address_family--ipv6_labeled_unicast--nexthop_local"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.nexthop_local`

Optional:

- `unchanged` (Boolean) Leave link-local nexthop unchanged for this peer


<a id="nestedatt--address_family--ipv6_labeled_unicast--nexthop_self"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


<a id="nestedatt--address_family--ipv6_labeled_unicast--prefix_list"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.prefix_list`

Optional:

- `export` (String) Prefix-list to filter outgoing route updates to this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv6 prefix-list  |
- `import` (String) Prefix-list to filter incoming route updates from this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv6 prefix-list  |


<a id="nestedatt--address_family--ipv6_labeled_unicast--remove_private_as"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


<a id="nestedatt--address_family--ipv6_labeled_unicast--route_map"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `import` (String) Route-map to filter incoming route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv6_labeled_unicast--soft_reconfiguration"></a>
### Nested Schema for `address_family.ipv6_labeled_unicast.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



<a id="nestedatt--address_family--ipv6_unicast"></a>
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

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `nexthop_local` (Attributes) Nexthop attributes (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--nexthop_local))
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--nexthop_self))
- `prefix_list` (Attributes) Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `weight` (Number) Default weight for routes from this peer

    |  Format   &emsp;|  Description     |
    |-----------|------------------|
    |  1-65535  &emsp;|  Default weight  |

<a id="nestedatt--address_family--ipv6_unicast--allowas_in"></a>
### Nested Schema for `address_family.ipv6_unicast.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    |  Format  &emsp;|  Description                            |
    |----------|-----------------------------------------|
    |  1-10    &emsp;|  Number of times AS is allowed in path  |


<a id="nestedatt--address_family--ipv6_unicast--attribute_unchanged"></a>
### Nested Schema for `address_family.ipv6_unicast.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


<a id="nestedatt--address_family--ipv6_unicast--capability"></a>
### Nested Schema for `address_family.ipv6_unicast.capability`

Optional:

- `orf` (Attributes) Advertise ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--capability--orf))

<a id="nestedatt--address_family--ipv6_unicast--capability--orf"></a>
### Nested Schema for `address_family.ipv6_unicast.capability.orf`

Optional:

- `prefix_list` (Attributes) Advertise prefix-list ORF capability to this peer (see [below for nested schema](#nestedatt--address_family--ipv6_unicast--capability--orf--prefix_list))

<a id="nestedatt--address_family--ipv6_unicast--capability--orf--prefix_list"></a>
### Nested Schema for `address_family.ipv6_unicast.capability.orf.prefix_list`

Optional:

- `receive` (Boolean) Capability to receive the ORF
- `send` (Boolean) Capability to send the ORF




<a id="nestedatt--address_family--ipv6_unicast--conditionally_advertise"></a>
### Nested Schema for `address_family.ipv6_unicast.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv6_unicast--default_originate"></a>
### Nested Schema for `address_family.ipv6_unicast.default_originate`

Optional:

- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv6_unicast--disable_send_community"></a>
### Nested Schema for `address_family.ipv6_unicast.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


<a id="nestedatt--address_family--ipv6_unicast--distribute_list"></a>
### Nested Schema for `address_family.ipv6_unicast.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    |  Format   &emsp;|  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    |  Format   &emsp;|  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter incoming route updates from this peer-group  |


<a id="nestedatt--address_family--ipv6_unicast--filter_list"></a>
### Nested Schema for `address_family.ipv6_unicast.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


<a id="nestedatt--address_family--ipv6_unicast--nexthop_local"></a>
### Nested Schema for `address_family.ipv6_unicast.nexthop_local`

Optional:

- `unchanged` (Boolean) Leave link-local nexthop unchanged for this peer


<a id="nestedatt--address_family--ipv6_unicast--nexthop_self"></a>
### Nested Schema for `address_family.ipv6_unicast.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


<a id="nestedatt--address_family--ipv6_unicast--prefix_list"></a>
### Nested Schema for `address_family.ipv6_unicast.prefix_list`

Optional:

- `export` (String) Prefix-list to filter outgoing route updates to this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv6 prefix-list  |
- `import` (String) Prefix-list to filter incoming route updates from this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv6 prefix-list  |


<a id="nestedatt--address_family--ipv6_unicast--remove_private_as"></a>
### Nested Schema for `address_family.ipv6_unicast.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


<a id="nestedatt--address_family--ipv6_unicast--route_map"></a>
### Nested Schema for `address_family.ipv6_unicast.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `import` (String) Route-map to filter incoming route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv6_unicast--soft_reconfiguration"></a>
### Nested Schema for `address_family.ipv6_unicast.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



<a id="nestedatt--address_family--ipv6_vpn"></a>
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

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `maximum_prefix_out` (Number) Maximum number of prefixes to be sent to this peer

    |  Format        &emsp;|  Description   |
    |----------------|----------------|
    |  1-4294967295  &emsp;|  Prefix limit  |
- `nexthop_local` (Attributes) Nexthop attributes (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--nexthop_local))
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--nexthop_self))
- `prefix_list` (Attributes) Prefix-list to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--prefix_list))
- `remove_private_as` (Attributes) Remove private AS numbers from AS path in outbound route updates (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--remove_private_as))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--ipv6_vpn--soft_reconfiguration))
- `unsuppress_map` (String) Route-map to selectively unsuppress suppressed routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `weight` (Number) Default weight for routes from this peer

    |  Format   &emsp;|  Description     |
    |-----------|------------------|
    |  1-65535  &emsp;|  Default weight  |

<a id="nestedatt--address_family--ipv6_vpn--allowas_in"></a>
### Nested Schema for `address_family.ipv6_vpn.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    |  Format  &emsp;|  Description                            |
    |----------|-----------------------------------------|
    |  1-10    &emsp;|  Number of times AS is allowed in path  |


<a id="nestedatt--address_family--ipv6_vpn--attribute_unchanged"></a>
### Nested Schema for `address_family.ipv6_vpn.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


<a id="nestedatt--address_family--ipv6_vpn--conditionally_advertise"></a>
### Nested Schema for `address_family.ipv6_vpn.conditionally_advertise`

Optional:

- `advertise_map` (String) Route-map to conditionally advertise routes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `exist_map` (String) Advertise routes only if prefixes in exist-map are installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `non_exist_map` (String) Advertise routes only if prefixes in non-exist-map are not installed in BGP table

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv6_vpn--disable_send_community"></a>
### Nested Schema for `address_family.ipv6_vpn.disable_send_community`

Optional:

- `extended` (Boolean) Disable sending extended community attributes to this peer
- `standard` (Boolean) Disable sending standard community attributes to this peer


<a id="nestedatt--address_family--ipv6_vpn--distribute_list"></a>
### Nested Schema for `address_family.ipv6_vpn.distribute_list`

Optional:

- `export` (Number) Access-list to filter outgoing route updates to this peer-group

    |  Format   &emsp;|  Description                                                      |
    |-----------|-------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter outgoing route updates to this peer-group  |
- `import` (Number) Access-list to filter incoming route updates from this peer-group

    |  Format   &emsp;|  Description                                                        |
    |-----------|---------------------------------------------------------------------|
    |  1-65535  &emsp;|  Access-list to filter incoming route updates from this peer-group  |


<a id="nestedatt--address_family--ipv6_vpn--filter_list"></a>
### Nested Schema for `address_family.ipv6_vpn.filter_list`

Optional:

- `export` (String) As-path-list to filter outgoing route updates to this peer
- `import` (String) As-path-list to filter incoming route updates from this peer


<a id="nestedatt--address_family--ipv6_vpn--nexthop_local"></a>
### Nested Schema for `address_family.ipv6_vpn.nexthop_local`

Optional:

- `unchanged` (Boolean) Leave link-local nexthop unchanged for this peer


<a id="nestedatt--address_family--ipv6_vpn--nexthop_self"></a>
### Nested Schema for `address_family.ipv6_vpn.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


<a id="nestedatt--address_family--ipv6_vpn--prefix_list"></a>
### Nested Schema for `address_family.ipv6_vpn.prefix_list`

Optional:

- `export` (String) Prefix-list to filter outgoing route updates to this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv6 prefix-list  |
- `import` (String) Prefix-list to filter incoming route updates from this peer

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  txt     &emsp;|  Name of IPv6 prefix-list  |


<a id="nestedatt--address_family--ipv6_vpn--remove_private_as"></a>
### Nested Schema for `address_family.ipv6_vpn.remove_private_as`

Optional:

- `all` (Boolean) Remove private AS numbers to all AS numbers in outbound route updates


<a id="nestedatt--address_family--ipv6_vpn--route_map"></a>
### Nested Schema for `address_family.ipv6_vpn.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `import` (String) Route-map to filter incoming route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--ipv6_vpn--soft_reconfiguration"></a>
### Nested Schema for `address_family.ipv6_vpn.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration



<a id="nestedatt--address_family--l2vpn_evpn"></a>
### Nested Schema for `address_family.l2vpn_evpn`

Optional:

- `allowas_in` (Attributes) Accept route that contains the local-as in the as-path (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn--allowas_in))
- `attribute_unchanged` (Attributes) BGP attributes are sent unchanged (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn--attribute_unchanged))
- `nexthop_self` (Attributes) Disable the next hop calculation for this peer (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn--nexthop_self))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn--route_map))
- `route_reflector_client` (Boolean) Peer is a route reflector client
- `route_server_client` (Boolean) Peer is a route server client
- `soft_reconfiguration` (Attributes) Soft reconfiguration for peer (see [below for nested schema](#nestedatt--address_family--l2vpn_evpn--soft_reconfiguration))

<a id="nestedatt--address_family--l2vpn_evpn--allowas_in"></a>
### Nested Schema for `address_family.l2vpn_evpn.allowas_in`

Optional:

- `number` (Number) Number of occurrences of AS number

    |  Format  &emsp;|  Description                            |
    |----------|-----------------------------------------|
    |  1-10    &emsp;|  Number of times AS is allowed in path  |


<a id="nestedatt--address_family--l2vpn_evpn--attribute_unchanged"></a>
### Nested Schema for `address_family.l2vpn_evpn.attribute_unchanged`

Optional:

- `as_path` (Boolean) Send AS path unchanged
- `med` (Boolean) Send multi-exit discriminator unchanged
- `next_hop` (Boolean) Send nexthop unchanged


<a id="nestedatt--address_family--l2vpn_evpn--nexthop_self"></a>
### Nested Schema for `address_family.l2vpn_evpn.nexthop_self`

Optional:

- `force` (Boolean) Set the next hop to self for reflected routes


<a id="nestedatt--address_family--l2vpn_evpn--route_map"></a>
### Nested Schema for `address_family.l2vpn_evpn.route_map`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `import` (String) Route-map to filter incoming route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--address_family--l2vpn_evpn--soft_reconfiguration"></a>
### Nested Schema for `address_family.l2vpn_evpn.soft_reconfiguration`

Optional:

- `inbound` (Boolean) Enable inbound soft reconfiguration




<a id="nestedatt--bfd"></a>
### Nested Schema for `bfd`

Optional:

- `check_control_plane_failure` (Boolean) Allow to write CBIT independence in BFD outgoing packets and read both C-BIT value of BFD and lookup BGP peer status
- `profile` (String) Use settings from BFD profile

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  txt     &emsp;|  BFD profile name  |


<a id="nestedatt--capability"></a>
### Nested Schema for `capability`

Optional:

- `dynamic` (Boolean) Advertise dynamic capability to this neighbor
- `extended_nexthop` (Boolean) Advertise extended-nexthop capability to this neighbor
- `software_version` (Boolean) Advertise Software Version capability to the peer


<a id="nestedatt--path_attribute"></a>
### Nested Schema for `path_attribute`

Optional:

- `discard` (List of Number) Drop specified attributes from incoming UPDATE messages

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  1-255   &emsp;|  Attribute number  |
- `treat_as_withdraw` (Number) Treat-as-withdraw any incoming BGP UPDATE messages that contain the specified attribute

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  1-255   &emsp;|  Attribute number  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


<a id="nestedatt--ttl_security"></a>
### Nested Schema for `ttl_security`

Optional:

- `hops` (Number) Number of the maximum number of hops to the BGP peer

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  1-254   &emsp;|  Number of hops  |
