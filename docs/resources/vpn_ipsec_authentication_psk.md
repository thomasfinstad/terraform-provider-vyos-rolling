---
page_title: "vyos_vpn_ipsec_authentication_psk Resource - vyos"

subcategory: "Vpn"

description: |- 
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯Authentication⯯Pre-shared key name
---

# vyos_vpn_ipsec_authentication_psk (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
Authentication  
⯯  
**Pre-shared key name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `dhcp_interface` (List of String) DHCP interface supplying next-hop IP address

    |Format  &emsp;|Description          |
    |----------|-----------------------|
    |txt     &emsp;|DHCP interface name  |
- `id_param` (List of String) ID for authentication

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |txt     &emsp;|ID used for authentication  |
- `secret` (String) IKE pre-shared secret key

    |Format  &emsp;|Description                |
    |----------|-----------------------------|
    |txt     &emsp;|IKE pre-shared secret key  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `psk` (String) Pre-shared key name


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
