---
page_title: "vyos_system_flow_accounting_sflow_server Resource - vyos"

subcategory: "System"

description: |- 
  system⯯Flow accounting settings⯯sFlow settings⯯sFlow destination server
---

# vyos_system_flow_accounting_sflow_server (Resource)
<center>

*system*  
⯯  
Flow accounting settings  
⯯  
sFlow settings  
⯯  
**sFlow destination server**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `port` (Number) sFlow port number

    &emsp;|Format      &emsp;|Description        |
    |--------------|---------------------|
    &emsp;|1025-65535  &emsp;|sFlow port number  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `server` (String) sFlow destination server

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 server to export sFlow  |
    &emsp;|ipv6    &emsp;|IPv6 server to export sFlow  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
