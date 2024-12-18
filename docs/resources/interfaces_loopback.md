---
page_title: "vyos_interfaces_loopback Resource - vyos"

subcategory: "Interfaces"

description: |-
  interfaces⯯Loopback Interface
---

# vyos_interfaces_loopback (Resource)
<center>


*interfaces*  
⯯  
**Loopback Interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_interfaces_loopback (Resource)](#vyos_interfaces_loopback-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [address](#address)
      - [description](#description)
      - [ip](#ip)
      - [mirror](#mirror)
      - [redirect](#redirect)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `ip`](#nested-schema-for-ip)
    - [Nested Schema for `mirror`](#nested-schema-for-mirror)
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
- `address` (List of String) IP address

    |  Format   &emsp;|  Description                     |
    |-----------|----------------------------------|
    |  ipv4net  &emsp;|  IPv4 address and prefix length  |
    |  ipv6net  &emsp;|  IPv6 address and prefix length  |
#### description
- `description` (String) Description

    |  Format  &emsp;|  Description  |
    |----------|---------------|
    |  txt     &emsp;|  Description  |
#### ip
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
#### mirror
- `mirror` (Attributes) Mirror ingress/egress packets (see [below for nested schema](#nestedatt--mirror))
#### redirect
- `redirect` (String) Redirect incoming packet to destination

    |  Format  &emsp;|  Description                 |
    |----------|------------------------------|
    |  txt     &emsp;|  Destination interface name  |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `loopback` (String) Loopback Interface

    |  Format  &emsp;|  Description         |
    |----------|----------------------|
    |  lo      &emsp;|  Loopback interface  |


<a id="nestedatt--ip"></a>
### Nested Schema for `ip`

Optional:

- `source_validation` (String) Source validation by reversed path (RFC3704)

    |  Format   &emsp;|  Description                                                  |
    |-----------|---------------------------------------------------------------|
    |  strict   &emsp;|  Enable Strict Reverse Path Forwarding as defined in RFC3704  |
    |  loose    &emsp;|  Enable Loose Reverse Path Forwarding as defined in RFC3704   |
    |  disable  &emsp;|  No source validation                                         |


<a id="nestedatt--mirror"></a>
### Nested Schema for `mirror`

Optional:

- `egress` (String) Mirror egress traffic to destination interface

    |  Format  &emsp;|  Description                 |
    |----------|------------------------------|
    |  txt     &emsp;|  Destination interface name  |
- `ingress` (String) Mirror ingress traffic to destination interface

    |  Format  &emsp;|  Description                 |
    |----------|------------------------------|
    |  txt     &emsp;|  Destination interface name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_interfaces_loopback.example "interfaces__loopback__<loopback>"
```
