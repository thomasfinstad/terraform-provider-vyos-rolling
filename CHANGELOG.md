
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 2.7.202409090 (2024-09-09 08-28-41 UTC)](#release-27202409090-2024-09-09-08-28-41-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
    - [Schema changes](#schema-changes)
      - [Features](#features)
        - [Resources](#resources)
  - [Release 2.6.202408230 (2024-08-30 00-43-22 UTC)](#release-26202408230-2024-08-30-00-43-22-utc)
    - [Project changes](#project-changes-1)
      - [Notes](#notes-1)
      - [Bug fixes](#bug-fixes)
    - [Schema changes](#schema-changes-1)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources-1)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 2.7.202409090 (2024-09-09 08-28-41 UTC)
### Project changes
#### Notes
* update to rolling release 2024-09-09T00:06:05Z

### Schema changes
#### Features

##### Resources
* Modified Resource `vyos_firewall_global_options_apply_to_bridged_traffic`
	* New attribute `invalid_connections`

* Modified Resource `vyos_firewall_bridge_prerouting_filter_rule`
	* New attribute `ethernet_type`

* Modified Resource `vyos_firewall_bridge_input_filter_rule`
	* New attribute `ethernet_type`

* Modified Resource `vyos_service_router_advert_interface`
	* New attribute `no_send_interval`

* Modified Resource `vyos_firewall_bridge_name_rule`
	* New attribute `ethernet_type`

* Modified Resource `vyos_firewall_bridge_output_filter_rule`
	* New attribute `ethernet_type`

* Modified Resource `vyos_firewall_bridge_forward_filter_rule`
	* New attribute `ethernet_type`









## Release 2.6.202408230 (2024-08-30 00-43-22 UTC)
### Project changes
#### Notes
* update to rolling release 2024-08-23T00:21:01Z
#### Bug fixes
* resource identifier attribute

### Schema changes
#### BREAKING CHANGES

##### Resources
* Modified Resource `vyos_protocols_static_table_route_next_hop`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_rpki_cache`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_priority_queue`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_tftp_server_listen_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_ip_protocol`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv4_input_filter_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_labeled_unicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_naptr`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ipoe_server_client_ipv6_pool_delegate`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_access_list_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_failover_route_next_hop`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_route_next_hop_bfd_multi_hop_source`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_ip_protocol`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bridge_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_vif_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_sstp_client_ipv6_pool_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_container_registry`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_babel_distribute_list_ipv6_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_webproxy_url_filtering_squidguard_time_period`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_static_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_openvpn_server_push_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_syslog_host_facility`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_babel_distribute_list_ipv4_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_neighbor_local_as`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_pppoe_server_client_ipv6_pool`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_static_route6_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_community_list`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif_s_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_bridge_input_filter_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_rip_network_distance`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_table_route6_next_hop`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_shaper_class`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv6_name_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_route6_next_hop`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_vpn_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_nat_cgnat_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_router_advert_interface_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_as_path_list_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_sysctl_parameter`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_srv_entry`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_container_name_device`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_interface_authentication_md5_key_id`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_segment_routing_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_high_availability_vrrp_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_vif_s`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_vif_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_sstp_client_ip_pool`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_openvpn`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pppoe`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_nat_static_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_pptp_remote_access_authentication_local_users_username`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireguard`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_table_route_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_segment_routing_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_route6_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_pki_key_pair`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_stunnel_server_psk`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_vif`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_bridge_prerouting_filter_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_babel_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_high_availability_vrrp_group_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_console_server_device`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_vif_s`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_lldp_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_sstp_authentication_local_users_username`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_mx`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ipoe_server_authentication_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dhcp_server_shared_network_name`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_snmp_listen_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_wan_rule_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_redistribute_table`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_large_community_list`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_pptp_remote_access_client_ipv6_pool`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_distance_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_area_virtual_link`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_peer_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_limiter`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_random_detect_precedence`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_isis_segment_routing_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv4_name_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_igmp_proxy_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_parameters_distance_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dhcp_server_shared_network_name_subnet_range_option_static_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_pppoe_server_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_zone_from`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_netns_name`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_failover_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_pki_ca`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_router_advert_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_neighbor_local_role`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wwan_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_reverse_proxy_service_http_response_headers`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_remote_access_pool`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_pptp_remote_access_client_ipv6_pool_delegate`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_unicast_distance_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_multicast_distance_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_group_dynamic_group_address_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_event_handler_event`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_ipv6_protocol`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_downstream_index`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_container_name_port`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_nhrp_tunnel_dynamic_map`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_flow_accounting_netflow_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_high_availability_virtual_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_nat_cgnat_pool_external`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_extcommunity_list`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_shaper_hfsc`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_ptr`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_bridge_name_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_access_list`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ipoe_server_client_ipv6_pool_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_multicast_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_pppoe_server_authentication_local_users_username`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_pim_rp_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_nat_destination_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_container_name`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_route6_next_hop_bfd_multi_hop_source`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_task_scheduler_task`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_neighbor_local_as`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_bridge_forward_filter_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dhcp_server_shared_network_name_subnet`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_ns`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_neighbor`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_area`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_vif_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_vif_s_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_fair_queue`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_labeled_unicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_snmp_v3_view`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_route_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dhcp_server_shared_network_name_subnet_range`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_sstp_client_ipv6_pool_delegate`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_vif_s_vif_c_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_wan_interface_health`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_profile`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_pim_interface_igmp_join`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_distance_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_area_virtual_link`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_area_range`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_snmp_script_extensions_extension_name`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_router_advert_interface_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_labeled_unicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_prefix_list`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_table_route_next_hop_bfd_multi_hop_source`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_esp_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_parameters_distance_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_high_availability_virtual_server_real_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_login_radius_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dhcp_server_shared_network_name_subnet_static_mapping_option_static_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_conntrack_sync_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_spf`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_vif_s_vif_c_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_vxlan_vlan_to_vni`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_route_map`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ipoe_server_authentication_interface_mac`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif_s`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bfd_profile`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_redistribute_table`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_vif_s_vif_c`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_flow_accounting_sflow_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_unicast_distance_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_peer_group_local_role`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_area`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_vif_s_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_security_wpa_radius_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_webproxy_url_filtering_squidguard_source_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_unicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_segment_routing_srv6_locator`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_priority_queue_class_match`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_reverse_proxy_backend_logging_facility`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_multicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_bridge_output_filter_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_unicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_name_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif_s_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_traffic_match_group_match`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv4_output_filter_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wwan`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bfd_peer`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_neighbor_proxy_and`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_route6`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_vif`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_multicast_interface_route_next_hop_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ipoe_server_client_ipv6_pool`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_site_to_site_peer`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_vpn_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_vif_s_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_nat_cgnat_pool_internal`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_isis_fast_reroute_lfa_remote_prefix_list`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_webproxy_cache_peer`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_labeled_unicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_listen_range`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_mpls_ldp_neighbor`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_route6`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_isis_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_dummy`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_nat_destination_rule_load_balance_backend`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_table`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_broadcast_relay_id`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_static_route6_next_hop`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_unicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_node_protecting_index`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_conntrack_ignore_ipv4_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_openconnect_accounting_radius_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_vif_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_reverse_proxy_global_parameters_logging_facility`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_drop_tail`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_rate_control`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_snmp_v3_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_labeled_unicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_vif_s_vif_c`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_console_device`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_vif_s_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_vif_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_login_tacacs_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_rip_distribute_list_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_community_list_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_isis_segment_routing_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_srv`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif_s_vif_c`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wwan_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_peer_group_local_as`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_peer_group_local_role`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_round_robin_class_match`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_webproxy_listen_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bridge_vif_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_stunnel_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_suricata_port_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_mx_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_esp_group_proposal`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_prefix_list6_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_pppoe_server_pado_delay`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireguard_peer`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_container_name_volume`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_vpn_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_vif_s_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_multicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_suricata_address_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_remote_access_connection`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_node_protecting_index`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_reverse_proxy_backend_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_sflow_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_static_route_next_hop`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_a`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_pppoe_server_client_ipv6_pool_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv6_input_filter_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_local_route6_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_syslog_console_facility`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_pppoe_server_authentication_radius_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_snmp_trap_target`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bridge_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospfv3_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_group_domain_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_webproxy_url_filtering_squidguard_time_period_days`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_reverse_proxy_service_logging_facility`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_multicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_bmp_target`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_nhrp_tunnel_shortcut_target`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_neighbor_local_role`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_container_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_round_robin`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dhcp_server_shared_network_name_option_static_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_group_mac_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif_s_vif_c_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_multicast_route_next_hop`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospfv3_area`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_static_route_next_hop_bfd_multi_hop_source`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_container_name_environment`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_vif`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_login_user_authentication_public_keys`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_group_address_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv6_prerouting_raw_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_vif_s_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_multicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_network_emulator`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_zone`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_vif_s_vif_c_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_local_tiebreaker_lowest_backup_metric_index`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_multicast_distance_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_static_route6`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_table_route6_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_conntrack_timeout_custom_ipv6_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_vif_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_naptr_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_area_virtual_link_authentication_md5_key_id`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_bmp_target`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_conntrack_timeout_custom_ipv4_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_lowest_backup_metric_index`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_snmp_v3_trap_target`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_ike_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_extcommunity_list_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_router_advert_interface_nat64prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_sstp_authentication_radius_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_bridge_name`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_vif_s_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_l2vpn_evpn_vni`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_downstream_index`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_area_range`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dhcp_server_shared_network_name_subnet_option_static_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv4_prerouting_raw_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_nat_source_rule_load_balance_backend`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ntp_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_route_map_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_conntrack_ignore_ipv6_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_syslog_host`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_pptp_remote_access_client_ip_pool`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_peer_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_aaaa`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_monitoring_zabbix_agent_server_active`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_authentication_psk`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_distance_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_openconnect_authentication_local_users_username`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_login_user`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_container_name_sysctl_parameter`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_sstpc`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_pki_openssh`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dhcp_server_shared_network_name_subnet_static_mapping`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_domain_name_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_https_api_keys_id`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_webproxy_url_filtering_squidguard_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_vif_s_vif_c_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_openconnect_authentication_radius_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pppoe_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_reverse_proxy_service_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_geneve`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_vxlan`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_local_route_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_reverse_proxy_service`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_syslog_global_facility`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_as_path_list`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_prefix_list_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_unicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_labeled_unicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_unicast_distance_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_high_availability_vrrp_group_excluded_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_vif_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_macsec_security_static_peer`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_tunnel`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_reverse_proxy_backend_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_pki_openvpn_shared_secret`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_large_community_list_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospfv3_area_range`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_table_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ndp_proxy_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_vif_s`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_flowtable`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_reverse_proxy_backend_http_response_headers`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_prefix_list6`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_table_route6_next_hop_bfd_multi_hop_source`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_shaper_hfsc_class`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_vif_s_vif_c`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_area_virtual_link_authentication_md5_key_id`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_cake`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_fq_codel`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_group_ipv6_network_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_wan_interface_health_test`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_nat_source_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_listen_range`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_group_network_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bridge_member_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_route6_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_arp_interface_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_limiter_class`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv6_output_filter_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_table_route6`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif_s_vif_c_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ripng_distribute_list_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_neighbor_proxy_arp`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_vif_s_vif_c_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_multicast_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_nhrp_tunnel_map`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_neighbor`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_interface_authentication_md5_key_id`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_syslog_user`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv4_name`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_shaper`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ripng_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_txt`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ipoe_server_client_ip_pool`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_pki_certificate`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv4_forward_filter_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv4_output_raw_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_access_list`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_snmp_community`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_stunnel_client`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_static_host_mapping_host_name`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_syslog_file`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_input`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_arp_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ndp_proxy_interface_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_pppoe_server_client_ip_pool`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_macsec`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_unicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_wan_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_segment_routing_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bonding_vif`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_vif_s_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_nat_cgnat_pool_external_range`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_access_list6`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_domain`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_container_name_label`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv6_output_raw_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_access_list6_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_remote_access_connection_authentication_local_users_username`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bridge_vif`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_nhrp_tunnel`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_virtual_ethernet_vif_s_vif_c_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_static_route6_next_hop_bfd_multi_hop_source`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_vti`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_summary_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_group_ipv6_address_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_vif_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_random_detect`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ipoe_server_authentication_radius_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_pppoe_server_client_ipv6_pool_delegate`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_snmp_v3_user`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_event_handler_event_script_environment`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_ike_group_proposal`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_remote_access_radius_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_group_interface_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_snmp_v3_view_oid`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_rip_interface_authentication_md5`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_ipsec_site_to_site_peer_tunnel`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_labeled_unicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_shaper_class_match`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_pptp_remote_access_client_ipv6_pool_prefix`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_l2vpn_evpn_vni`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_isis_fast_reroute_lfa_remote_prefix_list`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_isis_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_priority_queue_class`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_sstp_client_ipv6_pool`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_neighbor`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_ethernet_vif_s`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_syslog_file_facility`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_system_syslog_user_facility`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_high_availability_vrrp_sync_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_loopback`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv6_forward_filter_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_load_balancing_reverse_proxy_backend`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bridge_vif_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_dynamic_name`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_macsec_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_bridge`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_pppoe_dhcpv6_options_pd`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_wireless_vif_s_vif_c`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain_records_cname`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_labeled_unicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_shaper_hfsc_class_match`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_limiter_class_match`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_group_port_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_interfaces_macsec_dhcpv6_options_pd_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_policy_round_robin_class`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_peer_group_local_as`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_ipv6_name`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_vpn_network`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_rip_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_route_next_hop`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_qos_traffic_match_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_firewall_group_dynamic_group_ipv6_address_group`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vpn_pptp_remote_access_authentication_radius_server`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_unicast_aggregate_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_summary_address`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_stunnel_client_psk`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_ospf_access_list`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_ospf_neighbor`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_pki_dh`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_policy_route_rule`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_pim_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_protocols_static_multicast_interface_route`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_service_ipoe_server_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single

* Modified Resource `vyos_vrf_name_protocols_static_route_interface`
	* Modified attribute `identifier.` nested attribute mode changed to single








## Previous changelogs
For older versions see changelog archive [directory](data/changelogs/) or changelog for [version 1](data/changelogs/CHANGELOG-v1.md)
