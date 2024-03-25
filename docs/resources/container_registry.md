---
page_title: "vyos_container_registry Resource - vyos"
subcategory: "container"
description: |- 
  Container applications⯯Registry Name
---

# vyos_container_registry (Resource)
<center>

Container applications  
⯯  
**Registry Name**


</center>

## Schema

### Required

- `registry_id` (String) Registry Name

### Optional

- `authentication` (Attributes) Authentication settings (see [below for nested schema](#nestedatt--authentication))
- `disable` (Boolean) Disable instance
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--authentication&#34;&gt;&lt;/a&gt;
### Nested Schema for `authentication`

Optional:

- `password` (String) Password used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Password     |
- `username` (String) Username used for authentication

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Username     |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
