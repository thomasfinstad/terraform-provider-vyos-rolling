---
page_title: "vyos_firewall_global_options Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Firewall  
  ⯯Global Options
---

# vyos_firewall_global_options (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

	Firewall  
⯯  
**Global Options**


</center>

## Schema

### Optional

- `all_ping` (String) Policy for handling of all IPv4 ICMP echo requests

    &emsp;|Format   &emsp;|Description                                        |
    |-----------|-----------------------------------------------------|
    &emsp;|enable   &emsp;|Enable processing of all IPv4 ICMP echo requests   |
    &emsp;|disable  &emsp;|Disable processing of all IPv4 ICMP echo requests  |
- `broadcast_ping` (String) Policy for handling broadcast IPv4 ICMP echo and timestamp requests

    &emsp;|Format   &emsp;|Description                                                        |
    |-----------|---------------------------------------------------------------------|
    &emsp;|enable   &emsp;|Enable processing of broadcast IPv4 ICMP echo/timestamp requests   |
    &emsp;|disable  &emsp;|Disable processing of broadcast IPv4 ICMP echo/timestamp requests  |
- `ip_src_route` (String) Policy for handling IPv4 packets with source route option

    &emsp;|Format   &emsp;|Description                                                  |
    |-----------|---------------------------------------------------------------|
    &emsp;|enable   &emsp;|Enable processing of IPv4 packets with source route option   |
    &emsp;|disable  &emsp;|Disable processing of IPv4 packets with source route option  |
- `ipv6_receive_redirects` (String) Policy for handling received ICMPv6 redirect messages

    &emsp;|Format   &emsp;|Description                                              |
    |-----------|-----------------------------------------------------------|
    &emsp;|enable   &emsp;|Enable processing of received ICMPv6 redirect messages   |
    &emsp;|disable  &emsp;|Disable processing of received ICMPv6 redirect messages  |
- `ipv6_source_validation` (String) Policy for IPv6 source validation by reversed path, as specified in RFC3704

    &emsp;|Format   &emsp;|Description                                                       |
    |-----------|--------------------------------------------------------------------|
    &emsp;|strict   &emsp;|Enable IPv6 Strict Reverse Path Forwarding as defined in RFC3704  |
    &emsp;|loose    &emsp;|Enable IPv6 Loose Reverse Path Forwarding as defined in RFC3704   |
    &emsp;|disable  &emsp;|No IPv6 source validation                                         |
- `ipv6_src_route` (String) Policy for handling IPv6 packets with routing extension header

    &emsp;|Format   &emsp;|Description                                                   |
    |-----------|----------------------------------------------------------------|
    &emsp;|enable   &emsp;|Enable processing of IPv6 packets with routing header type 2  |
    &emsp;|disable  &emsp;|Disable processing of IPv6 packets with routing header        |
- `log_martians` (String) Policy for logging IPv4 packets with invalid addresses

    &emsp;|Format   &emsp;|Description                                             |
    |-----------|----------------------------------------------------------|
    &emsp;|enable   &emsp;|Enable logging of IPv4 packets with invalid addresses   |
    &emsp;|disable  &emsp;|Disable logging of Ipv4 packets with invalid addresses  |
- `receive_redirects` (String) Policy for handling received IPv4 ICMP redirect messages

    &emsp;|Format   &emsp;|Description                                                 |
    |-----------|--------------------------------------------------------------|
    &emsp;|enable   &emsp;|Enable processing of received IPv4 ICMP redirect messages   |
    &emsp;|disable  &emsp;|Disable processing of received IPv4 ICMP redirect messages  |
- `resolver_cache` (Boolean) Retains last successful value if domain resolution fails
- `resolver_interval` (Number) Domain resolver update interval

    &emsp;|Format   &emsp;|Description         |
    |-----------|----------------------|
    &emsp;|10-3600  &emsp;|Interval (seconds)  |
- `send_redirects` (String) Policy for sending IPv4 ICMP redirect messages

    &emsp;|Format   &emsp;|Description                                  |
    |-----------|-----------------------------------------------|
    &emsp;|enable   &emsp;|Enable sending IPv4 ICMP redirect messages   |
    &emsp;|disable  &emsp;|Disable sending IPv4 ICMP redirect messages  |
- `source_validation` (String) Policy for IPv4 source validation by reversed path, as specified in RFC3704

    &emsp;|Format   &emsp;|Description                                                       |
    |-----------|--------------------------------------------------------------------|
    &emsp;|strict   &emsp;|Enable IPv4 Strict Reverse Path Forwarding as defined in RFC3704  |
    &emsp;|loose    &emsp;|Enable IPv4 Loose Reverse Path Forwarding as defined in RFC3704   |
    &emsp;|disable  &emsp;|No IPv4 source validation                                         |
- `syn_cookies` (String) Policy for using TCP SYN cookies with IPv4

    &emsp;|Format   &emsp;|Description                               |
    |-----------|--------------------------------------------|
    &emsp;|enable   &emsp;|Enable use of TCP SYN cookies with IPv4   |
    &emsp;|disable  &emsp;|Disable use of TCP SYN cookies with IPv4  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `twa_hazards_protection` (String) RFC1337 TCP TIME-WAIT assasination hazards protection

    &emsp;|Format   &emsp;|Description                                   |
    |-----------|------------------------------------------------|
    &emsp;|enable   &emsp;|Enable RFC1337 TIME-WAIT hazards protection   |
    &emsp;|disable  &emsp;|Disable RFC1337 TIME-WAIT hazards protection  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
