---
page_title: "vyos_protocols_bgp_bmp_target Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Border Gateway Protocol (BGP)⯯BGP Monitoring Protocol (BMP)⯯BMP target
---

# vyos_protocols_bgp_bmp_target (Resource)
<center>

*protocols*  
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

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|ipv4    &emsp;|IPv4 address  |
    &emsp;|ipv6    &emsp;|IPv6 address  |
- `max_retry` (Number) Maximum connection retry interval

    &emsp;|Format          &emsp;|Description                        |
    |------------------|-------------------------------------|
    &emsp;|100-4294967295  &emsp;|Maximum connection retry interval  |
- `min_retry` (Number) Minimum connection retry interval (in milliseconds)

    &emsp;|Format        &emsp;|Description                        |
    |----------------|-------------------------------------|
    &emsp;|100-86400000  &emsp;|Minimum connection retry interval  |
- `mirror` (Boolean) Send BMP route mirroring messages
- `monitor` (Attributes) Send BMP route monitoring messages (see [below for nested schema](#nestedatt--monitor))
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `target` (String) BMP target


&lt;a id=&#34;nestedatt--monitor&#34;&gt;&lt;/a&gt;
### Nested Schema for `monitor`

Optional:

- `ipv4_unicast` (Attributes) Address family IPv4 unicast (see [below for nested schema](#nestedatt--monitor--ipv4_unicast))
- `ipv6_unicast` (Attributes) Address family IPv6 unicast (see [below for nested schema](#nestedatt--monitor--ipv6_unicast))

&lt;a id=&#34;nestedatt--monitor--ipv4_unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `monitor.ipv4_unicast`

Optional:

- `post_policy` (Boolean) Send state with policy and filters applied
- `pre_policy` (Boolean) Send state before policy and filter processing


&lt;a id=&#34;nestedatt--monitor--ipv6_unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `monitor.ipv6_unicast`

Optional:

- `post_policy` (Boolean) Send state with policy and filters applied
- `pre_policy` (Boolean) Send state before policy and filter processing



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
