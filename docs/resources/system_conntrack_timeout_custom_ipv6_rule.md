---
page_title: "vyos_system_conntrack_timeout_custom_ipv6_rule Resource - vyos"

subcategory: "System"

description: |- 
  system⯯Connection Tracking Engine Options⯯Connection timeout options⯯Define custom timeouts per connection⯯IPv6 rules⯯Rule number
---

# vyos_system_conntrack_timeout_custom_ipv6_rule (Resource)
<center>

*system*  
⯯  
Connection Tracking Engine Options  
⯯  
Connection timeout options  
⯯  
Define custom timeouts per connection  
⯯  
IPv6 rules  
⯯  
**Rule number**


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
- `destination` (Attributes) Destination parameters (see [below for nested schema](#nestedatt--destination))
- `inbound_interface` (String) Interface to apply custom connection timers on
- `protocol` (Attributes) Customize protocol specific timers, one protocol configuration per rule (see [below for nested schema](#nestedatt--protocol))
- `source` (Attributes) Source parameters (see [below for nested schema](#nestedatt--source))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `rule` (Number) Rule number

    &emsp;|Format    &emsp;|Description               |
    |------------|----------------------------|
    &emsp;|1-999999  &emsp;|Number of conntrack rule  |


&lt;a id=&#34;nestedatt--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination`

Optional:

- `address` (String) IP address, subnet, or range

    &emsp;|Format      &emsp;|Description                                    |
    |--------------|-------------------------------------------------|
    &emsp;|ipv6        &emsp;|IP address to match                            |
    &emsp;|ipv6net     &emsp;|Subnet to match                                |
    &emsp;|ipv6range   &emsp;|IP range to match                              |
    &emsp;|!ipv6       &emsp;|Match everything except the specified address  |
    &emsp;|!ipv6net    &emsp;|Match everything except the specified prefix   |
    &emsp;|!ipv6range  &emsp;|Match everything except the specified range    |
- `port` (String) Port number

    &emsp;|Format     &emsp;|Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt        &emsp;|Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    &emsp;|1-65535    &emsp;|Numeric IP port                                                                                                                                                          |
    &emsp;|start-end  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    &emsp;|           &emsp;|\n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using &#39;!&#39;.\nFor example: &#39;!22,telnet,http,123,1001-1005&#39;  |


&lt;a id=&#34;nestedatt--protocol&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocol`

Optional:

- `tcp` (Attributes) TCP connection timeout options (see [below for nested schema](#nestedatt--protocol--tcp))
- `udp` (Attributes) UDP timeout options (see [below for nested schema](#nestedatt--protocol--udp))

&lt;a id=&#34;nestedatt--protocol--tcp&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocol.tcp`

Optional:

- `close` (Number) TCP CLOSE timeout in seconds

    &emsp;|Format      &emsp;|Description                   |
    |--------------|--------------------------------|
    &emsp;|1-21474836  &emsp;|TCP CLOSE timeout in seconds  |
- `close_wait` (Number) TCP CLOSE-WAIT timeout in seconds

    &emsp;|Format      &emsp;|Description                        |
    |--------------|-------------------------------------|
    &emsp;|1-21474836  &emsp;|TCP CLOSE-WAIT timeout in seconds  |
- `established` (Number) TCP ESTABLISHED timeout in seconds

    &emsp;|Format      &emsp;|Description                         |
    |--------------|--------------------------------------|
    &emsp;|1-21474836  &emsp;|TCP ESTABLISHED timeout in seconds  |
- `fin_wait` (Number) TCP FIN-WAIT timeout in seconds

    &emsp;|Format      &emsp;|Description                      |
    |--------------|-----------------------------------|
    &emsp;|1-21474836  &emsp;|TCP FIN-WAIT timeout in seconds  |
- `last_ack` (Number) TCP LAST-ACK timeout in seconds

    &emsp;|Format      &emsp;|Description                      |
    |--------------|-----------------------------------|
    &emsp;|1-21474836  &emsp;|TCP LAST-ACK timeout in seconds  |
- `syn_recv` (Number) TCP SYN-RECEIVED timeout in seconds

    &emsp;|Format      &emsp;|Description                          |
    |--------------|---------------------------------------|
    &emsp;|1-21474836  &emsp;|TCP SYN-RECEIVED timeout in seconds  |
- `syn_sent` (Number) TCP SYN-SENT timeout in seconds

    &emsp;|Format      &emsp;|Description                      |
    |--------------|-----------------------------------|
    &emsp;|1-21474836  &emsp;|TCP SYN-SENT timeout in seconds  |
- `time_wait` (Number) TCP TIME-WAIT timeout in seconds

    &emsp;|Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    &emsp;|1-21474836  &emsp;|TCP TIME-WAIT timeout in seconds  |


&lt;a id=&#34;nestedatt--protocol--udp&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocol.udp`

Optional:

- `replied` (Number) Timeout for UDP connection seen in both directions

    &emsp;|Format      &emsp;|Description                                         |
    |--------------|------------------------------------------------------|
    &emsp;|1-21474836  &emsp;|Timeout for UDP connection seen in both directions  |
- `unreplied` (Number) Timeout for unreplied UDP

    &emsp;|Format      &emsp;|Description                |
    |--------------|-----------------------------|
    &emsp;|1-21474836  &emsp;|Timeout for unreplied UDP  |



&lt;a id=&#34;nestedatt--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `source`

Optional:

- `address` (String) IP address, subnet, or range

    &emsp;|Format      &emsp;|Description                                    |
    |--------------|-------------------------------------------------|
    &emsp;|ipv6        &emsp;|IP address to match                            |
    &emsp;|ipv6net     &emsp;|Subnet to match                                |
    &emsp;|ipv6range   &emsp;|IP range to match                              |
    &emsp;|!ipv6       &emsp;|Match everything except the specified address  |
    &emsp;|!ipv6net    &emsp;|Match everything except the specified prefix   |
    &emsp;|!ipv6range  &emsp;|Match everything except the specified range    |
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
