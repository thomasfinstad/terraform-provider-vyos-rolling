---
page_title: "vyos_system_login_user_authentication_public_keys Resource - vyos"

subcategory: "System"

description: |- 
  system⯯System User Login Configuration⯯Local user account information⯯Authentication settings⯯Remote access public keys
---

# vyos_system_login_user_authentication_public_keys (Resource)
<center>

*system*  
⯯  
System User Login Configuration  
⯯  
Local user account information  
⯯  
Authentication settings  
⯯  
**Remote access public keys**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `key` (String) Public key value (Base64 encoded)
- `options` (String) Optional public key options
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `type` (String) SSH public key type

    &emsp;|Format                              &emsp;|Description                                    |
    |--------------------------------------|-------------------------------------------------|
    &emsp;|ssh-dss                             &emsp;|Digital Signature Algorithm (DSA) key support  |
    &emsp;|ssh-rsa                             &emsp;|Key pair based on RSA algorithm                |
    &emsp;|ecdsa-sha2-nistp256                 &emsp;|Elliptic Curve DSA with NIST P-256 curve       |
    &emsp;|ecdsa-sha2-nistp384                 &emsp;|Elliptic Curve DSA with NIST P-384 curve       |
    &emsp;|ecdsa-sha2-nistp521                 &emsp;|Elliptic Curve DSA with NIST P-521 curve       |
    &emsp;|ssh-ed25519                         &emsp;|Edwards-curve DSA with elliptic curve 25519    |
    &emsp;|sk-ecdsa-sha2-nistp256@openssh.com  &emsp;|Elliptic Curve DSA security key                |
    &emsp;|sk-ssh-ed25519@openssh.com          &emsp;|Elliptic curve 25519 security key              |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `public_keys` (String) Remote access public keys

    &emsp;|Format  &emsp;|Description                                                    |
    |----------|-----------------------------------------------------------------|
    &emsp;|txt     &emsp;|Key identifier used by ssh-keygen (usually of form user@host)  |
- `user` (String) Local user account information


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
