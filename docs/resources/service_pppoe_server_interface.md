---
page_title: "vyos_service_pppoe_server_interface Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Point to Point over Ethernet (PPPoE) Server⯯interface(s) to listen on
---

# vyos_service_pppoe_server_interface (Resource)
<center>

*service*  
⯯  
Point to Point over Ethernet (PPPoE) Server  
⯯  
**interface(s) to listen on**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vlan` (List of String) VLAN monitor for automatic creation of VLAN interfaces

    &emsp;|Format     &emsp;|Description                                      |
    |-------------|---------------------------------------------------|
    &emsp;|1-4094     &emsp;|VLAN for automatic creation                      |
    &emsp;|start-end  &emsp;|VLAN range for automatic creation (e.g. 1-4094)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) interface(s) to listen on


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
