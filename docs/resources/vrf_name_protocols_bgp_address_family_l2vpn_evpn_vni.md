---
page_title: "vyos_vrf_name_protocols_bgp_address_family_l2vpn_evpn_vni Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Border Gateway Protocol (BGP)⯯BGP address-family parameters⯯L2VPN EVPN BGP settings⯯VXLAN Network Identifier
---

# vyos_vrf_name_protocols_bgp_address_family_l2vpn_evpn_vni (Resource)
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
L2VPN EVPN BGP settings  
⯯  
**VXLAN Network Identifier**


</center>

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |
- `vni_id` (Number) VXLAN Network Identifier

    &emsp;|Format      &emsp;|Description  |
    |--------------|---------------|
    &emsp;|1-16777215  &emsp;|VNI number   |

### Optional

- `advertise_default_gw` (Boolean) Advertise All default g/w mac-ip routes in EVPN
- `advertise_svi_ip` (Boolean) Advertise svi mac-ip routes in EVPN
- `rd` (String) Route Distinguisher

    &emsp;|Format                   &emsp;|Description                                   |
    |---------------------------|------------------------------------------------|
    &emsp;|ASN:NN_OR_IP-ADDRESS:NN  &emsp;|Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
- `route_target` (Attributes) Route Target (see [below for nested schema](#nestedatt--route_target))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--route_target&#34;&gt;&lt;/a&gt;
### Nested Schema for `route_target`

Optional:

- `both` (List of String) Route Target both import and export

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|txt     &emsp;|Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
- `export` (List of String) Route Target export

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|txt     &emsp;|Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
- `import` (List of String) Route Target import

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|txt     &emsp;|Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
