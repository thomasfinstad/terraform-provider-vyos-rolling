---
page_title: "vyos_protocols_eigrp Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Enhanced Interior Gateway Routing Protocol (EIGRP)
---

# vyos_protocols_eigrp (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
**Enhanced Interior Gateway Routing Protocol (EIGRP)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `maximum_paths` (Number) Forward packets over multiple paths

    |Format  &emsp;|Description      |
    |----------|-------------------|
    |1-32    &emsp;|Number of paths  |
- `network` (List of String) Enable routing on an IP network

    |Format   &emsp;|Description           |
    |-----------|------------------------|
    |ipv4net  &emsp;|EIGRP network prefix  |
- `passive_interface` (List of String) Suppress routing updates on an interface
- `redistribute` (List of String) Redistribute information from another routing protocol

    |Format     &emsp;|Description                          |
    |-------------|---------------------------------------|
    |bgp        &emsp;|Border Gateway Protocol (BGP)        |
    |connected  &emsp;|Connected routes                     |
    |nhrp       &emsp;|Next Hop Resolution Protocol (NHRP)  |
    |ospf       &emsp;|Open Shortest Path First (OSPFv2)    |
    |rip        &emsp;|Routing Information Protocol (RIP)   |
    |babel      &emsp;|Babel routing protocol (Babel)       |
    |static     &emsp;|Statically configured routes         |
    |vnc        &emsp;|Virtual Network Control (VNC)        |
- `router_id` (String) Override default router identifier

    |Format  &emsp;|Description                     |
    |----------|----------------------------------|
    |ipv4    &emsp;|Router-ID in IP address format  |
- `system_as` (Number) Autonomous System Number (ASN)

    |Format   &emsp;|Description               |
    |-----------|----------------------------|
    |1-65535  &emsp;|Autonomous System Number  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `variance` (Number) Control load balancing variance

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |1-128   &emsp;|Metric variance multiplier  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
