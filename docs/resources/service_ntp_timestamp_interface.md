---
page_title: "vyos_service_ntp_timestamp_interface Resource - vyos"

subcategory: "Service"

description: |-
  service⯯Network Time Protocol (NTP) configuration⯯Enable timestamping of packets in the NIC hardware⯯Interface to enable timestamping on
---

# vyos_service_ntp_timestamp_interface (Resource)
<center>


*service*  
⯯  
Network Time Protocol (NTP) configuration  
⯯  
Enable timestamping of packets in the NIC hardware  
⯯  
**Interface to enable timestamping on**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_service_ntp_timestamp_interface (Resource)](#vyos_service_ntp_timestamp_interface-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [receive_filter](#receive_filter)
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

#### receive_filter
- `receive_filter` (String) Selects which inbound packets are timestamped by the NIC

    |  Format  &emsp;|  Description                                                      |
    |----------|-------------------------------------------------------------------|
    |  all     &emsp;|  All packets are timestamped                                      |
    |  ntp     &emsp;|  Only NTP packets are timestamped                                 |
    |  ptp     &emsp;|  Only PTP or NTP packets using the PTP transport are timestamped  |
    |  none    &emsp;|  No packet is timestamped                                         |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface to enable timestamping on

    |  Format  &emsp;|  Description            |
    |----------|-------------------------|
    |  all     &emsp;|  Select all interfaces  |
    |  txt     &emsp;|  Interface name         |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_service_ntp_timestamp_interface.example "service__ntp__timestamp__interface__<interface>"
```