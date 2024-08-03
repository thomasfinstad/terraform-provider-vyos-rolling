---
page_title: "vyos_vpn_sstp_ppp_options Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  vpn⯯Secure Socket Tunneling Protocol (SSTP) server⯯Advanced protocol options
---

# vyos_vpn_sstp_ppp_options (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*vpn*  
⯯  
Secure Socket Tunneling Protocol (SSTP) server  
⯯  
**Advanced protocol options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `disable_ccp` (Boolean) Disable Compression Control Protocol (CCP)
- `interface_cache` (Number) PPP interface cache

    &emsp;|Format    &emsp;|Description                           |
    |------------|----------------------------------------|
    &emsp;|1-256000  &emsp;|Count of interfaces to keep in cache  |
- `ipv4` (String) IPv4 (IPCP) negotiation algorithm

    &emsp;|Format   &emsp;|Description                                                 |
    |-----------|--------------------------------------------------------------|
    &emsp;|deny     &emsp;|Do not negotiate IPv4                                       |
    &emsp;|allow    &emsp;|Negotiate IPv4 only if client requests                      |
    &emsp;|prefer   &emsp;|Ask client for IPv4 negotiation, do not fail if it rejects  |
    &emsp;|require  &emsp;|Require IPv4 negotiation                                    |
- `ipv6` (String) IPv6 (IPCP6) negotiation algorithm

    &emsp;|Format   &emsp;|Description                                                 |
    |-----------|--------------------------------------------------------------|
    &emsp;|deny     &emsp;|Do not negotiate IPv6                                       |
    &emsp;|allow    &emsp;|Negotiate IPv6 only if client requests                      |
    &emsp;|prefer   &emsp;|Ask client for IPv6 negotiation, do not fail if it rejects  |
    &emsp;|require  &emsp;|Require IPv6 negotiation                                    |
- `ipv6_accept_peer_interface_id` (Boolean) Accept peer interface identifier
- `ipv6_interface_id` (String) Fixed or random interface identifier for IPv6

    &emsp;|Format   &emsp;|Description                            |
    |-----------|-----------------------------------------|
    &emsp;|random   &emsp;|Random interface identifier for IPv6   |
    &emsp;|x:x:x:x  &emsp;|specify interface identifier for IPv6  |
- `ipv6_peer_interface_id` (String) Peer interface identifier for IPv6

    &emsp;|Format       &emsp;|Description                                                                |
    |---------------|-----------------------------------------------------------------------------|
    &emsp;|x:x:x:x      &emsp;|Interface identifier for IPv6                                              |
    &emsp;|random       &emsp;|Use a random interface identifier for IPv6                                 |
    &emsp;|ipv4-addr    &emsp;|Calculate interface identifier from IPv4 address, for example 192:168:0:1  |
    &emsp;|calling-sid  &emsp;|Calculate interface identifier from calling-station-id                     |
- `lcp_echo_failure` (String) Maximum number of Echo-Requests may be sent without valid reply
- `lcp_echo_interval` (String) LCP echo-requests/sec
- `lcp_echo_timeout` (String) Timeout in seconds to wait for any peer activity. If this option specified it turns on adaptive lcp echo functionality and &#34;lcp-echo-failure&#34; is not used.
- `min_mtu` (String) Minimum acceptable MTU (68-65535)
- `mppe` (String) Specifies mppe negotiation preferences

    &emsp;|Format   &emsp;|Description                                                |
    |-----------|-------------------------------------------------------------|
    &emsp;|require  &emsp;|send mppe request, if client rejects, drop the connection  |
    &emsp;|prefer   &emsp;|send mppe request, if client rejects continue              |
    &emsp;|deny     &emsp;|drop all mppe                                              |
- `mru` (String) Preferred MRU (68-65535)
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
