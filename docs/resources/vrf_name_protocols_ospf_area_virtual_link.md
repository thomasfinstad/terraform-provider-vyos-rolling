---
page_title: "vyos_vrf_name_protocols_ospf_area_virtual_link Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Open Shortest Path First (OSPF)⯯OSPF area settings⯯Virtual link
---

# vyos_vrf_name_protocols_ospf_area_virtual_link (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Open Shortest Path First (OSPF)  
⯯  
OSPF area settings  
⯯  
**Virtual link**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `authentication` (Attributes) Authentication (see [below for nested schema](#nestedatt--authentication))
- `dead_interval` (Number) Interval after which a neighbor is declared dead

    |Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    |1-65535  &emsp;|Neighbor dead interval (seconds)  |
- `hello_interval` (Number) Interval between hello packets

    |Format   &emsp;|Description               |
    |-----------|----------------------------|
    |1-65535  &emsp;|Hello interval (seconds)  |
- `retransmit_interval` (Number) Interval between retransmitting lost link state advertisements

    |Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    |1-65535  &emsp;|Retransmit interval (seconds)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `transmit_delay` (Number) Link state transmit delay

    |Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    |1-65535  &emsp;|Link state transmit delay (seconds)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `area` (String) OSPF area settings

    |Format  &emsp;|Description                                  |
    |----------|-----------------------------------------------|
    |u32     &emsp;|OSPF area number in decimal notation         |
    |ipv4    &emsp;|OSPF area number in dotted decimal notation  |
- `name` (String) Virtual Routing and Forwarding instance

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |
- `virtual_link` (String) Virtual link

    |Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    |ipv4    &emsp;|OSPF area in dotted decimal notation  |


<a id="nestedatt--authentication"></a>
### Nested Schema for `authentication`

Optional:

- `md5` (Attributes) MD5 key id (see [below for nested schema](#nestedatt--authentication--md5))
- `plaintext_password` (String) Plain text password

    |Format  &emsp;|Description                                 |
    |----------|----------------------------------------------|
    |txt     &emsp;|Plain text password (8 characters or less)  |

<a id="nestedatt--authentication--md5"></a>
### Nested Schema for `authentication.md5`



<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
