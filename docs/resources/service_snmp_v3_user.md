---
page_title: "vyos_service_snmp_v3_user Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Simple Network Management Protocol (SNMP)⯯Simple Network Management Protocol (SNMP) v3⯯Specifies the user with name username
---

# vyos_service_snmp_v3_user (Resource)
<center>

*service*  
⯯  
Simple Network Management Protocol (SNMP)  
⯯  
Simple Network Management Protocol (SNMP) v3  
⯯  
**Specifies the user with name username**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `auth` (Attributes) Specifies the auth (see [below for nested schema](#nestedatt--auth))
- `group` (String) Specifies group for user name
- `mode` (String) Define access permission

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|ro      &emsp;|Read-Only    |
    &emsp;|rw      &emsp;|read write   |
- `privacy` (Attributes) Defines the privacy (see [below for nested schema](#nestedatt--privacy))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `user` (String) Specifies the user with name username


&lt;a id=&#34;nestedatt--auth&#34;&gt;&lt;/a&gt;
### Nested Schema for `auth`

Optional:

- `encrypted_password` (String) Defines the encrypted key for authentication
- `plaintext_password` (String) Defines the clear text key for authentication
- `type` (String) Define used protocol

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|md5     &emsp;|Message Digest 5       |
    &emsp;|sha     &emsp;|Secure Hash Algorithm  |


&lt;a id=&#34;nestedatt--privacy&#34;&gt;&lt;/a&gt;
### Nested Schema for `privacy`

Optional:

- `encrypted_password` (String) Defines the encrypted key for privacy protocol
- `plaintext_password` (String) Defines the clear text key for privacy protocol
- `type` (String) Defines the protocol for privacy

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|des     &emsp;|Data Encryption Standard      |
    &emsp;|aes     &emsp;|Advanced Encryption Standard  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
