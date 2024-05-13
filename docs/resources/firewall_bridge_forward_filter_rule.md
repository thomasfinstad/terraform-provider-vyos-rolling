---
page_title: "vyos_firewall_bridge_forward_filter_rule Resource - vyos"

subcategory: "Firewall"

description: |- 
  Firewall⯯Bridge firewall⯯Bridge forward firewall⯯Bridge firewall forward filter⯯Bridge Firewall forward filter rule number
---

# vyos_firewall_bridge_forward_filter_rule (Resource)
<center>

Firewall  
⯯  
Bridge firewall  
⯯  
Bridge forward firewall  
⯯  
Bridge firewall forward filter  
⯯  
**Bridge Firewall forward filter rule number**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `rule_id` (Number) Bridge Firewall forward filter rule number

    &emsp;|Format    &emsp;|Description                    |
    |------------|---------------------------------|
    &emsp;|1-999999  &emsp;|Number for this firewall rule  |

### Optional

- `action` (String) Rule action

    &emsp;|Format    &emsp;|Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    &emsp;|accept    &emsp;|Accept matching entries                                                        |
    &emsp;|continue  &emsp;|Continue parsing next rule                                                     |
    &emsp;|jump      &emsp;|Jump to another chain                                                          |
    &emsp;|return    &emsp;|Return from the current chain and continue at the next rule of the last chain  |
    &emsp;|drop      &emsp;|Drop matching entries                                                          |
    &emsp;|queue     &emsp;|Enqueue packet to userspace                                                    |
- `destination` (Attributes) Destination parameters (see [below for nested schema](#nestedatt--destination))
- `disable` (Boolean) Disable instance
- `inbound_interface` (Attributes) Match inbound-interface (see [below for nested schema](#nestedatt--inbound_interface))
- `jump_target` (String) Set jump target. Action jump must be defined to use this setting
- `log` (Boolean) Log packets hitting this rule
- `log_options` (Attributes) Log options (see [below for nested schema](#nestedatt--log_options))
- `outbound_interface` (Attributes) Match outbound-interface (see [below for nested schema](#nestedatt--outbound_interface))
- `queue` (Number) Queue target to use. Action queue must be defined to use this setting

    &emsp;|Format   &emsp;|Description   |
    |-----------|----------------|
    &emsp;|0-65535  &emsp;|Queue target  |
- `queue_options` (List of String) Options used for queue target. Action queue must be defined to use this setting

    &emsp;|Format  &emsp;|Description                                                      |
    |----------|-------------------------------------------------------------------|
    &emsp;|bypass  &emsp;|Let packets go through if userspace application cannot back off  |
    &emsp;|fanout  &emsp;|Distribute packets between several queues                        |
- `source` (Attributes) Source parameters (see [below for nested schema](#nestedatt--source))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vlan` (Attributes) VLAN parameters (see [below for nested schema](#nestedatt--vlan))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination`

Optional:

- `mac_address` (String) MAC address

    &emsp;|Format    &emsp;|Description                                        |
    |------------|-----------------------------------------------------|
    &emsp;|macaddr   &emsp;|MAC address to match                               |
    &emsp;|!macaddr  &emsp;|Match everything except the specified MAC address  |


&lt;a id=&#34;nestedatt--inbound_interface&#34;&gt;&lt;/a&gt;
### Nested Schema for `inbound_interface`

Optional:

- `group` (String) Match interface-group

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|txt     &emsp;|Interface-group name to match           |
    &emsp;|!txt    &emsp;|Inverted interface-group name to match  |
- `name` (String) Match interface

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|txt     &emsp;|Interface name                    |
    &emsp;|txt&amp;    &emsp;|Interface name with wildcard      |
    &emsp;|!txt    &emsp;|Inverted interface name to match  |


&lt;a id=&#34;nestedatt--log_options&#34;&gt;&lt;/a&gt;
### Nested Schema for `log_options`

Optional:

- `group` (Number) Set log group

    &emsp;|Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    &emsp;|0-65535  &emsp;|Log group to send messages to  |
- `level` (String) Set log-level

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|emerg   &emsp;|Emerg log level     |
    &emsp;|alert   &emsp;|Alert log level     |
    &emsp;|crit    &emsp;|Critical log level  |
    &emsp;|err     &emsp;|Error log level     |
    &emsp;|warn    &emsp;|Warning log level   |
    &emsp;|notice  &emsp;|Notice log level    |
    &emsp;|info    &emsp;|Info log level      |
    &emsp;|debug   &emsp;|Debug log level     |
- `queue_threshold` (Number) Number of packets to queue inside the kernel before sending them to userspace

    &emsp;|Format   &emsp;|Description                                                                    |
    |-----------|---------------------------------------------------------------------------------|
    &emsp;|0-65535  &emsp;|Number of packets to queue inside the kernel before sending them to userspace  |
- `snapshot_length` (Number) Length of packet payload to include in netlink message

    &emsp;|Format  &emsp;|Description                                             |
    |----------|----------------------------------------------------------|
    &emsp;|0-9000  &emsp;|Length of packet payload to include in netlink message  |


&lt;a id=&#34;nestedatt--outbound_interface&#34;&gt;&lt;/a&gt;
### Nested Schema for `outbound_interface`

Optional:

- `group` (String) Match interface-group

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|txt     &emsp;|Interface-group name to match           |
    &emsp;|!txt    &emsp;|Inverted interface-group name to match  |
- `name` (String) Match interface

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|txt     &emsp;|Interface name                    |
    &emsp;|txt&amp;    &emsp;|Interface name with wildcard      |
    &emsp;|!txt    &emsp;|Inverted interface name to match  |


&lt;a id=&#34;nestedatt--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `source`

Optional:

- `mac_address` (String) MAC address

    &emsp;|Format    &emsp;|Description                                        |
    |------------|-----------------------------------------------------|
    &emsp;|macaddr   &emsp;|MAC address to match                               |
    &emsp;|!macaddr  &emsp;|Match everything except the specified MAC address  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--vlan&#34;&gt;&lt;/a&gt;
### Nested Schema for `vlan`

Optional:

- `id` (String) Vlan id

    &emsp;|Format       &emsp;|Description             |
    |---------------|--------------------------|
    &emsp;|0-4096       &emsp;|Vlan id                 |
    &emsp;|&lt;start-end&gt;  &emsp;|Vlan id range to match  |
- `priority` (String) Vlan priority(pcp)

    &emsp;|Format       &emsp;|Description                   |
    |---------------|--------------------------------|
    &emsp;|0-7          &emsp;|Vlan priority                 |
    &emsp;|&lt;start-end&gt;  &emsp;|Vlan priority range to match  |  
