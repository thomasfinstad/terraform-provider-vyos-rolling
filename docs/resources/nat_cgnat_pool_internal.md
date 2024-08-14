---
page_title: "vyos_nat_cgnat_pool_internal Resource - vyos"

subcategory: "Nat"

description: |- 
  nat⯯Carrier-grade NAT (CGNAT) parameters⯯External and internal pool parameters⯯Internal pool name
---

# vyos_nat_cgnat_pool_internal (Resource)
<center>

*nat*  
⯯  
Carrier-grade NAT (CGNAT) parameters  
⯯  
External and internal pool parameters  
⯯  
**Internal pool name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `range` (List of String) Range of IP addresses

    |Format     &emsp;|Description         |
    |-------------|----------------------|
    |ipv4net    &emsp;|IPv4 prefix         |
    |ipv4range  &emsp;|IPv4 address range  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `internal` (String) Internal pool name

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |txt     &emsp;|Internal pool name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
