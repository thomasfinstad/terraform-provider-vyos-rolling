---
page_title: "vyos_vrf_name_protocols_bgp_address_family_ipv4_vpn_network Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
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
  Unicast VPN IPv4 BGP settings
  ⯯
  Import BGP network/prefix into unicast VPN IPv4 RIB
---

# vyos_vrf_name_protocols_bgp_address_family_ipv4_vpn_network (Resource)
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
Unicast VPN IPv4 BGP settings
⯯
**Import BGP network/prefix into unicast VPN IPv4 RIB**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `network_id` (String) Import BGP network/prefix into unicast VPN IPv4 RIB

    &emsp;|Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    &emsp;|ipv4net  &emsp;|Unicast VPN IPv4 BGP network/prefix  |

### Optional

- `label` (Number) MPLS label value assigned to route

    &emsp;|Format     &emsp;|Description       |
    |-------------|--------------------|
    &emsp;|0-1048575  &emsp;|MPLS label value  |
- `rd` (String) Route Distinguisher

    &emsp;|Format                   &emsp;|Description                                   |
    |---------------------------|------------------------------------------------|
    &emsp;|ASN:NN_OR_IP-ADDRESS:NN  &emsp;|Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
