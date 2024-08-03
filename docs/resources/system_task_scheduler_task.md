---
page_title: "vyos_system_task_scheduler_task Resource - vyos"

subcategory: "System"

description: |- 
  system⯯Task scheduler settings⯯Scheduled task
---

# vyos_system_task_scheduler_task (Resource)
<center>

*system*  
⯯  
Task scheduler settings  
⯯  
**Scheduled task**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `crontab_spec` (String) UNIX crontab time specification string
- `executable` (Attributes) Executable path and arguments (see [below for nested schema](#nestedatt--executable))
- `interval` (String) Execution interval

    &emsp;|Format      &emsp;|Description                    |
    |--------------|---------------------------------|
    &emsp;|&lt;minutes&gt;   &emsp;|Execution interval in minutes  |
    &emsp;|&lt;minutes&gt;m  &emsp;|Execution interval in minutes  |
    &emsp;|&lt;hours&gt;h    &emsp;|Execution interval in hours    |
    &emsp;|&lt;days&gt;d     &emsp;|Execution interval in days     |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `task` (String) Scheduled task

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Task name    |


&lt;a id=&#34;nestedatt--executable&#34;&gt;&lt;/a&gt;
### Nested Schema for `executable`

Optional:

- `arguments` (String) Arguments passed to the executable
- `path` (String) Path to executable


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
