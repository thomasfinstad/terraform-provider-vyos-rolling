
# CHANGELOG

<!--TOC-->

- [CHANGELOG](#changelog)
  - [Release 10.0.202411130 (2024-11-14 08-31-32 UTC)](#release-100202411130-2024-11-14-08-31-32-utc)
    - [Project changes](#project-changes)
      - [Notes](#notes)
    - [Schema changes](#schema-changes)
      - [BREAKING CHANGES](#breaking-changes)
        - [Resources](#resources)
      - [Notes](#notes-1)
        - [Resources](#resources-1)
  - [Previous changelogs](#previous-changelogs)

<!--TOC-->


## Release 10.0.202411130 (2024-11-14 08-31-32 UTC)
### Project changes
#### Notes
* update to rolling release 2024-11-13T00:05:53Z

### Schema changes
#### BREAKING CHANGES

##### Resources
* Modified Resource `vyos_protocols_babel_redistribute_ipv4`
	* Modified attribute `kernel` changed description
	* **Removed attribute** `nhrp`
	* Modified attribute `static` changed description
	* Modified attribute `bgp` changed description
	* **Removed attribute** `eigrp`
	* Modified attribute `isis` changed description
	* Modified attribute `connected` changed description
	* New attribute `openfabric`

* Modified Resource `vyos_protocols_babel_redistribute_ipv6`
	* Modified attribute `isis` changed description
	* Modified attribute `kernel` changed description
	* **Removed attribute** `nhrp`
	* Modified attribute `ospfv3` changed description
	* Modified attribute `ripng` changed description
	* Modified attribute `bgp` changed description
	* Modified attribute `connected` changed description
	* Modified attribute `static` changed description
	* New attribute `openfabric`





#### Notes

##### Resources
* Modified Resource `vyos_protocols_babel_distribute_list_ipv6_access_list`
	* Modified attributes under: `in`
		* New attribute ``
		* Modified attribute `` changed description
	* Modified attributes under: `out`
		* New attribute ``
		* Modified attribute `` changed description

* Modified Resource `vyos_protocols_ripng_distribute_list_access_list`
	* Modified attributes under: `in`
		* New attribute ``
		* Modified attribute `` changed description
	* Modified attributes under: `out`
		* New attribute ``
		* Modified attribute `` changed description

* Modified Resource `vyos_protocols_ripng_distribute_list_interface`
	* Modified attributes under: `access_list`
		* Modified attributes under: `in`
			* New attribute ``
			* Modified attribute `` changed description
		* Modified attributes under: `out`
			* New attribute ``
			* Modified attribute `` changed description

* Modified Resource `vyos_protocols_babel_distribute_list_ipv6_interface`
	* Modified attributes under: `access_list`
		* Modified attributes under: `in`
			* New attribute ``
			* Modified attribute `` changed description
		* Modified attributes under: `out`
			* New attribute ``
			* Modified attribute `` changed description








## Previous changelogs
For previous version see [changelog for v9](data/changelogs/CHANGELOG-v9.md) or older archives [directory](data/changelogs/)
