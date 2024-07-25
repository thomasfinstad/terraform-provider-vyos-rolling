---
page_title: "vyos_vrf_name Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance
---

# vyos_vrf_name (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
**Virtual Routing and Forwarding instance**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `disable` (Boolean) Administratively disable interface
- `ip` (Attributes) IPv4 routing parameters (see [below for nested schema](#nestedatt--ip))
- `ipv6` (Attributes) IPv6 routing parameters (see [below for nested schema](#nestedatt--ipv6))
- `protocols` (Attributes) Routing protocol parameters (see [below for nested schema](#nestedatt--protocols))
- `table` (Number) Routing table associated with this instance

    &emsp;|Format     &emsp;|Description       |
    |-------------|--------------------|
    &emsp;|100-65535  &emsp;|Routing table ID  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vni` (Number) Virtual Network Identifier

    &emsp;|Format      &emsp;|Description                       |
    |--------------|------------------------------------|
    &emsp;|0-16777214  &emsp;|VXLAN virtual network identifier  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--ip&#34;&gt;&lt;/a&gt;
### Nested Schema for `ip`

Optional:

- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `nht` (Attributes) Filter Next Hop tracking route resolution (see [below for nested schema](#nestedatt--ip--nht))

&lt;a id=&#34;nestedatt--ip--nht&#34;&gt;&lt;/a&gt;
### Nested Schema for `ip.nht`

Optional:

- `no_resolve_via_default` (Boolean) Do not resolve via default route



&lt;a id=&#34;nestedatt--ipv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6`

Optional:

- `disable_forwarding` (Boolean) Disable IP forwarding on this interface
- `nht` (Attributes) Filter Next Hop tracking route resolution (see [below for nested schema](#nestedatt--ipv6--nht))

&lt;a id=&#34;nestedatt--ipv6--nht&#34;&gt;&lt;/a&gt;
### Nested Schema for `ipv6.nht`

Optional:

- `no_resolve_via_default` (Boolean) Do not resolve via default route



&lt;a id=&#34;nestedatt--protocols&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols`

Optional:

- `bgp` (Attributes) Border Gateway Protocol (BGP) (see [below for nested schema](#nestedatt--protocols--bgp))
- `eigrp` (Attributes) Enhanced Interior Gateway Routing Protocol (EIGRP) (see [below for nested schema](#nestedatt--protocols--eigrp))
- `isis` (Attributes) Intermediate System to Intermediate System (IS-IS) (see [below for nested schema](#nestedatt--protocols--isis))
- `ospf` (Attributes) Open Shortest Path First (OSPF) (see [below for nested schema](#nestedatt--protocols--ospf))
- `ospfv3` (Attributes) Open Shortest Path First (OSPF) for IPv6 (see [below for nested schema](#nestedatt--protocols--ospfv3))
- `static` (Attributes) Static Routing (see [below for nested schema](#nestedatt--protocols--static))

&lt;a id=&#34;nestedatt--protocols--bgp&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp`

Optional:

- `address_family` (Attributes) BGP address-family parameters (see [below for nested schema](#nestedatt--protocols--bgp--address_family))
- `bmp` (Attributes) BGP Monitoring Protocol (BMP) (see [below for nested schema](#nestedatt--protocols--bgp--bmp))
- `listen` (Attributes) Listen for and accept BGP dynamic neighbors from range (see [below for nested schema](#nestedatt--protocols--bgp--listen))
- `parameters` (Attributes) BGP parameters (see [below for nested schema](#nestedatt--protocols--bgp--parameters))
- `sid` (Attributes) SID value for VRF (see [below for nested schema](#nestedatt--protocols--bgp--sid))
- `srv6` (Attributes) Segment-Routing SRv6 configuration (see [below for nested schema](#nestedatt--protocols--bgp--srv6))
- `system_as` (Number) Autonomous System Number (ASN)

    &emsp;|Format        &emsp;|Description               |
    |----------------|----------------------------|
    &emsp;|1-4294967294  &emsp;|Autonomous System Number  |
- `timers` (Attributes) BGP protocol timers (see [below for nested schema](#nestedatt--protocols--bgp--timers))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family`

Optional:

- `ipv4_flowspec` (Attributes) Flowspec IPv4 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_flowspec))
- `ipv4_labeled_unicast` (Attributes) Labeled Unicast IPv4 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_labeled_unicast))
- `ipv4_multicast` (Attributes) Multicast IPv4 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_multicast))
- `ipv4_unicast` (Attributes) IPv4 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast))
- `ipv4_vpn` (Attributes) Unicast VPN IPv4 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_vpn))
- `ipv6_flowspec` (Attributes) Flowspec IPv6 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_flowspec))
- `ipv6_labeled_unicast` (Attributes) Labeled Unicast IPv6 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_labeled_unicast))
- `ipv6_multicast` (Attributes) Multicast IPv6 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_multicast))
- `ipv6_unicast` (Attributes) IPv6 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast))
- `ipv6_vpn` (Attributes) Unicast VPN IPv6 BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_vpn))
- `l2vpn_evpn` (Attributes) L2VPN EVPN BGP settings (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_flowspec&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_flowspec`

Optional:

- `local_install` (Attributes) Apply local policy routing to interface (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_flowspec--local_install))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_flowspec--local_install&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_flowspec.local_install`

Optional:

- `interface` (List of String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_labeled_unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_labeled_unicast`

Optional:

- `maximum_paths` (Attributes) Forward packets over multiple paths (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_labeled_unicast--maximum_paths))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_labeled_unicast--maximum_paths&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_labeled_unicast.maximum_paths`

Optional:

- `ebgp` (Number) eBGP maximum paths

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|1-256   &emsp;|Number of paths to consider  |
- `ibgp` (Number) iBGP maximum paths

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|1-256   &emsp;|Number of paths to consider  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_multicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_multicast`

Optional:

- `distance` (Attributes) Administrative distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_multicast--distance))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_multicast--distance&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_multicast.distance`

Optional:

- `external` (Number) eBGP routes administrative distance

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|1-255   &emsp;|eBGP routes administrative distance  |
- `internal` (Number) iBGP routes administrative distance

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|1-255   &emsp;|iBGP routes administrative distance  |
- `local` (Number) Locally originated BGP routes administrative distance

    &emsp;|Format  &emsp;|Description                                            |
    |----------|---------------------------------------------------------|
    &emsp;|1-255   &emsp;|Locally originated BGP routes administrative distance  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast`

Optional:

- `distance` (Attributes) Administrative distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--distance))
- `export` (Attributes) Export routes from this address-family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--export))
- `import` (Attributes) Import routes to this address-family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--import))
- `label` (Attributes) Label value for VRF (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--label))
- `maximum_paths` (Attributes) Forward packets over multiple paths (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--maximum_paths))
- `nexthop` (Attributes) Specify next hop to use for VRF advertised prefixes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--nexthop))
- `rd` (Attributes) Specify route distinguisher (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--rd))
- `redistribute` (Attributes) Redistribute routes from other protocols into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--route_map))
- `route_target` (Attributes) Specify route target list (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--route_target))
- `sid` (Attributes) SID value for VRF (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--sid))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--distance&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.distance`

Optional:

- `external` (Number) eBGP routes administrative distance

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|1-255   &emsp;|eBGP routes administrative distance  |
- `internal` (Number) iBGP routes administrative distance

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|1-255   &emsp;|iBGP routes administrative distance  |
- `local` (Number) Locally originated BGP routes administrative distance

    &emsp;|Format  &emsp;|Description                                            |
    |----------|---------------------------------------------------------|
    &emsp;|1-255   &emsp;|Locally originated BGP routes administrative distance  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--export&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.export`

Optional:

- `vpn` (Boolean) to/from default instance VPN RIB


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--import&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.import`

Optional:

- `vpn` (Boolean) to/from default instance VPN RIB
- `vrf` (List of String) VRF to import from

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--label&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.label`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--label--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--label--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.label.vpn`

Optional:

- `allocation_mode` (Attributes) Label allocation mode (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--label--vpn--allocation_mode))
- `export` (String) For routes leaked from current address-family to VPN

    &emsp;|Format     &emsp;|Description                   |
    |-------------|--------------------------------|
    &emsp;|auto       &emsp;|Automatically assign a label  |
    &emsp;|0-1048575  &emsp;|Label Value                   |

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--label--vpn--allocation_mode&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.label.vpn.allocation_mode`

Optional:

- `per_nexthop` (Boolean) Allocate a label per connected next-hop in the VRF




&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--maximum_paths&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.maximum_paths`

Optional:

- `ebgp` (Number) eBGP maximum paths

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|1-256   &emsp;|Number of paths to consider  |
- `ibgp` (Number) iBGP maximum paths

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|1-256   &emsp;|Number of paths to consider  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--nexthop&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.nexthop`

Optional:

- `vpn` (Attributes) Between current address-family and vpn (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--nexthop--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--nexthop--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.nexthop.vpn`

Optional:

- `export` (String) For routes leaked from current address-family to vpn

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|ipv4    &emsp;|BGP neighbor IP address    |
    &emsp;|ipv6    &emsp;|BGP neighbor IPv6 address  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--rd&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.rd`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--rd--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--rd--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.rd.vpn`

Optional:

- `export` (String) For routes leaked from current address-family to VPN

    &emsp;|Format                   &emsp;|Description                                   |
    |---------------------------|------------------------------------------------|
    &emsp;|ASN:NN_OR_IP-ADDRESS:NN  &emsp;|Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute`

Optional:

- `babel` (Attributes) Redistribute Babel routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--babel))
- `connected` (Attributes) Redistribute connected routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--connected))
- `isis` (Attributes) Redistribute IS-IS routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--isis))
- `kernel` (Attributes) Redistribute kernel routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--kernel))
- `ospf` (Attributes) Redistribute OSPF routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--ospf))
- `rip` (Attributes) Redistribute RIP routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--rip))
- `static` (Attributes) Redistribute static routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--static))
- `table` (String) Redistribute non-main Kernel Routing Table

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--babel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.babel`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--connected&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.connected`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--isis&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.isis`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--kernel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.kernel`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--ospf&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.ospf`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--rip&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.rip`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--redistribute--static&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.redistribute.static`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_map`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--route_map--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--route_map--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_map.vpn`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--route_target&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_target`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--route_target--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--route_target--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.route_target.vpn`

Optional:

- `both` (String) Route Target both import and export

    &emsp;|Format  &emsp;|Description                                                     |
    |----------|------------------------------------------------------------------|
    &emsp;|txt     &emsp;|Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
- `export` (String) Route Target export

    &emsp;|Format  &emsp;|Description                                                     |
    |----------|------------------------------------------------------------------|
    &emsp;|txt     &emsp;|Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
- `import` (String) Route Target import

    &emsp;|Format  &emsp;|Description                                                     |
    |----------|------------------------------------------------------------------|
    &emsp;|txt     &emsp;|Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--sid&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.sid`

Optional:

- `vpn` (Attributes) Between current VRF and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv4_unicast--sid--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_unicast--sid--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_unicast.sid.vpn`

Optional:

- `export` (String) For routes leaked from current VRF to VPN

    &emsp;|Format     &emsp;|Description                   |
    |-------------|--------------------------------|
    &emsp;|1-1048575  &emsp;|SID allocation index          |
    &emsp;|auto       &emsp;|Automatically assign a label  |




&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv4_vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv4_vpn`


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_flowspec&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_flowspec`

Optional:

- `local_install` (Attributes) Apply local policy routing to interface (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_flowspec--local_install))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_flowspec--local_install&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_flowspec.local_install`

Optional:

- `interface` (List of String) Interface



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_labeled_unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_labeled_unicast`


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_multicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_multicast`

Optional:

- `distance` (Attributes) Administrative distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_multicast--distance))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_multicast--distance&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_multicast.distance`

Optional:

- `external` (Number) eBGP routes administrative distance

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|1-255   &emsp;|eBGP routes administrative distance  |
- `internal` (Number) iBGP routes administrative distance

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|1-255   &emsp;|iBGP routes administrative distance  |
- `local` (Number) Locally originated BGP routes administrative distance

    &emsp;|Format  &emsp;|Description                                            |
    |----------|---------------------------------------------------------|
    &emsp;|1-255   &emsp;|Locally originated BGP routes administrative distance  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast`

Optional:

- `distance` (Attributes) Administrative distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--distance))
- `export` (Attributes) Export routes from this address-family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--export))
- `import` (Attributes) Import routes to this address-family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--import))
- `label` (Attributes) Label value for VRF (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--label))
- `maximum_paths` (Attributes) Forward packets over multiple paths (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--maximum_paths))
- `nexthop` (Attributes) Specify next hop to use for VRF advertised prefixes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--nexthop))
- `rd` (Attributes) Specify route distinguisher (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--rd))
- `redistribute` (Attributes) Redistribute routes from other protocols into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute))
- `route_map` (Attributes) Route-map to filter route updates to/from this peer (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--route_map))
- `route_target` (Attributes) Specify route target list (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--route_target))
- `sid` (Attributes) SID value for VRF (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--sid))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--distance&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.distance`

Optional:

- `external` (Number) eBGP routes administrative distance

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|1-255   &emsp;|eBGP routes administrative distance  |
- `internal` (Number) iBGP routes administrative distance

    &emsp;|Format  &emsp;|Description                          |
    |----------|---------------------------------------|
    &emsp;|1-255   &emsp;|iBGP routes administrative distance  |
- `local` (Number) Locally originated BGP routes administrative distance

    &emsp;|Format  &emsp;|Description                                            |
    |----------|---------------------------------------------------------|
    &emsp;|1-255   &emsp;|Locally originated BGP routes administrative distance  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--export&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.export`

Optional:

- `vpn` (Boolean) to/from default instance VPN RIB


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--import&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.import`

Optional:

- `vpn` (Boolean) to/from default instance VPN RIB
- `vrf` (List of String) VRF to import from

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--label&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.label`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--label--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--label--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.label.vpn`

Optional:

- `allocation_mode` (Attributes) Label allocation mode (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--label--vpn--allocation_mode))
- `export` (String) For routes leaked from current address-family to VPN

    &emsp;|Format     &emsp;|Description                   |
    |-------------|--------------------------------|
    &emsp;|auto       &emsp;|Automatically assign a label  |
    &emsp;|0-1048575  &emsp;|Label Value                   |

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--label--vpn--allocation_mode&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.label.vpn.allocation_mode`

Optional:

- `per_nexthop` (Boolean) Allocate a label per connected next-hop in the VRF




&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--maximum_paths&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.maximum_paths`

Optional:

- `ebgp` (Number) eBGP maximum paths

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|1-256   &emsp;|Number of paths to consider  |
- `ibgp` (Number) iBGP maximum paths

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|1-256   &emsp;|Number of paths to consider  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--nexthop&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.nexthop`

Optional:

- `vpn` (Attributes) Between current address-family and vpn (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--nexthop--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--nexthop--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.nexthop.vpn`

Optional:

- `export` (String) For routes leaked from current address-family to vpn

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|ipv4    &emsp;|BGP neighbor IP address    |
    &emsp;|ipv6    &emsp;|BGP neighbor IPv6 address  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--rd&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.rd`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--rd--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--rd--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.rd.vpn`

Optional:

- `export` (String) For routes leaked from current address-family to VPN

    &emsp;|Format                   &emsp;|Description                                   |
    |---------------------------|------------------------------------------------|
    &emsp;|ASN:NN_OR_IP-ADDRESS:NN  &emsp;|Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute`

Optional:

- `babel` (Attributes) Redistribute Babel routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--babel))
- `connected` (Attributes) Redistribute connected routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--connected))
- `kernel` (Attributes) Redistribute kernel routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--kernel))
- `ospfv3` (Attributes) Redistribute OSPFv3 routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--ospfv3))
- `ripng` (Attributes) Redistribute RIPng routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--ripng))
- `static` (Attributes) Redistribute static routes into BGP (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--static))
- `table` (String) Redistribute non-main Kernel Routing Table

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--babel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.babel`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--connected&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.connected`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--kernel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.kernel`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--ospfv3&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.ospfv3`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--ripng&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.ripng`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--redistribute--static&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.redistribute.static`

Optional:

- `metric` (Number) Metric for redistributed routes

    &emsp;|Format        &emsp;|Description                      |
    |----------------|-----------------------------------|
    &emsp;|1-4294967295  &emsp;|Metric for redistributed routes  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--route_map&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_map`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--route_map--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--route_map--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_map.vpn`

Optional:

- `export` (String) Route-map to filter outgoing route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `import` (String) Route-map to filter incoming route updates

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--route_target&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_target`

Optional:

- `vpn` (Attributes) Between current address-family and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--route_target--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--route_target--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.route_target.vpn`

Optional:

- `both` (String) Route Target both import and export

    &emsp;|Format  &emsp;|Description                                                     |
    |----------|------------------------------------------------------------------|
    &emsp;|txt     &emsp;|Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
- `export` (String) Route Target export

    &emsp;|Format  &emsp;|Description                                                     |
    |----------|------------------------------------------------------------------|
    &emsp;|txt     &emsp;|Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |
- `import` (String) Route Target import

    &emsp;|Format  &emsp;|Description                                                     |
    |----------|------------------------------------------------------------------|
    &emsp;|txt     &emsp;|Space separated route target list (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--sid&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.sid`

Optional:

- `vpn` (Attributes) Between current VRF and VPN (see [below for nested schema](#nestedatt--protocols--bgp--address_family--ipv6_unicast--sid--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_unicast--sid--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_unicast.sid.vpn`

Optional:

- `export` (String) For routes leaked from current VRF to VPN

    &emsp;|Format     &emsp;|Description                   |
    |-------------|--------------------------------|
    &emsp;|1-1048575  &emsp;|SID allocation index          |
    &emsp;|auto       &emsp;|Automatically assign a label  |




&lt;a id=&#34;nestedatt--protocols--bgp--address_family--ipv6_vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.ipv6_vpn`


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn`

Optional:

- `advertise` (Attributes) Advertise prefix routes (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise))
- `advertise_all_vni` (Boolean) Advertise All local VNIs
- `advertise_default_gw` (Boolean) Advertise All default g/w mac-ip routes in EVPN
- `advertise_pip` (String) EVPN system primary IP

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|ipv4    &emsp;|IP address   |
- `advertise_svi_ip` (Boolean) Advertise svi mac-ip routes in EVPN
- `default_originate` (Attributes) Originate a default route (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--default_originate))
- `disable_ead_evi_rx` (Boolean) Activate PE on EAD-ES even if EAD-EVI is not received
- `disable_ead_evi_tx` (Boolean) Do not advertise EAD-EVI for local ESs
- `ead_es_frag` (Attributes) EAD ES fragment config (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--ead_es_frag))
- `ead_es_route_target` (Attributes) EAD ES Route Target (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--ead_es_route_target))
- `flooding` (Attributes) Specify handling for BUM packets (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--flooding))
- `mac_vrf` (Attributes) EVPN MAC-VRF (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--mac_vrf))
- `rd` (String) Route Distinguisher

    &emsp;|Format                   &emsp;|Description                                   |
    |---------------------------|------------------------------------------------|
    &emsp;|ASN:NN_OR_IP-ADDRESS:NN  &emsp;|Route Distinguisher, (x.x.x.x:yyy|xxxx:yyyy)  |
- `route_target` (Attributes) Route Target (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--route_target))
- `rt_auto_derive` (Boolean) Auto derivation of Route Target (RFC8365)

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise`

Optional:

- `ipv4` (Attributes) IPv4 address family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv4))
- `ipv6` (Attributes) IPv6 address family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv6))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv4&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv4`

Optional:

- `unicast` (Attributes) IPv4 address family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv4--unicast))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv4--unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv4.unicast`

Optional:

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv6`

Optional:

- `unicast` (Attributes) IPv4 address family (see [below for nested schema](#nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv6--unicast))

&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--advertise--ipv6--unicast&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.advertise.ipv6.unicast`

Optional:

- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |




&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--default_originate&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.default_originate`

Optional:

- `ipv4` (Boolean) IPv4 address family
- `ipv6` (Boolean) IPv6 address family


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--ead_es_frag&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.ead_es_frag`

Optional:

- `evi_limit` (Number) EVIs per-fragment

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1-1000  &emsp;|limit        |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--ead_es_route_target&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.ead_es_route_target`

Optional:

- `export` (List of String) Route Target export

    &emsp;|Format  &emsp;|Description                                |
    |----------|---------------------------------------------|
    &emsp;|txt     &emsp;|Route target (A.B.C.D:MN|EF:OPQR|GHJK:MN)  |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--flooding&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.flooding`

Optional:

- `disable` (Boolean) Disable instance
- `head_end_replication` (Boolean) Flood BUM packets using head-end replication


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--mac_vrf&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.mac_vrf`

Optional:

- `soo` (String) Site-of-Origin extended community

    &emsp;|Format  &emsp;|Description                                                         |
    |----------|----------------------------------------------------------------------|
    &emsp;|ASN:NN  &emsp;|based on autonomous system number in format &lt;0-65535:0-4294967295&gt;  |
    &emsp;|IP:NN   &emsp;|Based on a router-id IP address in format &lt;IP:0-65535&gt;              |


&lt;a id=&#34;nestedatt--protocols--bgp--address_family--l2vpn_evpn--route_target&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.address_family.l2vpn_evpn.route_target`

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




&lt;a id=&#34;nestedatt--protocols--bgp--bmp&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.bmp`

Optional:

- `mirror_buffer_limit` (Number) Maximum memory used for buffered mirroring messages (in bytes)

    &emsp;|Format        &emsp;|Description     |
    |----------------|------------------|
    &emsp;|0-4294967294  &emsp;|Limit in bytes  |


&lt;a id=&#34;nestedatt--protocols--bgp--listen&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.listen`

Optional:

- `limit` (Number) Maximum number of dynamic neighbors that can be created

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|1-5000  &emsp;|BGP neighbor limit  |


&lt;a id=&#34;nestedatt--protocols--bgp--parameters&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters`

Optional:

- `allow_martian_nexthop` (Boolean) Allow Martian nexthops to be received in the NLRI from a peer
- `always_compare_med` (Boolean) Always compare MEDs from different neighbors
- `bestpath` (Attributes) Default bestpath selection mechanism (see [below for nested schema](#nestedatt--protocols--bgp--parameters--bestpath))
- `cluster_id` (String) Route-reflector cluster-id

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|ipv4    &emsp;|Route-reflector cluster-id  |
- `conditional_advertisement` (Attributes) Conditional advertisement settings (see [below for nested schema](#nestedatt--protocols--bgp--parameters--conditional_advertisement))
- `confederation` (Attributes) AS confederation parameters (see [below for nested schema](#nestedatt--protocols--bgp--parameters--confederation))
- `dampening` (Attributes) Enable route-flap dampening (see [below for nested schema](#nestedatt--protocols--bgp--parameters--dampening))
- `default` (Attributes) BGP defaults (see [below for nested schema](#nestedatt--protocols--bgp--parameters--default))
- `deterministic_med` (Boolean) Compare MEDs between different peers in the same AS
- `disable_ebgp_connected_route_check` (Boolean) Disable checking if nexthop is connected on eBGP session
- `distance` (Attributes) Administratives distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--parameters--distance))
- `ebgp_requires_policy` (Boolean) Require in and out policy for eBGP peers (RFC8212)
- `fast_convergence` (Boolean) Teardown sessions immediately whenever peer becomes unreachable
- `graceful_restart` (Attributes) Graceful restart capability parameters (see [below for nested schema](#nestedatt--protocols--bgp--parameters--graceful_restart))
- `graceful_shutdown` (Boolean) Graceful shutdown
- `labeled_unicast` (String) BGP Labeled-unicast options

    &emsp;|Format              &emsp;|Description                                                 |
    |----------------------|--------------------------------------------------------------|
    &emsp;|explicit-null       &emsp;|Use explicit-null label values for all local prefixes       |
    &emsp;|ipv4-explicit-null  &emsp;|Use IPv4 explicit-null label value for IPv4 local prefixes  |
    &emsp;|ipv6-explicit-null  &emsp;|Use IPv6 explicit-null label value for IPv4 local prefixes  |
- `log_neighbor_changes` (Boolean) Log neighbor up/down changes and reset reason
- `minimum_holdtime` (Number) BGP minimum holdtime

    &emsp;|Format   &emsp;|Description                  |
    |-----------|-------------------------------|
    &emsp;|1-65535  &emsp;|Minimum holdtime in seconds  |
- `network_import_check` (Boolean) Enable IGP route check for network statements
- `no_client_to_client_reflection` (Boolean) Disable client to client route reflection
- `no_fast_external_failover` (Boolean) Disable immediate session reset on peer link down event
- `no_hard_administrative_reset` (Boolean) Do not send hard reset CEASE Notification for &#39;Administrative Reset&#39;
- `no_suppress_duplicates` (Boolean) Disable suppress duplicate updates if the route actually not changed
- `reject_as_sets` (Boolean) Reject routes with AS_SET or AS_CONFED_SET flag
- `route_reflector_allow_outbound_policy` (Boolean) Route reflector client allow policy outbound
- `router_id` (String) Override default router identifier

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv4    &emsp;|Router-ID in IP address format  |
- `shutdown` (Boolean) Administrative shutdown of the BGP instance
- `suppress_fib_pending` (Boolean) Advertise only routes that are programmed in kernel to peers
- `tcp_keepalive` (Attributes) TCP keepalive parameters (see [below for nested schema](#nestedatt--protocols--bgp--parameters--tcp_keepalive))

&lt;a id=&#34;nestedatt--protocols--bgp--parameters--bestpath&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.bestpath`

Optional:

- `as_path` (Attributes) AS-path attribute comparison parameters (see [below for nested schema](#nestedatt--protocols--bgp--parameters--bestpath--as_path))
- `bandwidth` (String) Link Bandwidth attribute

    &emsp;|Format                      &emsp;|Description                                                            |
    |------------------------------|-------------------------------------------------------------------------|
    &emsp;|default-weight-for-missing  &emsp;|Assign low default weight (1) to paths not having link bandwidth       |
    &emsp;|ignore                      &emsp;|Ignore link bandwidth (do regular ECMP, not weighted)                  |
    &emsp;|skip-missing                &emsp;|Ignore paths without link bandwidth for ECMP (if other paths have it)  |
- `compare_routerid` (Boolean) Compare the router-id for identical EBGP paths
- `med` (List of String) MED attribute comparison parameters

    &emsp;|Format            &emsp;|Description                                              |
    |--------------------|-----------------------------------------------------------|
    &emsp;|confed            &emsp;|Compare MEDs among confederation paths                   |
    &emsp;|missing-as-worst  &emsp;|Treat missing route as a MED as the least preferred one  |
- `peer_type` (Attributes) Peer type (see [below for nested schema](#nestedatt--protocols--bgp--parameters--bestpath--peer_type))

&lt;a id=&#34;nestedatt--protocols--bgp--parameters--bestpath--as_path&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.bestpath.as_path`

Optional:

- `confed` (Boolean) Compare AS-path lengths including confederation sets and sequences
- `ignore` (Boolean) Ignore AS-path length in selecting a route
- `multipath_relax` (Boolean) Allow load sharing across routes that have different AS paths (but same length)


&lt;a id=&#34;nestedatt--protocols--bgp--parameters--bestpath--peer_type&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.bestpath.peer_type`

Optional:

- `multipath_relax` (Boolean) Allow load sharing across routes learned from different peer types



&lt;a id=&#34;nestedatt--protocols--bgp--parameters--conditional_advertisement&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.conditional_advertisement`

Optional:

- `timer` (Number) Set period to rescan BGP table to check if condition is met

    &emsp;|Format  &emsp;|Description                                                    |
    |----------|-----------------------------------------------------------------|
    &emsp;|5-240   &emsp;|Period to rerun the conditional advertisement scanner process  |


&lt;a id=&#34;nestedatt--protocols--bgp--parameters--confederation&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.confederation`

Optional:

- `identifier` (Number) Confederation AS identifier

    &emsp;|Format        &emsp;|Description          |
    |----------------|-----------------------|
    &emsp;|1-4294967294  &emsp;|Confederation AS id  |
- `peers` (List of Number) Peer ASs in the BGP confederation

    &emsp;|Format        &emsp;|Description     |
    |----------------|------------------|
    &emsp;|1-4294967294  &emsp;|Peer AS number  |


&lt;a id=&#34;nestedatt--protocols--bgp--parameters--dampening&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.dampening`

Optional:

- `half_life` (Number) Half-life time for dampening

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|1-45    &emsp;|Half-life penalty in minutes  |
- `max_suppress_time` (Number) Maximum duration to suppress a stable route

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|1-255   &emsp;|Maximum suppress duration in minutes  |
- `re_use` (Number) Threshold to start reusing a route

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|1-20000  &emsp;|Re-use penalty points  |
- `start_suppress_time` (Number) When to start suppressing a route

    &emsp;|Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    &emsp;|1-20000  &emsp;|Start-suppress penalty points  |


&lt;a id=&#34;nestedatt--protocols--bgp--parameters--default&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.default`

Optional:

- `local_pref` (Number) Default local preference

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|u32     &emsp;|Local preference  |


&lt;a id=&#34;nestedatt--protocols--bgp--parameters--distance&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.distance`

Optional:

- `global` (Attributes) Global administratives distances for BGP routes (see [below for nested schema](#nestedatt--protocols--bgp--parameters--distance--global))

&lt;a id=&#34;nestedatt--protocols--bgp--parameters--distance--global&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.distance.global`

Optional:

- `external` (Number) Administrative distance for external BGP routes

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|1-255   &emsp;|Administrative distance for external BGP routes  |
- `internal` (Number) Administrative distance for internal BGP routes

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|1-255   &emsp;|Administrative distance for internal BGP routes  |
- `local` (Number) Administrative distance for local BGP routes

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|1-255   &emsp;|Administrative distance for internal BGP routes  |



&lt;a id=&#34;nestedatt--protocols--bgp--parameters--graceful_restart&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.graceful_restart`

Optional:

- `stalepath_time` (Number) Maximum time to hold onto restarting neighbors stale paths

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|1-3600  &emsp;|Hold time in seconds  |


&lt;a id=&#34;nestedatt--protocols--bgp--parameters--tcp_keepalive&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.parameters.tcp_keepalive`

Optional:

- `idle` (Number) TCP keepalive idle time

    &emsp;|Format   &emsp;|Description           |
    |-----------|------------------------|
    &emsp;|1-65535  &emsp;|Idle time in seconds  |
- `interval` (Number) TCP keepalive interval

    &emsp;|Format   &emsp;|Description          |
    |-----------|-----------------------|
    &emsp;|1-65535  &emsp;|Interval in seconds  |
- `probes` (Number) TCP keepalive maximum probes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|1-30    &emsp;|Maximum probes  |



&lt;a id=&#34;nestedatt--protocols--bgp--sid&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.sid`

Optional:

- `vpn` (Attributes) Between current VRF and VPN (see [below for nested schema](#nestedatt--protocols--bgp--sid--vpn))

&lt;a id=&#34;nestedatt--protocols--bgp--sid--vpn&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.sid.vpn`

Optional:

- `per_vrf` (Attributes) SID per-VRF (both IPv4 and IPv6 address families) (see [below for nested schema](#nestedatt--protocols--bgp--sid--vpn--per_vrf))

&lt;a id=&#34;nestedatt--protocols--bgp--sid--vpn--per_vrf&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.sid.vpn.per_vrf`

Optional:

- `export` (String) For routes leaked from current VRF to VPN

    &emsp;|Format     &emsp;|Description                   |
    |-------------|--------------------------------|
    &emsp;|1-1048575  &emsp;|SID allocation index          |
    &emsp;|auto       &emsp;|Automatically assign a label  |




&lt;a id=&#34;nestedatt--protocols--bgp--srv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.srv6`

Optional:

- `locator` (String) Specify SRv6 locator

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|SRv6 locator name  |


&lt;a id=&#34;nestedatt--protocols--bgp--timers&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.bgp.timers`

Optional:

- `holdtime` (String) Hold timer

    &emsp;|Format   &emsp;|Description            |
    |-----------|-------------------------|
    &emsp;|1-65535  &emsp;|Hold timer in seconds  |
    &emsp;|0        &emsp;|Disable hold timer     |
- `keepalive` (Number) BGP keepalive interval for this neighbor

    &emsp;|Format   &emsp;|Description                    |
    |-----------|---------------------------------|
    &emsp;|1-65535  &emsp;|Keepalive interval in seconds  |



&lt;a id=&#34;nestedatt--protocols--eigrp&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.eigrp`

Optional:

- `maximum_paths` (Number) Forward packets over multiple paths

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|1-32    &emsp;|Number of paths  |
- `metric` (Attributes) Modify metrics and parameters for advertisement (see [below for nested schema](#nestedatt--protocols--eigrp--metric))
- `network` (List of String) Enable routing on an IP network

    &emsp;|Format   &emsp;|Description           |
    |-----------|------------------------|
    &emsp;|ipv4net  &emsp;|EIGRP network prefix  |
- `passive_interface` (List of String) Suppress routing updates on an interface
- `redistribute` (List of String) Redistribute information from another routing protocol

    &emsp;|Format     &emsp;|Description                          |
    |-------------|---------------------------------------|
    &emsp;|bgp        &emsp;|Border Gateway Protocol (BGP)        |
    &emsp;|connected  &emsp;|Connected routes                     |
    &emsp;|nhrp       &emsp;|Next Hop Resolution Protocol (NHRP)  |
    &emsp;|ospf       &emsp;|Open Shortest Path First (OSPFv2)    |
    &emsp;|rip        &emsp;|Routing Information Protocol (RIP)   |
    &emsp;|babel      &emsp;|Babel routing protocol (Babel)       |
    &emsp;|static     &emsp;|Statically configured routes         |
    &emsp;|vnc        &emsp;|Virtual Network Control (VNC)        |
- `router_id` (String) Override default router identifier

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv4    &emsp;|Router-ID in IP address format  |
- `system_as` (Number) Autonomous System Number (ASN)

    &emsp;|Format   &emsp;|Description               |
    |-----------|----------------------------|
    &emsp;|1-65535  &emsp;|Autonomous System Number  |
- `variance` (Number) Control load balancing variance

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|1-128   &emsp;|Metric variance multiplier  |

&lt;a id=&#34;nestedatt--protocols--eigrp--metric&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.eigrp.metric`

Optional:

- `weights` (Number) Modify metric coefficients

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|0-255   &emsp;|K1           |



&lt;a id=&#34;nestedatt--protocols--isis&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis`

Optional:

- `advertise_high_metrics` (Boolean) Advertise high metric value on all interfaces
- `advertise_passive_only` (Boolean) Advertise prefixes of passive interfaces only
- `area_password` (Attributes) Configure the authentication password for an area (see [below for nested schema](#nestedatt--protocols--isis--area_password))
- `default_information` (Attributes) Control distribution of default information (see [below for nested schema](#nestedatt--protocols--isis--default_information))
- `domain_password` (Attributes) Set the authentication password for a routing domain (see [below for nested schema](#nestedatt--protocols--isis--domain_password))
- `dynamic_hostname` (Boolean) Dynamic hostname for IS-IS
- `fast_reroute` (Attributes) IS-IS fast reroute configuration (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute))
- `ldp_sync` (Attributes) Protocol wide LDP-IGP synchronization configuration (see [below for nested schema](#nestedatt--protocols--isis--ldp_sync))
- `level` (String) IS-IS level number

    &emsp;|Format     &emsp;|Description                               |
    |-------------|--------------------------------------------|
    &emsp;|level-1    &emsp;|Act as a station router                   |
    &emsp;|level-1-2  &emsp;|Act as both a station and an area router  |
    &emsp;|level-2    &emsp;|Act as an area router                     |
- `log_adjacency_changes` (Boolean) Log adjacency state changes
- `lsp_gen_interval` (Number) Minimum interval between regenerating same LSP

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|1-120   &emsp;|Minimum interval in seconds  |
- `lsp_mtu` (Number) Configure the maximum size of generated LSPs

    &emsp;|Format    &emsp;|Description                     |
    |------------|----------------------------------|
    &emsp;|128-4352  &emsp;|Maximum size of generated LSPs  |
- `lsp_refresh_interval` (Number) LSP refresh interval

    &emsp;|Format   &emsp;|Description                      |
    |-----------|-----------------------------------|
    &emsp;|1-65235  &emsp;|LSP refresh interval in seconds  |
- `max_lsp_lifetime` (Number) Maximum LSP lifetime

    &emsp;|Format     &emsp;|Description              |
    |-------------|---------------------------|
    &emsp;|350-65535  &emsp;|LSP lifetime in seconds  |
- `metric_style` (String) Use old-style (ISO 10589) or new-style packet formats

    &emsp;|Format      &emsp;|Description                                            |
    |--------------|---------------------------------------------------------|
    &emsp;|narrow      &emsp;|Use old style of TLVs with narrow metric               |
    &emsp;|transition  &emsp;|Send and accept both styles of TLVs during transition  |
    &emsp;|wide        &emsp;|Use new style of TLVs to carry wider metric            |
- `net` (String) A Network Entity Title for this process (ISO only)

    &emsp;|Format                &emsp;|Description                 |
    |------------------------|------------------------------|
    &emsp;|XX.XXXX. ... .XXX.XX  &emsp;|Network entity title (NET)  |
- `purge_originator` (Boolean) Use the RFC 6232 purge-originator
- `redistribute` (Attributes) Redistribute information from another routing protocol (see [below for nested schema](#nestedatt--protocols--isis--redistribute))
- `segment_routing` (Attributes) Segment-Routing (SPRING) settings (see [below for nested schema](#nestedatt--protocols--isis--segment_routing))
- `set_attached_bit` (Boolean) Set attached bit to identify as L1/L2 router for inter-area traffic
- `set_overload_bit` (Boolean) Set overload bit to avoid any transit traffic
- `spf_delay_ietf` (Attributes) IETF SPF delay algorithm (see [below for nested schema](#nestedatt--protocols--isis--spf_delay_ietf))
- `spf_interval` (Number) Minimum interval between SPF calculations

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|1-120   &emsp;|Interval in seconds  |
- `topology` (String) Configure IS-IS topologies

    &emsp;|Format          &emsp;|Description                   |
    |------------------|--------------------------------|
    &emsp;|ipv4-multicast  &emsp;|Use IPv4 multicast topology   |
    &emsp;|ipv4-mgmt       &emsp;|Use IPv4 management topology  |
    &emsp;|ipv6-unicast    &emsp;|Use IPv6 unicast topology     |
    &emsp;|ipv6-multicast  &emsp;|Use IPv6 multicast topology   |
    &emsp;|ipv6-mgmt       &emsp;|Use IPv6 management topology  |
    &emsp;|ipv6-dstsrc     &emsp;|Use IPv6 dst-src topology     |
- `traffic_engineering` (Attributes) IS-IS traffic engineering extensions (see [below for nested schema](#nestedatt--protocols--isis--traffic_engineering))

&lt;a id=&#34;nestedatt--protocols--isis--area_password&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.area_password`

Optional:

- `md5` (String) MD5 authentication type

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|txt     &emsp;|Level-wide password  |
- `plaintext_password` (String) Plain-text authentication type

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|Circuit password  |


&lt;a id=&#34;nestedatt--protocols--isis--default_information&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.default_information`

Optional:

- `originate` (Attributes) Distribute a default route (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate))

&lt;a id=&#34;nestedatt--protocols--isis--default_information--originate&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.default_information.originate`

Optional:

- `ipv4` (Attributes) Distribute default route for IPv4 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv4))
- `ipv6` (Attributes) Distribute default route for IPv6 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv6))

&lt;a id=&#34;nestedatt--protocols--isis--default_information--originate--ipv4&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.default_information.originate.ipv4`

Optional:

- `level_1` (Attributes) Distribute default route into level-1 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv4--level_1))
- `level_2` (Attributes) Distribute default route into level-2 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv4--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--default_information--originate--ipv4--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.default_information.originate.ipv4.level_1`

Optional:

- `always` (Boolean) Always advertise default route
- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--default_information--originate--ipv4--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.default_information.originate.ipv4.level_2`

Optional:

- `always` (Boolean) Always advertise default route
- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--default_information--originate--ipv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.default_information.originate.ipv6`

Optional:

- `level_1` (Attributes) Distribute default route into level-1 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv6--level_1))
- `level_2` (Attributes) Distribute default route into level-2 (see [below for nested schema](#nestedatt--protocols--isis--default_information--originate--ipv6--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--default_information--originate--ipv6--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.default_information.originate.ipv6.level_1`

Optional:

- `always` (Boolean) Always advertise default route
- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--default_information--originate--ipv6--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.default_information.originate.ipv6.level_2`

Optional:

- `always` (Boolean) Always advertise default route
- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |





&lt;a id=&#34;nestedatt--protocols--isis--domain_password&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.domain_password`

Optional:

- `md5` (String) MD5 authentication type

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|txt     &emsp;|Level-wide password  |
- `plaintext_password` (String) Plain-text authentication type

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|txt     &emsp;|Circuit password  |


&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute`

Optional:

- `lfa` (Attributes) Loop free alternate functionality (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa))

&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa`

Optional:

- `local` (Attributes) Local loop free alternate options (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local))
- `remote` (Attributes) Remote loop free alternate options (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--remote))

&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local`

Optional:

- `load_sharing` (Attributes) Load share prefixes across multiple backups (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--load_sharing))
- `priority_limit` (Attributes) Limit backup computation up to the prefix priority (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit))
- `tiebreaker` (Attributes) Configure tiebreaker for multiple backups (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--tiebreaker))

&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local--load_sharing&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.load_sharing`

Optional:

- `disable` (Attributes) Disable load sharing (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--load_sharing--disable))

&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local--load_sharing--disable&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.load_sharing.disable`

Optional:

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes



&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit`

Optional:

- `critical` (Attributes) Compute for critical priority prefixes only (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--critical))
- `high` (Attributes) Compute for critical, and high priority prefixes (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--high))
- `medium` (Attributes) Compute for critical, high, and medium priority prefixes (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--medium))

&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--critical&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit.critical`

Optional:

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes


&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--high&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit.high`

Optional:

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes


&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local--priority_limit--medium&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.priority_limit.medium`

Optional:

- `level_1` (Boolean) Match on IS-IS level-1 routes
- `level_2` (Boolean) Match on IS-IS level-2 routes



&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local--tiebreaker&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.tiebreaker`

Optional:

- `downstream` (Attributes) Prefer backup path via downstream node (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--tiebreaker--downstream))
- `lowest_backup_metric` (Attributes) Prefer backup path with lowest total metric (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--tiebreaker--lowest_backup_metric))
- `node_protecting` (Attributes) Prefer node protecting backup path (see [below for nested schema](#nestedatt--protocols--isis--fast_reroute--lfa--local--tiebreaker--node_protecting))

&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local--tiebreaker--downstream&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.tiebreaker.downstream`


&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local--tiebreaker--lowest_backup_metric&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.tiebreaker.lowest_backup_metric`


&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--local--tiebreaker--node_protecting&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.local.tiebreaker.node_protecting`




&lt;a id=&#34;nestedatt--protocols--isis--fast_reroute--lfa--remote&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.fast_reroute.lfa.remote`




&lt;a id=&#34;nestedatt--protocols--isis--ldp_sync&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.ldp_sync`

Optional:

- `holddown` (Number) Hold down timer for LDP-IGP cost restoration

    &emsp;|Format   &emsp;|Description                                                                                   |
    |-----------|------------------------------------------------------------------------------------------------|
    &emsp;|0-10000  &emsp;|Time to wait in seconds for LDP-IGP synchronization to occur before restoring interface cost  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute`

Optional:

- `ipv4` (Attributes) Redistribute IPv4 routes (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4))
- `ipv6` (Attributes) Redistribute IPv6 routes (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4`

Optional:

- `babel` (Attributes) Redistribute Babel routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--babel))
- `bgp` (Attributes) Border Gateway Protocol (BGP) (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--bgp))
- `connected` (Attributes) Redistribute connected routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--connected))
- `kernel` (Attributes) Redistribute kernel routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--kernel))
- `ospf` (Attributes) Redistribute OSPF routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--ospf))
- `rip` (Attributes) Redistribute RIP routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--rip))
- `static` (Attributes) Redistribute static routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--static))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--babel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.babel`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--babel--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--babel--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--babel--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.babel.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--babel--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.babel.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--bgp&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.bgp`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--bgp--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--bgp--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--bgp--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.bgp.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--bgp--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.bgp.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--connected&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.connected`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--connected--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--connected--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--connected--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.connected.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--connected--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.connected.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--kernel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.kernel`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--kernel--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--kernel--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--kernel--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.kernel.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--kernel--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.kernel.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--ospf&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.ospf`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--ospf--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--ospf--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--ospf--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.ospf.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--ospf--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.ospf.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--rip&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.rip`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--rip--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--rip--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--rip--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.rip.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--rip--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.rip.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--static&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.static`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--static--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv4--static--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--static--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.static.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv4--static--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv4.static.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |




&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6`

Optional:

- `babel` (Attributes) Redistribute Babel routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--babel))
- `bgp` (Attributes) Redistribute BGP routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--bgp))
- `connected` (Attributes) Redistribute connected routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--connected))
- `kernel` (Attributes) Redistribute kernel routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--kernel))
- `ospf6` (Attributes) Redistribute OSPFv3 routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ospf6))
- `ripng` (Attributes) Redistribute RIPng routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ripng))
- `static` (Attributes) Redistribute static routes into IS-IS (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--static))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--babel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.babel`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--babel--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--babel--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--babel--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.babel.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--babel--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.babel.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--bgp&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.bgp`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--bgp--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--bgp--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--bgp--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.bgp.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--bgp--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.bgp.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--connected&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.connected`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--connected--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--connected--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--connected--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.connected.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--connected--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.connected.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--kernel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.kernel`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--kernel--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--kernel--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--kernel--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.kernel.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--kernel--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.kernel.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--ospf6&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.ospf6`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ospf6--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ospf6--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--ospf6--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.ospf6.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--ospf6--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.ospf6.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--ripng&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.ripng`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ripng--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--ripng--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--ripng--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.ripng.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--ripng--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.ripng.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--static&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.static`

Optional:

- `level_1` (Attributes) Redistribute into level-1 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--static--level_1))
- `level_2` (Attributes) Redistribute into level-2 (see [below for nested schema](#nestedatt--protocols--isis--redistribute--ipv6--static--level_2))

&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--static--level_1&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.static.level_1`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--isis--redistribute--ipv6--static--level_2&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.redistribute.ipv6.static.level_2`

Optional:

- `metric` (Number) Set default metric for circuit

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Default metric value  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |





&lt;a id=&#34;nestedatt--protocols--isis--segment_routing&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.segment_routing`

Optional:

- `global_block` (Attributes) Segment Routing Global Block label range (see [below for nested schema](#nestedatt--protocols--isis--segment_routing--global_block))
- `local_block` (Attributes) Segment Routing Local Block label range (see [below for nested schema](#nestedatt--protocols--isis--segment_routing--local_block))
- `maximum_label_depth` (Number) Maximum MPLS labels allowed for this router

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|1-16    &emsp;|MPLS label depth  |

&lt;a id=&#34;nestedatt--protocols--isis--segment_routing--global_block&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.segment_routing.global_block`

Optional:

- `high_label_value` (Number) MPLS label upper bound

    &emsp;|Format      &emsp;|Description  |
    |--------------|---------------|
    &emsp;|16-1048575  &emsp;|Label value  |
- `low_label_value` (Number) MPLS label lower bound

    &emsp;|Format      &emsp;|Description                                   |
    |--------------|------------------------------------------------|
    &emsp;|16-1048575  &emsp;|Label value (recommended minimum value: 300)  |


&lt;a id=&#34;nestedatt--protocols--isis--segment_routing--local_block&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.segment_routing.local_block`

Optional:

- `high_label_value` (Number) MPLS label upper bound

    &emsp;|Format      &emsp;|Description  |
    |--------------|---------------|
    &emsp;|16-1048575  &emsp;|Label value  |
- `low_label_value` (Number) MPLS label lower bound

    &emsp;|Format      &emsp;|Description                                   |
    |--------------|------------------------------------------------|
    &emsp;|16-1048575  &emsp;|Label value (recommended minimum value: 300)  |



&lt;a id=&#34;nestedatt--protocols--isis--spf_delay_ietf&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.spf_delay_ietf`

Optional:

- `holddown` (Number) Time with no received IGP events before considering IGP stable

    &emsp;|Format   &emsp;|Description                                                           |
    |-----------|------------------------------------------------------------------------|
    &emsp;|0-60000  &emsp;|Time with no received IGP events before considering IGP stable in ms  |
- `init_delay` (Number) Delay used while in QUIET state

    &emsp;|Format   &emsp;|Description                              |
    |-----------|-------------------------------------------|
    &emsp;|0-60000  &emsp;|Delay used while in QUIET state (in ms)  |
- `long_delay` (Number) Delay used while in LONG_WAIT

    &emsp;|Format   &emsp;|Description                                |
    |-----------|---------------------------------------------|
    &emsp;|0-60000  &emsp;|Delay used while in LONG_WAIT state in ms  |
- `short_delay` (Number) Delay used while in SHORT_WAIT state

    &emsp;|Format   &emsp;|Description                                   |
    |-----------|------------------------------------------------|
    &emsp;|0-60000  &emsp;|Delay used while in SHORT_WAIT state (in ms)  |
- `time_to_learn` (Number) Maximum duration needed to learn all the events related to a single failure

    &emsp;|Format   &emsp;|Description                                                                        |
    |-----------|-------------------------------------------------------------------------------------|
    &emsp;|0-60000  &emsp;|Maximum duration needed to learn all the events related to a single failure in ms  |


&lt;a id=&#34;nestedatt--protocols--isis--traffic_engineering&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.isis.traffic_engineering`

Optional:

- `address` (String) MPLS traffic engineering router ID

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|ipv4    &emsp;|IPv4 address  |
- `enable` (Boolean) Enable MPLS traffic engineering extensions



&lt;a id=&#34;nestedatt--protocols--ospf&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf`

Optional:

- `aggregation` (Attributes) External route aggregation (see [below for nested schema](#nestedatt--protocols--ospf--aggregation))
- `auto_cost` (Attributes) Calculate interface cost according to bandwidth (see [below for nested schema](#nestedatt--protocols--ospf--auto_cost))
- `capability` (Attributes) Enable specific OSPF features (see [below for nested schema](#nestedatt--protocols--ospf--capability))
- `default_information` (Attributes) Default route advertisment settings (see [below for nested schema](#nestedatt--protocols--ospf--default_information))
- `default_metric` (Number) Metric of redistributed routes

    &emsp;|Format      &emsp;|Description                     |
    |--------------|----------------------------------|
    &emsp;|0-16777214  &emsp;|Metric of redistributed routes  |
- `distance` (Attributes) Administrative distance (see [below for nested schema](#nestedatt--protocols--ospf--distance))
- `graceful_restart` (Attributes) Graceful Restart (see [below for nested schema](#nestedatt--protocols--ospf--graceful_restart))
- `ldp_sync` (Attributes) Protocol wide LDP-IGP synchronization configuration (see [below for nested schema](#nestedatt--protocols--ospf--ldp_sync))
- `log_adjacency_changes` (Attributes) Log adjacency state changes (see [below for nested schema](#nestedatt--protocols--ospf--log_adjacency_changes))
- `max_metric` (Attributes) OSPF maximum and infinite-distance metric (see [below for nested schema](#nestedatt--protocols--ospf--max_metric))
- `maximum_paths` (Number) Maximum multiple paths (ECMP)

    &emsp;|Format  &emsp;|Description                    |
    |----------|---------------------------------|
    &emsp;|1-64    &emsp;|Maximum multiple paths (ECMP)  |
- `mpls_te` (Attributes) MultiProtocol Label Switching-Traffic Engineering (MPLS-TE) parameters (see [below for nested schema](#nestedatt--protocols--ospf--mpls_te))
- `parameters` (Attributes) OSPF specific parameters (see [below for nested schema](#nestedatt--protocols--ospf--parameters))
- `passive_interface` (String) Suppress routing updates on an interface

    &emsp;|Format   &emsp;|Description                                            |
    |-----------|---------------------------------------------------------|
    &emsp;|default  &emsp;|Default to suppress routing updates on all interfaces  |
- `redistribute` (Attributes) Redistribute information from another routing protocol (see [below for nested schema](#nestedatt--protocols--ospf--redistribute))
- `refresh` (Attributes) Adjust refresh parameters (see [below for nested schema](#nestedatt--protocols--ospf--refresh))
- `segment_routing` (Attributes) Segment-Routing (SPRING) settings (see [below for nested schema](#nestedatt--protocols--ospf--segment_routing))
- `timers` (Attributes) Adjust routing timers (see [below for nested schema](#nestedatt--protocols--ospf--timers))

&lt;a id=&#34;nestedatt--protocols--ospf--aggregation&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.aggregation`

Optional:

- `timer` (Number) Delay timer

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|5-1800  &emsp;|Timer interval in seconds  |


&lt;a id=&#34;nestedatt--protocols--ospf--auto_cost&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.auto_cost`

Optional:

- `reference_bandwidth` (Number) Reference bandwidth method to assign cost

    &emsp;|Format     &emsp;|Description                            |
    |-------------|-----------------------------------------|
    &emsp;|1-4294967  &emsp;|Reference bandwidth cost in Mbits/sec  |


&lt;a id=&#34;nestedatt--protocols--ospf--capability&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.capability`

Optional:

- `opaque` (Boolean) Opaque LSA


&lt;a id=&#34;nestedatt--protocols--ospf--default_information&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.default_information`

Optional:

- `originate` (Attributes) Distribute a default route (see [below for nested schema](#nestedatt--protocols--ospf--default_information--originate))

&lt;a id=&#34;nestedatt--protocols--ospf--default_information--originate&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.default_information.originate`

Optional:

- `always` (Boolean) Always advertise a default route
- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--ospf--distance&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.distance`

Optional:

- `global` (Number) Administrative distance

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Administrative distance  |
- `ospf` (Attributes) OSPF administrative distance (see [below for nested schema](#nestedatt--protocols--ospf--distance--ospf))

&lt;a id=&#34;nestedatt--protocols--ospf--distance--ospf&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.distance.ospf`

Optional:

- `external` (Number) Distance for external routes

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|1-255   &emsp;|Distance for external routes  |
- `inter_area` (Number) Distance for inter-area routes

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|1-255   &emsp;|Distance for inter-area routes  |
- `intra_area` (Number) Distance for intra-area routes

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|1-255   &emsp;|Distance for intra-area routes  |



&lt;a id=&#34;nestedatt--protocols--ospf--graceful_restart&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.graceful_restart`

Optional:

- `grace_period` (Number) Maximum length of the grace period

    &emsp;|Format  &emsp;|Description                                    |
    |----------|-------------------------------------------------|
    &emsp;|1-1800  &emsp;|Maximum length of the grace period in seconds  |
- `helper` (Attributes) OSPF graceful-restart helpers (see [below for nested schema](#nestedatt--protocols--ospf--graceful_restart--helper))

&lt;a id=&#34;nestedatt--protocols--ospf--graceful_restart--helper&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.graceful_restart.helper`

Optional:

- `enable` (Attributes) Enable helper support (see [below for nested schema](#nestedatt--protocols--ospf--graceful_restart--helper--enable))
- `no_strict_lsa_checking` (Boolean) Disable strict LSA check
- `planned_only` (Boolean) Supported only planned restart
- `supported_grace_time` (Number) Supported grace timer

    &emsp;|Format   &emsp;|Description                |
    |-----------|-----------------------------|
    &emsp;|10-1800  &emsp;|Grace interval in seconds  |

&lt;a id=&#34;nestedatt--protocols--ospf--graceful_restart--helper--enable&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.graceful_restart.helper.enable`

Optional:

- `router_id` (List of String) Advertising Router-ID

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv4    &emsp;|Router-ID in IP address format  |




&lt;a id=&#34;nestedatt--protocols--ospf--ldp_sync&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.ldp_sync`

Optional:

- `holddown` (Number) Hold down timer for LDP-IGP cost restoration

    &emsp;|Format   &emsp;|Description                                                                                   |
    |-----------|------------------------------------------------------------------------------------------------|
    &emsp;|0-10000  &emsp;|Time to wait in seconds for LDP-IGP synchronization to occur before restoring interface cost  |


&lt;a id=&#34;nestedatt--protocols--ospf--log_adjacency_changes&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.log_adjacency_changes`

Optional:

- `detail` (Boolean) Log all state changes


&lt;a id=&#34;nestedatt--protocols--ospf--max_metric&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.max_metric`

Optional:

- `router_lsa` (Attributes) Advertise own Router-LSA with infinite distance (stub router) (see [below for nested schema](#nestedatt--protocols--ospf--max_metric--router_lsa))

&lt;a id=&#34;nestedatt--protocols--ospf--max_metric--router_lsa&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.max_metric.router_lsa`

Optional:

- `administrative` (Boolean) Administratively apply, for an indefinite period
- `on_shutdown` (Number) Advertise stub-router prior to full shutdown of OSPF

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|5-100   &emsp;|Time (seconds) to advertise self as stub-router  |
- `on_startup` (Number) Automatically advertise stub Router-LSA on startup of OSPF

    &emsp;|Format   &emsp;|Description                                      |
    |-----------|---------------------------------------------------|
    &emsp;|5-86400  &emsp;|Time (seconds) to advertise self as stub-router  |



&lt;a id=&#34;nestedatt--protocols--ospf--mpls_te&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.mpls_te`

Optional:

- `enable` (Boolean) Enable MPLS-TE functionality
- `router_address` (String) Stable IP address of the advertising router

    &emsp;|Format  &emsp;|Description                                  |
    |----------|-----------------------------------------------|
    &emsp;|ipv4    &emsp;|Stable IP address of the advertising router  |


&lt;a id=&#34;nestedatt--protocols--ospf--parameters&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.parameters`

Optional:

- `abr_type` (String) OSPF ABR type

    &emsp;|Format    &emsp;|Description        |
    |------------|---------------------|
    &emsp;|cisco     &emsp;|Cisco ABR type     |
    &emsp;|ibm       &emsp;|IBM ABR type       |
    &emsp;|shortcut  &emsp;|Shortcut ABR type  |
    &emsp;|standard  &emsp;|Standard ABR type  |
- `opaque_lsa` (Boolean) Enable the Opaque-LSA capability (rfc2370)
- `rfc1583_compatibility` (Boolean) Enable RFC1583 criteria for handling AS external routes
- `router_id` (String) Override default router identifier

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv4    &emsp;|Router-ID in IP address format  |


&lt;a id=&#34;nestedatt--protocols--ospf--redistribute&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.redistribute`

Optional:

- `babel` (Attributes) Redistribute Babel routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--babel))
- `bgp` (Attributes) Redistribute BGP routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--bgp))
- `connected` (Attributes) Redistribute connected routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--connected))
- `isis` (Attributes) Redistribute IS-IS routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--isis))
- `kernel` (Attributes) Redistribute Kernel routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--kernel))
- `rip` (Attributes) Redistribute RIP routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--rip))
- `static` (Attributes) Redistribute statically configured routes (see [below for nested schema](#nestedatt--protocols--ospf--redistribute--static))

&lt;a id=&#34;nestedatt--protocols--ospf--redistribute--babel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.redistribute.babel`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospf--redistribute--bgp&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.redistribute.bgp`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospf--redistribute--connected&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.redistribute.connected`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospf--redistribute--isis&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.redistribute.isis`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospf--redistribute--kernel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.redistribute.kernel`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospf--redistribute--rip&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.redistribute.rip`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospf--redistribute--static&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.redistribute.static`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--ospf--refresh&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.refresh`

Optional:

- `timers` (Number) Refresh timer

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|10-1800  &emsp;|Timer value in seconds  |


&lt;a id=&#34;nestedatt--protocols--ospf--segment_routing&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.segment_routing`

Optional:

- `global_block` (Attributes) Segment Routing Global Block label range (see [below for nested schema](#nestedatt--protocols--ospf--segment_routing--global_block))
- `local_block` (Attributes) Segment Routing Local Block label range (see [below for nested schema](#nestedatt--protocols--ospf--segment_routing--local_block))
- `maximum_label_depth` (Number) Maximum MPLS labels allowed for this router

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|1-16    &emsp;|MPLS label depth  |

&lt;a id=&#34;nestedatt--protocols--ospf--segment_routing--global_block&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.segment_routing.global_block`

Optional:

- `high_label_value` (Number) MPLS label upper bound

    &emsp;|Format      &emsp;|Description  |
    |--------------|---------------|
    &emsp;|16-1048575  &emsp;|Label value  |
- `low_label_value` (Number) MPLS label lower bound

    &emsp;|Format      &emsp;|Description                                   |
    |--------------|------------------------------------------------|
    &emsp;|16-1048575  &emsp;|Label value (recommended minimum value: 300)  |


&lt;a id=&#34;nestedatt--protocols--ospf--segment_routing--local_block&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.segment_routing.local_block`

Optional:

- `high_label_value` (Number) MPLS label upper bound

    &emsp;|Format      &emsp;|Description  |
    |--------------|---------------|
    &emsp;|16-1048575  &emsp;|Label value  |
- `low_label_value` (Number) MPLS label lower bound

    &emsp;|Format      &emsp;|Description                                   |
    |--------------|------------------------------------------------|
    &emsp;|16-1048575  &emsp;|Label value (recommended minimum value: 300)  |



&lt;a id=&#34;nestedatt--protocols--ospf--timers&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.timers`

Optional:

- `throttle` (Attributes) Throttling adaptive timers (see [below for nested schema](#nestedatt--protocols--ospf--timers--throttle))

&lt;a id=&#34;nestedatt--protocols--ospf--timers--throttle&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.timers.throttle`

Optional:

- `spf` (Attributes) OSPF SPF timers (see [below for nested schema](#nestedatt--protocols--ospf--timers--throttle--spf))

&lt;a id=&#34;nestedatt--protocols--ospf--timers--throttle--spf&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospf.timers.throttle.spf`

Optional:

- `delay` (Number) Delay from the first change received to SPF calculation

    &emsp;|Format    &emsp;|Description            |
    |------------|-------------------------|
    &emsp;|0-600000  &emsp;|Delay in milliseconds  |
- `initial_holdtime` (Number) Initial hold time between consecutive SPF calculations

    &emsp;|Format    &emsp;|Description                        |
    |------------|-------------------------------------|
    &emsp;|0-600000  &emsp;|Initial hold time in milliseconds  |
- `max_holdtime` (Number) Maximum hold time

    &emsp;|Format    &emsp;|Description                    |
    |------------|---------------------------------|
    &emsp;|0-600000  &emsp;|Max hold time in milliseconds  |





&lt;a id=&#34;nestedatt--protocols--ospfv3&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3`

Optional:

- `auto_cost` (Attributes) Calculate interface cost according to bandwidth (see [below for nested schema](#nestedatt--protocols--ospfv3--auto_cost))
- `default_information` (Attributes) Default route advertisment settings (see [below for nested schema](#nestedatt--protocols--ospfv3--default_information))
- `distance` (Attributes) Administrative distance (see [below for nested schema](#nestedatt--protocols--ospfv3--distance))
- `graceful_restart` (Attributes) Graceful Restart (see [below for nested schema](#nestedatt--protocols--ospfv3--graceful_restart))
- `log_adjacency_changes` (Attributes) Log adjacency state changes (see [below for nested schema](#nestedatt--protocols--ospfv3--log_adjacency_changes))
- `parameters` (Attributes) OSPFv3 specific parameters (see [below for nested schema](#nestedatt--protocols--ospfv3--parameters))
- `redistribute` (Attributes) Redistribute information from another routing protocol (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute))

&lt;a id=&#34;nestedatt--protocols--ospfv3--auto_cost&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.auto_cost`

Optional:

- `reference_bandwidth` (Number) Reference bandwidth method to assign cost

    &emsp;|Format     &emsp;|Description                            |
    |-------------|-----------------------------------------|
    &emsp;|1-4294967  &emsp;|Reference bandwidth cost in Mbits/sec  |


&lt;a id=&#34;nestedatt--protocols--ospfv3--default_information&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.default_information`

Optional:

- `originate` (Attributes) Distribute a default route (see [below for nested schema](#nestedatt--protocols--ospfv3--default_information--originate))

&lt;a id=&#34;nestedatt--protocols--ospfv3--default_information--originate&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.default_information.originate`

Optional:

- `always` (Boolean) Always advertise a default route
- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |



&lt;a id=&#34;nestedatt--protocols--ospfv3--distance&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.distance`

Optional:

- `global` (Number) Administrative distance

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Administrative distance  |
- `ospfv3` (Attributes) OSPFv3 administrative distance (see [below for nested schema](#nestedatt--protocols--ospfv3--distance--ospfv3))

&lt;a id=&#34;nestedatt--protocols--ospfv3--distance--ospfv3&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.distance.ospfv3`

Optional:

- `external` (Number) Distance for external routes

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|1-255   &emsp;|Distance for external routes  |
- `inter_area` (Number) Distance for inter-area routes

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|1-255   &emsp;|Distance for inter-area routes  |
- `intra_area` (Number) Distance for intra-area routes

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|1-255   &emsp;|Distance for intra-area routes  |



&lt;a id=&#34;nestedatt--protocols--ospfv3--graceful_restart&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.graceful_restart`

Optional:

- `grace_period` (Number) Maximum length of the grace period

    &emsp;|Format  &emsp;|Description                                    |
    |----------|-------------------------------------------------|
    &emsp;|1-1800  &emsp;|Maximum length of the grace period in seconds  |
- `helper` (Attributes) OSPF graceful-restart helpers (see [below for nested schema](#nestedatt--protocols--ospfv3--graceful_restart--helper))

&lt;a id=&#34;nestedatt--protocols--ospfv3--graceful_restart--helper&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.graceful_restart.helper`

Optional:

- `enable` (Attributes) Enable helper support (see [below for nested schema](#nestedatt--protocols--ospfv3--graceful_restart--helper--enable))
- `lsa_check_disable` (Boolean) Disable strict LSA check
- `planned_only` (Boolean) Supported only planned restart
- `supported_grace_time` (Number) Supported grace timer

    &emsp;|Format   &emsp;|Description                |
    |-----------|-----------------------------|
    &emsp;|10-1800  &emsp;|Grace interval in seconds  |

&lt;a id=&#34;nestedatt--protocols--ospfv3--graceful_restart--helper--enable&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.graceful_restart.helper.enable`

Optional:

- `router_id` (List of String) Advertising Router-ID

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv4    &emsp;|Router-ID in IP address format  |




&lt;a id=&#34;nestedatt--protocols--ospfv3--log_adjacency_changes&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.log_adjacency_changes`

Optional:

- `detail` (Boolean) Log all state changes


&lt;a id=&#34;nestedatt--protocols--ospfv3--parameters&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.parameters`

Optional:

- `router_id` (String) Override default router identifier

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|ipv4    &emsp;|Router-ID in IP address format  |


&lt;a id=&#34;nestedatt--protocols--ospfv3--redistribute&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.redistribute`

Optional:

- `babel` (Attributes) Redistribute Babel routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--babel))
- `bgp` (Attributes) Redistribute BGP routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--bgp))
- `connected` (Attributes) Redistribute connected routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--connected))
- `isis` (Attributes) Redistribute IS-IS routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--isis))
- `kernel` (Attributes) Redistribute kernel routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--kernel))
- `ripng` (Attributes) Redistribute RIPNG routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--ripng))
- `static` (Attributes) Redistribute static routes (see [below for nested schema](#nestedatt--protocols--ospfv3--redistribute--static))

&lt;a id=&#34;nestedatt--protocols--ospfv3--redistribute--babel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.redistribute.babel`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospfv3--redistribute--bgp&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.redistribute.bgp`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospfv3--redistribute--connected&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.redistribute.connected`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospfv3--redistribute--isis&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.redistribute.isis`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospfv3--redistribute--kernel&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.redistribute.kernel`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospfv3--redistribute--ripng&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.redistribute.ripng`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |


&lt;a id=&#34;nestedatt--protocols--ospfv3--redistribute--static&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.ospfv3.redistribute.static`

Optional:

- `metric` (Number) OSPF default metric

    &emsp;|Format      &emsp;|Description     |
    |--------------|------------------|
    &emsp;|0-16777214  &emsp;|Default metric  |
- `metric_type` (Number) OSPF metric type for default routes

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|1-2     &emsp;|Set OSPF External Type 1/2 metrics  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |




&lt;a id=&#34;nestedatt--protocols--static&#34;&gt;&lt;/a&gt;
### Nested Schema for `protocols.static`



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
