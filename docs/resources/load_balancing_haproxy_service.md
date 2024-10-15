---
page_title: "vyos_load_balancing_haproxy_service Resource - vyos"

subcategory: "Load Balancing"

description: |-
  load-balancing⯯Configure haproxy⯯Frontend service name
---

# vyos_load_balancing_haproxy_service (Resource)
<center>

*load-balancing*  
⯯  
Configure haproxy  
⯯  
**Frontend service name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_load_balancing_haproxy_service (Resource)](#vyos_load_balancing_haproxy_service-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [backend](#backend)
      - [description](#description)
      - [listen_address](#listen_address)
      - [logging](#logging)
      - [mode](#mode)
      - [port](#port)
      - [redirect_http_to_https](#redirect_http_to_https)
      - [ssl](#ssl)
      - [tcp_request](#tcp_request)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `logging`](#nested-schema-for-logging)
    - [Nested Schema for `ssl`](#nested-schema-for-ssl)
    - [Nested Schema for `tcp_request`](#nested-schema-for-tcp_request)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### backend
- `backend` (List of String) Backend member

    |  Format  &emsp;|  Description                     |
    |----------|----------------------------------|
    |  txt     &emsp;|  Name of haproxy backend system  |
#### description
- `description` (String) Description

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  txt     &emsp;|  Description  |
#### listen_address
- `listen_address` (List of String) Local IP addresses to listen on

    |  Format  &emsp;|  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    &emsp;|  IPv4 address to listen for incoming connections  |
    |  ipv6    &emsp;|  IPv6 address to listen for incoming connections  |
#### logging
- `logging` (Attributes) Logging parameters (see [below for nested schema](#nestedatt--logging))
#### mode
- `mode` (String) Proxy mode

    |  Format  &emsp;|  Description      |
    |----------|-------------------|
    |  http    &emsp;|  HTTP proxy mode  |
    |  tcp     &emsp;|  TCP proxy mode   |
#### port
- `port` (Number) Port number used by connection

    |  Format   &emsp;|  Description      |
    |-----------|-------------------|
    |  1-65535  &emsp;|  Numeric IP port  |
#### redirect_http_to_https
- `redirect_http_to_https` (Boolean) Redirect HTTP to HTTPS
#### ssl
- `ssl` (Attributes) SSL Certificate, SSL Key and CA (see [below for nested schema](#nestedatt--ssl))
#### tcp_request
- `tcp_request` (Attributes) TCP request directive (see [below for nested schema](#nestedatt--tcp_request))
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `service` (String) Frontend service name


<a id="nestedatt--logging"></a>
### Nested Schema for `logging`


<a id="nestedatt--ssl"></a>
### Nested Schema for `ssl`

Optional:

- `certificate` (List of String) Certificate in PKI configuration

    |  Format  &emsp;|  Description                               |
    |----------|--------------------------------------------|
    |  txt     &emsp;|  Name of certificate in PKI configuration  |


<a id="nestedatt--tcp_request"></a>
### Nested Schema for `tcp_request`

Optional:

- `inspect_delay` (Number) Set the maximum allowed time to wait for data during content inspection

    |  Format   &emsp;|  Description                                  |
    |-----------|-----------------------------------------------|
    |  1-65535  &emsp;|  The timeout value specified in milliseconds  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).