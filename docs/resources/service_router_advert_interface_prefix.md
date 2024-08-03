---
page_title: "vyos_service_router_advert_interface_prefix Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯IPv6 Router Advertisements (RAs) service⯯Interface to send RA on⯯IPv6 prefix to be advertised in Router Advertisements (RAs)
---

# vyos_service_router_advert_interface_prefix (Resource)
<center>

*service*  
⯯  
IPv6 Router Advertisements (RAs) service  
⯯  
Interface to send RA on  
⯯  
**IPv6 prefix to be advertised in Router Advertisements (RAs)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `decrement_lifetime` (Boolean) Lifetime is decremented by the number of seconds since the last RA - use in conjunction with a DHCPv6-PD prefix
- `deprecate_prefix` (Boolean) Upon shutdown, this option will deprecate the prefix by announcing it in the shutdown RA
- `no_autonomous_flag` (Boolean) Prefix can not be used for stateless address auto-configuration
- `no_on_link_flag` (Boolean) Prefix can not be used for on-link determination
- `preferred_lifetime` (String) Time in seconds that the prefix will remain preferred

    &emsp;|Format    &emsp;|Description                                            |
    |------------|---------------------------------------------------------|
    &emsp;|u32       &emsp;|Time in seconds that the prefix will remain preferred  |
    &emsp;|infinity  &emsp;|Prefix will remain preferred forever                   |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `valid_lifetime` (String) Time in seconds that the prefix will remain valid

    &emsp;|Format        &emsp;|Description                                        |
    |----------------|-----------------------------------------------------|
    &emsp;|1-4294967295  &emsp;|Time in seconds that the prefix will remain valid  |
    &emsp;|infinity      &emsp;|Prefix will remain preferred forever               |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface to send RA on
- `prefix` (String) IPv6 prefix to be advertised in Router Advertisements (RAs)

    &emsp;|Format   &emsp;|Description                   |
    |-----------|--------------------------------|
    &emsp;|ipv6net  &emsp;|IPv6 prefix to be advertized  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
