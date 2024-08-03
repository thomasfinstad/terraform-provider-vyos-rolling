---
page_title: "vyos_protocols_mpls_ldp_discovery Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Multiprotocol Label Switching (MPLS)⯯Label Distribution Protocol (LDP)⯯Discovery parameters
---

# vyos_protocols_mpls_ldp_discovery (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Multiprotocol Label Switching (MPLS)  
⯯  
Label Distribution Protocol (LDP)  
⯯  
**Discovery parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `hello_ipv4_holdtime` (Number) Hello IPv4 hold time

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Time in seconds  |
- `hello_ipv4_interval` (Number) Hello IPv4 interval

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Time in seconds  |
- `hello_ipv6_holdtime` (Number) Hello IPv6 hold time

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Time in seconds  |
- `hello_ipv6_interval` (Number) Hello IPv6 interval

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Time in seconds  |
- `session_ipv4_holdtime` (Number) Session IPv4 hold time

    &emsp;|Format    &emsp;|Description      |
    |------------|-------------------|
    &emsp;|15-65535  &emsp;|Time in seconds  |
- `session_ipv6_holdtime` (Number) Session IPv6 hold time

    &emsp;|Format    &emsp;|Description      |
    |------------|-------------------|
    &emsp;|15-65535  &emsp;|Time in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `transport_ipv4_address` (String) Transport IPv4 address

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|ipv4    &emsp;|IPv4 bind as transport  |
- `transport_ipv6_address` (String) Transport IPv6 address

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|ipv6    &emsp;|IPv6 bind as transport  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
