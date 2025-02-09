---
page_title: "vyos_protocols_isis_fast_reroute_lfa_remote_prefix_list Resource - vyos"

subcategory: "Protocols"

description: |-
  protocols⯯Intermediate System to Intermediate System (IS-IS)⯯IS-IS fast reroute configuration⯯Loop free alternate functionality⯯Remote loop free alternate options⯯Filter PQ node router ID based on prefix list
---

# vyos_protocols_isis_fast_reroute_lfa_remote_prefix_list (Resource)
<center>


*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
IS-IS fast reroute configuration  
⯯  
Loop free alternate functionality  
⯯  
Remote loop free alternate options  
⯯  
**Filter PQ node router ID based on prefix list**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_protocols_isis_fast_reroute_lfa_remote_prefix_list (Resource)](#vyos_protocols_isis_fast_reroute_lfa_remote_prefix_list-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [level_1](#level_1)
      - [level_2](#level_2)
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

#### level_1
- `level_1` (Boolean) Match on IS-IS level-1 routes
#### level_2
- `level_2` (Boolean) Match on IS-IS level-2 routes
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `prefix_list` (String) Filter PQ node router ID based on prefix list

    |  Format  &emsp;|  Description                    |
    |----------|---------------------------------|
    |  txt     &emsp;|  Name of IPv4/IPv6 prefix-list  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_protocols_isis_fast_reroute_lfa_remote_prefix_list.example "protocols__isis__fast-reroute__lfa__remote__prefix-list__<prefix-list>"
```
