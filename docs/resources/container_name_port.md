---
page_title: "vyos_container_name_port Resource - vyos"

subcategory: "Container"

description: |- 
  Container applications⯯Container name⯯Publish port to the container
---

# vyos_container_name_port (Resource)
<center>

Container applications  
⯯  
Container name  
⯯  
**Publish port to the container**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `destination` (String) Destination container port

    &emsp;|Format     &emsp;|Description                                          |
    |-------------|-------------------------------------------------------|
    &emsp;|1-65535    &emsp;|Destination container port                           |
    &emsp;|start-end  &emsp;|Destination container port range (e.g. 10025-10030)  |
- `listen_address` (List of String) Local IP addresses to listen on

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address to listen for incoming connections  |
    &emsp;|ipv6    &emsp;|IPv6 address to listen for incoming connections  |
- `protocol` (String) Transport protocol used for port mapping

    &emsp;|Format  &emsp;|Description                                       |
    |----------|----------------------------------------------------|
    &emsp;|tcp     &emsp;|Use Transmission Control Protocol for given port  |
    &emsp;|udp     &emsp;|Use User Datagram Protocol for given port         |
- `source` (String) Source host port

    &emsp;|Format     &emsp;|Description                                |
    |-------------|---------------------------------------------|
    &emsp;|1-65535    &emsp;|Source host port                           |
    &emsp;|start-end  &emsp;|Source host port range (e.g. 10025-10030)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `name` (String) Container name
- `port` (String) Publish port to the container


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
