
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 20.0.202502190 (2025-02-19 08-31-26 UTC)](#release-200202502190-2025-02-19-08-31-26-utc)
    - [Project changes](#project-changes)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
      - [Features](#features)
        - [Resources](#resources-1)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 20.0.202502190 (2025-02-19 08-31-26 UTC)
### Project changes

### Schema changes
#### BREAKING CHANGES

##### Resources
* **Removed Resource** `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute`

* Modified Resource `vyos_vrf_name`
	* Modified attribute `protocols`
		* Modified attribute `protocols.bgp.bgp`
			* Modified attribute `protocols.bgp.bgp.address_family.address_family`
				* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv4_unicast.ipv4_unicast`
					* Attribute `redistribute`**Removed attribute** `protocols.bgp.bgp.address_family.address_family.ipv4_unicast.ipv4_unicast.redistribute.redistribute.table`
				* Modified attribute `protocols.bgp.bgp.address_family.address_family.ipv6_unicast.ipv6_unicast`
					* Attribute `redistribute`**Removed attribute** `protocols.bgp.bgp.address_family.address_family.ipv6_unicast.ipv6_unicast.redistribute.redistribute.table`

* **Removed Resource** `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute`





#### Features

##### Resources
* New Resource `vyos_vrf_name_protocols_bgp_address_family_ipv6_unicast_redistribute_table`

* New Resource `vyos_protocols_bgp_address_family_ipv4_unicast_redistribute_table`

* New Resource `vyos_protocols_bgp_address_family_ipv6_unicast_redistribute_table`

* New Resource `vyos_vrf_name_protocols_bgp_address_family_ipv4_unicast_redistribute_table`








## Previous changelogs
For previous version see [changelog for v19](data/changelogs/CHANGELOG-v19.md) or older archives [directory](data/changelogs/)
