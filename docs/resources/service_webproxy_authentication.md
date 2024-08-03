---
page_title: "vyos_service_webproxy_authentication Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Webproxy service settings⯯Proxy Authentication Settings
---

# vyos_service_webproxy_authentication (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Webproxy service settings  
⯯  
**Proxy Authentication Settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `children` (String) Number of authentication helper processes

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|n       &emsp;|Number of authentication helper processes  |
- `credentials_ttl` (String) Authenticated session time to live in minutes

    &emsp;|Format  &emsp;|Description                    |
    |----------|---------------------------------|
    &emsp;|n       &emsp;|Authenticated session timeout  |
- `method` (String) Authentication Method

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|ldap    &emsp;|Lightweight Directory Access Protocol  |
- `realm` (String) Name of authentication realm (e.g. &#34;My Company proxy server&#34;)
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
