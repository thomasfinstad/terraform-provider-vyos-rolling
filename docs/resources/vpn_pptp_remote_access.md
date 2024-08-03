---
page_title: "vyos_vpn_pptp_remote_access Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  vpn⯯Point to Point Tunneling Protocol (PPTP) Virtual Private Network (VPN)⯯Remote access PPTP VPN
---

# vyos_vpn_pptp_remote_access (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*vpn*  
⯯  
Point to Point Tunneling Protocol (PPTP) Virtual Private Network (VPN)  
⯯  
**Remote access PPTP VPN**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `default_ipv6_pool` (String) Default client IPv6 pool name

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|Default IPv6 pool  |
- `default_pool` (String) Default client IP pool name

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|Default IP pool  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `gateway_address` (String) Gateway IP address

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|ipv4    &emsp;|Default Gateway send to the client  |
- `max_concurrent_sessions` (Number) Maximum number of concurrent session start attempts

    &emsp;|Format   &emsp;|Description                                          |
    |-----------|-------------------------------------------------------|
    &emsp;|0-65535  &emsp;|Maximum number of concurrent session start attempts  |
- `mtu` (String) Maximum Transmission Unit (MTU)
- `name_server` (List of String) Domain Name Servers (DNS) addresses

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|ipv4    &emsp;|Domain Name Server (DNS) IPv4 address  |
    &emsp;|ipv6    &emsp;|Domain Name Server (DNS) IPv6 address  |
- `outside_address` (String) External IP address to which VPN clients will connect
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `wins_server` (List of String) Windows Internet Name Service (WINS) servers propagated to client

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|ipv4    &emsp;|Domain Name Server (DNS) IPv4 address  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
