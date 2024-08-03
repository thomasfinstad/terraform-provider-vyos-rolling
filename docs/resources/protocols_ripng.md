---
page_title: "vyos_protocols_ripng Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Routing Information Protocol (RIPng) parameters
---

# vyos_protocols_ripng (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
**Routing Information Protocol (RIPng) parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `aggregate_address` (List of String) Aggregate RIPng route announcement

    &emsp;|Format   &emsp;|Description                         |
    |-----------|--------------------------------------|
    &emsp;|ipv6net  &emsp;|Aggregate RIPng route announcement  |
- `default_metric` (Number) Metric of redistributed routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|1-16    &emsp;|Default metric  |
- `network` (List of String) RIPng network

    &emsp;|Format   &emsp;|Description    |
    |-----------|-----------------|
    &emsp;|ipv6net  &emsp;|RIPng network  |
- `passive_interface` (List of String) Passive interface

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|txt     &emsp;|Suppress routing updates on interface  |
- `route` (List of String) RIPng static route

    &emsp;|Format   &emsp;|Description         |
    |-----------|----------------------|
    &emsp;|ipv6net  &emsp;|RIPng static route  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
