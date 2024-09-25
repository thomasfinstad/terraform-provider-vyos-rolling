---
page_title: "vyos_protocols_ospfv3_graceful_restart Resource - vyos"

subcategory: "Protocols"

description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Open Shortest Path First (OSPF) for IPv6⯯Graceful Restart
---

# vyos_protocols_ospfv3_graceful_restart (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Open Shortest Path First (OSPF) for IPv6  
⯯  
**Graceful Restart**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_protocols_ospfv3_graceful_restart (Resource)](#vyos_protocols_ospfv3_graceful_restart-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [grace_period](#grace_period)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### grace_period
- `grace_period` (Number) Maximum length of the grace period

    |  Format  &emsp;|  Description                                    |
    |----------|-------------------------------------------------|
    |  1-1800  &emsp;|  Maximum length of the grace period in seconds  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).