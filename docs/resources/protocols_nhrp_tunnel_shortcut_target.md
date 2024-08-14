---
page_title: "vyos_protocols_nhrp_tunnel_shortcut_target Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Next Hop Resolution Protocol (NHRP) parameters⯯Tunnel for NHRP⯯Defines an off-NBMA network prefix for which the GRE interface will act as a gateway
---

# vyos_protocols_nhrp_tunnel_shortcut_target (Resource)
<center>

*protocols*  
⯯  
Next Hop Resolution Protocol (NHRP) parameters  
⯯  
Tunnel for NHRP  
⯯  
**Defines an off-NBMA network prefix for which the GRE interface will act as a gateway**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `holding_time` (String) Holding time in seconds
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `shortcut_target` (String) Defines an off-NBMA network prefix for which the GRE interface will act as a gateway
- `tunnel` (String) Tunnel for NHRP

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |tunN    &emsp;|NHRP tunnel name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
