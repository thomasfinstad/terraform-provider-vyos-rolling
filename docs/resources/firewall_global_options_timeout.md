---
page_title: "vyos_firewall_global_options_timeout Resource - vyos"

subcategory: "Firewall"

description: |-
  ~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.
  Firewall⯯Global Options⯯Connection timeout options
---

# vyos_firewall_global_options_timeout (Resource)
<center>

~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

Firewall  
⯯  
Global Options  
⯯  
**Connection timeout options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_firewall_global_options_timeout (Resource)](#vyos_firewall_global_options_timeout-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [icmp](#icmp)
      - [other](#other)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### icmp
- `icmp` (Number) ICMP timeout in seconds

    |  Format      &emsp;|  Description              |
    |--------------|---------------------------|
    |  1-21474836  &emsp;|  ICMP timeout in seconds  |
#### other
- `other` (Number) Generic connection timeout in seconds

    |  Format      &emsp;|  Description                            |
    |--------------|-----------------------------------------|
    |  1-21474836  &emsp;|  Generic connection timeout in seconds  |
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
terraform import vyos_firewall_global_options_timeout.example "firewall__global-options__timeout"
```
