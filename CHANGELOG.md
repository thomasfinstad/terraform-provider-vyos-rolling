
# CHANGELOG

<!-- ToC start -->
1. [CHANGELOG](#changelog)
   1. [0.4.20240813 (2024-08-13 08-24-24 UTC)](#0420240813-2024-08-13-08-24-24-utc)
   1. [0.4.20240806 (2024-08-06 08-24-19 UTC)](#0420240806-2024-08-06-08-24-19-utc)
   1. [0.4.20240803 (2024-08-03 08-23-39 UTC)](#0420240803-2024-08-03-08-23-39-utc)
   1. [0.4.20240801 (2024-08-01 08-24-17 UTC)](#0420240801-2024-08-01-08-24-17-utc)
   1. [0.4.20240725 (2024-07-30 08-24-02 UTC)](#0420240725-2024-07-30-08-24-02-utc)
   1. [0.3.20240725 (2024-07-25 08-24-01 UTC)](#0320240725-2024-07-25-08-24-01-utc)
   1. [0.3.20240717 (2024-07-21 08-22-11 UTC)](#0320240717-2024-07-21-08-22-11-utc)
   1. [0.3.20240710 (2024-07-10 08-29-07 UTC)](#0320240710-2024-07-10-08-29-07-utc)
   1. [0.3.20240709 (2024-07-09 08-22-43 UTC)](#0320240709-2024-07-09-08-22-43-utc)
   1. [0.3.20240707 (2024-07-07 08-22-13 UTC)](#0320240707-2024-07-07-08-22-13-utc)
   1. [0.3.20240706 (2024-07-06 08-23-33 UTC)](#0320240706-2024-07-06-08-23-33-utc)
   1. [0.3.20240705 (2024-07-05 08-24-23 UTC)](#0320240705-2024-07-05-08-24-23-utc)
   1. [0.3.20240704 (2024-07-04 08-23-56 UTC)](#0320240704-2024-07-04-08-23-56-utc)
   1. [0.3.20240702 (2024-07-03 08-23-41 UTC)](#0320240702-2024-07-03-08-23-41-utc)
   1. [0.3.20240701 (2024-07-01 08-23-44 UTC)](#0320240701-2024-07-01-08-23-44-utc)
   1. [0.3.20240630 (2024-06-30 08-21-28 UTC)](#0320240630-2024-06-30-08-21-28-utc)
   1. [0.3.20240629 (2024-06-29 08-22-20 UTC)](#0320240629-2024-06-29-08-22-20-utc)
   1. [0.3.20240628 (2024-06-28 08-22-19 UTC)](#0320240628-2024-06-28-08-22-19-utc)
   1. [0.3.20240627 (2024-06-27 08-23-29 UTC)](#0320240627-2024-06-27-08-23-29-utc)
   1. [0.3.20240626 (2024-06-26 08-23-08 UTC)](#0320240626-2024-06-26-08-23-08-utc)
   1. [0.3.20240625 (2024-06-25 08-23-29 UTC)](#0320240625-2024-06-25-08-23-29-utc)
   1. [0.3.20240623 (2024-06-23 08-21-01 UTC)](#0320240623-2024-06-23-08-21-01-utc)
   1. [0.3.20240621 (2024-06-21 08-22-55 UTC)](#0320240621-2024-06-21-08-22-55-utc)
   1. [0.3.20240620 (2024-06-20 08-23-39 UTC)](#0320240620-2024-06-20-08-23-39-utc)
   1. [0.3.20240619 (2024-06-19 08-23-09 UTC)](#0320240619-2024-06-19-08-23-09-utc)
   1. [0.3.20240617 (2024-06-17 08-23-38 UTC)](#0320240617-2024-06-17-08-23-38-utc)
   1. [0.3.20240613 (2024-06-13 08-23-27 UTC)](#0320240613-2024-06-13-08-23-27-utc)
   1. [0.3.20240612 (2024-06-12 08-22-53 UTC)](#0320240612-2024-06-12-08-22-53-utc)
   1. [0.3.20240606 (2024-06-06 08-23-02 UTC)](#0320240606-2024-06-06-08-23-02-utc)
   1. [0.3.20240602 (2024-06-02 08-21-22 UTC)](#0320240602-2024-06-02-08-21-22-utc)
   1. [0.3.20240601 (2024-06-01 08-23-16 UTC)](#0320240601-2024-06-01-08-23-16-utc)
   1. [0.3.20240531 (2024-05-31 08-27-03 UTC)](#0320240531-2024-05-31-08-27-03-utc)
   1. [0.3.20240528 (2024-05-28 08-24-26 UTC)](#0320240528-2024-05-28-08-24-26-utc)
   1. [0.3.20240527 (2024-05-27 08-23-01 UTC)](#0320240527-2024-05-27-08-23-01-utc)
   1. [0.3.20240526 (2024-05-26 08-22-01 UTC)](#0320240526-2024-05-26-08-22-01-utc)
   1. [0.3.20240524 (2024-05-24 08-23-39 UTC)](#0320240524-2024-05-24-08-23-39-utc)
   1. [0.3.20240522 (2024-05-22 08-23-17 UTC)](#0320240522-2024-05-22-08-23-17-utc)
   1. [0.3.20240521 (2024-05-21 08-22-14 UTC)](#0320240521-2024-05-21-08-22-14-utc)
   1. [0.3.20240520 (2024-05-20 08-22-55 UTC)](#0320240520-2024-05-20-08-22-55-utc)
   1. [0.3.20240518 (2024-05-19 08-20-42 UTC)](#0320240518-2024-05-19-08-20-42-utc)
   1. [0.3.20240514 (2024-05-14 08-23-36 UTC)](#0320240514-2024-05-14-08-23-36-utc)
   1. [0.3.20240512 (2024-05-13 08-23-13 UTC)](#0320240512-2024-05-13-08-23-13-utc)
   1. [0.3.20240510 (2024-05-10 08-21-55 UTC)](#0320240510-2024-05-10-08-21-55-utc)
   1. [0.3.20240509 (2024-05-09 08-22-09 UTC)](#0320240509-2024-05-09-08-22-09-utc)
   1. [0.3.20240508 (2024-05-08 08-23-55 UTC)](#0320240508-2024-05-08-08-23-55-utc)
   1. [0.3.20240507 (2024-05-07 08-21-46 UTC)](#0320240507-2024-05-07-08-21-46-utc)
   1. [0.3.20240506 (2024-05-06 10-32-57 UTC)](#0320240506-2024-05-06-10-32-57-utc)
   1. [0.3.20240505 (2024-05-05 13-02-39 UTC)](#0320240505-2024-05-05-13-02-39-utc)
   1. [0.3.20240323 (2024-05-01 15-38-08 UTC)](#0320240323-2024-05-01-15-38-08-utc)
   1. [0.2.20240323 (2024-04-28 12-32-19 UTC)](#0220240323-2024-04-28-12-32-19-utc)
   1. [0.2.20240324 (2024-04-28 12-20-45 UTC)](#0220240324-2024-04-28-12-20-45-utc)
   1. [0.1.20240323 (2024-04-27 17-21-15 UTC)](#0120240323-2024-04-27-17-21-15-utc)
   1. [0.0.20240322 (2024-03-22 00-00-00 UTC)](#0020240322-2024-03-22-00-00-00-utc)
<!-- ToC end -->

## 0.4.20240813 (2024-08-13 08-24-24 UTC)
NOTES:

* update to rolling release 2024-08-13T00-21-40Z

* update to rolling release 2024-08-12T00-21-49Z

* update to rolling release 2024-08-11T00-23-16Z

* update to rolling release 2024-08-10T00-20-35Z

* update to rolling release 2024-08-09T00-20-49Z

* update to rolling release 2024-08-08T00-20-40Z





## 0.4.20240806 (2024-08-06 08-24-19 UTC)
NOTES:

* update to rolling release 2024-08-06T00-20-41Z

* update to rolling release 2024-08-04T00-22-49Z





## 0.4.20240803 (2024-08-03 08-23-39 UTC)
NOTES:

* update to rolling release 2024-08-03T00-20-04Z

* update to rolling release 2024-08-01T10-47-52Z





## 0.4.20240801 (2024-08-01 08-24-17 UTC)
NOTES:

* update to rolling release 2024-08-01T00-23-04Z





## 0.4.20240725 (2024-07-30 08-24-02 UTC)
NOTES:

* update to rolling release 2024-07-30T00-20-39Z

* update to rolling release 2024-07-28T00-23-03Z

* update to rolling release 2024-07-27T00-19-54Z

* update to rolling release 2024-07-26T00-20-13Z


FEATURES:

* Resource/`vyos_firewall_ipv4_forward_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_forward_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv4_input_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_input_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv4_name_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_name_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv4_output_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_output_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv4_output_raw_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_output_raw_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv4_prerouting_raw_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv4_prerouting_raw_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_forward_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_forward_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_input_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_input_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_name_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_name_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_output_filter_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_output_filter_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_output_raw_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_output_raw_rule` Added attribute `match_none`

* Resource/`vyos_firewall_ipv6_prerouting_raw_rule` Added attribute `match_ipsec`

* Resource/`vyos_firewall_ipv6_prerouting_raw_rule` Added attribute `match_none`





## 0.3.20240725 (2024-07-25 08-24-01 UTC)
NOTES:

* update to rolling release 2024-07-25T00-20-19Z


BUG FIXES:

* generate files with relative path in autogen comment

* move autogen comment to use relative path to avoid changes between machines





## 0.3.20240717 (2024-07-21 08-22-11 UTC)
NOTES:

* update to rolling release 2024-07-17T17-05-42Z





## 0.3.20240710 (2024-07-10 08-29-07 UTC)
NOTES:

* update to rolling release 2024-07-10T00-20-26Z





## 0.3.20240709 (2024-07-09 08-22-43 UTC)
NOTES:

* update to rolling release 2024-07-09T00-20-00Z





## 0.3.20240707 (2024-07-07 08-22-13 UTC)
NOTES:

* update to rolling release 2024-07-07T00-22-17Z





## 0.3.20240706 (2024-07-06 08-23-33 UTC)
NOTES:

* update to rolling release 2024-07-06T00-19-12Z





## 0.3.20240705 (2024-07-05 08-24-23 UTC)
NOTES:

* update to rolling release 2024-07-05T00-19-47Z





## 0.3.20240704 (2024-07-04 08-23-56 UTC)
NOTES:

* update to rolling release 2024-07-04T00-19-47Z





## 0.3.20240702 (2024-07-03 08-23-41 UTC)
NOTES:

* update to rolling release 2024-07-02T18-39-02Z





## 0.3.20240701 (2024-07-01 08-23-44 UTC)
NOTES:

* update to rolling release 2024-07-01T00-23-26Z





## 0.3.20240630 (2024-06-30 08-21-28 UTC)
NOTES:

* update to rolling release 2024-06-30T00-22-05Z





## 0.3.20240629 (2024-06-29 08-22-20 UTC)
NOTES:

* update to rolling release 2024-06-29T00-19-21Z





## 0.3.20240628 (2024-06-28 08-22-19 UTC)
NOTES:

* update to rolling release 2024-06-28T00-19-42Z





## 0.3.20240627 (2024-06-27 08-23-29 UTC)
NOTES:

* update to rolling release 2024-06-27T00-19-46Z





## 0.3.20240626 (2024-06-26 08-23-08 UTC)
NOTES:

* update to rolling release 2024-06-26T00-19-30Z





## 0.3.20240625 (2024-06-25 08-23-29 UTC)
NOTES:

* update to rolling release 2024-06-25T00-19-32Z





## 0.3.20240623 (2024-06-23 08-21-01 UTC)
NOTES:

* update to rolling release 2024-06-23T00-21-32Z





## 0.3.20240621 (2024-06-21 08-22-55 UTC)
NOTES:

* update to rolling release 2024-06-21T00-19-18Z





## 0.3.20240620 (2024-06-20 08-23-39 UTC)
NOTES:

* update to rolling release 2024-06-20T00-19-27Z





## 0.3.20240619 (2024-06-19 08-23-09 UTC)
NOTES:

* update to rolling release 2024-06-19T00-19-47Z





## 0.3.20240617 (2024-06-17 08-23-38 UTC)
NOTES:

* update to rolling release 2024-06-17T00-20-56Z





## 0.3.20240613 (2024-06-13 08-23-27 UTC)
NOTES:

* update to rolling release 2024-06-13T00-19-41Z





## 0.3.20240612 (2024-06-12 08-22-53 UTC)
NOTES:

* update to rolling release 2024-06-12T00-19-44Z





## 0.3.20240606 (2024-06-06 08-23-02 UTC)
NOTES:

* update to rolling release 2024-06-06T00-19-21Z





## 0.3.20240602 (2024-06-02 08-21-22 UTC)
NOTES:

* update to rolling release 2024-06-02T00-21-06Z





## 0.3.20240601 (2024-06-01 08-23-16 UTC)
NOTES:

* update to rolling release 2024-06-01T00-20-56Z





## 0.3.20240531 (2024-05-31 08-27-03 UTC)
NOTES:

* update to rolling release 2024-05-31T00-18-53Z





## 0.3.20240528 (2024-05-28 08-24-26 UTC)
NOTES:

* update to rolling release 2024-05-28T00-19-17Z





## 0.3.20240527 (2024-05-27 08-23-01 UTC)
NOTES:

* update to rolling release 2024-05-27T00-19-48Z





## 0.3.20240526 (2024-05-26 08-22-01 UTC)
NOTES:

* update to rolling release 2024-05-26T00-21-03Z





## 0.3.20240524 (2024-05-24 08-23-39 UTC)
NOTES:

* update to rolling release 2024-05-24T00-19-23Z





## 0.3.20240522 (2024-05-22 08-23-17 UTC)
NOTES:

* update to rolling release 2024-05-22T00-18-44Z





## 0.3.20240521 (2024-05-21 08-22-14 UTC)
NOTES:

* update to rolling release 2024-05-21T00-19-19Z





## 0.3.20240520 (2024-05-20 08-22-55 UTC)
NOTES:

* update to rolling release 2024-05-20T00-19-28Z





## 0.3.20240518 (2024-05-19 08-20-42 UTC)
NOTES:

* update to rolling release 2024-05-18T12-11-15Z





## 0.3.20240514 (2024-05-14 08-23-36 UTC)
NOTES:

* update to rolling release 2024-05-14T00-18-58Z





## 0.3.20240512 (2024-05-13 08-23-13 UTC)
NOTES:

* update to rolling release 2024-05-12T14-02-26Z

* Add notetice about provider being incomplete





## 0.3.20240510 (2024-05-10 08-21-55 UTC)
NOTES:

* update to rolling release 2024-05-10T00-18-53Z





## 0.3.20240509 (2024-05-09 08-22-09 UTC)
NOTES:

* update to rolling release 2024-05-09T00-18-37Z





## 0.3.20240508 (2024-05-08 08-23-55 UTC)
NOTES:

* update to rolling release 2024-05-08T00-15-46Z





## 0.3.20240507 (2024-05-07 08-21-46 UTC)
NOTES:

* update to rolling release 2024-05-07T00-18-45Z





## 0.3.20240506 (2024-05-06 10-32-57 UTC)
NOTES:

* update to rolling release 2024-05-06T00-19-12Z

* attempt to add changelog to release notes





## 0.3.20240505 (2024-05-05 13-02-39 UTC)
BUG FIXES:

* makefile improvements and CI fixes





## 0.3.20240323 (2024-05-01 15-38-08 UTC)
FEATURES:

* enable daily automatic release to match rolling release

* prepare ci-update make target

* Move devcontainer to Dockerfile

* Resource/`vyos_container_name` Added attribute `cap_add`

* Resource/`vyos_high_availability_vrrp_group` Added attribute `excluded_address`


BUG FIXES:

* auto update github action





## 0.2.20240323 (2024-04-28 12-32-19 UTC)
FEATURES:

* use goreleaser to release provider


BUG FIXES:

* add missing manifest required by goreleaser

* removed unintended increase in fix version





## 0.2.20240324 (2024-04-28 12-20-45 UTC)
FEATURES:

* use goreleaser to release provider





## 0.1.20240323 (2024-04-27 17-21-15 UTC)
NOTES:

* beginning of changelog generation


FEATURES:

* Autogenerate changelog


BUG FIXES:

* Improve publishing target

* publish packaging zips and docs





## 0.0.20240322 (2024-03-22 00-00-00 UTC)

Initial release
