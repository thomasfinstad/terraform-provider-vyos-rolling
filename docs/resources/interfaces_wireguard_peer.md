---
page_title: "vyos_interfaces_wireguard_peer Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯WireGuard Interface⯯peer alias
---

# vyos_interfaces_wireguard_peer (Resource)
<center>

*interfaces*  
⯯  
WireGuard Interface  
⯯  
**peer alias**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (String) IP address of tunnel endpoint

    |Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    |ipv4    &emsp;|IPv4 address of remote tunnel endpoint  |
    |ipv6    &emsp;|IPv6 address of remote tunnel endpoint  |
- `allowed_ips` (List of String) IP addresses allowed to traverse the peer
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `disable` (Boolean) Disable instance
- `persistent_keepalive` (Number) Interval to send keepalive messages

    |Format   &emsp;|Description          |
    |-----------|-----------------------|
    |1-65535  &emsp;|Interval in seconds  |
- `port` (Number) Port number used by connection

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |1-65535  &emsp;|Numeric IP port  |
- `preshared_key` (String) base64 encoded preshared key
- `public_key` (String) base64 encoded public key
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `peer` (String) peer alias
- `wireguard` (String) WireGuard Interface

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |wgN     &emsp;|WireGuard interface name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
