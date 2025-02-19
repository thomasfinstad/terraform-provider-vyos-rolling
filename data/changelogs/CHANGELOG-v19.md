
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 19.1.202502130 (2025-02-14 08-30-19 UTC)](#release-191202502130-2025-02-14-08-30-19-utc)
    - [Project changes](#project-changes)
    - [Schema changes](#schema-changes)
      - [Notes](#notes)
        - [Resources](#resources)
      - [Features](#features)
        - [Resources](#resources-1)
  - [Release 19.0.202502110 (2025-02-11 08-30-27 UTC)](#release-190202502110-2025-02-11-08-30-27-utc)
    - [Project changes](#project-changes-1)
    - [Schema changes](#schema-changes-1)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources-2)
      - [Features](#features-1)
        - [Resources](#resources-3)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 19.1.202502130 (2025-02-14 08-30-19 UTC)
### Project changes

### Schema changes
#### Notes

##### Resources
* Modified Resource `vyos_protocols_static_table_route6_next_hop`
	* Modified attribute `identifier`
		* Attribute `table`changed `description`

* Modified Resource `vyos_protocols_static_table`
	* Modified attribute `identifier`
		* Attribute `table`changed `description`

* Modified Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute`
	* Modified attribute `table`
		* type changed to `list of number`
		* changed `description`

* Modified Resource `vyos_protocols_static_table_route_interface`
	* Modified attribute `identifier`
		* Attribute `table`changed `description`

* Modified Resource `vyos_protocols_static_table_route6_interface`
	* Modified attribute `identifier`
		* Attribute `table`changed `description`

* Modified Resource `vyos_protocols_static_table_route6`
	* Modified attribute `identifier`
		* Attribute `table`changed `description`

* Modified Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute`
	* Modified attribute `table`
		* type changed to `list of number`
		* changed `description`

* Modified Resource `vyos_protocols_static_table_route_next_hop`
	* Modified attribute `identifier`
		* Attribute `table`changed `description`

* Modified Resource `vyos_protocols_static_table_route`
	* Modified attribute `identifier`
		* Attribute `table`changed `description`





#### Features

##### Resources
* Modified Resource `vyos_vrf_name`
	* Modified attribute `protocols`
		* Modified attribute `protocols.bgp.bgp`
			* Modified attribute `protocols.bgp.bgp.address_family.address_family`
				* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv4_unicast.ipv4_unicast`
					* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv4_unicast.ipv4_unicast.redistribute.redistribute`
						* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv4_unicast.ipv4_unicast.redistribute.redistribute.table.table`
							* type changed to `list of number`
							* changed `description`
				* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv6_unicast.ipv6_unicast`
					* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv6_unicast.ipv6_unicast.redistribute.redistribute`
						* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv6_unicast.ipv6_unicast.redistribute.redistribute.table.table`
							* type changed to `list of number`
							* changed `description`
						* New attribute `protocols.bgp.bgp.address_family.address_family.ipv6_unicast.ipv6_unicast.redistribute.redistribute.isis.isis`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_isis`









## Release 19.0.202502110 (2025-02-11 08-30-27 UTC)
### Project changes

### Schema changes
#### BREAKING CHANGES

##### Resources
* **Removed Resource** `vyos_system_syslog_global_facility`

* **Removed Resource** `vyos_system_syslog_global_marker`

* **Removed Resource** `vyos_system_syslog_global`

* **Removed Resource** `vyos_system_syslog_user`

* **Removed Resource** `vyos_system_syslog_user_facility`

* **Removed Resource** `vyos_system_syslog_file_facility`

* Modified Resource `vyos_system_syslog`
	* New attribute `preserve_fqdn`
	* **Removed attribute** `vrf`

* **Removed Resource** `vyos_system_syslog_host_facility`

* **Removed Resource** `vyos_system_syslog_host`

* **Removed Resource** `vyos_system_syslog_file`





#### Features

##### Resources
* New Resource `vyos_system_syslog_remote_facility`

* New Resource `vyos_system_syslog_local_facility`

* New Resource `vyos_system_syslog_remote`

* New Resource `vyos_system_syslog_marker`








## Previous changelogs
For previous version see [changelog for v18](CHANGELOG-v18.md) or older archives [directory](data/changelogs/)
