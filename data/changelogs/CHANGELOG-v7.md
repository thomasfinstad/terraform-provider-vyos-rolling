
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 7.0.202411060 (2024-11-06 08-31-38 UTC)](#release-70202411060-2024-11-06-08-31-38-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
  - [Release 7.0.202410300 (2024-11-05 17-36-12 UTC)](#release-70202410300-2024-11-05-17-36-12-utc)
    - [Project changes](#project-changes-1)
      - [Notes](#notes-1)
      - [Features](#features)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
      - [Features](#features-1)
        - [Resources](#resources-1)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 7.0.202411060 (2024-11-06 08-31-38 UTC)
### Project changes
#### Notes
* update to rolling release 2024-11-06T06:01:10Z


## Release 7.0.202410300 (2024-11-05 17-36-12 UTC)
### Project changes
#### Notes
* update to rolling release 2024-10-30T00:06:09Z
#### Features
* allow named resources to be merged under global resources

### Schema changes
#### BREAKING CHANGES

##### Resources
* **Removed Resource** `vyos_service_dhcp_server_shared_network_name_subnet`





#### Features

##### Resources
* Modified Resource `vyos_vpn_openconnect`
	* New attribute `accounting`

* Modified Resource `vyos_protocols_bgp`
	* New attribute `address_family`
	* New attribute `sid`

* Modified Resource `vyos_vpn_l2tp_remote_access_authentication`
	* New attribute `local_users`

* Modified Resource `vyos_service_conntrack_sync`
	* New attribute `failover_mechanism`

* Modified Resource `vyos_load_balancing_haproxy_global_parameters`
	* New attribute `logging`

* Modified Resource `vyos_vpn_pptp_remote_access_authentication`
	* New attribute `local_users`

* Modified Resource `vyos_firewall_global_options`
	* New attribute `state_policy`

* Modified Resource `vyos_service_pppoe_server_authentication`
	* New attribute `local_users`

* Modified Resource `vyos_protocols_isis`
	* New attribute `default_information`
	* New attribute `fast_reroute`
	* New attribute `redistribute`

* Modified Resource `vyos_container_name`
	* New attribute `network`

* Modified Resource `vyos_service_ntp`
	* New attribute `server`
	* New attribute `allow_client`

* Modified Resource `vyos_nat_cgnat`
	* New attribute `pool`

* Modified Resource `vyos_protocols_rip`
	* New attribute `redistribute`
	* New attribute `distribute_list`

* Modified Resource `vyos_protocols_pim`
	* New attribute `spt_switchover`

* Modified Resource `vyos_interfaces_openvpn`
	* New attribute `server.client`
	* New attribute `local_address`

* Modified Resource `vyos_protocols_ospf`
	* New attribute `default_information`
	* New attribute `max_metric`
	* New attribute `redistribute`
	* New attribute `timers`

* Modified Resource `vyos_system_conntrack_log`
	* New attribute `event`

* Modified Resource `vyos_service_suricata`
	* New attribute `log`

* Modified Resource `vyos_system_ip`
	* New attribute `tcp`

* Modified Resource `vyos_protocols_ripng`
	* New attribute `distribute_list`
	* New attribute `redistribute`

* Modified Resource `vyos_service_ssh`
	* New attribute `access_control`

* Modified Resource `vyos_system`
	* New attribute `sysctl`
	* New attribute `task_scheduler`
	* New attribute `logs`
	* New attribute `static_host_mapping`

* Modified Resource `vyos_vpn_openconnect_authentication`
	* New attribute `local_users`

* Modified Resource `vyos_protocols_mpls_ldp`
	* New attribute `allocation`
	* New attribute `export`
	* New attribute `import`
	* New attribute `targeted_neighbor`

* Modified Resource `vyos_vpn_ipsec`
	* New attribute `remote_access`
	* New attribute `site_to_site`
	* New attribute `authentication`

* Modified Resource `vyos_service_https`
	* New attribute `api`

* Modified Resource `vyos_system_conntrack`
	* New attribute `ignore`
	* New attribute `timeout`

* Modified Resource `vyos_protocols_bgp_address_family_l2vpn_evpn`
	* New attribute `advertise`

* Modified Resource `vyos_vpn_sstp_authentication`
	* New attribute `local_users`

* Modified Resource `vyos_service_snmp`
	* New attribute `script_extensions`

* Modified Resource `vyos_system_syslog`
	* New attribute `console`

* Modified Resource `vyos_service_ids_ddos_protection`
	* New attribute `threshold`

* Modified Resource `vyos_protocols_static`
	* New attribute `multicast`
	* New attribute `neighbor_proxy`
	* New attribute `arp`

* Modified Resource `vyos_protocols_bgp_parameters`
	* New attribute `distance`

* Modified Resource `vyos_service_ntp_ptp`
	* New attribute `timestamp`

* Modified Resource `vyos_service_dhcp_server_shared_network_name`
	* New attribute `subnet`








## Previous changelogs
For previous version see [changelog for v6](CHANGELOG-v6.md) or older archives [directory](data/changelogs/)
