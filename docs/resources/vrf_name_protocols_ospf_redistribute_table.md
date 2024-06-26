---
page_title: "vyos_vrf_name_protocols_ospf_redistribute_table Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Open Shortest Path First (OSPF)⯯Redistribute information from another routing protocol⯯Redistribute non-main Kernel Routing Table
---

# vyos_vrf_name_protocols_ospf_redistribute_table (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Open Shortest Path First (OSPF)  
⯯  
Redistribute information from another routing protocol  
⯯  
**Redistribute non-main Kernel Routing Table**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `table_id` (Number) Redistribute non-main Kernel Routing Table

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|1-200   &emsp;|Policy route table number  |

### Optional

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
