---
page_title: "vyos_firewall_global_options_apply_to_bridged_traffic Resource - vyos"

subcategory: "Firewall"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Firewall⯯Global Options⯯Apply configured firewall rules to traffic switched by bridges
---

# vyos_firewall_global_options_apply_to_bridged_traffic (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Firewall  
⯯  
Global Options  
⯯  
**Apply configured firewall rules to traffic switched by bridges**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `ipv4` (Boolean) Apply configured IPv4 firewall rules
- `ipv6` (Boolean) Apply configured IPv6 firewall rules
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
