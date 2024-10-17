
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 5.1.202410170 (2024-10-17 14-56-24 UTC)](#release-51202410170-2024-10-17-14-56-24-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
      - [Features](#features)
  - [Release 5.0.202410170 (2024-10-17 10-59-50 UTC)](#release-50202410170-2024-10-17-10-59-50-utc)
    - [Project changes](#project-changes-1)
      - [Notes](#notes-1)
      - [Bug fixes](#bug-fixes)
  - [Release 5.0.202410150 (2024-10-15 15-53-09 UTC)](#release-50202410150-2024-10-15-15-53-09-utc)
    - [Project changes](#project-changes-2)
      - [Notes](#notes-2)
      - [Bug fixes](#bug-fixes-1)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 5.1.202410170 (2024-10-17 14-56-24 UTC)
### Project changes
#### Notes
* update to rolling release 2024-10-17T00:06:01Z
#### Features
* implement resource importing


## Release 5.0.202410170 (2024-10-17 10-59-50 UTC)
### Project changes
#### Notes
* update to rolling release 2024-10-17T00:06:01Z
#### Bug fixes
* unconfigured lists caused panic


## Release 5.0.202410150 (2024-10-15 15-53-09 UTC)
### Project changes
#### Notes
* update to rolling release 2024-10-15T00:05:58Z
* initial res gen template merging
* split resource templates
* Improve code generation reliability
#### Bug fixes
* allows for IP addresses to be used in named resource identifiers as described in #222
* generate go from xsd with installed tooling instead of imported tooling

### Schema changes
#### BREAKING CHANGES

##### Resources
* **Removed Resource** `vyos_service_ntp_allow_client`

* **Removed Resource** `vyos_service_ntp_server`








## Previous changelogs
For previous version see [changelog for 4](data/changelogs/CHANGELOG-4.md) or older archives [directory](data/changelogs/)
