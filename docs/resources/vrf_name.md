---
page_title: "vyos_vrf_name Resource - vyos"

subcategory: "Vrf"

description: |-
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance
---

# vyos_vrf_name (Resource)
<center>


Virtual Routing and Forwarding  
⯯  
**Virtual Routing and Forwarding instance**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_vrf_name (Resource)](#vyos_vrf_name-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [description](#description)
      - [disable](#disable)
      - [ip](#ip)
      - [ipv6](#ipv6)
      - [protocols](#protocols)
      - [table](#table)
      - [timeouts](#timeouts)
      - [vni](#vni)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `ip`](#nested-schema-for-ip)
    - [Nested Schema for `ip.nht`](#nested-schema-for-ipnht)
    - [Nested Schema for `ipv6`](#nested-schema-for-ipv6)
    - [Nested Schema for `ipv6.nht`](#nested-schema-for-ipv6nht)
    - [Nested Schema for `protocols`](#nested-schema-for-protocols)
    - [Nested Schema for `protocols.bgp`](#nested-schema-for-protocolsbgp)
    - [Nested Schema for `protocols.bgp.address_family`](#nested-schema-for-protocolsbgpaddress_family)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_flowspec`](#nested-schema-for-protocolsbgpaddress_familyipv4_flowspec)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_flowspec.local_install`](#nested-schema-for-protocolsbgpaddress_familyipv4_flowspeclocal_install)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_labeled_unicast`](#nested-schema-for-protocolsbgpaddress_familyipv4_labeled_unicast)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_labeled_unicast.maximum_paths`](#nested-schema-for-protocolsbgpaddress_familyipv4_labeled_unicastmaximum_paths)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_multicast`](#nested-schema-for-protocolsbgpaddress_familyipv4_multicast)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_multicast.distance`](#nested-schema-for-protocolsbgpaddress_familyipv4_multicastdistance)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicast)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.distance`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastdistance)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.export`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastexport)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.import`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastimport)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.label`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastlabel)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.label.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastlabelvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.label.vpn.allocation_mode`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastlabelvpnallocation_mode)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.maximum_paths`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastmaximum_paths)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.nexthop`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastnexthop)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.nexthop.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastnexthopvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.rd`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastrd)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.rd.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastrdvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastredistribute)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.babel`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastredistributebabel)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.connected`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastredistributeconnected)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.isis`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastredistributeisis)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.kernel`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastredistributekernel)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.ospf`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastredistributeospf)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.rip`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastredistributerip)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.static`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastredistributestatic)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_map`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastroute_map)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_map.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastroute_mapvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_target`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastroute_target)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_target.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastroute_targetvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.sid`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastsid)
    - [Nested Schema for `protocols.bgp.address_family.ipv4_unicast.sid.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv4_unicastsidvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_flowspec`](#nested-schema-for-protocolsbgpaddress_familyipv6_flowspec)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_flowspec.local_install`](#nested-schema-for-protocolsbgpaddress_familyipv6_flowspeclocal_install)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_multicast`](#nested-schema-for-protocolsbgpaddress_familyipv6_multicast)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_multicast.distance`](#nested-schema-for-protocolsbgpaddress_familyipv6_multicastdistance)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicast)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.distance`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastdistance)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.export`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastexport)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.import`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastimport)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.label`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastlabel)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.label.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastlabelvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.label.vpn.allocation_mode`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastlabelvpnallocation_mode)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.maximum_paths`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastmaximum_paths)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.nexthop`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastnexthop)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.nexthop.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastnexthopvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.rd`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastrd)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.rd.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastrdvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastredistribute)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.babel`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastredistributebabel)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.connected`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastredistributeconnected)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.isis`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastredistributeisis)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.kernel`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastredistributekernel)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.ospfv3`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastredistributeospfv3)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.ripng`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastredistributeripng)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.static`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastredistributestatic)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_map`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastroute_map)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_map.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastroute_mapvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_target`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastroute_target)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_target.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastroute_targetvpn)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.sid`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastsid)
    - [Nested Schema for `protocols.bgp.address_family.ipv6_unicast.sid.vpn`](#nested-schema-for-protocolsbgpaddress_familyipv6_unicastsidvpn)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpn)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpnadvertise)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv4`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpnadvertiseipv4)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv4.unicast`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpnadvertiseipv4unicast)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv6`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpnadvertiseipv6)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv6.unicast`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpnadvertiseipv6unicast)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.default_originate`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpndefault_originate)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.ead_es_frag`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpnead_es_frag)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.ead_es_route_target`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpnead_es_route_target)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.flooding`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpnflooding)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.mac_vrf`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpnmac_vrf)
    - [Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.route_target`](#nested-schema-for-protocolsbgpaddress_familyl2vpn_evpnroute_target)
    - [Nested Schema for `protocols.bgp.bmp`](#nested-schema-for-protocolsbgpbmp)
    - [Nested Schema for `protocols.bgp.listen`](#nested-schema-for-protocolsbgplisten)
    - [Nested Schema for `protocols.bgp.parameters`](#nested-schema-for-protocolsbgpparameters)
    - [Nested Schema for `protocols.bgp.parameters.bestpath`](#nested-schema-for-protocolsbgpparametersbestpath)
    - [Nested Schema for `protocols.bgp.parameters.bestpath.as_path`](#nested-schema-for-protocolsbgpparametersbestpathas_path)
    - [Nested Schema for `protocols.bgp.parameters.bestpath.peer_type`](#nested-schema-for-protocolsbgpparametersbestpathpeer_type)
    - [Nested Schema for `protocols.bgp.parameters.conditional_advertisement`](#nested-schema-for-protocolsbgpparametersconditional_advertisement)
    - [Nested Schema for `protocols.bgp.parameters.confederation`](#nested-schema-for-protocolsbgpparametersconfederation)
    - [Nested Schema for `protocols.bgp.parameters.dampening`](#nested-schema-for-protocolsbgpparametersdampening)
    - [Nested Schema for `protocols.bgp.parameters.default`](#nested-schema-for-protocolsbgpparametersdefault)
    - [Nested Schema for `protocols.bgp.parameters.distance`](#nested-schema-for-protocolsbgpparametersdistance)
    - [Nested Schema for `protocols.bgp.parameters.distance.global`](#nested-schema-for-protocolsbgpparametersdistanceglobal)
    - [Nested Schema for `protocols.bgp.parameters.graceful_restart`](#nested-schema-for-protocolsbgpparametersgraceful_restart)
    - [Nested Schema for `protocols.bgp.parameters.tcp_keepalive`](#nested-schema-for-protocolsbgpparameterstcp_keepalive)
    - [Nested Schema for `protocols.bgp.sid`](#nested-schema-for-protocolsbgpsid)
    - [Nested Schema for `protocols.bgp.sid.vpn`](#nested-schema-for-protocolsbgpsidvpn)
    - [Nested Schema for `protocols.bgp.sid.vpn.per_vrf`](#nested-schema-for-protocolsbgpsidvpnper_vrf)
    - [Nested Schema for `protocols.bgp.srv6`](#nested-schema-for-protocolsbgpsrv6)
    - [Nested Schema for `protocols.bgp.timers`](#nested-schema-for-protocolsbgptimers)
    - [Nested Schema for `protocols.eigrp`](#nested-schema-for-protocolseigrp)
    - [Nested Schema for `protocols.eigrp.metric`](#nested-schema-for-protocolseigrpmetric)
    - [Nested Schema for `protocols.isis`](#nested-schema-for-protocolsisis)
    - [Nested Schema for `protocols.isis.area_password`](#nested-schema-for-protocolsisisarea_password)
    - [Nested Schema for `protocols.isis.default_information`](#nested-schema-for-protocolsisisdefault_information)
    - [Nested Schema for `protocols.isis.default_information.originate`](#nested-schema-for-protocolsisisdefault_informationoriginate)
    - [Nested Schema for `protocols.isis.default_information.originate.ipv4`](#nested-schema-for-protocolsisisdefault_informationoriginateipv4)
    - [Nested Schema for `protocols.isis.default_information.originate.ipv4.level_1`](#nested-schema-for-protocolsisisdefault_informationoriginateipv4level_1)
    - [Nested Schema for `protocols.isis.default_information.originate.ipv4.level_2`](#nested-schema-for-protocolsisisdefault_informationoriginateipv4level_2)
    - [Nested Schema for `protocols.isis.default_information.originate.ipv6`](#nested-schema-for-protocolsisisdefault_informationoriginateipv6)
    - [Nested Schema for `protocols.isis.default_information.originate.ipv6.level_1`](#nested-schema-for-protocolsisisdefault_informationoriginateipv6level_1)
    - [Nested Schema for `protocols.isis.default_information.originate.ipv6.level_2`](#nested-schema-for-protocolsisisdefault_informationoriginateipv6level_2)
    - [Nested Schema for `protocols.isis.domain_password`](#nested-schema-for-protocolsisisdomain_password)
    - [Nested Schema for `protocols.isis.fast_reroute`](#nested-schema-for-protocolsisisfast_reroute)
    - [Nested Schema for `protocols.isis.fast_reroute.lfa`](#nested-schema-for-protocolsisisfast_reroutelfa)
    - [Nested Schema for `protocols.isis.fast_reroute.lfa.local`](#nested-schema-for-protocolsisisfast_reroutelfalocal)
    - [Nested Schema for `protocols.isis.fast_reroute.lfa.local.load_sharing`](#nested-schema-for-protocolsisisfast_reroutelfalocalload_sharing)
    - [Nested Schema for `protocols.isis.fast_reroute.lfa.local.load_sharing.disable`](#nested-schema-for-protocolsisisfast_reroutelfalocalload_sharingdisable)
    - [Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit`](#nested-schema-for-protocolsisisfast_reroutelfalocalpriority_limit)
    - [Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit.critical`](#nested-schema-for-protocolsisisfast_reroutelfalocalpriority_limitcritical)
    - [Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit.high`](#nested-schema-for-protocolsisisfast_reroutelfalocalpriority_limithigh)
    - [Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit.medium`](#nested-schema-for-protocolsisisfast_reroutelfalocalpriority_limitmedium)
    - [Nested Schema for `protocols.isis.ldp_sync`](#nested-schema-for-protocolsisisldp_sync)
    - [Nested Schema for `protocols.isis.redistribute`](#nested-schema-for-protocolsisisredistribute)
    - [Nested Schema for `protocols.isis.redistribute.ipv4`](#nested-schema-for-protocolsisisredistributeipv4)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.babel`](#nested-schema-for-protocolsisisredistributeipv4babel)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.babel.level_1`](#nested-schema-for-protocolsisisredistributeipv4babellevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.babel.level_2`](#nested-schema-for-protocolsisisredistributeipv4babellevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.bgp`](#nested-schema-for-protocolsisisredistributeipv4bgp)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.bgp.level_1`](#nested-schema-for-protocolsisisredistributeipv4bgplevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.bgp.level_2`](#nested-schema-for-protocolsisisredistributeipv4bgplevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.connected`](#nested-schema-for-protocolsisisredistributeipv4connected)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.connected.level_1`](#nested-schema-for-protocolsisisredistributeipv4connectedlevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.connected.level_2`](#nested-schema-for-protocolsisisredistributeipv4connectedlevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.kernel`](#nested-schema-for-protocolsisisredistributeipv4kernel)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.kernel.level_1`](#nested-schema-for-protocolsisisredistributeipv4kernellevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.kernel.level_2`](#nested-schema-for-protocolsisisredistributeipv4kernellevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.ospf`](#nested-schema-for-protocolsisisredistributeipv4ospf)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.ospf.level_1`](#nested-schema-for-protocolsisisredistributeipv4ospflevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.ospf.level_2`](#nested-schema-for-protocolsisisredistributeipv4ospflevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.rip`](#nested-schema-for-protocolsisisredistributeipv4rip)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.rip.level_1`](#nested-schema-for-protocolsisisredistributeipv4riplevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.rip.level_2`](#nested-schema-for-protocolsisisredistributeipv4riplevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.static`](#nested-schema-for-protocolsisisredistributeipv4static)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.static.level_1`](#nested-schema-for-protocolsisisredistributeipv4staticlevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv4.static.level_2`](#nested-schema-for-protocolsisisredistributeipv4staticlevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv6`](#nested-schema-for-protocolsisisredistributeipv6)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.babel`](#nested-schema-for-protocolsisisredistributeipv6babel)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.babel.level_1`](#nested-schema-for-protocolsisisredistributeipv6babellevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.babel.level_2`](#nested-schema-for-protocolsisisredistributeipv6babellevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.bgp`](#nested-schema-for-protocolsisisredistributeipv6bgp)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.bgp.level_1`](#nested-schema-for-protocolsisisredistributeipv6bgplevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.bgp.level_2`](#nested-schema-for-protocolsisisredistributeipv6bgplevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.connected`](#nested-schema-for-protocolsisisredistributeipv6connected)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.connected.level_1`](#nested-schema-for-protocolsisisredistributeipv6connectedlevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.connected.level_2`](#nested-schema-for-protocolsisisredistributeipv6connectedlevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.kernel`](#nested-schema-for-protocolsisisredistributeipv6kernel)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.kernel.level_1`](#nested-schema-for-protocolsisisredistributeipv6kernellevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.kernel.level_2`](#nested-schema-for-protocolsisisredistributeipv6kernellevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.ospf6`](#nested-schema-for-protocolsisisredistributeipv6ospf6)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.ospf6.level_1`](#nested-schema-for-protocolsisisredistributeipv6ospf6level_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.ospf6.level_2`](#nested-schema-for-protocolsisisredistributeipv6ospf6level_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.ripng`](#nested-schema-for-protocolsisisredistributeipv6ripng)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.ripng.level_1`](#nested-schema-for-protocolsisisredistributeipv6ripnglevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.ripng.level_2`](#nested-schema-for-protocolsisisredistributeipv6ripnglevel_2)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.static`](#nested-schema-for-protocolsisisredistributeipv6static)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.static.level_1`](#nested-schema-for-protocolsisisredistributeipv6staticlevel_1)
    - [Nested Schema for `protocols.isis.redistribute.ipv6.static.level_2`](#nested-schema-for-protocolsisisredistributeipv6staticlevel_2)
    - [Nested Schema for `protocols.isis.segment_routing`](#nested-schema-for-protocolsisissegment_routing)
    - [Nested Schema for `protocols.isis.segment_routing.global_block`](#nested-schema-for-protocolsisissegment_routingglobal_block)
    - [Nested Schema for `protocols.isis.segment_routing.local_block`](#nested-schema-for-protocolsisissegment_routinglocal_block)
    - [Nested Schema for `protocols.isis.spf_delay_ietf`](#nested-schema-for-protocolsisisspf_delay_ietf)
    - [Nested Schema for `protocols.isis.traffic_engineering`](#nested-schema-for-protocolsisistraffic_engineering)
    - [Nested Schema for `protocols.ospf`](#nested-schema-for-protocolsospf)
    - [Nested Schema for `protocols.ospf.aggregation`](#nested-schema-for-protocolsospfaggregation)
    - [Nested Schema for `protocols.ospf.auto_cost`](#nested-schema-for-protocolsospfauto_cost)
    - [Nested Schema for `protocols.ospf.capability`](#nested-schema-for-protocolsospfcapability)
    - [Nested Schema for `protocols.ospf.default_information`](#nested-schema-for-protocolsospfdefault_information)
    - [Nested Schema for `protocols.ospf.default_information.originate`](#nested-schema-for-protocolsospfdefault_informationoriginate)
    - [Nested Schema for `protocols.ospf.distance`](#nested-schema-for-protocolsospfdistance)
    - [Nested Schema for `protocols.ospf.distance.ospf`](#nested-schema-for-protocolsospfdistanceospf)
    - [Nested Schema for `protocols.ospf.graceful_restart`](#nested-schema-for-protocolsospfgraceful_restart)
    - [Nested Schema for `protocols.ospf.graceful_restart.helper`](#nested-schema-for-protocolsospfgraceful_restarthelper)
    - [Nested Schema for `protocols.ospf.graceful_restart.helper.enable`](#nested-schema-for-protocolsospfgraceful_restarthelperenable)
    - [Nested Schema for `protocols.ospf.ldp_sync`](#nested-schema-for-protocolsospfldp_sync)
    - [Nested Schema for `protocols.ospf.log_adjacency_changes`](#nested-schema-for-protocolsospflog_adjacency_changes)
    - [Nested Schema for `protocols.ospf.max_metric`](#nested-schema-for-protocolsospfmax_metric)
    - [Nested Schema for `protocols.ospf.max_metric.router_lsa`](#nested-schema-for-protocolsospfmax_metricrouter_lsa)
    - [Nested Schema for `protocols.ospf.mpls_te`](#nested-schema-for-protocolsospfmpls_te)
    - [Nested Schema for `protocols.ospf.parameters`](#nested-schema-for-protocolsospfparameters)
    - [Nested Schema for `protocols.ospf.redistribute`](#nested-schema-for-protocolsospfredistribute)
    - [Nested Schema for `protocols.ospf.redistribute.babel`](#nested-schema-for-protocolsospfredistributebabel)
    - [Nested Schema for `protocols.ospf.redistribute.bgp`](#nested-schema-for-protocolsospfredistributebgp)
    - [Nested Schema for `protocols.ospf.redistribute.connected`](#nested-schema-for-protocolsospfredistributeconnected)
    - [Nested Schema for `protocols.ospf.redistribute.isis`](#nested-schema-for-protocolsospfredistributeisis)
    - [Nested Schema for `protocols.ospf.redistribute.kernel`](#nested-schema-for-protocolsospfredistributekernel)
    - [Nested Schema for `protocols.ospf.redistribute.rip`](#nested-schema-for-protocolsospfredistributerip)
    - [Nested Schema for `protocols.ospf.redistribute.static`](#nested-schema-for-protocolsospfredistributestatic)
    - [Nested Schema for `protocols.ospf.refresh`](#nested-schema-for-protocolsospfrefresh)
    - [Nested Schema for `protocols.ospf.segment_routing`](#nested-schema-for-protocolsospfsegment_routing)
    - [Nested Schema for `protocols.ospf.segment_routing.global_block`](#nested-schema-for-protocolsospfsegment_routingglobal_block)
    - [Nested Schema for `protocols.ospf.segment_routing.local_block`](#nested-schema-for-protocolsospfsegment_routinglocal_block)
    - [Nested Schema for `protocols.ospf.timers`](#nested-schema-for-protocolsospftimers)
    - [Nested Schema for `protocols.ospf.timers.throttle`](#nested-schema-for-protocolsospftimersthrottle)
    - [Nested Schema for `protocols.ospf.timers.throttle.spf`](#nested-schema-for-protocolsospftimersthrottlespf)
    - [Nested Schema for `protocols.ospfv3`](#nested-schema-for-protocolsospfv3)
    - [Nested Schema for `protocols.ospfv3.auto_cost`](#nested-schema-for-protocolsospfv3auto_cost)
    - [Nested Schema for `protocols.ospfv3.default_information`](#nested-schema-for-protocolsospfv3default_information)
    - [Nested Schema for `protocols.ospfv3.default_information.originate`](#nested-schema-for-protocolsospfv3default_informationoriginate)
    - [Nested Schema for `protocols.ospfv3.distance`](#nested-schema-for-protocolsospfv3distance)
    - [Nested Schema for `protocols.ospfv3.distance.ospfv3`](#nested-schema-for-protocolsospfv3distanceospfv3)
    - [Nested Schema for `protocols.ospfv3.graceful_restart`](#nested-schema-for-protocolsospfv3graceful_restart)
    - [Nested Schema for `protocols.ospfv3.graceful_restart.helper`](#nested-schema-for-protocolsospfv3graceful_restarthelper)
    - [Nested Schema for `protocols.ospfv3.graceful_restart.helper.enable`](#nested-schema-for-protocolsospfv3graceful_restarthelperenable)
    - [Nested Schema for `protocols.ospfv3.log_adjacency_changes`](#nested-schema-for-protocolsospfv3log_adjacency_changes)
    - [Nested Schema for `protocols.ospfv3.parameters`](#nested-schema-for-protocolsospfv3parameters)
    - [Nested Schema for `protocols.ospfv3.redistribute`](#nested-schema-for-protocolsospfv3redistribute)
    - [Nested Schema for `protocols.ospfv3.redistribute.babel`](#nested-schema-for-protocolsospfv3redistributebabel)
    - [Nested Schema for `protocols.ospfv3.redistribute.bgp`](#nested-schema-for-protocolsospfv3redistributebgp)
    - [Nested Schema for `protocols.ospfv3.redistribute.connected`](#nested-schema-for-protocolsospfv3redistributeconnected)
    - [Nested Schema for `protocols.ospfv3.redistribute.isis`](#nested-schema-for-protocolsospfv3redistributeisis)
    - [Nested Schema for `protocols.ospfv3.redistribute.kernel`](#nested-schema-for-protocolsospfv3redistributekernel)
    - [Nested Schema for `protocols.ospfv3.redistribute.ripng`](#nested-schema-for-protocolsospfv3redistributeripng)
    - [Nested Schema for `protocols.ospfv3.redistribute.static`](#nested-schema-for-protocolsospfv3redistributestatic)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### description
- `description` (String) Description

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  txt     &emsp;|  Description  |
#### disable
- `disable` (Boolean) Administratively disable interface
#### ip
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
#### ipv6
- `ipv6` (Attributes) IPv6 routing parameters (see [below for nested schema](#nestedatt--ipv6))
#### protocols
- `protocols` (Attributes) Routing protocol parameters (see [below for nested schema](#nestedatt--protocols))
#### table
- `table` (Number) Routing table associated with this instance

    |  Format     &emsp;|  Description       |
    |-------------|--------------------|
    |  100-65535  &emsp;|  Routing table ID  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
#### vni
- `vni` (Number) Virtual Network Identifier

    |  Format      &emsp;|  Description                       |
    |--------------|------------------------------------|
    |  0-16777214  &emsp;|  VXLAN virtual network identifier  |

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `name` (String) Virtual Routing and Forwarding instance

    |  Format  &emsp;|  Description        |
    |----------|---------------------|
    |  txt     &emsp;|  VRF instance name  |


<a id="nestedatt--ip"></a>
### Nested Schema for `ip`

Optional:

- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `nht` (Attributes) Filter Next Hop tracking route resolution (see [below for nested schema](#nestedatt--ip--nht))

<a id="nestedatt--ip--nht"></a>
### Nested Schema for `ip.nht`

Optional:

- `no_resolve_via_default` (Boolean) Do not resolve via default route



<a id="nestedatt--ipv6"></a>
### Nested Schema for `ipv6`

Optional:

- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `nht` (Attributes) Filter Next Hop tracking route resolution (see [below for nested schema](#nestedatt--ipv6--nht))

<a id="nestedatt--ipv6--nht"></a>
### Nested Schema for `ipv6.nht`

Optional:

- `no_resolve_via_default` (Boolean) Do not resolve via default route



<a id="nestedatt--protocols"></a>
### Nested Schema for `protocols`

Optional:

- `bgp` (Attributes) Border Gateway Protocol (BGP) (see [below for nested schema](#nestedatt--protocols--bgp))
- `eigrp` (Attributes) Enhanced Interior Gateway Routing Protocol (EIGRP) (see [below for nested schema](#nestedatt--protocols--eigrp))
- `isis` (Attributes) Intermediate System to Intermediate System (IS-IS) (see [below for nested schema](#nestedatt--protocols--isis))
- `ospf` (Attributes) Open Shortest Path First (OSPF) (see [below for nested schema](#nestedatt--protocols--ospf))
- `ospfv3` (Attributes) Open Shortest Path First (OSPF) for IPv6 (see [below for nested schema](#nestedatt--protocols--ospfv3))

<a id="nestedatt--protocols--bgp"></a>
### Nested Schema for `protocols.bgp`

Optional:

- `address_family` (Attributes) BGP address-family parameters (see [below for nested schema](#nestedatt--protocols--bgp--address_family))
- `bmp` (Attributes) BGP Monitoring Protocol (BMP) (see [below for nested schema](#nestedatt--protocols--bgp--bmp))
- `listen` (Attributes) Listen for and accept BGP dynamic neighbors from range (see [below for nested schema](#nestedatt--protocols--bgp--listen))
- `parameters` (Attributes) BGP parameters (see [below for nested schema](#nestedatt--protocols--bgp--parameters))
- `sid` (Attributes) SID value for VRF (see [below for nested schema](#nestedatt--protocols--bgp--sid))
- `srv6` (Attributes) Segment-Routing SRv6 configuration (see [below for nested schema](#nestedatt--protocols--bgp--srv6))
- `system_as` (Number) Autonomous System Number (ASN)

    |  Format        &emsp;|  Description               |
    |----------------|----------------------------|
    |  1-4294967294  &emsp;|  Autonomous System Number  |
- `timers` (Attributes) BGP protocol timers (see [below for nested schema](#nestedatt--protocols--bgp--timers))

<a id="nestedatt--protocols--bgp--address_family"></a>
### Nested Schema for `protocols.bgp.address_family`

Optional:

- `ipv4_flowspec` (Attributes) Flowspec IPv4 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_flowspec))
- `ipv4_labeled_unicast` (Attributes) Labeled Unicast IPv4 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_labeled_unicast))
- `ipv4_multicast` (Attributes) Multicast IPv4 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_multicast))
- `ipv4_unicast` (Attributes) IPv4 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast))
- `ipv6_flowspec` (Attributes) Flowspec IPv6 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_flowspec))
- `ipv6_multicast` (Attributes) Multicast IPv6 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_multicast))
- `ipv6_unicast` (Attributes) IPv6 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast))
- `l2vpn_evpn` (Attributes) L2VPN EVPN BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn))

<a id="nestedatt--protocols--bgp--address_family--ipv4_flowspec"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_flowspec`

Optional:

- `local_install` (Attributes) Apply local policy routing to interface (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_flowspec--local_install))

<a id="nestedatt--protocols--bgp--address_family--ipv4_flowspec--local_install"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_flowspec.local_install`

Optional:

- `interface` (List of String) Interface

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Interface name  |



<a id="nestedatt--protocols--bgp--address_family--ipv4_labeled_unicast"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_labeled_unicast`

Optional:

- `maximum_paths` (Attributes) Forward packets over multiple paths (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_labeled_unicast--maximum_paths))

<a id="nestedatt--protocols--bgp--address_family--ipv4_labeled_unicast--maximum_paths"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_labeled_unicast.maximum_paths`

Optional:

- `ebgp` (Number) eBGP maximum paths

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  1-256   &emsp;|  Number of paths to consider  |
- `ibgp` (Number) iBGP maximum paths

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  1-256   &emsp;|  Number of paths to consider  |



<a id="nestedatt--protocols--bgp--address_family--ipv4_multicast"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_multicast`

Optional:

- `distance` (Attributes) Administrative distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_multicast--distance))

<a id="nestedatt--protocols--bgp--address_family--ipv4_multicast--distance"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_multicast.distance`

Optional:

- `external` (Number) eBGP routes administrative distance

    |  Format  &emsp;|  Description                          |
    |----------|---------------------------------------|
    |  1-255   &emsp;|  eBGP routes administrative distance  |
- `internal` (Number) iBGP routes administrative distance

    |  Format  &emsp;|  Description                          |
    |----------|---------------------------------------|
    |  1-255   &emsp;|  iBGP routes administrative distance  |
- `local` (Number) Locally originated BGP routes administrative distance

    |  Format  &emsp;|  Description                                            |
    |----------|---------------------------------------------------------|
    |  1-255   &emsp;|  Locally originated BGP routes administrative distance  |



<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast`

Optional:

- `distance` (Attributes) Administrative distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--distance))
- `export` (Attributes) Export routes from this address-family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--export))
- `import` (Attributes) Import routes to this address-family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--import))
- `label` (Attributes) Label value for VRF (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--label))
- `maximum_paths` (Attributes) Forward packets over multiple paths (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--maximum_paths))
- `nexthop` (Attributes) Specify next hop to use for VRF advertised prefixes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--nexthop))
- `rd` (Attributes) Specify route distinguisher (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--rd))
- `redistribute` (Attributes) Redistribute routes from other protocols into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--route_map))
- `route_target` (Attributes) Specify route target list (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--route_target))
- `sid` (Attributes) SID value for VRF (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--sid))

<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--distance"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.distance`

Optional:

- `external` (Number) eBGP routes administrative distance

    |  Format  &emsp;|  Description                          |
    |----------|---------------------------------------|
    |  1-255   &emsp;|  eBGP routes administrative distance  |
- `internal` (Number) iBGP routes administrative distance

    |  Format  &emsp;|  Description                          |
    |----------|---------------------------------------|
    |  1-255   &emsp;|  iBGP routes administrative distance  |
- `local` (Number) Locally originated BGP routes administrative distance

    |  Format  &emsp;|  Description                                            |
    |----------|---------------------------------------------------------|
    |  1-255   &emsp;|  Locally originated BGP routes administrative distance  |


<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--export"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.export`

Optional:

- `vpn` (Boolean) to/from default instance VPN RIB


<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--import"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.import`

Optional:

- `vpn` (Boolean) to/from default instance VPN RIB
- `vrf` (List of String) VRF to import from

    |  Format  &emsp;|  Description        |
    |----------|---------------------|
    |  txt     &emsp;|  VRF instance name  |


<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--label"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.label`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--label--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--label--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.label.vpn`

Optional:

- `allocation_mode` (Attributes) Label allocation mode (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--label--vpn--allocation_mode))
- `export` (String) For routes leaked from current address-family to VPN

    |  Format     &emsp;|  Description                   |
    |-------------|--------------------------------|
    |  auto       &emsp;|  Automatically assign a label  |
    |  0-1048575  &emsp;|  Label Value                   |

<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--label--vpn--allocation_mode"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.label.vpn.allocation_mode`

Optional:

- `per_nexthop` (Boolean) Allocate a label per connected next-hop in the VRF




<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--maximum_paths"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.maximum_paths`

Optional:

- `ebgp` (Number) eBGP maximum paths

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  1-256   &emsp;|  Number of paths to consider  |
- `ibgp` (Number) iBGP maximum paths

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  1-256   &emsp;|  Number of paths to consider  |


<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--nexthop"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.nexthop`

Optional:

- `vpn` (Attributes) Between current address-family and vpn (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--nexthop--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--nexthop--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.nexthop.vpn`

Optional:

- `export` (String) For routes leaked from current address-family to vpn

    |  Format  &emsp;|  Description                |
    |----------|-----------------------------|
    |  ipv4    &emsp;|  BGP neighbor IP address    |
    |  ipv6    &emsp;|  BGP neighbor IPv6 address  |



<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--rd"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.rd`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--rd--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--rd--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.rd.vpn`

Optional:

- `export` (String) For routes leaked from current address-family to VPN

    |  Format                   &emsp;|  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy&emsp;|xxxx:yyyy)  |



<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute`

Optional:

- `babel` (Attributes) Redistribute Babel routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--babel))
- `connected` (Attributes) Redistribute connected routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--connected))
- `isis` (Attributes) Redistribute IS-IS routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--isis))
- `kernel` (Attributes) Redistribute kernel routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--kernel))
- `ospf` (Attributes) Redistribute OSPF routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--ospf))
- `rip` (Attributes) Redistribute RIP routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--rip))
- `static` (Attributes) Redistribute static routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--static))

<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--babel"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.babel`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--connected"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.connected`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--isis"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.isis`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--kernel"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.kernel`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--ospf"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.ospf`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--rip"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.rip`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--static"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.static`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--route_map"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_map`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--route_map--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--route_map--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_map.vpn`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `import` (String) Route-map to filter incoming route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--route_target"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_target`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--route_target--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--route_target--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_target.vpn`

Optional:

- `both` (String) Route Target both import and export

    |  Format  &emsp;|  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |
- `export` (String) Route Target export

    |  Format  &emsp;|  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |
- `import` (String) Route Target import

    |  Format  &emsp;|  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |



<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--sid"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.sid`

Optional:

- `vpn` (Attributes) Between current VRF and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--sid--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv4_unicast--sid--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.sid.vpn`

Optional:

- `export` (String) For routes leaked from current VRF to VPN

    |  Format     &emsp;|  Description                   |
    |-------------|--------------------------------|
    |  1-1048575  &emsp;|  SID allocation index          |
    |  auto       &emsp;|  Automatically assign a label  |




<a id="nestedatt--protocols--bgp--address_family--ipv6_flowspec"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_flowspec`

Optional:

- `local_install` (Attributes) Apply local policy routing to interface (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_flowspec--local_install))

<a id="nestedatt--protocols--bgp--address_family--ipv6_flowspec--local_install"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_flowspec.local_install`

Optional:

- `interface` (List of String) Interface

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Interface name  |



<a id="nestedatt--protocols--bgp--address_family--ipv6_multicast"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_multicast`

Optional:

- `distance` (Attributes) Administrative distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_multicast--distance))

<a id="nestedatt--protocols--bgp--address_family--ipv6_multicast--distance"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_multicast.distance`

Optional:

- `external` (Number) eBGP routes administrative distance

    |  Format  &emsp;|  Description                          |
    |----------|---------------------------------------|
    |  1-255   &emsp;|  eBGP routes administrative distance  |
- `internal` (Number) iBGP routes administrative distance

    |  Format  &emsp;|  Description                          |
    |----------|---------------------------------------|
    |  1-255   &emsp;|  iBGP routes administrative distance  |
- `local` (Number) Locally originated BGP routes administrative distance

    |  Format  &emsp;|  Description                                            |
    |----------|---------------------------------------------------------|
    |  1-255   &emsp;|  Locally originated BGP routes administrative distance  |



<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast`

Optional:

- `distance` (Attributes) Administrative distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--distance))
- `export` (Attributes) Export routes from this address-family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--export))
- `import` (Attributes) Import routes to this address-family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--import))
- `label` (Attributes) Label value for VRF (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--label))
- `maximum_paths` (Attributes) Forward packets over multiple paths (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--maximum_paths))
- `nexthop` (Attributes) Specify next hop to use for VRF advertised prefixes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--nexthop))
- `rd` (Attributes) Specify route distinguisher (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--rd))
- `redistribute` (Attributes) Redistribute routes from other protocols into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--route_map))
- `route_target` (Attributes) Specify route target list (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--route_target))
- `sid` (Attributes) SID value for VRF (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--sid))

<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--distance"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.distance`

Optional:

- `external` (Number) eBGP routes administrative distance

    |  Format  &emsp;|  Description                          |
    |----------|---------------------------------------|
    |  1-255   &emsp;|  eBGP routes administrative distance  |
- `internal` (Number) iBGP routes administrative distance

    |  Format  &emsp;|  Description                          |
    |----------|---------------------------------------|
    |  1-255   &emsp;|  iBGP routes administrative distance  |
- `local` (Number) Locally originated BGP routes administrative distance

    |  Format  &emsp;|  Description                                            |
    |----------|---------------------------------------------------------|
    |  1-255   &emsp;|  Locally originated BGP routes administrative distance  |


<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--export"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.export`

Optional:

- `vpn` (Boolean) to/from default instance VPN RIB


<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--import"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.import`

Optional:

- `vpn` (Boolean) to/from default instance VPN RIB
- `vrf` (List of String) VRF to import from

    |  Format  &emsp;|  Description        |
    |----------|---------------------|
    |  txt     &emsp;|  VRF instance name  |


<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--label"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.label`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--label--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--label--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.label.vpn`

Optional:

- `allocation_mode` (Attributes) Label allocation mode (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--label--vpn--allocation_mode))
- `export` (String) For routes leaked from current address-family to VPN

    |  Format     &emsp;|  Description                   |
    |-------------|--------------------------------|
    |  auto       &emsp;|  Automatically assign a label  |
    |  0-1048575  &emsp;|  Label Value                   |

<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--label--vpn--allocation_mode"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.label.vpn.allocation_mode`

Optional:

- `per_nexthop` (Boolean) Allocate a label per connected next-hop in the VRF




<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--maximum_paths"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.maximum_paths`

Optional:

- `ebgp` (Number) eBGP maximum paths

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  1-256   &emsp;|  Number of paths to consider  |
- `ibgp` (Number) iBGP maximum paths

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  1-256   &emsp;|  Number of paths to consider  |


<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--nexthop"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.nexthop`

Optional:

- `vpn` (Attributes) Between current address-family and vpn (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--nexthop--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--nexthop--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.nexthop.vpn`

Optional:

- `export` (String) For routes leaked from current address-family to vpn

    |  Format  &emsp;|  Description                |
    |----------|-----------------------------|
    |  ipv4    &emsp;|  BGP neighbor IP address    |
    |  ipv6    &emsp;|  BGP neighbor IPv6 address  |



<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--rd"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.rd`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--rd--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--rd--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.rd.vpn`

Optional:

- `export` (String) For routes leaked from current address-family to VPN

    |  Format                   &emsp;|  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy&emsp;|xxxx:yyyy)  |



<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute`

Optional:

- `babel` (Attributes) Redistribute Babel routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--babel))
- `connected` (Attributes) Redistribute connected routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--connected))
- `isis` (Attributes) Redistribute IS-IS routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--isis))
- `kernel` (Attributes) Redistribute kernel routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--kernel))
- `ospfv3` (Attributes) Redistribute OSPFv3 routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--ospfv3))
- `ripng` (Attributes) Redistribute RIPng routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--ripng))
- `static` (Attributes) Redistribute static routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--static))

<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--babel"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.babel`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--connected"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.connected`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--isis"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.isis`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--kernel"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.kernel`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--ospfv3"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.ospfv3`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--ripng"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.ripng`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--static"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.static`

Optional:

- `metric` (Number) Metric for redistributed routes

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--route_map"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_map`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--route_map--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--route_map--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_map.vpn`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |
- `import` (String) Route-map to filter incoming route updates

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--route_target"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_target`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--route_target--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--route_target--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_target.vpn`

Optional:

- `both` (String) Route Target both import and export

    |  Format  &emsp;|  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |
- `export` (String) Route Target export

    |  Format  &emsp;|  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |
- `import` (String) Route Target import

    |  Format  &emsp;|  Description                                                     |
    |----------|------------------------------------------------------------------|
    |  txt     |  Space separated route target list (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |



<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--sid"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.sid`

Optional:

- `vpn` (Attributes) Between current VRF and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--sid--vpn))

<a id="nestedatt--protocols--bgp--address_family--ipv6_unicast--sid--vpn"></a>
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.sid.vpn`

Optional:

- `export` (String) For routes leaked from current VRF to VPN

    |  Format     &emsp;|  Description                   |
    |-------------|--------------------------------|
    |  1-1048575  &emsp;|  SID allocation index          |
    |  auto       &emsp;|  Automatically assign a label  |




<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn`

Optional:

- `advertise` (Attributes) Advertise prefix routes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise))
- `advertise_all_vni` (Boolean) Advertise All local VNIs
- `advertise_default_gw` (Boolean) Advertise All default g/w mac-ip routes in EVPN
- `advertise_pip` (String) EVPN system primary IP

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  ipv4    &emsp;|  IP address   |
- `advertise_svi_ip` (Boolean) Advertise svi mac-ip routes in EVPN
- `default_originate` (Attributes) Originate a default route (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--default_originate))
- `disable_ead_evi_rx` (Boolean) Activate PE on EAD-ES even if EAD-EVI is not received
- `disable_ead_evi_tx` (Boolean) Do not advertise EAD-EVI for local ESs
- `ead_es_frag` (Attributes) EAD ES fragment config (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--ead_es_frag))
- `ead_es_route_target` (Attributes) EAD ES Route Target (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--ead_es_route_target))
- `flooding` (Attributes) Specify handling for BUM packets (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--flooding))
- `mac_vrf` (Attributes) EVPN MAC-VRF (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--mac_vrf))
- `rd` (String) Route Distinguisher

    |  Format                   &emsp;|  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy&emsp;|xxxx:yyyy)  |
- `route_target` (Attributes) Route Target (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--route_target))
- `rt_auto_derive` (Boolean) Auto derivation of Route Target (RFC8365)

<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise`

Optional:

- `ipv4` (Attributes) IPv4 address family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv4))
- `ipv6` (Attributes) IPv6 address family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv6))

<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv4"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv4`

Optional:

- `unicast` (Attributes) IPv4 address family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv4--unicast))

<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv4--unicast"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv4.unicast`

Optional:

- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv6"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv6`

Optional:

- `unicast` (Attributes) IPv4 address family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv6--unicast))

<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv6--unicast"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv6.unicast`

Optional:

- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |




<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--default_originate"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.default_originate`

Optional:

- `ipv4` (Boolean) IPv4 address family
- `ipv6` (Boolean) IPv6 address family


<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--ead_es_frag"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.ead_es_frag`

Optional:

- `evi_limit` (Number) EVIs per-fragment

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  1-1000  &emsp;|  limit        |


<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--ead_es_route_target"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.ead_es_route_target`

Optional:

- `export` (List of String) Route Target export

    |  Format  &emsp;|  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |


<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--flooding"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.flooding`

Optional:

- `disable` (Boolean) Disable instance
- `head_end_replication` (Boolean) Flood BUM packets using head-end replication


<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--mac_vrf"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.mac_vrf`

Optional:

- `soo` (String) Site-of-Origin extended community

    |  Format  &emsp;|  Description                                                         |
    |----------|----------------------------------------------------------------------|
    |  ASN:NN  &emsp;|  based on autonomous system number in format &lt;0-65535:0-4294967295&gt;  |
    |  IP:NN   &emsp;|  Based on a router-id IP address in format &lt;IP:0-65535&gt;              |


<a id="nestedatt--protocols--bgp--address_family--l2vpn_evpn--route_target"></a>
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.route_target`

Optional:

- `both` (List of String) Route Target both import and export

    |  Format  &emsp;|  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |
- `export` (List of String) Route Target export

    |  Format  &emsp;|  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |
- `import` (List of String) Route Target import

    |  Format  &emsp;|  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |




<a id="nestedatt--protocols--bgp--bmp"></a>
### Nested Schema for `protocols.bgp.bmp`

Optional:

- `mirror_buffer_limit` (Number) Maximum memory used for buffered mirroring messages (in bytes)

    |  Format        &emsp;|  Description     |
    |----------------|------------------|
    |  0-4294967294  &emsp;|  Limit in bytes  |


<a id="nestedatt--protocols--bgp--listen"></a>
### Nested Schema for `protocols.bgp.listen`

Optional:

- `limit` (Number) Maximum number of dynamic neighbors that can be created

    |  Format  &emsp;|  Description         |
    |----------|----------------------|
    |  1-5000  &emsp;|  BGP neighbor limit  |


<a id="nestedatt--protocols--bgp--parameters"></a>
### Nested Schema for `protocols.bgp.parameters`

Optional:

- `allow_martian_nexthop` (Boolean) Allow Martian nexthops to be received in the NLRI from a peer
- `always_compare_med` (Boolean) Always compare MEDs from different neighbors
- `bestpath` (Attributes) Default bestpath selection mechanism (see [below for nested schema](#nestedatt--protocols--bgp--parameters--bestpath))
- `cluster_id` (String) Route-reflector cluster-id

    |  Format  &emsp;|  Description                 |
    |----------|------------------------------|
    |  ipv4    &emsp;|  Route-reflector cluster-id  |
- `conditional_advertisement` (Attributes) Conditional advertisement settings (see [below for nested schema](#nestedatt--protocols--bgp--parameters--conditional_advertisement))
- `confederation` (Attributes) AS confederation parameters (see [below for nested schema](#nestedatt--protocols--bgp--parameters--confederation))
- `dampening` (Attributes) Enable route-flap dampening (see [below for nested schema](#nestedatt--protocols--bgp--parameters--dampening))
- `default` (Attributes) BGP defaults (see [below for nested schema](#nestedatt--protocols--bgp--parameters--default))
- `deterministic_med` (Boolean) Compare MEDs between different peers in the same AS
- `disable_ebgp_connected_route_check` (Boolean) Disable checking if nexthop is connected on eBGP session
- `distance` (Attributes) Administratives distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--parameters--distance))
- `ebgp_requires_policy` (Boolean) Require in and out policy for eBGP peers (RFC8212)
- `fast_convergence` (Boolean) Teardown sessions immediately whenever peer becomes unreachable
- `graceful_restart` (Attributes) Graceful restart capability parameters (see [below for nested schema](#nestedatt--protocols--bgp--parameters--graceful_restart))
- `graceful_shutdown` (Boolean) Graceful shutdown
- `labeled_unicast` (String) BGP Labeled-unicast options

    |  Format              &emsp;|  Description                                                 |
    |----------------------|--------------------------------------------------------------|
    |  explicit-null       &emsp;|  Use explicit-null label values for all local prefixes       |
    |  ipv4-explicit-null  &emsp;|  Use IPv4 explicit-null label value for IPv4 local prefixes  |
    |  ipv6-explicit-null  &emsp;|  Use IPv6 explicit-null label value for IPv4 local prefixes  |
- `log_neighbor_changes` (Boolean) Log neighbor up/down changes and reset reason
- `minimum_holdtime` (Number) BGP minimum holdtime

    |  Format   &emsp;|  Description                  |
    |-----------|-------------------------------|
    |  1-65535  &emsp;|  Minimum holdtime in seconds  |
- `network_import_check` (Boolean) Enable IGP route check for network statements
- `no_client_to_client_reflection` (Boolean) Disable client to client route reflection
- `no_fast_external_failover` (Boolean) Disable immediate session reset on peer link down event
- `no_hard_administrative_reset` (Boolean) Do not send hard reset CEASE Notification for &#39;Administrative Reset&#39;
- `no_suppress_duplicates` (Boolean) Disable suppress duplicate updates if the route actually not changed
- `reject_as_sets` (Boolean) Reject routes with AS_SET or AS_CONFED_SET flag
- `route_reflector_allow_outbound_policy` (Boolean) Route reflector client allow policy outbound
- `router_id` (String) Override default router identifier

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  ipv4    &emsp;|  Router-ID in IP address format  |
- `shutdown` (Boolean) Administrative shutdown of the BGP instance
- `suppress_fib_pending` (Boolean) Advertise only routes that are programmed in kernel to peers
- `tcp_keepalive` (Attributes) TCP keepalive parameters (see [below for nested schema](#nestedatt--protocols--bgp--parameters--tcp_keepalive))

<a id="nestedatt--protocols--bgp--parameters--bestpath"></a>
### Nested Schema for `protocols.bgp.parameters.bestpath`

Optional:

- `as_path` (Attributes) AS-path attribute comparison parameters (see [below for nested schema](#nestedatt--protocols--bgp--parameters--bestpath--as_path))
- `bandwidth` (String) Link Bandwidth attribute

    |  Format                      &emsp;|  Description                                                            |
    |------------------------------|-------------------------------------------------------------------------|
    |  default-weight-for-missing  &emsp;|  Assign low default weight (1) to paths not having link bandwidth       |
    |  ignore                      &emsp;|  Ignore link bandwidth (do regular ECMP, not weighted)                  |
    |  skip-missing                &emsp;|  Ignore paths without link bandwidth for ECMP (if other paths have it)  |
- `compare_routerid` (Boolean) Compare the router-id for identical EBGP paths
- `med` (List of String) MED attribute comparison parameters

    |  Format            &emsp;|  Description                                              |
    |--------------------|-----------------------------------------------------------|
    |  confed            &emsp;|  Compare MEDs among confederation paths                   |
    |  missing-as-worst  &emsp;|  Treat missing route as a MED as the least preferred one  |
- `peer_type` (Attributes) Peer type (see [below for nested schema](#nestedatt--protocols--bgp--parameters--bestpath--peer_type))

<a id="nestedatt--protocols--bgp--parameters--bestpath--as_path"></a>
### Nested Schema for `protocols.bgp.parameters.bestpath.as_path`

Optional:

- `confed` (Boolean) Compare AS-path lengths including confederation sets and sequences
- `ignore` (Boolean) Ignore AS-path length in selecting a route
- `multipath_relax` (Boolean) Allow load sharing across routes that have different AS paths (but same length)


<a id="nestedatt--protocols--bgp--parameters--bestpath--peer_type"></a>
### Nested Schema for `protocols.bgp.parameters.bestpath.peer_type`

Optional:

- `multipath_relax` (Boolean) Allow load sharing across routes learned from different peer types



<a id="nestedatt--protocols--bgp--parameters--conditional_advertisement"></a>
### Nested Schema for `protocols.bgp.parameters.conditional_advertisement`

Optional:

- `timer` (Number) Set period to rescan BGP table to check if condition is met

    |  Format  &emsp;|  Description                                                    |
    |----------|-----------------------------------------------------------------|
    |  5-240   &emsp;|  Period to rerun the conditional advertisement scanner process  |


<a id="nestedatt--protocols--bgp--parameters--confederation"></a>
### Nested Schema for `protocols.bgp.parameters.confederation`

Optional:

- `identifier` (Number) Confederation AS identifier

    |  Format        &emsp;|  Description          |
    |----------------|-----------------------|
    |  1-4294967294  &emsp;|  Confederation AS id  |
- `peers` (List of Number) Peer ASs in the BGP confederation

    |  Format        &emsp;|  Description     |
    |----------------|------------------|
    |  1-4294967294  &emsp;|  Peer AS number  |


<a id="nestedatt--protocols--bgp--parameters--dampening"></a>
### Nested Schema for `protocols.bgp.parameters.dampening`

Optional:

- `half_life` (Number) Half-life time for dampening

    |  Format  &emsp;|  Description                   |
    |----------|--------------------------------|
    |  1-45    &emsp;|  Half-life penalty in minutes  |
- `max_suppress_time` (Number) Maximum duration to suppress a stable route

    |  Format  &emsp;|  Description                           |
    |----------|----------------------------------------|
    |  1-255   &emsp;|  Maximum suppress duration in minutes  |
- `re_use` (Number) Threshold to start reusing a route

    |  Format   &emsp;|  Description            |
    |-----------|-------------------------|
    |  1-20000  &emsp;|  Re-use penalty points  |
- `start_suppress_time` (Number) When to start suppressing a route

    |  Format   &emsp;|  Description                    |
    |-----------|---------------------------------|
    |  1-20000  &emsp;|  Start-suppress penalty points  |


<a id="nestedatt--protocols--bgp--parameters--default"></a>
### Nested Schema for `protocols.bgp.parameters.default`

Optional:

- `local_pref` (Number) Default local preference

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  u32     &emsp;|  Local preference  |


<a id="nestedatt--protocols--bgp--parameters--distance"></a>
### Nested Schema for `protocols.bgp.parameters.distance`

Optional:

- `global` (Attributes) Global administratives distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--parameters--distance--global))

<a id="nestedatt--protocols--bgp--parameters--distance--global"></a>
### Nested Schema for `protocols.bgp.parameters.distance.global`

Optional:

- `external` (Number) Administrative distance for external BGP routes

    |  Format  &emsp;|  Description                                      |
    |----------|---------------------------------------------------|
    |  1-255   &emsp;|  Administrative distance for external BGP routes  |
- `internal` (Number) Administrative distance for internal BGP routes

    |  Format  &emsp;|  Description                                      |
    |----------|---------------------------------------------------|
    |  1-255   &emsp;|  Administrative distance for internal BGP routes  |
- `local` (Number) Administrative distance for local BGP routes

    |  Format  &emsp;|  Description                                      |
    |----------|---------------------------------------------------|
    |  1-255   &emsp;|  Administrative distance for internal BGP routes  |



<a id="nestedatt--protocols--bgp--parameters--graceful_restart"></a>
### Nested Schema for `protocols.bgp.parameters.graceful_restart`

Optional:

- `stalepath_time` (Number) Maximum time to hold onto restarting neighbors stale paths

    |  Format  &emsp;|  Description           |
    |----------|------------------------|
    |  1-3600  &emsp;|  Hold time in seconds  |


<a id="nestedatt--protocols--bgp--parameters--tcp_keepalive"></a>
### Nested Schema for `protocols.bgp.parameters.tcp_keepalive`

Optional:

- `idle` (Number) TCP keepalive idle time

    |  Format   &emsp;|  Description           |
    |-----------|------------------------|
    |  1-65535  &emsp;|  Idle time in seconds  |
- `interval` (Number) TCP keepalive interval

    |  Format   &emsp;|  Description          |
    |-----------|-----------------------|
    |  1-65535  &emsp;|  Interval in seconds  |
- `probes` (Number) TCP keepalive maximum probes

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  1-30    &emsp;|  Maximum probes  |



<a id="nestedatt--protocols--bgp--sid"></a>
### Nested Schema for `protocols.bgp.sid`

Optional:

- `vpn` (Attributes) Between current VRF and VPN (see [below for nested schema](#nestedatt--protocols--bgp--sid--vpn))

<a id="nestedatt--protocols--bgp--sid--vpn"></a>
### Nested Schema for `protocols.bgp.sid.vpn`

Optional:

- `per_vrf` (Attributes) SID per-VRF (both IPv4 and IPv6 address families) (see [below for nested schema](#nestedatt--protocols--bgp--sid--vpn--per_vrf))

<a id="nestedatt--protocols--bgp--sid--vpn--per_vrf"></a>
### Nested Schema for `protocols.bgp.sid.vpn.per_vrf`

Optional:

- `export` (String) For routes leaked from current VRF to VPN

    |  Format     &emsp;|  Description                   |
    |-------------|--------------------------------|
    |  1-1048575  &emsp;|  SID allocation index          |
    |  auto       &emsp;|  Automatically assign a label  |




<a id="nestedatt--protocols--bgp--srv6"></a>
### Nested Schema for `protocols.bgp.srv6`

Optional:

- `locator` (String) Specify SRv6 locator

    |  Format  &emsp;|  Description        |
    |----------|---------------------|
    |  txt     &emsp;|  SRv6 locator name  |


<a id="nestedatt--protocols--bgp--timers"></a>
### Nested Schema for `protocols.bgp.timers`

Optional:

- `holdtime` (String) Hold timer

    |  Format   &emsp;|  Description            |
    |-----------|-------------------------|
    |  1-65535  &emsp;|  Hold timer in seconds  |
    |  0        &emsp;|  Disable hold timer     |
- `keepalive` (Number) BGP keepalive interval for this neighbor

    |  Format   &emsp;|  Description                    |
    |-----------|---------------------------------|
    |  1-65535  &emsp;|  Keepalive interval in seconds  |



<a id="nestedatt--protocols--eigrp"></a>
### Nested Schema for `protocols.eigrp`

Optional:

- `maximum_paths` (Number) Forward packets over multiple paths

    |  Format  &emsp;|  Description      |
    |----------|-------------------|
    |  1-32    &emsp;|  Number of paths  |
- `metric` (Attributes) Modify metrics and parameters for advertisement (see [below for nested schema](#nestedatt--protocols--eigrp--metric))
- `network` (List of String) Enable routing on an IP network

    |  Format   &emsp;|  Description           |
    |-----------|------------------------|
    |  ipv4net  &emsp;|  EIGRP network prefix  |
- `passive_interface` (List of String) Suppress routing updates on an interface
- `redistribute` (List of String) Redistribute information from another routing protocol

    |  Format     &emsp;|  Description                          |
    |-------------|---------------------------------------|
    |  bgp        &emsp;|  Border Gateway Protocol (BGP)        |
    |  connected  &emsp;|  Connected routes                     |
    |  nhrp       &emsp;|  Next Hop Resolution Protocol (NHRP)  |
    |  ospf       &emsp;|  Open Shortest Path First (OSPFv2)    |
    |  rip        &emsp;|  Routing Information Protocol (RIP)   |
    |  babel      &emsp;|  Babel routing protocol (Babel)       |
    |  static     &emsp;|  Statically configured routes         |
    |  vnc        &emsp;|  Virtual Network Control (VNC)        |
- `router_id` (String) Override default router identifier

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  ipv4    &emsp;|  Router-ID in IP address format  |
- `system_as` (Number) Autonomous System Number (ASN)

    |  Format   &emsp;|  Description               |
    |-----------|----------------------------|
    |  1-65535  &emsp;|  Autonomous System Number  |
- `variance` (Number) Control load balancing variance

    |  Format  &emsp;|  Description                 |
    |----------|------------------------------|
    |  1-128   &emsp;|  Metric variance multiplier  |

<a id="nestedatt--protocols--eigrp--metric"></a>
### Nested Schema for `protocols.eigrp.metric`

Optional:

- `weights` (Number) Modify metric coefficients

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  0-255   &emsp;|  K1           |



<a id="nestedatt--protocols--isis"></a>
### Nested Schema for `protocols.isis`

Optional:

- `advertise_high_metrics` (Boolean) Advertise high metric value on all interfaces
- `advertise_passive_only` (Boolean) Advertise prefixes of passive interfaces only
- `area_password` (Attributes) Configure the authentication password for an area (see [below for nested schema](#nestedatt--protocols--isis--area_password))
- `default_information` (Attributes) Control distribution of default information (see [below for nested schema](#nestedatt--protocols--isis--default_information))
- `domain_password` (Attributes) Set the authentication password for a routing domain (see [below for nested schema](#nestedatt--protocols--isis--domain_password))
- `dynamic_hostname` (Boolean) Dynamic hostname for IS-IS
- `fast_reroute` (Attributes) IS-IS fast reroute configuration (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute))
- `ldp_sync` (Attributes) Protocol wide LDP-IGP synchronization configuration (see [below for nested schema](#nestedatt--protocols--isis--ldp_sync))
- `level` (String) IS-IS level number

    |  Format     &emsp;|  Description                               |
    |-------------|--------------------------------------------|
    |  level-1    &emsp;|  Act as a station router                   |
    |  level-1-2  &emsp;|  Act as both a station and an area router  |
    |  level-2    &emsp;|  Act as an area router                     |
- `log_adjacency_changes` (Boolean) Log changes in adjacency state
- `lsp_gen_interval` (Number) Minimum interval between regenerating same LSP

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  1-120   &emsp;|  Minimum interval in seconds  |
- `lsp_mtu` (Number) Configure the maximum size of generated LSPs

    |  Format    &emsp;|  Description                     |
    |------------|----------------------------------|
    |  128-4352  &emsp;|  Maximum size of generated LSPs  |
- `lsp_refresh_interval` (Number) LSP refresh interval

    |  Format   &emsp;|  Description                      |
    |-----------|-----------------------------------|
    |  1-65235  &emsp;|  LSP refresh interval in seconds  |
- `max_lsp_lifetime` (Number) Maximum LSP lifetime

    |  Format     &emsp;|  Description              |
    |-------------|---------------------------|
    |  350-65535  &emsp;|  LSP lifetime in seconds  |
- `metric_style` (String) Use old-style (ISO 10589) or new-style packet formats

    |  Format      &emsp;|  Description                                            |
    |--------------|---------------------------------------------------------|
    |  narrow      &emsp;|  Use old style of TLVs with narrow metric               |
    |  transition  &emsp;|  Send and accept both styles of TLVs during transition  |
    |  wide        &emsp;|  Use new style of TLVs to carry wider metric            |
- `net` (String) A Network Entity Title for the process (ISO only)

    |  Format                &emsp;|  Description                 |
    |------------------------|------------------------------|
    |  XX.XXXX. ... .XXX.XX  &emsp;|  Network entity title (NET)  |
- `purge_originator` (Boolean) Use the RFC 6232 purge-originator
- `redistribute` (Attributes) Redistribute information from another routing protocol (see [below for nested schema](#nestedatt--protocols--isis--redistribute))
- `segment_routing` (Attributes) Segment-Routing (SPRING) settings (see [below for nested schema](#nestedatt--protocols--isis--segment_routing))
- `set_attached_bit` (Boolean) Set attached bit to identify as L1/L2 router for inter-area traffic
- `set_overload_bit` (Boolean) Set overload bit to avoid any transit traffic
- `spf_delay_ietf` (Attributes) IETF SPF delay algorithm (see [below for nested schema](#nestedatt--protocols--isis--spf_delay_ietf))
- `spf_interval` (Number) Minimum interval between SPF calculations

    |  Format  &emsp;|  Description          |
    |----------|-----------------------|
    |  1-120   &emsp;|  Interval in seconds  |
- `topology` (String) Configure IS-IS topologies

    |  Format          &emsp;|  Description                   |
    |------------------|--------------------------------|
    |  ipv4-multicast  &emsp;|  Use IPv4 multicast topology   |
    |  ipv4-mgmt       &emsp;|  Use IPv4 management topology  |
    |  ipv6-unicast    &emsp;|  Use IPv6 unicast topology     |
    |  ipv6-multicast  &emsp;|  Use IPv6 multicast topology   |
    |  ipv6-mgmt       &emsp;|  Use IPv6 management topology  |
    |  ipv6-dstsrc     &emsp;|  Use IPv6 dst-src topology     |
- `traffic_engineering` (Attributes) IS-IS traffic engineering extensions (see [below for nested schema](#nestedatt--protocols--isis--traffic_engineering))

<a id="nestedatt--protocols--isis--area_password"></a>
### Nested Schema for `protocols.isis.area_password`

Optional:

- `md5` (String) MD5 authentication type

    |  Format  &emsp;|  Description          |
    |----------|-----------------------|
    |  txt     &emsp;|  Level-wide password  |
- `plaintext_password` (String) Plain-text authentication type

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  txt     &emsp;|  Circuit password  |


<a id="nestedatt--protocols--isis--default_information"></a>
### Nested Schema for `protocols.isis.default_information`

Optional:

- `originate` (Attributes) Distribute a default route (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate))

<a id="nestedatt--protocols--isis--default_information--originate"></a>
### Nested Schema for `protocols.isis.default_information.originate`

Optional:

- `ipv4` (Attributes) Distribute default route for IPv4 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv4))
- `ipv6` (Attributes) Distribute default route for IPv6 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv6))

<a id="nestedatt--protocols--isis--default_information--originate--ipv4"></a>
### Nested Schema for `protocols.isis.default_information.originate.ipv4`

Optional:

- `level_1` (Attributes) Distribute default route into level-1 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv4--level_1))
- `level_2` (Attributes) Distribute default route into level-2 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv4--level_2))

<a id="nestedatt--protocols--isis--default_information--originate--ipv4--level_1"></a>
### Nested Schema for `protocols.isis.default_information.originate.ipv4.level_1`

Optional:

- `always` (Boolean) Always advertise default route
- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--default_information--originate--ipv4--level_2"></a>
### Nested Schema for `protocols.isis.default_information.originate.ipv4.level_2`

Optional:

- `always` (Boolean) Always advertise default route
- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--default_information--originate--ipv6"></a>
### Nested Schema for `protocols.isis.default_information.originate.ipv6`

Optional:

- `level_1` (Attributes) Distribute default route into level-1 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv6--level_1))
- `level_2` (Attributes) Distribute default route into level-2 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv6--level_2))

<a id="nestedatt--protocols--isis--default_information--originate--ipv6--level_1"></a>
### Nested Schema for `protocols.isis.default_information.originate.ipv6.level_1`

Optional:

- `always` (Boolean) Always advertise default route
- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--default_information--originate--ipv6--level_2"></a>
### Nested Schema for `protocols.isis.default_information.originate.ipv6.level_2`

Optional:

- `always` (Boolean) Always advertise default route
- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |





<a id="nestedatt--protocols--isis--domain_password"></a>
### Nested Schema for `protocols.isis.domain_password`

Optional:

- `md5` (String) MD5 authentication type

    |  Format  &emsp;|  Description          |
    |----------|-----------------------|
    |  txt     &emsp;|  Level-wide password  |
- `plaintext_password` (String) Plain-text authentication type

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  txt     &emsp;|  Circuit password  |


<a id="nestedatt--protocols--isis--fast_reroute"></a>
### Nested Schema for `protocols.isis.fast_reroute`

Optional:

- `lfa` (Attributes) Loop free alternate functionality (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa))

<a id="nestedatt--protocols--isis--fast_reroute--lfa"></a>
### Nested Schema for `protocols.isis.fast_reroute.lfa`

Optional:

- `local` (Attributes) Local loop free alternate options (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local))

<a id="nestedatt--protocols--isis--fast_reroute--lfa--local"></a>
### Nested Schema for `protocols.isis.fast_reroute.lfa.local`

Optional:

- `load_sharing` (Attributes) Load share prefixes across multiple backups (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--load_sharing))
- `priority_limit` (Attributes) Limit backup computation up to the prefix priority (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit))

<a id="nestedatt--protocols--isis--fast_reroute--lfa--local--load_sharing"></a>
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.load_sharing`

Optional:

- `disable` (Attributes) Disable load sharing (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--load_sharing--disable))

<a id="nestedatt--protocols--isis--fast_reroute--lfa--local--load_sharing--disable"></a>
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.load_sharing.disable`

Optional:

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes



<a id="nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit"></a>
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit`

Optional:

- `critical` (Attributes) Compute for critical priority prefixes only (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--critical))
- `high` (Attributes) Compute for critical, and high priority prefixes (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--high))
- `medium` (Attributes) Compute for critical, high, and medium priority prefixes (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--medium))

<a id="nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--critical"></a>
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit.critical`

Optional:

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes


<a id="nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--high"></a>
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit.high`

Optional:

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes


<a id="nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--medium"></a>
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit.medium`

Optional:

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes






<a id="nestedatt--protocols--isis--ldp_sync"></a>
### Nested Schema for `protocols.isis.ldp_sync`

Optional:

- `holddown` (Number) Hold down timer for LDP-IGP cost restoration

    |  Format   &emsp;|  Description                                                                                   |
    |-----------|------------------------------------------------------------------------------------------------|
    |  0-10000  &emsp;|  Time to wait in seconds for LDP-IGP synchronization to occur before restoring interface cost  |


<a id="nestedatt--protocols--isis--redistribute"></a>
### Nested Schema for `protocols.isis.redistribute`

Optional:

- `ipv4` (Attributes) Redistribute IPv4 routes (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4))
- `ipv6` (Attributes) Redistribute IPv6 routes (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6))

<a id="nestedatt--protocols--isis--redistribute--ipv4"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4`

Optional:

- `babel` (Attributes) Redistribute Babel routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--babel))
- `bgp` (Attributes) Border Gateway Protocol (BGP) (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--bgp))
- `connected` (Attributes) Redistribute connected routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--connected))
- `kernel` (Attributes) Redistribute kernel routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--kernel))
- `ospf` (Attributes) Redistribute OSPF routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--ospf))
- `rip` (Attributes) Redistribute RIP routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--rip))
- `static` (Attributes) Redistribute static routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--static))

<a id="nestedatt--protocols--isis--redistribute--ipv4--babel"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.babel`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--babel--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--babel--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv4--babel--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.babel.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv4--babel--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.babel.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv4--bgp"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.bgp`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--bgp--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--bgp--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv4--bgp--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.bgp.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv4--bgp--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.bgp.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv4--connected"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.connected`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--connected--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--connected--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv4--connected--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.connected.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv4--connected--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.connected.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv4--kernel"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.kernel`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--kernel--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--kernel--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv4--kernel--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.kernel.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv4--kernel--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.kernel.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv4--ospf"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.ospf`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--ospf--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--ospf--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv4--ospf--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.ospf.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv4--ospf--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.ospf.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv4--rip"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.rip`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--rip--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--rip--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv4--rip--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.rip.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv4--rip--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.rip.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv4--static"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.static`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--static--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--static--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv4--static--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.static.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv4--static--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv4.static.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |




<a id="nestedatt--protocols--isis--redistribute--ipv6"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6`

Optional:

- `babel` (Attributes) Redistribute Babel routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--babel))
- `bgp` (Attributes) Redistribute BGP routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--bgp))
- `connected` (Attributes) Redistribute connected routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--connected))
- `kernel` (Attributes) Redistribute kernel routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--kernel))
- `ospf6` (Attributes) Redistribute OSPFv3 routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ospf6))
- `ripng` (Attributes) Redistribute RIPng routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ripng))
- `static` (Attributes) Redistribute static routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--static))

<a id="nestedatt--protocols--isis--redistribute--ipv6--babel"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.babel`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--babel--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--babel--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv6--babel--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.babel.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv6--babel--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.babel.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv6--bgp"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.bgp`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--bgp--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--bgp--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv6--bgp--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.bgp.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv6--bgp--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.bgp.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv6--connected"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.connected`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--connected--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--connected--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv6--connected--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.connected.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv6--connected--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.connected.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv6--kernel"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.kernel`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--kernel--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--kernel--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv6--kernel--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.kernel.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv6--kernel--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.kernel.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv6--ospf6"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.ospf6`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ospf6--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ospf6--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv6--ospf6--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.ospf6.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv6--ospf6--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.ospf6.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv6--ripng"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.ripng`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ripng--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ripng--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv6--ripng--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.ripng.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv6--ripng--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.ripng.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--isis--redistribute--ipv6--static"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.static`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--static--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--static--level_2))

<a id="nestedatt--protocols--isis--redistribute--ipv6--static--level_1"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.static.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--isis--redistribute--ipv6--static--level_2"></a>
### Nested Schema for `protocols.isis.redistribute.ipv6.static.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    |  Format      &emsp;|  Description           |
    |--------------|------------------------|
    |  0-16777215  &emsp;|  Default metric value  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |





<a id="nestedatt--protocols--isis--segment_routing"></a>
### Nested Schema for `protocols.isis.segment_routing`

Optional:

- `global_block` (Attributes) Segment Routing Global Block label range (see [below for nested schema](#nestedatt--protocols--isis--segment_routing--global_block))
- `local_block` (Attributes) Segment Routing Local Block label range (see [below for nested schema](#nestedatt--protocols--isis--segment_routing--local_block))
- `maximum_label_depth` (Number) Maximum MPLS labels allowed for this router

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  1-16    &emsp;|  MPLS label depth  |

<a id="nestedatt--protocols--isis--segment_routing--global_block"></a>
### Nested Schema for `protocols.isis.segment_routing.global_block`

Optional:

- `high_label_value` (Number) MPLS label upper bound

    |  Format      &emsp;|  Description  |
    |--------------|---------------|
    |  16-1048575  &emsp;|  Label value  |
- `low_label_value` (Number) MPLS label lower bound

    |  Format      &emsp;|  Description                                   |
    |--------------|------------------------------------------------|
    |  16-1048575  &emsp;|  Label value (recommended minimum value: 300)  |


<a id="nestedatt--protocols--isis--segment_routing--local_block"></a>
### Nested Schema for `protocols.isis.segment_routing.local_block`

Optional:

- `high_label_value` (Number) MPLS label upper bound

    |  Format      &emsp;|  Description  |
    |--------------|---------------|
    |  16-1048575  &emsp;|  Label value  |
- `low_label_value` (Number) MPLS label lower bound

    |  Format      &emsp;|  Description                                   |
    |--------------|------------------------------------------------|
    |  16-1048575  &emsp;|  Label value (recommended minimum value: 300)  |



<a id="nestedatt--protocols--isis--spf_delay_ietf"></a>
### Nested Schema for `protocols.isis.spf_delay_ietf`

Optional:

- `holddown` (Number) Time with no received IGP events before considering IGP stable

    |  Format   &emsp;|  Description                                                           |
    |-----------|------------------------------------------------------------------------|
    |  0-60000  &emsp;|  Time with no received IGP events before considering IGP stable in ms  |
- `init_delay` (Number) Delay used while in QUIET state

    |  Format   &emsp;|  Description                              |
    |-----------|-------------------------------------------|
    |  0-60000  &emsp;|  Delay used while in QUIET state (in ms)  |
- `long_delay` (Number) Delay used while in LONG_WAIT

    |  Format   &emsp;|  Description                                |
    |-----------|---------------------------------------------|
    |  0-60000  &emsp;|  Delay used while in LONG_WAIT state in ms  |
- `short_delay` (Number) Delay used while in SHORT_WAIT state

    |  Format   &emsp;|  Description                                   |
    |-----------|------------------------------------------------|
    |  0-60000  &emsp;|  Delay used while in SHORT_WAIT state (in ms)  |
- `time_to_learn` (Number) Maximum duration needed to learn all the events related to a single failure

    |  Format   &emsp;|  Description                                                                        |
    |-----------|-------------------------------------------------------------------------------------|
    |  0-60000  &emsp;|  Maximum duration needed to learn all the events related to a single failure in ms  |


<a id="nestedatt--protocols--isis--traffic_engineering"></a>
### Nested Schema for `protocols.isis.traffic_engineering`

Optional:

- `address` (String) MPLS traffic engineering router ID

    |  Format  &emsp;|  Description   |
    |----------|----------------|
    |  ipv4    &emsp;|  IPv4 address  |
- `enable` (Boolean) Enable MPLS traffic engineering extensions



<a id="nestedatt--protocols--ospf"></a>
### Nested Schema for `protocols.ospf`

Optional:

- `aggregation` (Attributes) External route aggregation (see [below for nested schema](#nestedatt--protocols--ospf--aggregation))
- `auto_cost` (Attributes) Calculate interface cost according to bandwidth (see [below for nested schema](#nestedatt--protocols--ospf--auto_cost))
- `capability` (Attributes) Enable specific OSPF features (see [below for nested schema](#nestedatt--protocols--ospf--capability))
- `default_information` (Attributes) Default route advertisment settings (see [below for nested schema](#nestedatt--protocols--ospf--default_information))
- `default_metric` (Number) Metric of redistributed routes

    |  Format      &emsp;|  Description                     |
    |--------------|----------------------------------|
    |  0-16777214  &emsp;|  Metric of redistributed routes  |
- `distance` (Attributes) Administrative distance (see [below for nested schema](#nestedatt--protocols--ospf--distance))
- `graceful_restart` (Attributes) Graceful Restart (see [below for nested schema](#nestedatt--protocols--ospf--graceful_restart))
- `ldp_sync` (Attributes) Protocol wide LDP-IGP synchronization configuration (see [below for nested schema](#nestedatt--protocols--ospf--ldp_sync))
- `log_adjacency_changes` (Attributes) Log adjacency state changes (see [below for nested schema](#nestedatt--protocols--ospf--log_adjacency_changes))
- `max_metric` (Attributes) OSPF maximum and infinite-distance metric (see [below for nested schema](#nestedatt--protocols--ospf--max_metric))
- `maximum_paths` (Number) Maximum multiple paths (ECMP)

    |  Format  &emsp;|  Description                    |
    |----------|---------------------------------|
    |  1-64    &emsp;|  Maximum multiple paths (ECMP)  |
- `mpls_te` (Attributes) MultiProtocol Label Switching-Traffic Engineering (MPLS-TE) parameters (see [below for nested schema](#nestedatt--protocols--ospf--mpls_te))
- `parameters` (Attributes) OSPF specific parameters (see [below for nested schema](#nestedatt--protocols--ospf--parameters))
- `passive_interface` (String) Suppress routing updates on an interface

    |  Format   &emsp;|  Description                                            |
    |-----------|---------------------------------------------------------|
    |  default  &emsp;|  Default to suppress routing updates on all interfaces  |
- `redistribute` (Attributes) Redistribute information from another routing protocol (see [below for nested schema](#nestedatt--protocols--ospf--redistribute))
- `refresh` (Attributes) Adjust refresh parameters (see [below for nested schema](#nestedatt--protocols--ospf--refresh))
- `segment_routing` (Attributes) Segment-Routing (SPRING) settings (see [below for nested schema](#nestedatt--protocols--ospf--segment_routing))
- `timers` (Attributes) Adjust routing timers (see [below for nested schema](#nestedatt--protocols--ospf--timers))

<a id="nestedatt--protocols--ospf--aggregation"></a>
### Nested Schema for `protocols.ospf.aggregation`

Optional:

- `timer` (Number) Delay timer

    |  Format  &emsp;|  Description                |
    |----------|-----------------------------|
    |  5-1800  &emsp;|  Timer interval in seconds  |


<a id="nestedatt--protocols--ospf--auto_cost"></a>
### Nested Schema for `protocols.ospf.auto_cost`

Optional:

- `reference_bandwidth` (Number) Reference bandwidth method to assign cost

    |  Format     &emsp;|  Description                            |
    |-------------|-----------------------------------------|
    |  1-4294967  &emsp;|  Reference bandwidth cost in Mbits/sec  |


<a id="nestedatt--protocols--ospf--capability"></a>
### Nested Schema for `protocols.ospf.capability`

Optional:

- `opaque` (Boolean) Opaque LSA


<a id="nestedatt--protocols--ospf--default_information"></a>
### Nested Schema for `protocols.ospf.default_information`

Optional:

- `originate` (Attributes) Distribute a default route (see [below for nested schema](#nestedatt--protocols--ospf--default_information--originate))

<a id="nestedatt--protocols--ospf--default_information--originate"></a>
### Nested Schema for `protocols.ospf.default_information.originate`

Optional:

- `always` (Boolean) Always advertise a default route
- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--ospf--distance"></a>
### Nested Schema for `protocols.ospf.distance`

Optional:

- `global` (Number) Administrative distance

    |  Format  &emsp;|  Description              |
    |----------|---------------------------|
    |  1-255   &emsp;|  Administrative distance  |
- `ospf` (Attributes) OSPF administrative distance (see [below for nested schema](#nestedatt--protocols--ospf--distance--ospf))

<a id="nestedatt--protocols--ospf--distance--ospf"></a>
### Nested Schema for `protocols.ospf.distance.ospf`

Optional:

- `external` (Number) Distance for external routes

    |  Format  &emsp;|  Description                   |
    |----------|--------------------------------|
    |  1-255   &emsp;|  Distance for external routes  |
- `inter_area` (Number) Distance for inter-area routes

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  1-255   &emsp;|  Distance for inter-area routes  |
- `intra_area` (Number) Distance for intra-area routes

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  1-255   &emsp;|  Distance for intra-area routes  |



<a id="nestedatt--protocols--ospf--graceful_restart"></a>
### Nested Schema for `protocols.ospf.graceful_restart`

Optional:

- `grace_period` (Number) Maximum length of the grace period

    |  Format  &emsp;|  Description                                    |
    |----------|-------------------------------------------------|
    |  1-1800  &emsp;|  Maximum length of the grace period in seconds  |
- `helper` (Attributes) OSPF graceful-restart helpers (see [below for nested schema](#nestedatt--protocols--ospf--graceful_restart--helper))

<a id="nestedatt--protocols--ospf--graceful_restart--helper"></a>
### Nested Schema for `protocols.ospf.graceful_restart.helper`

Optional:

- `enable` (Attributes) Enable helper support (see [below for nested schema](#nestedatt--protocols--ospf--graceful_restart--helper--enable))
- `no_strict_lsa_checking` (Boolean) Disable strict LSA check
- `planned_only` (Boolean) Supported only planned restart
- `supported_grace_time` (Number) Supported grace timer

    |  Format   &emsp;|  Description                |
    |-----------|-----------------------------|
    |  10-1800  &emsp;|  Grace interval in seconds  |

<a id="nestedatt--protocols--ospf--graceful_restart--helper--enable"></a>
### Nested Schema for `protocols.ospf.graceful_restart.helper.enable`

Optional:

- `router_id` (List of String) Advertising Router-ID

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  ipv4    &emsp;|  Router-ID in IP address format  |




<a id="nestedatt--protocols--ospf--ldp_sync"></a>
### Nested Schema for `protocols.ospf.ldp_sync`

Optional:

- `holddown` (Number) Hold down timer for LDP-IGP cost restoration

    |  Format   &emsp;|  Description                                                                                   |
    |-----------|------------------------------------------------------------------------------------------------|
    |  0-10000  &emsp;|  Time to wait in seconds for LDP-IGP synchronization to occur before restoring interface cost  |


<a id="nestedatt--protocols--ospf--log_adjacency_changes"></a>
### Nested Schema for `protocols.ospf.log_adjacency_changes`

Optional:

- `detail` (Boolean) Log all state changes


<a id="nestedatt--protocols--ospf--max_metric"></a>
### Nested Schema for `protocols.ospf.max_metric`

Optional:

- `router_lsa` (Attributes) Advertise own Router-LSA with infinite distance (stub router) (see [below for nested schema](#nestedatt--protocols--ospf--max_metric--router_lsa))

<a id="nestedatt--protocols--ospf--max_metric--router_lsa"></a>
### Nested Schema for `protocols.ospf.max_metric.router_lsa`

Optional:

- `administrative` (Boolean) Administratively apply, for an indefinite period
- `on_shutdown` (Number) Advertise stub-router prior to full shutdown of OSPF

    |  Format  &emsp;|  Description                                      |
    |----------|---------------------------------------------------|
    |  5-100   &emsp;|  Time (seconds) to advertise self as stub-router  |
- `on_startup` (Number) Automatically advertise stub Router-LSA on startup of OSPF

    |  Format   &emsp;|  Description                                      |
    |-----------|---------------------------------------------------|
    |  5-86400  &emsp;|  Time (seconds) to advertise self as stub-router  |



<a id="nestedatt--protocols--ospf--mpls_te"></a>
### Nested Schema for `protocols.ospf.mpls_te`

Optional:

- `enable` (Boolean) Enable MPLS-TE functionality
- `router_address` (String) Stable IP address of the advertising router

    |  Format  &emsp;|  Description                                  |
    |----------|-----------------------------------------------|
    |  ipv4    &emsp;|  Stable IP address of the advertising router  |


<a id="nestedatt--protocols--ospf--parameters"></a>
### Nested Schema for `protocols.ospf.parameters`

Optional:

- `abr_type` (String) OSPF ABR type

    |  Format    &emsp;|  Description        |
    |------------|---------------------|
    |  cisco     &emsp;|  Cisco ABR type     |
    |  ibm       &emsp;|  IBM ABR type       |
    |  shortcut  &emsp;|  Shortcut ABR type  |
    |  standard  &emsp;|  Standard ABR type  |
- `opaque_lsa` (Boolean) Enable the Opaque-LSA capability (rfc2370)
- `rfc1583_compatibility` (Boolean) Enable RFC1583 criteria for handling AS external routes
- `router_id` (String) Override default router identifier

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  ipv4    &emsp;|  Router-ID in IP address format  |


<a id="nestedatt--protocols--ospf--redistribute"></a>
### Nested Schema for `protocols.ospf.redistribute`

Optional:

- `babel` (Attributes) Redistribute Babel routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--babel))
- `bgp` (Attributes) Redistribute BGP routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--bgp))
- `connected` (Attributes) Redistribute connected routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--connected))
- `isis` (Attributes) Redistribute IS-IS routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--isis))
- `kernel` (Attributes) Redistribute Kernel routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--kernel))
- `rip` (Attributes) Redistribute RIP routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--rip))
- `static` (Attributes) Redistribute statically configured routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--static))

<a id="nestedatt--protocols--ospf--redistribute--babel"></a>
### Nested Schema for `protocols.ospf.redistribute.babel`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospf--redistribute--bgp"></a>
### Nested Schema for `protocols.ospf.redistribute.bgp`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospf--redistribute--connected"></a>
### Nested Schema for `protocols.ospf.redistribute.connected`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospf--redistribute--isis"></a>
### Nested Schema for `protocols.ospf.redistribute.isis`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospf--redistribute--kernel"></a>
### Nested Schema for `protocols.ospf.redistribute.kernel`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospf--redistribute--rip"></a>
### Nested Schema for `protocols.ospf.redistribute.rip`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospf--redistribute--static"></a>
### Nested Schema for `protocols.ospf.redistribute.static`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--ospf--refresh"></a>
### Nested Schema for `protocols.ospf.refresh`

Optional:

- `timers` (Number) Refresh timer

    |  Format   &emsp;|  Description             |
    |-----------|--------------------------|
    |  10-1800  &emsp;|  Timer value in seconds  |


<a id="nestedatt--protocols--ospf--segment_routing"></a>
### Nested Schema for `protocols.ospf.segment_routing`

Optional:

- `global_block` (Attributes) Segment Routing Global Block label range (see [below for nested schema](#nestedatt--protocols--ospf--segment_routing--global_block))
- `local_block` (Attributes) Segment Routing Local Block label range (see [below for nested schema](#nestedatt--protocols--ospf--segment_routing--local_block))
- `maximum_label_depth` (Number) Maximum MPLS labels allowed for this router

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  1-16    &emsp;|  MPLS label depth  |

<a id="nestedatt--protocols--ospf--segment_routing--global_block"></a>
### Nested Schema for `protocols.ospf.segment_routing.global_block`

Optional:

- `high_label_value` (Number) MPLS label upper bound

    |  Format      &emsp;|  Description  |
    |--------------|---------------|
    |  16-1048575  &emsp;|  Label value  |
- `low_label_value` (Number) MPLS label lower bound

    |  Format      &emsp;|  Description                                   |
    |--------------|------------------------------------------------|
    |  16-1048575  &emsp;|  Label value (recommended minimum value: 300)  |


<a id="nestedatt--protocols--ospf--segment_routing--local_block"></a>
### Nested Schema for `protocols.ospf.segment_routing.local_block`

Optional:

- `high_label_value` (Number) MPLS label upper bound

    |  Format      &emsp;|  Description  |
    |--------------|---------------|
    |  16-1048575  &emsp;|  Label value  |
- `low_label_value` (Number) MPLS label lower bound

    |  Format      &emsp;|  Description                                   |
    |--------------|------------------------------------------------|
    |  16-1048575  &emsp;|  Label value (recommended minimum value: 300)  |



<a id="nestedatt--protocols--ospf--timers"></a>
### Nested Schema for `protocols.ospf.timers`

Optional:

- `throttle` (Attributes) Throttling adaptive timers (see [below for nested schema](#nestedatt--protocols--ospf--timers--throttle))

<a id="nestedatt--protocols--ospf--timers--throttle"></a>
### Nested Schema for `protocols.ospf.timers.throttle`

Optional:

- `spf` (Attributes) OSPF SPF timers (see [below for nested schema](#nestedatt--protocols--ospf--timers--throttle--spf))

<a id="nestedatt--protocols--ospf--timers--throttle--spf"></a>
### Nested Schema for `protocols.ospf.timers.throttle.spf`

Optional:

- `delay` (Number) Delay from the first change received to SPF calculation

    |  Format    &emsp;|  Description            |
    |------------|-------------------------|
    |  0-600000  &emsp;|  Delay in milliseconds  |
- `initial_holdtime` (Number) Initial hold time between consecutive SPF calculations

    |  Format    &emsp;|  Description                        |
    |------------|-------------------------------------|
    |  0-600000  &emsp;|  Initial hold time in milliseconds  |
- `max_holdtime` (Number) Maximum hold time

    |  Format    &emsp;|  Description                    |
    |------------|---------------------------------|
    |  0-600000  &emsp;|  Max hold time in milliseconds  |





<a id="nestedatt--protocols--ospfv3"></a>
### Nested Schema for `protocols.ospfv3`

Optional:

- `auto_cost` (Attributes) Calculate interface cost according to bandwidth (see [below for nested schema](#nestedatt--protocols--ospfv3--auto_cost))
- `default_information` (Attributes) Default route advertisment settings (see [below for nested schema](#nestedatt--protocols--ospfv3--default_information))
- `distance` (Attributes) Administrative distance (see [below for nested schema](#nestedatt--protocols--ospfv3--distance))
- `graceful_restart` (Attributes) Graceful Restart (see [below for nested schema](#nestedatt--protocols--ospfv3--graceful_restart))
- `log_adjacency_changes` (Attributes) Log adjacency state changes (see [below for nested schema](#nestedatt--protocols--ospfv3--log_adjacency_changes))
- `parameters` (Attributes) OSPFv3 specific parameters (see [below for nested schema](#nestedatt--protocols--ospfv3--parameters))
- `redistribute` (Attributes) Redistribute information from another routing protocol (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute))

<a id="nestedatt--protocols--ospfv3--auto_cost"></a>
### Nested Schema for `protocols.ospfv3.auto_cost`

Optional:

- `reference_bandwidth` (Number) Reference bandwidth method to assign cost

    |  Format     &emsp;|  Description                            |
    |-------------|-----------------------------------------|
    |  1-4294967  &emsp;|  Reference bandwidth cost in Mbits/sec  |


<a id="nestedatt--protocols--ospfv3--default_information"></a>
### Nested Schema for `protocols.ospfv3.default_information`

Optional:

- `originate` (Attributes) Distribute a default route (see [below for nested schema](#nestedatt--protocols--ospfv3--default_information--originate))

<a id="nestedatt--protocols--ospfv3--default_information--originate"></a>
### Nested Schema for `protocols.ospfv3.default_information.originate`

Optional:

- `always` (Boolean) Always advertise a default route
- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |



<a id="nestedatt--protocols--ospfv3--distance"></a>
### Nested Schema for `protocols.ospfv3.distance`

Optional:

- `global` (Number) Administrative distance

    |  Format  &emsp;|  Description              |
    |----------|---------------------------|
    |  1-255   &emsp;|  Administrative distance  |
- `ospfv3` (Attributes) OSPFv3 administrative distance (see [below for nested schema](#nestedatt--protocols--ospfv3--distance--ospfv3))

<a id="nestedatt--protocols--ospfv3--distance--ospfv3"></a>
### Nested Schema for `protocols.ospfv3.distance.ospfv3`

Optional:

- `external` (Number) Distance for external routes

    |  Format  &emsp;|  Description                   |
    |----------|--------------------------------|
    |  1-255   &emsp;|  Distance for external routes  |
- `inter_area` (Number) Distance for inter-area routes

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  1-255   &emsp;|  Distance for inter-area routes  |
- `intra_area` (Number) Distance for intra-area routes

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  1-255   &emsp;|  Distance for intra-area routes  |



<a id="nestedatt--protocols--ospfv3--graceful_restart"></a>
### Nested Schema for `protocols.ospfv3.graceful_restart`

Optional:

- `grace_period` (Number) Maximum length of the grace period

    |  Format  &emsp;|  Description                                    |
    |----------|-------------------------------------------------|
    |  1-1800  &emsp;|  Maximum length of the grace period in seconds  |
- `helper` (Attributes) OSPF graceful-restart helpers (see [below for nested schema](#nestedatt--protocols--ospfv3--graceful_restart--helper))

<a id="nestedatt--protocols--ospfv3--graceful_restart--helper"></a>
### Nested Schema for `protocols.ospfv3.graceful_restart.helper`

Optional:

- `enable` (Attributes) Enable helper support (see [below for nested schema](#nestedatt--protocols--ospfv3--graceful_restart--helper--enable))
- `lsa_check_disable` (Boolean) Disable strict LSA check
- `planned_only` (Boolean) Supported only planned restart
- `supported_grace_time` (Number) Supported grace timer

    |  Format   &emsp;|  Description                |
    |-----------|-----------------------------|
    |  10-1800  &emsp;|  Grace interval in seconds  |

<a id="nestedatt--protocols--ospfv3--graceful_restart--helper--enable"></a>
### Nested Schema for `protocols.ospfv3.graceful_restart.helper.enable`

Optional:

- `router_id` (List of String) Advertising Router-ID

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  ipv4    &emsp;|  Router-ID in IP address format  |




<a id="nestedatt--protocols--ospfv3--log_adjacency_changes"></a>
### Nested Schema for `protocols.ospfv3.log_adjacency_changes`

Optional:

- `detail` (Boolean) Log all state changes


<a id="nestedatt--protocols--ospfv3--parameters"></a>
### Nested Schema for `protocols.ospfv3.parameters`

Optional:

- `router_id` (String) Override default router identifier

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  ipv4    &emsp;|  Router-ID in IP address format  |


<a id="nestedatt--protocols--ospfv3--redistribute"></a>
### Nested Schema for `protocols.ospfv3.redistribute`

Optional:

- `babel` (Attributes) Redistribute Babel routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--babel))
- `bgp` (Attributes) Redistribute BGP routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--bgp))
- `connected` (Attributes) Redistribute connected routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--connected))
- `isis` (Attributes) Redistribute IS-IS routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--isis))
- `kernel` (Attributes) Redistribute kernel routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--kernel))
- `ripng` (Attributes) Redistribute RIPNG routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--ripng))
- `static` (Attributes) Redistribute static routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--static))

<a id="nestedatt--protocols--ospfv3--redistribute--babel"></a>
### Nested Schema for `protocols.ospfv3.redistribute.babel`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospfv3--redistribute--bgp"></a>
### Nested Schema for `protocols.ospfv3.redistribute.bgp`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospfv3--redistribute--connected"></a>
### Nested Schema for `protocols.ospfv3.redistribute.connected`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospfv3--redistribute--isis"></a>
### Nested Schema for `protocols.ospfv3.redistribute.isis`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospfv3--redistribute--kernel"></a>
### Nested Schema for `protocols.ospfv3.redistribute.kernel`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospfv3--redistribute--ripng"></a>
### Nested Schema for `protocols.ospfv3.redistribute.ripng`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |


<a id="nestedatt--protocols--ospfv3--redistribute--static"></a>
### Nested Schema for `protocols.ospfv3.redistribute.static`

Optional:

- `metric` (Number) OSPF default metric

    |  Format      &emsp;|  Description     |
    |--------------|------------------|
    |  0-16777214  &emsp;|  Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    |  Format  &emsp;|  Description                         |
    |----------|--------------------------------------|
    |  1-2     &emsp;|  Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Route map name  |





<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_vrf_name.example "vrf__name__<name>"
```
