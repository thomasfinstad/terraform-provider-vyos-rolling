---
page_title: "vyos_protocols_pim_rp_address Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Protocol Independent Multicast (PIM) and IGMP⯯Rendezvous Point⯯Rendezvous Point address
---

# vyos_protocols_pim_rp_address (Resource)
<center>

*protocols*  
⯯  
Protocol Independent Multicast (PIM) and IGMP  
⯯  
Rendezvous Point  
⯯  
**Rendezvous Point address**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `group` (List of String) Group Address range

    |Format   &emsp;|Description                   |
    |-----------|--------------------------------|
    |ipv4net  &emsp;|Group Address range RFC 3171  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `address` (String) Rendezvous Point address

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |ipv4    &emsp;|Rendezvous Point address  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
