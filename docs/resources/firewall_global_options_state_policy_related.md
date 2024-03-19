---
page_title: "vyos_firewall_global_options_state_policy_related Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Firewall
  ⯯
  Global Options
  ⯯
  Global firewall state-policy
  ⯯
  Global firewall policy for packets part of a related connection
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

## Schema

### Optional

- `action` (String) Action for packets

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|accept  &emsp;|Action to accept  |
    &emsp;|drop    &emsp;|Action to drop    |
    &emsp;|reject  &emsp;|Action to reject  |
- `log` (Boolean) Log packets hitting this rule
- `log_level` (String) Set log-level. Log must be enable.

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|emerg   &emsp;|Emerg log level     |
    &emsp;|alert   &emsp;|Alert log level     |
    &emsp;|crit    &emsp;|Critical log level  |
    &emsp;|err     &emsp;|Error log level     |
    &emsp;|warn    &emsp;|Warning log level   |
    &emsp;|notice  &emsp;|Notice log level    |
    &emsp;|info    &emsp;|Info log level      |
    &emsp;|debug   &emsp;|Debug log level     |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
