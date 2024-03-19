---
page_title: "vyos_pki_ca Resource - terraform-provider-vyos"
subcategory: "pki"
description: |-
  Public key infrastructure (PKI)
  ⯯
  Certificate Authority
---

# vyos_pki_ca (Resource)
<center>

Public key infrastructure (PKI)
⯯
**Certificate Authority**


</center>

## Schema

### Required

- `ca_id` (String) Certificate Authority

### Optional

- `certificate` (String) Certificate in PEM format
- `crl` (List of String) Certificate revocation list in PEM format
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `private` (Attributes) CA private key in PEM format (see [below for nested schema](#nestedatt--private))
- `revoke` (Boolean) Include certificate in parent CRL

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--private&#34;&gt;&lt;/a&gt;
### Nested Schema for `private`

Optional:

- `key` (String) Private key in PEM format
- `password_protected` (Boolean) Private key portion is password protected  &emsp;|
