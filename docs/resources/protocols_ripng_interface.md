---
page_title: "vyos_protocols_ripng_interface Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Routing Information Protocol (RIPng) parameters⯯Interface name
---

# vyos_protocols_ripng_interface (Resource)
<center>

*protocols*  
⯯  
Routing Information Protocol (RIPng) parameters  
⯯  
**Interface name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `split_horizon` (Attributes) Split horizon parameters (see [below for nested schema](#nestedatt--split_horizon))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface name

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |


<a id="nestedatt--split_horizon"></a>
### Nested Schema for `split_horizon`

Optional:

- `disable` (Boolean) Disable instance
- `poison_reverse` (Boolean) Disable split horizon on specified interface


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
