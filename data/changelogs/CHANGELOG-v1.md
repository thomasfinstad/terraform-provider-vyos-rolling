
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 1.6.202408210 (2024-08-21 08-30-09 UTC)](#release-16202408210-2024-08-21-08-30-09-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
    - [Schema changes](#schema-changes)
      - [Features](#features)
        - [Resources](#resources)
  - [Release 1.5.202408170 (2024-08-17 08-23-43 UTC)](#release-15202408170-2024-08-17-08-23-43-utc)
    - [Project changes](#project-changes-1)
      - [Notes](#notes-1)
    - [Schema changes](#schema-changes-1)
      - [Features](#features-1)
        - [Resources](#resources-1)
  - [Release 1.4.202408160 (2024-08-16 16-29-03 UTC)](#release-14202408160-2024-08-16-16-29-03-utc)
    - [Project changes](#project-changes-2)
      - [Notes](#notes-2)
  - [Release 1.4.202408152 (2024-08-15 17-55-10 UTC)](#release-14202408152-2024-08-15-17-55-10-utc)
    - [Project changes](#project-changes-3)
      - [Notes](#notes-3)
      - [Bug fixes](#bug-fixes)
  - [Release 1.4.202408151 (2024-08-15 16-16-24 UTC)](#release-14202408151-2024-08-15-16-16-24-utc)
    - [Project changes](#project-changes-4)
      - [Notes](#notes-4)
      - [Bug fixes](#bug-fixes-1)
    - [Schema changes](#schema-changes-2)
      - [Notes](#notes-5)
        - [Resources](#resources-2)
  - [Release 1.4.202408150 (2024-08-15 11-31-13 UTC)](#release-14202408150-2024-08-15-11-31-13-utc)
    - [Project changes](#project-changes-5)
      - [Notes](#notes-6)
      - [Features](#features-2)
      - [Bug fixes](#bug-fixes-2)
    - [Schema changes](#schema-changes-3)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources-3)
      - [Features](#features-3)
        - [Resources](#resources-4)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 1.6.202408210 (2024-08-21 08-30-09 UTC)
### Project changes
#### Notes
* update to rolling release 2024-08-21T00:20:57Z

### Schema changes
#### Features

##### Resources
* Modified Resource `vyos_interfaces_openvpn`
	* New attribute `ip_version`









## Release 1.5.202408170 (2024-08-17 08-23-43 UTC)
### Project changes
#### Notes
* update to rolling release 2024-08-17T00:20:00Z

### Schema changes
#### Features

##### Resources
* Modified Resource `vyos_service_ipoe_server_interface`
	* New attribute `vlan_mon`

* Modified Resource `vyos_service_pppoe_server_interface`
	* New attribute `vlan_mon`









## Release 1.4.202408160 (2024-08-16 16-29-03 UTC)
### Project changes
#### Notes
* update to rolling release 2024-08-16T00:20:42Z
* improve change detection and build target during ci run


## Release 1.4.202408152 (2024-08-15 17-55-10 UTC)
### Project changes
#### Notes
* Fix load balancer resources subcategory
#### Bug fixes
* Makefile shell substitution and git cmd error


## Release 1.4.202408151 (2024-08-15 16-16-24 UTC)
### Project changes
#### Notes
* improve documentation with extra focus on adding better table of contents
#### Bug fixes
* remove redundant and error prone pre-commit run

### Schema changes
#### Notes

##### Resources
* Modified Resource `vyos_firewall_ipv4_output_raw_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_firewall_bridge_forward_filter_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_firewall_ipv6_name_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_system_conntrack_timeout_custom_ipv4_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_firewall_ipv4_prerouting_raw_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_firewall_ipv6_prerouting_raw_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_firewall_ipv6_output_raw_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_firewall_bridge_output_filter_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_firewall_bridge_input_filter_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_nat_destination_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_firewall_ipv6_output_filter_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_system_conntrack_ignore_ipv6_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_system_conntrack_timeout_custom_ipv6_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_firewall_ipv4_name_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_policy_route_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_firewall_bridge_name_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_system_conntrack_ignore_ipv4_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_firewall_bridge_prerouting_filter_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_firewall_ipv4_forward_filter_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_firewall_ipv6_forward_filter_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_policy_route6_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_firewall_ipv6_input_filter_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_firewall_ipv4_output_filter_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_load_balancing_wan_rule`
	* Modified attribute `source.port.` changed description
	* Modified attribute `destination.port.` changed description

* Modified Resource `vyos_firewall_ipv4_input_filter_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description

* Modified Resource `vyos_nat_source_rule`
	* Modified attribute `destination.port.` changed description
	* Modified attribute `source.port.` changed description









## Release 1.4.202408150 (2024-08-15 11-31-13 UTC)
### Project changes
#### Notes
* Create changelog archiving
* update to rolling release 2024-08-15T00:20:15Z
* improve provider documentation readability
* improve versioning and changelog generation
* update to rolling release 2024-08-14T00:20:52Z
#### Features
* reenable generation of missing resources and move resource identifying information into dedicated parameter
#### Bug fixes
* base changelog generation on previous tag instead of previous commit
* fixed error in versioning
* changelog and release procedures

### Schema changes
#### BREAKING CHANGES

##### Resources
* Modified Resource `vyos_high_availability_vrrp_group_address`
	* **Removed attribute** `address_id`
	* **Removed attribute** `group_id`
	* New attribute `identifier`

* Modified Resource `vyos_netns_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_community_list_rule`
	* **Removed attribute** `community_list_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_access_list`
	* **Removed attribute** `name_id`
	* **Removed attribute** `access_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route6`
	* **Removed attribute** `name_id`
	* **Removed attribute** `route6_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_interface_group`
	* **Removed attribute** `interface_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_prerouting_raw_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_ca`
	* **Removed attribute** `ca_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_round_robin_class_match`
	* **Removed attribute** `match_id`
	* **Removed attribute** `round_robin_id`
	* **Removed attribute** `class_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_port`
	* **Removed attribute** `name_id`
	* **Removed attribute** `port_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_area_virtual_link_authentication_md5_key_id`
	* **Removed attribute** `virtual_link_id`
	* **Removed attribute** `area_id`
	* **Removed attribute** `key_id_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_forward_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_name_rule`
	* **Removed attribute** `name_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_nat_source_rule_load_balance_backend`
	* **Removed attribute** `backend_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_access_list_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `access_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_drop_tail`
	* **Removed attribute** `drop_tail_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_bmp_target`
	* **Removed attribute** `name_id`
	* **Removed attribute** `target_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_dynamic_group_address_group`
	* **Removed attribute** `address_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_vrrp_sync_group`
	* **Removed attribute** `sync_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_prefix_list`
	* **Removed attribute** `prefix_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route6_next_hop`
	* **Removed attribute** `next_hop_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `route6_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_network`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_limiter`
	* **Removed attribute** `limiter_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_network_emulator`
	* **Removed attribute** `network_emulator_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_rate_control`
	* **Removed attribute** `rate_control_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper_hfsc_class`
	* **Removed attribute** `class_id`
	* **Removed attribute** `shaper_hfsc_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_labeled_unicast_aggregate_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `aggregate_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_peer_group_local_as`
	* **Removed attribute** `local_as_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `peer_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_dh`
	* **Removed attribute** `dh_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_prefix_list_rule`
	* **Removed attribute** `prefix_list_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper_hfsc_class_match`
	* **Removed attribute** `shaper_hfsc_id`
	* **Removed attribute** `class_id`
	* **Removed attribute** `match_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_traffic_match_group`
	* **Removed attribute** `traffic_match_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospfv3_interface`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_input_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_prerouting_raw_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_access_list`
	* **Removed attribute** `access_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_priority_queue_class_match`
	* **Removed attribute** `priority_queue_id`
	* **Removed attribute** `class_id`
	* **Removed attribute** `match_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_unicast_aggregate_address`
	* **Removed attribute** `aggregate_address_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_neighbor`
	* **Removed attribute** `neighbor_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_peer_group_local_role`
	* **Removed attribute** `local_role_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `peer_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_ipv6_address_group`
	* **Removed attribute** `ipv6_address_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route_interface`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `route_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_area_range`
	* **Removed attribute** `name_id`
	* **Removed attribute** `range_id`
	* **Removed attribute** `area_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_parameters_distance_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_node_protecting_index`
	* **Removed attribute** `index_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_environment`
	* **Removed attribute** `name_id`
	* **Removed attribute** `environment_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_traffic_match_group_match`
	* **Removed attribute** `match_id`
	* **Removed attribute** `traffic_match_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_vpn_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_registry`
	* **Removed attribute** `registry_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_prerouting_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_domain_group`
	* **Removed attribute** `domain_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_network_group`
	* **Removed attribute** `network_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_nat_destination_rule_load_balance_backend`
	* **Removed attribute** `backend_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_community_list`
	* **Removed attribute** `community_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_input_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_random_detect_precedence`
	* **Removed attribute** `precedence_id`
	* **Removed attribute** `random_detect_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_aggregate_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `aggregate_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_vpn_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_route_map_rule`
	* **Removed attribute** `route_map_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_large_community_list_rule`
	* **Removed attribute** `large_community_list_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_area_virtual_link`
	* **Removed attribute** `name_id`
	* **Removed attribute** `virtual_link_id`
	* **Removed attribute** `area_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_output_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_mac_group`
	* **Removed attribute** `mac_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_as_path_list`
	* **Removed attribute** `as_path_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper_class_match`
	* **Removed attribute** `shaper_id`
	* **Removed attribute** `class_id`
	* **Removed attribute** `match_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_aggregate_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `aggregate_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route_next_hop_bfd_multi_hop_source`
	* **Removed attribute** `name_id`
	* **Removed attribute** `next_hop_id`
	* **Removed attribute** `route_id`
	* **Removed attribute** `source_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_fq_codel`
	* **Removed attribute** `fq_codel_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospfv3_area_range`
	* **Removed attribute** `range_id`
	* **Removed attribute** `area_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_forward_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_zone_from`
	* **Removed attribute** `from_id`
	* **Removed attribute** `zone_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_priority_queue`
	* **Removed attribute** `priority_queue_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_random_detect`
	* **Removed attribute** `random_detect_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_l2vpn_evpn_vni`
	* **Removed attribute** `name_id`
	* **Removed attribute** `vni_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_segment_routing_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_ipv6_network_group`
	* **Removed attribute** `ipv6_network_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_key_pair`
	* **Removed attribute** `key_pair_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_extcommunity_list`
	* **Removed attribute** `extcommunity_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_distance_prefix`
	* **Removed attribute** `prefix_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_forward_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_nat_source_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_certificate`
	* **Removed attribute** `certificate_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_openvpn_shared_secret`
	* **Removed attribute** `shared_secret_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_round_robin`
	* **Removed attribute** `round_robin_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_unicast_distance_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_volume`
	* **Removed attribute** `name_id`
	* **Removed attribute** `volume_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_dynamic_group_ipv6_address_group`
	* **Removed attribute** `ipv6_address_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_port_group`
	* **Removed attribute** `port_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_output_raw_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_large_community_list`
	* **Removed attribute** `large_community_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_limiter_class`
	* **Removed attribute** `limiter_id`
	* **Removed attribute** `class_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_summary_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `summary_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_label`
	* **Removed attribute** `label_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_nat_static_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_extcommunity_list_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `extcommunity_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_cake`
	* **Removed attribute** `cake_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper`
	* **Removed attribute** `shaper_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_labeled_unicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_nat_destination_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_priority_queue_class`
	* **Removed attribute** `priority_queue_id`
	* **Removed attribute** `class_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_labeled_unicast_aggregate_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `aggregate_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_downstream_index`
	* **Removed attribute** `index_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route`
	* **Removed attribute** `name_id`
	* **Removed attribute** `route_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route6_next_hop_bfd_multi_hop_source`
	* **Removed attribute** `name_id`
	* **Removed attribute** `next_hop_id`
	* **Removed attribute** `route6_id`
	* **Removed attribute** `source_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_sysctl_parameter`
	* **Removed attribute** `name_id`
	* **Removed attribute** `parameter_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_prefix_list6`
	* **Removed attribute** `prefix_list6_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper_class`
	* **Removed attribute** `shaper_id`
	* **Removed attribute** `class_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_distance_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_output_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_output_raw_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_name_rule`
	* **Removed attribute** `name_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_address_group`
	* **Removed attribute** `address_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_device`
	* **Removed attribute** `name_id`
	* **Removed attribute** `device_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_vrrp_group`
	* **Removed attribute** `group_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_vrrp_group_excluded_address`
	* **Removed attribute** `excluded_address_id`
	* **Removed attribute** `group_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_openssh`
	* **Removed attribute** `openssh_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_access_list6_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `access_list6_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_as_path_list_rule`
	* **Removed attribute** `as_path_list_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_ip_protocol`
	* **Removed attribute** `name_id`
	* **Removed attribute** `protocol_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_lowest_backup_metric_index`
	* **Removed attribute** `name_id`
	* **Removed attribute** `index_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_virtual_server`
	* **Removed attribute** `virtual_server_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_segment_routing_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_output_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_prefix_list6_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `prefix_list6_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_limiter_class_match`
	* **Removed attribute** `class_id`
	* **Removed attribute** `limiter_id`
	* **Removed attribute** `match_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper_hfsc`
	* **Removed attribute** `shaper_hfsc_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_route_map`
	* **Removed attribute** `route_map_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_ipv6_protocol`
	* **Removed attribute** `name_id`
	* **Removed attribute** `protocol_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_labeled_unicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_neighbor_local_role`
	* **Removed attribute** `local_role_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `neighbor_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_area`
	* **Removed attribute** `area_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_fair_queue`
	* **Removed attribute** `fair_queue_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_round_robin_class`
	* **Removed attribute** `round_robin_id`
	* **Removed attribute** `class_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_neighbor_local_as`
	* **Removed attribute** `local_as_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `neighbor_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_peer_group`
	* **Removed attribute** `name_id`
	* **Removed attribute** `peer_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_interface_authentication_md5_key_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `key_id_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_flowtable`
	* **Removed attribute** `flowtable_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_interface`
	* **Removed attribute** `name_id`
	* **Removed attribute** `interface_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_neighbor`
	* **Removed attribute** `name_id`
	* **Removed attribute** `neighbor_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospfv3_area`
	* **Removed attribute** `area_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_unicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_access_list6`
	* **Removed attribute** `access_list6_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_input_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_distance_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_interface`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* **Removed Resource** `vyos_container_name_network`

* Modified Resource `vyos_firewall_ipv6_name_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_interface`
	* **Removed attribute** `interface_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_remote_prefix_list`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_interface`
	* **Removed attribute** `name_id`
	* **Removed attribute** `interface_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_redistribute_table`
	* **Removed attribute** `name_id`
	* **Removed attribute** `table_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route6_interface`
	* **Removed attribute** `route6_id`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route_next_hop`
	* **Removed attribute** `route_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `next_hop_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_aggregate_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `aggregate_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_virtual_server_real_server`
	* **Removed attribute** `real_server_id`
	* **Removed attribute** `virtual_server_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_listen_range`
	* **Removed attribute** `name_id`
	* **Removed attribute** `range_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_zone`
	* **Removed attribute** `zone_id`
	* New attribute `identifier`





#### Features

##### Resources
* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_priority_limit_high`

* New Resource `vyos_protocols_isis_redistribute_ipv6_ospf6_level_2`

* New Resource `vyos_protocols_failover_route`

* New Resource `vyos_system_login_radius`

* New Resource `vyos_protocols_rip_redistribute_isis`

* New Resource `vyos_load_balancing_reverse_proxy`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_distance_prefix`

* New Resource `vyos_protocols_isis_redistribute_ipv4_static_level_2`

* New Resource `vyos_protocols_pim_ecmp`

* New Resource `vyos_policy_route6_rule`

* New Resource `vyos_protocols_bgp_address_family_ipv4_labeled_unicast_maximum_paths`

* New Resource `vyos_system_conntrack_ignore_ipv4_rule`

* New Resource `vyos_service_ipoe_server_snmp`

* New Resource `vyos_interfaces_bonding_vif_s_vif_c`

* New Resource `vyos_protocols_static_route6_next_hop_bfd_multi_hop_source`

* New Resource `vyos_vpn_pptp_remote_access_authentication_radius`

* New Resource `vyos_vpn_sstp_client_ipv6_pool`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn`

* New Resource `vyos_vpn_ipsec_log`

* New Resource `vyos_protocols_bgp_timers`

* New Resource `vyos_system_conntrack_log_event_new`

* New Resource `vyos_system_login_user_authentication_public_keys`

* New Resource `vyos_vpn_ipsec_site_to_site_peer_tunnel`

* New Resource `vyos_load_balancing_reverse_proxy_global_parameters_logging_facility`

* New Resource `vyos_load_balancing_wan_interface_health_test`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_ptr`

* New Resource `vyos_service_ipoe_server_authentication_radius_dynamic_author`

* New Resource `vyos_system_update_check`

* New Resource `vyos_interfaces_tunnel`

* New Resource `vyos_protocols_rip_timers`

* New Resource `vyos_protocols_isis_redistribute_ipv6_ospf6_level_1`

* New Resource `vyos_service_config_sync`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_mx_server`

* New Resource `vyos_service_pppoe_server_authentication`

* New Resource `vyos_vpn_pptp_remote_access_authentication`

* New Resource `vyos_protocols_ospf_redistribute_connected`

* New Resource `vyos_protocols_static_route`

* New Resource `vyos_service_dns_forwarding_domain`

* New Resource `vyos_service_dns_forwarding_name_server`

* New Resource `vyos_service_pppoe_server_snmp`

* New Resource `vyos_vpn_openconnect_listen_ports`

* New Resource `vyos_protocols_isis_redistribute_ipv4_connected_level_1`

* New Resource `vyos_protocols_static_arp_interface`

* New Resource `vyos_protocols_static_table_route`

* New Resource `vyos_service_lldp`

* New Resource `vyos_service_ssh_access_control_deny`

* New Resource `vyos_service_webproxy_url_filtering_squidguard`

* New Resource `vyos_vpn_sstp_authentication_radius_dynamic_author`

* New Resource `vyos_protocols_isis_area_password`

* New Resource `vyos_vpn_openconnect_authentication_local_users_username`

* New Resource `vyos_interfaces_macsec_dhcpv6_options_pd_interface`

* New Resource `vyos_load_balancing_reverse_proxy_backend_logging_facility`

* New Resource `vyos_system_option_http_client`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet_range`

* New Resource `vyos_service_stunnel_server_psk`

* New Resource `vyos_interfaces_pseudo_ethernet_vif`

* New Resource `vyos_service_snmp_v3_trap_target`

* New Resource `vyos_protocols_bgp_bmp`

* New Resource `vyos_protocols_ospf_segment_routing_prefix`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s_dhcpv6_options_pd`

* New Resource `vyos_service_pppoe_server_limits`

* New Resource `vyos_protocols_isis_redistribute_ipv6_static_level_2`

* New Resource `vyos_protocols_ospf_area_virtual_link`

* New Resource `vyos_protocols_static_table_route_next_hop_bfd_multi_hop_source`

* New Resource `vyos_system_sysctl_parameter`

* New Resource `vyos_interfaces_macsec`

* New Resource `vyos_protocols_failover_route_next_hop`

* New Resource `vyos_service_dns_forwarding_authoritative_domain`

* New Resource `vyos_vpn_sstp_limits`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_nexthop_vpn`

* New Resource `vyos_service_dhcp_server_shared_network_name_option_static_route`

* New Resource `vyos_system_login_radius_server`

* New Resource `vyos_protocols_isis_redistribute_ipv6_ripng_level_2`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_kernel`

* New Resource `vyos_protocols_pim_igmp`

* New Resource `vyos_interfaces_bonding_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_wwan`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_route_target_vpn`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_advertise_ipv6_unicast`

* New Resource `vyos_interfaces_ethernet_vif`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s`

* New Resource `vyos_protocols_isis_interface`

* New Resource `vyos_service_event_handler_event_script_environment`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_spf`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_a`

* New Resource `vyos_protocols_static_multicast_interface_route_next_hop_interface`

* New Resource `vyos_service_event_handler_event`

* New Resource `vyos_service_ipoe_server`

* New Resource `vyos_service_ssh_dynamic_protection`

* New Resource `vyos_protocols_ripng_redistribute_connected`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet`

* New Resource `vyos_protocols_rip`

* New Resource `vyos_system_syslog_global`

* New Resource `vyos_interfaces_wireless_vif_s_vif_c_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_babel_redistribute_ipv6`

* New Resource `vyos_protocols_bgp_address_family_ipv6_labeled_unicast_network`

* New Resource `vyos_protocols_bgp_address_family_ipv6_vpn_network`

* New Resource `vyos_protocols_isis_redistribute_ipv6_bgp_level_2`

* New Resource `vyos_protocols_ospf_distance`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_mx`

* New Resource `vyos_protocols_isis_redistribute_ipv4_rip_level_2`

* New Resource `vyos_protocols_ospf_area_virtual_link_authentication_md5_key_id`

* New Resource `vyos_vpn_ipsec_remote_access_connection_authentication_local_users_username`

* New Resource `vyos_vpn_pptp_remote_access_authentication_radius_rate_limit`

* New Resource `vyos_protocols_bgp_address_family_ipv4_multicast_aggregate_address`

* New Resource `vyos_protocols_isis_redistribute_ipv4_ospf_level_1`

* New Resource `vyos_vpn_pptp_remote_access_snmp`

* New Resource `vyos_protocols_bgp_address_family_ipv6_labeled_unicast_aggregate_address`

* New Resource `vyos_system_conntrack_tcp`

* New Resource `vyos_system_console_device`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_cname`

* New Resource `vyos_protocols_static_route6_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv6_multicast_network`

* New Resource `vyos_interfaces_wireless_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_pim_interface`

* New Resource `vyos_service_config_sync_section_service`

* New Resource `vyos_service_snmp`

* New Resource `vyos_interfaces_wireless_vif_dhcpv6_options_pd`

* New Resource `vyos_policy_route`

* New Resource `vyos_protocols_bgp_parameters_confederation`

* New Resource `vyos_system_console`

* New Resource `vyos_system_wireless`

* New Resource `vyos_service_pppoe_server_interface`

* New Resource `vyos_service_ipoe_server_authentication`

* New Resource `vyos_service_monitoring_zabbix_agent_limits`

* New Resource `vyos_service_config_sync_section`

* New Resource `vyos_system_logs_logrotate_atop`

* New Resource `vyos_service_monitoring_telegraf_azure_data_explorer_authentication`

* New Resource `vyos_protocols_rip_redistribute_bgp`

* New Resource `vyos_protocols_isis_redistribute_ipv6_connected_level_2`

* New Resource `vyos_service_ipoe_server_client_ipv6_pool`

* New Resource `vyos_service_ntp_server`

* New Resource `vyos_load_balancing_reverse_proxy_backend_http_response_headers`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_static`

* New Resource `vyos_protocols_ripng_distribute_list_access_list`

* New Resource `vyos_protocols_ospf_timers_throttle_spf`

* New Resource `vyos_protocols_isis_redistribute_ipv4_babel_level_2`

* New Resource `vyos_protocols_rpki_cache`

* New Resource `vyos_service_ipoe_server_extended_scripts`

* New Resource `vyos_service_snmp_trap_target`

* New Resource `vyos_service_snmp_v3_user`

* New Resource `vyos_protocols_ripng_default_information`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_maximum_paths`

* New Resource `vyos_service_dhcp_relay`

* New Resource `vyos_service_dhcp_relay_relay_options`

* New Resource `vyos_interfaces_bridge_dhcpv6_options_pd`

* New Resource `vyos_protocols_ospf_log_adjacency_changes`

* New Resource `vyos_service_ids_ddos_protection_threshold_general`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_priority_limit_medium`

* New Resource `vyos_protocols_ripng_timers`

* New Resource `vyos_interfaces_bridge`

* New Resource `vyos_interfaces_loopback`

* New Resource `vyos_protocols_ospf_aggregation`

* New Resource `vyos_vpn_openconnect_authentication`

* New Resource `vyos_interfaces_wireless_vif_s`

* New Resource `vyos_protocols_bgp_parameters_tcp_keepalive`

* New Resource `vyos_interfaces_wwan_dhcpv6_options_pd`

* New Resource `vyos_service_ipoe_server_authentication_radius`

* New Resource `vyos_interfaces_ethernet_vif_s_vif_c`

* New Resource `vyos_service_pppoe_server_client_ip_pool`

* New Resource `vyos_protocols_ospf`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_nexthop_vpn`

* New Resource `vyos_protocols_bgp_neighbor_local_as`

* New Resource `vyos_vpn_ipsec_options`

* New Resource `vyos_vpn_sstp_ssl`

* New Resource `vyos_interfaces_dummy`

* New Resource `vyos_service_dns_forwarding_domain_name_server`

* New Resource `vyos_interfaces_openvpn_server_push_route`

* New Resource `vyos_service_pppoe_server_client_ipv6_pool_delegate`

* New Resource `vyos_system_login_banner`

* New Resource `vyos_system_login_tacacs_server`

* New Resource `vyos_vpn_pptp_remote_access`

* New Resource `vyos_protocols_isis_domain_password`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_node_protecting_index`

* New Resource `vyos_service_aws_glb_status`

* New Resource `vyos_vpn_openconnect`

* New Resource `vyos_interfaces_wireless`

* New Resource `vyos_protocols_igmp_proxy`

* New Resource `vyos_service_sla_twamp_server`

* New Resource `vyos_system_flow_accounting`

* New Resource `vyos_vpn_pptp_remote_access_limits`

* New Resource `vyos_vpn_sstp_log`

* New Resource `vyos_protocols_rpki`

* New Resource `vyos_service_console_server_device`

* New Resource `vyos_vpn_ipsec_ike_group`

* New Resource `vyos_protocols_mpls_ldp_export_ipv6`

* New Resource `vyos_interfaces_ethernet_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_load_balancing_reverse_proxy_service_rule`

* New Resource `vyos_protocols_isis_spf_delay_ietf`

* New Resource `vyos_protocols_ripng_redistribute_ospfv3`

* New Resource `vyos_protocols_bgp_interface`

* New Resource `vyos_protocols_mpls_ldp_discovery`

* New Resource `vyos_interfaces_wireless_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_ospf_graceful_restart_helper`

* New Resource `vyos_system_syslog_console_facility`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_aggregate_address`

* New Resource `vyos_protocols_isis_redistribute_ipv6_ripng_level_1`

* New Resource `vyos_protocols_nhrp_tunnel`

* New Resource `vyos_service_ipoe_server_client_ipv6_pool_delegate`

* New Resource `vyos_service_pppoe_server_pado_delay`

* New Resource `vyos_system_conntrack_ignore_ipv6_rule`

* New Resource `vyos_protocols_babel_distribute_list_ipv6_prefix_list`

* New Resource `vyos_protocols_isis_redistribute_ipv6_babel_level_2`

* New Resource `vyos_vpn_sstp_authentication_radius_server`

* New Resource `vyos_load_balancing_wan_interface_health`

* New Resource `vyos_system_frr`

* New Resource `vyos_protocols_rip_default_information`

* New Resource `vyos_protocols_bgp_address_family_ipv6_multicast_distance_prefix`

* New Resource `vyos_protocols_bgp_parameters_bestpath`

* New Resource `vyos_protocols_ospf_area_range`

* New Resource `vyos_system_syslog_global_marker`

* New Resource `vyos_vpn_ipsec_authentication_psk`

* New Resource `vyos_vpn_pptp_remote_access_log`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_system_ip_protocol`

* New Resource `vyos_nat_cgnat_rule`

* New Resource `vyos_service_pppoe_server_authentication_local_users_username`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_remote_prefix_list`

* New Resource `vyos_protocols_ospf_default_information_originate`

* New Resource `vyos_service_https`

* New Resource `vyos_system_flow_accounting_sflow`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_label_vpn`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_rd_vpn`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_route_map_vpn`

* New Resource `vyos_service_monitoring_telegraf_splunk_authentication`

* New Resource `vyos_service_conntrack_sync_failover_mechanism_vrrp`

* New Resource `vyos_vpn_ipsec_remote_access_dhcp`

* New Resource `vyos_protocols_bgp_address_family_ipv4_multicast_distance_prefix`

* New Resource `vyos_protocols_isis_redistribute_ipv4_bgp_level_2`

* New Resource `vyos_protocols_static_neighbor_proxy_and`

* New Resource `vyos_protocols_static_route6`

* New Resource `vyos_protocols_mpls_ldp_targeted_neighbor_ipv4`

* New Resource `vyos_protocols_ospf_redistribute_kernel`

* New Resource `vyos_system_syslog`

* New Resource `vyos_protocols_ospf_distance_ospf`

* New Resource `vyos_service_ntp_allow_client`

* New Resource `vyos_service_tftp_server_listen_address`

* New Resource `vyos_vpn_ipsec_remote_access_connection`

* New Resource `vyos_vpn_openconnect_network_settings`

* New Resource `vyos_load_balancing_reverse_proxy_backend_server`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_isis`

* New Resource `vyos_vpn_openconnect_authentication_identity_based_config`

* New Resource `vyos_protocols_ospf_capability`

* New Resource `vyos_service_pppoe_server_authentication_radius_dynamic_author`

* New Resource `vyos_service_router_advert_interface`

* New Resource `vyos_service_snmp_community`

* New Resource `vyos_load_balancing_wan`

* New Resource `vyos_protocols_isis_default_information_originate_ipv4_level_1`

* New Resource `vyos_service_mdns_repeater`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_label_vpn_allocation_mode`

* New Resource `vyos_service_ntp`

* New Resource `vyos_load_balancing_reverse_proxy_backend`

* New Resource `vyos_protocols_bgp_parameters_dampening`

* New Resource `vyos_protocols_rip_distribute_list_interface`

* New Resource `vyos_system_login`

* New Resource `vyos_protocols_rip_distribute_list_prefix_list`

* New Resource `vyos_protocols_rip_network_distance`

* New Resource `vyos_service_ids_ddos_protection_sflow`

* New Resource `vyos_service_monitoring_telegraf_azure_data_explorer`

* New Resource `vyos_service_sla_owamp_server`

* New Resource `vyos_vpn_sstp_extended_scripts`

* New Resource `vyos_nat_cgnat_pool_external`

* New Resource `vyos_protocols_mpls`

* New Resource `vyos_protocols_bgp_neighbor`

* New Resource `vyos_protocols_isis_redistribute_ipv6_babel_level_1`

* New Resource `vyos_service_config_sync_section_interfaces`

* New Resource `vyos_system_flow_accounting_sflow_server`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_babel_parameters`

* New Resource `vyos_protocols_ospf_segment_routing`

* New Resource `vyos_service_pppoe_server_shaper`

* New Resource `vyos_protocols_ospf_redistribute_isis`

* New Resource `vyos_vpn_ipsec_ike_group_proposal`

* New Resource `vyos_interfaces_bridge_member_interface`

* New Resource `vyos_interfaces_pppoe`

* New Resource `vyos_load_balancing_wan_sticky_connections`

* New Resource `vyos_protocols_bgp_parameters`

* New Resource `vyos_protocols_isis`

* New Resource `vyos_protocols_isis_default_information_originate_ipv4_level_2`

* New Resource `vyos_service_conntrack_sync`

* New Resource `vyos_interfaces_virtual_ethernet_dhcpv6_options_pd`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_aaaa`

* New Resource `vyos_interfaces_wireless_vif_s_vif_c_dhcpv6_options_pd`

* New Resource `vyos_protocols_eigrp`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_ead_es_route_target`

* New Resource `vyos_protocols_ospf_graceful_restart_helper_enable`

* New Resource `vyos_service_broadcast_relay_id`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet_static_mapping`

* New Resource `vyos_service_suricata_port_group`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_label_vpn_allocation_mode`

* New Resource `vyos_protocols_static_route_interface`

* New Resource `vyos_service_ipoe_server_authentication_interface`

* New Resource `vyos_vpn_openconnect_authentication_radius_server`

* New Resource `vyos_interfaces_wireless_security_wpa_radius_server`

* New Resource `vyos_nat_cgnat_pool_external_range`

* New Resource `vyos_service_pppoe_server_client_ipv6_pool`

* New Resource `vyos_interfaces_bonding_vif_s_vif_c_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_ripng_redistribute_static`

* New Resource `vyos_service_ipoe_server_authentication_interface_mac`

* New Resource `vyos_interfaces_ethernet_vif_s_dhcpv6_options_pd`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_naptr`

* New Resource `vyos_system_flow_accounting_netflow_timeout`

* New Resource `vyos_system_frr_snmp`

* New Resource `vyos_vpn_openconnect_network_settings_client_ip_settings`

* New Resource `vyos_protocols_bgp_parameters_distance_global`

* New Resource `vyos_protocols_ospf_redistribute_static`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_txt`

* New Resource `vyos_service_stunnel_server`

* New Resource `vyos_policy_route6`

* New Resource `vyos_protocols_isis_redistribute_ipv4_static_level_1`

* New Resource `vyos_protocols_bgp_parameters_bestpath_peer_type`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_vni`

* New Resource `vyos_service_dhcp_server`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_naptr_rule`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_srv`

* New Resource `vyos_service_stunnel_client_psk`

* New Resource `vyos_service_webproxy_url_filtering_squidguard_auto_update`

* New Resource `vyos_protocols_ospf_graceful_restart`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s_vif_c_dhcpv6_options_pd`

* New Resource `vyos_protocols_static_neighbor_proxy_arp`

* New Resource `vyos_service_monitoring_telegraf_prometheus_client_authentication`

* New Resource `vyos_interfaces_virtual_ethernet_vif`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_route_map_vpn`

* New Resource `vyos_protocols_isis_segment_routing_prefix`

* New Resource `vyos_system_option_ssh_client`

* New Resource `vyos_protocols_nhrp_tunnel_shortcut_target`

* New Resource `vyos_protocols_rip_redistribute_ospf`

* New Resource `vyos_protocols_bgp_neighbor_local_role`

* New Resource `vyos_protocols_pim_spt_switchover_infinity_and_beyond`

* New Resource `vyos_protocols_ripng_distribute_list_prefix_list`

* New Resource `vyos_service_aws_glb_threads`

* New Resource `vyos_interfaces_bonding_vif_s_vif_c_dhcpv6_options_pd`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_export`

* New Resource `vyos_protocols_bgp_bmp_target`

* New Resource `vyos_interfaces_pppoe_dhcpv6_options_pd`

* New Resource `vyos_protocols_static_table`

* New Resource `vyos_system_syslog_host_facility`

* New Resource `vyos_interfaces_bridge_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_bridge_vif_dhcpv6_options_pd`

* New Resource `vyos_nat_cgnat`

* New Resource `vyos_service_config_sync_section_qos`

* New Resource `vyos_protocols_rip_redistribute_connected`

* New Resource `vyos_service_suricata_log_eve`

* New Resource `vyos_protocols_pim_rp`

* New Resource `vyos_protocols_ripng_distribute_list_interface`

* New Resource `vyos_service_webproxy_url_filtering_squidguard_time_period_days`

* New Resource `vyos_service_https_api_keys_id`

* New Resource `vyos_service_ipoe_server_limits`

* New Resource `vyos_protocols_bgp_address_family_ipv6_multicast_aggregate_address`

* New Resource `vyos_service_ndp_proxy_interface`

* New Resource `vyos_load_balancing_wan_rule_interface`

* New Resource `vyos_protocols_static_route_next_hop_bfd_multi_hop_source`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s_vif_c_dhcpv6_options_pd`

* New Resource `vyos_protocols_isis_redistribute_ipv6_kernel_level_1`

* New Resource `vyos_service_aws_glb_script`

* New Resource `vyos_system_static_host_mapping_host_name`

* New Resource `vyos_protocols_bgp_parameters_bestpath_as_path`

* New Resource `vyos_interfaces_input`

* New Resource `vyos_interfaces_virtual_ethernet`

* New Resource `vyos_protocols_isis_redistribute_ipv4_kernel_level_1`

* New Resource `vyos_protocols_mpls_ldp_allocation_ipv6`

* New Resource `vyos_service_router_advert_interface_prefix`

* New Resource `vyos_vpn_ipsec_remote_access_radius`

* New Resource `vyos_interfaces_bonding_vif_dhcpv6_options_pd`

* New Resource `vyos_protocols_isis_default_information_originate_ipv6_level_2`

* New Resource `vyos_protocols_ospf_max_metric_router_lsa`

* New Resource `vyos_protocols_rip_redistribute_babel`

* New Resource `vyos_service_ssh`

* New Resource `vyos_system_acceleration`

* New Resource `vyos_load_balancing_reverse_proxy_service_http_response_headers`

* New Resource `vyos_protocols_bfd_profile`

* New Resource `vyos_protocols_bgp_address_family_ipv6_multicast_distance`

* New Resource `vyos_service_pppoe_server_authentication_radius_rate_limit`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_flooding`

* New Resource `vyos_protocols_isis_segment_routing`

* New Resource `vyos_protocols_ospf_interface_authentication_md5_key_id`

* New Resource `vyos_protocols_ospf_redistribute_table`

* New Resource `vyos_system_lcd`

* New Resource `vyos_interfaces_geneve`

* New Resource `vyos_protocols_rip_redistribute_static`

* New Resource `vyos_system_option`

* New Resource `vyos_protocols_isis_ldp_sync`

* New Resource `vyos_protocols_isis_redistribute_ipv4_rip_level_1`

* New Resource `vyos_interfaces_sstpc`

* New Resource `vyos_service_monitoring_telegraf_influxdb`

* New Resource `vyos_service_suricata_address_group`

* New Resource `vyos_interfaces_ethernet_vif_dhcpv6_options_pd`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_babel_distribute_list_ipv6_access_list`

* New Resource `vyos_service_pppoe_server_ppp_options`

* New Resource `vyos_service_webproxy_listen_address`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_sid_vpn`

* New Resource `vyos_service_dns_forwarding_options`

* New Resource `vyos_system_proxy`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s_vif_c`

* New Resource `vyos_service_lldp_interface`

* New Resource `vyos_service_ipoe_server_client_ipv6_pool_prefix`

* New Resource `vyos_system_config_management`

* New Resource `vyos_service_webproxy_cache_peer`

* New Resource `vyos_vpn_sstp`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_rip`

* New Resource `vyos_system_login_user`

* New Resource `vyos_service_ssh_access_control_allow`

* New Resource `vyos_interfaces_ethernet_dhcpv6_options_pd`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s_vif_c`

* New Resource `vyos_protocols_isis_traffic_engineering`

* New Resource `vyos_protocols_pim_ssm`

* New Resource `vyos_system_syslog_global_facility`

* New Resource `vyos_protocols_isis_redistribute_ipv4_babel_level_1`

* New Resource `vyos_protocols_static_table_route6_next_hop_bfd_multi_hop_source`

* New Resource `vyos_protocols_isis_segment_routing_global_block`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_ns`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s_dhcpv6_options_pd`

* New Resource `vyos_nat_cgnat_pool_internal`

* New Resource `vyos_interfaces_bonding_dhcpv6_options_pd`

* New Resource `vyos_protocols_babel_redistribute_ipv4`

* New Resource `vyos_protocols_ospf_neighbor`

* New Resource `vyos_protocols_isis_default_information_originate_ipv6_level_1`

* New Resource `vyos_service_ids_ddos_protection_threshold_udp`

* New Resource `vyos_interfaces_wireless_vif_s_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_wireless_vif_s_dhcpv6_options_pd`

* New Resource `vyos_vpn_ipsec_esp_group`

* New Resource `vyos_vpn_pptp_remote_access_shaper`

* New Resource `vyos_protocols_static`

* New Resource `vyos_protocols_isis_redistribute_ipv6_bgp_level_1`

* New Resource `vyos_protocols_static_multicast_route`

* New Resource `vyos_interfaces_bonding_vif_s_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_maximum_paths`

* New Resource `vyos_protocols_bgp_srv6`

* New Resource `vyos_service_broadcast_relay`

* New Resource `vyos_service_dns_forwarding`

* New Resource `vyos_service_https_certificates`

* New Resource `vyos_service_webproxy_url_filtering_squidguard_time_period`

* New Resource `vyos_protocols_rip_distribute_list_access_list`

* New Resource `vyos_service_config_sync_section_protocols`

* New Resource `vyos_service_https_allow_client`

* New Resource `vyos_system`

* New Resource `vyos_system_flow_accounting_netflow`

* New Resource `vyos_protocols_bgp_listen_range`

* New Resource `vyos_vpn_openconnect_authentication_mode`

* New Resource `vyos_protocols_bgp_address_family_ipv4_labeled_unicast_aggregate_address`

* New Resource `vyos_protocols_bgp_peer_group`

* New Resource `vyos_protocols_pim_interface_igmp_join`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_label_vpn`

* New Resource `vyos_protocols_mpls_ldp_parameters`

* New Resource `vyos_service_ids_ddos_protection_threshold_icmp`

* New Resource `vyos_service_snmp_v3_view_oid`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_isis_redistribute_ipv6_connected_level_1`

* New Resource `vyos_service_https_api_cors`

* New Resource `vyos_service_salt_minion`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_lowest_backup_metric_index`

* New Resource `vyos_protocols_ospf_redistribute_bgp`

* New Resource `vyos_interfaces_macsec_security_static_peer`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_kernel`

* New Resource `vyos_vpn_pptp_remote_access_authentication_local_users_username`

* New Resource `vyos_service_dhcp_server_shared_network_name`

* New Resource `vyos_service_webproxy_url_filtering_squidguard_source_group`

* New Resource `vyos_system_syslog_user`

* New Resource `vyos_interfaces_macsec_dhcpv6_options_pd`

* New Resource `vyos_protocols_bgp_sid_vpn_per_vrf`

* New Resource `vyos_protocols_isis_redistribute_ipv4_ospf_level_2`

* New Resource `vyos_interfaces_pseudo_ethernet`

* New Resource `vyos_protocols_isis_segment_routing_local_block`

* New Resource `vyos_service_stunnel_client`

* New Resource `vyos_system_ip`

* New Resource `vyos_protocols_rip_interface_authentication_md5`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_ead_es_frag`

* New Resource `vyos_system_conntrack_log_event_update`

* New Resource `vyos_interfaces_ethernet_vif_s_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_sid_vpn`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_import`

* New Resource `vyos_interfaces_bonding_vif`

* New Resource `vyos_interfaces_vxlan`

* New Resource `vyos_interfaces_wireless_dhcpv6_options_pd`

* New Resource `vyos_system_syslog_file_facility`

* New Resource `vyos_system_syslog_host`

* New Resource `vyos_system_task_scheduler_task`

* New Resource `vyos_protocols_ospf_segment_routing_local_block`

* New Resource `vyos_service_https_api`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface`

* New Resource `vyos_service_snmp_listen_address`

* New Resource `vyos_service_snmp_v3`

* New Resource `vyos_service_webproxy_authentication_ldap`

* New Resource `vyos_protocols_mpls_ldp_export_ipv4_export_filter`

* New Resource `vyos_protocols_segment_routing_interface`

* New Resource `vyos_protocols_ripng_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_network`

* New Resource `vyos_protocols_ripng_redistribute_babel`

* New Resource `vyos_system_ip_arp`

* New Resource `vyos_vpn_pptp_remote_access_authentication_radius_dynamic_author`

* New Resource `vyos_vpn_sstp_snmp`

* New Resource `vyos_protocols_ospf_segment_routing_global_block`

* New Resource `vyos_protocols_rip_redistribute_kernel`

* New Resource `vyos_protocols_ripng_redistribute_kernel`

* New Resource `vyos_protocols_static_table_route6_next_hop`

* New Resource `vyos_protocols_mpls_ldp_export_ipv4`

* New Resource `vyos_protocols_segment_routing_srv6_locator`

* New Resource `vyos_service_https_api_graphql_authentication`

* New Resource `vyos_system_conntrack_log`

* New Resource `vyos_system_conntrack_timeout_custom_ipv6_rule`

* New Resource `vyos_policy_route_rule`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_rd_vpn`

* New Resource `vyos_protocols_isis_redistribute_ipv4_bgp_level_1`

* New Resource `vyos_service_dns_dynamic_name`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_connected`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_aggregate_address`

* New Resource `vyos_interfaces_bonding_vif_s`

* New Resource `vyos_policy_local_route6_rule`

* New Resource `vyos_service_pppoe_server_authentication_radius`

* New Resource `vyos_vpn_openconnect_network_settings_client_ipv6_pool`

* New Resource `vyos_protocols_mpls_ldp_allocation_ipv4`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_ripng`

* New Resource `vyos_service_ipoe_server_shaper`

* New Resource `vyos_protocols_pim`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_load_sharing_disable`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_static`

* New Resource `vyos_service_monitoring_zabbix_agent`

* New Resource `vyos_vpn_sstp_authentication_radius_rate_limit`

* New Resource `vyos_service_pppoe_server_log`

* New Resource `vyos_service_router_advert_interface_nat64prefix`

* New Resource `vyos_interfaces_wireless_vif`

* New Resource `vyos_protocols_bgp_address_family_ipv6_flowspec_local_install`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_route_target_vpn`

* New Resource `vyos_service_dhcp_server_high_availability`

* New Resource `vyos_system_login_tacacs`

* New Resource `vyos_interfaces_pppoe_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_virtual_ethernet_dhcpv6_options_pd_interface`

* New Resource `vyos_system_conntrack_timeout_custom_ipv4_rule`

* New Resource `vyos_service_ids_ddos_protection_threshold_tcp`

* New Resource `vyos_service_webproxy_url_filtering`

* New Resource `vyos_vpn_openconnect_ssl`

* New Resource `vyos_protocols_babel_interface`

* New Resource `vyos_protocols_ospf_summary_address`

* New Resource `vyos_system_ip_multipath`

* New Resource `vyos_protocols_bgp_address_family_ipv4_labeled_unicast_network`

* New Resource `vyos_interfaces_wwan_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_ospf_redistribute_babel`

* New Resource `vyos_protocols_static_multicast_route_next_hop`

* New Resource `vyos_protocols_nhrp_tunnel_dynamic_map`

* New Resource `vyos_system_conntrack_modules`

* New Resource `vyos_system_flow_accounting_netflow_server`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_srv_entry`

* New Resource `vyos_service_ipoe_server_client_ip_pool`

* New Resource `vyos_system_option_kernel_debug`

* New Resource `vyos_vpn_sstp_shaper`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_downstream_index`

* New Resource `vyos_service_ipoe_server_interface`

* New Resource `vyos_protocols_ospf_auto_cost`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet_static_mapping_option_static_route`

* New Resource `vyos_system_config_management_commit_archive`

* New Resource `vyos_vpn_ipsec_remote_access_pool`

* New Resource `vyos_vpn_openconnect_accounting_radius_server`

* New Resource `vyos_protocols_bgp_address_family_ipv4_multicast_distance`

* New Resource `vyos_service_ipoe_server_authentication_radius_rate_limit`

* New Resource `vyos_service_pppoe_server_authentication_radius_server`

* New Resource `vyos_protocols_nhrp_tunnel_map`

* New Resource `vyos_load_balancing_reverse_proxy_service_logging_facility`

* New Resource `vyos_service_pppoe_server_client_ipv6_pool_prefix`

* New Resource `vyos_vpn_sstp_authentication`

* New Resource `vyos_protocols_isis_redistribute_ipv4_connected_level_2`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_priority_limit_critical`

* New Resource `vyos_protocols_ospf_interface`

* New Resource `vyos_vpn_sstp_authentication_local_users_username`

* New Resource `vyos_service_config_sync_secondary`

* New Resource `vyos_service_ssh_rekey`

* New Resource `vyos_interfaces_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_route_target`

* New Resource `vyos_protocols_rip_interface`

* New Resource `vyos_service_snmp_script_extensions_extension_name`

* New Resource `vyos_service_stunnel_log`

* New Resource `vyos_interfaces_vti`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_network`

* New Resource `vyos_service_monitoring_telegraf_loki`

* New Resource `vyos_service_monitoring_zabbix_agent_log`

* New Resource `vyos_vpn_ipsec_site_to_site_peer`

* New Resource `vyos_interfaces_pseudo_ethernet_dhcpv6_options_pd`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s`

* New Resource `vyos_service_conntrack_sync_interface`

* New Resource `vyos_vpn_sstp_ppp_options`

* New Resource `vyos_system_conntrack_log_event_destroy`

* New Resource `vyos_protocols_pim_rp_address`

* New Resource `vyos_protocols_bgp_address_family_ipv4_flowspec_local_install`

* New Resource `vyos_protocols_bgp_address_family_ipv4_multicast_network`

* New Resource `vyos_protocols_igmp_proxy_interface`

* New Resource `vyos_service_dns_dynamic`

* New Resource `vyos_service_ids_ddos_protection`

* New Resource `vyos_vpn_pptp_remote_access_client_ipv6_pool`

* New Resource `vyos_interfaces_bridge_vif`

* New Resource `vyos_interfaces_virtual_ethernet_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_policy_local_route_rule`

* New Resource `vyos_protocols_static_multicast_interface_route`

* New Resource `vyos_service_monitoring_telegraf_loki_authentication`

* New Resource `vyos_service_snmp_mib`

* New Resource `vyos_service_suricata`

* New Resource `vyos_vpn_pptp_remote_access_ppp_options`

* New Resource `vyos_protocols_bgp_peer_group_local_as`

* New Resource `vyos_service_config_sync_section_system`

* New Resource `vyos_protocols_mpls_parameters`

* New Resource `vyos_protocols_ospf_refresh`

* New Resource `vyos_vpn_sstp_client_ipv6_pool_prefix`

* New Resource `vyos_interfaces_ethernet_vif_s_vif_c_dhcpv6_options_pd`

* New Resource `vyos_interfaces_wireguard_peer`

* New Resource `vyos_service_router_advert_interface_route`

* New Resource `vyos_system_option_kernel`

* New Resource `vyos_vpn_pptp_remote_access_client_ipv6_pool_delegate`

* New Resource `vyos_vpn_sstp_client_ipv6_pool_delegate`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_export`

* New Resource `vyos_protocols_isis_redistribute_ipv4_kernel_level_2`

* New Resource `vyos_system_sflow_server`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_distance_prefix`

* New Resource `vyos_protocols_bgp_parameters_graceful_restart`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_distance`

* New Resource `vyos_protocols_static_table_route_interface`

* New Resource `vyos_service_webproxy_authentication`

* New Resource `vyos_protocols_bgp_peer_group_local_role`

* New Resource `vyos_protocols_mpls_ldp`

* New Resource `vyos_service_tftp_server`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_dhcpv6_options_pd`

* New Resource `vyos_protocols_ospf_access_list`

* New Resource `vyos_protocols_ripng`

* New Resource `vyos_service_ndp_proxy_interface_prefix`

* New Resource `vyos_system_logs_logrotate_messages`

* New Resource `vyos_vpn_sstp_client_ip_pool`

* New Resource `vyos_load_balancing_wan_rule`

* New Resource `vyos_service_webproxy_url_filtering_squidguard_rule`

* New Resource `vyos_system_syslog_file`

* New Resource `vyos_vpn_pptp_remote_access_client_ipv6_pool_prefix`

* New Resource `vyos_interfaces_ethernet_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_pim_register_accept_list`

* New Resource `vyos_service_webproxy`

* New Resource `vyos_protocols_isis_redistribute_ipv6_static_level_1`

* New Resource `vyos_service_https_api_graphql`

* New Resource `vyos_interfaces_virtual_ethernet_vif_dhcpv6_options_pd`

* New Resource `vyos_protocols_bgp_parameters_conditional_advertisement`

* New Resource `vyos_interfaces_bridge_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_static_route_next_hop`

* New Resource `vyos_protocols_ospf_parameters`

* New Resource `vyos_protocols_mpls_ldp_neighbor`

* New Resource `vyos_protocols_static_arp_interface_address`

* New Resource `vyos_system_sflow`

* New Resource `vyos_service_monitoring_telegraf`

* New Resource `vyos_protocols_ospf_area`

* New Resource `vyos_service_ndp_proxy`

* New Resource `vyos_service_pppoe_server_extended_scripts`

* New Resource `vyos_system_ip_tcp_mss`

* New Resource `vyos_vpn_openconnect_accounting_mode`

* New Resource `vyos_protocols_babel_distribute_list_ipv6_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv4_vpn_network`

* New Resource `vyos_protocols_ospf_redistribute_rip`

* New Resource `vyos_load_balancing_reverse_proxy_service`

* New Resource `vyos_protocols_bgp_parameters_default`

* New Resource `vyos_protocols_ospf_ldp_sync`

* New Resource `vyos_service_monitoring_telegraf_prometheus_client`

* New Resource `vyos_protocols_static_table_route_next_hop`

* New Resource `vyos_service_monitoring_telegraf_influxdb_authentication`

* New Resource `vyos_interfaces_wireguard`

* New Resource `vyos_interfaces_pseudo_ethernet_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_bgp`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_ospf`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_ospfv3`

* New Resource `vyos_service_lldp_legacy_protocols`

* New Resource `vyos_interfaces_bonding_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_openvpn`

* New Resource `vyos_system_syslog_user_facility`

* New Resource `vyos_protocols_babel_distribute_list_ipv4_prefix_list`

* New Resource `vyos_protocols_bgp_parameters_distance_prefix`

* New Resource `vyos_vpn_sstp_authentication_radius`

* New Resource `vyos_protocols_bgp_listen`

* New Resource `vyos_protocols_eigrp_metric`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet_option_static_route`

* New Resource `vyos_vpn_ipsec_esp_group_proposal`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_distance`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_import`

* New Resource `vyos_protocols_ripng_redistribute_bgp`

* New Resource `vyos_protocols_static_table_route6`

* New Resource `vyos_load_balancing_reverse_proxy_backend_rule`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_connected`

* New Resource `vyos_system_ip_nht`

* New Resource `vyos_protocols_isis_redistribute_ipv6_kernel_level_2`

* New Resource `vyos_service_monitoring_zabbix_agent_server_active`

* New Resource `vyos_service_snmp_v3_view`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_mac_vrf`

* New Resource `vyos_service_monitoring_telegraf_splunk`

* New Resource `vyos_service_pppoe_server`

* New Resource `vyos_protocols_ospf_mpls_te`

* New Resource `vyos_service_ipoe_server_authentication_radius_server`

* New Resource `vyos_protocols_mpls_ldp_export_ipv6_export_filter`

* New Resource `vyos_protocols_mpls_ldp_targeted_neighbor_ipv6`

* New Resource `vyos_interfaces_bonding`

* New Resource `vyos_service_snmp_v3_group`

* New Resource `vyos_protocols_babel_distribute_list_ipv4_access_list`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_default_originate`

* New Resource `vyos_vpn_pptp_remote_access_extended_scripts`

* New Resource `vyos_load_balancing_reverse_proxy_global_parameters`

* New Resource `vyos_protocols_bfd_peer`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_babel`

* New Resource `vyos_protocols_mpls_ldp_import_ipv6_import_filter`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet_range_option_static_route`

* New Resource `vyos_service_ipoe_server_log`

* New Resource `vyos_interfaces_bonding_vif_s_dhcpv6_options_pd`

* New Resource `vyos_interfaces_ethernet`

* New Resource `vyos_system_conntrack`

* New Resource `vyos_vpn_ipsec_remote_access_radius_server`

* New Resource `vyos_vpn_pptp_remote_access_authentication_radius_server`

* New Resource `vyos_interfaces_ethernet_vif_s`

* New Resource `vyos_vpn_ipsec_profile`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_advertise_ipv4_unicast`

* New Resource `vyos_protocols_mpls_ldp_import_ipv4_import_filter`

* New Resource `vyos_interfaces_wireless_vif_s_vif_c`

* New Resource `vyos_vpn_ipsec`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute`

* New Resource `vyos_protocols_static_table_route6_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_babel`

* New Resource `vyos_protocols_static_route6_next_hop`

* New Resource `vyos_vpn_openconnect_authentication_radius`

* New Resource `vyos_vpn_pptp_remote_access_client_ip_pool`

* New Resource `vyos_interfaces_vxlan_vlan_to_vni`

* New Resource `vyos_protocols_babel_distribute_list_ipv4_interface`








## Previous changelogs
For older versions see changelog archive [directory](data/changelogs/) or changelog for [version v0](CHANGELOG-v0.md)
