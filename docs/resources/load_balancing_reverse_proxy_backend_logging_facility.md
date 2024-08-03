---
page_title: "vyos_load_balancing_reverse_proxy_backend_logging_facility Resource - vyos"

subcategory: "Load"

description: |- 
  load-balancing⯯Configure reverse-proxy⯯Backend server name⯯Logging parameters⯯Facility for logging
---

# vyos_load_balancing_reverse_proxy_backend_logging_facility (Resource)
<center>

*load-balancing*  
⯯  
Configure reverse-proxy  
⯯  
Backend server name  
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

    &emsp;|Format   &emsp;|Description                         |
    |-----------|--------------------------------------|
    &emsp;|emerg    &emsp;|Emergency messages                  |
    &emsp;|alert    &emsp;|Urgent messages                     |
    &emsp;|crit     &emsp;|Critical messages                   |
    &emsp;|err      &emsp;|Error messages                      |
    &emsp;|warning  &emsp;|Warning messages                    |
    &emsp;|notice   &emsp;|Messages for further investigation  |
    &emsp;|info     &emsp;|Informational messages              |
    &emsp;|debug    &emsp;|Debug messages                      |
    &emsp;|all      &emsp;|Log everything                      |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `backend` (String) Backend server name
- `facility` (String) Facility for logging

    &emsp;|Format    &emsp;|Description                       |
    |------------|------------------------------------|
    &emsp;|all       &emsp;|All facilities excluding &#34;mark&#34;   |
    &emsp;|auth      &emsp;|Authentication and authorization  |
    &emsp;|authpriv  &emsp;|Non-system authorization          |
    &emsp;|cron      &emsp;|Cron daemon                       |
    &emsp;|daemon    &emsp;|System daemons                    |
    &emsp;|kern      &emsp;|Kernel                            |
    &emsp;|lpr       &emsp;|Line printer spooler              |
    &emsp;|mail      &emsp;|Mail subsystem                    |
    &emsp;|mark      &emsp;|Timestamp                         |
    &emsp;|news      &emsp;|USENET subsystem                  |
    &emsp;|syslog    &emsp;|Authentication and authorization  |
    &emsp;|user      &emsp;|Application processes             |
    &emsp;|uucp      &emsp;|UUCP subsystem                    |
    &emsp;|local0    &emsp;|Local facility 0                  |
    &emsp;|local1    &emsp;|Local facility 1                  |
    &emsp;|local2    &emsp;|Local facility 2                  |
    &emsp;|local3    &emsp;|Local facility 3                  |
    &emsp;|local4    &emsp;|Local facility 4                  |
    &emsp;|local5    &emsp;|Local facility 5                  |
    &emsp;|local6    &emsp;|Local facility 6                  |
    &emsp;|local7    &emsp;|Local facility 7                  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
