
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 17.0.202501210 (2025-01-21 11-37-09 UTC)](#release-170202501210-2025-01-21-11-37-09-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
  - [Release 17.0.202501140 (2025-01-15 08-30-06 UTC)](#release-170202501140-2025-01-15-08-30-06-utc)
    - [Project changes](#project-changes-1)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
      - [Features](#features)
        - [Resources](#resources-1)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 17.0.202501210 (2025-01-21 11-37-09 UTC)
### Project changes
#### Notes
* change TagNode identifier into descriptive struct


## Release 17.0.202501140 (2025-01-15 08-30-06 UTC)
### Project changes

### Schema changes
#### BREAKING CHANGES

##### Resources
* **Removed Resource** `vyos_protocols_nhrp_tunnel_shortcut_target`

* **Removed Resource** `vyos_protocols_nhrp_tunnel_dynamic_map`

* Modified Resource `vyos_protocols_nhrp_tunnel`
	* Modified attribute `multicast`
		* type changed to `list of string`
		* changed `description`
	* New attribute `authentication`
	* New attribute `holdtime`
	* New attribute `registration_no_unique`
	* New attribute `mtu`
	* New attribute `network_id`
	* **Removed attribute** `shortcut_destination`
	* **Removed attribute** `cisco_authentication`
	* **Removed attribute** `non_caching`
	* **Removed attribute** `holding_time`

* **Removed Resource** `vyos_protocols_nhrp_tunnel_map`





#### Features

##### Resources
* New Resource `vyos_protocols_nhrp_tunnel_nhs_tunnel_ip`

* New Resource `vyos_protocols_nhrp_tunnel_map_tunnel_ip`








## Previous changelogs
For previous version see [changelog for v16](CHANGELOG-v16.md) or older archives [directory](data/changelogs/)
