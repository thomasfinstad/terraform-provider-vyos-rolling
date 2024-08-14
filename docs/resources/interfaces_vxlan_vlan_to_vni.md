---
page_title: "vyos_interfaces_vxlan_vlan_to_vni Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Virtual Extensible LAN (VXLAN) Interface⯯Configuring VLAN-to-VNI mappings for EVPN-VXLAN
---

# vyos_interfaces_vxlan_vlan_to_vni (Resource)
<center>

*interfaces*  
⯯  
Virtual Extensible LAN (VXLAN) Interface  
⯯  
**Configuring VLAN-to-VNI mappings for EVPN-VXLAN**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vni` (String) Virtual Network Identifier

    &emsp;|Format       &emsp;|Description                                             |
    |---------------|----------------------------------------------------------|
    &emsp;|0-16777214   &emsp;|VXLAN virtual network identifier                        |
    &emsp;|&lt;start-end&gt;  &emsp;|VXLAN virtual network IDs range (use &#39;-&#39; as delimiter)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `vlan_to_vni` (String) Configuring VLAN-to-VNI mappings for EVPN-VXLAN

    &emsp;|Format       &emsp;|Description                            |
    |---------------|-----------------------------------------|
    &emsp;|0-4094       &emsp;|Virtual Local Area Network (VLAN) ID   |
    &emsp;|&lt;start-end&gt;  &emsp;|VLAN IDs range (use &#39;-&#39; as delimiter)  |
- `vxlan` (String) Virtual Extensible LAN (VXLAN) Interface

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|vxlanN  &emsp;|VXLAN interface name  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  