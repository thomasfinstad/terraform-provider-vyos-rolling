---
page_title: "vyos_vpn_openconnect_authentication_mode Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  vpn⯯SSL VPN OpenConnect, AnyConnect compatible server⯯Authentication for remote access SSL VPN Server⯯Authentication mode used by this server
---

# vyos_vpn_openconnect_authentication_mode (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*vpn*  
⯯  
SSL VPN OpenConnect, AnyConnect compatible server  
⯯  
Authentication for remote access SSL VPN Server  
⯯  
**Authentication mode used by this server**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `local` (String) Use local username/password configuration (OTP supported)

    &emsp;|Format        &emsp;|Description                                  |
    |----------------|-----------------------------------------------|
    &emsp;|password      &emsp;|Password-only local authentication           |
    &emsp;|otp           &emsp;|OTP-only local authentication                |
    &emsp;|password-otp  &emsp;|Password (first) + OTP local authentication  |
- `radius` (Boolean) Use RADIUS server for user autentication
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  