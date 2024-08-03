---
page_title: "vyos_interfaces_bridge_member_interface Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Bridge Interface⯯Bridge member interfaces⯯Member interface name
---

# vyos_interfaces_bridge_member_interface (Resource)
<center>

*interfaces*  
⯯  
Bridge Interface  
⯯  
Bridge member interfaces  
⯯  
**Member interface name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `allowed_vlan` (List of String) Specify VLAN id which is allowed in this trunk interface

    &emsp;|Format       &emsp;|Description                                                     |
    |---------------|------------------------------------------------------------------|
    &emsp;|&lt;id&gt;         &emsp;|VLAN id allowed to pass this interface                          |
    &emsp;|&lt;idN&gt;-&lt;idM&gt;  &emsp;|VLAN id range allowed on this interface (use &#39;-&#39; as delimiter)  |
- `cost` (Number) Bridge port cost

    &emsp;|Format   &emsp;|Description                                 |
    |-----------|----------------------------------------------|
    &emsp;|1-65535  &emsp;|Path cost value for Spanning Tree Protocol  |
- `isolated` (Boolean) Port is isolated (also known as Private-VLAN)
- `native_vlan` (Number) Specify VLAN id which should natively be present on the link

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|1-4094  &emsp;|Virtual Local Area Network (VLAN) ID  |
- `priority` (Number) Bridge port priority

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|0-63    &emsp;|Bridge port priority  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `bridge` (String) Bridge Interface

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|brN     &emsp;|Bridge interface name  |
- `interface` (String) Member interface name


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
