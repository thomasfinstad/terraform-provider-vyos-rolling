---
page_title: "vyos_service_ipoe_server_authentication_radius_dynamic_author Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Internet Protocol over Ethernet (IPoE) Server⯯Client authentication methods⯯RADIUS based user authentication⯯Dynamic Authorization Extension/Change of Authorization server
---

# vyos_service_ipoe_server_authentication_radius_dynamic_author (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Internet Protocol over Ethernet (IPoE) Server  
⯯  
Client authentication methods  
⯯  
RADIUS based user authentication  
⯯  
**Dynamic Authorization Extension/Change of Authorization server**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `key` (String) Shared secret for Dynamic Authorization Extension server
- `port` (Number) Port for Dynamic Authorization Extension server (DM/CoA)

    |Format   &emsp;|Description  |
    |-----------|---------------|
    |1-65535  &emsp;|TCP port     |
- `server` (String) IP address for Dynamic Authorization Extension server (DM/CoA)

    |Format  &emsp;|Description                                    |
    |----------|-------------------------------------------------|
    |ipv4    &emsp;|IPv4 address for dynamic authorization server  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
