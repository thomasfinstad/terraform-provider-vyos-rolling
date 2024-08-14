---
page_title: "vyos_protocols_segment_routing_srv6_locator Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Segment Routing⯯Segment-Routing SRv6 configuration⯯Segment Routing SRv6 locator
---

# vyos_protocols_segment_routing_srv6_locator (Resource)
<center>

*protocols*  
⯯  
Segment Routing  
⯯  
Segment-Routing SRv6 configuration  
⯯  
**Segment Routing SRv6 locator**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `behavior_usid` (Boolean) Set SRv6 behavior uSID
- `block_len` (Number) Configure SRv6 locator block length in bits

    |Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    |16-64   &emsp;|Specify SRv6 locator block length in bits  |
- `func_bits` (Number) Configure SRv6 locator function length in bits

    |Format  &emsp;|Description                                   |
    |----------|------------------------------------------------|
    |0-64    &emsp;|Specify SRv6 locator function length in bits  |
- `node_len` (Number) Configure SRv6 locator node length in bits

    |Format  &emsp;|Description                                 |
    |----------|----------------------------------------------|
    |16-64   &emsp;|Configure SRv6 locator node length in bits  |
- `prefix` (String) SRv6 locator prefix

    |Format   &emsp;|Description          |
    |-----------|-----------------------|
    |ipv6net  &emsp;|SRv6 locator prefix  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `locator` (String) Segment Routing SRv6 locator


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
