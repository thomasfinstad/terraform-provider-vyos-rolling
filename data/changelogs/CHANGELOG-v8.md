
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 8.0.202411060 (2024-11-06 15-38-46 UTC)](#release-80202411060-2024-11-06-15-38-46-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
      - [Bug fixes](#bug-fixes)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 8.0.202411060 (2024-11-06 15-38-46 UTC)
### Project changes
#### Notes
* update to rolling release 2024-11-06T06:01:10Z
#### Bug fixes
* remove child nodes from schema and resource model where they should not be

### Schema changes
#### BREAKING CHANGES

##### Resources
* Modified Resource `vyos_load_balancing_haproxy_service`
	* **Removed attribute** `logging`

* Modified Resource `vyos_service_ssh`
	* **Removed attribute** `access_control`

* Modified Resource `vyos_protocols_static_table_route6_next_hop`
	* **Removed attribute** `bfd.multi_hop`

* Modified Resource `vyos_service_https`
	* **Removed attribute** `api`

* Modified Resource `vyos_vpn_ipsec_remote_access_connection`
	* **Removed attribute** `authentication.local_users`

* Modified Resource `vyos_vrf_name_protocols_ospf_interface`
	* **Removed attribute** `authentication.md5`

* Modified Resource `vyos_protocols_rip`
	* **Removed attribute** `distribute_list`
	* **Removed attribute** `redistribute`

* Modified Resource `vyos_protocols_ospf`
	* **Removed attribute** `timers`
	* **Removed attribute** `default_information`
	* **Removed attribute** `max_metric`
	* **Removed attribute** `redistribute`

* Modified Resource `vyos_system`
	* **Removed attribute** `task_scheduler`
	* **Removed attribute** `static_host_mapping`
	* **Removed attribute** `sysctl`
	* **Removed attribute** `logs`

* Modified Resource `vyos_vpn_openconnect`
	* **Removed attribute** `accounting`

* Modified Resource `vyos_vrf_name_protocols_static_route_next_hop`
	* **Removed attribute** `bfd.multi_hop`

* Modified Resource `vyos_service_snmp`
	* **Removed attribute** `script_extensions`

* Modified Resource `vyos_vpn_openconnect_authentication`
	* **Removed attribute** `local_users`

* Modified Resource `vyos_firewall_global_options`
	* **Removed attribute** `state_policy`

* Modified Resource `vyos_protocols_static_route_next_hop`
	* **Removed attribute** `bfd.multi_hop`

* Modified Resource `vyos_vrf_name`
	* Modified attributes under: `protocols`
		* **Removed attribute** `static`
		* Modified attributes under: `bgp.address_family`
			* **Removed attribute** `ipv6_vpn`
			* **Removed attribute** `ipv4_vpn`
			* **Removed attribute** `ipv6_labeled_unicast`
		* Modified attributes under: `isis.fast_reroute.lfa`
			* **Removed attribute** `local.tiebreaker`
			* **Removed attribute** `remote`

* Modified Resource `vyos_interfaces_bridge`
	* **Removed attribute** `member`

* Modified Resource `vyos_service_ids_ddos_protection`
	* **Removed attribute** `threshold`

* Modified Resource `vyos_service_dns_forwarding_authoritative_domain`
	* **Removed attribute** `records`

* Modified Resource `vyos_vrf_name_protocols_ospf_area_virtual_link`
	* **Removed attribute** `authentication.md5`

* Modified Resource `vyos_protocols_bgp_address_family_l2vpn_evpn`
	* **Removed attribute** `advertise`

* Modified Resource `vyos_system_conntrack`
	* **Removed attribute** `ignore`
	* **Removed attribute** `timeout`

* Modified Resource `vyos_protocols_mpls_ldp`
	* **Removed attribute** `export`
	* **Removed attribute** `import`
	* **Removed attribute** `targeted_neighbor`
	* **Removed attribute** `allocation`

* Modified Resource `vyos_container_name`
	* **Removed attribute** `sysctl`

* Modified Resource `vyos_service_pppoe_server_authentication`
	* **Removed attribute** `local_users`

* Modified Resource `vyos_protocols_bgp_parameters`
	* **Removed attribute** `distance`

* Modified Resource `vyos_protocols_static_route6_next_hop`
	* **Removed attribute** `bfd.multi_hop`

* Modified Resource `vyos_vpn_ipsec`
	* **Removed attribute** `authentication`
	* **Removed attribute** `remote_access`
	* **Removed attribute** `site_to_site`

* Modified Resource `vyos_load_balancing_haproxy_global_parameters`
	* **Removed attribute** `logging`

* Modified Resource `vyos_system_conntrack_log`
	* **Removed attribute** `event`

* Modified Resource `vyos_protocols_static`
	* **Removed attribute** `arp`
	* **Removed attribute** `multicast`
	* **Removed attribute** `neighbor_proxy`

* Modified Resource `vyos_service_ntp_ptp`
	* **Removed attribute** `timestamp`

* Modified Resource `vyos_service_dhcpv6_server_shared_network_name_subnet`
	* **Removed attribute** `prefix_delegation`

* Modified Resource `vyos_system_syslog`
	* **Removed attribute** `console`

* Modified Resource `vyos_protocols_isis`
	* **Removed attribute** `fast_reroute`
	* **Removed attribute** `redistribute`
	* **Removed attribute** `default_information`

* Modified Resource `vyos_vrf_name_protocols_static_route6_next_hop`
	* **Removed attribute** `bfd.multi_hop`

* Modified Resource `vyos_service_conntrack_sync`
	* **Removed attribute** `failover_mechanism`

* Modified Resource `vyos_protocols_static_table_route_next_hop`
	* **Removed attribute** `bfd.multi_hop`

* Modified Resource `vyos_protocols_ospf_area_virtual_link`
	* **Removed attribute** `authentication.md5`

* Modified Resource `vyos_protocols_ripng`
	* **Removed attribute** `distribute_list`
	* **Removed attribute** `redistribute`

* Modified Resource `vyos_protocols_ospf_interface`
	* **Removed attribute** `authentication.md5`

* Modified Resource `vyos_load_balancing_haproxy_backend`
	* **Removed attribute** `logging`

* Modified Resource `vyos_nat_cgnat`
	* **Removed attribute** `pool`

* Modified Resource `vyos_vpn_sstp_authentication`
	* **Removed attribute** `local_users`

* Modified Resource `vyos_vpn_pptp_remote_access_authentication`
	* **Removed attribute** `local_users`

* Modified Resource `vyos_system_ip`
	* **Removed attribute** `tcp`

* Modified Resource `vyos_vpn_l2tp_remote_access_authentication`
	* **Removed attribute** `local_users`

* Modified Resource `vyos_protocols_bgp`
	* **Removed attribute** `address_family`
	* **Removed attribute** `sid`

* Modified Resource `vyos_service_suricata`
	* **Removed attribute** `log`

* Modified Resource `vyos_protocols_pim`
	* **Removed attribute** `spt_switchover`








## Previous changelogs
For previous version see [changelog for v7](CHANGELOG-v7.md) or older archives [directory](data/changelogs/)
