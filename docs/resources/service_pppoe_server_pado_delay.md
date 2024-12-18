---
page_title: "vyos_service_pppoe_server_pado_delay Resource - vyos"

subcategory: "Service"

description: |-
  service⯯Point to Point over Ethernet (PPPoE) Server⯯PADO delays
---

# vyos_service_pppoe_server_pado_delay (Resource)
<center>


*service*  
⯯  
Point to Point over Ethernet (PPPoE) Server  
⯯  
**PADO delays**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_service_pppoe_server_pado_delay (Resource)](#vyos_service_pppoe_server_pado_delay-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [sessions](#sessions)
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

#### sessions
- `sessions` (Number) Number of sessions

    |  Format    &emsp;|  Description         |
    |------------|----------------------|
    |  1-999999  &emsp;|  Number of sessions  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `pado_delay` (String) PADO delays

    |  Format    &emsp;|  Description              |
    |------------|---------------------------|
    |  disable   &emsp;|  Disable new connections  |
    |  1-999999  &emsp;|  Number in ms             |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_service_pppoe_server_pado_delay.example "service__pppoe-server__pado-delay__<pado-delay>"
```
