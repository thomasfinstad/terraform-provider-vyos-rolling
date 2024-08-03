---
page_title: "vyos_service_config_sync_section_service Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Configuration synchronization⯯Section for synchronization⯯System services
---

# vyos_service_config_sync_section_service (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Configuration synchronization  
⯯  
Section for synchronization  
⯯  
**System services**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `console_server` (Boolean) Serial Console Server
- `dhcp_relay` (Boolean) Host Configuration Protocol (DHCP) relay agent
- `dhcp_server` (Boolean) Dynamic Host Configuration Protocol (DHCP) for DHCP server
- `dhcpv6_relay` (Boolean) DHCPv6 Relay Agent parameters
- `dhcpv6_server` (Boolean) DHCP for IPv6 (DHCPv6) server
- `dns` (Boolean) Domain Name System (DNS) related services
- `lldp` (Boolean) LLDP settings
- `mdns` (Boolean) Multicast DNS (mDNS) parameters
- `monitoring` (Boolean) Monitoring services
- `ndp_proxy` (Boolean) Neighbor Discovery Protocol (NDP) Proxy
- `ntp` (Boolean) Network Time Protocol (NTP) configuration
- `snmp` (Boolean) Simple Network Management Protocol (SNMP)
- `tftp_server` (Boolean) Trivial File Transfer Protocol (TFTP) server
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `webproxy` (Boolean) Webproxy service settings

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
