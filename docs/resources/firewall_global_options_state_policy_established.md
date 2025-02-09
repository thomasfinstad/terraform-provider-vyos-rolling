---
page_title: "vyos_firewall_global_options_state_policy_established Resource - vyos"

subcategory: "Firewall"

description: |-
  ~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.
  Firewall⯯Global Options⯯Global firewall state-policy⯯Global firewall policy for packets part of an established connection
---

# vyos_firewall_global_options_state_policy_established (Resource)
<center>

~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

Firewall  
⯯  
Global Options  
⯯  
Global firewall state-policy  
⯯  
**Global firewall policy for packets part of an established connection**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_firewall_global_options_state_policy_established (Resource)](#vyos_firewall_global_options_state_policy_established-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [action](#action)
      - [log](#log)
      - [log_level](#log_level)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### action
- `action` (String) Action for packets

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  accept  &emsp;|  Action to accept  |
    |  drop    &emsp;|  Action to drop    |
    |  reject  &emsp;|  Action to reject  |
#### log
- `log` (Boolean) Log packets hitting this rule
#### log_level
- `log_level` (String) Set log-level. Log must be enable.

    |  Format  &emsp;|  Description         |
    |----------|----------------------|
    |  emerg   &emsp;|  Emerg log level     |
    |  alert   &emsp;|  Alert log level     |
    |  crit    &emsp;|  Critical log level  |
    |  err     &emsp;|  Error log level     |
    |  warn    &emsp;|  Warning log level   |
    |  notice  &emsp;|  Notice log level    |
    |  info    &emsp;|  Info log level      |
    |  debug   &emsp;|  Debug log level     |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_firewall_global_options_state_policy_established.example "firewall__global-options__state-policy__established"
```
