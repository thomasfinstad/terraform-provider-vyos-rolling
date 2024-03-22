---
page_title: "vyos_container_name_environment Resource - terraform-provider-vyos"
subcategory: "container"
description: |-
  Container applications⯯Container name⯯Add custom environment variables
---

# vyos_container_name_environment (Resource)
<center>

Container applications
⯯
Container name
⯯
**Add custom environment variables**


</center>

## Schema

### Required

- `environment_id` (String) Add custom environment variables
- `name_id` (String) Container name

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `value` (String) Set environment option value

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|txt     &emsp;|Set environment option value  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
