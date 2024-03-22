---
page_title: "vyos_pki_openssh Resource - terraform-provider-vyos"
subcategory: "pki"
description: |-
  Public key infrastructure (PKI)⯯OpenSSH public and private keys
---

# vyos_pki_openssh (Resource)
<center>

Public key infrastructure (PKI)
⯯
**OpenSSH public and private keys**


</center>

## Schema

### Required

- `openssh_id` (String) OpenSSH public and private keys

### Optional

- `private` (Attributes) Private key (see [below for nested schema](#nestedatt--private))
- `public` (Attributes) Public key (see [below for nested schema](#nestedatt--public))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--private&#34;&gt;&lt;/a&gt;
### Nested Schema for `private`

Optional:

- `key` (String) Private key in PEM format
- `password_protected` (Boolean) Private key portion is password protected


&lt;a id=&#34;nestedatt--public&#34;&gt;&lt;/a&gt;
### Nested Schema for `public`

Optional:

- `key` (String) Public key in PEM format
- `type` (String) SSH public key type

    &emsp;|Format   &emsp;|Description                      |
    |-----------|-----------------------------------|
    &emsp;|ssh-rsa  &emsp;|Key pair based on RSA algorithm  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
