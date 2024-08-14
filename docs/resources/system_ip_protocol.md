---
page_title: "vyos_system_ip_protocol Resource - vyos"

subcategory: "System"

description: |- 
  system⯯IPv4 Settings⯯Filter routing info exchanged between routing protocol and zebra
---

# vyos_system_ip_protocol (Resource)
<center>

*system*  
⯯  
IPv4 Settings  
⯯  
**Filter routing info exchanged between routing protocol and zebra**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `route_map` (String) Specify route-map name to use

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Route map name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `protocol` (String) Filter routing info exchanged between routing protocol and zebra

    |Format     &emsp;|Description                                          |
    |-------------|-------------------------------------------------------|
    |any        &emsp;|Any of the above protocols                           |
    |babel      &emsp;|Babel routing protocol                               |
    |bgp        &emsp;|Border Gateway Protocol                              |
    |connected  &emsp;|Connected routes (directly attached subnet or host)  |
    |eigrp      &emsp;|Enhanced Interior Gateway Routing Protocol           |
    |isis       &emsp;|Intermediate System to Intermediate System           |
    |kernel     &emsp;|Kernel routes (not installed via the zebra RIB)      |
    |ospf       &emsp;|Open Shortest Path First (OSPFv2)                    |
    |rip        &emsp;|Routing Information Protocol                         |
    |static     &emsp;|Statically configured routes                         |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
