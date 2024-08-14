---
page_title: "vyos_system_flow_accounting Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯Flow accounting settings
---

# vyos_system_flow_accounting (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
**Flow accounting settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `buffer_size` (Number) Buffer size

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |u32     &emsp;|Buffer size in MiB  |
- `disable_imt` (Boolean) Disable in memory table plugin
- `enable_egress` (Boolean) Enable egress flow accounting
- `interface` (List of String) Interface to use

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |
- `packet_length` (Number) Specifies the maximum number of bytes to capture for each packet

    |Format   &emsp;|Description             |
    |-----------|--------------------------|
    |128-750  &emsp;|Packet length in bytes  |
- `syslog_facility` (String) Syslog facility for flow-accounting

    |Format     &emsp;|Description                       |
    |-------------|------------------------------------|
    |auth       &emsp;|Authentication and authorization  |
    |authpriv   &emsp;|Non-system authorization          |
    |cron       &emsp;|Cron daemon                       |
    |daemon     &emsp;|System daemons                    |
    |kern       &emsp;|Kernel                            |
    |lpr        &emsp;|Line printer spooler              |
    |mail       &emsp;|Mail subsystem                    |
    |mark       &emsp;|Timestamp                         |
    |news       &emsp;|USENET subsystem                  |
    |protocols  &emsp;|Routing protocols (local7)        |
    |security   &emsp;|Authentication and authorization  |
    |syslog     &emsp;|Authentication and authorization  |
    |user       &emsp;|Application processes             |
    |uucp       &emsp;|UUCP subsystem                    |
    |local0     &emsp;|Local facility 0                  |
    |local1     &emsp;|Local facility 1                  |
    |local2     &emsp;|Local facility 2                  |
    |local3     &emsp;|Local facility 3                  |
    |local4     &emsp;|Local facility 4                  |
    |local5     &emsp;|Local facility 5                  |
    |local6     &emsp;|Local facility 6                  |
    |local7     &emsp;|Local facility 7                  |
    |all        &emsp;|Authentication and authorization  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF instance name

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
