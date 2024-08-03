---
page_title: "vyos_load_balancing_reverse_proxy_backend_server Resource - vyos"

subcategory: "Load"

description: |- 
  load-balancing⯯Configure reverse-proxy⯯Backend server name⯯Backend server name
---

# vyos_load_balancing_reverse_proxy_backend_server (Resource)
<center>

*load-balancing*  
⯯  
Configure reverse-proxy  
⯯  
Backend server name  
⯯  
**Backend server name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (String) Backend server address

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|ipv4    &emsp;|IPv4 unicast peer address  |
    &emsp;|ipv6    &emsp;|IPv6 unicast peer address  |
- `backup` (Boolean) Use backup server if other servers are not available
- `check` (Boolean) Active health check backend server
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `send_proxy` (Boolean) Send a Proxy Protocol version 1 header (text format)
- `send_proxy_v2` (Boolean) Send a Proxy Protocol version 2 header (binary format)
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `backend` (String) Backend server name
- `server` (String) Backend server name


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
