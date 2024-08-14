---
page_title: "vyos_load_balancing_reverse_proxy_global_parameters_logging_facility Resource - vyos"

subcategory: "Load"

description: |- 
  load-balancing⯯Configure reverse-proxy⯯Global perfomance parameters and limits⯯Logging parameters⯯Facility for logging
---

# vyos_load_balancing_reverse_proxy_global_parameters_logging_facility (Resource)
<center>

*load-balancing*  
⯯  
Configure reverse-proxy  
⯯  
Global perfomance parameters and limits  
⯯  
Logging parameters  
⯯  
**Facility for logging**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `level` (String) Logging level

    |Format   &emsp;|Description                         |
    |-----------|--------------------------------------|
    |emerg    &emsp;|Emergency messages                  |
    |alert    &emsp;|Urgent messages                     |
    |crit     &emsp;|Critical messages                   |
    |err      &emsp;|Error messages                      |
    |warning  &emsp;|Warning messages                    |
    |notice   &emsp;|Messages for further investigation  |
    |info     &emsp;|Informational messages              |
    |debug    &emsp;|Debug messages                      |
    |all      &emsp;|Log everything                      |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `facility` (String) Facility for logging

    |Format    &emsp;|Description                       |
    |------------|------------------------------------|
    |all       &emsp;|All facilities excluding &#34;mark&#34;   |
    |auth      &emsp;|Authentication and authorization  |
    |authpriv  &emsp;|Non-system authorization          |
    |cron      &emsp;|Cron daemon                       |
    |daemon    &emsp;|System daemons                    |
    |kern      &emsp;|Kernel                            |
    |lpr       &emsp;|Line printer spooler              |
    |mail      &emsp;|Mail subsystem                    |
    |mark      &emsp;|Timestamp                         |
    |news      &emsp;|USENET subsystem                  |
    |syslog    &emsp;|Authentication and authorization  |
    |user      &emsp;|Application processes             |
    |uucp      &emsp;|UUCP subsystem                    |
    |local0    &emsp;|Local facility 0                  |
    |local1    &emsp;|Local facility 1                  |
    |local2    &emsp;|Local facility 2                  |
    |local3    &emsp;|Local facility 3                  |
    |local4    &emsp;|Local facility 4                  |
    |local5    &emsp;|Local facility 5                  |
    |local6    &emsp;|Local facility 6                  |
    |local7    &emsp;|Local facility 7                  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
