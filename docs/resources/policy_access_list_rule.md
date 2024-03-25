---
page_title: "vyos_policy_access_list_rule Resource - vyos"
subcategory: "policy"
description: |- 
  Routing policy⯯IP access-list filter⯯Rule for this access-list
---

# vyos_policy_access_list_rule (Resource)
<center>

Routing policy  
⯯  
IP access-list filter  
⯯  
**Rule for this access-list**


</center>

## Schema

### Required

- `access_list_id` (Number) IP access-list filter

    &emsp;|Format     &emsp;|Description                               |
    |-------------|--------------------------------------------|
    &emsp;|1-99       &emsp;|IP standard access list                   |
    &emsp;|100-199    &emsp;|IP extended access list                   |
    &emsp;|1300-1999  &emsp;|IP standard access list (expanded range)  |
    &emsp;|2000-2699  &emsp;|IP extended access list (expanded range)  |
- `rule_id` (Number) Rule for this access-list

    &emsp;|Format   &emsp;|Description              |
    |-----------|---------------------------|
    &emsp;|1-65535  &emsp;|Access-list rule number  |

### Optional

- `action` (String) Action to take on entries matching this rule

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|permit  &emsp;|Permit matching entries  |
    &emsp;|deny    &emsp;|Deny matching entries    |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `destination` (Attributes) Destination network or address (see [below for nested schema](#nestedatt--destination))
- `source` (Attributes) Source network or address to match (see [below for nested schema](#nestedatt--source))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--destination&#34;&gt;&lt;/a&gt;
### Nested Schema for `destination`

Optional:

- `any` (Boolean) Any IP address to match
- `host` (String) Single host IP address to match

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|ipv4    &emsp;|Host address to match  |
- `inverse_mask` (String) Network/netmask to match (requires network be defined)

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|ipv4    &emsp;|Inverse-mask to match  |
- `network` (String) Network/netmask to match (requires inverse-mask be defined)

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|ipv4net  &emsp;|Inverse-mask to match  |


&lt;a id=&#34;nestedatt--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `source`

Optional:

- `any` (Boolean) Any IP address to match
- `host` (String) Single host IP address to match

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|ipv4    &emsp;|Host address to match  |
- `inverse_mask` (String) Network/netmask to match (requires network be defined)

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|ipv4    &emsp;|Inverse-mask to match  |
- `network` (String) Network/netmask to match (requires inverse-mask be defined)

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|ipv4net  &emsp;|Inverse-mask to match  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
