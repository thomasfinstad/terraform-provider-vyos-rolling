---
page_title: "vyos_service_router_advert_interface_nat64prefix Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯IPv6 Router Advertisements (RAs) service⯯Interface to send RA on⯯NAT64 prefix included in the router advertisements
---

# vyos_service_router_advert_interface_nat64prefix (Resource)
<center>

*service*  
⯯  
IPv6 Router Advertisements (RAs) service  
⯯  
Interface to send RA on  
⯯  
**NAT64 prefix included in the router advertisements**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `valid_lifetime` (Number) Time in seconds that the prefix will remain valid

    |Format   &emsp;|Description                                        |
    |-----------|-----------------------------------------------------|
    |4-65528  &emsp;|Time in seconds that the prefix will remain valid  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface to send RA on
- `nat64prefix` (String) NAT64 prefix included in the router advertisements

    |Format   &emsp;|Description                   |
    |-----------|--------------------------------|
    |ipv6net  &emsp;|IPv6 prefix to be advertized  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
