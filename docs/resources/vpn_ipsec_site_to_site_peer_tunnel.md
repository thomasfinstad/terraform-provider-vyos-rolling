---
page_title: "vyos_vpn_ipsec_site_to_site_peer_tunnel Resource - vyos"

subcategory: "Vpn"

description: |- 
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯Site-to-site VPN⯯Connection name of the peer⯯Peer tunnel
---

# vyos_vpn_ipsec_site_to_site_peer_tunnel (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
Site-to-site VPN  
⯯  
Connection name of the peer  
⯯  
**Peer tunnel**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `disable` (Boolean) Disable instance
- `esp_group` (String) Encapsulating Security Payloads (ESP) group name
- `local` (Attributes) Local parameters for interesting traffic (see [below for nested schema](#nestedatt--local))
- `priority` (Number) Priority for IPsec policy (lowest value more preferable)

    &emsp;|Format  &emsp;|Description                                               |
    |----------|------------------------------------------------------------|
    &emsp;|1-100   &emsp;|Priority for IPsec policy (lowest value more preferable)  |
- `protocol` (String) Protocol

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|txt     &emsp;|Protocol name  |
- `remote` (Attributes) Match remote addresses (see [below for nested schema](#nestedatt--remote))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `peer` (String) Connection name of the peer

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|txt     &emsp;|Connection name of the peer  |
- `tunnel` (Number) Peer tunnel

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|u32     &emsp;|Peer tunnel  |


&lt;a id=&#34;nestedatt--local&#34;&gt;&lt;/a&gt;
### Nested Schema for `local`

Optional:

- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `prefix` (List of String) Local IPv4 or IPv6 prefix

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv4net  &emsp;|Local IPv4 prefix  |
    &emsp;|ipv6net  &emsp;|Local IPv6 prefix  |


&lt;a id=&#34;nestedatt--remote&#34;&gt;&lt;/a&gt;
### Nested Schema for `remote`

Optional:

- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `prefix` (List of String) Remote IPv4 or IPv6 prefix

    &emsp;|Format   &emsp;|Description         |
    |-----------|----------------------|
    &emsp;|ipv4net  &emsp;|Remote IPv4 prefix  |
    &emsp;|ipv6net  &emsp;|Remote IPv6 prefix  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
