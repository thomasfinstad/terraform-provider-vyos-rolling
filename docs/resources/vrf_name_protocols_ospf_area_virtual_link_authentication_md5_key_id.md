---
page_title: "vyos_vrf_name_protocols_ospf_area_virtual_link_authentication_md5_key_id Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Open Shortest Path First (OSPF)⯯OSPF area settings⯯Virtual link⯯Authentication⯯MD5 key id⯯MD5 key id
---

# vyos_vrf_name_protocols_ospf_area_virtual_link_authentication_md5_key_id (Resource)
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
Virtual link  
⯯  
Authentication  
⯯  
MD5 key id  
⯯  
**MD5 key id**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `area_id` (String) OSPF area settings

    &emsp;|Format  &emsp;|Description                                  |
    |----------|-----------------------------------------------|
    &emsp;|u32     &emsp;|OSPF area number in decimal notation         |
    &emsp;|ipv4    &emsp;|OSPF area number in dotted decimal notation  |
- `key_id_id` (Number) MD5 key id

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1-255   &emsp;|MD5 key id   |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `virtual_link_id` (String) Virtual link

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|ipv4    &emsp;|OSPF area in dotted decimal notation  |

### Optional

- `md5_key` (String) MD5 authentication type

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|MD5 Key (16 characters or less)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
