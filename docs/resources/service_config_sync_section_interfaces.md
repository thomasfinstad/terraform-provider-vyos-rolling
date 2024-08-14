---
page_title: "vyos_service_config_sync_section_interfaces Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Configuration synchronization⯯Section for synchronization⯯Interfaces
---

# vyos_service_config_sync_section_interfaces (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Configuration synchronization  
⯯  
Section for synchronization  
⯯  
**Interfaces**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `bonding` (Boolean) Bonding interface
- `bridge` (Boolean) Bridge interface
- `dummy` (Boolean) Dummy interface
- `ethernet` (Boolean) Ethernet interface
- `geneve` (Boolean) GENEVE interface
- `input` (Boolean) Input interface
- `l2tpv3` (Boolean) L2TPv3 interface
- `loopback` (Boolean) Loopback interface
- `macsec` (Boolean) MACsec interface
- `openvpn` (Boolean) OpenVPN interface
- `pppoe` (Boolean) PPPoE interface
- `pseudo_ethernet` (Boolean) Pseudo-Ethernet interface
- `sstpc` (Boolean) SSTP client interface
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `tunnel` (Boolean) Tunnel interface
- `virtual_ethernet` (Boolean) Virtual Ethernet interface
- `vti` (Boolean) Virtual tunnel interface
- `vxlan` (Boolean) VXLAN interface
- `wireguard` (Boolean) Wireguard interface
- `wireless` (Boolean) Wireless interface
- `wwan` (Boolean) WWAN interface

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
