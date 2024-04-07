---
page_title: "vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_distance_prefix Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Border Gateway Protocol (BGP)⯯BGP address-family parameters⯯Multicast IPv4 BGP settings⯯Administrative distances for BGP routes⯯Administrative distance for a specific BGP prefix
---

# vyos_vrf_name_protocols_bgp_address_family_ipv4_multicast_distance_prefix (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Border Gateway Protocol (BGP)  
⯯  
BGP address-family parameters  
⯯  
Multicast IPv4 BGP settings  
⯯  
Administrative distances for BGP routes  
⯯  
**Administrative distance for a specific BGP prefix**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `prefix_id` (String) Administrative distance for a specific BGP prefix

    &emsp;|Format   &emsp;|Description                                        |
    |-----------|-----------------------------------------------------|
    &emsp;|ipv4net  &emsp;|Administrative distance for a specific BGP prefix  |

### Optional

- `distance` (Number) Administrative distance for prefix

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|1-255   &emsp;|Administrative distance for external BGP routes  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
