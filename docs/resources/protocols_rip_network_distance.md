---
page_title: "vyos_protocols_rip_network_distance Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Routing Information Protocol (RIP) parameters⯯Source network
---

# vyos_protocols_rip_network_distance (Resource)
<center>

*protocols*  
⯯  
Routing Information Protocol (RIP) parameters  
⯯  
**Source network**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `access_list` (String) Access list

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Access list  |
- `distance` (Number) Distance for this route

    |Format  &emsp;|Description              |
    |----------|---------------------------|
    |1-255   &emsp;|Distance for this route  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `network_distance` (String) Source network

    |Format   &emsp;|Description     |
    |-----------|------------------|
    |ipv4net  &emsp;|Source network  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
