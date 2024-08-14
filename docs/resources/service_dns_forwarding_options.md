---
page_title: "vyos_service_dns_forwarding_options Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Domain Name System (DNS) related services⯯DNS forwarding⯯DNS server options
---

# vyos_service_dns_forwarding_options (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Domain Name System (DNS) related services  
⯯  
DNS forwarding  
⯯  
**DNS server options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `ecs_add_for` (List of String) Client netmask for which EDNS Client Subnet will be added

    |Format    &emsp;|Description                                        |
    |------------|-----------------------------------------------------|
    |ipv4net   &emsp;|IPv4 prefix to match                               |
    |!ipv4net  &emsp;|Match everything except the specified IPv4 prefix  |
    |ipv6net   &emsp;|IPv6 prefix to match                               |
    |!ipv6net  &emsp;|Match everything except the specified IPv6 prefix  |
- `ecs_ipv4_bits` (Number) Number of bits of IPv4 address to pass for EDNS Client Subnet

    |Format  &emsp;|Description                     |
    |----------|----------------------------------|
    |0-32    &emsp;|Number of bits of IPv4 address  |
- `edns_subnet_allow_list` (List of String) Netmask or domain that we should enable EDNS subnet for

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|Netmask or domain  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
