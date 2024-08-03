---
page_title: "vyos_protocols_bfd_profile Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Bidirectional Forwarding Detection (BFD)⯯Configure BFD profile used by individual peer
---

# vyos_protocols_bfd_profile (Resource)
<center>

*protocols*  
⯯  
Bidirectional Forwarding Detection (BFD)  
⯯  
**Configure BFD profile used by individual peer**


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
- `passive` (Boolean) Do not attempt to start sessions
- `shutdown` (Boolean) Disable this peer
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `profile` (String) Configure BFD profile used by individual peer

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|txt     &emsp;|Name of BFD profile  |


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


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
