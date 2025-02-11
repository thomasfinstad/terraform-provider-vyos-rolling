
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 19.0.202502110 (2025-02-11 08-30-27 UTC)](#release-190202502110-2025-02-11-08-30-27-utc)
    - [Project changes](#project-changes)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
      - [Features](#features)
        - [Resources](#resources-1)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


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
For previous version see [changelog for v18](data/changelogs/CHANGELOG-v18.md) or older archives [directory](data/changelogs/)
