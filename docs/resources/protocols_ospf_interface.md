---
page_title: "vyos_protocols_ospf_interface Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Open Shortest Path First (OSPF)⯯Interface configuration
---

# vyos_protocols_ospf_interface (Resource)
<center>

*protocols*  
⯯  
Open Shortest Path First (OSPF)  
⯯  
**Interface configuration**


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
- `authentication` (Attributes) Authentication (see [below for nested schema](#nestedatt--authentication))
- `bandwidth` (Number) Interface bandwidth (Mbit/s)

    |Format    &emsp;|Description                                           |
    |------------|--------------------------------------------------------|
    |1-100000  &emsp;|Bandwidth in Megabit/sec (for calculating OSPF cost)  |
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
- `hello_multiplier` (Number) Hello multiplier factor

    |Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    |1-10    &emsp;|Number of Hellos to send each second  |
- `ldp_sync` (Attributes) LDP-IGP synchronization configuration for interface (see [below for nested schema](#nestedatt--ldp_sync))
- `mtu_ignore` (Boolean) Disable Maximum Transmission Unit (MTU) mismatch detection
- `network` (String) Network type

    |Format               &emsp;|Description                       |
    |-----------------------|------------------------------------|
    |broadcast            &emsp;|Broadcast network type            |
    |non-broadcast        &emsp;|Non-broadcast network type        |
    |point-to-multipoint  &emsp;|Point-to-multipoint network type  |
    |point-to-point       &emsp;|Point-to-point network type       |
- `passive` (Attributes) Suppress routing updates on an interface (see [below for nested schema](#nestedatt--passive))
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

- `interface` (String) Interface configuration

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |


<a id="nestedatt--authentication"></a>
### Nested Schema for `authentication`

Optional:

- `md5` (Attributes) MD5 key id (see [below for nested schema](#nestedatt--authentication--md5))
- `plaintext_password` (String) Plain text password

    |Format  &emsp;|Description                                 |
    |----------|----------------------------------------------|
    |txt     &emsp;|Plain text password (8 characters or less)  |

<a id="nestedatt--authentication--md5"></a>
### Nested Schema for `authentication.md5`



<a id="nestedatt--bfd"></a>
### Nested Schema for `bfd`

Optional:

- `profile` (String) Use settings from BFD profile

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |txt     &emsp;|BFD profile name  |


<a id="nestedatt--ldp_sync"></a>
### Nested Schema for `ldp_sync`

Optional:

- `disable` (Boolean) Disable instance
- `holddown` (Number) Hold down timer for LDP-IGP cost restoration

    |Format   &emsp;|Description                                                                                   |
    |-----------|------------------------------------------------------------------------------------------------|
    |0-10000  &emsp;|Time to wait in seconds for LDP-IGP synchronization to occur before restoring interface cost  |


<a id="nestedatt--passive"></a>
### Nested Schema for `passive`

Optional:

- `disable` (Boolean) Disable instance


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
