---
page_title: "vyos_service_https Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯HTTPS configuration
---

# vyos_service_https (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
**HTTPS configuration**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `enable_http_redirect` (Boolean) Enable HTTP to HTTPS redirect
- `listen_address` (List of String) Local IP addresses to listen on

    |Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    |ipv4    &emsp;|IPv4 address to listen for incoming connections  |
    |ipv6    &emsp;|IPv6 address to listen for incoming connections  |
- `port` (Number) Port number used by connection

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |1-65535  &emsp;|Numeric IP port  |
- `request_body_size_limit` (Number) Maximum request body size in megabytes

    |Format  &emsp;|Description                     |
    |----------|----------------------------------|
    |1-256   &emsp;|Request body size in megabytes  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `tls_version` (List of String) Specify available TLS version(s)

    |Format  &emsp;|Description  |
    |----------|---------------|
    |1.2     &emsp;|TLSv1.2      |
    |1.3     &emsp;|TLSv1.3      |
- `vrf` (String) VRF instance name

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
