---
page_title: "vyos_policy_prefix_list_rule Resource - vyos"
subcategory: "policy"
description: |- 
  Routing policy⯯IP prefix-list filter⯯Rule for this prefix-list
---

# vyos_policy_prefix_list_rule (Resource)
<center>

Routing policy  
⯯  
IP prefix-list filter  
⯯  
**Rule for this prefix-list**


</center>

## Schema

### Required

- `prefix_list_id` (String) IP prefix-list filter

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv4 prefix-list  |
- `rule_id` (Number) Rule for this prefix-list

    &emsp;|Format   &emsp;|Description              |
    |-----------|---------------------------|
    &emsp;|1-65535  &emsp;|Prefix-list rule number  |

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
- `ge` (Number) Prefix length to match a netmask greater than or equal to it

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|0-32    &emsp;|Netmask greater than length  |
- `le` (Number) Prefix length to match a netmask less than or equal to it

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|0-32    &emsp;|Netmask less than length  |
- `prefix` (String) Prefix to match

    &emsp;|Format   &emsp;|Description              |
    |-----------|---------------------------|
    &emsp;|ipv4net  &emsp;|Prefix to match against  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
