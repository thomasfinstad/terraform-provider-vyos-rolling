---
page_title: "vyos_system_flow_accounting_netflow_timeout Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯Flow accounting settings⯯NetFlow settings⯯NetFlow timeout values
---

# vyos_system_flow_accounting_netflow_timeout (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
Flow accounting settings  
⯯  
NetFlow settings  
⯯  
**NetFlow timeout values**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `expiry_interval` (Number) Expiry scan interval

    |Format        &emsp;|Description           |
    |----------------|------------------------|
    |0-2147483647  &emsp;|Expiry scan interval  |
- `flow_generic` (Number) Generic flow timeout value

    |Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    |0-2147483647  &emsp;|Generic flow timeout in seconds  |
- `icmp` (Number) ICMP timeout value

    |Format        &emsp;|Description              |
    |----------------|---------------------------|
    |0-2147483647  &emsp;|ICMP timeout in seconds  |
- `max_active_life` (Number) Max active timeout value

    |Format        &emsp;|Description                    |
    |----------------|---------------------------------|
    |0-2147483647  &emsp;|Max active timeout in seconds  |
- `tcp_fin` (Number) TCP finish timeout value

    |Format        &emsp;|Description                 |
    |----------------|------------------------------|
    |0-2147483647  &emsp;|TCP FIN timeout in seconds  |
- `tcp_generic` (Number) TCP generic timeout value

    |Format        &emsp;|Description                     |
    |----------------|----------------------------------|
    |0-2147483647  &emsp;|TCP generic timeout in seconds  |
- `tcp_rst` (Number) TCP reset timeout value

    |Format        &emsp;|Description                 |
    |----------------|------------------------------|
    |0-2147483647  &emsp;|TCP RST timeout in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `udp` (Number) UDP timeout value

    |Format        &emsp;|Description             |
    |----------------|--------------------------|
    |0-2147483647  &emsp;|UDP timeout in seconds  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
