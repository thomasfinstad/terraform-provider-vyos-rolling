---
page_title: "vyos_service_tftp_server_listen_address Resource - vyos"

subcategory: "Service"

description: |-
  service⯯Trivial File Transfer Protocol (TFTP) server⯯Local IP addresses to listen on
---

# vyos_service_tftp_server_listen_address (Resource)
<center>


*service*  
⯯  
Trivial File Transfer Protocol (TFTP) server  
⯯  
**Local IP addresses to listen on**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_service_tftp_server_listen_address (Resource)](#vyos_service_tftp_server_listen_address-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [timeouts](#timeouts)
      - [vrf](#vrf)
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

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `listen_address` (String) Local IP addresses to listen on

    |  Format  &emsp;|  Description                                      |
    |----------|---------------------------------------------------|
    |  ipv4    &emsp;|  IPv4 address to listen for incoming connections  |
    |  ipv6    &emsp;|  IPv6 address to listen for incoming connections  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_service_tftp_server_listen_address.example "service__tftp-server__listen-address__<listen-address>"
```
