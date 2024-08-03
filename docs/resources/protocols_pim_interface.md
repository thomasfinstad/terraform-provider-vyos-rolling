---
page_title: "vyos_protocols_pim_interface Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Protocol Independent Multicast (PIM) and IGMP⯯PIM interface
---

# vyos_protocols_pim_interface (Resource)
<center>

*protocols*  
⯯  
Protocol Independent Multicast (PIM) and IGMP  
⯯  
**PIM interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `bfd` (Attributes) Enable Bidirectional Forwarding Detection (BFD) (see [below for nested schema](#nestedatt--bfd))
- `dr_priority` (Number) Designated router election priority

    &emsp;|Format        &emsp;|Description  |
    |----------------|---------------|
    &emsp;|1-4294967295  &emsp;|DR Priority  |
- `hello` (Number) Hello Interval

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|1-180   &emsp;|Hello Interval in seconds  |
- `igmp` (Attributes) Internet Group Management Protocol (IGMP) options (see [below for nested schema](#nestedatt--igmp))
- `no_bsm` (Boolean) Do not process bootstrap messages
- `no_unicast_bsm` (Boolean) Do not process unicast bootstrap messages
- `passive` (Boolean) Disable sending and receiving PIM control packets on the interface
- `source_address` (String) IPv4 source address used to initiate connection

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|ipv4    &emsp;|IPv4 source address  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) PIM interface


&lt;a id=&#34;nestedatt--bfd&#34;&gt;&lt;/a&gt;
### Nested Schema for `bfd`

Optional:

- `profile` (String) Use settings from BFD profile

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|BFD profile name  |


&lt;a id=&#34;nestedatt--igmp&#34;&gt;&lt;/a&gt;
### Nested Schema for `igmp`

Optional:

- `disable` (Boolean) Disable instance
- `query_interval` (Number) IGMP host query interval

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|1-1800  &emsp;|Query interval in seconds  |
- `query_max_response_time` (Number) IGMP max query response time

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|10-250  &emsp;|Query response value in deci-seconds  |
- `version` (String) Interface IGMP version

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|2       &emsp;|IGMP version 2  |
    &emsp;|3       &emsp;|IGMP version 3  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
