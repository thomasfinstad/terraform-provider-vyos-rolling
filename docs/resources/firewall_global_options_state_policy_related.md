---
page_title: "vyos_firewall_global_options_state_policy_related Resource - vyos"

subcategory: "Firewall"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Firewall⯯Global Options⯯Global firewall state-policy⯯Global firewall policy for packets part of a related connection
---

# vyos_firewall_global_options_state_policy_related (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Firewall  
⯯  
Global Options  
⯯  
Global firewall state-policy  
⯯  
**Global firewall policy for packets part of a related connection**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `action` (String) Action for packets

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |accept  &emsp;|Action to accept  |
    |drop    &emsp;|Action to drop    |
    |reject  &emsp;|Action to reject  |
- `log` (Boolean) Log packets hitting this rule
- `log_level` (String) Set log-level. Log must be enable.

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |emerg   &emsp;|Emerg log level     |
    |alert   &emsp;|Alert log level     |
    |crit    &emsp;|Critical log level  |
    |err     &emsp;|Error log level     |
    |warn    &emsp;|Warning log level   |
    |notice  &emsp;|Notice log level    |
    |info    &emsp;|Info log level      |
    |debug   &emsp;|Debug log level     |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
