---
page_title: "vyos_vpn_ipsec_profile Resource - vyos"

subcategory: "Vpn"

description: |- 
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯VPN IPsec profile
---

# vyos_vpn_ipsec_profile (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
**VPN IPsec profile**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `authentication` (Attributes) Authentication (see [below for nested schema](#nestedatt--authentication))
- `bind` (Attributes) DMVPN tunnel configuration (see [below for nested schema](#nestedatt--bind))
- `disable` (Boolean) Disable instance
- `esp_group` (String) Encapsulating Security Payloads (ESP) group name
- `ike_group` (String) Internet Key Exchange (IKE) group name
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `profile` (String) VPN IPsec profile

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|txt     &emsp;|Profile name  |


&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `mode` (String) Authentication mode

    &emsp;|Format             &emsp;|Description                  |
    |---------------------|-------------------------------|
    &emsp;|pre-shared-secret  &emsp;|Use a pre-shared secret key  |
- `pre_shared_secret` (String) Pre-shared secret key

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|txt     &emsp;|Pre-shared secret key  |


&lt;a id=&#34;nestedatt--bind&#34;&gt;&lt;/a&gt;
### Nested Schema for `bind`

Optional:

- `tunnel` (List of String) Tunnel interface associated with this profile

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|txt     &emsp;|Associated interface to this profile  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
