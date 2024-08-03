---
page_title: "vyos_vpn_pptp_remote_access_authentication_local_users_username Resource - vyos"

subcategory: "Vpn"

description: |- 
  vpn⯯Point to Point Tunneling Protocol (PPTP) Virtual Private Network (VPN)⯯Remote access PPTP VPN⯯Authentication for remote access PPTP VPN⯯Local user authentication for PPPoE server⯯User name for authentication
---

# vyos_vpn_pptp_remote_access_authentication_local_users_username (Resource)
<center>

*vpn*  
⯯  
Point to Point Tunneling Protocol (PPTP) Virtual Private Network (VPN)  
⯯  
Remote access PPTP VPN  
⯯  
Authentication for remote access PPTP VPN  
⯯  
Local user authentication for PPPoE server  
⯯  
**User name for authentication**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `disable` (Boolean) Disable instance
- `password` (String) Password for authentication
- `rate_limit` (Attributes) Upload/Download speed limits (see [below for nested schema](#nestedatt--rate_limit))
- `static_ip` (String) Static client IP address
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `username` (String) User name for authentication


&lt;a id=&#34;nestedatt--rate_limit&#34;&gt;&lt;/a&gt;
### Nested Schema for `rate_limit`

Optional:

- `download` (String) Download bandwidth limit in kbits/sec
- `upload` (String) Upload bandwidth limit in kbits/sec


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
