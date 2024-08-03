---
page_title: "vyos_service_ipoe_server_client_ip_pool Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Internet Protocol over Ethernet (IPoE) Server⯯Client IP pool
---

# vyos_service_ipoe_server_client_ip_pool (Resource)
<center>

*service*  
⯯  
Internet Protocol over Ethernet (IPoE) Server  
⯯  
**Client IP pool**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `next_pool` (String) Next pool name

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|Name of IP pool  |
- `range` (List of String) Range of IP addresses

    &emsp;|Format     &emsp;|Description                            |
    |-------------|-----------------------------------------|
    &emsp;|ipv4net    &emsp;|IPv4 prefix                            |
    &emsp;|ipv4range  &emsp;|IPv4 address range inside /24 network  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `client_ip_pool` (String) Client IP pool

    &emsp;|Format  &emsp;|Description      |
    |----------|-------------------|
    &emsp;|txt     &emsp;|Name of IP pool  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
