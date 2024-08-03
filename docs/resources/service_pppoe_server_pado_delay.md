---
page_title: "vyos_service_pppoe_server_pado_delay Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Point to Point over Ethernet (PPPoE) Server⯯PADO delays
---

# vyos_service_pppoe_server_pado_delay (Resource)
<center>

*service*  
⯯  
Point to Point over Ethernet (PPPoE) Server  
⯯  
**PADO delays**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `sessions` (Number) Number of sessions

    &emsp;|Format    &emsp;|Description         |
    |------------|----------------------|
    &emsp;|1-999999  &emsp;|Number of sessions  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `pado_delay` (String) PADO delays

    &emsp;|Format    &emsp;|Description              |
    |------------|---------------------------|
    &emsp;|disable   &emsp;|Disable new connections  |
    &emsp;|1-999999  &emsp;|Number in ms             |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
