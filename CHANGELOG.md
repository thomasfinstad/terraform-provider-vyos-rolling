
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 5.0.202410150 (2024-10-15 15-53-09 UTC)](#release-50202410150-2024-10-15-15-53-09-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
      - [Bug fixes](#bug-fixes)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


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
