---
page_title: "vyos_protocols_pim_igmp Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Protocol Independent Multicast (PIM) and IGMP⯯Internet Group Management Protocol (IGMP) options
---

# vyos_protocols_pim_igmp (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Protocol Independent Multicast (PIM) and IGMP  
⯯  
**Internet Group Management Protocol (IGMP) options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `watermark_warning` (Number) Configure group limit for watermark warning

    &emsp;|Format   &emsp;|Description                                |
    |-----------|---------------------------------------------|
    &emsp;|1-65535  &emsp;|Group count to generate watermark warning  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
