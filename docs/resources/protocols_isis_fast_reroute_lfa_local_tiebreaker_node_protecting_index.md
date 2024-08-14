---
page_title: "vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_node_protecting_index Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Intermediate System to Intermediate System (IS-IS)⯯IS-IS fast reroute configuration⯯Loop free alternate functionality⯯Local loop free alternate options⯯Configure tiebreaker for multiple backups⯯Prefer node protecting backup path⯯Set preference order among tiebreakers
---

# vyos_protocols_isis_fast_reroute_lfa_local_tiebreaker_node_protecting_index (Resource)
<center>

*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
IS-IS fast reroute configuration  
⯯  
Loop free alternate functionality  
⯯  
Local loop free alternate options  
⯯  
Configure tiebreaker for multiple backups  
⯯  
Prefer node protecting backup path  
⯯  
**Set preference order among tiebreakers**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `index` (Number) Set preference order among tiebreakers

    |Format  &emsp;|Description              |
    |----------|---------------------------|
    |1-255   &emsp;|The index integer value  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
