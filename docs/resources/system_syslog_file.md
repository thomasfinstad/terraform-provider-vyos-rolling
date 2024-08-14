---
page_title: "vyos_system_syslog_file Resource - vyos"

subcategory: "System"

description: |- 
  system⯯System logging⯯Logging to a file
---

# vyos_system_syslog_file (Resource)
<center>

*system*  
⯯  
System logging  
⯯  
**Logging to a file**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `archive` (Attributes) Log file size and rotation characteristics (see [below for nested schema](#nestedatt--archive))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `file` (String) Logging to a file


<a id="nestedatt--archive"></a>
### Nested Schema for `archive`

Optional:

- `file` (String) Number of saved files
- `size` (String) Size of log files in kbytes


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
