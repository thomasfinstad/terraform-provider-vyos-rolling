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

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|u32     &emsp;|Buffer size in MiB  |
- `disable_imt` (Boolean) Disable in memory table plugin
- `enable_egress` (Boolean) Enable egress flow accounting
- `interface` (List of String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `packet_length` (Number) Specifies the maximum number of bytes to capture for each packet

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|128-750  &emsp;|Packet length in bytes  |
- `syslog_facility` (String) Syslog facility for flow-accounting

    &emsp;|Format     &emsp;|Description                       |
    |-------------|------------------------------------|
    &emsp;|auth       &emsp;|Authentication and authorization  |
    &emsp;|authpriv   &emsp;|Non-system authorization          |
    &emsp;|cron       &emsp;|Cron daemon                       |
    &emsp;|daemon     &emsp;|System daemons                    |
    &emsp;|kern       &emsp;|Kernel                            |
    &emsp;|lpr        &emsp;|Line printer spooler              |
    &emsp;|mail       &emsp;|Mail subsystem                    |
    &emsp;|mark       &emsp;|Timestamp                         |
    &emsp;|news       &emsp;|USENET subsystem                  |
    &emsp;|protocols  &emsp;|Routing protocols (local7)        |
    &emsp;|security   &emsp;|Authentication and authorization  |
    &emsp;|syslog     &emsp;|Authentication and authorization  |
    &emsp;|user       &emsp;|Application processes             |
    &emsp;|uucp       &emsp;|UUCP subsystem                    |
    &emsp;|local0     &emsp;|Local facility 0                  |
    &emsp;|local1     &emsp;|Local facility 1                  |
    &emsp;|local2     &emsp;|Local facility 2                  |
    &emsp;|local3     &emsp;|Local facility 3                  |
    &emsp;|local4     &emsp;|Local facility 4                  |
    &emsp;|local5     &emsp;|Local facility 5                  |
    &emsp;|local6     &emsp;|Local facility 6                  |
    &emsp;|local7     &emsp;|Local facility 7                  |
    &emsp;|all        &emsp;|Authentication and authorization  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF instance name

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
