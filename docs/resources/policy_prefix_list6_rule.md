---
page_title: "vyos_policy_prefix_list6_rule Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy
  ⯯
  IPv6 prefix-list filter
  ⯯
  Rule for this prefix-list6
---

# vyos_policy_prefix_list6_rule (Resource)
<center>

Routing policy
⯯
IPv6 prefix-list filter
⯯
**Rule for this prefix-list6**


</center>

## Schema

### Required

- `prefix_list6_id` (String) IPv6 prefix-list filter

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|txt     &emsp;|Name of IPv6 prefix-list  |
- `rule_id` (Number) Rule for this prefix-list6

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
    &emsp;|0-128   &emsp;|Netmask greater than length  |
- `le` (Number) Prefix length to match a netmask less than or equal to it

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|0-128   &emsp;|Netmask less than length  |
- `prefix` (String) Prefix to match

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|ipv6net  &emsp;|IPv6 prefix  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
