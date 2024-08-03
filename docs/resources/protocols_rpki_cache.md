---
page_title: "vyos_protocols_rpki_cache Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Resource Public Key Infrastructure (RPKI)⯯RPKI cache server address
---

# vyos_protocols_rpki_cache (Resource)
<center>

*protocols*  
⯯  
Resource Public Key Infrastructure (RPKI)  
⯯  
**RPKI cache server address**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `preference` (Number) Preference of the cache server

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|1-255   &emsp;|Preference of the cache server  |
- `ssh` (Attributes) RPKI SSH connection settings (see [below for nested schema](#nestedatt--ssh))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `cache` (String) RPKI cache server address

    &emsp;|Format    &emsp;|Description                                 |
    |------------|----------------------------------------------|
    &emsp;|ipv4      &emsp;|IP address of RPKI server                   |
    &emsp;|ipv6      &emsp;|IPv6 address of RPKI server                 |
    &emsp;|hostname  &emsp;|Fully qualified domain name of RPKI server  |


&lt;a id=&#34;nestedatt--ssh&#34;&gt;&lt;/a&gt;
### Nested Schema for `ssh`

Optional:

- `key` (String) OpenSSH key in PKI configuration

    &emsp;|Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    &emsp;|txt     &emsp;|Name of OpenSSH key in PKI configuration  |
- `username` (String) Username used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Username     |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
