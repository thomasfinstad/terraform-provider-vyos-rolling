---
page_title: "vyos_service_webproxy_authentication_ldap Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Webproxy service settings⯯Proxy Authentication Settings⯯LDAP authentication settings
---

# vyos_service_webproxy_authentication_ldap (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Webproxy service settings  
⯯  
Proxy Authentication Settings  
⯯  
**LDAP authentication settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `base_dn` (String) LDAP Base DN to search
- `bind_dn` (String) LDAP DN used to bind to server
- `filter_expression` (String) Filter expression to perform LDAP search with
- `password` (String) LDAP password to bind with
- `persistent_connection` (Boolean) Use persistent LDAP connection
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `server` (String) LDAP server to use
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `use_ssl` (Boolean) Use SSL/TLS for LDAP connection
- `username_attribute` (String) LDAP username attribute
- `version` (String) LDAP protocol version

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|2       &emsp;|LDAP protocol version 2  |
    &emsp;|3       &emsp;|LDAP protocol version 2  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
