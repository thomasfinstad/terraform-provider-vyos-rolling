---
page_title: "vyos_protocols_isis_segment_routing_prefix Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Intermediate System to Intermediate System (IS-IS)⯯Segment-Routing (SPRING) settings⯯Static IPv4/IPv6 prefix segment/label mapping
---

# vyos_protocols_isis_segment_routing_prefix (Resource)
<center>

*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
Segment-Routing (SPRING) settings  
⯯  
**Static IPv4/IPv6 prefix segment/label mapping**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `absolute` (Attributes) Specify the absolute value of prefix segment/label ID (see [below for nested schema](#nestedatt--absolute))
- `index` (Attributes) Specify the index value of prefix segment/label ID (see [below for nested schema](#nestedatt--index))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `prefix` (String) Static IPv4/IPv6 prefix segment/label mapping

    |Format   &emsp;|Description          |
    |-----------|-----------------------|
    |ipv4net  &emsp;|IPv4 prefix segment  |
    |ipv6net  &emsp;|IPv6 prefix segment  |


<a id="nestedatt--absolute"></a>
### Nested Schema for `absolute`

Optional:

- `explicit_null` (Boolean) Request upstream neighbor to replace segment/label with explicit null label
- `no_php_flag` (Boolean) Do not request penultimate hop popping for segment/label
- `value` (Number) Specify the absolute value of prefix segment/label ID

    |Format      &emsp;|Description                          |
    |--------------|---------------------------------------|
    |16-1048575  &emsp;|The absolute segment/label ID value  |


<a id="nestedatt--index"></a>
### Nested Schema for `index`

Optional:

- `explicit_null` (Boolean) Request upstream neighbor to replace segment/label with explicit null label
- `no_php_flag` (Boolean) Do not request penultimate hop popping for segment/label
- `value` (Number) Specify the index value of prefix segment/label ID

    |Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    |0-65535  &emsp;|The index segment/label ID value  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
