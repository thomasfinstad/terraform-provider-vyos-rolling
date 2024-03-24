---
page_title: "vyos_vrf_name_protocols_ospf_area_virtual_link Resource - terraform-provider-vyos"
subcategory: "vrf"
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

## Schema

### Required

- `area_id` (String) OSPF area settings

    &emsp;|Format  &emsp;|Description                                  |
    |----------|-----------------------------------------------|
    &emsp;|u32     &emsp;|OSPF area number in decimal notation         |
    &emsp;|ipv4    &emsp;|OSPF area number in dotted decimal notation  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `virtual_link_id` (String) Virtual link

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|ipv4    &emsp;|OSPF area in dotted decimal notation  |

### Optional

- `authentication` (Attributes) Authentication (see [below for nested schema](#nestedatt--authentication))
- `dead_interval` (Number) Interval after which a neighbor is declared dead

    &emsp;|Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    &emsp;|1-65535  &emsp;|Neighbor dead interval (seconds)  |
- `hello_interval` (Number) Interval between hello packets

    &emsp;|Format   &emsp;|Description               |
    |-----------|----------------------------|
    &emsp;|1-65535  &emsp;|Hello interval (seconds)  |
- `retransmit_interval` (Number) Interval between retransmitting lost link state advertisements

    &emsp;|Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    &emsp;|1-65535  &emsp;|Retransmit interval (seconds)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `transmit_delay` (Number) Link state transmit delay

    &emsp;|Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    &emsp;|1-65535  &emsp;|Link state transmit delay (seconds)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `md5` (Attributes) MD5 key id (see [below for nested schema](#nestedatt--authentication--md5))
- `plaintext_password` (String) Plain text password

    &emsp;|Format  &emsp;|Description                                 |
    |----------|----------------------------------------------|
    &emsp;|txt     &emsp;|Plain text password (8 characters or less)  |

&lt;a id=&#34;nestedatt--authentication--md5&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication.md5`



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
