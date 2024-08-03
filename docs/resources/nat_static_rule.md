---
page_title: "vyos_nat_static_rule Resource - vyos"

subcategory: "Nat"

description: |- 
  Network Address Translation (NAT) parameters⯯Static NAT (one-to-one)⯯Rule number for NAT
---

# vyos_nat_static_rule (Resource)
<center>

Network Address Translation (NAT) parameters  
⯯  
Static NAT (one-to-one)  
⯯  
**Rule number for NAT**


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
- `destination` (Attributes) NAT destination parameters (see [below for nested schema](#nestedatt--destination))
- `inbound_interface` (String) Inbound interface of NAT traffic
- `log` (Boolean) Log packets hitting this rule
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `translation` (Attributes) Translation address or prefix (see [below for nested schema](#nestedatt--translation))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `rule` (String) Rule number for NAT


&lt;a id=&#34;nestedatt--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination`

Optional:

- `address` (String) IP address, prefix

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|ipv4     &emsp;|IPv4 address to match  |
    &emsp;|ipv4net  &emsp;|IPv4 prefix to match   |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--translation&#34;&gt;&lt;/a&gt;
### Nested Schema for `translation`

Optional:

- `address` (String) IP address, prefix

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|ipv4     &emsp;|IPv4 address to match  |
    &emsp;|ipv4net  &emsp;|IPv4 prefix to match   |  
