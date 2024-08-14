---
page_title: "vyos_system_conntrack_timeout_custom_ipv4_rule Resource - vyos"

subcategory: "System"

description: |- 
  system⯯Connection Tracking Engine Options⯯Connection timeout options⯯Define custom timeouts per connection⯯IPv4 rules⯯Rule number
---

# vyos_system_conntrack_timeout_custom_ipv4_rule (Resource)
<center>

*system*  
⯯  
Connection Tracking Engine Options  
⯯  
Connection timeout options  
⯯  
Define custom timeouts per connection  
⯯  
IPv4 rules  
⯯  
**Rule number**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `destination` (Attributes) Destination parameters (see [below for nested schema](#nestedatt--destination))
- `inbound_interface` (String) Interface to apply custom connection timers on
- `protocol` (Attributes) Customize protocol specific timers, one protocol configuration per rule (see [below for nested schema](#nestedatt--protocol))
- `source` (Attributes) Source parameters (see [below for nested schema](#nestedatt--source))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `rule` (Number) Rule number

    |Format    &emsp;|Description               |
    |------------|----------------------------|
    |1-999999  &emsp;|Number of conntrack rule  |


<a id="nestedatt--destination"></a>
### Nested Schema for `destination`

Optional:

- `address` (String) IP address, subnet, or range

    |Format      &emsp;|Description                                    |
    |--------------|-------------------------------------------------|
    |ipv4        &emsp;|IPv4 address to match                          |
    |ipv4net     &emsp;|IPv4 prefix to match                           |
    |ipv4range   &emsp;|IPv4 address range to match                    |
    |!ipv4       &emsp;|Match everything except the specified address  |
    |!ipv4net    &emsp;|Match everything except the specified prefix   |
    |!ipv4range  &emsp;|Match everything except the specified range    |
- `port` (String) Port number

    |Format     &emsp;|Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    |txt        &emsp;|Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    |1-65535    &emsp;|Numeric IP port                                                                                                                                                          |
    |start-end  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    |           &emsp;|\n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using &#39;!&#39;.\nFor example: &#39;!22,telnet,http,123,1001-1005&#39;  |


<a id="nestedatt--protocol"></a>
### Nested Schema for `protocol`

Optional:

- `tcp` (Attributes) TCP connection timeout options (see [below for nested schema](#nestedatt--protocol--tcp))
- `udp` (Attributes) UDP timeout options (see [below for nested schema](#nestedatt--protocol--udp))

<a id="nestedatt--protocol--tcp"></a>
### Nested Schema for `protocol.tcp`

Optional:

- `close` (Number) TCP CLOSE timeout in seconds

    |Format      &emsp;|Description                   |
    |--------------|--------------------------------|
    |1-21474836  &emsp;|TCP CLOSE timeout in seconds  |
- `close_wait` (Number) TCP CLOSE-WAIT timeout in seconds

    |Format      &emsp;|Description                        |
    |--------------|-------------------------------------|
    |1-21474836  &emsp;|TCP CLOSE-WAIT timeout in seconds  |
- `established` (Number) TCP ESTABLISHED timeout in seconds

    |Format      &emsp;|Description                         |
    |--------------|--------------------------------------|
    |1-21474836  &emsp;|TCP ESTABLISHED timeout in seconds  |
- `fin_wait` (Number) TCP FIN-WAIT timeout in seconds

    |Format      &emsp;|Description                      |
    |--------------|-----------------------------------|
    |1-21474836  &emsp;|TCP FIN-WAIT timeout in seconds  |
- `last_ack` (Number) TCP LAST-ACK timeout in seconds

    |Format      &emsp;|Description                      |
    |--------------|-----------------------------------|
    |1-21474836  &emsp;|TCP LAST-ACK timeout in seconds  |
- `syn_recv` (Number) TCP SYN-RECEIVED timeout in seconds

    |Format      &emsp;|Description                          |
    |--------------|---------------------------------------|
    |1-21474836  &emsp;|TCP SYN-RECEIVED timeout in seconds  |
- `syn_sent` (Number) TCP SYN-SENT timeout in seconds

    |Format      &emsp;|Description                      |
    |--------------|-----------------------------------|
    |1-21474836  &emsp;|TCP SYN-SENT timeout in seconds  |
- `time_wait` (Number) TCP TIME-WAIT timeout in seconds

    |Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    |1-21474836  &emsp;|TCP TIME-WAIT timeout in seconds  |


<a id="nestedatt--protocol--udp"></a>
### Nested Schema for `protocol.udp`

Optional:

- `replied` (Number) Timeout for UDP connection seen in both directions

    |Format      &emsp;|Description                                         |
    |--------------|------------------------------------------------------|
    |1-21474836  &emsp;|Timeout for UDP connection seen in both directions  |
- `unreplied` (Number) Timeout for unreplied UDP

    |Format      &emsp;|Description                |
    |--------------|-----------------------------|
    |1-21474836  &emsp;|Timeout for unreplied UDP  |



<a id="nestedatt--source"></a>
### Nested Schema for `source`

Optional:

- `address` (String) IP address, subnet, or range

    |Format      &emsp;|Description                                    |
    |--------------|-------------------------------------------------|
    |ipv4        &emsp;|IPv4 address to match                          |
    |ipv4net     &emsp;|IPv4 prefix to match                           |
    |ipv4range   &emsp;|IPv4 address range to match                    |
    |!ipv4       &emsp;|Match everything except the specified address  |
    |!ipv4net    &emsp;|Match everything except the specified prefix   |
    |!ipv4range  &emsp;|Match everything except the specified range    |
- `port` (String) Port number

    |Format     &emsp;|Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    |txt        &emsp;|Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    |1-65535    &emsp;|Numeric IP port                                                                                                                                                          |
    |start-end  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    |           &emsp;|\n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using &#39;!&#39;.\nFor example: &#39;!22,telnet,http,123,1001-1005&#39;  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
