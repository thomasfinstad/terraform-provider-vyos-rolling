---
page_title: "vyos_vpn_openconnect Resource - vyos"

subcategory: "Vpn"

description: |-
  ~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.
  Virtual Private Network (VPN)⯯SSL VPN OpenConnect, AnyConnect compatible server
---

# vyos_vpn_openconnect (Resource)
<center>

~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

Virtual Private Network (VPN)  
⯯  
**SSL VPN OpenConnect, AnyConnect compatible server**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_vpn_openconnect (Resource)](#vyos_vpn_openconnect-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [http_security_headers](#http_security_headers)
      - [listen_address](#listen_address)
      - [timeouts](#timeouts)
      - [tls_version_min](#tls_version_min)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### http_security_headers
- `http_security_headers` (Boolean) Enable HTTP security headers
#### listen_address
- `listen_address` (String) Local IPv4 addresses to listen on

    |  Format  &emsp;|  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    &emsp;|  IPv4 address to listen for incoming connections  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
#### tls_version_min
- `tls_version_min` (String) Specify the minimum required TLS version

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  1.0     &emsp;|  TLS v1.0     |
    |  1.1     &emsp;|  TLS v1.1     |
    |  1.2     &emsp;|  TLS v1.2     |
    |  1.3     &emsp;|  TLS v1.3     |

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
terraform import vyos_vpn_openconnect.example "vpn__openconnect"
```
