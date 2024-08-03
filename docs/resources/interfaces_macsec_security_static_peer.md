---
page_title: "vyos_interfaces_macsec_security_static_peer Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯MACsec Interface (802.1ae)⯯Security/Encryption Settings⯯Use static keys for MACsec [static Secure Authentication Key (SAK) mode]⯯MACsec peer name
---

# vyos_interfaces_macsec_security_static_peer (Resource)
<center>

*interfaces*  
⯯  
MACsec Interface (802.1ae)  
⯯  
Security/Encryption Settings  
⯯  
Use static keys for MACsec [static Secure Authentication Key (SAK) mode]  
⯯  
**MACsec peer name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `disable` (Boolean) Disable instance
- `key` (String) MACsec static key

    &emsp;|Format  &emsp;|Description                                                                                                                   |
    |----------|--------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt     &emsp;|16-byte (128-bit) hex-string (32 hex-digits) for gcm-aes-128 or 32-byte (256-bit) hex-string (64 hex-digits) for gcm-aes-256  |
- `mac` (String) Media Access Control (MAC) address

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|macaddr  &emsp;|Hardware (MAC) address  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `macsec` (String) MACsec Interface (802.1ae)

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|macsecN  &emsp;|MACsec interface name  |
- `peer` (String) MACsec peer name


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
