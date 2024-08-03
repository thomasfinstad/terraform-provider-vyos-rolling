---
page_title: "vyos_vpn_sstp_client_ipv6_pool_prefix Resource - vyos"

subcategory: "Vpn"

description: |- 
  vpn⯯Secure Socket Tunneling Protocol (SSTP) server⯯Pool of client IPv6 addresses⯯Pool of addresses used to assign to clients
---

# vyos_vpn_sstp_client_ipv6_pool_prefix (Resource)
<center>

*vpn*  
⯯  
Secure Socket Tunneling Protocol (SSTP) server  
⯯  
Pool of client IPv6 addresses  
⯯  
**Pool of addresses used to assign to clients**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `mask` (Number) Prefix length used for individual client

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|48-128  &emsp;|Client prefix length  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `client_ipv6_pool` (String) Pool of client IPv6 addresses

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|Name of IPv6 pool  |
- `prefix` (String) Pool of addresses used to assign to clients

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv6net  &emsp;|IPv6 address and prefix length  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
