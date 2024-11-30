
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 12.0.202411270 (2024-11-29 08-37-33 UTC)](#release-120202411270-2024-11-29-08-37-33-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
      - [Bug fixes](#bug-fixes)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
      - [Features](#features)
        - [Resources](#resources-1)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 12.0.202411270 (2024-11-29 08-37-33 UTC)
### Project changes
#### Notes
* update to rolling release 2024-11-27T00:06:17Z
#### Bug fixes
* merged conntrack-sync resources
* merge dhcp-server subnet range

### Schema changes
#### BREAKING CHANGES

##### Resources
* **Removed Resource** `vyos_service_dhcp_server_shared_network_name_subnet_range`

* **Removed Resource** `vyos_service_conntrack_sync_interface`

* **Removed Resource** `vyos_service_conntrack_sync_failover_mechanism_vrrp`

* **Removed Resource** `vyos_service_dhcp_server_shared_network_name_subnet`





#### Features

##### Resources
* Modified Resource `vyos_service_dhcp_server_shared_network_name`
	* New attribute `subnet`

* Modified Resource `vyos_service_conntrack_sync`
	* New attribute `failover_mechanism`
	* New attribute `interface`








## Previous changelogs
For previous version see [changelog for v11](CHANGELOG-v11.md) or older archives [directory](data/changelogs/)
