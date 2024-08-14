---
page_title: "vyos_vpn_pptp_remote_access_authentication_radius_rate_limit Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  vpn⯯Point to Point Tunneling Protocol (PPTP) Virtual Private Network (VPN)⯯Remote access PPTP VPN⯯Authentication for remote access PPTP VPN⯯RADIUS based user authentication⯯Upload/Download speed limits
---

# vyos_vpn_pptp_remote_access_authentication_radius_rate_limit (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*vpn*  
⯯  
Point to Point Tunneling Protocol (PPTP) Virtual Private Network (VPN)  
⯯  
Remote access PPTP VPN  
⯯  
Authentication for remote access PPTP VPN  
⯯  
RADIUS based user authentication  
⯯  
**Upload/Download speed limits**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `attribute` (String) RADIUS attribute that contains rate information
- `enable` (Boolean) Enable bandwidth shaping via RADIUS
- `multiplier` (String) Shaper multiplier

    &emsp;|Format        &emsp;|Description        |
    |----------------|---------------------|
    &emsp;|&lt;0.001-1000&gt;  &emsp;|Shaper multiplier  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vendor` (String) Vendor dictionary

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  