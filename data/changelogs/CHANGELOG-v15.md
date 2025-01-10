
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 15.2.202501050 (2025-01-06 08-32-01 UTC)](#release-152202501050-2025-01-06-08-32-01-utc)
    - [Project changes](#project-changes)
    - [Schema changes](#schema-changes)
      - [Features](#features)
        - [Resources](#resources)
  - [Release 15.1.202501030 (2025-01-03 08-30-14 UTC)](#release-151202501030-2025-01-03-08-30-14-utc)
    - [Project changes](#project-changes-1)
    - [Schema changes](#schema-changes-1)
      - [Features](#features-1)
        - [Resources](#resources-1)
  - [Release 15.0.202412310 (2024-12-31 08-30-06 UTC)](#release-150202412310-2024-12-31-08-30-06-utc)
    - [Project changes](#project-changes-2)
    - [Schema changes](#schema-changes-2)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources-2)
      - [Notes](#notes)
        - [Resources](#resources-3)
      - [Features](#features-2)
        - [Resources](#resources-4)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 15.2.202501050 (2025-01-06 08-32-01 UTC)
### Project changes

### Schema changes
#### Features

##### Resources
* New Resource `vyos_service_monitoring_prometheus_node_exporter_collectors`









## Release 15.1.202501030 (2025-01-03 08-30-14 UTC)
### Project changes

### Schema changes
#### Features

##### Resources
* Modified Resource `vyos_protocols_segment_routing_srv6_locator`
	* New attribute `format`

* New Resource `vyos_service_monitoring_prometheus_blackbox_exporter_modules_icmp_name`

* New Resource `vyos_service_monitoring_prometheus_blackbox_exporter`

* New Resource `vyos_service_monitoring_prometheus_blackbox_exporter_modules_dns_name`









## Release 15.0.202412310 (2024-12-31 08-30-06 UTC)
### Project changes

### Schema changes
#### BREAKING CHANGES

##### Resources
* **Removed Resource** `vyos_protocols_static_multicast_interface_route_next_hop_interface`

* **Removed Resource** `vyos_vrf_name_protocols_static_route_next_hop_bfd_multi_hop_source`

* **Removed Resource** `vyos_vrf_name_protocols_static_route6_next_hop_bfd_multi_hop_source`

* **Removed Resource** `vyos_protocols_static_multicast_interface_route`

* **Removed Resource** `vyos_protocols_static_multicast_route_next_hop`

* **Removed Resource** `vyos_protocols_static_table_route6_next_hop_bfd_multi_hop_source`

* **Removed Resource** `vyos_service_monitoring_frr_exporter`

* **Removed Resource** `vyos_protocols_static_table_route_next_hop_bfd_multi_hop_source`

* **Removed Resource** `vyos_service_monitoring_node_exporter`

* **Removed Resource** `vyos_protocols_static_route_next_hop_bfd_multi_hop_source`

* **Removed Resource** `vyos_protocols_static_multicast_route`

* **Removed Resource** `vyos_protocols_static_route6_next_hop_bfd_multi_hop_source`





#### Notes

##### Resources
* Modified Resource `vyos_protocols_mpls_ldp`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_high_availability_vrrp_group`
	* Modified attribute `track`
		* Attribute `interface`changed `description`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_flowspec_local_install`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_service_stunnel_server_psk`
	* Attribute `secret`changed `description`

* Modified Resource `vyos_qos_policy_limiter_class_match`
	* Attribute `interface`changed `description`
	* Modified attribute `ether`
		* Attribute `protocol`changed `description`

* Modified Resource `vyos_service_dhcp_relay`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_service_ipoe_server_authentication_radius`
	* Attribute `source_address`changed `description`

* Modified Resource `vyos_high_availability_vrrp_group_excluded_address`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_qos_policy_round_robin_class_match`
	* Modified attribute `ether`
		* Attribute `protocol`changed `description`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_flowspec_local_install`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_service_mdns_repeater`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_vrf_name`
	* Modified attribute `protocols`
		* Modified attribute `protocols.bgp.bgp`
			* Modified attribute `protocols.bgp.bgp.address_family.address_family`
				* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv6_flowspec.ipv6_flowspec`
					* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv6_flowspec.ipv6_flowspec.local_install.local_install`
						* Attribute `interface`changed `description`
				* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv4_flowspec.ipv4_flowspec`
					* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv4_flowspec.ipv4_flowspec.local_install.local_install`
						* Attribute `interface`changed `description`

* Modified Resource `vyos_protocols_pim_interface`
	* Attribute `source_address`changed `description`

* Modified Resource `vyos_protocols_static_table_route`
	* Attribute `dhcp_interface`type changed to `list of string`

* Modified Resource `vyos_qos_policy_priority_queue_class_match`
	* Modified attribute `ether`
		* Attribute `protocol`changed `description`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_vpn_l2tp_remote_access_authentication_radius`
	* Attribute `source_address`changed `description`

* Modified Resource `vyos_vpn_ipsec_options`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_vrf_name_protocols_static_route`
	* Attribute `dhcp_interface`type changed to `list of string`

* Modified Resource `vyos_service_broadcast_relay_id`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_service_pppoe_server_authentication_radius`
	* Attribute `source_address`changed `description`

* Modified Resource `vyos_vpn_ipsec`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_system_sflow`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_service_suricata`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_system_flow_accounting`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_service_dns_dynamic_name`
	* Modified attribute `address`
		* Attribute `interface`changed `description`

* Modified Resource `vyos_interfaces_wireless`
	* Modified attribute `security`
		* Modified attribute `security.wpa.wpa`
			* Modified attribute `security.wpa.wpa.radius.radius`
				* Attribute `source_address`changed `description`

* Modified Resource `vyos_vpn_pptp_remote_access_authentication_radius`
	* Attribute `source_address`changed `description`

* Modified Resource `vyos_service_ntp`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_service_stunnel_client_psk`
	* Attribute `secret`changed `description`

* Modified Resource `vyos_vpn_ipsec_remote_access_dhcp`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_protocols_static_neighbor_proxy_and`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_policy_route_map_rule`
	* Modified attribute `match`
		* Attribute `interface`changed `description`

* Modified Resource `vyos_qos_policy_shaper_class_match`
	* Attribute `interface`changed `description`
	* Modified attribute `ether`
		* Attribute `protocol`changed `description`

* Modified Resource `vyos_system_login_tacacs`
	* Attribute `source_address`changed `description`

* Modified Resource `vyos_vpn_sstp_authentication_radius`
	* Attribute `source_address`changed `description`

* Modified Resource `vyos_protocols_failover_route_next_hop`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_protocols_static_route`
	* Attribute `dhcp_interface`type changed to `list of string`

* Modified Resource `vyos_service_dhcpv6_server_shared_network_name_subnet`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_service_dhcp_server_high_availability`
	* Attribute `source_address`changed `description`

* Modified Resource `vyos_vpn_openconnect_authentication_radius`
	* Attribute `source_address`changed `description`

* Modified Resource `vyos_protocols_mpls`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_service_ndp_proxy_interface_prefix`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_policy_route`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_protocols_static_neighbor_proxy_arp`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_high_availability_vrrp_group_address`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_policy_route6`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_qos_policy_shaper_hfsc_class_match`
	* Attribute `interface`changed `description`
	* Modified attribute `ether`
		* Attribute `protocol`changed `description`

* Modified Resource `vyos_firewall_global_options_apply_to_bridged_traffic`
	* Attribute `invalid_connections`changed `description`

* Modified Resource `vyos_service_dhcpv6_server_shared_network_name`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_protocols_bfd_peer`
	* Modified attribute `source`
		* Attribute `interface`changed `description`

* Modified Resource `vyos_vpn_ipsec_remote_access_radius`
	* Attribute `source_address`changed `description`





#### Features

##### Resources
* Modified Resource `vyos_protocols_ospf_area_virtual_link`
	* New attribute `retransmit_window`

* Modified Resource `vyos_vrf_name_protocols_static_route_next_hop`
	* Attribute `interface`changed `description`
	* Attribute `bfd`New attribute `bfd.multi_hop`

* Modified Resource `vyos_interfaces_ethernet`
	* New attribute `switchdev`

* Modified Resource `vyos_service_pppoe_server_interface`
	* New attribute `combined`

* Modified Resource `vyos_service_ipoe_server_authentication_interface_mac`
	* New attribute `static_ip`

* Modified Resource `vyos_vrf_name_protocols_ospf_interface`
	* New attribute `retransmit_window`

* Modified Resource `vyos_protocols_static_table_route_next_hop`
	* Attribute `bfd`New attribute `bfd.multi_hop`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_protocols_static_table_route6_next_hop`
	* Attribute `bfd`New attribute `bfd.multi_hop`
	* Attribute `interface`changed `description`

* Modified Resource `vyos_protocols_static_route_next_hop`
	* Attribute `interface`changed `description`
	* Attribute `bfd`New attribute `bfd.multi_hop`

* Modified Resource `vyos_protocols_ospf_interface`
	* New attribute `retransmit_window`

* Modified Resource `vyos_vrf_name_protocols_static_route6_next_hop`
	* Attribute `interface`changed `description`
	* Attribute `bfd`New attribute `bfd.multi_hop`

* Modified Resource `vyos_vrf_name_protocols_ospf_area_virtual_link`
	* New attribute `retransmit_window`

* Modified Resource `vyos_protocols_static_route6_next_hop`
	* Attribute `interface`changed `description`
	* Attribute `bfd`New attribute `bfd.multi_hop`

* New Resource `vyos_protocols_static_mroute`

* New Resource `vyos_service_monitoring_zabbix_agent_authentication`

* New Resource `vyos_service_monitoring_zabbix_agent_authentication_psk`

* New Resource `vyos_service_ssh_trusted_user_ca_key`

* New Resource `vyos_service_monitoring_prometheus_frr_exporter`

* New Resource `vyos_protocols_static_mroute_next_hop`

* New Resource `vyos_protocols_static_mroute_interface`

* New Resource `vyos_service_monitoring_prometheus_node_exporter`








## Previous changelogs
For previous version see [changelog for v14](CHANGELOG-v14.md) or older archives [directory](data/changelogs/)
