---
page_title: "vyos_protocols_babel_redistribute_ipv4 Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Babel Routing Protocol⯯Redistribute information from another routing protocol⯯Redistribute IPv4 routes
---

# vyos_protocols_babel_redistribute_ipv4 (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Babel Routing Protocol  
⯯  
Redistribute information from another routing protocol  
⯯  
**Redistribute IPv4 routes**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `bgp` (Boolean) Redistribute BGP routes
- `connected` (Boolean) Redistribute connected routes
- `eigrp` (Boolean) Redistribute EIGRP routes
- `isis` (Boolean) Redistribute IS-IS routes
- `kernel` (Boolean) Redistribute kernel routes
- `nhrp` (Boolean) Redistribute NHRP routes
- `ospf` (Boolean) Redistribute OSPF routes
- `rip` (Boolean) Redistribute RIP routes
- `static` (Boolean) Redistribute static routes
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
