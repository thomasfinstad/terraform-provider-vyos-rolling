---
page_title: "vyos_service_ipoe_server_log Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Internet Protocol over Ethernet (IPoE) Server⯯**Server logging **
---

# vyos_service_ipoe_server_log (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Internet Protocol over Ethernet (IPoE) Server  
⯯  
**Server logging **


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `level` (String) Specifies log level

    &emsp;|Format  &emsp;|Description                                          |
    |----------|-------------------------------------------------------|
    &emsp;|0       &emsp;|Turn off logging                                     |
    &emsp;|1       &emsp;|Log only error messages                              |
    &emsp;|2       &emsp;|Log error and warning messages                       |
    &emsp;|3       &emsp;|Log error, warning and minimum information messages  |
    &emsp;|4       &emsp;|Log error, warning and full information messages     |
    &emsp;|5       &emsp;|Log all messages including debug messages            |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  