---
page_title: "vyos_protocols_ospf_area_range Resource - vyos"

subcategory: "Protocols"

description: |- 
  protocols⯯Open Shortest Path First (OSPF)⯯OSPF area settings⯯Summarize routes matching a prefix (border routers only)
---

# vyos_protocols_ospf_area_range (Resource)
<center>

*protocols*  
⯯  
Open Shortest Path First (OSPF)  
⯯  
OSPF area settings  
⯯  
**Summarize routes matching a prefix (border routers only)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `cost` (Number) Metric for this range

    &emsp;|Format      &emsp;|Description            |
    |--------------|-------------------------|
    &emsp;|0-16777215  &emsp;|Metric for this range  |
- `not_advertise` (Boolean) Do not advertise this range
- `substitute` (String) Advertise area range as another prefix

    &emsp;|Format   &emsp;|Description                             |
    |-----------|------------------------------------------|
    &emsp;|ipv4net  &emsp;|Advertise area range as another prefix  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `area` (String) OSPF area settings

    &emsp;|Format  &emsp;|Description                                  |
    |----------|-----------------------------------------------|
    &emsp;|u32     &emsp;|OSPF area number in decimal notation         |
    &emsp;|ipv4    &emsp;|OSPF area number in dotted decimal notation  |
- `range` (String) Summarize routes matching a prefix (border routers only)

    &emsp;|Format   &emsp;|Description        |
    |-----------|---------------------|
    &emsp;|ipv4net  &emsp;|Area range prefix  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  