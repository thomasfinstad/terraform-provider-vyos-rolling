---
page_title: "vyos_protocols_bgp_sid_vpn_per_vrf Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Border Gateway Protocol (BGP)⯯SID value for VRF⯯Between current VRF and VPN⯯SID per-VRF (both IPv4 and IPv6 address families)
---

# vyos_protocols_bgp_sid_vpn_per_vrf (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
SID value for VRF  
⯯  
Between current VRF and VPN  
⯯  
**SID per-VRF (both IPv4 and IPv6 address families)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `export` (String) For routes leaked from current VRF to VPN

    &emsp;|Format     &emsp;|Description                   |
    |-------------|--------------------------------|
    &emsp;|1-1048575  &emsp;|SID allocation index          |
    &emsp;|auto       &emsp;|Automatically assign a label  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  