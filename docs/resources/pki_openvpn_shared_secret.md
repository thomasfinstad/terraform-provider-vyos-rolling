---
page_title: "vyos_pki_openvpn_shared_secret Resource - terraform-provider-vyos"
subcategory: "pki"
description: |-
  Public key infrastructure (PKI)
  ⯯
  OpenVPN keys
  ⯯
  OpenVPN shared secret key
---

# vyos_pki_openvpn_shared_secret (Resource)
<center>

Public key infrastructure (PKI)
⯯
OpenVPN keys
⯯
**OpenVPN shared secret key**


</center>

## Schema

### Required

- `shared_secret_id` (String) OpenVPN shared secret key

### Optional

- `key` (String) OpenVPN shared secret key data
- `version` (String) OpenVPN shared secret key version

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
