---
page_title: "vyos_load_balancing_wan Resource - vyos"

subcategory: "Load"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Configure load-balancing⯯Configure Wide Area Network (WAN) load-balancing
---

# vyos_load_balancing_wan (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Configure load-balancing  
⯯  
**Configure Wide Area Network (WAN) load-balancing**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `disable_source_nat` (Boolean) Disable source NAT rules from being configured for WAN load balancing
- `enable_local_traffic` (Boolean) Enable WAN load balancing for locally sourced traffic
- `flush_connections` (Boolean) Flush connection tracking tables on connection state change
- `hook` (String) Script to be executed on interface status change

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|txt     &emsp;|Script in /config/scripts  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
