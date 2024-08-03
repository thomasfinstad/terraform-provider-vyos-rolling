---
page_title: "vyos_protocols_ospf_neighbor Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Open Shortest Path First (OSPF)⯯Specify neighbor router
---

# vyos_protocols_ospf_neighbor (Resource)
<center>

*protocols*  
⯯  
Open Shortest Path First (OSPF)  
⯯  
**Specify neighbor router**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `poll_interval` (Number) Dead neighbor polling interval

    &emsp;|Format   &emsp;|Description                                     |
    |-----------|--------------------------------------------------|
    &emsp;|1-65535  &emsp;|Seconds between dead neighbor polling interval  |
- `priority` (Number) Neighbor priority in seconds

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|0-255   &emsp;|Neighbor priority  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `neighbor` (String) Specify neighbor router

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|ipv4    &emsp;|Neighbor IP address  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
