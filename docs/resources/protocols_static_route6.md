---
page_title: "vyos_protocols_static_route6 Resource - vyos"

subcategory: "Protocols"

description: |- 
  Routing protocols⯯Static Routing⯯Static IPv6 route
---

# vyos_protocols_static_route6 (Resource)
<center>

Routing protocols  
⯯  
Static Routing  
⯯  
**Static IPv6 route**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `blackhole` (Attributes) Silently discard pkts when matched (see [below for nested schema](#nestedatt--blackhole))
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `reject` (Attributes) Emit an ICMP unreachable when matched (see [below for nested schema](#nestedatt--reject))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `route6` (String) Static IPv6 route

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv6net  &emsp;|IPv6 static route  |


&lt;a id=&#34;nestedatt--blackhole&#34;&gt;&lt;/a&gt;
### Nested Schema for `blackhole`

Optional:

- `distance` (Number) Distance for this route

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Distance for this route  |
- `tag` (Number) Tag value for this route

    &emsp;|Format        &emsp;|Description               |
    |----------------|----------------------------|
    &emsp;|1-4294967295  &emsp;|Tag value for this route  |


&lt;a id=&#34;nestedatt--reject&#34;&gt;&lt;/a&gt;
### Nested Schema for `reject`

Optional:

- `distance` (Number) Distance for this route

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Distance for this route  |
- `tag` (Number) Tag value for this route

    &emsp;|Format        &emsp;|Description               |
    |----------------|----------------------------|
    &emsp;|1-4294967295  &emsp;|Tag value for this route  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
