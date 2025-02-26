---
page_title: "vyos_service_broadcast_relay_id Resource - vyos"

subcategory: "Service"

description: |-
  service⯯UDP broadcast relay service⯯Unique ID for each UDP port to forward
---

# vyos_service_broadcast_relay_id (Resource)
<center>


*service*  
⯯  
UDP broadcast relay service  
⯯  
**Unique ID for each UDP port to forward**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_service_broadcast_relay_id (Resource)](#vyos_service_broadcast_relay_id-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [address](#address)
      - [description](#description)
      - [disable](#disable)
      - [interface](#interface)
      - [port](#port)
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

#### address
- `address` (String) Set source IP of forwarded packets, otherwise original senders address is used

    |  Format  &emsp;|  Description                                    |
    |----------|-------------------------------------------------|
    |  ipv4    &emsp;|  Optional source address for forwarded packets  |
#### description
- `description` (String) Description

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  txt     &emsp;|  Description  |
#### disable
- `disable` (Boolean) Disable instance
#### interface
- `interface` (List of String) Interface

    |  Format  &emsp;|  Description     |
    |----------|------------------|
    |  txt     &emsp;|  Interface name  |
#### port
- `port` (Number) Port number used by connection

    |  Format   &emsp;|  Description      |
    |-----------|-------------------|
    |  1-65535  &emsp;|  Numeric IP port  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `id` (Number) Unique ID for each UDP port to forward

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  1-99    &emsp;|  Broadcast relay instance ID  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_service_broadcast_relay_id.example "service__broadcast-relay__id__<id>"
```
