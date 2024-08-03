---
page_title: "vyos_vpn_sstp_authentication Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  vpn⯯Secure Socket Tunneling Protocol (SSTP) server⯯Authentication for remote access SSTP Server
---

# vyos_vpn_sstp_authentication (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*vpn*  
⯯  
Secure Socket Tunneling Protocol (SSTP) server  
⯯  
**Authentication for remote access SSTP Server**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `mode` (String) Authentication mode used by this server

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|local   &emsp;|Use local username/password configuration  |
    &emsp;|radius  &emsp;|Use RADIUS server for user autentication   |
    &emsp;|noauth  &emsp;|Authentication disabled                    |
- `protocols` (List of String) Authentication protocol for remote access peer

    &emsp;|Format     &emsp;|Description                                                                                      |
    |-------------|---------------------------------------------------------------------------------------------------|
    &emsp;|pap        &emsp;|Authentication via PAP (Password Authentication Protocol)                                        |
    &emsp;|chap       &emsp;|Authentication via CHAP (Challenge Handshake Authentication Protocol)                            |
    &emsp;|mschap     &emsp;|Authentication via MS-CHAP (Microsoft Challenge Handshake Authentication Protocol)               |
    &emsp;|mschap-v2  &emsp;|Authentication via MS-CHAPv2 (Microsoft Challenge Handshake Authentication Protocol, version 2)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
