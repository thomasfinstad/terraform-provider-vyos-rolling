---
page_title: "vyos_service_https_api_graphql_authentication Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯HTTPS configuration⯯VyOS HTTP API configuration⯯GraphQL support⯯GraphQL authentication
---

# vyos_service_https_api_graphql_authentication (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
HTTPS configuration  
⯯  
VyOS HTTP API configuration  
⯯  
GraphQL support  
⯯  
**GraphQL authentication**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `expiration` (Number) Token time to expire in seconds

    &emsp;|Format       &emsp;|Description                |
    |---------------|-----------------------------|
    &emsp;|60-31536000  &emsp;|Token lifetime in seconds  |
- `secret_length` (Number) Length of shared secret in bytes

    &emsp;|Format    &emsp;|Description                             |
    |------------|------------------------------------------|
    &emsp;|16-65535  &emsp;|Byte length of generated shared secret  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `type` (String) Authentication type

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|key     &emsp;|Use API keys   |
    &emsp;|token   &emsp;|Use JWT token  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  