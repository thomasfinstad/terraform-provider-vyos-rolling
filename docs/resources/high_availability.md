---
page_title: "vyos_high_availability Resource - vyos"
subcategory: "high"
description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  **High availability settings**
---

# vyos_high_availability (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

	**High availability settings**


</center>

## Schema

### Optional

- `disable` (Boolean) Disable instance
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
