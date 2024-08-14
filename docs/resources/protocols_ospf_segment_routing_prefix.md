---
page_title: "vyos_protocols_ospf_segment_routing_prefix Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Open Shortest Path First (OSPF)⯯Segment-Routing (SPRING) settings⯯Static IPv4 prefix segment/label mapping
---

# vyos_protocols_ospf_segment_routing_prefix (Resource)
<center>

*protocols*  
⯯  
Open Shortest Path First (OSPF)  
⯯  
Segment-Routing (SPRING) settings  
⯯  
**Static IPv4 prefix segment/label mapping**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `index` (Attributes) Specify the index value of prefix segment/label ID (see [below for nested schema](#nestedatt--index))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `prefix` (String) Static IPv4 prefix segment/label mapping

    &emsp;|Format   &emsp;|Description          |
    |-----------|-----------------------|
    &emsp;|ipv4net  &emsp;|IPv4 prefix segment  |


&lt;a id=&#34;nestedatt--index&#34;&gt;&lt;/a&gt;
### Nested Schema for `index`

Optional:

- `explicit_null` (Boolean) Request upstream neighbor to replace segment/label with explicit null label
- `no_php_flag` (Boolean) Do not request penultimate hop popping for segment/label
- `value` (Number) Specify the index value of prefix segment/label ID

    &emsp;|Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    &emsp;|0-65535  &emsp;|The index segment/label ID value  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  