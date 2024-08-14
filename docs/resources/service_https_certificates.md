---
page_title: "vyos_service_https_certificates Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯HTTPS configuration⯯TLS certificates
---

# vyos_service_https_certificates (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
HTTPS configuration  
⯯  
**TLS certificates**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `ca_certificate` (String) Certificate Authority in PKI configuration

    |Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    |txt     &emsp;|Name of CA in PKI configuration  |
- `certificate` (String) Certificate in PKI configuration

    |Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    |txt     &emsp;|Name of certificate in PKI configuration  |
- `dh_params` (String) Diffie Hellman parameters (server only)
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
