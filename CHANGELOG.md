
# CHANGELOG

<!-- ToC start -->
   1. [Release 1.4.202408140 (2024-08-14 15:40:33 UTC)](#release-14202408140-2024-08-14-15:40:33-utc)
      1. [Project changes](#project-changes)
         1. [Notes](#notes)
         1. [Features](#features)
         1. [Bug fixes](#bug-fixes)
      1. [Schema changes](#schema-changes)
         1. [BREAKING CHANGES](#breaking-changes)
            1. [Resources](#resources)
         1. [Features](#features-1)
            1. [Resources](#resources-1)
   1. [0.4.20240813 (2024-08-13 08:24:24 UTC)](#0420240813-2024-08-13-08:24:24-utc)
   1. [0.4.20240806 (2024-08-06 08:24:19 UTC)](#0420240806-2024-08-06-08:24:19-utc)
   1. [0.4.20240803 (2024-08-03 08:23:39 UTC)](#0420240803-2024-08-03-08:23:39-utc)
   1. [0.4.20240801 (2024-08-01 08:24:17 UTC)](#0420240801-2024-08-01-08:24:17-utc)
   1. [0.4.20240725 (2024-07-30 08:24:02 UTC)](#0420240725-2024-07-30-08:24:02-utc)
   1. [0.3.20240725 (2024-07-25 08:24:01 UTC)](#0320240725-2024-07-25-08:24:01-utc)
   1. [0.3.20240717 (2024-07-21 08:22:11 UTC)](#0320240717-2024-07-21-08:22:11-utc)
   1. [0.3.20240710 (2024-07-10 08:29:07 UTC)](#0320240710-2024-07-10-08:29:07-utc)
   1. [0.3.20240709 (2024-07-09 08:22:43 UTC)](#0320240709-2024-07-09-08:22:43-utc)
   1. [0.3.20240707 (2024-07-07 08:22:13 UTC)](#0320240707-2024-07-07-08:22:13-utc)
   1. [0.3.20240706 (2024-07-06 08:23:33 UTC)](#0320240706-2024-07-06-08:23:33-utc)
   1. [0.3.20240705 (2024-07-05 08:24:23 UTC)](#0320240705-2024-07-05-08:24:23-utc)
   1. [0.3.20240704 (2024-07-04 08:23:56 UTC)](#0320240704-2024-07-04-08:23:56-utc)
   1. [0.3.20240702 (2024-07-03 08:23:41 UTC)](#0320240702-2024-07-03-08:23:41-utc)
   1. [0.3.20240701 (2024-07-01 08:23:44 UTC)](#0320240701-2024-07-01-08:23:44-utc)
   1. [0.3.20240630 (2024-06-30 08:21:28 UTC)](#0320240630-2024-06-30-08:21:28-utc)
   1. [0.3.20240629 (2024-06-29 08:22:20 UTC)](#0320240629-2024-06-29-08:22:20-utc)
   1. [0.3.20240628 (2024-06-28 08:22:19 UTC)](#0320240628-2024-06-28-08:22:19-utc)
   1. [0.3.20240627 (2024-06-27 08:23:29 UTC)](#0320240627-2024-06-27-08:23:29-utc)
   1. [0.3.20240626 (2024-06-26 08:23:08 UTC)](#0320240626-2024-06-26-08:23:08-utc)
   1. [0.3.20240625 (2024-06-25 08:23:29 UTC)](#0320240625-2024-06-25-08:23:29-utc)
   1. [0.3.20240623 (2024-06-23 08:21:01 UTC)](#0320240623-2024-06-23-08:21:01-utc)
   1. [0.3.20240621 (2024-06-21 08:22:55 UTC)](#0320240621-2024-06-21-08:22:55-utc)
   1. [0.3.20240620 (2024-06-20 08:23:39 UTC)](#0320240620-2024-06-20-08:23:39-utc)
   1. [0.3.20240619 (2024-06-19 08:23:09 UTC)](#0320240619-2024-06-19-08:23:09-utc)
   1. [0.3.20240617 (2024-06-17 08:23:38 UTC)](#0320240617-2024-06-17-08:23:38-utc)
   1. [0.3.20240613 (2024-06-13 08:23:27 UTC)](#0320240613-2024-06-13-08:23:27-utc)
   1. [0.3.20240612 (2024-06-12 08:22:53 UTC)](#0320240612-2024-06-12-08:22:53-utc)
   1. [0.3.20240606 (2024-06-06 08:23:02 UTC)](#0320240606-2024-06-06-08:23:02-utc)
   1. [0.3.20240602 (2024-06-02 08:21:22 UTC)](#0320240602-2024-06-02-08:21:22-utc)
   1. [0.3.20240601 (2024-06-01 08:23:16 UTC)](#0320240601-2024-06-01-08:23:16-utc)
   1. [0.3.20240531 (2024-05-31 08:27:03 UTC)](#0320240531-2024-05-31-08:27:03-utc)
   1. [0.3.20240528 (2024-05-28 08:24:26 UTC)](#0320240528-2024-05-28-08:24:26-utc)
   1. [0.3.20240527 (2024-05-27 08:23:01 UTC)](#0320240527-2024-05-27-08:23:01-utc)
   1. [0.3.20240526 (2024-05-26 08:22:01 UTC)](#0320240526-2024-05-26-08:22:01-utc)
   1. [0.3.20240524 (2024-05-24 08:23:39 UTC)](#0320240524-2024-05-24-08:23:39-utc)
   1. [0.3.20240522 (2024-05-22 08:23:17 UTC)](#0320240522-2024-05-22-08:23:17-utc)
   1. [0.3.20240521 (2024-05-21 08:22:14 UTC)](#0320240521-2024-05-21-08:22:14-utc)
   1. [0.3.20240520 (2024-05-20 08:22:55 UTC)](#0320240520-2024-05-20-08:22:55-utc)
   1. [0.3.20240518 (2024-05-19 08:20:42 UTC)](#0320240518-2024-05-19-08:20:42-utc)
   1. [0.3.20240514 (2024-05-14 08:23:36 UTC)](#0320240514-2024-05-14-08:23:36-utc)
   1. [0.3.20240512 (2024-05-13 08:23:13 UTC)](#0320240512-2024-05-13-08:23:13-utc)
   1. [0.3.20240510 (2024-05-10 08:21:55 UTC)](#0320240510-2024-05-10-08:21:55-utc)
   1. [0.3.20240509 (2024-05-09 08:22:09 UTC)](#0320240509-2024-05-09-08:22:09-utc)
   1. [0.3.20240508 (2024-05-08 08:23:55 UTC)](#0320240508-2024-05-08-08:23:55-utc)
   1. [0.3.20240507 (2024-05-07 08:21:46 UTC)](#0320240507-2024-05-07-08:21:46-utc)
   1. [0.3.20240506 (2024-05-06 10:32:57 UTC)](#0320240506-2024-05-06-10:32:57-utc)
   1. [0.3.20240505 (2024-05-05 13:02:39 UTC)](#0320240505-2024-05-05-13:02:39-utc)
   1. [0.3.20240323 (2024-05-01 15:38:08 UTC)](#0320240323-2024-05-01-15:38:08-utc)
   1. [0.2.20240323 (2024-04-28 12:32:19 UTC)](#0220240323-2024-04-28-12:32:19-utc)
   1. [0.2.20240324 (2024-04-28 12:20:45 UTC)](#0220240324-2024-04-28-12:20:45-utc)
   1. [0.1.20240323 (2024-04-27 17:21:15 UTC)](#0120240323-2024-04-27-17:21:15-utc)
   1. [0.0.20240322 (2024-03-22 00:00:00 UTC)](#0020240322-2024-03-22-00:00:00-utc)
<!-- ToC end -->


## Release 1.4.202408140 (2024-08-14 15:40:33 UTC)
### Project changes
#### Notes
* improve provider documentation readability
* improve versioning and changelog generation
* update to rolling release 2024-08-14T00:20:52Z
#### Features
* reenable generation of missing resources and move resource identifying information into dedicated parameter
#### Bug fixes
* changelog and release procedures

### Schema changes
#### BREAKING CHANGES

##### Resources
* Modified Resource `vyos_firewall_group_ipv6_address_group`
	* **Removed attribute** `ipv6_address_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_name_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_route_map_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `route_map_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_aggregate_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `aggregate_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_vpn_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_interface`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_area`
	* **Removed attribute** `name_id`
	* **Removed attribute** `area_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_input_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper_hfsc_class_match`
	* **Removed attribute** `match_id`
	* **Removed attribute** `class_id`
	* **Removed attribute** `shaper_hfsc_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route6_interface`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `route6_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_virtual_server`
	* **Removed attribute** `virtual_server_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_vrrp_sync_group`
	* **Removed attribute** `sync_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_openssh`
	* **Removed attribute** `openssh_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_limiter`
	* **Removed attribute** `limiter_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_limiter_class_match`
	* **Removed attribute** `class_id`
	* **Removed attribute** `match_id`
	* **Removed attribute** `limiter_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_area_virtual_link`
	* **Removed attribute** `name_id`
	* **Removed attribute** `virtual_link_id`
	* **Removed attribute** `area_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_interface_group`
	* **Removed attribute** `interface_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_as_path_list_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `as_path_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_route_map`
	* **Removed attribute** `route_map_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_network_emulator`
	* **Removed attribute** `network_emulator_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_distance_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospfv3_interface`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_dh`
	* **Removed attribute** `dh_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_access_list6`
	* **Removed attribute** `access_list6_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_cake`
	* **Removed attribute** `cake_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_priority_queue_class_match`
	* **Removed attribute** `match_id`
	* **Removed attribute** `priority_queue_id`
	* **Removed attribute** `class_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_access_list`
	* **Removed attribute** `name_id`
	* **Removed attribute** `access_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_address_group`
	* **Removed attribute** `address_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_input_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_access_list`
	* **Removed attribute** `access_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_prefix_list_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `prefix_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper_class_match`
	* **Removed attribute** `match_id`
	* **Removed attribute** `shaper_id`
	* **Removed attribute** `class_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_bmp_target`
	* **Removed attribute** `name_id`
	* **Removed attribute** `target_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_area_virtual_link_authentication_md5_key_id`
	* **Removed attribute** `key_id_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `virtual_link_id`
	* **Removed attribute** `area_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospfv3_area`
	* **Removed attribute** `name_id`
	* **Removed attribute** `area_id`
	* New attribute `identifier`

* **Removed Resource** `vyos_container_name_network`

* Modified Resource `vyos_pki_ca`
	* **Removed attribute** `ca_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_certificate`
	* **Removed attribute** `certificate_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_community_list`
	* **Removed attribute** `community_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_community_list_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `community_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_interface`
	* **Removed attribute** `interface_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_distance_prefix`
	* **Removed attribute** `prefix_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route6_next_hop_bfd_multi_hop_source`
	* **Removed attribute** `name_id`
	* **Removed attribute** `next_hop_id`
	* **Removed attribute** `route6_id`
	* **Removed attribute** `source_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_aggregate_address`
	* **Removed attribute** `aggregate_address_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_nat_destination_rule_load_balance_backend`
	* **Removed attribute** `backend_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_large_community_list`
	* **Removed attribute** `large_community_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route_next_hop_bfd_multi_hop_source`
	* **Removed attribute** `name_id`
	* **Removed attribute** `next_hop_id`
	* **Removed attribute** `route_id`
	* **Removed attribute** `source_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_ipv6_network_group`
	* **Removed attribute** `ipv6_network_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_output_raw_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_zone_from`
	* **Removed attribute** `from_id`
	* **Removed attribute** `zone_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_vrrp_group_excluded_address`
	* **Removed attribute** `excluded_address_id`
	* **Removed attribute** `group_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_access_list_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `access_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_listen_range`
	* **Removed attribute** `name_id`
	* **Removed attribute** `range_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_area_range`
	* **Removed attribute** `name_id`
	* **Removed attribute** `range_id`
	* **Removed attribute** `area_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route_interface`
	* **Removed attribute** `name_id`
	* **Removed attribute** `route_id`
	* **Removed attribute** `interface_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_forward_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_output_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_access_list6_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `access_list6_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_round_robin_class`
	* **Removed attribute** `round_robin_id`
	* **Removed attribute** `class_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_vpn_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_unicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_peer_group_local_as`
	* **Removed attribute** `name_id`
	* **Removed attribute** `peer_group_id`
	* **Removed attribute** `local_as_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route6`
	* **Removed attribute** `name_id`
	* **Removed attribute** `route6_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_forward_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_random_detect_precedence`
	* **Removed attribute** `random_detect_id`
	* **Removed attribute** `precedence_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_labeled_unicast_aggregate_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `aggregate_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_labeled_unicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_neighbor`
	* **Removed attribute** `neighbor_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_environment`
	* **Removed attribute** `environment_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_parameters_distance_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_netns_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_traffic_match_group_match`
	* **Removed attribute** `traffic_match_group_id`
	* **Removed attribute** `match_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_l2vpn_evpn_vni`
	* **Removed attribute** `vni_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_dynamic_group_ipv6_address_group`
	* **Removed attribute** `ipv6_address_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_lowest_backup_metric_index`
	* **Removed attribute** `name_id`
	* **Removed attribute** `index_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_flowtable`
	* **Removed attribute** `flowtable_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_limiter_class`
	* **Removed attribute** `class_id`
	* **Removed attribute** `limiter_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper_class`
	* **Removed attribute** `class_id`
	* **Removed attribute** `shaper_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_port`
	* **Removed attribute** `name_id`
	* **Removed attribute** `port_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_output_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_key_pair`
	* **Removed attribute** `key_pair_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_large_community_list_rule`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `large_community_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_downstream_index`
	* **Removed attribute** `name_id`
	* **Removed attribute** `index_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_remote_prefix_list`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_registry`
	* **Removed attribute** `registry_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_rate_control`
	* **Removed attribute** `rate_control_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_unicast_distance_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_node_protecting_index`
	* **Removed attribute** `name_id`
	* **Removed attribute** `index_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_network_group`
	* **Removed attribute** `network_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_random_detect`
	* **Removed attribute** `random_detect_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_isis_segment_routing_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_segment_routing_prefix`
	* **Removed attribute** `name_id`
	* **Removed attribute** `prefix_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route_next_hop`
	* **Removed attribute** `route_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `next_hop_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_fair_queue`
	* **Removed attribute** `fair_queue_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_dynamic_group_address_group`
	* **Removed attribute** `address_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_prerouting_raw_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_nat_static_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_priority_queue`
	* **Removed attribute** `priority_queue_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_round_robin`
	* **Removed attribute** `round_robin_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_neighbor`
	* **Removed attribute** `neighbor_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route`
	* **Removed attribute** `name_id`
	* **Removed attribute** `route_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_unicast_aggregate_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `aggregate_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_name_rule`
	* **Removed attribute** `name_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_round_robin_class_match`
	* **Removed attribute** `class_id`
	* **Removed attribute** `match_id`
	* **Removed attribute** `round_robin_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper_hfsc`
	* **Removed attribute** `shaper_hfsc_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_distance_prefix`
	* **Removed attribute** `prefix_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospfv3_area_range`
	* **Removed attribute** `area_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `range_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_domain_group`
	* **Removed attribute** `domain_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_nat_source_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_pki_openvpn_shared_secret`
	* **Removed attribute** `shared_secret_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_aggregate_address`
	* **Removed attribute** `aggregate_address_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_prerouting_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_vrrp_group_address`
	* **Removed attribute** `address_id`
	* **Removed attribute** `group_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_extcommunity_list`
	* **Removed attribute** `extcommunity_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_ipv6_protocol`
	* **Removed attribute** `name_id`
	* **Removed attribute** `protocol_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_name_rule`
	* **Removed attribute** `name_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_output_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv6_prerouting_raw_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_virtual_server_real_server`
	* **Removed attribute** `real_server_id`
	* **Removed attribute** `virtual_server_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper`
	* **Removed attribute** `shaper_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_peer_group_local_role`
	* **Removed attribute** `local_role_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `peer_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_drop_tail`
	* **Removed attribute** `drop_tail_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_input_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_high_availability_vrrp_group`
	* **Removed attribute** `group_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_prefix_list6_rule`
	* **Removed attribute** `prefix_list6_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_priority_queue_class`
	* **Removed attribute** `class_id`
	* **Removed attribute** `priority_queue_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_traffic_match_group`
	* **Removed attribute** `traffic_match_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_mac_group`
	* **Removed attribute** `mac_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_group_port_group`
	* **Removed attribute** `port_group_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_ipv4_output_raw_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_nat_source_rule_load_balance_backend`
	* **Removed attribute** `rule_id`
	* **Removed attribute** `backend_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_as_path_list`
	* **Removed attribute** `as_path_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_interface`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_volume`
	* **Removed attribute** `volume_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_zone`
	* **Removed attribute** `zone_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_extcommunity_list_rule`
	* **Removed attribute** `extcommunity_list_id`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_prefix_list`
	* **Removed attribute** `prefix_list_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_shaper_hfsc_class`
	* **Removed attribute** `class_id`
	* **Removed attribute** `shaper_hfsc_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_labeled_unicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_network`
	* **Removed attribute** `name_id`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_label`
	* **Removed attribute** `label_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_interface`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_name`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_nat_destination_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_policy_prefix_list6`
	* **Removed attribute** `prefix_list6_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_peer_group`
	* **Removed attribute** `peer_group_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_static_route6_next_hop`
	* **Removed attribute** `name_id`
	* **Removed attribute** `route6_id`
	* **Removed attribute** `next_hop_id`
	* New attribute `identifier`

* Modified Resource `vyos_firewall_bridge_forward_filter_rule`
	* **Removed attribute** `rule_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_labeled_unicast_aggregate_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `aggregate_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_neighbor_local_as`
	* **Removed attribute** `local_as_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `neighbor_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_summary_address`
	* **Removed attribute** `name_id`
	* **Removed attribute** `summary_address_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_network`
	* **Removed attribute** `network_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_bgp_neighbor_local_role`
	* **Removed attribute** `local_role_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `neighbor_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_interface_authentication_md5_key_id`
	* **Removed attribute** `name_id`
	* **Removed attribute** `interface_id`
	* **Removed attribute** `key_id_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_device`
	* **Removed attribute** `device_id`
	* **Removed attribute** `name_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_protocols_ospf_redistribute_table`
	* **Removed attribute** `name_id`
	* **Removed attribute** `table_id`
	* New attribute `identifier`

* Modified Resource `vyos_qos_policy_fq_codel`
	* **Removed attribute** `fq_codel_id`
	* New attribute `identifier`

* Modified Resource `vyos_vrf_name_ip_protocol`
	* **Removed attribute** `name_id`
	* **Removed attribute** `protocol_id`
	* New attribute `identifier`

* Modified Resource `vyos_container_name_sysctl_parameter`
	* **Removed attribute** `name_id`
	* **Removed attribute** `parameter_id`
	* New attribute `identifier`





#### Features

##### Resources
* New Resource `vyos_protocols_bgp`

* New Resource `vyos_protocols_isis_area_password`

* New Resource `vyos_service_dhcp_relay_relay_options`

* New Resource `vyos_protocols_ospf_redistribute_babel`

* New Resource `vyos_protocols_rip_distribute_list_access_list`

* New Resource `vyos_protocols_ripng_distribute_list_interface`

* New Resource `vyos_system_sysctl_parameter`

* New Resource `vyos_system_ip_protocol`

* New Resource `vyos_protocols_mpls_ldp_import_ipv6_import_filter`

* New Resource `vyos_service_monitoring_telegraf_azure_data_explorer_authentication`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_label_vpn_allocation_mode`

* New Resource `vyos_service_aws_glb_threads`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_ns`

* New Resource `vyos_interfaces_wireguard_peer`

* New Resource `vyos_protocols_rip_distribute_list_interface`

* New Resource `vyos_interfaces_bridge_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_ospf_default_information_originate`

* New Resource `vyos_vpn_ipsec_site_to_site_peer`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_import`

* New Resource `vyos_service_ipoe_server_client_ip_pool`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_route_map_vpn`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet_static_mapping`

* New Resource `vyos_protocols_bgp_parameters_bestpath`

* New Resource `vyos_service_snmp_v3_group`

* New Resource `vyos_service_tftp_server`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_txt`

* New Resource `vyos_interfaces_ethernet_dhcpv6_options_pd`

* New Resource `vyos_protocols_rip_timers`

* New Resource `vyos_protocols_ripng_redistribute_bgp`

* New Resource `vyos_interfaces_virtual_ethernet`

* New Resource `vyos_service_pppoe_server_snmp`

* New Resource `vyos_service_webproxy_url_filtering_squidguard_time_period_days`

* New Resource `vyos_vpn_ipsec_remote_access_connection_authentication_local_users_username`

* New Resource `vyos_service_config_sync_section_protocols`

* New Resource `vyos_vpn_pptp_remote_access_client_ipv6_pool`

* New Resource `vyos_vpn_pptp_remote_access_log`

* New Resource `vyos_protocols_rip_redistribute_kernel`

* New Resource `vyos_protocols_rip_redistribute_static`

* New Resource `vyos_service_pppoe_server_client_ipv6_pool_prefix`

* New Resource `vyos_vpn_ipsec_ike_group`

* New Resource `vyos_interfaces_dummy`

* New Resource `vyos_interfaces_virtual_ethernet_dhcpv6_options_pd`

* New Resource `vyos_service_pppoe_server_authentication`

* New Resource `vyos_protocols_ripng_default_information`

* New Resource `vyos_service_conntrack_sync_interface`

* New Resource `vyos_service_ntp_server`

* New Resource `vyos_system_flow_accounting`

* New Resource `vyos_interfaces_vxlan_vlan_to_vni`

* New Resource `vyos_protocols_isis_default_information_originate_ipv6_level_2`

* New Resource `vyos_service_monitoring_telegraf_azure_data_explorer`

* New Resource `vyos_service_sla_owamp_server`

* New Resource `vyos_protocols_isis_default_information_originate_ipv4_level_2`

* New Resource `vyos_protocols_static_table_route_next_hop_bfd_multi_hop_source`

* New Resource `vyos_interfaces_ethernet_vif_s_vif_c`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_sid_vpn`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_ospfv3`

* New Resource `vyos_protocols_ospf_redistribute_bgp`

* New Resource `vyos_protocols_static_neighbor_proxy_and`

* New Resource `vyos_service_ids_ddos_protection_threshold_udp`

* New Resource `vyos_protocols_bgp_address_family_ipv4_labeled_unicast_network`

* New Resource `vyos_protocols_bgp_srv6`

* New Resource `vyos_protocols_isis_redistribute_ipv4_rip_level_2`

* New Resource `vyos_protocols_rip_redistribute_ospf`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_nexthop_vpn`

* New Resource `vyos_system_conntrack_log_event_new`

* New Resource `vyos_system_syslog_console_facility`

* New Resource `vyos_protocols_mpls_parameters`

* New Resource `vyos_service_event_handler_event`

* New Resource `vyos_vpn_ipsec_remote_access_radius`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_spf`

* New Resource `vyos_service_event_handler_event_script_environment`

* New Resource `vyos_interfaces_tunnel`

* New Resource `vyos_protocols_bgp_address_family_ipv6_multicast_network`

* New Resource `vyos_protocols_isis_redistribute_ipv6_connected_level_1`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_rd_vpn`

* New Resource `vyos_protocols_bgp_address_family_ipv6_multicast_distance`

* New Resource `vyos_service_ids_ddos_protection_threshold_general`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_label_vpn`

* New Resource `vyos_vpn_ipsec_remote_access_dhcp`

* New Resource `vyos_vpn_openconnect`

* New Resource `vyos_interfaces_ethernet_dhcpv6_options_pd_interface`

* New Resource `vyos_service_ntp`

* New Resource `vyos_service_stunnel_client`

* New Resource `vyos_interfaces_pseudo_ethernet`

* New Resource `vyos_protocols_babel_distribute_list_ipv6_access_list`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_priority_limit_high`

* New Resource `vyos_protocols_ospf_access_list`

* New Resource `vyos_interfaces_ethernet_vif_s`

* New Resource `vyos_policy_local_route_rule`

* New Resource `vyos_vpn_ipsec_remote_access_radius_server`

* New Resource `vyos_protocols_ospf_redistribute_table`

* New Resource `vyos_service_salt_minion`

* New Resource `vyos_protocols_mpls_ldp_export_ipv6`

* New Resource `vyos_protocols_ospf_graceful_restart_helper_enable`

* New Resource `vyos_service_config_sync_section_system`

* New Resource `vyos_service_tftp_server_listen_address`

* New Resource `vyos_protocols_rpki_cache`

* New Resource `vyos_service_pppoe_server_client_ip_pool`

* New Resource `vyos_protocols_bgp_bmp`

* New Resource `vyos_service_conntrack_sync_failover_mechanism_vrrp`

* New Resource `vyos_system_login_tacacs_server`

* New Resource `vyos_interfaces_wireless_dhcpv6_options_pd_interface`

* New Resource `vyos_load_balancing_reverse_proxy_backend_server`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_aggregate_address`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_network`

* New Resource `vyos_protocols_static_table_route_interface`

* New Resource `vyos_protocols_babel_redistribute_ipv6`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_label_vpn_allocation_mode`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_kernel`

* New Resource `vyos_protocols_rip_redistribute_isis`

* New Resource `vyos_system_logs_logrotate_messages`

* New Resource `vyos_interfaces_vti`

* New Resource `vyos_protocols_ospf_auto_cost`

* New Resource `vyos_service_snmp_trap_target`

* New Resource `vyos_protocols_pim_rp_address`

* New Resource `vyos_system_config_management`

* New Resource `vyos_service_broadcast_relay`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_node_protecting_index`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_mac_vrf`

* New Resource `vyos_vpn_pptp_remote_access_ppp_options`

* New Resource `vyos_protocols_ospf_redistribute_static`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_advertise_ipv4_unicast`

* New Resource `vyos_protocols_bgp_parameters_bestpath_as_path`

* New Resource `vyos_protocols_isis_spf_delay_ietf`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_babel`

* New Resource `vyos_service_webproxy_listen_address`

* New Resource `vyos_interfaces_input`

* New Resource `vyos_vpn_sstp_ssl`

* New Resource `vyos_system_login_radius_server`

* New Resource `vyos_vpn_sstp_client_ipv6_pool_prefix`

* New Resource `vyos_interfaces_macsec`

* New Resource `vyos_protocols_nhrp_tunnel`

* New Resource `vyos_protocols_ripng_interface`

* New Resource `vyos_system_flow_accounting_netflow_timeout`

* New Resource `vyos_vpn_ipsec_esp_group_proposal`

* New Resource `vyos_protocols_isis_segment_routing_global_block`

* New Resource `vyos_service_ipoe_server_authentication_radius`

* New Resource `vyos_service_webproxy_url_filtering_squidguard_source_group`

* New Resource `vyos_system_login_user_authentication_public_keys`

* New Resource `vyos_interfaces_openvpn_server_push_route`

* New Resource `vyos_system_conntrack_tcp`

* New Resource `vyos_protocols_ospf_aggregation`

* New Resource `vyos_service_monitoring_telegraf_splunk_authentication`

* New Resource `vyos_system_console_device`

* New Resource `vyos_load_balancing_reverse_proxy_service_logging_facility`

* New Resource `vyos_protocols_static_route`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_aaaa`

* New Resource `vyos_interfaces_bridge_vif_dhcpv6_options_pd`

* New Resource `vyos_load_balancing_reverse_proxy_service`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_nexthop_vpn`

* New Resource `vyos_system_conntrack_log_event_update`

* New Resource `vyos_service_ids_ddos_protection_threshold_icmp`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_export`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_distance_prefix`

* New Resource `vyos_protocols_isis_traffic_engineering`

* New Resource `vyos_service_monitoring_telegraf`

* New Resource `vyos_service_webproxy_url_filtering`

* New Resource `vyos_protocols_failover_route`

* New Resource `vyos_system_login`

* New Resource `vyos_interfaces_pppoe_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_pseudo_ethernet_dhcpv6_options_pd`

* New Resource `vyos_interfaces_virtual_ethernet_vif_dhcpv6_options_pd`

* New Resource `vyos_service_pppoe_server_shaper`

* New Resource `vyos_protocols_bgp_parameters_confederation`

* New Resource `vyos_protocols_isis_redistribute_ipv4_bgp_level_1`

* New Resource `vyos_protocols_pim_igmp`

* New Resource `vyos_interfaces_ethernet_vif_dhcpv6_options_pd`

* New Resource `vyos_protocols_isis_redistribute_ipv6_babel_level_2`

* New Resource `vyos_protocols_rip_redistribute_babel`

* New Resource `vyos_service_snmp_v3_user`

* New Resource `vyos_protocols_babel_distribute_list_ipv6_prefix_list`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_distance_prefix`

* New Resource `vyos_protocols_isis_redistribute_ipv6_static_level_2`

* New Resource `vyos_service_snmp_mib`

* New Resource `vyos_interfaces_pseudo_ethernet_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv6_labeled_unicast_aggregate_address`

* New Resource `vyos_protocols_bgp_peer_group_local_role`

* New Resource `vyos_protocols_failover_route_next_hop`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet`

* New Resource `vyos_service_ipoe_server_interface`

* New Resource `vyos_protocols_mpls_ldp_export_ipv4`

* New Resource `vyos_vpn_openconnect_accounting_radius_server`

* New Resource `vyos_interfaces_openvpn`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_ripng`

* New Resource `vyos_protocols_ripng_redistribute_kernel`

* New Resource `vyos_interfaces_bonding`

* New Resource `vyos_protocols_bgp_parameters_bestpath_peer_type`

* New Resource `vyos_interfaces_bridge_vif`

* New Resource `vyos_protocols_bgp_address_family_ipv4_labeled_unicast_maximum_paths`

* New Resource `vyos_vpn_ipsec`

* New Resource `vyos_interfaces_virtual_ethernet_vif`

* New Resource `vyos_system_flow_accounting_netflow`

* New Resource `vyos_protocols_isis_segment_routing_prefix`

* New Resource `vyos_protocols_bgp_address_family_ipv6_flowspec_local_install`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_load_sharing_disable`

* New Resource `vyos_protocols_mpls_ldp_export_ipv4_export_filter`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_mx`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s_dhcpv6_options_pd`

* New Resource `vyos_load_balancing_reverse_proxy_service_http_response_headers`

* New Resource `vyos_service_ndp_proxy_interface_prefix`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_lowest_backup_metric_index`

* New Resource `vyos_protocols_static`

* New Resource `vyos_service_ipoe_server_shaper`

* New Resource `vyos_vpn_pptp_remote_access_limits`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_rip`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_flooding`

* New Resource `vyos_vpn_ipsec_esp_group`

* New Resource `vyos_service_suricata`

* New Resource `vyos_vpn_pptp_remote_access_extended_scripts`

* New Resource `vyos_protocols_pim`

* New Resource `vyos_protocols_ripng_redistribute_babel`

* New Resource `vyos_service_conntrack_sync`

* New Resource `vyos_interfaces_macsec_dhcpv6_options_pd`

* New Resource `vyos_protocols_isis_redistribute_ipv4_rip_level_1`

* New Resource `vyos_protocols_rpki`

* New Resource `vyos_policy_route6`

* New Resource `vyos_service_aws_glb_script`

* New Resource `vyos_protocols_bgp_address_family_ipv4_multicast_distance_prefix`

* New Resource `vyos_protocols_ospf_interface_authentication_md5_key_id`

* New Resource `vyos_service_ipoe_server_authentication_radius_server`

* New Resource `vyos_protocols_bgp_address_family_ipv6_multicast_aggregate_address`

* New Resource `vyos_protocols_isis_redistribute_ipv6_ripng_level_2`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv4_vpn_network`

* New Resource `vyos_nat_cgnat_pool_internal`

* New Resource `vyos_protocols_rip_redistribute_bgp`

* New Resource `vyos_system_option_kernel`

* New Resource `vyos_vpn_pptp_remote_access_authentication_local_users_username`

* New Resource `vyos_protocols_ospf_max_metric_router_lsa`

* New Resource `vyos_service_dns_forwarding`

* New Resource `vyos_vpn_ipsec_log`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_static`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_mx_server`

* New Resource `vyos_protocols_bgp_address_family_ipv6_multicast_distance_prefix`

* New Resource `vyos_vpn_ipsec_options`

* New Resource `vyos_protocols_bgp_listen_range`

* New Resource `vyos_protocols_ripng_redistribute_static`

* New Resource `vyos_service_monitoring_telegraf_influxdb`

* New Resource `vyos_system_flow_accounting_sflow_server`

* New Resource `vyos_system_wireless`

* New Resource `vyos_protocols_isis_redistribute_ipv4_static_level_2`

* New Resource `vyos_service_lldp`

* New Resource `vyos_load_balancing_reverse_proxy_service_rule`

* New Resource `vyos_load_balancing_wan`

* New Resource `vyos_system_syslog_global_marker`

* New Resource `vyos_vpn_pptp_remote_access`

* New Resource `vyos_protocols_rip_interface_authentication_md5`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_srv_entry`

* New Resource `vyos_service_monitoring_telegraf_loki_authentication`

* New Resource `vyos_system_task_scheduler_task`

* New Resource `vyos_protocols_bgp_address_family_ipv6_vpn_network`

* New Resource `vyos_protocols_mpls_ldp_targeted_neighbor_ipv6`

* New Resource `vyos_protocols_pim_rp`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_rd_vpn`

* New Resource `vyos_protocols_mpls_ldp_import_ipv4_import_filter`

* New Resource `vyos_system_conntrack_log_event_destroy`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_babel_interface`

* New Resource `vyos_service_ssh_dynamic_protection`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn`

* New Resource `vyos_protocols_igmp_proxy_interface`

* New Resource `vyos_protocols_ospf_refresh`

* New Resource `vyos_vpn_openconnect_network_settings_client_ip_settings`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_dhcpv6_options_pd`

* New Resource `vyos_protocols_ospf_area_virtual_link_authentication_md5_key_id`

* New Resource `vyos_service_ipoe_server_authentication_radius_rate_limit`

* New Resource `vyos_interfaces_virtual_ethernet_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv4_labeled_unicast_aggregate_address`

* New Resource `vyos_service_dns_dynamic`

* New Resource `vyos_interfaces_bridge`

* New Resource `vyos_load_balancing_wan_interface_health_test`

* New Resource `vyos_protocols_ripng`

* New Resource `vyos_protocols_static_route6`

* New Resource `vyos_protocols_ospf`

* New Resource `vyos_service_pppoe_server_interface`

* New Resource `vyos_service_snmp_script_extensions_extension_name`

* New Resource `vyos_system_login_tacacs`

* New Resource `vyos_interfaces_wireless_vif_s_dhcpv6_options_pd`

* New Resource `vyos_system_config_management_commit_archive`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s_dhcpv6_options_pd`

* New Resource `vyos_load_balancing_wan_interface_health`

* New Resource `vyos_protocols_bgp_address_family_ipv4_multicast_aggregate_address`

* New Resource `vyos_service_config_sync_secondary`

* New Resource `vyos_protocols_ospf_redistribute_rip`

* New Resource `vyos_service_https_certificates`

* New Resource `vyos_service_ssh`

* New Resource `vyos_vpn_sstp_snmp`

* New Resource `vyos_protocols_babel_distribute_list_ipv4_interface`

* New Resource `vyos_service_router_advert_interface`

* New Resource `vyos_service_suricata_log_eve`

* New Resource `vyos_protocols_bgp_parameters_conditional_advertisement`

* New Resource `vyos_protocols_isis_redistribute_ipv4_babel_level_2`

* New Resource `vyos_protocols_bgp_neighbor`

* New Resource `vyos_protocols_mpls_ldp_export_ipv6_export_filter`

* New Resource `vyos_system_flow_accounting_sflow`

* New Resource `vyos_interfaces_bonding_vif`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_aggregate_address`

* New Resource `vyos_protocols_bgp_neighbor_local_role`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_cname`

* New Resource `vyos_interfaces_pppoe_dhcpv6_options_pd`

* New Resource `vyos_service_config_sync_section`

* New Resource `vyos_vpn_pptp_remote_access_authentication_radius_server`

* New Resource `vyos_vpn_pptp_remote_access_client_ip_pool`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_babel`

* New Resource `vyos_protocols_isis_segment_routing_local_block`

* New Resource `vyos_vpn_sstp_shaper`

* New Resource `vyos_protocols_static_table_route6`

* New Resource `vyos_nat_cgnat_pool_external`

* New Resource `vyos_protocols_isis_redistribute_ipv4_bgp_level_2`

* New Resource `vyos_protocols_static_table_route6_next_hop`

* New Resource `vyos_service_pppoe_server_client_ipv6_pool`

* New Resource `vyos_vpn_sstp_authentication`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s`

* New Resource `vyos_protocols_isis_redistribute_ipv4_ospf_level_1`

* New Resource `vyos_vpn_openconnect_accounting_mode`

* New Resource `vyos_vpn_openconnect_authentication`

* New Resource `vyos_service_ipoe_server`

* New Resource `vyos_system_sflow_server`

* New Resource `vyos_interfaces_ethernet_vif_s_dhcpv6_options_pd`

* New Resource `vyos_load_balancing_wan_rule_interface`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_priority_limit_critical`

* New Resource `vyos_service_ipoe_server_client_ipv6_pool_delegate`

* New Resource `vyos_vpn_ipsec_site_to_site_peer_tunnel`

* New Resource `vyos_protocols_ospf_redistribute_kernel`

* New Resource `vyos_service_https_api_cors`

* New Resource `vyos_interfaces_ethernet`

* New Resource `vyos_protocols_isis`

* New Resource `vyos_protocols_ospf_segment_routing_local_block`

* New Resource `vyos_service_pppoe_server_authentication_radius_dynamic_author`

* New Resource `vyos_system_conntrack_timeout_custom_ipv4_rule`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_priority_limit_medium`

* New Resource `vyos_service_monitoring_zabbix_agent_server_active`

* New Resource `vyos_interfaces_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_sid_vpn`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet_range_option_static_route`

* New Resource `vyos_interfaces_bonding_vif_s_vif_c_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s_vif_c_dhcpv6_options_pd`

* New Resource `vyos_protocols_ospf_timers_throttle_spf`

* New Resource `vyos_protocols_isis_redistribute_ipv6_ospf6_level_1`

* New Resource `vyos_service_stunnel_log`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_export`

* New Resource `vyos_protocols_static_table_route6_next_hop_bfd_multi_hop_source`

* New Resource `vyos_system_sflow`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_connected`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_route_target_vpn`

* New Resource `vyos_protocols_mpls_ldp_allocation_ipv4`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet_static_mapping_option_static_route`

* New Resource `vyos_service_ids_ddos_protection_threshold_tcp`

* New Resource `vyos_service_lldp_legacy_protocols`

* New Resource `vyos_protocols_bgp_address_family_ipv4_multicast_network`

* New Resource `vyos_protocols_isis_redistribute_ipv4_kernel_level_2`

* New Resource `vyos_service_monitoring_telegraf_prometheus_client`

* New Resource `vyos_system_ip_arp`

* New Resource `vyos_interfaces_bonding_vif_s_dhcpv6_options_pd`

* New Resource `vyos_protocols_bgp_parameters_dampening`

* New Resource `vyos_protocols_mpls_ldp_discovery`

* New Resource `vyos_system_conntrack_modules`

* New Resource `vyos_system_conntrack`

* New Resource `vyos_interfaces_bonding_vif_s_dhcpv6_options_pd_interface`

* New Resource `vyos_load_balancing_reverse_proxy_backend_logging_facility`

* New Resource `vyos_protocols_pim_interface_igmp_join`

* New Resource `vyos_service_suricata_address_group`

* New Resource `vyos_interfaces_bonding_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_service_ipoe_server_client_ipv6_pool_prefix`

* New Resource `vyos_vpn_openconnect_network_settings_client_ipv6_pool`

* New Resource `vyos_protocols_isis_redistribute_ipv6_babel_level_1`

* New Resource `vyos_system_syslog`

* New Resource `vyos_protocols_isis_redistribute_ipv6_ospf6_level_2`

* New Resource `vyos_protocols_mpls`

* New Resource `vyos_protocols_pim_register_accept_list`

* New Resource `vyos_service_dns_dynamic_name`

* New Resource `vyos_interfaces_wireless`

* New Resource `vyos_protocols_bgp_address_family_ipv4_multicast_distance`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_remote_prefix_list`

* New Resource `vyos_vpn_ipsec_remote_access_connection`

* New Resource `vyos_service_https_api_graphql`

* New Resource `vyos_system_login_user`

* New Resource `vyos_system_static_host_mapping_host_name`

* New Resource `vyos_system_syslog_file_facility`

* New Resource `vyos_interfaces_wireless_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_ospf_segment_routing`

* New Resource `vyos_service_snmp_v3`

* New Resource `vyos_system_acceleration`

* New Resource `vyos_service_dns_forwarding_options`

* New Resource `vyos_service_https`

* New Resource `vyos_system_ip_tcp_mss`

* New Resource `vyos_protocols_rip_network_distance`

* New Resource `vyos_protocols_static_multicast_route`

* New Resource `vyos_system_conntrack_ignore_ipv6_rule`

* New Resource `vyos_service_https_allow_client`

* New Resource `vyos_service_webproxy_url_filtering_squidguard_auto_update`

* New Resource `vyos_protocols_static_table_route_next_hop`

* New Resource `vyos_system_option_kernel_debug`

* New Resource `vyos_protocols_bgp_peer_group`

* New Resource `vyos_protocols_eigrp`

* New Resource `vyos_protocols_isis_redistribute_ipv4_connected_level_1`

* New Resource `vyos_service_pppoe_server_pado_delay`

* New Resource `vyos_protocols_isis_redistribute_ipv4_ospf_level_2`

* New Resource `vyos_protocols_ripng_redistribute_ospfv3`

* New Resource `vyos_vpn_pptp_remote_access_client_ipv6_pool_delegate`

* New Resource `vyos_vpn_sstp`

* New Resource `vyos_vpn_pptp_remote_access_snmp`

* New Resource `vyos_system_syslog_global`

* New Resource `vyos_interfaces_wireless_vif_s`

* New Resource `vyos_service_ipoe_server_log`

* New Resource `vyos_service_webproxy_url_filtering_squidguard_time_period`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_nhrp_tunnel_dynamic_map`

* New Resource `vyos_service_monitoring_zabbix_agent_limits`

* New Resource `vyos_system_frr`

* New Resource `vyos_service_stunnel_server_psk`

* New Resource `vyos_system_update_check`

* New Resource `vyos_vpn_openconnect_authentication_identity_based_config`

* New Resource `vyos_protocols_bgp_parameters_distance_prefix`

* New Resource `vyos_system_syslog_user`

* New Resource `vyos_vpn_openconnect_authentication_mode`

* New Resource `vyos_protocols_bgp_bmp_target`

* New Resource `vyos_service_dhcp_server_shared_network_name_option_static_route`

* New Resource `vyos_service_https_api`

* New Resource `vyos_vpn_pptp_remote_access_authentication`

* New Resource `vyos_interfaces_macsec_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_wireless_security_wpa_radius_server`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_connected`

* New Resource `vyos_system_login_radius`

* New Resource `vyos_interfaces_wireless_vif_s_vif_c_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_static_arp_interface`

* New Resource `vyos_system_syslog_host`

* New Resource `vyos_protocols_bgp_parameters_distance_global`

* New Resource `vyos_system_option_http_client`

* New Resource `vyos_vpn_sstp_log`

* New Resource `vyos_system_conntrack_log`

* New Resource `vyos_interfaces_wwan_dhcpv6_options_pd_interface`

* New Resource `vyos_service_console_server_device`

* New Resource `vyos_service_snmp_community`

* New Resource `vyos_policy_local_route6_rule`

* New Resource `vyos_service_ndp_proxy_interface`

* New Resource `vyos_service_snmp`

* New Resource `vyos_protocols_isis_redistribute_ipv6_static_level_1`

* New Resource `vyos_protocols_segment_routing_interface`

* New Resource `vyos_protocols_static_table`

* New Resource `vyos_service_config_sync`

* New Resource `vyos_interfaces_sstpc`

* New Resource `vyos_protocols_bgp_sid_vpn_per_vrf`

* New Resource `vyos_service_stunnel_server`

* New Resource `vyos_system_conntrack_timeout_custom_ipv6_rule`

* New Resource `vyos_interfaces_wwan_dhcpv6_options_pd`

* New Resource `vyos_vpn_openconnect_network_settings`

* New Resource `vyos_vpn_sstp_extended_scripts`

* New Resource `vyos_system`

* New Resource `vyos_service_webproxy_url_filtering_squidguard`

* New Resource `vyos_protocols_bgp_parameters`

* New Resource `vyos_protocols_pim_ssm`

* New Resource `vyos_protocols_rip_default_information`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_advertise_ipv6_unicast`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_isis`

* New Resource `vyos_protocols_ospf_area_virtual_link`

* New Resource `vyos_service_router_advert_interface_prefix`

* New Resource `vyos_protocols_bfd_profile`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_route_map_vpn`

* New Resource `vyos_interfaces_macsec_security_static_peer`

* New Resource `vyos_protocols_isis_segment_routing`

* New Resource `vyos_system_proxy`

* New Resource `vyos_interfaces_bonding_vif_dhcpv6_options_pd`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_distance`

* New Resource `vyos_protocols_rip_redistribute_connected`

* New Resource `vyos_vpn_openconnect_authentication_radius_server`

* New Resource `vyos_service_ipoe_server_authentication_radius_dynamic_author`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_ospf`

* New Resource `vyos_protocols_ospf_neighbor`

* New Resource `vyos_service_aws_glb_status`

* New Resource `vyos_service_sla_twamp_server`

* New Resource `vyos_vpn_sstp_authentication_radius_server`

* New Resource `vyos_nat_cgnat_pool_external_range`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_default_originate`

* New Resource `vyos_service_pppoe_server_authentication_radius_rate_limit`

* New Resource `vyos_interfaces_wwan`

* New Resource `vyos_service_config_sync_section_interfaces`

* New Resource `vyos_service_pppoe_server_log`

* New Resource `vyos_interfaces_bonding_vif_s_vif_c`

* New Resource `vyos_protocols_ospf_parameters`

* New Resource `vyos_protocols_ospf_redistribute_connected`

* New Resource `vyos_service_webproxy_cache_peer`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_kernel`

* New Resource `vyos_protocols_ospf_segment_routing_prefix`

* New Resource `vyos_interfaces_wireless_vif_s_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_static_route_next_hop`

* New Resource `vyos_service_config_sync_section_qos`

* New Resource `vyos_protocols_isis_redistribute_ipv4_babel_level_1`

* New Resource `vyos_service_webproxy`

* New Resource `vyos_vpn_sstp_client_ipv6_pool`

* New Resource `vyos_protocols_bgp_parameters_default`

* New Resource `vyos_service_config_sync_section_service`

* New Resource `vyos_service_pppoe_server_client_ipv6_pool_delegate`

* New Resource `vyos_protocols_pim_interface`

* New Resource `vyos_service_dhcp_server_shared_network_name`

* New Resource `vyos_load_balancing_reverse_proxy_backend_http_response_headers`

* New Resource `vyos_protocols_ospf_area_range`

* New Resource `vyos_protocols_ripng_distribute_list_access_list`

* New Resource `vyos_system_ip_multipath`

* New Resource `vyos_service_pppoe_server_authentication_local_users_username`

* New Resource `vyos_service_dhcp_server`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s`

* New Resource `vyos_protocols_isis_default_information_originate_ipv6_level_1`

* New Resource `vyos_protocols_static_route6_interface`

* New Resource `vyos_load_balancing_reverse_proxy_global_parameters_logging_facility`

* New Resource `vyos_protocols_static_route6_next_hop`

* New Resource `vyos_protocols_static_route_next_hop_bfd_multi_hop_source`

* New Resource `vyos_system_frr_snmp`

* New Resource `vyos_system_logs_logrotate_atop`

* New Resource `vyos_system_syslog_host_facility`

* New Resource `vyos_vpn_sstp_limits`

* New Resource `vyos_protocols_ripng_redistribute_connected`

* New Resource `vyos_service_dns_forwarding_authoritative_domain`

* New Resource `vyos_protocols_ospf_summary_address`

* New Resource `vyos_service_pppoe_server_authentication_radius_server`

* New Resource `vyos_protocols_bgp_parameters_tcp_keepalive`

* New Resource `vyos_service_monitoring_telegraf_prometheus_client_authentication`

* New Resource `vyos_service_webproxy_url_filtering_squidguard_rule`

* New Resource `vyos_vpn_ipsec_ike_group_proposal`

* New Resource `vyos_protocols_babel_distribute_list_ipv4_prefix_list`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_naptr_rule`

* New Resource `vyos_service_mdns_repeater`

* New Resource `vyos_service_ipoe_server_authentication_interface_mac`

* New Resource `vyos_policy_route`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_static`

* New Resource `vyos_protocols_ospf_graceful_restart_helper`

* New Resource `vyos_protocols_ospf_interface`

* New Resource `vyos_system_ip`

* New Resource `vyos_interfaces_bonding_dhcpv6_options_pd`

* New Resource `vyos_service_https_api_graphql_authentication`

* New Resource `vyos_service_suricata_port_group`

* New Resource `vyos_protocols_igmp_proxy`

* New Resource `vyos_vpn_pptp_remote_access_shaper`

* New Resource `vyos_load_balancing_reverse_proxy_backend_rule`

* New Resource `vyos_policy_route6_rule`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_label_vpn`

* New Resource `vyos_protocols_static_table_route6_interface`

* New Resource `vyos_service_pppoe_server_extended_scripts`

* New Resource `vyos_interfaces_wireguard`

* New Resource `vyos_load_balancing_reverse_proxy_backend`

* New Resource `vyos_protocols_mpls_ldp_parameters`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_ptr`

* New Resource `vyos_service_monitoring_telegraf_splunk`

* New Resource `vyos_service_pppoe_server_authentication_radius`

* New Resource `vyos_interfaces_bonding_vif_s_vif_c_dhcpv6_options_pd`

* New Resource `vyos_interfaces_wireless_vif`

* New Resource `vyos_protocols_isis_redistribute_ipv6_ripng_level_1`

* New Resource `vyos_interfaces_bonding_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_ospf_segment_routing_global_block`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_maximum_paths`

* New Resource `vyos_service_dns_forwarding_name_server`

* New Resource `vyos_protocols_bgp_address_family_ipv6_labeled_unicast_network`

* New Resource `vyos_protocols_static_route6_next_hop_bfd_multi_hop_source`

* New Resource `vyos_vpn_pptp_remote_access_authentication_radius`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s_vif_c`

* New Resource `vyos_interfaces_wireless_dhcpv6_options_pd`

* New Resource `vyos_protocols_bgp_address_family_ipv4_flowspec_local_install`

* New Resource `vyos_protocols_isis_redistribute_ipv6_kernel_level_2`

* New Resource `vyos_vpn_sstp_authentication_radius`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_network`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute`

* New Resource `vyos_interfaces_bonding_vif_s`

* New Resource `vyos_interfaces_geneve`

* New Resource `vyos_protocols_static_multicast_interface_route_next_hop_interface`

* New Resource `vyos_service_dhcp_relay`

* New Resource `vyos_service_ipoe_server_authentication_interface`

* New Resource `vyos_interfaces_wireless_vif_s_vif_c_dhcpv6_options_pd`

* New Resource `vyos_load_balancing_reverse_proxy_global_parameters`

* New Resource `vyos_protocols_ospf_mpls_te`

* New Resource `vyos_vpn_ipsec_remote_access_pool`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_vni`

* New Resource `vyos_service_ids_ddos_protection`

* New Resource `vyos_protocols_isis_interface`

* New Resource `vyos_vpn_sstp_authentication_radius_dynamic_author`

* New Resource `vyos_vpn_openconnect_authentication_local_users_username`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_ead_es_route_target`

* New Resource `vyos_service_snmp_v3_view_oid`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_route_target`

* New Resource `vyos_system_option`

* New Resource `vyos_protocols_babel_parameters`

* New Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_ead_es_frag`

* New Resource `vyos_protocols_isis_redistribute_ipv6_kernel_level_1`

* New Resource `vyos_service_ipoe_server_client_ipv6_pool`

* New Resource `vyos_service_ids_ddos_protection_sflow`

* New Resource `vyos_service_pppoe_server_ppp_options`

* New Resource `vyos_service_webproxy_authentication`

* New Resource `vyos_system_syslog_user_facility`

* New Resource `vyos_nat_cgnat_rule`

* New Resource `vyos_protocols_mpls_ldp`

* New Resource `vyos_protocols_mpls_ldp_targeted_neighbor_ipv4`

* New Resource `vyos_vpn_ipsec_authentication_psk`

* New Resource `vyos_vpn_ipsec_profile`

* New Resource `vyos_interfaces_ethernet_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_ethernet_vif_s_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_isis_redistribute_ipv6_connected_level_2`

* New Resource `vyos_protocols_ospf_redistribute_isis`

* New Resource `vyos_service_broadcast_relay_id`

* New Resource `vyos_protocols_isis_redistribute_ipv6_bgp_level_2`

* New Resource `vyos_protocols_pim_spt_switchover_infinity_and_beyond`

* New Resource `vyos_system_ip_nht`

* New Resource `vyos_protocols_ospf_capability`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet_range`

* New Resource `vyos_service_ntp_allow_client`

* New Resource `vyos_protocols_isis_default_information_originate_ipv4_level_1`

* New Resource `vyos_service_monitoring_zabbix_agent`

* New Resource `vyos_service_ipoe_server_snmp`

* New Resource `vyos_vpn_openconnect_ssl`

* New Resource `vyos_service_ipoe_server_limits`

* New Resource `vyos_service_ndp_proxy`

* New Resource `vyos_system_login_banner`

* New Resource `vyos_interfaces_wireless_vif_s_vif_c`

* New Resource `vyos_protocols_static_route_interface`

* New Resource `vyos_service_dns_forwarding_domain_name_server`

* New Resource `vyos_system_flow_accounting_netflow_server`

* New Resource `vyos_interfaces_pppoe`

* New Resource `vyos_protocols_isis_redistribute_ipv4_kernel_level_1`

* New Resource `vyos_protocols_ospf_graceful_restart`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute`

* New Resource `vyos_service_dns_forwarding_domain`

* New Resource `vyos_protocols_ospf_log_adjacency_changes`

* New Resource `vyos_protocols_pim_ecmp`

* New Resource `vyos_service_webproxy_authentication_ldap`

* New Resource `vyos_protocols_isis_domain_password`

* New Resource `vyos_protocols_mpls_ldp_neighbor`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_import`

* New Resource `vyos_protocols_ospf_distance`

* New Resource `vyos_protocols_rip_distribute_list_prefix_list`

* New Resource `vyos_protocols_ospf_ldp_sync`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_a`

* New Resource `vyos_service_router_advert_interface_route`

* New Resource `vyos_protocols_rip`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_srv`

* New Resource `vyos_service_ssh_rekey`

* New Resource `vyos_vpn_pptp_remote_access_authentication_radius_rate_limit`

* New Resource `vyos_protocols_isis_redistribute_ipv4_connected_level_2`

* New Resource `vyos_vpn_sstp_ppp_options`

* New Resource `vyos_interfaces_ethernet_vif`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_babel_distribute_list_ipv6_interface`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_route_target_vpn`

* New Resource `vyos_protocols_nhrp_tunnel_map`

* New Resource `vyos_protocols_static_neighbor_proxy_arp`

* New Resource `vyos_system_syslog_file`

* New Resource `vyos_interfaces_bridge_member_interface`

* New Resource `vyos_load_balancing_wan_sticky_connections`

* New Resource `vyos_protocols_isis_ldp_sync`

* New Resource `vyos_protocols_mpls_ldp_allocation_ipv6`

* New Resource `vyos_protocols_ospf_distance_ospf`

* New Resource `vyos_service_snmp_v3_view`

* New Resource `vyos_protocols_babel_redistribute_ipv4`

* New Resource `vyos_protocols_isis_redistribute_ipv6_bgp_level_1`

* New Resource `vyos_service_dhcp_server_high_availability`

* New Resource `vyos_service_stunnel_client_psk`

* New Resource `vyos_protocols_bgp_neighbor_local_as`

* New Resource `vyos_interfaces_pseudo_ethernet_vif_s_vif_c_dhcpv6_options_pd`

* New Resource `vyos_interfaces_virtual_ethernet_dhcpv6_options_pd_interface`

* New Resource `vyos_interfaces_vxlan`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_maximum_paths`

* New Resource `vyos_interfaces_pseudo_ethernet_vif`

* New Resource `vyos_system_syslog_global_facility`

* New Resource `vyos_vpn_sstp_client_ipv6_pool_delegate`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_distance`

* New Resource `vyos_protocols_eigrp_metric`

* New Resource `vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_downstream_index`

* New Resource `vyos_protocols_static_multicast_interface_route`

* New Resource `vyos_protocols_bgp_timers`

* New Resource `vyos_protocols_isis_redistribute_ipv4_static_level_1`

* New Resource `vyos_protocols_ospf_area`

* New Resource `vyos_system_console`

* New Resource `vyos_interfaces_bridge_dhcpv6_options_pd`

* New Resource `vyos_protocols_bgp_interface`

* New Resource `vyos_vpn_openconnect_authentication_radius`

* New Resource `vyos_interfaces_loopback`

* New Resource `vyos_vpn_openconnect_listen_ports`

* New Resource `vyos_protocols_bfd_peer`

* New Resource `vyos_service_https_api_keys_id`

* New Resource `vyos_service_ssh_access_control_deny`

* New Resource `vyos_service_ipoe_server_extended_scripts`

* New Resource `vyos_system_option_ssh_client`

* New Resource `vyos_interfaces_virtual_ethernet_vif_s_vif_c`

* New Resource `vyos_load_balancing_wan_rule`

* New Resource `vyos_service_snmp_listen_address`

* New Resource `vyos_system_conntrack_ignore_ipv4_rule`

* New Resource `vyos_nat_cgnat`

* New Resource `vyos_vpn_sstp_client_ip_pool`

* New Resource `vyos_service_pppoe_server`

* New Resource `vyos_protocols_static_multicast_route_next_hop`

* New Resource `vyos_protocols_static_table_route`

* New Resource `vyos_service_monitoring_telegraf_loki`

* New Resource `vyos_service_ipoe_server_authentication`

* New Resource `vyos_service_monitoring_zabbix_agent_log`

* New Resource `vyos_system_lcd`

* New Resource `vyos_interfaces_wireless_vif_dhcpv6_options_pd`

* New Resource `vyos_protocols_babel_distribute_list_ipv4_access_list`

* New Resource `vyos_service_dhcp_server_shared_network_name_subnet_option_static_route`

* New Resource `vyos_interfaces_bridge_vif_dhcpv6_options_pd_interface`

* New Resource `vyos_protocols_ripng_timers`

* New Resource `vyos_protocols_static_arp_interface_address`

* New Resource `vyos_vpn_pptp_remote_access_client_ipv6_pool_prefix`

* New Resource `vyos_load_balancing_reverse_proxy`

* New Resource `vyos_service_dns_forwarding_authoritative_domain_records_naptr`

* New Resource `vyos_service_snmp_v3_trap_target`

* New Resource `vyos_protocols_bgp_peer_group_local_as`

* New Resource `vyos_service_lldp_interface`

* New Resource `vyos_protocols_ripng_distribute_list_prefix_list`

* New Resource `vyos_service_ssh_access_control_allow`

* New Resource `vyos_vpn_sstp_authentication_local_users_username`

* New Resource `vyos_protocols_bgp_parameters_graceful_restart`

* New Resource `vyos_service_monitoring_telegraf_influxdb_authentication`

* New Resource `vyos_service_pppoe_server_limits`

* New Resource `vyos_service_router_advert_interface_nat64prefix`

* New Resource `vyos_interfaces_ethernet_vif_s_vif_c_dhcpv6_options_pd`

* New Resource `vyos_policy_route_rule`

* New Resource `vyos_vpn_pptp_remote_access_authentication_radius_dynamic_author`

* New Resource `vyos_protocols_segment_routing_srv6_locator`

* New Resource `vyos_vpn_sstp_authentication_radius_rate_limit`

* New Resource `vyos_protocols_bgp_listen`

* New Resource `vyos_protocols_nhrp_tunnel_shortcut_target`

* New Resource `vyos_protocols_rip_interface`









## 0.4.20240813 (2024-08-13 08:24:24 UTC)
NOTES:

* update to rolling release 2024-08-13T00:21:40Z

* update to rolling release 2024-08-12T00:21:49Z

* update to rolling release 2024-08-11T00:23:16Z

* update to rolling release 2024-08-10T00:20:35Z

* update to rolling release 2024-08-09T00:20:49Z

* update to rolling release 2024-08-08T00:20:40Z





## 0.4.20240806 (2024-08-06 08:24:19 UTC)
NOTES:

* update to rolling release 2024-08-06T00:20:41Z

* update to rolling release 2024-08-04T00:22:49Z





## 0.4.20240803 (2024-08-03 08:23:39 UTC)
NOTES:

* update to rolling release 2024-08-03T00:20:04Z

* update to rolling release 2024-08-01T10:47:52Z





## 0.4.20240801 (2024-08-01 08:24:17 UTC)
NOTES:

* update to rolling release 2024-08-01T00:23:04Z





## 0.4.20240725 (2024-07-30 08:24:02 UTC)
NOTES:

* update to rolling release 2024-07-30T00:20:39Z

* update to rolling release 2024-07-28T00:23:03Z

* update to rolling release 2024-07-27T00:19:54Z

* update to rolling release 2024-07-26T00:20:13Z


FEATURES:

* Resource/`vyos_firewall_ipv4_forward_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_forward_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv4_input_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_input_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv4_name_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_name_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv4_output_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_output_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv4_output_raw_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_output_raw_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv4_prerouting_raw_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_prerouting_raw_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_forward_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_forward_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_input_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_input_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_name_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_name_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_output_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_output_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_output_raw_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_output_raw_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_prerouting_raw_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_prerouting_raw_rule` Added attribute `match_none`





## 0.3.20240725 (2024-07-25 08:24:01 UTC)
NOTES:

* update to rolling release 2024-07-25T00:20:19Z


BUG FIXES:

* generate files with relative path in autogen comment

* move autogen comment to use relative path to avoid changes between machines





## 0.3.20240717 (2024-07-21 08:22:11 UTC)
NOTES:

* update to rolling release 2024-07-17T17:05:42Z





## 0.3.20240710 (2024-07-10 08:29:07 UTC)
NOTES:

* update to rolling release 2024-07-10T00:20:26Z





## 0.3.20240709 (2024-07-09 08:22:43 UTC)
NOTES:

* update to rolling release 2024-07-09T00:20:00Z





## 0.3.20240707 (2024-07-07 08:22:13 UTC)
NOTES:

* update to rolling release 2024-07-07T00:22:17Z





## 0.3.20240706 (2024-07-06 08:23:33 UTC)
NOTES:

* update to rolling release 2024-07-06T00:19:12Z





## 0.3.20240705 (2024-07-05 08:24:23 UTC)
NOTES:

* update to rolling release 2024-07-05T00:19:47Z





## 0.3.20240704 (2024-07-04 08:23:56 UTC)
NOTES:

* update to rolling release 2024-07-04T00:19:47Z





## 0.3.20240702 (2024-07-03 08:23:41 UTC)
NOTES:

* update to rolling release 2024-07-02T18:39:02Z





## 0.3.20240701 (2024-07-01 08:23:44 UTC)
NOTES:

* update to rolling release 2024-07-01T00:23:26Z





## 0.3.20240630 (2024-06-30 08:21:28 UTC)
NOTES:

* update to rolling release 2024-06-30T00:22:05Z





## 0.3.20240629 (2024-06-29 08:22:20 UTC)
NOTES:

* update to rolling release 2024-06-29T00:19:21Z





## 0.3.20240628 (2024-06-28 08:22:19 UTC)
NOTES:

* update to rolling release 2024-06-28T00:19:42Z





## 0.3.20240627 (2024-06-27 08:23:29 UTC)
NOTES:

* update to rolling release 2024-06-27T00:19:46Z





## 0.3.20240626 (2024-06-26 08:23:08 UTC)
NOTES:

* update to rolling release 2024-06-26T00:19:30Z





## 0.3.20240625 (2024-06-25 08:23:29 UTC)
NOTES:

* update to rolling release 2024-06-25T00:19:32Z





## 0.3.20240623 (2024-06-23 08:21:01 UTC)
NOTES:

* update to rolling release 2024-06-23T00:21:32Z





## 0.3.20240621 (2024-06-21 08:22:55 UTC)
NOTES:

* update to rolling release 2024-06-21T00:19:18Z





## 0.3.20240620 (2024-06-20 08:23:39 UTC)
NOTES:

* update to rolling release 2024-06-20T00:19:27Z





## 0.3.20240619 (2024-06-19 08:23:09 UTC)
NOTES:

* update to rolling release 2024-06-19T00:19:47Z





## 0.3.20240617 (2024-06-17 08:23:38 UTC)
NOTES:

* update to rolling release 2024-06-17T00:20:56Z





## 0.3.20240613 (2024-06-13 08:23:27 UTC)
NOTES:

* update to rolling release 2024-06-13T00:19:41Z





## 0.3.20240612 (2024-06-12 08:22:53 UTC)
NOTES:

* update to rolling release 2024-06-12T00:19:44Z





## 0.3.20240606 (2024-06-06 08:23:02 UTC)
NOTES:

* update to rolling release 2024-06-06T00:19:21Z





## 0.3.20240602 (2024-06-02 08:21:22 UTC)
NOTES:

* update to rolling release 2024-06-02T00:21:06Z





## 0.3.20240601 (2024-06-01 08:23:16 UTC)
NOTES:

* update to rolling release 2024-06-01T00:20:56Z





## 0.3.20240531 (2024-05-31 08:27:03 UTC)
NOTES:

* update to rolling release 2024-05-31T00:18:53Z





## 0.3.20240528 (2024-05-28 08:24:26 UTC)
NOTES:

* update to rolling release 2024-05-28T00:19:17Z





## 0.3.20240527 (2024-05-27 08:23:01 UTC)
NOTES:

* update to rolling release 2024-05-27T00:19:48Z





## 0.3.20240526 (2024-05-26 08:22:01 UTC)
NOTES:

* update to rolling release 2024-05-26T00:21:03Z





## 0.3.20240524 (2024-05-24 08:23:39 UTC)
NOTES:

* update to rolling release 2024-05-24T00:19:23Z





## 0.3.20240522 (2024-05-22 08:23:17 UTC)
NOTES:

* update to rolling release 2024-05-22T00:18:44Z





## 0.3.20240521 (2024-05-21 08:22:14 UTC)
NOTES:

* update to rolling release 2024-05-21T00:19:19Z





## 0.3.20240520 (2024-05-20 08:22:55 UTC)
NOTES:

* update to rolling release 2024-05-20T00:19:28Z





## 0.3.20240518 (2024-05-19 08:20:42 UTC)
NOTES:

* update to rolling release 2024-05-18T12:11:15Z





## 0.3.20240514 (2024-05-14 08:23:36 UTC)
NOTES:

* update to rolling release 2024-05-14T00:18:58Z





## 0.3.20240512 (2024-05-13 08:23:13 UTC)
NOTES:

* update to rolling release 2024-05-12T14:02:26Z

* Add notetice about provider being incomplete





## 0.3.20240510 (2024-05-10 08:21:55 UTC)
NOTES:

* update to rolling release 2024-05-10T00:18:53Z





## 0.3.20240509 (2024-05-09 08:22:09 UTC)
NOTES:

* update to rolling release 2024-05-09T00:18:37Z





## 0.3.20240508 (2024-05-08 08:23:55 UTC)
NOTES:

* update to rolling release 2024-05-08T00:15:46Z





## 0.3.20240507 (2024-05-07 08:21:46 UTC)
NOTES:

* update to rolling release 2024-05-07T00:18:45Z





## 0.3.20240506 (2024-05-06 10:32:57 UTC)
NOTES:

* update to rolling release 2024-05-06T00:19:12Z

* attempt to add changelog to release notes





## 0.3.20240505 (2024-05-05 13:02:39 UTC)
BUG FIXES:

* makefile improvements and CI fixes





## 0.3.20240323 (2024-05-01 15:38:08 UTC)
FEATURES:

* enable daily automatic release to match rolling release

* prepare ci-update make target

* Move devcontainer to Dockerfile

* Resource/`vyos_container_name` Added attribute `cap_add`

* Resource/`vyos_high_availability_vrrp_group` Added attribute `excluded_address`


BUG FIXES:

* auto update github action





## 0.2.20240323 (2024-04-28 12:32:19 UTC)
FEATURES:

* use goreleaser to release provider


BUG FIXES:

* add missing manifest required by goreleaser

* removed unintended increase in fix version





## 0.2.20240324 (2024-04-28 12:20:45 UTC)
FEATURES:

* use goreleaser to release provider





## 0.1.20240323 (2024-04-27 17:21:15 UTC)
NOTES:

* beginning of changelog generation


FEATURES:

* Autogenerate changelog


BUG FIXES:

* Improve publishing target

* publish packaging zips and docs





## 0.0.20240322 (2024-03-22 00:00:00 UTC)

Initial release
