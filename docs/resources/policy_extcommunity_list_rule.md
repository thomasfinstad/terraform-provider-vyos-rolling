---
page_title: "vyos_policy_extcommunity_list_rule Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy
  ⯯
  Add a BGP extended community list entry
  ⯯
  Rule for this BGP extended community list
---

# vyos_policy_extcommunity_list_rule (Resource)
<center>

Routing policy
⯯
Add a BGP extended community list entry
⯯
**Rule for this BGP extended community list**


</center>

## Schema

### Required

- `extcommunity_list_id` (String) Add a BGP extended community list entry

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|txt     &emsp;|BGP extended community-list name  |
- `rule_id` (Number) Rule for this BGP extended community list

    &emsp;|Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    &emsp;|1-65535  &emsp;|Extended community-list rule number  |

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
- `regex` (String) Regular expression to match against an extended community list

    &emsp;|Format          &emsp;|Description                                 |
    |------------------|----------------------------------------------|
    &emsp;|&lt;aa:nn:nn&gt;      &emsp;|Extended community list regular expression  |
    &emsp;|&lt;rt aa:nn:nn&gt;   &emsp;|Route Target regular expression             |
    &emsp;|&lt;soo aa:nn:nn&gt;  &emsp;|Site of Origin regular expression           |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
