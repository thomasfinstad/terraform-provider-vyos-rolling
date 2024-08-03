---
page_title: "vyos_system_conntrack_ignore_ipv4_rule Resource - vyos"

subcategory: "System"

description: |- 
  system⯯Connection Tracking Engine Options⯯Customized rules to ignore selective connection tracking⯯IPv4 rules⯯Rule number
---

# vyos_system_conntrack_ignore_ipv4_rule (Resource)
<center>

*system*  
⯯  
Connection Tracking Engine Options  
⯯  
Customized rules to ignore selective connection tracking  
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

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `destination` (Attributes) Destination parameters (see [below for nested schema](#nestedatt--destination))
- `inbound_interface` (String) Interface to ignore connections tracking on
- `protocol` (String) Protocol

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|txt     &emsp;|Protocol name  |
- `source` (Attributes) Source parameters (see [below for nested schema](#nestedatt--source))
- `tcp` (Attributes) TCP options to match (see [below for nested schema](#nestedatt--tcp))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `rule` (Number) Rule number

    &emsp;|Format    &emsp;|Description                      |
    |------------|-----------------------------------|
    &emsp;|1-999999  &emsp;|Number of conntrack ignore rule  |


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
- `group` (Attributes) Group (see [below for nested schema](#nestedatt--destination--group))
- `port` (String) Port number

    &emsp;|Format     &emsp;|Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt        &emsp;|Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    &emsp;|1-65535    &emsp;|Numeric IP port                                                                                                                                                          |
    &emsp;|start-end  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    &emsp;|           &emsp;|\n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using &#39;!&#39;.\nFor example: &#39;!22,telnet,http,123,1001-1005&#39;  |

&lt;a id=&#34;nestedatt--destination--group&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination.group`

Optional:

- `address_group` (String) Group of addresses
- `domain_group` (String) Group of domains
- `network_group` (String) Group of networks
- `port_group` (String) Group of ports



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
- `group` (Attributes) Group (see [below for nested schema](#nestedatt--source--group))
- `port` (String) Port number

    &emsp;|Format     &emsp;|Description                                                                                                                                                              |
    |-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|txt        &emsp;|Named port (any name in /etc/services, e.g., http)                                                                                                                       |
    &emsp;|1-65535    &emsp;|Numeric IP port                                                                                                                                                          |
    &emsp;|start-end  &emsp;|Numbered port range (e.g. 1001-1005)                                                                                                                                     |
    &emsp;|           &emsp;|\n\nMultiple destination ports can be specified as a comma-separated list.\nThe whole list can also be negated using &#39;!&#39;.\nFor example: &#39;!22,telnet,http,123,1001-1005&#39;  |

&lt;a id=&#34;nestedatt--source--group&#34;&gt;&lt;/a&gt;
### Nested Schema for `source.group`

Optional:

- `address_group` (String) Group of addresses
- `domain_group` (String) Group of domains
- `network_group` (String) Group of networks
- `port_group` (String) Group of ports



&lt;a id=&#34;nestedatt--tcp&#34;&gt;&lt;/a&gt;
### Nested Schema for `tcp`

Optional:

- `flags` (Attributes) TCP flags to match (see [below for nested schema](#nestedatt--tcp--flags))

&lt;a id=&#34;nestedatt--tcp--flags&#34;&gt;&lt;/a&gt;
### Nested Schema for `tcp.flags`

Optional:

- `ack` (Boolean) Acknowledge flag
- `cwr` (Boolean) Congestion Window Reduced flag
- `ecn` (Boolean) Explicit Congestion Notification flag
- `fin` (Boolean) Finish flag
- `not` (Attributes) Match flags not set (see [below for nested schema](#nestedatt--tcp--flags--not))
- `psh` (Boolean) Push flag
- `rst` (Boolean) Reset flag
- `syn` (Boolean) Synchronise flag
- `urg` (Boolean) Urgent flag

&lt;a id=&#34;nestedatt--tcp--flags--not&#34;&gt;&lt;/a&gt;
### Nested Schema for `tcp.flags.not`

Optional:

- `ack` (Boolean) Acknowledge flag
- `cwr` (Boolean) Congestion Window Reduced flag
- `ecn` (Boolean) Explicit Congestion Notification flag
- `fin` (Boolean) Finish flag
- `psh` (Boolean) Push flag
- `rst` (Boolean) Reset flag
- `syn` (Boolean) Synchronise flag
- `urg` (Boolean) Urgent flag




&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
