---
page_title: "vyos_protocols_ospf_graceful_restart_helper Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Open Shortest Path First (OSPF)⯯Graceful Restart⯯OSPF graceful-restart helpers
---

# vyos_protocols_ospf_graceful_restart_helper (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Open Shortest Path First (OSPF)  
⯯  
Graceful Restart  
⯯  
**OSPF graceful-restart helpers**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `no_strict_lsa_checking` (Boolean) Disable strict LSA check
- `planned_only` (Boolean) Supported only planned restart
- `supported_grace_time` (Number) Supported grace timer

    &emsp;|Format   &emsp;|Description                |
    |-----------|-----------------------------|
    &emsp;|10-1800  &emsp;|Grace interval in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
