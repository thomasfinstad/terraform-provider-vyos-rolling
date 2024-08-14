---
page_title: "vyos_service_pppoe_server_authentication_radius_rate_limit Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Point to Point over Ethernet (PPPoE) Server⯯Authentication for remote access PPPoE Server⯯RADIUS based user authentication⯯Upload/Download speed limits
---

# vyos_service_pppoe_server_authentication_radius_rate_limit (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Point to Point over Ethernet (PPPoE) Server  
⯯  
Authentication for remote access PPPoE Server  
⯯  
RADIUS based user authentication  
⯯  
**Upload/Download speed limits**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `attribute` (String) RADIUS attribute that contains rate information
- `enable` (Boolean) Enable bandwidth shaping via RADIUS
- `multiplier` (String) Shaper multiplier

    |Format        &emsp;|Description        |
    |----------------|---------------------|
    |&lt;0.001-1000&gt;  &emsp;|Shaper multiplier  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vendor` (String) Vendor dictionary

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
