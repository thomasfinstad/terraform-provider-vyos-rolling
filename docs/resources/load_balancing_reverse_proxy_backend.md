---
page_title: "vyos_load_balancing_reverse_proxy_backend Resource - vyos"

subcategory: "Load"

description: |- 
  load-balancing⯯Configure reverse-proxy⯯Backend server name
---

# vyos_load_balancing_reverse_proxy_backend (Resource)
<center>

*load-balancing*  
⯯  
Configure reverse-proxy  
⯯  
**Backend server name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `balance` (String) Load-balancing algorithm

    |Format            &emsp;|Description                         |
    |--------------------|--------------------------------------|
    |source-address    &emsp;|Based on hash of source IP address  |
    |round-robin       &emsp;|Round robin                         |
    |least-connection  &emsp;|Least connection                    |
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `health_check` (String) Non HTTP health check options

    |Format  &emsp;|Description                |
    |----------|-----------------------------|
    |ldap    &emsp;|LDAP protocol check        |
    |mysql   &emsp;|MySQL protocol check       |
    |pgsql   &emsp;|PostgreSQL protocol check  |
    |redis   &emsp;|Redis protocol check       |
    |smtp    &emsp;|SMTP protocol check        |
- `http_check` (Attributes) HTTP check configuration (see [below for nested schema](#nestedatt--http_check))
- `logging` (Attributes) Logging parameters (see [below for nested schema](#nestedatt--logging))
- `mode` (String) Proxy mode

    |Format  &emsp;|Description      |
    |----------|-------------------|
    |http    &emsp;|HTTP proxy mode  |
    |tcp     &emsp;|TCP proxy mode   |
- `ssl` (Attributes) SSL Certificate, SSL Key and CA (see [below for nested schema](#nestedatt--ssl))
- `timeout` (Attributes) Timeout options (see [below for nested schema](#nestedatt--timeout))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `backend` (String) Backend server name


<a id="nestedatt--http_check"></a>
### Nested Schema for `http_check`

Optional:

- `expect` (Attributes) Expected response for the health check to pass (see [below for nested schema](#nestedatt--http_check--expect))
- `method` (String) HTTP method used for health check

    |Format                     &emsp;|Description                           |
    |-----------------------------|----------------------------------------|
    |options|head|get|post|put  &emsp;|HTTP method used for health checking  |
- `uri` (String) URI used for HTTP health check (Example: &#39;/&#39; or &#39;/health&#39;)

<a id="nestedatt--http_check--expect"></a>
### Nested Schema for `http_check.expect`

Optional:

- `status` (Number) Expected response status code for the health check to pass

    |Format   &emsp;|Description             |
    |-----------|--------------------------|
    |200-399  &emsp;|Expected response code  |
- `string` (String) Expected to be in response body for the health check to pass

    |Format  &emsp;|Description                              |
    |----------|-------------------------------------------|
    |txt     &emsp;|A string expected to be in the response  |



<a id="nestedatt--logging"></a>
### Nested Schema for `logging`


<a id="nestedatt--ssl"></a>
### Nested Schema for `ssl`

Optional:

- `ca_certificate` (String) Certificate Authority in PKI configuration

    |Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    |txt     &emsp;|Name of CA in PKI configuration  |
- `no_verify` (Boolean) Do not attempt to verify SSL certificates for backend servers


<a id="nestedatt--timeout"></a>
### Nested Schema for `timeout`

Optional:

- `check` (Number) Timeout in seconds for established connections

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |1-3600  &emsp;|Check timeout in seconds  |
- `connect` (Number) Set the maximum time to wait for a connection attempt to a server to succeed

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |1-3600  &emsp;|Connect timeout in seconds  |
- `server` (Number) Set the maximum inactivity time on the server side

    |Format  &emsp;|Description                |
    |----------|-----------------------------|
    |1-3600  &emsp;|Server timeout in seconds  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
