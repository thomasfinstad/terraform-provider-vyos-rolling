---
page_title: "vyos_container_name_sysctl_parameter Resource - vyos"

subcategory: "Container"

description: |- 
  Container applications⯯Container name⯯Configure namespaced kernel parameters of the container⯯Sysctl key name
---

# vyos_container_name_sysctl_parameter (Resource)
<center>

Container applications  
⯯  
Container name  
⯯  
Configure namespaced kernel parameters of the container  
⯯  
**Sysctl key name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `name_id` (String) Container name
- `parameter_id` (String) Sysctl key name

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|Sysctl key name  |

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `value` (String) Sysctl configuration value

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
