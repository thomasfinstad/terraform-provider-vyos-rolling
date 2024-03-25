---
page_title: "vyos_policy_as_path_list Resource - vyos"
subcategory: "policy"
description: |- 
  Routing policy⯯Add a BGP autonomous system path filter
---

# vyos_policy_as_path_list (Resource)
<center>

Routing policy  
⯯  
**Add a BGP autonomous system path filter**


</center>

## Schema

### Required

- `as_path_list_id` (String) Add a BGP autonomous system path filter

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|AS path list name  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
