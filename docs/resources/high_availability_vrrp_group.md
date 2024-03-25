---
page_title: "vyos_high_availability_vrrp_group Resource - vyos"
subcategory: "high"
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

## Schema

### Required

- `group_id` (String) VRRP group

### Optional

- `advertise_interval` (Number) Advertise interval

    &emsp;|Format  &emsp;|Description                    |
    |----------|---------------------------------|
    &emsp;|1-255   &emsp;|Advertise interval in seconds  |
- `authentication` (Attributes) VRRP authentication (see [below for nested schema](#nestedatt--authentication))
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `disable` (Boolean) Disable instance
- `excluded_address` (List of String) Virtual address (If you need additional IPv4 and IPv6 in same group)

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|ipv4    &emsp;|IP address    |
    &emsp;|ipv6    &emsp;|IPv6 address  |
- `garp` (Attributes) Gratuitous ARP parameters (see [below for nested schema](#nestedatt--garp))
- `health_check` (Attributes) Health check (see [below for nested schema](#nestedatt--health_check))
- `hello_source_address` (String) VRRP hello source address

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|ipv4    &emsp;|IPv4 hello source address  |
    &emsp;|ipv6    &emsp;|IPv6 hello source address  |
- `interface` (String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `no_preempt` (Boolean) Disable master preemption
- `peer_address` (List of String) Unicast VRRP peer address

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|ipv4    &emsp;|IPv4 unicast peer address  |
    &emsp;|ipv6    &emsp;|IPv6 unicast peer address  |
- `preempt_delay` (Number) Preempt delay (in seconds)

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|0-1000  &emsp;|preempt delay  |
- `priority` (Number) Router priority

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|1-255   &emsp;|Router priority  |
- `rfc3768_compatibility` (Boolean) Use VRRP virtual MAC address as per RFC3768
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `track` (Attributes) Track settings (see [below for nested schema](#nestedatt--track))
- `transition_script` (Attributes) VRRP transition scripts (see [below for nested schema](#nestedatt--transition_script))
- `vrid` (Number) Virtual router identifier

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|1-255   &emsp;|Virtual router identifier  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `password` (String) VRRP password

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|txt     &emsp;|Password string (up to 8 characters)  |
- `type` (String) Authentication type

    &emsp;|Format              &emsp;|Description                   |
    |----------------------|--------------------------------|
    &emsp;|plaintext-password  &emsp;|Simple password string        |
    &emsp;|ah                  &emsp;|AH - IPSEC (not recommended)  |


&lt;a id=&#34;nestedatt--garp&#34;&gt;&lt;/a&gt;
### Nested Schema for `garp`

Optional:

- `interval` (String) Interval between Gratuitous ARP

    &emsp;|Format        &emsp;|Description                                   |
    |----------------|------------------------------------------------|
    &emsp;|&lt;0.000-1000&gt;  &emsp;|Interval in seconds, resolution microseconds  |
- `master_delay` (Number) Delay for second set of gratuitous ARPs after transition to master

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|1-1000  &emsp;|Delay in seconds  |
- `master_refresh` (Number) Minimum time interval for refreshing gratuitous ARPs while beeing master

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|0       &emsp;|No refresh           |
    &emsp;|1-255   &emsp;|Interval in seconds  |
- `master_refresh_repeat` (Number) Number of gratuitous ARP messages to send at a time while beeing master

    &emsp;|Format  &emsp;|Description                        |
    |----------|-------------------------------------|
    &emsp;|1-255   &emsp;|Number of gratuitous ARP messages  |
- `master_repeat` (Number) Number of gratuitous ARP messages to send at a time after transition to master

    &emsp;|Format  &emsp;|Description                        |
    |----------|-------------------------------------|
    &emsp;|1-255   &emsp;|Number of gratuitous ARP messages  |


&lt;a id=&#34;nestedatt--health_check&#34;&gt;&lt;/a&gt;
### Nested Schema for `health_check`

Optional:

- `failure_count` (String) Health check failure count required for transition to fault
- `interval` (String) Health check execution interval in seconds
- `ping` (String) ICMP ping health check

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|ipv4    &emsp;|IPv4 ping target address  |
    &emsp;|ipv6    &emsp;|IPv6 ping target address  |
- `script` (String) Health check script file


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--track&#34;&gt;&lt;/a&gt;
### Nested Schema for `track`

Optional:

- `exclude_vrrp_interface` (Boolean) Disable track state of main interface
- `interface` (List of String) Interface name state check

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |


&lt;a id=&#34;nestedatt--transition_script&#34;&gt;&lt;/a&gt;
### Nested Schema for `transition_script`

Optional:

- `backup` (String) Script to run on VRRP state transition to backup
- `fault` (String) Script to run on VRRP state transition to fault
- `master` (String) Script to run on VRRP state transition to master
- `stop` (String) Script to run on VRRP state transition to stop  &emsp;|
