---
page_title: "vyos_system_conntrack_log_event_update Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯Connection Tracking Engine Options⯯Log connection tracking⯯Event type and protocol⯯Log connection updates
---

# vyos_system_conntrack_log_event_update (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
Connection Tracking Engine Options  
⯯  
Log connection tracking  
⯯  
Event type and protocol  
⯯  
**Log connection updates**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `icmp` (Boolean) Log connection tracking events for ICMP
- `other` (Boolean) Log connection tracking events for all protocols other than TCP, UDP and ICMP
- `tcp` (Boolean) Log connection tracking events for TCP
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `udp` (Boolean) Log connection tracking events for UDP

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
