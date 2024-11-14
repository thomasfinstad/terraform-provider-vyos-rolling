
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 9.0.202411130 (2024-11-13 08-30-45 UTC)](#release-90202411130-2024-11-13-08-30-45-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
    - [Schema changes](#schema-changes)
      - [Notes](#notes-1)
        - [Resources](#resources)
  - [Release 9.0.202411071 (2024-11-10 08-27-14 UTC)](#release-90202411071-2024-11-10-08-27-14-utc)
    - [Project changes](#project-changes-1)
      - [Notes](#notes-2)
  - [Release 9.0.202411070 (2024-11-07 08-11-32 UTC)](#release-90202411070-2024-11-07-08-11-32-utc)
    - [Project changes](#project-changes-2)
      - [Notes](#notes-3)
      - [Bug fixes](#bug-fixes)
    - [Schema changes](#schema-changes-1)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources-1)
      - [Features](#features)
        - [Resources](#resources-2)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 9.0.202411130 (2024-11-13 08-30-45 UTC)
### Project changes
#### Notes
* update to rolling release 2024-11-13T00:05:53Z

### Schema changes
#### Notes

##### Resources
* Modified Resource `vyos_system_option`
	* Modified attributes under: `performance`
		* New attribute ``
		* Modified attribute `` changed description









## Release 9.0.202411071 (2024-11-10 08-27-14 UTC)
### Project changes
#### Notes
* update to rolling release 2024-11-07T00:05:48Z
* fix error in import doc


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
For previous version see [changelog for v8](CHANGELOG-v8.md) or older archives [directory](data/changelogs/)
