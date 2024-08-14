---
page_title: "vyos_high_availability_vrrp_group Resource - vyos"

subcategory: "High Availability"

description: |- 
  High availability settings⯯Virtual Router Redundancy Protocol settings⯯VRRP group
---

# vyos_high_availability_vrrp_group (Resource)
<center>

High availability settings  
⯯  
Virtual Router Redundancy Protocol settings  
⯯  
**VRRP group**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `advertise_interval` (Number) Advertise interval

    |Format  &emsp;|Description                    |
    |----------|---------------------------------|
    |1-255   &emsp;|Advertise interval in seconds  |
- `authentication` (Attributes) VRRP authentication (see [below for nested schema](#nestedatt--authentication))
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `disable` (Boolean) Disable instance
- `garp` (Attributes) Gratuitous ARP parameters (see [below for nested schema](#nestedatt--garp))
- `health_check` (Attributes) Health check (see [below for nested schema](#nestedatt--health_check))
- `hello_source_address` (String) VRRP hello source address

    |Format  &emsp;|Description                |
    |----------|-----------------------------|
    |ipv4    &emsp;|IPv4 hello source address  |
    |ipv6    &emsp;|IPv6 hello source address  |
- `interface` (String) Interface to use

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |
- `no_preempt` (Boolean) Disable master preemption
- `peer_address` (List of String) Unicast VRRP peer address

    |Format  &emsp;|Description                |
    |----------|-----------------------------|
    |ipv4    &emsp;|IPv4 unicast peer address  |
    |ipv6    &emsp;|IPv6 unicast peer address  |
- `preempt_delay` (Number) Preempt delay (in seconds)

    |Format  &emsp;|Description    |
    |----------|-----------------|
    |0-1000  &emsp;|preempt delay  |
- `priority` (Number) Router priority

    |Format  &emsp;|Description      |
    |----------|-------------------|
    |1-255   &emsp;|Router priority  |
- `rfc3768_compatibility` (Boolean) Use VRRP virtual MAC address as per RFC3768
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `track` (Attributes) Track settings (see [below for nested schema](#nestedatt--track))
- `transition_script` (Attributes) VRRP transition scripts (see [below for nested schema](#nestedatt--transition_script))
- `vrid` (Number) Virtual router identifier

    |Format  &emsp;|Description                |
    |----------|-----------------------------|
    |1-255   &emsp;|Virtual router identifier  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `group` (String) VRRP group


<a id="nestedatt--authentication"></a>
### Nested Schema for `authentication`

Optional:

- `password` (String) VRRP password

    |Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    |txt     &emsp;|Password string (up to 8 characters)  |
- `type` (String) Authentication type

    |Format              &emsp;|Description                   |
    |----------------------|--------------------------------|
    |plaintext-password  &emsp;|Simple password string        |
    |ah                  &emsp;|AH - IPSEC (not recommended)  |


<a id="nestedatt--garp"></a>
### Nested Schema for `garp`

Optional:

- `interval` (String) Interval between Gratuitous ARP

    |Format        &emsp;|Description                                   |
    |----------------|------------------------------------------------|
    |&lt;0.000-1000&gt;  &emsp;|Interval in seconds, resolution microseconds  |
- `master_delay` (Number) Delay for second set of gratuitous ARPs after transition to master

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |1-1000  &emsp;|Delay in seconds  |
- `master_refresh` (Number) Minimum time interval for refreshing gratuitous ARPs while beeing master

    |Format  &emsp;|Description          |
    |----------|-----------------------|
    |0       &emsp;|No refresh           |
    |1-255   &emsp;|Interval in seconds  |
- `master_refresh_repeat` (Number) Number of gratuitous ARP messages to send at a time while beeing master

    |Format  &emsp;|Description                        |
    |----------|-------------------------------------|
    |1-255   &emsp;|Number of gratuitous ARP messages  |
- `master_repeat` (Number) Number of gratuitous ARP messages to send at a time after transition to master

    |Format  &emsp;|Description                        |
    |----------|-------------------------------------|
    |1-255   &emsp;|Number of gratuitous ARP messages  |


<a id="nestedatt--health_check"></a>
### Nested Schema for `health_check`

Optional:

- `failure_count` (String) Health check failure count required for transition to fault
- `interval` (String) Health check execution interval in seconds
- `ping` (String) ICMP ping health check

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |ipv4    &emsp;|IPv4 ping target address  |
    |ipv6    &emsp;|IPv6 ping target address  |
- `script` (String) Health check script file


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


<a id="nestedatt--track"></a>
### Nested Schema for `track`

Optional:

- `exclude_vrrp_interface` (Boolean) Disable track state of main interface
- `interface` (List of String) Interface name state check

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |


<a id="nestedatt--transition_script"></a>
### Nested Schema for `transition_script`

Optional:

- `backup` (String) Script to run on VRRP state transition to backup
- `fault` (String) Script to run on VRRP state transition to fault
- `master` (String) Script to run on VRRP state transition to master
- `stop` (String) Script to run on VRRP state transition to stop  
