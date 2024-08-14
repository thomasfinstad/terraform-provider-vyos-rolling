---
page_title: "vyos_vrf_name_protocols_bgp_bmp_target Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Border Gateway Protocol (BGP)⯯BGP Monitoring Protocol (BMP)⯯BMP target
---

# vyos_vrf_name_protocols_bgp_bmp_target (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP Monitoring Protocol (BMP)  
⯯  
**BMP target**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (String) IP address

    |Format  &emsp;|Description   |
    |----------|----------------|
    |ipv4    &emsp;|IPv4 address  |
    |ipv6    &emsp;|IPv6 address  |
- `max_retry` (Number) Maximum connection retry interval

    |Format          &emsp;|Description                        |
    |------------------|-------------------------------------|
    |100-4294967295  &emsp;|Maximum connection retry interval  |
- `min_retry` (Number) Minimum connection retry interval (in milliseconds)

    |Format        &emsp;|Description                        |
    |----------------|-------------------------------------|
    |100-86400000  &emsp;|Minimum connection retry interval  |
- `mirror` (Boolean) Send BMP route mirroring messages
- `monitor` (Attributes) Send BMP route monitoring messages (see [below for nested schema](#nestedatt--monitor))
- `port` (Number) Port number used by connection

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |1-65535  &emsp;|Numeric IP port  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `name` (String) Virtual Routing and Forwarding instance

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |
- `target` (String) BMP target


<a id="nestedatt--monitor"></a>
### Nested Schema for `monitor`

Optional:

- `ipv4_unicast` (Attributes) Address family IPv4 unicast (see [below for nested schema](#nestedatt--monitor--ipv4_unicast))
- `ipv6_unicast` (Attributes) Address family IPv6 unicast (see [below for nested schema](#nestedatt--monitor--ipv6_unicast))

<a id="nestedatt--monitor--ipv4_unicast"></a>
### Nested Schema for `monitor.ipv4_unicast`

Optional:

- `post_policy` (Boolean) Send state with policy and filters applied
- `pre_policy` (Boolean) Send state before policy and filter processing


<a id="nestedatt--monitor--ipv6_unicast"></a>
### Nested Schema for `monitor.ipv6_unicast`

Optional:

- `post_policy` (Boolean) Send state with policy and filters applied
- `pre_policy` (Boolean) Send state before policy and filter processing



<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
