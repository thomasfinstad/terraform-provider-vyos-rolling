---
page_title: "vyos_service_event_handler_event Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Service event handler⯯Event handler name
---

# vyos_service_event_handler_event (Resource)
<center>

*service*  
⯯  
Service event handler  
⯯  
**Event handler name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `filter` (Attributes) Logs filter settings (see [below for nested schema](#nestedatt--filter))
- `script` (Attributes) Event handler script file (see [below for nested schema](#nestedatt--script))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `event` (String) Event handler name


&lt;a id=&#34;nestedatt--filter&#34;&gt;&lt;/a&gt;
### Nested Schema for `filter`

Optional:

- `pattern` (String) Match pattern (regex)
- `syslog_identifier` (String) Identifier of a process in syslog (string)


&lt;a id=&#34;nestedatt--script&#34;&gt;&lt;/a&gt;
### Nested Schema for `script`

Optional:

- `arguments` (String) Script arguments
- `path` (String) Path to the script


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
