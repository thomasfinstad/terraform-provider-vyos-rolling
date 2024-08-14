---
page_title: "vyos_system_flow_accounting_netflow_server Resource - vyos"

subcategory: "System"

description: |- 
  system⯯Flow accounting settings⯯NetFlow settings⯯NetFlow destination server
---

# vyos_system_flow_accounting_netflow_server (Resource)
<center>

*system*  
⯯  
Flow accounting settings  
⯯  
NetFlow settings  
⯯  
**NetFlow destination server**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `port` (Number) NetFlow port number

    |Format      &emsp;|Description          |
    |--------------|-----------------------|
    |1025-65535  &emsp;|NetFlow port number  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `server` (String) NetFlow destination server

    |Format  &emsp;|Description                    |
    |----------|---------------------------------|
    |ipv4    &emsp;|IPv4 server to export NetFlow  |
    |ipv6    &emsp;|IPv6 server to export NetFlow  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
