---
page_title: "vyos_pki_key_pair Resource - vyos"

subcategory: "Pki"

description: |- 
  Public key infrastructure (PKI)⯯Public and private keys
---

# vyos_pki_key_pair (Resource)
<center>

Public key infrastructure (PKI)  
⯯  
**Public and private keys**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `key_pair_id` (String) Public and private keys

### Optional

- `private` (Attributes) Private key (see [below for nested schema](#nestedatt--private))
- `public` (Attributes) Public key (see [below for nested schema](#nestedatt--public))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--private&#34;&gt;&lt;/a&gt;
### Nested Schema for `private`

Optional:

- `key` (String) Private key in PEM format
- `password_protected` (Boolean) Private key portion is password protected


&lt;a id=&#34;nestedatt--public&#34;&gt;&lt;/a&gt;
### Nested Schema for `public`

Optional:

- `key` (String) Public key in PEM format


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
