---
page_title: "vyos_nat_cgnat_rule Resource - vyos"

subcategory: "Nat"

description: |- 
  nat⯯Carrier-grade NAT (CGNAT) parameters⯯Rule
---

# vyos_nat_cgnat_rule (Resource)
<center>

*nat*  
⯯  
Carrier-grade NAT (CGNAT) parameters  
⯯  
**Rule**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `source` (Attributes) Source parameters (see [below for nested schema](#nestedatt--source))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `translation` (Attributes) Translation parameters (see [below for nested schema](#nestedatt--translation))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `rule` (Number) Rule

    &emsp;|Format    &emsp;|Description                 |
    |------------|------------------------------|
    &emsp;|1-999999  &emsp;|Number for this CGNAT rule  |


&lt;a id=&#34;nestedatt--source&#34;&gt;&lt;/a&gt;
### Nested Schema for `source`

Optional:

- `pool` (String) Source internal pool

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|txt     &emsp;|Source internal pool name  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--translation&#34;&gt;&lt;/a&gt;
### Nested Schema for `translation`

Optional:

- `pool` (String) Translation external pool

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|txt     &emsp;|Translation external pool name  |  
