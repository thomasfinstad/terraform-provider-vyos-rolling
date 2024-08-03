---
page_title: "vyos_vpn_sstp_client_ipv6_pool Resource - vyos"

subcategory: "Vpn"

description: |- 
  vpn⯯Secure Socket Tunneling Protocol (SSTP) server⯯Pool of client IPv6 addresses
---

# vyos_vpn_sstp_client_ipv6_pool (Resource)
<center>

*vpn*  
⯯  
Secure Socket Tunneling Protocol (SSTP) server  
⯯  
**Pool of client IPv6 addresses**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

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


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
