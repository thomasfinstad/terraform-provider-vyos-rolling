---
page_title: "vyos_vpn_l2tp_remote_access_authentication_radius_server Resource - vyos"

subcategory: "Vpn"

description: |-
  Virtual Private Network (VPN)⯯L2TP Virtual Private Network (VPN)⯯Remote access L2TP VPN⯯Authentication for remote access L2TP VPN⯯RADIUS based user authentication⯯RADIUS server configuration
---

# vyos_vpn_l2tp_remote_access_authentication_radius_server (Resource)
<center>

Virtual Private Network (VPN)  
⯯  
L2TP Virtual Private Network (VPN)  
⯯  
Remote access L2TP VPN  
⯯  
Authentication for remote access L2TP VPN  
⯯  
RADIUS based user authentication  
⯯  
**RADIUS server configuration**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_vpn_l2tp_remote_access_authentication_radius_server (Resource)](#vyos_vpn_l2tp_remote_access_authentication_radius_server-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [acct_port](#acct_port)
      - [backup](#backup)
      - [disable](#disable)
      - [disable_accounting](#disable_accounting)
      - [fail_time](#fail_time)
      - [key](#key)
      - [port](#port)
      - [priority](#priority)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### acct_port
- `acct_port` (Number) Accounting port

    |  Format   &emsp;|  Description      |
    |-----------|-------------------|
    |  1-65535  &emsp;|  Numeric IP port  |
#### backup
- `backup` (Boolean) Use backup server if other servers are not available
#### disable
- `disable` (Boolean) Disable instance
#### disable_accounting
- `disable_accounting` (Boolean) Disable accounting
#### fail_time
- `fail_time` (Number) Mark server unavailable for &lt;n&gt; seconds on failure

    |  Format  &emsp;|  Description        |
    |----------|---------------------|
    |  0-600   &emsp;|  Fail time penalty  |
#### key
- `key` (String) Shared secret key

    |  Format  &emsp;|  Description            |
    |----------|-------------------------|
    |  txt     &emsp;|  Password string (key)  |
#### port
- `port` (Number) Port number used by connection

    |  Format   &emsp;|  Description      |
    |-----------|-------------------|
    |  1-65535  &emsp;|  Numeric IP port  |
#### priority
- `priority` (Number) Server priority

    |  Format  &emsp;|  Description      |
    |----------|-------------------|
    |  1-255   &emsp;|  Server priority  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `server` (String) RADIUS server configuration

    |  Format  &emsp;|  Description                 |
    |----------|------------------------------|
    |  ipv4    &emsp;|  RADIUS server IPv4 address  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).