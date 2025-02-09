---
page_title: "vyos_service_dhcpv6_relay Resource - vyos"

subcategory: "Service"

description: |-
  ~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.
  service⯯DHCPv6 Relay Agent parameters
---

# vyos_service_dhcpv6_relay (Resource)
<center>

~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*service*  
⯯  
**DHCPv6 Relay Agent parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_service_dhcpv6_relay (Resource)](#vyos_service_dhcpv6_relay-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [disable](#disable)
      - [max_hop_count](#max_hop_count)
      - [timeouts](#timeouts)
      - [use_interface_id_option](#use_interface_id_option)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### disable
- `disable` (Boolean) Disable instance
#### max_hop_count
- `max_hop_count` (Number) Maximum hop count for which requests will be processed

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  1-255   &emsp;|  Hop count    |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
#### use_interface_id_option
- `use_interface_id_option` (Boolean) Option to set DHCPv6 interface-ID option

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
terraform import vyos_service_dhcpv6_relay.example "service__dhcpv6-relay"
```
