
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 9.0.202411070 (2024-11-07 08-11-32 UTC)](#release-90202411070-2024-11-07-08-11-32-utc)
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


## Release 9.0.202411070 (2024-11-07 08-11-32 UTC)
### Project changes
#### Notes
* update to rolling release 2024-11-07T00:05:48Z
#### Bug fixes
* removed dhcp-server subnet merge

### Schema changes
#### BREAKING CHANGES

##### Resources
* Modified Resource `vyos_service_dhcp_server_shared_network_name`
	* **Removed attribute** `subnet`





#### Features

##### Resources
* New Resource `vyos_service_dhcp_server_shared_network_name_subnet`








## Previous changelogs
For previous version see [changelog for v8](data/changelogs/CHANGELOG-v8.md) or older archives [directory](data/changelogs/)
