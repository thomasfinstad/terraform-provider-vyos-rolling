---
page_title: "vyos_vrf_name_protocols_bgp_address_family_ipv6_vpn_network Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Border Gateway Protocol (BGP)⯯BGP address-family parameters⯯Unicast VPN IPv6 BGP settings⯯Import BGP network/prefix into unicast VPN IPv6 RIB
---

# vyos_vrf_name_protocols_bgp_address_family_ipv6_vpn_network (Resource)
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
Unicast VPN IPv6 BGP settings  
⯯  
**Import BGP network/prefix into unicast VPN IPv6 RIB**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `label` (Number) MPLS label value assigned to route

    |Format     &emsp;|Description       |
    |-------------|--------------------|
    |0-1048575  &emsp;|MPLS label value  |
- `rd` (String) Route Distinguisher

    |Format                   &emsp;|Description                                   |
    |---------------------------|------------------------------------------------|
    |ASN:NN_OR_IP-ADDRESS:NN  &emsp;|Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `name` (String) Virtual Routing and Forwarding instance

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |
- `network` (String) Import BGP network/prefix into unicast VPN IPv6 RIB

    |Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    |ipv6net  &emsp;|Unicast VPN IPv6 BGP network/prefix  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
