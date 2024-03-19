---
page_title: "vyos_vrf_name_protocols_static_route6_next_hop Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  Routing protocol parameters
  ⯯
  Static Routing
  ⯯
  Static IPv6 route
  ⯯
  IPv6 gateway address
---

# vyos_vrf_name_protocols_static_route6_next_hop (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Static Routing
⯯
Static IPv6 route
⯯
**IPv6 gateway address**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `next_hop_id` (String) IPv6 gateway address

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|ipv6    &emsp;|Next-hop IPv6 router  |
- `route6_id` (String) Static IPv6 route

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv6net  &emsp;|IPv6 static route  |

### Optional

- `bfd` (Attributes) BFD monitoring (see [below for nested schema](#nestedatt--bfd))
- `disable` (Boolean) Disable instance
- `distance` (Number) Distance for this route

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Distance for this route  |
- `interface` (String) Gateway interface name

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Gateway interface name  |
- `segments` (String) SRv6 segments

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Segs (SIDs)  |
- `vrf` (String) VRF to leak route

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Name of VRF to leak to  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--bfd&#34;&gt;&lt;/a&gt;
### Nested Schema for `bfd`

Optional:

- `multi_hop` (Attributes) Use BFD multi hop session (see [below for nested schema](#nestedatt--bfd--multi_hop))
- `profile` (String) Use settings from BFD profile

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|BFD profile name  |

&lt;a id=&#34;nestedatt--bfd--multi_hop&#34;&gt;&lt;/a&gt;
### Nested Schema for `bfd.multi_hop`  &emsp;|
