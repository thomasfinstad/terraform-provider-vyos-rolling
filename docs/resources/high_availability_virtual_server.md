---
page_title: "vyos_high_availability_virtual_server Resource - terraform-provider-vyos"
subcategory: "high"
description: |-
  High availability settings
  ⯯
  Load-balancing virtual server alias
---

# vyos_high_availability_virtual_server (Resource)
<center>

High availability settings
⯯
**Load-balancing virtual server alias**


</center>

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

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
