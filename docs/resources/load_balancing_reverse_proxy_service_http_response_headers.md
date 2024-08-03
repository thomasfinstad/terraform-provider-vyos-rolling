---
page_title: "vyos_load_balancing_reverse_proxy_service_http_response_headers Resource - vyos"

subcategory: "Load"

description: |- 
  load-balancing⯯Configure reverse-proxy⯯Frontend service name⯯Headers to include in HTTP response
---

# vyos_load_balancing_reverse_proxy_service_http_response_headers (Resource)
<center>

*load-balancing*  
⯯  
Configure reverse-proxy  
⯯  
Frontend service name  
⯯  
**Headers to include in HTTP response**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `value` (String) HTTP header value

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|HTTP header value  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `http_response_headers` (String) Headers to include in HTTP response

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|HTTP header name  |
- `service` (String) Frontend service name


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
