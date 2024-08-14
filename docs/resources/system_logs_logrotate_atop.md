---
page_title: "vyos_system_logs_logrotate_atop Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯Logging options⯯Logrotate options⯯Atop logs options (system resources usage)
---

# vyos_system_logs_logrotate_atop (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
Logging options  
⯯  
Logrotate options  
⯯  
**Atop logs options (system resources usage)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `max_size` (Number) Size of a single log file that triggers rotation

    |Format  &emsp;|Description  |
    |----------|---------------|
    |1-1024  &emsp;|Size in MB   |
- `rotate` (Number) Count of rotations before old logs will be deleted

    |Format  &emsp;|Description  |
    |----------|---------------|
    |1-100   &emsp;|Rotations    |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
