---
page_title: "vyos_vpn_openconnect_authentication_local_users_username Resource - vyos"

subcategory: "Vpn"

description: |- 
  vpn⯯SSL VPN OpenConnect, AnyConnect compatible server⯯Authentication for remote access SSL VPN Server⯯Local user authentication⯯Username used for authentication
---

# vyos_vpn_openconnect_authentication_local_users_username (Resource)
<center>

*vpn*  
⯯  
SSL VPN OpenConnect, AnyConnect compatible server  
⯯  
Authentication for remote access SSL VPN Server  
⯯  
Local user authentication  
⯯  
**Username used for authentication**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `disable` (Boolean) Disable instance
- `otp` (Attributes) 2FA OTP authentication parameters (see [below for nested schema](#nestedatt--otp))
- `password` (String) Password used for authentication
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `username` (String) Username used for authentication

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|txt     &emsp;|Username used for authentication  |


&lt;a id=&#34;nestedatt--otp&#34;&gt;&lt;/a&gt;
### Nested Schema for `otp`

Optional:

- `interval` (Number) Time tokens interval in seconds

    &emsp;|Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    &emsp;|5-86400  &emsp;|Time tokens interval in seconds.  |
- `key` (String) Token Key Secret key for the token algorithm (see RFC 4226)

    &emsp;|Format  &emsp;|Description                    |
    |----------|---------------------------------|
    &emsp;|txt     &emsp;|OTP key in hex-encoded format  |
- `otp_length` (Number) Number of digits in OTP code

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|6-8     &emsp;|Number of digits in OTP code  |
- `token_type` (String) Token type

    &emsp;|Format      &emsp;|Description                |
    |--------------|-----------------------------|
    &emsp;|hotp-time   &emsp;|Time-based OTP algorithm   |
    &emsp;|hotp-event  &emsp;|Event-based OTP algorithm  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  