---
page_title: "vyos_vrf_name_protocols_bgp_interface Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  Routing protocol parameters
  ⯯
  Border Gateway Protocol (BGP)
  ⯯
  Configure interface related parameters, e.g. MPLS
---

# vyos_vrf_name_protocols_bgp_interface (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Border Gateway Protocol (BGP)
⯯
**Configure interface related parameters, e.g. MPLS**


</center>

## Schema

### Required

- `interface_id` (String) Configure interface related parameters, e.g. MPLS

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Optional

- `mpls` (Attributes) MPLS options (see [below for nested schema](#nestedatt--mpls))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--mpls&#34;&gt;&lt;/a&gt;
### Nested Schema for `mpls`

Optional:

- `forwarding` (Boolean) Enable MPLS forwarding for eBGP directly connected peers  &emsp;|
