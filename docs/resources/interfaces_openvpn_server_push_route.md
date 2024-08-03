---
page_title: "vyos_interfaces_openvpn_server_push_route Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯OpenVPN Tunnel Interface⯯Server-mode options⯯Route to be pushed to all clients
---

# vyos_interfaces_openvpn_server_push_route (Resource)
<center>

*interfaces*  
⯯  
OpenVPN Tunnel Interface  
⯯  
Server-mode options  
⯯  
**Route to be pushed to all clients**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `metric` (Number) Set metric for this route

    &emsp;|Format        &emsp;|Description            |
    |----------------|-------------------------|
    &emsp;|0-4294967295  &emsp;|Metric for this route  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `openvpn` (String) OpenVPN Tunnel Interface

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|vtunN   &emsp;|OpenVPN interface name  |
- `push_route` (String) Route to be pushed to all clients

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 network and prefix length  |
    &emsp;|ipv6net  &emsp;|IPv6 network and prefix length  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
