---
page_title: "vyos_high_availability_vrrp_global_parameters Resource - vyos"

subcategory: "High Availability"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  High availability settings⯯Virtual Router Redundancy Protocol settings⯯VRRP global parameters
---

# vyos_high_availability_vrrp_global_parameters (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

High availability settings  
⯯  
Virtual Router Redundancy Protocol settings  
⯯  
**VRRP global parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `startup_delay` (Number) Time VRRP startup process (in seconds)

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|1-600   &emsp;|Interval in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `version` (String) Default VRRP version to use, IPv6 always uses VRRP version 3

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|2       &emsp;|VRRP version 2  |
    &emsp;|3       &emsp;|VRRP version 3  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
