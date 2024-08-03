---
page_title: "vyos_vpn_ipsec_remote_access_pool Resource - vyos"

subcategory: "Vpn"

description: |- 
  Virtual Private Network (VPN)⯯VPN IP security (IPsec) parameters⯯IKEv2 remote access VPN⯯IP address pool for remote access users
---

# vyos_vpn_ipsec_remote_access_pool (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
VPN IP security (IPsec) parameters  
⯯  
IKEv2 remote access VPN  
⯯  
**IP address pool for remote access users**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `exclude` (List of String) Local IPv4 or IPv6 pool prefix exclusions

    &emsp;|Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    &emsp;|ipv4net  &emsp;|Local IPv4 pool prefix exclusion  |
    &emsp;|ipv6net  &emsp;|Local IPv6 pool prefix exclusion  |
- `name_server` (List of String) Domain Name Servers (DNS) addresses

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|ipv4    &emsp;|Domain Name Server (DNS) IPv4 address  |
    &emsp;|ipv6    &emsp;|Domain Name Server (DNS) IPv6 address  |
- `prefix` (String) Local IPv4 or IPv6 pool prefix

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|ipv4net  &emsp;|Local IPv4 pool prefix  |
    &emsp;|ipv6net  &emsp;|Local IPv6 pool prefix  |
- `range` (Attributes) Local IPv4 or IPv6 pool range (see [below for nested schema](#nestedatt--range))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `pool` (String) IP address pool for remote access users


&lt;a id=&#34;nestedatt--range&#34;&gt;&lt;/a&gt;
### Nested Schema for `range`

Optional:

- `start` (String) First IP address for local pool range

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 start address of pool  |
    &emsp;|ipv6    &emsp;|IPv6 start address of pool  |
- `stop` (String) Last IP address for local pool range

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|ipv4    &emsp;|IPv4 end address of pool  |
    &emsp;|ipv6    &emsp;|IPv6 end address of pool  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
