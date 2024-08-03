---
page_title: "vyos_interfaces_sstpc Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Secure Socket Tunneling Protocol (SSTP) client Interface
---

# vyos_interfaces_sstpc (Resource)
<center>

*interfaces*  
⯯  
**Secure Socket Tunneling Protocol (SSTP) client Interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `authentication` (Attributes) Authentication settings (see [below for nested schema](#nestedatt--authentication))
- `default_route_distance` (Number) Distance for installed default route

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|1-255   &emsp;|Distance for the default route from DHCP server  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `disable` (Boolean) Administratively disable interface
- `mtu` (Number) Maximum Transmission Unit (MTU)

    &emsp;|Format   &emsp;|Description                        |
    |-----------|-------------------------------------|
    &emsp;|68-1500  &emsp;|Maximum Transmission Unit in byte  |
- `no_default_route` (Boolean) Do not install default route to system
- `no_peer_dns` (Boolean) Do not use DNS servers provided by the peer
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `server` (String) Remote server to connect to

    &emsp;|Format    &emsp;|Description           |
    |------------|------------------------|
    &emsp;|ipv4      &emsp;|Server IPv4 address   |
    &emsp;|hostname  &emsp;|Server hostname/FQDN  |
- `ssl` (Attributes) Secure Sockets Layer (SSL) configuration (see [below for nested schema](#nestedatt--ssl))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF instance name

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `sstpc` (String) Secure Socket Tunneling Protocol (SSTP) client Interface

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|sstpcN  &emsp;|Secure Socket Tunneling Protocol interface name  |


&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `password` (String) Password used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Password     |
- `username` (String) Username used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Username     |


&lt;a id=&#34;nestedatt--ssl&#34;&gt;&lt;/a&gt;
### Nested Schema for `ssl`

Optional:

- `ca_certificate` (String) Certificate Authority in PKI configuration

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|Name of CA in PKI configuration  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
