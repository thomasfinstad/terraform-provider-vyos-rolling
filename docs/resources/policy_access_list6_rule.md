---
page_title: "vyos_policy_access_list6_rule Resource - vyos"

subcategory: "Policy"

description: |- 
  Routing policy⯯IPv6 access-list filter⯯Rule for this access-list6
---

# vyos_policy_access_list6_rule (Resource)
<center>

Routing policy  
⯯  
IPv6 access-list filter  
⯯  
**Rule for this access-list6**


</center>

## Schema

### Required

- `access_list6_id` (String) IPv6 access-list filter

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 access-list  |
- `rule_id` (Number) Rule for this access-list6

    &emsp;|Format   &emsp;|Description               |
    |-----------|----------------------------|
    &emsp;|1-65535  &emsp;|Access-list6 rule number  |

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
- `source` (Attributes) Source IPv6 network to match (see [below for nested schema](#nestedatt--source))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `source`

Optional:

- `any` (Boolean) Any IP address to match
- `exact_match` (Boolean) Exact match of the network prefixes
- `network` (String) Network/netmask to match

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv6net  &emsp;|IPv6 address and prefix length  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
