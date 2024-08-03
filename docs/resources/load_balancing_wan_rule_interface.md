---
page_title: "vyos_load_balancing_wan_rule_interface Resource - vyos"

subcategory: "Load"

description: |- 
  Configure load-balancing⯯Configure Wide Area Network (WAN) load-balancing⯯Rule number (1-9999)⯯Interface name [REQUIRED]
---

# vyos_load_balancing_wan_rule_interface (Resource)
<center>

Configure load-balancing  
⯯  
Configure Wide Area Network (WAN) load-balancing  
⯯  
Rule number (1-9999)  
⯯  
**Interface name [REQUIRED]**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `weight` (Number) Load-balance weight

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|1-255   &emsp;|Interface weight  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface name [REQUIRED]
- `rule` (Number) Rule number (1-9999)

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1-9999  &emsp;|Rule number  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
