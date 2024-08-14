---
page_title: "vyos_protocols_bgp_address_family_l2vpn_evpn_route_target Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Border Gateway Protocol (BGP)⯯BGP address-family parameters⯯L2VPN EVPN BGP settings⯯Route Target
---

# vyos_protocols_bgp_address_family_l2vpn_evpn_route_target (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP address-family parameters  
⯯  
L2VPN EVPN BGP settings  
⯯  
**Route Target**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `both` (List of String) Route Target both import and export

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|txt     &emsp;|Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
- `export` (List of String) Route Target export

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|txt     &emsp;|Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
- `import` (List of String) Route Target import

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|txt     &emsp;|Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  