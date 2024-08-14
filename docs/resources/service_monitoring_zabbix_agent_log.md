---
page_title: "vyos_service_monitoring_zabbix_agent_log Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯monitoring⯯Zabbix-agent settings⯯Log settings
---

# vyos_service_monitoring_zabbix_agent_log (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
*monitoring*  
⯯  
Zabbix-agent settings  
⯯  
**Log settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `debug_level` (String) Debug level

    |Format          &emsp;|Description                 |
    |------------------|------------------------------|
    |basic           &emsp;|Basic information           |
    |critical        &emsp;|Critical information        |
    |error           &emsp;|Error information           |
    |warning         &emsp;|Warnings                    |
    |debug           &emsp;|Debug information           |
    |extended-debug  &emsp;|Extended debug information  |
- `remote_commands` (Boolean) Enable logging of executed shell commands as warnings
- `size` (Number) Log file size in megabytes

    |Format  &emsp;|Description  |
    |----------|---------------|
    |0-1024  &emsp;|Megabytes    |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
