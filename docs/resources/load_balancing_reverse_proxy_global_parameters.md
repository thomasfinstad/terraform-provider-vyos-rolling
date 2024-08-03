---
page_title: "vyos_load_balancing_reverse_proxy_global_parameters Resource - vyos"

subcategory: "Load"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  load-balancing⯯Configure reverse-proxy⯯Global perfomance parameters and limits
---

# vyos_load_balancing_reverse_proxy_global_parameters (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*load-balancing*  
⯯  
Configure reverse-proxy  
⯯  
**Global perfomance parameters and limits**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `max_connections` (Number) Maximum allowed connections

    &emsp;|Format     &emsp;|Description                  |
    |-------------|-------------------------------|
    &emsp;|1-2000000  &emsp;|Maximum allowed connections  |
- `ssl_bind_ciphers` (List of String) Cipher algorithms (&#34;cipher suite&#34;) used during SSL/TLS handshake for all frontend servers

    &emsp;|Format                         &emsp;|Description                    |
    |---------------------------------|---------------------------------|
    &emsp;|ecdhe-ecdsa-aes128-gcm-sha256  &emsp;|ecdhe-ecdsa-aes128-gcm-sha256  |
    &emsp;|ecdhe-rsa-aes128-gcm-sha256    &emsp;|ecdhe-rsa-aes128-gcm-sha256    |
    &emsp;|ecdhe-ecdsa-aes256-gcm-sha384  &emsp;|ecdhe-ecdsa-aes256-gcm-sha384  |
    &emsp;|ecdhe-rsa-aes256-gcm-sha384    &emsp;|ecdhe-rsa-aes256-gcm-sha384    |
    &emsp;|ecdhe-ecdsa-chacha20-poly1305  &emsp;|ecdhe-ecdsa-chacha20-poly1305  |
    &emsp;|ecdhe-rsa-chacha20-poly1305    &emsp;|ecdhe-rsa-chacha20-poly1305    |
    &emsp;|dhe-rsa-aes128-gcm-sha256      &emsp;|dhe-rsa-aes128-gcm-sha256      |
    &emsp;|dhe-rsa-aes256-gcm-sha384      &emsp;|dhe-rsa-aes256-gcm-sha384      |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `tls_version_min` (String) Specify the minimum required TLS version

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1.2     &emsp;|TLS v1.2     |
    &emsp;|1.3     &emsp;|TLS v1.3     |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
