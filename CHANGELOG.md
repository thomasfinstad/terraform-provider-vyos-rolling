
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 18.1.202501250 (2025-01-25 08-27-53 UTC)](#release-181202501250-2025-01-25-08-27-53-utc)
    - [Project changes](#project-changes)
    - [Schema changes](#schema-changes)
      - [Notes](#notes)
        - [Resources](#resources)
      - [Features](#features)
        - [Resources](#resources-1)
  - [Release 18.0.202501220 (2025-01-22 08-30-54 UTC)](#release-180202501220-2025-01-22-08-30-54-utc)
    - [Project changes](#project-changes-1)
    - [Schema changes](#schema-changes-1)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources-2)
      - [Features](#features-1)
        - [Resources](#resources-3)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 18.1.202501250 (2025-01-25 08-27-53 UTC)
### Project changes

### Schema changes
#### Notes

##### Resources
* Modified Resource `vyos_system_conntrack_log`
	* Attribute `log_level`changed `description`
	* Attribute `queue_size`changed `description`





#### Features

##### Resources
* Modified Resource `vyos_interfaces_wireguard_peer`
	* New attribute `host_name`

* Modified Resource `vyos_interfaces_wireguard`
	* New attribute `max_dns_retry`

* New Resource `vyos_service_monitoring_network_event`

* New Resource `vyos_service_monitoring_network_event_event`









## Release 18.0.202501220 (2025-01-22 08-30-54 UTC)
### Project changes

### Schema changes
#### BREAKING CHANGES

##### Resources
* **Removed Resource** `vyos_system_flow_accounting_sflow_server`

* **Removed Resource** `vyos_system_flow_accounting_sflow`





#### Features

##### Resources
* Modified Resource `vyos_system_sflow`
	* New attribute `enable_egress`








## Previous changelogs
For previous version see [changelog for v17](data/changelogs/CHANGELOG-v17.md) or older archives [directory](data/changelogs/)
