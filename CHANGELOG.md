
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 11.0.202411271 (2024-11-28 08-33-49 UTC)](#release-110202411271-2024-11-28-08-33-49-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
  - [Release 11.0.202411270 (2024-11-27 23-56-46 UTC)](#release-110202411270-2024-11-27-23-56-46-utc)
    - [Project changes](#project-changes-1)
      - [Notes](#notes-1)
      - [Bug fixes](#bug-fixes)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
      - [Features](#features)
        - [Resources](#resources-1)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 11.0.202411271 (2024-11-28 08-33-49 UTC)
### Project changes
#### Notes
* improved schema based changelog output


## Release 11.0.202411270 (2024-11-27 23-56-46 UTC)
### Project changes
#### Notes
* update to rolling release 2024-11-27T00:06:17Z
* improve build logging
* moved from rootless podman to privileged docker
#### Bug fixes
* remove redundant makefile dependency and debug prints
* nil pointer error during some change types

### Schema changes
#### BREAKING CHANGES

##### Resources
* Modified Resource `vyos_qos_policy_cake`
	* Modified attribute `flow_isolation`
		* type changed to `string`
		* changed `description`
	* changed `description`
	* New attribute `flow_isolation_nat`





#### Features

##### Resources
* Modified Resource `vyos_service_ipoe_server_interface`
	* New attribute `lua_username`

* Modified Resource `vyos_service_ipoe_server`
	* New attribute `lua_file`

* Modified Resource `vyos_vpn_ipsec_authentication_psk`
	* New attribute `secret_type`

* Modified Resource `vyos_service_mdns_repeater`
	* New attribute `cache_entries`








## Previous changelogs
For previous version see [changelog for v10](data/changelogs/CHANGELOG-v10.md) or older archives [directory](data/changelogs/)
