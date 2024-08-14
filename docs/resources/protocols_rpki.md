---
page_title: "vyos_protocols_rpki Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Resource Public Key Infrastructure (RPKI)
---

# vyos_protocols_rpki (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
**Resource Public Key Infrastructure (RPKI)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `expire_interval` (Number) Interval to wait before expiring the cache

    &emsp;|Format      &emsp;|Description          |
    |--------------|-----------------------|
    &emsp;|600-172800  &emsp;|Interval in seconds  |
- `polling_period` (Number) Cache polling interval

    &emsp;|Format   &emsp;|Description          |
    |-----------|-----------------------|
    &emsp;|1-86400  &emsp;|Interval in seconds  |
- `retry_interval` (Number) Retry interval to connect to the cache server

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|1-7200  &emsp;|Interval in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  