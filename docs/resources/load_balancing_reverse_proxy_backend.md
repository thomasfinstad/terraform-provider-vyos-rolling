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

    &emsp;|Format            &emsp;|Description                         |
    |--------------------|--------------------------------------|
    &emsp;|source-address    &emsp;|Based on hash of source IP address  |
    &emsp;|round-robin       &emsp;|Round robin                         |
    &emsp;|least-connection  &emsp;|Least connection                    |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `health_check` (String) Non HTTP health check options

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|ldap    &emsp;|LDAP protocol check        |
    &emsp;|mysql   &emsp;|MySQL protocol check       |
    &emsp;|pgsql   &emsp;|PostgreSQL protocol check  |
    &emsp;|redis   &emsp;|Redis protocol check       |
    &emsp;|smtp    &emsp;|SMTP protocol check        |
- `http_check` (Attributes) HTTP check configuration (see [below for nested schema](#nestedatt--http_check))
- `logging` (Attributes) Logging parameters (see [below for nested schema](#nestedatt--logging))
- `mode` (String) Proxy mode

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|http    &emsp;|HTTP proxy mode  |
    &emsp;|tcp     &emsp;|TCP proxy mode   |
- `ssl` (Attributes) SSL Certificate, SSL Key and CA (see [below for nested schema](#nestedatt--ssl))
- `timeout` (Attributes) Timeout options (see [below for nested schema](#nestedatt--timeout))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `backend` (String) Backend server name


&lt;a id=&#34;nestedatt--http_check&#34;&gt;&lt;/a&gt;
### Nested Schema for `http_check`

Optional:

- `expect` (Attributes) Expected response for the health check to pass (see [below for nested schema](#nestedatt--http_check--expect))
- `method` (String) HTTP method used for health check

    &emsp;|Format                     &emsp;|Description                           |
    |-----------------------------|----------------------------------------|
    &emsp;|options|head|get|post|put  &emsp;|HTTP method used for health checking  |
- `uri` (String) URI used for HTTP health check (Example: &#39;/&#39; or &#39;/health&#39;)

&lt;a id=&#34;nestedatt--http_check--expect&#34;&gt;&lt;/a&gt;
### Nested Schema for `http_check.expect`

Optional:

- `status` (Number) Expected response status code for the health check to pass

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|200-399  &emsp;|Expected response code  |
- `string` (String) Expected to be in response body for the health check to pass

    &emsp;|Format  &emsp;|Description                              |
    |----------|-------------------------------------------|
    &emsp;|txt     &emsp;|A string expected to be in the response  |



&lt;a id=&#34;nestedatt--logging&#34;&gt;&lt;/a&gt;
### Nested Schema for `logging`


&lt;a id=&#34;nestedatt--ssl&#34;&gt;&lt;/a&gt;
### Nested Schema for `ssl`

Optional:

- `ca_certificate` (String) Certificate Authority in PKI configuration

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|Name of CA in PKI configuration  |
- `no_verify` (Boolean) Do not attempt to verify SSL certificates for backend servers


&lt;a id=&#34;nestedatt--timeout&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeout`

Optional:

- `check` (Number) Timeout in seconds for established connections

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|1-3600  &emsp;|Check timeout in seconds  |
- `connect` (Number) Set the maximum time to wait for a connection attempt to a server to succeed

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|1-3600  &emsp;|Connect timeout in seconds  |
- `server` (Number) Set the maximum inactivity time on the server side

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|1-3600  &emsp;|Server timeout in seconds  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
