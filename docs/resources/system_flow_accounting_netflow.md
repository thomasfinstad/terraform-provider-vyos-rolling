---
page_title: "vyos_system_flow_accounting_netflow Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯Flow accounting settings⯯NetFlow settings
---

# vyos_system_flow_accounting_netflow (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
Flow accounting settings  
⯯  
**NetFlow settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `engine_id` (String) NetFlow engine-id

    &emsp;|Format                &emsp;|Description                       |
    |------------------------|------------------------------------|
    &emsp;|0-255 or 0-255:0-255  &emsp;|NetFlow engine-id for v5          |
    &emsp;|u32                   &emsp;|NetFlow engine-id for v9 / IPFIX  |
- `max_flows` (Number) NetFlow maximum flows

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|u32     &emsp;|NetFlow maximum flows  |
- `sampling_rate` (Number) NetFlow sampling-rate

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|u32     &emsp;|Sampling rate (1 in N packets)  |
- `source_address` (String) Source IP address used to initiate connection

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|ipv4    &emsp;|IPv4 source address  |
    &emsp;|ipv6    &emsp;|IPv6 source address  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `version` (String) NetFlow version to export

    &emsp;|Format  &emsp;|Description                                        |
    |----------|-----------------------------------------------------|
    &emsp;|5       &emsp;|NetFlow version 5                                  |
    &emsp;|9       &emsp;|NetFlow version 9                                  |
    &emsp;|10      &emsp;|Internet Protocol Flow Information Export (IPFIX)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  