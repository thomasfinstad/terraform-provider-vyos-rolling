---
page_title: "vyos_policy_access_list Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy
  ⯯
  IP access-list filter
---

# vyos_policy_access_list (Resource)
<center>

Routing policy
⯯
**IP access-list filter**


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

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
