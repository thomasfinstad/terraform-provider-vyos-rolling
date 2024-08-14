---
page_title: "vyos_system_login_user Resource - vyos"

subcategory: "System"

description: |- 
  system⯯System User Login Configuration⯯Local user account information
---

# vyos_system_login_user (Resource)
<center>

*system*  
⯯  
System User Login Configuration  
⯯  
**Local user account information**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `authentication` (Attributes) Authentication settings (see [below for nested schema](#nestedatt--authentication))
- `disable` (Boolean) Disable instance
- `full_name` (String) Full name of the user (use quotes for names with spaces)
- `home_directory` (String) Home directory

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Path to home directory  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `user` (String) Local user account information


&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `encrypted_password` (String) Encrypted password
- `otp` (Attributes) One-Time-Pad (two-factor) authentication parameters (see [below for nested schema](#nestedatt--authentication--otp))
- `plaintext_password` (String) Plaintext password used for encryption

&lt;a id=&#34;nestedatt--authentication--otp&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication.otp`

Optional:

- `key` (String) Key/secret the token algorithm (see RFC4226)

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Base32 encoded key/token  |
- `rate_limit` (Number) Limit number of logins (rate-limit) per rate-time

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|1-10    &emsp;|Number of attempts  |
- `rate_time` (Number) Limit number of logins (rate-limit) per rate-time

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|15-600  &emsp;|Time interval  |
- `window_size` (Number) Set window of concurrently valid codes

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1-21    &emsp;|Window size  |



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  