---
page_title: "vyos_service_dhcpv6_server_shared_network_name_subnet Resource - vyos"

subcategory: "Service"

description: |-
  service⯯DHCP for IPv6 (DHCPv6) server⯯DHCPv6 shared network name⯯IPv6 DHCP subnet for this shared network
---

# vyos_service_dhcpv6_server_shared_network_name_subnet (Resource)
<center>


*service*  
⯯  
DHCP for IPv6 (DHCPv6) server  
⯯  
DHCPv6 shared network name  
⯯  
**IPv6 DHCP subnet for this shared network**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_service_dhcpv6_server_shared_network_name_subnet (Resource)](#vyos_service_dhcpv6_server_shared_network_name_subnet-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [interface](#interface)
      - [lease_time](#lease_time)
      - [option](#option)
      - [subnet_id](#subnet_id)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `lease_time`](#nested-schema-for-lease_time)
    - [Nested Schema for `option`](#nested-schema-for-option)
    - [Nested Schema for `option.vendor_option`](#nested-schema-for-optionvendor_option)
    - [Nested Schema for `option.vendor_option.cisco`](#nested-schema-for-optionvendor_optioncisco)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### interface
- `interface` (String) Interface

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Interface name  |
#### lease_time
- `lease_time` (Attributes) Parameters relating to the lease time (see [below for nested schema](#nestedatt--lease_time))
#### option
- `option` (Attributes) DHCPv6 option (see [below for nested schema](#nestedatt--option))
#### subnet_id
- `subnet_id` (Number) Unique ID mapped to leases in the lease file

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  u32     &emsp;|  Unique subnet ID  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `shared_network_name` (String) DHCPv6 shared network name
- `subnet` (String) IPv6 DHCP subnet for this shared network

    |  Format   &emsp;|  Description                     |
    |-----------|----------------------------------|
    |  ipv6net  &emsp;|  IPv6 address and prefix length  |


<a id="nestedatt--lease_time"></a>
### Nested Schema for `lease_time`

Optional:

- `default` (Number) Default time (in seconds) that will be assigned to a lease

    |  Format        &emsp;|  Description            |
    |----------------|-------------------------|
    |  1-4294967295  &emsp;|  DHCPv6 valid lifetime  |
- `maximum` (Number) Maximum time (in seconds) that will be assigned to a lease

    |  Format        &emsp;|  Description                    |
    |----------------|---------------------------------|
    |  1-4294967295  &emsp;|  Maximum lease time in seconds  |
- `minimum` (Number) Minimum time (in seconds) that will be assigned to a lease

    |  Format        &emsp;|  Description                    |
    |----------------|---------------------------------|
    |  1-4294967295  &emsp;|  Minimum lease time in seconds  |


<a id="nestedatt--option"></a>
### Nested Schema for `option`

Optional:

- `captive_portal` (String) Captive portal API endpoint

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  txt     &emsp;|  Captive portal API endpoint  |
- `domain_search` (List of String) Client Domain Name search list
- `info_refresh_time` (Number) Time (in seconds) that stateless clients should wait between refreshing the information they were given

    |  Format        &emsp;|  Description                      |
    |----------------|-----------------------------------|
    |  1-4294967295  &emsp;|  DHCPv6 information refresh time  |
- `name_server` (List of String) Domain Name Servers (DNS) addresses

    |  Format  &emsp;|  Description                            |
    |----------|-----------------------------------------|
    |  ipv6    &emsp;|  Domain Name Server (DNS) IPv6 address  |
- `nis_domain` (String) NIS domain name for client to use
- `nis_server` (List of String) IPv6 address of a NIS Server

    |  Format  &emsp;|  Description                 |
    |----------|------------------------------|
    |  ipv6    &emsp;|  IPv6 address of NIS server  |
- `nisplus_domain` (String) NIS+ domain name for client to use
- `nisplus_server` (List of String) IPv6 address of a NIS+ Server

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  ipv6    &emsp;|  IPv6 address of NIS+ server  |
- `sip_server` (List of String) IPv6 address of SIP server

    |  Format    &emsp;|  Description                 |
    |------------|------------------------------|
    |  ipv6      &emsp;|  IPv6 address of SIP server  |
    |  hostname  &emsp;|  FQDN of SIP server          |
- `sntp_server` (List of String) IPv6 address of an SNTP server for client to use
- `vendor_option` (Attributes) Vendor Specific Options (see [below for nested schema](#nestedatt--option--vendor_option))

<a id="nestedatt--option--vendor_option"></a>
### Nested Schema for `option.vendor_option`

Optional:

- `cisco` (Attributes) Cisco specific parameters (see [below for nested schema](#nestedatt--option--vendor_option--cisco))

<a id="nestedatt--option--vendor_option--cisco"></a>
### Nested Schema for `option.vendor_option.cisco`

Optional:

- `tftp_server` (List of String) TFTP server name

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  ipv6    &emsp;|  TFTP server IPv6 address  |




<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_service_dhcpv6_server_shared_network_name_subnet.example "service__dhcpv6-server__shared-network-name__<shared-network-name>__subnet__<subnet>"
```
