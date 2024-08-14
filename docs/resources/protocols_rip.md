---
page_title: "vyos_protocols_rip Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Routing Information Protocol (RIP) parameters
---

# vyos_protocols_rip (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
**Routing Information Protocol (RIP) parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `default_distance` (Number) Administrative distance

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-255   &emsp;|Administrative distance  |
- `default_metric` (Number) Metric of redistributed routes

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|1-16    &emsp;|Default metric  |
- `neighbor` (List of String) Neighbor router

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|ipv4    &emsp;|Neighbor router  |
- `network` (List of String) RIP network

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|ipv4net  &emsp;|RIP network  |
- `passive_interface` (List of String) Suppress routing updates on an interface

    &emsp;|Format   &emsp;|Description                                              |
    |-----------|-----------------------------------------------------------|
    &emsp;|txt      &emsp;|Interface to be passive (i.e. suppress routing updates)  |
    &emsp;|default  &emsp;|Default to suppress routing updates on all interfaces    |
- `route` (List of String) RIP static route

    &emsp;|Format   &emsp;|Description       |
    |-----------|--------------------|
    &emsp;|ipv4net  &emsp;|RIP static route  |
- `route_map` (String) Specify route-map name to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `version` (String) Limit RIP protocol version

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|1       &emsp;|Allow RIPv1 only  |
    &emsp;|2       &emsp;|Allow RIPv2 only  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  