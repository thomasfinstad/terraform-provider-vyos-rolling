---
page_title: "vyos_policy_local_route6_rule Resource - vyos"

subcategory: "Policy"

description: |- 
  policy⯯IPv6 policy route of local traffic⯯IPv6 policy local-route rule set number
---

# vyos_policy_local_route6_rule (Resource)
<center>

*policy*  
⯯  
IPv6 policy route of local traffic  
⯯  
**IPv6 policy local-route rule set number**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `destination` (Attributes) Destination parameters (see [below for nested schema](#nestedatt--destination))
- `fwmark` (Number) Match fwmark value

    &emsp;|Format        &emsp;|Description               |
    |----------------|----------------------------|
    &emsp;|1-2147483647  &emsp;|Address to match against  |
- `inbound_interface` (String) Inbound Interface
- `protocol` (String) Protocol to match (protocol name or number)

    &emsp;|Format      &emsp;|Description         |
    |--------------|----------------------|
    &emsp;|0-255       &emsp;|IP protocol number  |
    &emsp;|&lt;protocol&gt;  &emsp;|IP protocol name    |
- `set` (Attributes) Packet modifications (see [below for nested schema](#nestedatt--set))
- `source` (Attributes) Source parameters (see [below for nested schema](#nestedatt--source))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `rule` (Number) IPv6 policy local-route rule set number

    &emsp;|Format   &emsp;|Description                        |
    |-----------|-------------------------------------|
    &emsp;|1-32765  &emsp;|Local-route rule number (1-32765)  |


&lt;a id=&#34;nestedatt--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination`

Optional:

- `address` (List of String) IPv6 address or prefix

    &emsp;|Format   &emsp;|Description               |
    |-----------|----------------------------|
    &emsp;|ipv6     &emsp;|Address to match against  |
    &emsp;|ipv6net  &emsp;|Prefix to match against   |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |


&lt;a id=&#34;nestedatt--set&#34;&gt;&lt;/a&gt;
### Nested Schema for `set`

Optional:

- `table` (Number) Routing table to forward packet with

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|1-200   &emsp;|Table number  |


&lt;a id=&#34;nestedatt--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `source`

Optional:

- `address` (List of String) IPv6 address or prefix

    &emsp;|Format   &emsp;|Description               |
    |-----------|----------------------------|
    &emsp;|ipv6     &emsp;|Address to match against  |
    &emsp;|ipv6net  &emsp;|Prefix to match against   |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
