---
page_title: "vyos_pki_openssh Resource - vyos"

subcategory: "Pki"

description: |- 
  Public key infrastructure (PKI)⯯OpenSSH public and private keys
---

# vyos_pki_openssh (Resource)
<center>

Public key infrastructure (PKI)  
⯯  
**OpenSSH public and private keys**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `private` (Attributes) Private key (see [below for nested schema](#nestedatt--private))
- `public` (Attributes) Public key (see [below for nested schema](#nestedatt--public))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `openssh` (String) OpenSSH public and private keys


<a id="nestedatt--private"></a>
### Nested Schema for `private`

Optional:

- `key` (String) Private key in PEM format
- `password_protected` (Boolean) Private key portion is password protected


<a id="nestedatt--public"></a>
### Nested Schema for `public`

Optional:

- `key` (String) Public key in PEM format
- `type` (String) SSH public key type

    |Format   &emsp;|Description                      |
    |-----------|-----------------------------------|
    |ssh-rsa  &emsp;|Key pair based on RSA algorithm  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
