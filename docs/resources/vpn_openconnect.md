---
page_title: "vyos_vpn_openconnect Resource - vyos"

subcategory: "Vpn"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  vpn⯯SSL VPN OpenConnect, AnyConnect compatible server
---

# vyos_vpn_openconnect (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*vpn*  
⯯  
**SSL VPN OpenConnect, AnyConnect compatible server**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `http_security_headers` (Boolean) Enable HTTP security headers
- `listen_address` (String) Local IPv4 addresses to listen on

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address to listen for incoming connections  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `tls_version_min` (String) Specify the minimum required TLS version

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1.0     &emsp;|TLS v1.0     |
    &emsp;|1.1     &emsp;|TLS v1.1     |
    &emsp;|1.2     &emsp;|TLS v1.2     |
    &emsp;|1.3     &emsp;|TLS v1.3     |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
