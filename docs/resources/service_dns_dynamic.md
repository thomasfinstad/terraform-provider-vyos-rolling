---
page_title: "vyos_service_dns_dynamic Resource - vyos"

subcategory: "Service"

description: |-
  ~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.
  service⯯Domain Name System (DNS) related services⯯Dynamic DNS
---

# vyos_service_dns_dynamic (Resource)
<center>

~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*service*  
⯯  
Domain Name System (DNS) related services  
⯯  
**Dynamic DNS**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_service_dns_dynamic (Resource)](#vyos_service_dns_dynamic-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [interval](#interval)
      - [timeouts](#timeouts)
      - [vrf](#vrf)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### interval
- `interval` (Number) Interval in seconds to wait between Dynamic DNS updates

    |  Format   &emsp;|  Description      |
    |-----------|-------------------|
    |  60-3600  &emsp;|  Time in seconds  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
#### vrf
- `vrf` (String) VRF instance name

    |  Format  &emsp;|  Description        |
    |----------|---------------------|
    |  txt     &emsp;|  VRF instance name  |

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
terraform import vyos_service_dns_dynamic.example "service__dns__dynamic"
```
