---
page_title: "vyos_vpn_sstp_authentication_local_users_username Resource - vyos"

subcategory: "Vpn"

description: |- 
  vpn⯯Secure Socket Tunneling Protocol (SSTP) server⯯Authentication for remote access SSTP Server⯯Local user authentication for PPPoE server⯯User name for authentication
---

# vyos_vpn_sstp_authentication_local_users_username (Resource)
<center>

*vpn*  
⯯  
Secure Socket Tunneling Protocol (SSTP) server  
⯯  
Authentication for remote access SSTP Server  
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

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `username` (String) User name for authentication


<a id="nestedatt--rate_limit"></a>
### Nested Schema for `rate_limit`

Optional:

- `download` (String) Download bandwidth limit in kbits/sec
- `upload` (String) Upload bandwidth limit in kbits/sec


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
