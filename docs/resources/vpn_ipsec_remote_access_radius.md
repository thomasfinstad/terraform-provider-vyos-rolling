---
page_title: "vyos_vpn_ipsec_remote_access_radius Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯IKEv2 remote access VPN⯯RADIUS based user authentication
---

# vyos_vpn_ipsec_remote_access_radius (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
IKEv2 remote access VPN  
⯯  
**RADIUS based user authentication**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `nas_identifier` (String) NAS-Identifier attribute sent to RADIUS
- `source_address` (String) IPv4 source address used to initiate connection

    |Format  &emsp;|Description          |
    |----------|-----------------------|
    |ipv4    &emsp;|IPv4 source address  |
- `timeout` (Number) Session timeout

    |Format  &emsp;|Description                              |
    |----------|-------------------------------------------|
    |1-240   &emsp;|Session timeout in seconds (default: 2)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
