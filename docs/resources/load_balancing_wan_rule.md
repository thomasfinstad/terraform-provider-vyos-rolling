---
page_title: "vyos_load_balancing_wan_rule Resource - vyos"

subcategory: "Load"

description: |- 
  Configure load-balancing⯯Configure Wide Area Network (WAN) load-balancing⯯Rule number (1-9999)
---

# vyos_load_balancing_wan_rule (Resource)
<center>

Configure load-balancing  
⯯  
Configure Wide Area Network (WAN) load-balancing  
⯯  
**Rule number (1-9999)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `destination` (Attributes) Destination (see [below for nested schema](#nestedatt--destination))
- `exclude` (Boolean) Exclude packets matching this rule from WAN load balance
- `failover` (Boolean) Enable failover for packets matching this rule from WAN load balance
- `inbound_interface` (String) Inbound interface name (e.g., &#34;eth0&#34;) [REQUIRED]
- `limit` (Attributes) Enable packet limit for this rule (see [below for nested schema](#nestedatt--limit))
- `per_packet_balancing` (Boolean) Option to match traffic per-packet instead of the default, per-flow
- `protocol` (String) Protocol to match (protocol name, number, or &#34;all&#34;)

    &emsp;|Format       &emsp;|Description         |
    |---------------|----------------------|
    &emsp;|all          &emsp;|All IP protocols    |
    &emsp;|tcp_udp      &emsp;|Both TCP and UDP    |
    &emsp;|0-255        &emsp;|IP protocol number  |
    &emsp;|&lt;protocol&gt;   &emsp;|IP protocol name    |
    &emsp;|!&lt;protocol&gt;  &emsp;|IP protocol name    |
- `source` (Attributes) Source information (see [below for nested schema](#nestedatt--source))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `rule` (Number) Rule number (1-9999)

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1-9999  &emsp;|Rule number  |


&lt;a id=&#34;nestedatt--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination`

Optional:

- `address` (String) IP address, subnet, or range

    &emsp;|Format      &emsp;|Description                                    |
    |--------------|-------------------------------------------------|
    &emsp;|ipv4        &emsp;|IPv4 address to match                          |
    &emsp;|ipv4net     &emsp;|IPv4 prefix to match                           |
    &emsp;|ipv4range   &emsp;|IPv4 address range to match                    |
    &emsp;|!ipv4       &emsp;|Match everything except the specified address  |
    &emsp;|!ipv4net    &emsp;|Match everything except the specified prefix   |
    &emsp;|!ipv4range  &emsp;|Match everything except the specified range    |
- `port` (String) Port number

    &emsp;|Format     &emsp;|Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt        &emsp;|Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    &emsp;|1-65535    &emsp;|Numeric IP port                                                                                                                                                          |
    &emsp;|start-end  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    &emsp;|           &emsp;|\n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using &#39;!&#39;.\nFor example: &#39;!22,telnet,http,123,1001-1005&#39;  |


&lt;a id=&#34;nestedatt--limit&#34;&gt;&lt;/a&gt;
### Nested Schema for `limit`

Optional:

- `burst` (Number) Burst limit for matching packets

    &emsp;|Format        &emsp;|Description                       |
    |----------------|------------------------------------|
    &emsp;|0-4294967295  &emsp;|Burst limit for matching packets  |
- `period` (String) Time window for rate calculation

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|hour    &emsp;|hour         |
    &emsp;|minute  &emsp;|minute       |
    &emsp;|second  &emsp;|second       |
- `rate` (Number) Number of packets used for rate limit

    &emsp;|Format        &emsp;|Description                            |
    |----------------|-----------------------------------------|
    &emsp;|0-4294967295  &emsp;|Number of packets used for rate limit  |
- `threshold` (String) Threshold behavior for limit

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|above   &emsp;|Above limit  |
    &emsp;|below   &emsp;|Below limit  |


&lt;a id=&#34;nestedatt--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `source`

Optional:

- `address` (String) IP address, subnet, or range

    &emsp;|Format      &emsp;|Description                                    |
    |--------------|-------------------------------------------------|
    &emsp;|ipv4        &emsp;|IPv4 address to match                          |
    &emsp;|ipv4net     &emsp;|IPv4 prefix to match                           |
    &emsp;|ipv4range   &emsp;|IPv4 address range to match                    |
    &emsp;|!ipv4       &emsp;|Match everything except the specified address  |
    &emsp;|!ipv4net    &emsp;|Match everything except the specified prefix   |
    &emsp;|!ipv4range  &emsp;|Match everything except the specified range    |
- `port` (String) Port number

    &emsp;|Format     &emsp;|Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt        &emsp;|Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    &emsp;|1-65535    &emsp;|Numeric IP port                                                                                                                                                          |
    &emsp;|start-end  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    &emsp;|           &emsp;|\n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using &#39;!&#39;.\nFor example: &#39;!22,telnet,http,123,1001-1005&#39;  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
