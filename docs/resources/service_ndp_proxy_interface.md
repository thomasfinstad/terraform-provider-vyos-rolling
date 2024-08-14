---
page_title: "vyos_service_ndp_proxy_interface Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Neighbor Discovery Protocol (NDP) Proxy⯯NDP proxy listener interface
---

# vyos_service_ndp_proxy_interface (Resource)
<center>

*service*  
⯯  
Neighbor Discovery Protocol (NDP) Proxy  
⯯  
**NDP proxy listener interface**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `disable` (Boolean) Disable instance
- `enable_router_bit` (Boolean) Enable router bit in Neighbor Advertisement messages
- `timeout` (Number) Timeout for Neighbor Advertisement after Neighbor Solicitation message

    |Format      &emsp;|Description              |
    |--------------|---------------------------|
    |500-120000  &emsp;|Timeout in milliseconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `ttl` (Number) Proxy entry cache Time-To-Live

    |Format        &emsp;|Description           |
    |----------------|------------------------|
    |10000-120000  &emsp;|Time in milliseconds  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) NDP proxy listener interface


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
