---
page_title: "vyos_vrf_name_protocols_ospfv3_interface Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Open Shortest Path First (OSPF) for IPv6⯯Enable routing on an IPv6 interface
---

# vyos_vrf_name_protocols_ospfv3_interface (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Open Shortest Path First (OSPF) for IPv6  
⯯  
**Enable routing on an IPv6 interface**


</center>

## Schema

### Required

- `interface_id` (String) Enable routing on an IPv6 interface

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|txt     &emsp;|Interface used for routing information exchange  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Optional

- `area` (String) Enable OSPF on this interface

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|u32     &emsp;|OSPF area ID as decimal notation     |
    &emsp;|ipv4    &emsp;|OSPF area ID in IP address notation  |
- `bfd` (Attributes) Enable Bidirectional Forwarding Detection (BFD) (see [below for nested schema](#nestedatt--bfd))
- `cost` (Number) Interface cost

    &emsp;|Format   &emsp;|Description          |
    |-----------|-----------------------|
    &emsp;|1-65535  &emsp;|OSPF interface cost  |
- `dead_interval` (Number) Interval after which a neighbor is declared dead

    &emsp;|Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    &emsp;|1-65535  &emsp;|Neighbor dead interval (seconds)  |
- `hello_interval` (Number) Interval between hello packets

    &emsp;|Format   &emsp;|Description               |
    |-----------|----------------------------|
    &emsp;|1-65535  &emsp;|Hello interval (seconds)  |
- `ifmtu` (Number) Interface MTU

    &emsp;|Format   &emsp;|Description    |
    |-----------|-----------------|
    &emsp;|1-65535  &emsp;|Interface MTU  |
- `instance_id` (Number) Instance ID

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|0-255   &emsp;|Instance Id  |
- `mtu_ignore` (Boolean) Disable Maximum Transmission Unit (MTU) mismatch detection
- `network` (String) Network type

    &emsp;|Format          &emsp;|Description                  |
    |------------------|-------------------------------|
    &emsp;|broadcast       &emsp;|Broadcast network type       |
    &emsp;|point-to-point  &emsp;|Point-to-point network type  |
- `passive` (Boolean) Configure passive mode for interface
- `priority` (Number) Router priority

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|0-255   &emsp;|OSPF router priority cost  |
- `retransmit_interval` (Number) Interval between retransmitting lost link state advertisements

    &emsp;|Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    &emsp;|1-65535  &emsp;|Retransmit interval (seconds)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `transmit_delay` (Number) Link state transmit delay

    &emsp;|Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    &emsp;|1-65535  &emsp;|Link state transmit delay (seconds)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--bfd&#34;&gt;&lt;/a&gt;
### Nested Schema for `bfd`

Optional:

- `profile` (String) Use settings from BFD profile

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|BFD profile name  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
