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

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|1-32    &emsp;|Number of paths  |
- `network` (List of String) Enable routing on an IP network

    &emsp;|Format   &emsp;|Description           |
    |-----------|------------------------|
    &emsp;|ipv4net  &emsp;|EIGRP network prefix  |
- `passive_interface` (List of String) Suppress routing updates on an interface
- `redistribute` (List of String) Redistribute information from another routing protocol

    &emsp;|Format     &emsp;|Description                          |
    |-------------|---------------------------------------|
    &emsp;|bgp        &emsp;|Border Gateway Protocol (BGP)        |
    &emsp;|connected  &emsp;|Connected routes                     |
    &emsp;|nhrp       &emsp;|Next Hop Resolution Protocol (NHRP)  |
    &emsp;|ospf       &emsp;|Open Shortest Path First (OSPFv2)    |
    &emsp;|rip        &emsp;|Routing Information Protocol (RIP)   |
    &emsp;|babel      &emsp;|Babel routing protocol (Babel)       |
    &emsp;|static     &emsp;|Statically configured routes         |
    &emsp;|vnc        &emsp;|Virtual Network Control (VNC)        |
- `router_id` (String) Override default router identifier

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv4    &emsp;|Router-ID in IP address format  |
- `system_as` (Number) Autonomous System Number (ASN)

    &emsp;|Format   &emsp;|Description               |
    |-----------|----------------------------|
    &emsp;|1-65535  &emsp;|Autonomous System Number  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `variance` (Number) Control load balancing variance

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|1-128   &emsp;|Metric variance multiplier  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
