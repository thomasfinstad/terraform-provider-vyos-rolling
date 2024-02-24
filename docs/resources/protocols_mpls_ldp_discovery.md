---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_mpls_ldp_discovery Resource - vyos"
subcategory: "protocols"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>protocols</i>

  <br>
  &darr;
  <br>
  Multiprotocol Label Switching (MPLS)

  <br>
  &darr;
  <br>
  Label Distribution Protocol (LDP)

  <br>
  &darr;
  <br>
  <b>
  Discovery parameters
  </b>
  </div>
---

# vyos_protocols_mpls_ldp_discovery (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Multiprotocol Label Switching (MPLS)

<br>
&darr;
<br>
Label Distribution Protocol (LDP)

<br>
&darr;
<br>
<b>
Discovery parameters
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `hello_ipv4_holdtime` (Number) Hello IPv4 hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Time in seconds  |
- `hello_ipv4_interval` (Number) Hello IPv4 interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Time in seconds  |
- `hello_ipv6_holdtime` (Number) Hello IPv6 hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Time in seconds  |
- `hello_ipv6_interval` (Number) Hello IPv6 interval

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Time in seconds  |
- `session_ipv4_holdtime` (Number) Session IPv4 hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 15-65535  &emsp; |  Time in seconds  |
- `session_ipv6_holdtime` (Number) Session IPv6 hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 15-65535  &emsp; |  Time in seconds  |
- `transport_ipv4_address` (String) Transport IPv4 address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 bind as transport  |
- `transport_ipv6_address` (String) Transport IPv6 address

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv6  &emsp; |  IPv6 bind as transport  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).