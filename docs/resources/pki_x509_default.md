---
page_title: "vyos_pki_x509_default Resource - terraform-provider-vyos"
subcategory: "pki"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Public key infrastructure (PKI)
  ⯯
  X509 Settings
  ⯯
  X509 Default Values
---

# vyos_pki_x509_default (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

	Public key infrastructure (PKI)
⯯
X509 Settings
⯯
**X509 Default Values**


</center>

## Schema

### Optional

- `country` (String) Default country
- `locality` (String) Default locality
- `organization` (String) Default organization
- `state` (String) Default state

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
