---
page_title: "vyos_vrf_name_protocols_isis_segment_routing_prefix Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  Routing protocol parameters
  ⯯
  Intermediate System to Intermediate System (IS-IS)
  ⯯
  Segment-Routing (SPRING) settings
  ⯯
  Static IPv4/IPv6 prefix segment/label mapping
---

# vyos_vrf_name_protocols_isis_segment_routing_prefix (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Intermediate System to Intermediate System (IS-IS)
⯯
Segment-Routing (SPRING) settings
⯯
**Static IPv4/IPv6 prefix segment/label mapping**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `prefix_id` (String) Static IPv4/IPv6 prefix segment/label mapping

    &emsp;|Format   &emsp;|Description          |
    |-----------|-----------------------|
    &emsp;|ipv4net  &emsp;|IPv4 prefix segment  |
    &emsp;|ipv6net  &emsp;|IPv6 prefix segment  |

### Optional

- `absolute` (Attributes) Specify the absolute value of prefix segment/label ID (see [below for nested schema](#nestedatt--absolute))
- `index` (Attributes) Specify the index value of prefix segment/label ID (see [below for nested schema](#nestedatt--index))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--absolute&#34;&gt;&lt;/a&gt;
### Nested Schema for `absolute`

Optional:

- `explicit_null` (Boolean) Request upstream neighbor to replace segment/label with explicit null label
- `no_php_flag` (Boolean) Do not request penultimate hop popping for segment/label
- `value` (Number) Specify the absolute value of prefix segment/label ID

    &emsp;|Format      &emsp;|Description                          |
    |--------------|---------------------------------------|
    &emsp;|16-1048575  &emsp;|The absolute segment/label ID value  |


&lt;a id=&#34;nestedatt--index&#34;&gt;&lt;/a&gt;
### Nested Schema for `index`

Optional:

- `explicit_null` (Boolean) Request upstream neighbor to replace segment/label with explicit null label
- `no_php_flag` (Boolean) Do not request penultimate hop popping for segment/label
- `value` (Number) Specify the index value of prefix segment/label ID

    &emsp;|Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    &emsp;|0-65535  &emsp;|The index segment/label ID value  |  &emsp;|
