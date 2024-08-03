---
page_title: "vyos_service_dns_dynamic_name Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Domain Name System (DNS) related services⯯Dynamic DNS⯯Dynamic DNS configuration
---

# vyos_service_dns_dynamic_name (Resource)
<center>

*service*  
⯯  
Domain Name System (DNS) related services  
⯯  
Dynamic DNS  
⯯  
**Dynamic DNS configuration**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (Attributes) Obtain IP address to send Dynamic DNS update for (see [below for nested schema](#nestedatt--address))
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `expiry_time` (Number) Time in seconds for the hostname to be marked expired in cache

    &emsp;|Format       &emsp;|Description      |
    |---------------|-------------------|
    &emsp;|300-2160000  &emsp;|Time in seconds  |
- `host_name` (List of String) Hostname to register with Dynamic DNS service
- `ip_version` (String) IP address version to use

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|_ipv4   &emsp;|Use only IPv4 address           |
    &emsp;|_ipv6   &emsp;|Use only IPv6 address           |
    &emsp;|both    &emsp;|Use both IPv4 and IPv6 address  |
- `key` (String) File containing TSIG authentication key for RFC2136 nsupdate on remote DNS server

    &emsp;|Format    &emsp;|Description                     |
    |------------|----------------------------------|
    &emsp;|filename  &emsp;|File in /config/auth directory  |
- `password` (String) Password used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Password     |
- `protocol` (String) ddclient protocol used for Dynamic DNS service
- `server` (String) Remote Dynamic DNS server to send updates to

    &emsp;|Format    &emsp;|Description                                       |
    |------------|----------------------------------------------------|
    &emsp;|ipv4      &emsp;|IPv4 address of the remote server                 |
    &emsp;|ipv6      &emsp;|IPv6 address of the remote server                 |
    &emsp;|hostname  &emsp;|Fully qualified domain name of the remote server  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `ttl` (Number) Time-to-live (TTL)

    &emsp;|Format        &emsp;|Description     |
    |----------------|------------------|
    &emsp;|0-2147483647  &emsp;|TTL in seconds  |
- `username` (String) Username used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Username     |
- `wait_time` (Number) Time in seconds to wait between update attempts

    &emsp;|Format    &emsp;|Description      |
    |------------|-------------------|
    &emsp;|60-86400  &emsp;|Time in seconds  |
- `zone` (String) DNS zone to be updated

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|Name of DNS zone  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `name` (String) Dynamic DNS configuration

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Dynamic DNS service name  |


&lt;a id=&#34;nestedatt--address&#34;&gt;&lt;/a&gt;
### Nested Schema for `address`

Optional:

- `interface` (String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `web` (Attributes) HTTP(S) web request to use (see [below for nested schema](#nestedatt--address--web))

&lt;a id=&#34;nestedatt--address--web&#34;&gt;&lt;/a&gt;
### Nested Schema for `address.web`

Optional:

- `skip` (String) Pattern to skip from the HTTP(S) respose

    &emsp;|Format  &emsp;|Description                                                                  |
    |----------|-------------------------------------------------------------------------------|
    &emsp;|txt     &emsp;|Pattern to skip from the HTTP(S) respose to extract the external IP address  |
- `url` (String) Remote URL

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|url     &emsp;|Remote HTTP(S) URL  |



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
