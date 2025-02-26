---
page_title: "vyos_protocols_static_table_route6 Resource - vyos"

subcategory: "Protocols"

description: |-
  protocols⯯Static Routing⯯Non-main Kernel Routing Table⯯Static IPv6 route
---

# vyos_protocols_static_table_route6 (Resource)
<center>


*protocols*  
⯯  
Static Routing  
⯯  
Non-main Kernel Routing Table  
⯯  
**Static IPv6 route**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_protocols_static_table_route6 (Resource)](#vyos_protocols_static_table_route6-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [blackhole](#blackhole)
      - [description](#description)
      - [reject](#reject)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `blackhole`](#nested-schema-for-blackhole)
    - [Nested Schema for `reject`](#nested-schema-for-reject)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### blackhole
- `blackhole` (Attributes) Silently discard pkts when matched (see [below for nested schema](#nestedatt--blackhole))
#### description
- `description` (String) Description

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  txt     &emsp;|  Description  |
#### reject
- `reject` (Attributes) Emit an ICMP unreachable when matched (see [below for nested schema](#nestedatt--reject))
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `route6` (String) Static IPv6 route

    |  Format   &emsp;|  Description        |
    |-----------|---------------------|
    |  ipv6net  &emsp;|  IPv6 static route  |
- `table` (Number) Non-main Kernel Routing Table

    |  Format  &emsp;|  Description                |
    |----------|-----------------------------|
    |  1-200   &emsp;|  Policy route table number  |


<a id="nestedatt--blackhole"></a>
### Nested Schema for `blackhole`

Optional:

- `distance` (Number) Distance for this route

    |  Format  &emsp;|  Description              |
    |----------|---------------------------|
    |  1-255   &emsp;|  Distance for this route  |
- `tag` (Number) Tag value for this route

    |  Format        &emsp;|  Description               |
    |----------------|----------------------------|
    |  1-4294967295  &emsp;|  Tag value for this route  |


<a id="nestedatt--reject"></a>
### Nested Schema for `reject`

Optional:

- `distance` (Number) Distance for this route

    |  Format  &emsp;|  Description              |
    |----------|---------------------------|
    |  1-255   &emsp;|  Distance for this route  |
- `tag` (Number) Tag value for this route

    |  Format        &emsp;|  Description               |
    |----------------|----------------------------|
    |  1-4294967295  &emsp;|  Tag value for this route  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_protocols_static_table_route6.example "protocols__static__table__<table>__route6__<route6>"
```
