---
page_title: "vyos_protocols_bgp_parameters_dampening Resource - vyos"

subcategory: "Protocols"

description: |-
  ~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.
  protocols⯯Border Gateway Protocol (BGP)⯯BGP parameters⯯Enable route-flap dampening
---

# vyos_protocols_bgp_parameters_dampening (Resource)
<center>

~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP parameters  
⯯  
**Enable route-flap dampening**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_protocols_bgp_parameters_dampening (Resource)](#vyos_protocols_bgp_parameters_dampening-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [half_life](#half_life)
      - [max_suppress_time](#max_suppress_time)
      - [re_use](#re_use)
      - [start_suppress_time](#start_suppress_time)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### half_life
- `half_life` (Number) Half-life time for dampening

    |  Format  &emsp;|  Description                   |
    |----------|--------------------------------|
    |  1-45    &emsp;|  Half-life penalty in minutes  |
#### max_suppress_time
- `max_suppress_time` (Number) Maximum duration to suppress a stable route

    |  Format  &emsp;|  Description                           |
    |----------|----------------------------------------|
    |  1-255   &emsp;|  Maximum suppress duration in minutes  |
#### re_use
- `re_use` (Number) Threshold to start reusing a route

    |  Format   &emsp;|  Description            |
    |-----------|-------------------------|
    |  1-20000  &emsp;|  Re-use penalty points  |
#### start_suppress_time
- `start_suppress_time` (Number) When to start suppressing a route

    |  Format   &emsp;|  Description                    |
    |-----------|---------------------------------|
    |  1-20000  &emsp;|  Start-suppress penalty points  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_protocols_bgp_parameters_dampening.example "protocols__bgp__parameters__dampening"
```
