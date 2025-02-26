---
page_title: "vyos_service_config_sync_section_interfaces Resource - vyos"

subcategory: "Service"

description: |-
  ~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.
  service⯯Configuration synchronization⯯Section for synchronization⯯Interfaces
---

# vyos_service_config_sync_section_interfaces (Resource)
<center>

~> This resource is global, having more than one resource of this type is likely to cause configuration drift / conflicts.

*service*  
⯯  
Configuration synchronization  
⯯  
Section for synchronization  
⯯  
**Interfaces**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

<!--TOC-->

- [vyos_service_config_sync_section_interfaces (Resource)](#vyos_service_config_sync_section_interfaces-resource)
  - [Schema](#schema)
    - [Optional](#optional)
      - [bonding](#bonding)
      - [bridge](#bridge)
      - [dummy](#dummy)
      - [ethernet](#ethernet)
      - [geneve](#geneve)
      - [input](#input)
      - [l2tpv3](#l2tpv3)
      - [loopback](#loopback)
      - [macsec](#macsec)
      - [openvpn](#openvpn)
      - [pppoe](#pppoe)
      - [pseudo_ethernet](#pseudo_ethernet)
      - [sstpc](#sstpc)
      - [timeouts](#timeouts)
      - [tunnel](#tunnel)
      - [virtual_ethernet](#virtual_ethernet)
      - [vti](#vti)
      - [vxlan](#vxlan)
      - [wireguard](#wireguard)
      - [wireless](#wireless)
      - [wwan](#wwan)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

#### bonding
- `bonding` (Boolean) Bonding interface
#### bridge
- `bridge` (Boolean) Bridge interface
#### dummy
- `dummy` (Boolean) Dummy interface
#### ethernet
- `ethernet` (Boolean) Ethernet interface
#### geneve
- `geneve` (Boolean) GENEVE interface
#### input
- `input` (Boolean) Input interface
#### l2tpv3
- `l2tpv3` (Boolean) L2TPv3 interface
#### loopback
- `loopback` (Boolean) Loopback interface
#### macsec
- `macsec` (Boolean) MACsec interface
#### openvpn
- `openvpn` (Boolean) OpenVPN interface
#### pppoe
- `pppoe` (Boolean) PPPoE interface
#### pseudo_ethernet
- `pseudo_ethernet` (Boolean) Pseudo-Ethernet interface
#### sstpc
- `sstpc` (Boolean) SSTP client interface
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
#### tunnel
- `tunnel` (Boolean) Tunnel interface
#### virtual_ethernet
- `virtual_ethernet` (Boolean) Virtual Ethernet interface
#### vti
- `vti` (Boolean) Virtual tunnel interface
#### vxlan
- `vxlan` (Boolean) VXLAN interface
#### wireguard
- `wireguard` (Boolean) Wireguard interface
#### wireless
- `wireless` (Boolean) Wireless interface
#### wwan
- `wwan` (Boolean) WWAN interface

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
terraform import vyos_service_config_sync_section_interfaces.example "service__config-sync__section__interfaces"
```
