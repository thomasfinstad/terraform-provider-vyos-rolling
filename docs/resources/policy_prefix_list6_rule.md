---
page_title: "vyos_policy_prefix_list6_rule Resource - vyos"

subcategory: "Policy"

description: |- 
  Routing policy⯯IPv6 prefix-list filter⯯Rule for this prefix-list6
---

# vyos_policy_prefix_list6_rule (Resource)
<center>

Routing policy  
⯯  
IPv6 prefix-list filter  
⯯  
**Rule for this prefix-list6**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `action` (String) Action to take on entries matching this rule

    |Format  &emsp;|Description              |
    |----------|---------------------------|
    |permit  &emsp;|Permit matching entries  |
    |deny    &emsp;|Deny matching entries    |
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `ge` (Number) Prefix length to match a netmask greater than or equal to it

    |Format  &emsp;|Description                  |
    |----------|-------------------------------|
    |0-128   &emsp;|Netmask greater than length  |
- `le` (Number) Prefix length to match a netmask less than or equal to it

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |0-128   &emsp;|Netmask less than length  |
- `prefix` (String) Prefix to match

    |Format   &emsp;|Description  |
    |-----------|---------------|
    |ipv6net  &emsp;|IPv6 prefix  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `prefix_list6` (String) IPv6 prefix-list filter

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |txt     &emsp;|Name of IPv6 prefix-list  |
- `rule` (Number) Rule for this prefix-list6

    |Format   &emsp;|Description              |
    |-----------|---------------------------|
    |1-65535  &emsp;|Prefix-list rule number  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
