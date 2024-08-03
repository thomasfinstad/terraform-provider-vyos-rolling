---
page_title: "vyos_service_snmp Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Simple Network Management Protocol (SNMP)
---

# vyos_service_snmp (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
**Simple Network Management Protocol (SNMP)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `contact` (String) Contact information
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `location` (String) Location information
- `oid_enable` (List of String) Enable specific OIDs that by default are disable

    &emsp;|Format                           &emsp;|Description                                           |
    |-----------------------------------|--------------------------------------------------------|
    &emsp;|ip-forward                       &emsp;|Enable ipForward: .1.3.6.1.2.1.4.24                   |
    &emsp;|ip-route-table                   &emsp;|Enable ipRouteTable: .1.3.6.1.2.1.4.21                |
    &emsp;|ip-net-to-media-table            &emsp;|Enable ipNetToMediaTable: .1.3.6.1.2.1.4.22           |
    &emsp;|ip-net-to-physical-phys-address  &emsp;|Enable ipNetToPhysicalPhysAddress: .1.3.6.1.2.1.4.35  |
- `protocol` (String) Protocol to be used (TCP/UDP)

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|udp     &emsp;|Listen protocol UDP  |
    &emsp;|tcp     &emsp;|Listen protocol TCP  |
- `smux_peer` (List of String) Register a subtree for SMUX-based processing

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|txt     &emsp;|SNMP Object Identifier  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `trap_source` (String) SNMP trap source address

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|ipv4    &emsp;|IPv4 address  |
    &emsp;|ipv6    &emsp;|IPv6 address  |
- `vrf` (String) VRF instance name

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
