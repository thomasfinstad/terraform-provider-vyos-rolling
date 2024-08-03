---
page_title: "vyos_protocols_bgp_address_family_l2vpn_evpn Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Border Gateway Protocol (BGP)⯯BGP address-family parameters⯯L2VPN EVPN BGP settings
---

# vyos_protocols_bgp_address_family_l2vpn_evpn (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP address-family parameters  
⯯  
**L2VPN EVPN BGP settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `advertise_all_vni` (Boolean) Advertise All local VNIs
- `advertise_default_gw` (Boolean) Advertise All default g/w mac-ip routes in EVPN
- `advertise_pip` (String) EVPN system primary IP

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|ipv4    &emsp;|IP address   |
- `advertise_svi_ip` (Boolean) Advertise svi mac-ip routes in EVPN
- `disable_ead_evi_rx` (Boolean) Activate PE on EAD-ES even if EAD-EVI is not received
- `disable_ead_evi_tx` (Boolean) Do not advertise EAD-EVI for local ESs
- `rd` (String) Route Distinguisher

    &emsp;|Format                   &emsp;|Description                                   |
    |---------------------------|------------------------------------------------|
    &emsp;|ASN:NN_OR_IP-ADDRESS:NN  &emsp;|Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
- `rt_auto_derive` (Boolean) Auto derivation of Route Target (RFC8365)
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
