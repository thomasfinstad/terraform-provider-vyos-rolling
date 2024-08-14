---
page_title: "vyos_protocols_babel_distribute_list_ipv4_interface Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Babel Routing Protocol⯯Filter networks in routing updates⯯Filter IPv4 routes⯯Apply filtering to an interface
---

# vyos_protocols_babel_distribute_list_ipv4_interface (Resource)
<center>

*protocols*  
⯯  
Babel Routing Protocol  
⯯  
Filter networks in routing updates  
⯯  
Filter IPv4 routes  
⯯  
**Apply filtering to an interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `access_list` (Attributes) Access-list (see [below for nested schema](#nestedatt--access_list))
- `prefix_list` (Attributes) Prefix-list (see [below for nested schema](#nestedatt--prefix_list))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Apply filtering to an interface

    |Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    |txt     &emsp;|Apply filtering to an interface  |


<a id="nestedatt--access_list"></a>
### Nested Schema for `access_list`

Optional:

- `in` (Number) Access list to apply to input packets

    |Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    |u32     &emsp;|Access list to apply to input packets  |
- `out` (Number) Access list to apply to output packets

    |Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    |u32     &emsp;|Access list to apply to output packets  |


<a id="nestedatt--prefix_list"></a>
### Nested Schema for `prefix_list`

Optional:

- `in` (String) Prefix-list to apply to input packets

    |Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    |txt     &emsp;|Prefix-list to apply to input packets  |
- `out` (String) Prefix-list to apply to output packets

    |Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    |txt     &emsp;|Prefix-list to apply to output packets  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
