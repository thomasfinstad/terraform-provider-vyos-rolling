---
page_title: "vyos_vrf_name_protocols_ospf_access_list Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Open Shortest Path First (OSPF)⯯Access list to filter networks in routing updates
---

# vyos_vrf_name_protocols_ospf_access_list (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Open Shortest Path First (OSPF)  
⯯  
**Access list to filter networks in routing updates**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `access_list_id` (Number) Access list to filter networks in routing updates

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|u32     &emsp;|Access-list number  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Optional

- `export` (List of String) Filter for outgoing routing update

    &emsp;|Format     &emsp;|Description              |
    |-------------|---------------------------|
    &emsp;|bgp        &emsp;|Filter BGP routes        |
    &emsp;|connected  &emsp;|Filter connected routes  |
    &emsp;|isis       &emsp;|Filter IS-IS routes      |
    &emsp;|kernel     &emsp;|Filter Kernel routes     |
    &emsp;|rip        &emsp;|Filter RIP routes        |
    &emsp;|static     &emsp;|Filter static routes     |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
