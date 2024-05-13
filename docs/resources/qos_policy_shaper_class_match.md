---
page_title: "vyos_qos_policy_shaper_class_match Resource - vyos"

subcategory: "Qos"

description: |- 
  Quality of Service (QoS)⯯Service Policy definitions⯯Traffic shaping based policy (Hierarchy Token Bucket)⯯Class ID⯯Class matching rule name
---

# vyos_qos_policy_shaper_class_match (Resource)
<center>

Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
Traffic shaping based policy (Hierarchy Token Bucket)  
⯯  
Class ID  
⯯  
**Class matching rule name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `class_id` (Number) Class ID

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|2-4095  &emsp;|Class Identifier  |
- `match_id` (String) Class matching rule name
- `shaper_id` (String) Traffic shaping based policy (Hierarchy Token Bucket)

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `ether` (Attributes) Ethernet header match (see [below for nested schema](#nestedatt--ether))
- `interface` (String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `ip` (Attributes) Match IP protocol header (see [below for nested schema](#nestedatt--ip))
- `ipv6` (Attributes) Match IPv6 protocol header (see [below for nested schema](#nestedatt--ipv6))
- `mark` (Number) Match on mark applied by firewall

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|u32     &emsp;|FW mark to match  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vif` (Number) Virtual Local Area Network (VLAN) ID for this match

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|0-4095  &emsp;|Virtual Local Area Network (VLAN) tag   |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--ether&#34;&gt;&lt;/a&gt;
### Nested Schema for `ether`

Optional:

- `destination` (String) Ethernet destination address for this match

    &emsp;|Format   &emsp;|Description           |
    |-----------|------------------------|
    &emsp;|macaddr  &emsp;|MAC address to match  |
- `protocol` (String) Ethernet protocol for this match

    &emsp;|Format   &emsp;|Description                      |
    |-----------|-----------------------------------|
    &emsp;|0-65535  &emsp;|Ethernet protocol number         |
    &emsp;|txt      &emsp;|Ethernet protocol name           |
    &emsp;|all      &emsp;|Any protocol                     |
    &emsp;|ip       &emsp;|Internet IP (IPv4)               |
    &emsp;|ipv6     &emsp;|Internet IP (IPv6)               |
    &emsp;|arp      &emsp;|Address Resolution Protocol      |
    &emsp;|atalk    &emsp;|Appletalk                        |
    &emsp;|ipx      &emsp;|Novell Internet Packet Exchange  |
    &emsp;|802.1Q   &emsp;|802.1Q VLAN tag                  |
- `source` (String) Ethernet source address for this match

    &emsp;|Format   &emsp;|Description           |
    |-----------|------------------------|
    &emsp;|macaddr  &emsp;|MAC address to match  |


&lt;a id=&#34;nestedatt--ip&#34;&gt;&lt;/a&gt;
### Nested Schema for `ip`

Optional:

- `destination` (Attributes) Match on destination port or address (see [below for nested schema](#nestedatt--ip--destination))
- `dscp` (String) Match on Differentiated Services Codepoint (DSCP)

    &emsp;|Format          &emsp;|Description                                      |
    |------------------|---------------------------------------------------|
    &emsp;|0-63            &emsp;|Differentiated Services Codepoint (DSCP) value   |
    &emsp;|default         &emsp;|match DSCP (000000)                              |
    &emsp;|reliability     &emsp;|match DSCP (000001)                              |
    &emsp;|throughput      &emsp;|match DSCP (000010)                              |
    &emsp;|lowdelay        &emsp;|match DSCP (000100)                              |
    &emsp;|priority        &emsp;|match DSCP (001000)                              |
    &emsp;|immediate       &emsp;|match DSCP (010000)                              |
    &emsp;|flash           &emsp;|match DSCP (011000)                              |
    &emsp;|flash-override  &emsp;|match DSCP (100000)                              |
    &emsp;|critical        &emsp;|match DSCP (101000)                              |
    &emsp;|internet        &emsp;|match DSCP (110000)                              |
    &emsp;|network         &emsp;|match DSCP (111000)                              |
    &emsp;|AF11            &emsp;|High-throughput data                             |
    &emsp;|AF12            &emsp;|High-throughput data                             |
    &emsp;|AF13            &emsp;|High-throughput data                             |
    &emsp;|AF21            &emsp;|Low-latency data                                 |
    &emsp;|AF22            &emsp;|Low-latency data                                 |
    &emsp;|AF23            &emsp;|Low-latency data                                 |
    &emsp;|AF31            &emsp;|Multimedia streaming                             |
    &emsp;|AF32            &emsp;|Multimedia streaming                             |
    &emsp;|AF33            &emsp;|Multimedia streaming                             |
    &emsp;|AF41            &emsp;|Multimedia conferencing                          |
    &emsp;|AF42            &emsp;|Multimedia conferencing                          |
    &emsp;|AF43            &emsp;|Multimedia conferencing                          |
    &emsp;|CS1             &emsp;|Low-priority data                                |
    &emsp;|CS2             &emsp;|OAM                                              |
    &emsp;|CS3             &emsp;|Broadcast video                                  |
    &emsp;|CS4             &emsp;|Real-time interactive                            |
    &emsp;|CS5             &emsp;|Signaling                                        |
    &emsp;|CS6             &emsp;|Network control                                  |
    &emsp;|CS7             &emsp;|N/A                                              |
    &emsp;|EF              &emsp;|Expedited Forwarding                             |
- `max_length` (Number) Maximum packet length

    &emsp;|Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    &emsp;|1-65535  &emsp;|Maximum packet/payload length  |
- `protocol` (String) Protocol

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|txt     &emsp;|Protocol name  |
- `source` (Attributes) Match on source port or address (see [below for nested schema](#nestedatt--ip--source))
- `tcp` (Attributes) TCP Flags matching (see [below for nested schema](#nestedatt--ip--tcp))

&lt;a id=&#34;nestedatt--ip--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `ip.destination`

Optional:

- `address` (String) IPv4 destination address for this match

    &emsp;|Format   &emsp;|Description   |
    |-----------|----------------|
    &emsp;|ipv4     &emsp;|IPv4 address  |
    &emsp;|ipv4net  &emsp;|IPv4 prefix   |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |


&lt;a id=&#34;nestedatt--ip--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `ip.source`

Optional:

- `address` (String) IPv4 destination address for this match

    &emsp;|Format   &emsp;|Description   |
    |-----------|----------------|
    &emsp;|ipv4     &emsp;|IPv4 address  |
    &emsp;|ipv4net  &emsp;|IPv4 prefix   |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |


&lt;a id=&#34;nestedatt--ip--tcp&#34;&gt;&lt;/a&gt;
### Nested Schema for `ip.tcp`

Optional:

- `ack` (Boolean) Match TCP ACK
- `syn` (Boolean) Match TCP SYN



&lt;a id=&#34;nestedatt--ipv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6`

Optional:

- `destination` (Attributes) Match on destination port or address (see [below for nested schema](#nestedatt--ipv6--destination))
- `dscp` (String) Match on Differentiated Services Codepoint (DSCP)

    &emsp;|Format          &emsp;|Description                                      |
    |------------------|---------------------------------------------------|
    &emsp;|0-63            &emsp;|Differentiated Services Codepoint (DSCP) value   |
    &emsp;|default         &emsp;|match DSCP (000000)                              |
    &emsp;|reliability     &emsp;|match DSCP (000001)                              |
    &emsp;|throughput      &emsp;|match DSCP (000010)                              |
    &emsp;|lowdelay        &emsp;|match DSCP (000100)                              |
    &emsp;|priority        &emsp;|match DSCP (001000)                              |
    &emsp;|immediate       &emsp;|match DSCP (010000)                              |
    &emsp;|flash           &emsp;|match DSCP (011000)                              |
    &emsp;|flash-override  &emsp;|match DSCP (100000)                              |
    &emsp;|critical        &emsp;|match DSCP (101000)                              |
    &emsp;|internet        &emsp;|match DSCP (110000)                              |
    &emsp;|network         &emsp;|match DSCP (111000)                              |
    &emsp;|AF11            &emsp;|High-throughput data                             |
    &emsp;|AF12            &emsp;|High-throughput data                             |
    &emsp;|AF13            &emsp;|High-throughput data                             |
    &emsp;|AF21            &emsp;|Low-latency data                                 |
    &emsp;|AF22            &emsp;|Low-latency data                                 |
    &emsp;|AF23            &emsp;|Low-latency data                                 |
    &emsp;|AF31            &emsp;|Multimedia streaming                             |
    &emsp;|AF32            &emsp;|Multimedia streaming                             |
    &emsp;|AF33            &emsp;|Multimedia streaming                             |
    &emsp;|AF41            &emsp;|Multimedia conferencing                          |
    &emsp;|AF42            &emsp;|Multimedia conferencing                          |
    &emsp;|AF43            &emsp;|Multimedia conferencing                          |
    &emsp;|CS1             &emsp;|Low-priority data                                |
    &emsp;|CS2             &emsp;|OAM                                              |
    &emsp;|CS3             &emsp;|Broadcast video                                  |
    &emsp;|CS4             &emsp;|Real-time interactive                            |
    &emsp;|CS5             &emsp;|Signaling                                        |
    &emsp;|CS6             &emsp;|Network control                                  |
    &emsp;|CS7             &emsp;|N/A                                              |
    &emsp;|EF              &emsp;|Expedited Forwarding                             |
- `max_length` (Number) Maximum packet length

    &emsp;|Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    &emsp;|1-65535  &emsp;|Maximum packet/payload length  |
- `protocol` (String) Protocol

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|txt     &emsp;|Protocol name  |
- `source` (Attributes) Match on source port or address (see [below for nested schema](#nestedatt--ipv6--source))
- `tcp` (Attributes) TCP Flags matching (see [below for nested schema](#nestedatt--ipv6--tcp))

&lt;a id=&#34;nestedatt--ipv6--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6.destination`

Optional:

- `address` (String) IPv6 destination address for this match

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv6net  &emsp;|IPv6 address and prefix length  |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |


&lt;a id=&#34;nestedatt--ipv6--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6.source`

Optional:

- `address` (String) IPv6 destination address for this match

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv6net  &emsp;|IPv6 address and prefix length  |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |


&lt;a id=&#34;nestedatt--ipv6--tcp&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6.tcp`

Optional:

- `ack` (Boolean) Match TCP ACK
- `syn` (Boolean) Match TCP SYN



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
