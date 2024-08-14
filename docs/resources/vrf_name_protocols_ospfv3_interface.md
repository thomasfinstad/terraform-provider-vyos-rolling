---
page_title: "vyos_vrf_name_protocols_ospfv3_interface Resource - vyos"

subcategory: "Vrf"

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

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `area` (String) Enable OSPF on this interface

    |Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    |u32     &emsp;|OSPF area ID as decimal notation     |
    |ipv4    &emsp;|OSPF area ID in IP address notation  |
- `bfd` (Attributes) Enable Bidirectional Forwarding Detection (BFD) (see [below for nested schema](#nestedatt--bfd))
- `cost` (Number) Interface cost

    |Format   &emsp;|Description          |
    |-----------|-----------------------|
    |1-65535  &emsp;|OSPF interface cost  |
- `dead_interval` (Number) Interval after which a neighbor is declared dead

    |Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    |1-65535  &emsp;|Neighbor dead interval (seconds)  |
- `hello_interval` (Number) Interval between hello packets

    |Format   &emsp;|Description               |
    |-----------|----------------------------|
    |1-65535  &emsp;|Hello interval (seconds)  |
- `ifmtu` (Number) Interface MTU

    |Format   &emsp;|Description    |
    |-----------|-----------------|
    |1-65535  &emsp;|Interface MTU  |
- `instance_id` (Number) Instance ID

    |Format  &emsp;|Description  |
    |----------|---------------|
    |0-255   &emsp;|Instance Id  |
- `mtu_ignore` (Boolean) Disable Maximum Transmission Unit (MTU) mismatch detection
- `network` (String) Network type

    |Format          &emsp;|Description                  |
    |------------------|-------------------------------|
    |broadcast       &emsp;|Broadcast network type       |
    |point-to-point  &emsp;|Point-to-point network type  |
- `passive` (Boolean) Configure passive mode for interface
- `priority` (Number) Router priority

    |Format  &emsp;|Description                |
    |----------|-----------------------------|
    |0-255   &emsp;|OSPF router priority cost  |
- `retransmit_interval` (Number) Interval between retransmitting lost link state advertisements

    |Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    |1-65535  &emsp;|Retransmit interval (seconds)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `transmit_delay` (Number) Link state transmit delay

    |Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    |1-65535  &emsp;|Link state transmit delay (seconds)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Enable routing on an IPv6 interface

    |Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    |txt     &emsp;|Interface used for routing information exchange  |
- `name` (String) Virtual Routing and Forwarding instance

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |


<a id="nestedatt--bfd"></a>
### Nested Schema for `bfd`

Optional:

- `profile` (String) Use settings from BFD profile

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |txt     &emsp;|BFD profile name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
