---
page_title: "vyos_system_conntrack_log Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯Connection Tracking Engine Options⯯Log connection tracking
---

# vyos_system_conntrack_log (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
Connection Tracking Engine Options  
⯯  
**Log connection tracking**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `log_level` (String) Set log-level. Log must be enable.

    |Format  &emsp;|Description      |
    |----------|-------------------|
    |info    &emsp;|Info log level   |
    |debug   &emsp;|Debug log level  |
- `queue_size` (Number) Internal message queue size

    |Format      &emsp;|Description  |
    |--------------|---------------|
    |100-999999  &emsp;|Queue size   |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `timestamp` (Boolean) Log connection tracking events include flow-based timestamp

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
