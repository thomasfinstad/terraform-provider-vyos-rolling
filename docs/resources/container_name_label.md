---
page_title: "vyos_container_name_label Resource - vyos"

subcategory: "Container"

description: |- 
  Container applications⯯Container name⯯Add label variables
---

# vyos_container_name_label (Resource)
<center>

Container applications  
⯯  
Container name  
⯯  
**Add label variables**


</center>

## Schema

### Required

- `label_id` (String) Add label variables
- `name_id` (String) Container name

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `value` (String) Set label option value

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|Set label option value  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
