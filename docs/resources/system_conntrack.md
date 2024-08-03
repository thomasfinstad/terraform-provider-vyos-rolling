---
page_title: "vyos_system_conntrack Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯Connection Tracking Engine Options
---

# vyos_system_conntrack (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
**Connection Tracking Engine Options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `expect_table_size` (Number) Size of connection tracking expect table

    &emsp;|Format      &emsp;|Description                                                    |
    |--------------|-----------------------------------------------------------------|
    &emsp;|1-50000000  &emsp;|Number of entries allowed in connection tracking expect table  |
- `flow_accounting` (Boolean) Enable connection tracking flow accounting
- `hash_size` (Number) Hash size for connection tracking table

    &emsp;|Format      &emsp;|Description                                        |
    |--------------|-----------------------------------------------------|
    &emsp;|1-50000000  &emsp;|Size of hash to use for connection tracking table  |
- `table_size` (Number) Size of connection tracking table

    &emsp;|Format      &emsp;|Description                                             |
    |--------------|----------------------------------------------------------|
    &emsp;|1-50000000  &emsp;|Number of entries allowed in connection tracking table  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
