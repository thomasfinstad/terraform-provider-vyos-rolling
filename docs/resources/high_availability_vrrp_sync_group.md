---
page_title: "vyos_high_availability_vrrp_sync_group Resource - vyos"

subcategory: "High Availability"

description: |-
  High availability settings⯯Virtual Router Redundancy Protocol settings⯯VRRP sync group
---

# vyos_high_availability_vrrp_sync_group (Resource)
<center>


High availability settings  
⯯  
Virtual Router Redundancy Protocol settings  
⯯  
**VRRP sync group**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_high_availability_vrrp_sync_group (Resource)](#vyos_high_availability_vrrp_sync_group-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [health_check](#health_check)
      - [member](#member)
      - [timeouts](#timeouts)
      - [transition_script](#transition_script)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `health_check`](#nested-schema-for-health_check)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
    - [Nested Schema for `transition_script`](#nested-schema-for-transition_script)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### health_check
- `health_check` (Attributes) Health check (see [below for nested schema](#nestedatt--health_check))
#### member
- `member` (List of String) Sync group member

    |  Format  &emsp;|  Description      |
    |----------|-------------------|
    |  txt     &emsp;|  VRRP group name  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
#### transition_script
- `transition_script` (Attributes) VRRP transition scripts (see [below for nested schema](#nestedatt--transition_script))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `sync_group` (String) VRRP sync group


<a id="nestedatt--health_check"></a>
### Nested Schema for `health_check`

Optional:

- `failure_count` (String) Health check failure count required for transition to fault
- `interval` (String) Health check execution interval in seconds
- `ping` (String) ICMP ping health check

    |  Format  &emsp;|  Description               |
    |----------|----------------------------|
    |  ipv4    &emsp;|  IPv4 ping target address  |
    |  ipv6    &emsp;|  IPv6 ping target address  |
- `script` (String) Health check script file


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


<a id="nestedatt--transition_script"></a>
### Nested Schema for `transition_script`

Optional:

- `backup` (String) Script to run on VRRP state transition to backup
- `fault` (String) Script to run on VRRP state transition to fault
- `master` (String) Script to run on VRRP state transition to master
- `stop` (String) Script to run on VRRP state transition to stop

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_high_availability_vrrp_sync_group.example "high-availability__vrrp__sync-group__<sync-group>"
```
