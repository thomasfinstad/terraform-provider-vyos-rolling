---
page_title: "vyos_service_config_sync_section_protocols Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Configuration synchronization⯯Section for synchronization⯯Routing protocols
---

# vyos_service_config_sync_section_protocols (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Configuration synchronization  
⯯  
Section for synchronization  
⯯  
**Routing protocols**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `babel` (Boolean) Babel Routing Protocol
- `bfd` (Boolean) Bidirectional Forwarding Detection (BFD)
- `bgp` (Boolean) Border Gateway Protocol (BGP)
- `failover` (Boolean) Failover route
- `igmp_proxy` (Boolean) Internet Group Management Protocol (IGMP) proxy
- `isis` (Boolean) Intermediate System to Intermediate System (IS-IS)
- `mpls` (Boolean) Multiprotocol Label Switching (MPLS)
- `nhrp` (Boolean) Next Hop Resolution Protocol (NHRP) parameters
- `ospf` (Boolean) Open Shortest Path First (OSPF)
- `ospfv3` (Boolean) Open Shortest Path First (OSPF) for IPv6
- `pim` (Boolean) Protocol Independent Multicast (PIM) and IGMP
- `pim6` (Boolean) Protocol Independent Multicast for IPv6 (PIMv6) and MLD
- `rip` (Boolean) Routing Information Protocol (RIP) parameters
- `ripng` (Boolean) Routing Information Protocol (RIPng) parameters
- `rpki` (Boolean) Resource Public Key Infrastructure (RPKI)
- `segment_routing` (Boolean) Segment Routing
- `static` (Boolean) Static Routing
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
