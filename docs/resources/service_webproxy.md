---
page_title: "vyos_service_webproxy Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Webproxy service settings
---

# vyos_service_webproxy (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
**Webproxy service settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `append_domain` (String) Default domain name

    |Format  &emsp;|Description                                       |
    |----------|----------------------------------------------------|
    |domain  &emsp;|Domain to use for urls that do not contain a &#39;.&#39;  |
- `cache_size` (String) Disk cache size in MB

    |Format  &emsp;|Description            |
    |----------|-------------------------|
    |u32     &emsp;|Disk cache size in MB  |
    |0       &emsp;|Disable disk caching   |
- `default_port` (Number) Default Proxy Port

    |Format      &emsp;|Description          |
    |--------------|-----------------------|
    |1025-65535  &emsp;|Default port number  |
- `disable_access_log` (Boolean) Disable logging of HTTP accesses
- `domain_block` (List of String) Domain name to block
- `domain_noncache` (List of String) Domain name to access without caching
- `maximum_object_size` (Number) Maximum size of object to be stored in cache in kilobytes

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |u32     &emsp;|Object size in KB  |
- `mem_cache_size` (Number) Memory cache size in MB

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |u32     &emsp;|Memory cache size in MB   |
- `minimum_object_size` (Number) Maximum size of object to be stored in cache in kilobytes

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |u32     &emsp;|Object size in KB  |
- `outgoing_address` (String) Outgoing IP address for webproxy
- `reply_block_mime` (List of String) MIME type to block
- `reply_body_max_size` (Number) Maximum reply body size in KB

    |Format  &emsp;|Description       |
    |----------|--------------------|
    |u32     &emsp;|Reply size in KB  |
- `safe_ports` (List of Number) Safe port ACL

    |Format  &emsp;|Description                                                                              |
    |----------|-------------------------------------------------------------------------------------------|
    |1-1024  &emsp;|Port number. Ports included by default: 21,70,80,210,280,443,488,591,777,873,1025-65535  |
- `ssl_safe_ports` (List of Number) SSL safe port

    |Format   &emsp;|Description                                  |
    |-----------|-----------------------------------------------|
    |1-65535  &emsp;|Port number. Ports included by default: 443  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
