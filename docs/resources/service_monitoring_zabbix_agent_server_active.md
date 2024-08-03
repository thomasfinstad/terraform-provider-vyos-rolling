---
page_title: "vyos_service_monitoring_zabbix_agent_server_active Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯monitoring⯯Zabbix-agent settings⯯Remote server address to get active checks from
---

# vyos_service_monitoring_zabbix_agent_server_active (Resource)
<center>

*service*  
⯯  
*monitoring*  
⯯  
Zabbix-agent settings  
⯯  
**Remote server address to get active checks from**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `server_active` (String) Remote server address to get active checks from

    &emsp;|Format    &emsp;|Description           |
    |------------|------------------------|
    &emsp;|ipv4      &emsp;|Server IPv4 address   |
    &emsp;|ipv6      &emsp;|Server IPv6 address   |
    &emsp;|hostname  &emsp;|Server hostname/FQDN  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
