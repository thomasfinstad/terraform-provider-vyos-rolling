---
page_title: "vyos_protocols_bfd_peer Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Bidirectional Forwarding Detection (BFD)⯯Configures BFD peer to listen and talk to
---

# vyos_protocols_bfd_peer (Resource)
<center>

*protocols*  
⯯  
Bidirectional Forwarding Detection (BFD)  
⯯  
**Configures BFD peer to listen and talk to**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `echo_mode` (Boolean) Enables the echo transmission mode
- `interval` (Attributes) Configure timer intervals (see [below for nested schema](#nestedatt--interval))
- `minimum_ttl` (Number) Expect packets with at least this TTL

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|1-254   &emsp;|Minimum TTL expected  |
- `multihop` (Boolean) Allow this BFD peer to not be directly connected
- `passive` (Boolean) Do not attempt to start sessions
- `profile` (String) Use settings from BFD profile

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|BFD profile name  |
- `shutdown` (Boolean) Disable this peer
- `source` (Attributes) Bind listener to specified interface/address, mandatory for IPv6 (see [below for nested schema](#nestedatt--source))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF instance name

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `peer` (String) Configures BFD peer to listen and talk to

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|ipv4    &emsp;|BFD peer IPv4 address  |
    &emsp;|ipv6    &emsp;|BFD peer IPv6 address  |


&lt;a id=&#34;nestedatt--interval&#34;&gt;&lt;/a&gt;
### Nested Schema for `interval`

Optional:

- `echo_interval` (Number) Echo receive transmission interval

    &emsp;|Format    &emsp;|Description                                                                             |
    |------------|------------------------------------------------------------------------------------------|
    &emsp;|10-60000  &emsp;|The minimal echo receive transmission interval that this system is capable of handling  |
- `multiplier` (Number) Multiplier to determine packet loss

    &emsp;|Format  &emsp;|Description                                                    |
    |----------|-----------------------------------------------------------------|
    &emsp;|2-255   &emsp;|Remote transmission interval will be multiplied by this value  |
- `receive` (Number) Minimum interval of receiving control packets

    &emsp;|Format    &emsp;|Description               |
    |------------|----------------------------|
    &emsp;|10-60000  &emsp;|Interval in milliseconds  |
- `transmit` (Number) Minimum interval of transmitting control packets

    &emsp;|Format    &emsp;|Description               |
    |------------|----------------------------|
    &emsp;|10-60000  &emsp;|Interval in milliseconds  |


&lt;a id=&#34;nestedatt--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `source`

Optional:

- `address` (String) Local address to bind our peer listener to

    &emsp;|Format  &emsp;|Description                                     |
    |----------|--------------------------------------------------|
    &emsp;|ipv4    &emsp;|Local IPv4 address used to connect to the peer  |
    &emsp;|ipv6    &emsp;|Local IPv6 address used to connect to the peer  |
- `interface` (String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
