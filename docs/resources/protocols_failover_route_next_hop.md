---
page_title: "vyos_protocols_failover_route_next_hop Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Failover Routing⯯Failover IPv4 route⯯Next-hop IPv4 router address
---

# vyos_protocols_failover_route_next_hop (Resource)
<center>

*protocols*  
⯯  
Failover Routing  
⯯  
Failover IPv4 route  
⯯  
**Next-hop IPv4 router address**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `check` (Attributes) Check target options (see [below for nested schema](#nestedatt--check))
- `interface` (String) Gateway interface name

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Gateway interface name  |
- `metric` (Number) Route metric for this gateway

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|1-255   &emsp;|Route metric  |
- `onlink` (Boolean) The next hop is directly connected to the interface, even if it does not match interface prefix
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `next_hop` (String) Next-hop IPv4 router address

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|ipv4    &emsp;|Next-hop router address  |
- `route` (String) Failover IPv4 route

    &emsp;|Format   &emsp;|Description          |
    |-----------|-----------------------|
    &emsp;|ipv4net  &emsp;|IPv4 failover route  |


&lt;a id=&#34;nestedatt--check&#34;&gt;&lt;/a&gt;
### Nested Schema for `check`

Optional:

- `policy` (String) Policy for check targets

    &emsp;|Format         &emsp;|Description                |
    |-----------------|-----------------------------|
    &emsp;|all-available  &emsp;|All targets must be alive  |
    &emsp;|any-available  &emsp;|Any target must be alive   |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `target` (List of String) Check target address

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|ipv4    &emsp;|Address to check  |
- `timeout` (Number) Timeout between checks

    &emsp;|Format  &emsp;|Description                        |
    |----------|-------------------------------------|
    &emsp;|1-300   &emsp;|Timeout in seconds between checks  |
- `type` (String) Check type

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|arp     &emsp;|Check target by ARP   |
    &emsp;|icmp    &emsp;|Check target by ICMP  |
    &emsp;|tcp     &emsp;|Check target by TCP   |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  