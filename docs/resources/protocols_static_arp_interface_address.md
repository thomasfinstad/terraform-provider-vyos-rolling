---
page_title: "vyos_protocols_static_arp_interface_address Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯static⯯Static ARP translation⯯Interface configuration⯯IP address for static ARP entry
---

# vyos_protocols_static_arp_interface_address (Resource)
<center>

*protocols*  
⯯  
*static*  
⯯  
Static ARP translation  
⯯  
Interface configuration  
⯯  
**IP address for static ARP entry**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `mac` (String) Media Access Control (MAC) address

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|macaddr  &emsp;|Hardware (MAC) address  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `address` (String) IP address for static ARP entry

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|ipv4    &emsp;|IPv4 destination address  |
- `interface` (String) Interface configuration

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
