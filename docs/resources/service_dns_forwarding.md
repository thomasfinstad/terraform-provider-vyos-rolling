---
page_title: "vyos_service_dns_forwarding Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Domain Name System (DNS) related services⯯DNS forwarding
---

# vyos_service_dns_forwarding (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Domain Name System (DNS) related services  
⯯  
**DNS forwarding**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `allow_from` (List of String) Networks allowed to query this server

    |Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    |ipv4net  &emsp;|IP address and prefix length    |
    |ipv6net  &emsp;|IPv6 address and prefix length  |
- `cache_size` (Number) DNS forwarding cache size

    |Format        &emsp;|Description                |
    |----------------|-----------------------------|
    |0-2147483647  &emsp;|DNS forwarding cache size  |
- `dhcp` (List of String) Interfaces whose DHCP client nameservers to forward requests to
- `dns64_prefix` (String) Help to communicate between IPv6-only client and IPv4-only server

    |Format   &emsp;|Description                              |
    |-----------|-------------------------------------------|
    |ipv6net  &emsp;|IPv6 address and /96 only prefix length  |
- `dnssec` (String) DNSSEC mode

    |Format               &emsp;|Description                                                                                      |
    |-----------------------|---------------------------------------------------------------------------------------------------|
    |off                  &emsp;|No DNSSEC processing whatsoever!                                                                 |
    |process-no-validate  &emsp;|Respond with DNSSEC records to clients that ask for it. No validation done at all!               |
    |process              &emsp;|Respond with DNSSEC records to clients that ask for it. Validation for clients that request it.  |
    |log-fail             &emsp;|Similar behaviour to process, but validate RRSIGs on responses and log bogus responses.          |
    |validate             &emsp;|Full blown DNSSEC validation. Send SERVFAIL to clients on bogus responses.                       |
- `exclude_throttle_address` (List of String) IP address or subnet

    |Format   &emsp;|Description            |
    |-----------|-------------------------|
    |ipv4     &emsp;|IPv4 address to match  |
    |ipv4net  &emsp;|IPv4 prefix to match   |
    |ipv6     &emsp;|IPv6 address           |
    |ipv6net  &emsp;|IPv6 address           |
- `ignore_hosts_file` (Boolean) Do not use local /etc/hosts file in name resolution
- `listen_address` (List of String) Local IP addresses to listen on

    |Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    |ipv4    &emsp;|IPv4 address to listen for incoming connections  |
    |ipv6    &emsp;|IPv6 address to listen for incoming connections  |
- `negative_ttl` (Number) Maximum amount of time negative entries are cached

    |Format  &emsp;|Description                        |
    |----------|-------------------------------------|
    |0-7200  &emsp;|Seconds to cache NXDOMAIN entries  |
- `no_serve_rfc1918` (Boolean) Makes the server authoritatively not aware of RFC1918 addresses
- `port` (Number) Port number used by connection

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |1-65535  &emsp;|Numeric IP port  |
- `serve_stale_extension` (Number) Number of times the expired TTL of a record is extended by 30 seconds when serving stale

    |Format   &emsp;|Description                        |
    |-----------|-------------------------------------|
    |0-65535  &emsp;|Number of times to extend the TTL  |
- `source_address` (List of String) Source IP address used to initiate connection

    |Format  &emsp;|Description          |
    |----------|-----------------------|
    |ipv4    &emsp;|IPv4 source address  |
    |ipv6    &emsp;|IPv6 source address  |
- `system` (Boolean) Use system name servers
- `timeout` (Number) Number of milliseconds to wait for a remote authoritative server to respond

    |Format    &emsp;|Description                      |
    |------------|-----------------------------------|
    |10-60000  &emsp;|Network timeout in milliseconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
