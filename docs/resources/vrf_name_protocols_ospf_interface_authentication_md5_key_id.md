---
page_title: "vyos_vrf_name_protocols_ospf_interface_authentication_md5_key_id Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Open Shortest Path First (OSPF)⯯Interface configuration⯯Authentication⯯MD5 key id⯯MD5 key id
---

# vyos_vrf_name_protocols_ospf_interface_authentication_md5_key_id (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Open Shortest Path First (OSPF)  
⯯  
Interface configuration  
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

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `md5_key` (String) MD5 authentication type

    |Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    |txt     &emsp;|MD5 Key (16 characters or less)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface configuration

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |
- `key_id` (Number) MD5 key id

    |Format  &emsp;|Description  |
    |----------|---------------|
    |1-255   &emsp;|MD5 key id   |
- `name` (String) Virtual Routing and Forwarding instance

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
