
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 13.0.202411270 (2024-11-30 19-26-29 UTC)](#release-130202411270-2024-11-30-19-26-29-utc)
    - [Project changes](#project-changes)
      - [Bug fixes](#bug-fixes)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 13.0.202411270 (2024-11-30 19-26-29 UTC)
### Project changes
#### Bug fixes
* changelog generation nil dereference error

### Schema changes
#### BREAKING CHANGES

##### Resources
* Modified Resource `vyos_container_network`
	* **Removed attribute** `mtu`

* Modified Resource `vyos_qos_policy_cake`
	* Attribute `rtt`changed `description`
	* Modified attribute `flow_isolation`
		* changed to `nested` attribute
		* changed `description`
	* **Removed attribute** `flow_isolation_nat`

* Modified Resource `vyos_vpn_ipsec_authentication_psk`
	* **Removed attribute** `secret_type`

* Modified Resource `vyos_service_ipoe_server`
	* **Removed attribute** `lua_file`

* Modified Resource `vyos_service_mdns_repeater`
	* **Removed attribute** `cache_entries`

* Modified Resource `vyos_service_ipoe_server_interface`
	* **Removed attribute** `lua_username`








## Previous changelogs
For previous version see [changelog for v12](data/changelogs/CHANGELOG-v12.md) or older archives [directory](data/changelogs/)
