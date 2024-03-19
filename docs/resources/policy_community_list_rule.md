---
page_title: "vyos_policy_community_list_rule Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy
  ⯯
  Add a BGP community list entry
  ⯯
  Rule for this BGP community list
---

# vyos_policy_community_list_rule (Resource)
<center>

Routing policy
⯯
Add a BGP community list entry
⯯
**Rule for this BGP community list**


</center>

## Schema

### Required

- `community_list_id` (String) Add a BGP community list entry

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|txt     &emsp;|BGP community-list name  |
- `rule_id` (Number) Rule for this BGP community list

    &emsp;|Format   &emsp;|Description                 |
    |-----------|------------------------------|
    &emsp;|1-65535  &emsp;|Community-list rule number  |

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
- `regex` (String) Regular expression to match against a community-list

    &emsp;|Format        &emsp;|Description                                                  |
    |----------------|---------------------------------------------------------------|
    &emsp;|&lt;aa:nn&gt;       &emsp;|Community number in AA:NN format                             |
    &emsp;|local-AS      &emsp;|Well-known communities value NO_EXPORT_SUBCONFED 0xFFFFFF03  |
    &emsp;|no-advertise  &emsp;|Well-known communities value NO_ADVERTISE 0xFFFFFF02         |
    &emsp;|no-export     &emsp;|Well-known communities value NO_EXPORT 0xFFFFFF01            |
    &emsp;|internet      &emsp;|Well-known communities value 0                               |
    &emsp;|additive      &emsp;|New value is appended to the existing value                  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
