---
page_title: "vyos_firewall_global_options_state_policy_established Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Firewall
  ⯯
  Global Options
  ⯯
  Global firewall state-policy
  ⯯
  Global firewall policy for packets part of an established connection
---

# vyos_firewall_global_options_state_policy_established (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

	Firewall
⯯
Global Options
⯯
Global firewall state-policy
⯯
**Global firewall policy for packets part of an established connection**


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
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
