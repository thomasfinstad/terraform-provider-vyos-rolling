---
page_title: "vyos_pki_certificate Resource - terraform-provider-vyos"
subcategory: "pki"
description: |-
  Public key infrastructure (PKI)
  ⯯
  Certificate
---

# vyos_pki_certificate (Resource)
<center>

Public key infrastructure (PKI)
⯯
**Certificate**


</center>

## Schema

### Required

- `certificate_id` (String) Certificate

### Optional

- `acme` (Attributes) Automatic Certificate Management Environment (ACME) request (see [below for nested schema](#nestedatt--acme))
- `certificate` (String) Certificate in PEM format
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `private` (Attributes) Certificate private key (see [below for nested schema](#nestedatt--private))
- `revoke` (Boolean) Include certificate in parent CRL

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--acme&#34;&gt;&lt;/a&gt;
### Nested Schema for `acme`

Optional:

- `domain_name` (List of String) Domain Name
- `email` (String) Email address to associate with certificate
- `listen_address` (String) Local IPv4 addresses to listen on

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address to listen for incoming connections  |
- `rsa_key_size` (String) Size of the RSA key

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|2048    &emsp;|RSA key length 2048 bit  |
    &emsp;|3072    &emsp;|RSA key length 3072 bit  |
    &emsp;|4096    &emsp;|RSA key length 4096 bit  |
- `url` (String) Remote URL

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|url     &emsp;|Remote HTTP(S) URL  |


&lt;a id=&#34;nestedatt--private&#34;&gt;&lt;/a&gt;
### Nested Schema for `private`

Optional:

- `key` (String) Private key in PEM format
- `password_protected` (Boolean) Private key portion is password protected  &emsp;|
