---
page_title: "vyos_pki_openvpn_shared_secret Resource - vyos"

subcategory: "Pki"

description: |- 
  Public key infrastructure (PKI)⯯OpenVPN keys⯯OpenVPN shared secret key
---

# vyos_pki_openvpn_shared_secret (Resource)
<center>

Public key infrastructure (PKI)  
⯯  
OpenVPN keys  
⯯  
**OpenVPN shared secret key**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `shared_secret_id` (String) OpenVPN shared secret key

### Optional

- `key` (String) OpenVPN shared secret key data
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `version` (String) OpenVPN shared secret key version

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
