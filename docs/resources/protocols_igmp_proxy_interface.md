---
page_title: "vyos_protocols_igmp_proxy_interface Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Internet Group Management Protocol (IGMP) proxy parameters⯯Interface for IGMP proxy
---

# vyos_protocols_igmp_proxy_interface (Resource)
<center>

*protocols*  
⯯  
Internet Group Management Protocol (IGMP) proxy parameters  
⯯  
**Interface for IGMP proxy**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `alt_subnet` (List of String) Unicast source networks allowed for multicast traffic to be proxyed

    |Format   &emsp;|Description   |
    |-----------|----------------|
    |ipv4net  &emsp;|IPv4 network  |
- `role` (String) IGMP interface role

    |Format      &emsp;|Description                          |
    |--------------|---------------------------------------|
    |upstream    &emsp;|Upstream interface (only 1 allowed)  |
    |downstream  &emsp;|Downstream interface(s)              |
    |disabled    &emsp;|Disabled interface                   |
- `threshold` (Number) TTL threshold

    |Format  &emsp;|Description                       |
    |----------|------------------------------------|
    |1-255   &emsp;|TTL threshold for the interfaces  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `whitelist` (List of String) Group to whitelist

    |Format   &emsp;|Description   |
    |-----------|----------------|
    |ipv4net  &emsp;|IPv4 network  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface for IGMP proxy


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
