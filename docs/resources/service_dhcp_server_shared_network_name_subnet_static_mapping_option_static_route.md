---
page_title: "vyos_service_dhcp_server_shared_network_name_subnet_static_mapping_option_static_route Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Dynamic Host Configuration Protocol (DHCP) for DHCP server⯯Name of DHCP shared network⯯DHCP subnet for shared network⯯Hostname for static mapping reservation⯯DHCP option⯯Classless static route destination subnet
---

# vyos_service_dhcp_server_shared_network_name_subnet_static_mapping_option_static_route (Resource)
<center>

*service*  
⯯  
Dynamic Host Configuration Protocol (DHCP) for DHCP server  
⯯  
Name of DHCP shared network  
⯯  
DHCP subnet for shared network  
⯯  
Hostname for static mapping reservation  
⯯  
DHCP option  
⯯  
**Classless static route destination subnet**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `next_hop` (String) IP address of router to be used to reach the destination subnet

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address of router  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `shared_network_name` (String) Name of DHCP shared network
- `static_mapping` (String) Hostname for static mapping reservation
- `static_route` (String) Classless static route destination subnet

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 address and prefix length  |
- `subnet` (String) DHCP subnet for shared network

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 address and prefix length  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
