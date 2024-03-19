---
page_title: "vyos_pki_openssh Resource - terraform-provider-vyos"
subcategory: "pki"
description: |-
  Public key infrastructure (PKI)
  ⯯
  OpenSSH public and private keys
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
    &emsp;|ssh-rsa  &emsp;|Key pair based on RSA algorithm  |  &emsp;|
