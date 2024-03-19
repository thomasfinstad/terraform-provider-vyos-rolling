---
page_title: "vyos_high_availability_vrrp_sync_group Resource - terraform-provider-vyos"
subcategory: "high"
description: |-
  High availability settings
  ⯯
  Virtual Router Redundancy Protocol settings
  ⯯
  VRRP sync group
---

# vyos_high_availability_vrrp_sync_group (Resource)
<center>

High availability settings
⯯
Virtual Router Redundancy Protocol settings
⯯
**VRRP sync group**


</center>

## Schema

### Required

- `sync_group_id` (String) VRRP sync group

### Optional

- `health_check` (Attributes) Health check (see [below for nested schema](#nestedatt--health_check))
- `member` (List of String) Sync group member

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|VRRP group name  |
- `transition_script` (Attributes) VRRP transition scripts (see [below for nested schema](#nestedatt--transition_script))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

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


&lt;a id=&#34;nestedatt--transition_script&#34;&gt;&lt;/a&gt;
### Nested Schema for `transition_script`

Optional:

- `backup` (String) Script to run on VRRP state transition to backup
- `fault` (String) Script to run on VRRP state transition to fault
- `master` (String) Script to run on VRRP state transition to master
- `stop` (String) Script to run on VRRP state transition to stop  &emsp;|
