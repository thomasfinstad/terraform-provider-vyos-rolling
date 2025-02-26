---
page_title: "vyos_protocols_nhrp_tunnel_map_tunnel_ip Resource - vyos"

subcategory: "Protocols"

description: |-
  protocols⯯Next Hop Resolution Protocol (NHRP) parameters⯯Tunnel for NHRP⯯Map tunnel IP to NBMA⯯Set a NHRP tunnel address
---

# vyos_protocols_nhrp_tunnel_map_tunnel_ip (Resource)
<center>


*protocols*  
⯯  
Next Hop Resolution Protocol (NHRP) parameters  
⯯  
Tunnel for NHRP  
⯯  
Map tunnel IP to NBMA  
⯯  
**Set a NHRP tunnel address**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_protocols_nhrp_tunnel_map_tunnel_ip (Resource)](#vyos_protocols_nhrp_tunnel_map_tunnel_ip-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [nbma](#nbma)
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

#### nbma
- `nbma` (String) Set NHRP NBMA address to map

    |  Format  &emsp;|  Description                |
    |----------|-----------------------------|
    |  ipv4    &emsp;|  Set the IP address to map  |
    |  local   &emsp;|  Set the local address      |
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `tunnel` (String) Tunnel for NHRP

    |  Format  &emsp;|  Description       |
    |----------|--------------------|
    |  tunN    &emsp;|  NHRP tunnel name  |
- `tunnel_ip` (String) Set a NHRP tunnel address

    |  Format  &emsp;|  Description                |
    |----------|-----------------------------|
    |  ipv4    &emsp;|  Set the IP address to map  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_protocols_nhrp_tunnel_map_tunnel_ip.example "protocols__nhrp__tunnel__<tunnel>__map__tunnel-ip__<tunnel-ip>"
```
