---
page_title: "vyos_service_dhcp_server_shared_network_name_subnet Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Dynamic Host Configuration Protocol (DHCP) for DHCP server⯯Name of DHCP shared network⯯DHCP subnet for shared network
---

# vyos_service_dhcp_server_shared_network_name_subnet (Resource)
<center>

*service*  
⯯  
Dynamic Host Configuration Protocol (DHCP) for DHCP server  
⯯  
Name of DHCP shared network  
⯯  
**DHCP subnet for shared network**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `disable` (Boolean) Disable instance
- `exclude` (List of String) IP address to exclude from DHCP lease range

    &emsp;|Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address to exclude from lease range  |
- `ignore_client_id` (Boolean) Ignore client identifier for lease lookups
- `lease` (Number) Lease timeout in seconds

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|u32     &emsp;|DHCP lease time in seconds  |
- `option` (Attributes) DHCP option (see [below for nested schema](#nestedatt--option))
- `subnet_id` (Number) Unique ID mapped to leases in the lease file

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|u32     &emsp;|Unique subnet ID  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `shared_network_name` (String) Name of DHCP shared network
- `subnet` (String) DHCP subnet for shared network

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 address and prefix length  |


&lt;a id=&#34;nestedatt--option&#34;&gt;&lt;/a&gt;
### Nested Schema for `option`

Optional:

- `bootfile_name` (String) Bootstrap file name
- `bootfile_server` (String) Server from which the initial boot file is to be loaded

    &emsp;|Format    &emsp;|Description                   |
    |------------|--------------------------------|
    &emsp;|ipv4      &emsp;|Bootfile server IPv4 address  |
    &emsp;|hostname  &emsp;|Bootfile server FQDN          |
- `bootfile_size` (Number) Bootstrap file size

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|1-16    &emsp;|Bootstrap file size in 512 byte blocks  |
- `captive_portal` (String) Captive portal API endpoint

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|txt     &emsp;|Captive portal API endpoint  |
- `client_prefix_length` (Number) Specifies the clients subnet mask as per RFC 950. If unset, subnet declaration is used.

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|0-32    &emsp;|DHCP client prefix length must be 0 to 32  |
- `default_router` (String) IP address of default router

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|ipv4    &emsp;|Default router IPv4 address  |
- `domain_name` (String) Client Domain Name
- `domain_search` (List of String) Client Domain Name search list
- `ip_forwarding` (Boolean) Enable IP forwarding on client
- `ipv6_only_preferred` (Number) Disable IPv4 on IPv6 only hosts (RFC 8925)

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|u32     &emsp;|Seconds      |
- `name_server` (List of String) Domain Name Servers (DNS) addresses

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|ipv4    &emsp;|Domain Name Server (DNS) IPv4 address  |
- `ntp_server` (List of String) IP address of NTP server

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|ipv4    &emsp;|NTP server IPv4 address  |
- `pop_server` (List of String) IP address of POP3 server

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|ipv4    &emsp;|POP3 server IPv4 address  |
- `server_identifier` (String) Address for DHCP server identifier

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|ipv4    &emsp;|DHCP server identifier IPv4 address  |
- `smtp_server` (List of String) IP address of SMTP server

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|ipv4    &emsp;|SMTP server IPv4 address  |
- `tftp_server_name` (String) TFTP server name

    &emsp;|Format    &emsp;|Description               |
    |------------|----------------------------|
    &emsp;|ipv4      &emsp;|TFTP server IPv4 address  |
    &emsp;|hostname  &emsp;|TFTP server FQDN          |
- `time_offset` (String) Client subnet offset in seconds from Coordinated Universal Time (UTC)

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|[-]N    &emsp;|Time offset (number, may be negative)  |
- `time_server` (List of String) IP address of time server

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|ipv4    &emsp;|Time server IPv4 address  |
- `time_zone` (String) Time zone to send to clients. Uses RFC4833 options 100 and 101
- `vendor_option` (Attributes) Vendor Specific Options (see [below for nested schema](#nestedatt--option--vendor_option))
- `wins_server` (List of String) IP address for Windows Internet Name Service (WINS) server

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|ipv4    &emsp;|WINS server IPv4 address  |
- `wpad_url` (String) Web Proxy Autodiscovery (WPAD) URL

&lt;a id=&#34;nestedatt--option--vendor_option&#34;&gt;&lt;/a&gt;
### Nested Schema for `option.vendor_option`

Optional:

- `ubiquiti` (Attributes) Ubiquiti specific parameters (see [below for nested schema](#nestedatt--option--vendor_option--ubiquiti))

&lt;a id=&#34;nestedatt--option--vendor_option--ubiquiti&#34;&gt;&lt;/a&gt;
### Nested Schema for `option.vendor_option.ubiquiti`

Optional:

- `unifi_controller` (String) Address of UniFi controller

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv4    &emsp;|IP address of UniFi controller  |




&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  