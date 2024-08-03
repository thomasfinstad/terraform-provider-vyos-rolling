---
page_title: "vyos_load_balancing_reverse_proxy_service Resource - vyos"

subcategory: "Load"

description: |- 
  load-balancing⯯Configure reverse-proxy⯯Frontend service name
---

# vyos_load_balancing_reverse_proxy_service (Resource)
<center>

*load-balancing*  
⯯  
Configure reverse-proxy  
⯯  
**Frontend service name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `backend` (List of String) Backend member

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|txt     &emsp;|Name of reverse-proxy backend system  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `listen_address` (List of String) Local IP addresses to listen on

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address to listen for incoming connections  |
    &emsp;|ipv6    &emsp;|IPv6 address to listen for incoming connections  |
- `logging` (Attributes) Logging parameters (see [below for nested schema](#nestedatt--logging))
- `mode` (String) Proxy mode

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|http    &emsp;|HTTP proxy mode  |
    &emsp;|tcp     &emsp;|TCP proxy mode   |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `redirect_http_to_https` (Boolean) Redirect HTTP to HTTPS
- `ssl` (Attributes) SSL Certificate, SSL Key and CA (see [below for nested schema](#nestedatt--ssl))
- `tcp_request` (Attributes) TCP request directive (see [below for nested schema](#nestedatt--tcp_request))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `service` (String) Frontend service name


&lt;a id=&#34;nestedatt--logging&#34;&gt;&lt;/a&gt;
### Nested Schema for `logging`


&lt;a id=&#34;nestedatt--ssl&#34;&gt;&lt;/a&gt;
### Nested Schema for `ssl`

Optional:

- `certificate` (List of String) Certificate in PKI configuration

    &emsp;|Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    &emsp;|txt     &emsp;|Name of certificate in PKI configuration  |


&lt;a id=&#34;nestedatt--tcp_request&#34;&gt;&lt;/a&gt;
### Nested Schema for `tcp_request`

Optional:

- `inspect_delay` (Number) Set the maximum allowed time to wait for data during content inspection

    &emsp;|Format   &emsp;|Description                                  |
    |-----------|-----------------------------------------------|
    &emsp;|1-65535  &emsp;|The timeout value specified in milliseconds  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
