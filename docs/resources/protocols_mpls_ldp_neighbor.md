---
page_title: "vyos_protocols_mpls_ldp_neighbor Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Multiprotocol Label Switching (MPLS)⯯Label Distribution Protocol (LDP)⯯LDP neighbor parameters
---

# vyos_protocols_mpls_ldp_neighbor (Resource)
<center>

*protocols*  
⯯  
Multiprotocol Label Switching (MPLS)  
⯯  
Label Distribution Protocol (LDP)  
⯯  
**LDP neighbor parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `password` (String) Neighbor password
- `session_holdtime` (Number) Session IPv4 hold time

    |Format    &emsp;|Description      |
    |------------|-------------------|
    |15-65535  &emsp;|Time in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `ttl_security` (String) Neighbor TTL security

    |Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    |1-254    &emsp;|TTL                            |
    |disable  &emsp;|Disable neighbor TTL security  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `neighbor` (String) LDP neighbor parameters

    |Format  &emsp;|Description            |
    |----------|-------------------------|
    |ipv4    &emsp;|Neighbor IPv4 address  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
