---
page_title: "vyos_service_router_advert_interface_route Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯IPv6 Router Advertisements (RAs) service⯯Interface to send RA on⯯IPv6 route to be advertised in Router Advertisements (RAs)
---

# vyos_service_router_advert_interface_route (Resource)
<center>

*service*  
⯯  
IPv6 Router Advertisements (RAs) service  
⯯  
Interface to send RA on  
⯯  
**IPv6 route to be advertised in Router Advertisements (RAs)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `no_remove_route` (Boolean) Do not announce this route with a zero second lifetime upon shutdown
- `route_preference` (String) Preference associated with the route,

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|low     &emsp;|Route has low preference     |
    &emsp;|medium  &emsp;|Route has medium preference  |
    &emsp;|high    &emsp;|Route has high preference    |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `valid_lifetime` (String) Time in seconds that the route will remain valid

    &emsp;|Format        &emsp;|Description                                       |
    |----------------|----------------------------------------------------|
    &emsp;|1-4294967295  &emsp;|Time in seconds that the route will remain valid  |
    &emsp;|infinity      &emsp;|Route will remain preferred forever               |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface to send RA on
- `route` (String) IPv6 route to be advertised in Router Advertisements (RAs)

    &emsp;|Format   &emsp;|Description                  |
    |-----------|-------------------------------|
    &emsp;|ipv6net  &emsp;|IPv6 route to be advertized  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
