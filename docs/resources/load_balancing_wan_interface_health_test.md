---
page_title: "vyos_load_balancing_wan_interface_health_test Resource - vyos"

subcategory: "Load Balancing"

description: |-
  load-balancing⯯Configure Wide Area Network (WAN) load-balancing⯯Interface name⯯Rule number
---

# vyos_load_balancing_wan_interface_health_test (Resource)
<center>


*load-balancing*  
⯯  
Configure Wide Area Network (WAN) load-balancing  
⯯  
Interface name  
⯯  
**Rule number**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_load_balancing_wan_interface_health_test (Resource)](#vyos_load_balancing_wan_interface_health_test-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [resp_time](#resp_time)
      - [target](#target)
      - [test_script](#test_script)
      - [timeouts](#timeouts)
      - [ttl_limit](#ttl_limit)
      - [type](#type)
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

#### resp_time
- `resp_time` (Number) Ping response time (seconds)

    |  Format  &emsp;|  Description              |
    |----------|---------------------------|
    |  1-30    &emsp;|  Response time (seconds)  |
#### target
- `target` (String) Health target address

    |  Format  &emsp;|  Description            |
    |----------|-------------------------|
    |  ipv4    &emsp;|  Health target address  |
#### test_script
- `test_script` (String) Path to user-defined script

    |  Format  &emsp;|  Description                |
    |----------|-----------------------------|
    |  txt     &emsp;|  Script in /config/scripts  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
#### ttl_limit
- `ttl_limit` (Number) TTL limit (hop count)

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  1-254   &emsp;|  Number of hops  |
#### type
- `type` (String) WLB test type

    |  Format        &emsp;|  Description                         |
    |----------------|--------------------------------------|
    |  ping          &emsp;|  Test with ICMP echo response        |
    |  ttl           &emsp;|  Test with UDP TTL expired response  |
    |  user-defined  &emsp;|  User-defined test script            |

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface_health` (String) Interface name
- `test` (Number) Rule number

    |  Format        &emsp;|  Description  |
    |----------------|---------------|
    |  0-4294967295  &emsp;|  Rule number  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_load_balancing_wan_interface_health_test.example "load-balancing__wan__interface-health__<interface-health>__test__<test>"
```
