---
page_title: "vyos_service_pppoe_server_extended_scripts Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Point to Point over Ethernet (PPPoE) Server⯯Extended script execution
---

# vyos_service_pppoe_server_extended_scripts (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Point to Point over Ethernet (PPPoE) Server  
⯯  
**Extended script execution**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `on_change` (String) Script to run when session interface changed by RADIUS CoA handling
- `on_down` (String) Script to run when session interface going to terminate
- `on_pre_up` (String) Script to run before session interface comes up
- `on_up` (String) Script to run when session interface is completely configured and started
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  