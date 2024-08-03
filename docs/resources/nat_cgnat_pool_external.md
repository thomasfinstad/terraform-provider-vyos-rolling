---
page_title: "vyos_nat_cgnat_pool_external Resource - vyos"

subcategory: "Nat"

description: |- 
  nat⯯Carrier-grade NAT (CGNAT) parameters⯯External and internal pool parameters⯯External pool name
---

# vyos_nat_cgnat_pool_external (Resource)
<center>

*nat*  
⯯  
Carrier-grade NAT (CGNAT) parameters  
⯯  
External and internal pool parameters  
⯯  
**External pool name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `external_port_range` (String) Port range

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|range   &emsp;|Numbered port range (e.g., 1001-1005)  |
- `per_user_limit` (Attributes) Per user limits for the pool (see [below for nested schema](#nestedatt--per_user_limit))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `external` (String) External pool name

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|txt     &emsp;|External pool name  |


&lt;a id=&#34;nestedatt--per_user_limit&#34;&gt;&lt;/a&gt;
### Nested Schema for `per_user_limit`

Optional:

- `port` (Number) Ports per user

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
