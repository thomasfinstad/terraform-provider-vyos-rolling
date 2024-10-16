---
page_title: "vyos_vpn_sstp_ssl Resource - vyos"

subcategory: "Vpn"

description: |-
  ~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.
  Virtual Private Network (VPN)⯯Secure Socket Tunneling Protocol (SSTP) server⯯SSL Certificate, SSL Key and CA
---

# vyos_vpn_sstp_ssl (Resource)
<center>

~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

Virtual Private Network (VPN)  
⯯  
Secure Socket Tunneling Protocol (SSTP) server  
⯯  
**SSL Certificate, SSL Key and CA**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_vpn_sstp_ssl (Resource)](#vyos_vpn_sstp_ssl-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [ca_certificate](#ca_certificate)
      - [certificate](#certificate)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### ca_certificate
- `ca_certificate` (String) Certificate Authority in PKI configuration

    |  Format  &emsp;|  Description                      |
    |----------|-----------------------------------|
    |  txt     &emsp;|  Name of CA in PKI configuration  |
#### certificate
- `certificate` (String) Certificate in PKI configuration

    |  Format  &emsp;|  Description                               |
    |----------|--------------------------------------------|
    |  txt     &emsp;|  Name of certificate in PKI configuration  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_vpn_sstp_ssl.example "vpn__sstp__ssl"
```
