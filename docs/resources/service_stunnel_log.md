---
page_title: "vyos_service_stunnel_log Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  System services⯯Stunnel TLS Proxy⯯Service logging
---

# vyos_service_stunnel_log (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

System services  
⯯  
Stunnel TLS Proxy  
⯯  
**Service logging**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `level` (String) Specifies log level.

    &emsp;|Format   &emsp;|Description         |
    |-----------|----------------------|
    &emsp;|emerg    &emsp;|Emerg log level     |
    &emsp;|alert    &emsp;|Alert log level     |
    &emsp;|crit     &emsp;|Critical log level  |
    &emsp;|err      &emsp;|Error log level     |
    &emsp;|warning  &emsp;|Warning log level   |
    &emsp;|notice   &emsp;|Notice log level    |
    &emsp;|info     &emsp;|Info log level      |
    &emsp;|debug    &emsp;|Debug log level     |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
