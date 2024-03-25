---
page_title: "vyos_pki_dh Resource - vyos"
subcategory: "pki"
description: |- 
  Public key infrastructure (PKI)⯯Diffie-Hellman parameters
---

# vyos_pki_dh (Resource)
<center>

Public key infrastructure (PKI)  
⯯  
**Diffie-Hellman parameters**


</center>

## Schema

### Required

- `dh_id` (String) Diffie-Hellman parameters

### Optional

- `parameters` (String) DH parameters in PEM format
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
