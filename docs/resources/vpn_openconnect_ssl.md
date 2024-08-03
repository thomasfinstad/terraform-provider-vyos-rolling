---
page_title: "vyos_vpn_openconnect_ssl Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  vpn⯯SSL VPN OpenConnect, AnyConnect compatible server⯯SSL Certificate, SSL Key and CA
---

# vyos_vpn_openconnect_ssl (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*vpn*  
⯯  
SSL VPN OpenConnect, AnyConnect compatible server  
⯯  
**SSL Certificate, SSL Key and CA**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `ca_certificate` (List of String) Certificate Authority chain in PKI configuration

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|Name of CA in PKI configuration  |
- `certificate` (String) Certificate in PKI configuration

    &emsp;|Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    &emsp;|txt     &emsp;|Name of certificate in PKI configuration  |
- `passphrase` (String) Private key passphrase

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|txt     &emsp;|Passphrase to decrypt the private key  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
