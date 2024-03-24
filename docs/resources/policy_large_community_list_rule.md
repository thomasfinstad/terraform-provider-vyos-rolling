---
page_title: "vyos_policy_large_community_list_rule Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy⯯Add a BGP large community list entry⯯Rule for this BGP extended community list
---

# vyos_policy_large_community_list_rule (Resource)
<center>

Routing policy  
⯯  
Add a BGP large community list entry  
⯯  
**Rule for this BGP extended community list**


</center>

## Schema

### Required

- `large_community_list_id` (String) Add a BGP large community list entry

    &emsp;|Format  &emsp;|Description                    |
    |----------|---------------------------------|
    &emsp;|txt     &emsp;|BGP large-community-list name  |
- `rule_id` (Number) Rule for this BGP extended community list

    &emsp;|Format   &emsp;|Description                       |
    |-----------|------------------------------------|
    &emsp;|1-65535  &emsp;|Large community-list rule number  |

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
- `regex` (String) Regular expression to match against a large community list

    &emsp;|Format     &emsp;|Description                                            |
    |-------------|---------------------------------------------------------|
    &emsp;|ASN:NN:NN  &emsp;|BGP large-community-list filter                        |
    &emsp;|IP:NN:NN   &emsp;|BGP large-community-list filter (IPv4 address format)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
