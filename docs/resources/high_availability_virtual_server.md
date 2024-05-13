---
page_title: "vyos_high_availability_virtual_server Resource - vyos"

subcategory: "High Availability"

description: |- 
  High availability settings⯯Load-balancing virtual server alias
---

# vyos_high_availability_virtual_server (Resource)
<center>

High availability settings  
⯯  
**Load-balancing virtual server alias**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `virtual_server_id` (String) Load-balancing virtual server alias

### Optional

- `address` (String) IP address

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|ipv4    &emsp;|IPv4 address  |
    &emsp;|ipv6    &emsp;|IPv6 address  |
- `algorithm` (String) Schedule algorithm (default - least-connection)

    &emsp;|Format                           &emsp;|Description                      |
    |-----------------------------------|-----------------------------------|
    &emsp;|round-robin                      &emsp;|Round robin                      |
    &emsp;|weighted-round-robin             &emsp;|Weighted round robin             |
    &emsp;|least-connection                 &emsp;|Least connection                 |
    &emsp;|weighted-least-connection        &emsp;|Weighted least connection        |
    &emsp;|source-hashing                   &emsp;|Source hashing                   |
    &emsp;|destination-hashing              &emsp;|Destination hashing              |
    &emsp;|locality-based-least-connection  &emsp;|Locality-Based least connection  |
- `delay_loop` (Number) Interval between health-checks (in seconds)

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|1-600   &emsp;|Interval in seconds  |
- `forward_method` (String) Forwarding method

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|direct  &emsp;|Direct routing  |
    &emsp;|nat     &emsp;|NAT             |
    &emsp;|tunnel  &emsp;|Tunneling       |
- `fwmark` (Number) Match fwmark value

    &emsp;|Format        &emsp;|Description                |
    |----------------|-----------------------------|
    &emsp;|1-2147483647  &emsp;|Match firewall mark value  |
- `persistence_timeout` (Number) Timeout for persistent connections

    &emsp;|Format   &emsp;|Description                         |
    |-----------|--------------------------------------|
    &emsp;|1-86400  &emsp;|Timeout for persistent connections  |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|0-65535  &emsp;|Numeric IP port  |
- `protocol` (String) Protocol for port checks

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|tcp     &emsp;|TCP          |
    &emsp;|udp     &emsp;|UDP          |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
