---
page_title: "vyos_nat_cgnat_pool_external_range Resource - vyos"

subcategory: "Nat"

description: |- 
  nat⯯Carrier-grade NAT (CGNAT) parameters⯯External and internal pool parameters⯯External pool name⯯Range of IP addresses
---

# vyos_nat_cgnat_pool_external_range (Resource)
<center>

*nat*  
⯯  
Carrier-grade NAT (CGNAT) parameters  
⯯  
External and internal pool parameters  
⯯  
External pool name  
⯯  
**Range of IP addresses**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `seq` (Number) Sequence

    |Format    &emsp;|Description      |
    |------------|-------------------|
    |1-999999  &emsp;|Sequence number  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `external` (String) External pool name

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |txt     &emsp;|External pool name  |
- `range` (String) Range of IP addresses

    |Format     &emsp;|Description         |
    |-------------|----------------------|
    |ipv4net    &emsp;|IPv4 prefix         |
    |ipv4range  &emsp;|IPv4 address range  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
