---
page_title: "vyos_nat_cgnat_pool_internal Resource - vyos"

subcategory: "Nat"

description: |-
  Network Address Translation (NAT) parameters⯯Carrier-grade NAT (CGNAT) parameters⯯External and internal pool parameters⯯Internal pool name
---

# vyos_nat_cgnat_pool_internal (Resource)
<center>


Network Address Translation (NAT) parameters  
⯯  
Carrier-grade NAT (CGNAT) parameters  
⯯  
External and internal pool parameters  
⯯  
**Internal pool name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_nat_cgnat_pool_internal (Resource)](#vyos_nat_cgnat_pool_internal-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [range](#range)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### range
- `range` (List of String) Range of IP addresses

    |  Format     &emsp;|  Description         |
    |-------------|----------------------|
    |  ipv4net    &emsp;|  IPv4 prefix         |
    |  ipv4range  &emsp;|  IPv4 address range  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `internal` (String) Internal pool name

    |  Format  &emsp;|  Description         |
    |----------|----------------------|
    |  txt     &emsp;|  Internal pool name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_nat_cgnat_pool_internal.example "nat__cgnat__pool__internal__<internal>"
```
