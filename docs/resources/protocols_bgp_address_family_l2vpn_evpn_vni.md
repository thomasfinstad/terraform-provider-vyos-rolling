---
page_title: "vyos_protocols_bgp_address_family_l2vpn_evpn_vni Resource - vyos"

subcategory: "Protocols"

description: |-
  protocols⯯Border Gateway Protocol (BGP)⯯BGP address-family parameters⯯L2VPN EVPN BGP settings⯯VXLAN Network Identifier
---

# vyos_protocols_bgp_address_family_l2vpn_evpn_vni (Resource)
<center>


*protocols*  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP address-family parameters  
⯯  
L2VPN EVPN BGP settings  
⯯  
**VXLAN Network Identifier**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_protocols_bgp_address_family_l2vpn_evpn_vni (Resource)](#vyos_protocols_bgp_address_family_l2vpn_evpn_vni-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [advertise_default_gw](#advertise_default_gw)
      - [advertise_svi_ip](#advertise_svi_ip)
      - [rd](#rd)
      - [route_target](#route_target)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `route_target`](#nested-schema-for-route_target)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### advertise_default_gw
- `advertise_default_gw` (Boolean) Advertise All default g/w mac-ip routes in EVPN
#### advertise_svi_ip
- `advertise_svi_ip` (Boolean) Advertise svi mac-ip routes in EVPN
#### rd
- `rd` (String) Route Distinguisher

    |  Format                   &emsp;|  Description                                   |
    |---------------------------|------------------------------------------------|
    |  ASN:NN_OR_IP-ADDRESS:NN  |  Route Distinguisher, (x.x.x.x:yyy&emsp;|xxxx:yyyy)  |
#### route_target
- `route_target` (Attributes) Route Target (see [below for nested schema](#nestedatt--route_target))
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `vni` (Number) VXLAN Network Identifier

    |  Format      &emsp;|  Description  |
    |--------------|---------------|
    |  1-16777215  &emsp;|  VNI number   |


<a id="nestedatt--route_target"></a>
### Nested Schema for `route_target`

Optional:

- `both` (List of String) Route Target both import and export

    |  Format  &emsp;|  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |
- `export` (List of String) Route Target export

    |  Format  &emsp;|  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |
- `import` (List of String) Route Target import

    |  Format  &emsp;|  Description                                |
    |----------|---------------------------------------------|
    |  txt     |  Route target (A.B.C.D:MN|EF:OPQR&emsp;|GHJK:MN)  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_protocols_bgp_address_family_l2vpn_evpn_vni.example "protocols__bgp__address-family__l2vpn-evpn__vni__<vni>"
```
