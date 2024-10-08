
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 3.11.202410081 (????)](#release-311202410081-)
  - [Release 3.0.202410080 (2024-10-08 16-52-50 UTC)](#release-30202410080-2024-10-08-16-52-50-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
      - [Bug fixes](#bug-fixes)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
      - [Notes](#notes-1)
        - [Resources](#resources-1)
      - [Features](#features)
        - [Resources](#resources-2)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->

## Release 3.11.202410081 (????)
bugged release

## Release 3.0.202410080 (2024-10-08 16-52-50 UTC)
### Project changes
#### Notes
* reduce wasted ci resources
* clean up CI instance disk space before building release
* update to rolling release 2024-10-08T12:08:53Z
* fix devcontainer tooling download
#### Bug fixes
* minor versioning reset bug and changelog typeos
* multivalue elements dereferencing during vyos api data generation

### Schema changes
#### BREAKING CHANGES

##### Resources
* **Removed Resource** `vyos_service_https_api`

* **Removed Resource** `vyos_service_https_api_cors`





#### Notes

##### Resources
* Modified Resource `vyos_system_option`
	* Modified attribute `keyboard_layout` changed description

* Modified Resource `vyos_policy_community_list_rule`
	* Modified attribute `regex` changed description





#### Features

##### Resources
* Modified Resource `vyos_firewall_ipv6_output_filter_rule`
	* New attribute `set`

* Modified Resource `vyos_firewall_bridge_output_filter_rule`
	* New attribute `set`

* Modified Resource `vyos_policy_local_route6_rule`
	* New attribute `set.vrf`

* Modified Resource `vyos_policy_local_route_rule`
	* New attribute `set.vrf`

* Modified Resource `vyos_firewall_ipv4_forward_filter_rule`
	* New attribute `set`

* Modified Resource `vyos_nat_source_rule`
	* New attribute `destination.fqdn`
	* New attribute `source.fqdn`

* Modified Resource `vyos_firewall_ipv6_output_raw_rule`
	* New attribute `set`

* Modified Resource `vyos_firewall_ipv4_name_rule`
	* New attribute `set`

* Modified Resource `vyos_firewall_bridge_prerouting_filter_rule`
	* New attribute `set`

* Modified Resource `vyos_firewall_ipv4_output_filter_rule`
	* New attribute `set`

* Modified Resource `vyos_firewall_ipv6_name_rule`
	* New attribute `set`

* Modified Resource `vyos_firewall_bridge_forward_filter_rule`
	* New attribute `set`

* Modified Resource `vyos_firewall_ipv4_prerouting_raw_rule`
	* New attribute `set`

* Modified Resource `vyos_firewall_ipv6_forward_filter_rule`
	* New attribute `set`

* Modified Resource `vyos_firewall_ipv6_prerouting_raw_rule`
	* New attribute `set`

* Modified Resource `vyos_nat_destination_rule`
	* New attribute `destination.fqdn`
	* New attribute `source.fqdn`

* Modified Resource `vyos_firewall_ipv4_output_raw_rule`
	* New attribute `set`

* New Resource `vyos_service_https_api_rest`

* New Resource `vyos_service_https_api_graphql_cors`

* New Resource `vyos_service_monitoring_node_exporter`








## Previous changelogs
For previous version see [changelog for v2](data/changelogs/CHANGELOG-v2.md) or older archives [directory](data/changelogs/)
