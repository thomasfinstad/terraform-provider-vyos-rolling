---
page_title: "vyos_pki_ca Resource - vyos"

subcategory: "Pki"

description: |-
  Public key infrastructure (PKI)⯯Certificate Authority
---

# vyos_pki_ca (Resource)
<center>


Public key infrastructure (PKI)  
⯯  
**Certificate Authority**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_pki_ca (Resource)](#vyos_pki_ca-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [certificate](#certificate)
      - [crl](#crl)
      - [description](#description)
      - [private](#private)
      - [revoke](#revoke)
      - [system_install](#system_install)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `private`](#nested-schema-for-private)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### certificate
- `certificate` (String) Certificate in PEM format
#### crl
- `crl` (List of String) Certificate revocation list in PEM format
#### description
- `description` (String) Description

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  txt     &emsp;|  Description  |
#### private
- `private` (Attributes) CA private key in PEM format (see [below for nested schema](#nestedatt--private))
#### revoke
- `revoke` (Boolean) Include certificate in parent CRL
#### system_install
- `system_install` (Boolean) Install into CA certificate store on router
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `ca` (String) Certificate Authority


<a id="nestedatt--private"></a>
### Nested Schema for `private`

Optional:

- `key` (String) Private key in PEM format
- `password_protected` (Boolean) Private key portion is password protected


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_pki_ca.example "pki__ca__<ca>"
```
