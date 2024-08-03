---
page_title: "vyos_service_ssh Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  System services⯯Secure Shell (SSH)
---

# vyos_service_ssh (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

System services  
⯯  
**Secure Shell (SSH)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `ciphers` (List of String) Allowed ciphers
- `client_keepalive_interval` (Number) Enable transmission of keepalives from server to client

    &emsp;|Format   &emsp;|Description                                     |
    |-----------|--------------------------------------------------|
    &emsp;|1-65535  &emsp;|Time interval in seconds for keepalive message  |
- `disable_host_validation` (Boolean) Disable IP Address to Hostname lookup
- `disable_password_authentication` (Boolean) Disable password-based authentication
- `hostkey_algorithm` (List of String) Allowed host key signature algorithms
- `key_exchange` (List of String) Allowed key exchange (KEX) algorithms
- `listen_address` (List of String) Local IP addresses to listen on

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address to listen for incoming connections  |
    &emsp;|ipv6    &emsp;|IPv6 address to listen for incoming connections  |
- `loglevel` (String) Log level

    &emsp;|Format   &emsp;|Description                              |
    |-----------|-------------------------------------------|
    &emsp;|quiet    &emsp;|stay silent                              |
    &emsp;|fatal    &emsp;|log fatals only                          |
    &emsp;|error    &emsp;|log errors and fatals only               |
    &emsp;|info     &emsp;|default log level                        |
    &emsp;|verbose  &emsp;|enable logging of failed login attempts  |
- `mac` (List of String) Allowed message authentication code (MAC) algorithms
- `port` (List of Number) Port for SSH service

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `pubkey_accepted_algorithm` (List of String) Allowed pubkey signature algorithms
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (List of String) VRF instance name

    &emsp;|Format   &emsp;|Description                      |
    |-----------|-----------------------------------|
    &emsp;|default  &emsp;|Explicitly start in default VRF  |
    &emsp;|txt      &emsp;|VRF instance name                |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
