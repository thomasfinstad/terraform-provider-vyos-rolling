---
page_title: "vyos_service_stunnel_client Resource - vyos"

subcategory: "Service"

description: |- 
  System services⯯Stunnel TLS Proxy⯯Stunnel client config
---

# vyos_service_stunnel_client (Resource)
<center>

System services  
⯯  
Stunnel TLS Proxy  
⯯  
**Stunnel client config**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `connect` (Attributes) Connect to a remote address (see [below for nested schema](#nestedatt--connect))
- `listen` (Attributes) Accept connections on specified address (see [below for nested schema](#nestedatt--listen))
- `options` (Attributes) Advanced protocol options (see [below for nested schema](#nestedatt--options))
- `protocol` (String) Application protocol to negotiate TLS

    &emsp;|Format   &emsp;|Description                                                                                           |
    |-----------|--------------------------------------------------------------------------------------------------------|
    &emsp;|cifs     &emsp;|Proprietary (undocummented) extension of CIFS protocol                                                |
    &emsp;|connect  &emsp;|Based on RFC 2817 - Upgrading to TLS Within HTTP/1.1, section 5.2 - Requesting a Tunnel with CONNECT  |
    &emsp;|imap     &emsp;|Based on RFC 2595 - Using TLS with IMAP, POP3 and ACAP                                                |
    &emsp;|nntp     &emsp;|Based on RFC 4642 - Using Transport Layer Security (TLS) with Network News Transfer Protocol (NNTP)   |
    &emsp;|pgsql    &emsp;|Based on PostgreSQL frontend/backend protocol                                                         |
    &emsp;|pop3     &emsp;|Based on RFC 2449 - POP3 Extension Mechanism                                                          |
    &emsp;|proxy    &emsp;|Passing of the original client IP address with HAProxy PROXY protocol version 1                       |
    &emsp;|smtp     &emsp;|Based on RFC 2487 - SMTP Service Extension for Secure SMTP over TLS                                   |
    &emsp;|socks    &emsp;|SOCKS versions 4, 4a, and 5 are supported                                                             |
- `ssl` (Attributes) SSL Certificate, SSL Key and CA (see [below for nested schema](#nestedatt--ssl))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `client` (String) Stunnel client config


&lt;a id=&#34;nestedatt--connect&#34;&gt;&lt;/a&gt;
### Nested Schema for `connect`

Optional:

- `address` (String) Hostname or IP address

    &emsp;|Format    &emsp;|Description   |
    |------------|----------------|
    &emsp;|ipv4      &emsp;|IPv4 address  |
    &emsp;|hostname  &emsp;|hostname      |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |


&lt;a id=&#34;nestedatt--listen&#34;&gt;&lt;/a&gt;
### Nested Schema for `listen`

Optional:

- `address` (String) Hostname or IP address

    &emsp;|Format    &emsp;|Description   |
    |------------|----------------|
    &emsp;|ipv4      &emsp;|IPv4 address  |
    &emsp;|hostname  &emsp;|hostname      |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |


&lt;a id=&#34;nestedatt--options&#34;&gt;&lt;/a&gt;
### Nested Schema for `options`

Optional:

- `authentication` (String) Authentication type for the protocol negotiations

    &emsp;|Format  &emsp;|Description                                                |
    |----------|-------------------------------------------------------------|
    &emsp;|basic   &emsp;|The default &#39;connect&#39; authentication type                  |
    &emsp;|ntlm    &emsp;|Supported authentication types for the &#39;connect&#39; protocol  |
    &emsp;|plain   &emsp;|The default &#39;smtp&#39; authentication type                     |
    &emsp;|login   &emsp;|Supported authentication types for the &#39;smtp&#39; protocol     |
- `domain` (String) Domain for the &#39;connect&#39; protocol.

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|domain  &emsp;|domain       |
- `host` (Attributes) Destination address for the &#39;connect&#39; protocol (see [below for nested schema](#nestedatt--options--host))
- `password` (String) Password for the protocol negotiations

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|txt     &emsp;|Authentication password  |
- `username` (String) Username for the protocol negotiations

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|txt     &emsp;|Authentication username  |

&lt;a id=&#34;nestedatt--options--host&#34;&gt;&lt;/a&gt;
### Nested Schema for `options.host`

Optional:

- `address` (String) Hostname or IP address

    &emsp;|Format    &emsp;|Description   |
    |------------|----------------|
    &emsp;|ipv4      &emsp;|IPv4 address  |
    &emsp;|hostname  &emsp;|hostname      |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |



&lt;a id=&#34;nestedatt--ssl&#34;&gt;&lt;/a&gt;
### Nested Schema for `ssl`

Optional:

- `ca_certificate` (List of String) Certificate Authority chain in PKI configuration

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|Name of CA in PKI configuration  |
- `certificate` (String) Certificate in PKI configuration

    &emsp;|Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    &emsp;|txt     &emsp;|Name of certificate in PKI configuration  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
