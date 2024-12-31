
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 14.1.202412140 (2024-12-14 08-30-30 UTC)](#release-141202412140-2024-12-14-08-30-30-utc)
    - [Project changes](#project-changes)
    - [Schema changes](#schema-changes)
      - [Notes](#notes)
        - [Resources](#resources)
  - [Release 14.1.202412100 (2024-12-10 08-33-29 UTC)](#release-141202412100-2024-12-10-08-33-29-utc)
    - [Project changes](#project-changes-1)
    - [Schema changes](#schema-changes-1)
      - [Features](#features)
        - [Resources](#resources-1)
  - [Release 14.0.202412030 (2024-12-04 08-33-51 UTC)](#release-140202412030-2024-12-04-08-33-51-utc)
    - [Project changes](#project-changes-2)
    - [Schema changes](#schema-changes-2)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources-2)
      - [Features](#features-1)
        - [Resources](#resources-3)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 14.1.202412140 (2024-12-14 08-30-30 UTC)
### Project changes

### Schema changes
#### Notes

##### Resources
* Modified Resource `vyos_interfaces_ethernet`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_ethernet_vif_s_vif_c`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_wireless_vif`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_virtual_ethernet_vif_s`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_sstpc`
	* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif_s_vif_c`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_bonding_vif`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_wireless_vif_s_vif_c`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_virtual_ethernet_vif_s_vif_c`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif_s`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_virtual_ethernet_vif`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_macsec`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_pppoe`
	* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_bridge_vif`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_ethernet_vif_s`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_wireless_vif_s`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_bridge`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_ethernet_vif`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_bonding_vif_s`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_pseudo_ethernet_vif`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_wwan`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_wireless`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_pki_certificate`
	* Modified attribute `acme`
		* Attribute `listen_address`changed `description`

* Modified Resource `vyos_interfaces_pseudo_ethernet`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_bonding`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_virtual_ethernet`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`

* Modified Resource `vyos_interfaces_bonding_vif_s_vif_c`
	* Modified attribute `dhcp_options`
		* Attribute `default_route_distance`changed `description`









## Release 14.1.202412100 (2024-12-10 08-33-29 UTC)
### Project changes

### Schema changes
#### Features

##### Resources
* Modified Resource `vyos_container_name`
	* New attribute `name_server`









## Release 14.0.202412030 (2024-12-04 08-33-51 UTC)
### Project changes

### Schema changes
#### BREAKING CHANGES

##### Resources
* Modified Resource `vyos_qos_policy_cake`
	* Attribute `rtt`changed `description`
	* Modified attribute `flow_isolation`
		* type changed to `string`
		* changed `description`
	* New attribute `flow_isolation_nat`

* **Removed Resource** `vyos_service_ntp_ptp_timestamp_interface`





#### Features

##### Resources
* Modified Resource `vyos_service_ipoe_server_interface`
	* New attribute `lua_username`
	* New attribute `start_session`

* Modified Resource `vyos_vpn_ipsec_authentication_psk`
	* New attribute `secret_type`

* Modified Resource `vyos_service_ipoe_server`
	* New attribute `lua_file`

* Modified Resource `vyos_pki_ca`
	* New attribute `system_install`

* Modified Resource `vyos_service_mdns_repeater`
	* New attribute `cache_entries`

* Modified Resource `vyos_container_network`
	* New attribute `mtu`

* New Resource `vyos_service_ntp_timestamp_interface`








## Previous changelogs
For previous version see [changelog for v13](CHANGELOG-v13.md) or older archives [directory](data/changelogs/)
