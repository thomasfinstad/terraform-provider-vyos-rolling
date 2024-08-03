---
page_title: "vyos_system_frr_snmp Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯Configure FRRouting parameters⯯Enable SNMP integration for next daemons
---

# vyos_system_frr_snmp (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
Configure FRRouting parameters  
⯯  
**Enable SNMP integration for next daemons**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `bgpd` (Boolean) BGP
- `isisd` (Boolean) IS-IS
- `ldpd` (Boolean) LDP
- `ospf6d` (Boolean) OSPFv3
- `ospfd` (Boolean) OSPFv2
- `ripd` (Boolean) RIP
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `zebra` (Boolean) Zebra (IP routing manager)

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
