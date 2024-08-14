---
page_title: "vyos_protocols_ospf_area Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Open Shortest Path First (OSPF)⯯OSPF area settings
---

# vyos_protocols_ospf_area (Resource)
<center>

*protocols*  
⯯  
Open Shortest Path First (OSPF)  
⯯  
**OSPF area settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `area_type` (Attributes) Area type (see [below for nested schema](#nestedatt--area_type))
- `authentication` (String) OSPF area authentication type

    |Format              &emsp;|Description                    |
    |----------------------|---------------------------------|
    |plaintext-password  &emsp;|Use plain-text authentication  |
    |md5                 &emsp;|Use MD5 authentication         |
- `export_list` (Number) Set the filter for networks announced to other areas

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |u32     &emsp;|Access-list number  |
- `import_list` (Number) Set the filter for networks from other areas announced

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |u32     &emsp;|Access-list number  |
- `network` (List of String) OSPF network

    |Format   &emsp;|Description   |
    |-----------|----------------|
    |ipv4net  &emsp;|OSPF network  |
- `shortcut` (String) Area shortcut mode

    |Format   &emsp;|Description                |
    |-----------|-----------------------------|
    |default  &emsp;|Set default                |
    |disable  &emsp;|Disable shortcutting mode  |
    |enable   &emsp;|Enable shortcutting mode   |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `area` (String) OSPF area settings

    |Format  &emsp;|Description                                  |
    |----------|-----------------------------------------------|
    |u32     &emsp;|OSPF area number in decimal notation         |
    |ipv4    &emsp;|OSPF area number in dotted decimal notation  |


<a id="nestedatt--area_type"></a>
### Nested Schema for `area_type`

Optional:

- `normal` (Boolean) Normal OSPF area
- `nssa` (Attributes) Not-So-Stubby OSPF area (see [below for nested schema](#nestedatt--area_type--nssa))
- `stub` (Attributes) Stub OSPF area (see [below for nested schema](#nestedatt--area_type--stub))

<a id="nestedatt--area_type--nssa"></a>
### Nested Schema for `area_type.nssa`

Optional:

- `default_cost` (Number) Summary-default cost of an NSSA area

    |Format      &emsp;|Description           |
    |--------------|------------------------|
    |0-16777215  &emsp;|Summary default cost  |
- `no_summary` (Boolean) Do not inject inter-area routes into stub
- `translate` (String) Configure NSSA-ABR

    |Format     &emsp;|Description                 |
    |-------------|------------------------------|
    |always     &emsp;|Always translate LSA types  |
    |candidate  &emsp;|Translate for election      |
    |never      &emsp;|Never translate LSA types   |


<a id="nestedatt--area_type--stub"></a>
### Nested Schema for `area_type.stub`

Optional:

- `default_cost` (Number) Summary-default cost

    |Format      &emsp;|Description           |
    |--------------|------------------------|
    |0-16777215  &emsp;|Summary default cost  |
- `no_summary` (Boolean) Do not inject inter-area routes into the stub



<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
