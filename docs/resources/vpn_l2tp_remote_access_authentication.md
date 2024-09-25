---
page_title: "vyos_vpn_l2tp_remote_access_authentication Resource - vyos"

subcategory: "Vpn"

description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Virtual Private Network (VPN)⯯L2TP Virtual Private Network (VPN)⯯Remote access L2TP VPN⯯Authentication for remote access L2TP VPN
---

# vyos_vpn_l2tp_remote_access_authentication (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Virtual Private Network (VPN)  
⯯  
L2TP Virtual Private Network (VPN)  
⯯  
Remote access L2TP VPN  
⯯  
**Authentication for remote access L2TP VPN**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_vpn_l2tp_remote_access_authentication (Resource)](#vyos_vpn_l2tp_remote_access_authentication-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [mode](#mode)
      - [protocols](#protocols)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### mode
- `mode` (String) Authentication mode used by this server

    |  Format  &emsp;|  Description                                |
    |----------|---------------------------------------------|
    |  local   &emsp;|  Use local username/password configuration  |
    |  radius  &emsp;|  Use RADIUS server for user autentication   |
    |  noauth  &emsp;|  Authentication disabled                    |
#### protocols
- `protocols` (List of String) Authentication protocol for remote access peer

    |  Format     &emsp;|  Description                                                                                      |
    |-------------|---------------------------------------------------------------------------------------------------|
    |  pap        &emsp;|  Authentication via PAP (Password Authentication Protocol)                                        |
    |  chap       &emsp;|  Authentication via CHAP (Challenge Handshake Authentication Protocol)                            |
    |  mschap     &emsp;|  Authentication via MS-CHAP (Microsoft Challenge Handshake Authentication Protocol)               |
    |  mschap-v2  &emsp;|  Authentication via MS-CHAPv2 (Microsoft Challenge Handshake Authentication Protocol, version 2)  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).