---
page_title: "vyos_vrf_name_protocols_ospf_area Resource - terraform-provider-vyos"
subcategory: "vrf"
description: |-
  Virtual Routing and Forwarding
  ⯯
  Virtual Routing and Forwarding instance
  ⯯
  Routing protocol parameters
  ⯯
  Open Shortest Path First (OSPF)
  ⯯
  OSPF area settings
---

# vyos_vrf_name_protocols_ospf_area (Resource)
<center>

Virtual Routing and Forwarding
⯯
Virtual Routing and Forwarding instance
⯯
Routing protocol parameters
⯯
Open Shortest Path First (OSPF)
⯯
**OSPF area settings**


</center>

## Schema

### Required

- `area_id` (String) OSPF area settings

    &emsp;|Format  &emsp;|Description                                  |
    |----------|-----------------------------------------------|
    &emsp;|u32     &emsp;|OSPF area number in decimal notation         |
    &emsp;|ipv4    &emsp;|OSPF area number in dotted decimal notation  |
- `name_id` (String) Virtual Routing and Forwarding instance

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Optional

- `area_type` (Attributes) Area type (see [below for nested schema](#nestedatt--area_type))
- `authentication` (String) OSPF area authentication type

    &emsp;|Format              &emsp;|Description                    |
    |----------------------|---------------------------------|
    &emsp;|plaintext-password  &emsp;|Use plain-text authentication  |
    &emsp;|md5                 &emsp;|Use MD5 authentication         |
- `export_list` (Number) Set the filter for networks announced to other areas

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|u32     &emsp;|Access-list number  |
- `import_list` (Number) Set the filter for networks from other areas announced

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|u32     &emsp;|Access-list number  |
- `network` (List of String) OSPF network

    &emsp;|Format   &emsp;|Description   |
    |-----------|----------------|
    &emsp;|ipv4net  &emsp;|OSPF network  |
- `shortcut` (String) Area shortcut mode

    &emsp;|Format   &emsp;|Description                |
    |-----------|-----------------------------|
    &emsp;|default  &emsp;|Set default                |
    &emsp;|disable  &emsp;|Disable shortcutting mode  |
    &emsp;|enable   &emsp;|Enable shortcutting mode   |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--area_type&#34;&gt;&lt;/a&gt;
### Nested Schema for `area_type`

Optional:

- `normal` (Boolean) Normal OSPF area
- `nssa` (Attributes) Not-So-Stubby OSPF area (see [below for nested schema](#nestedatt--area_type--nssa))
- `stub` (Attributes) Stub OSPF area (see [below for nested schema](#nestedatt--area_type--stub))

&lt;a id=&#34;nestedatt--area_type--nssa&#34;&gt;&lt;/a&gt;
### Nested Schema for `area_type.nssa`

Optional:

- `default_cost` (Number) Summary-default cost of an NSSA area

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Summary default cost  |
- `no_summary` (Boolean) Do not inject inter-area routes into stub
- `translate` (String) Configure NSSA-ABR

    &emsp;|Format     &emsp;|Description                 |
    |-------------|------------------------------|
    &emsp;|always     &emsp;|Always translate LSA types  |
    &emsp;|candidate  &emsp;|Translate for election      |
    &emsp;|never      &emsp;|Never translate LSA types   |


&lt;a id=&#34;nestedatt--area_type--stub&#34;&gt;&lt;/a&gt;
### Nested Schema for `area_type.stub`

Optional:

- `default_cost` (Number) Summary-default cost

    &emsp;|Format      &emsp;|Description           |
    |--------------|------------------------|
    &emsp;|0-16777215  &emsp;|Summary default cost  |
- `no_summary` (Boolean) Do not inject inter-area routes into the stub



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
