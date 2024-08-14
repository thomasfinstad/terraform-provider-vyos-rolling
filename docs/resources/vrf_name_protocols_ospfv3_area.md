---
page_title: "vyos_vrf_name_protocols_ospfv3_area Resource - vyos"

subcategory: "Vrf"

description: |- 
  Virtual Routing and Forwarding⯯Virtual Routing and Forwarding instance⯯Routing protocol parameters⯯Open Shortest Path First (OSPF) for IPv6⯯OSPFv3 Area
---

# vyos_vrf_name_protocols_ospfv3_area (Resource)
<center>

Virtual Routing and Forwarding  
⯯  
Virtual Routing and Forwarding instance  
⯯  
Routing protocol parameters  
⯯  
Open Shortest Path First (OSPF) for IPv6  
⯯  
**OSPFv3 Area**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `area_type` (Attributes) OSPFv3 Area type (see [below for nested schema](#nestedatt--area_type))
- `export_list` (String) Name of export-list
- `import_list` (String) Name of import-list
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `area` (String) OSPFv3 Area

    |Format  &emsp;|Description                  |
    |----------|-------------------------------|
    |u32     &emsp;|Area ID as a decimal value   |
    |ipv4    &emsp;|Area ID in IP address forma  |
- `name` (String) Virtual Routing and Forwarding instance

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |


<a id="nestedatt--area_type"></a>
### Nested Schema for `area_type`

Optional:

- `nssa` (Attributes) NSSA OSPFv3 area (see [below for nested schema](#nestedatt--area_type--nssa))
- `stub` (Attributes) Stub OSPFv3 area (see [below for nested schema](#nestedatt--area_type--stub))

<a id="nestedatt--area_type--nssa"></a>
### Nested Schema for `area_type.nssa`

Optional:

- `default_information_originate` (Boolean) Originate Type 7 default into NSSA area
- `no_summary` (Boolean) Do not inject inter-area routes into the stub


<a id="nestedatt--area_type--stub"></a>
### Nested Schema for `area_type.stub`

Optional:

- `no_summary` (Boolean) Do not inject inter-area routes into the stub



<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
