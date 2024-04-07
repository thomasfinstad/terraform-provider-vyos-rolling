---
page_title: "vyos_container_name_network Resource - vyos"

subcategory: "Container"

description: |- 
  Container applications⯯Container name⯯Attach user defined network to container
---

# vyos_container_name_network (Resource)
<center>

Container applications  
⯯  
Container name  
⯯  
**Attach user defined network to container**


</center>

## Schema

### Required

- `name_id` (String) Container name
- `network_id` (String) Attach user defined network to container

### Optional

- `address` (List of String) Assign static IP address to container

    &emsp;|Format  &emsp;|Description   |
    |----------|----------------|
    &emsp;|ipv4    &emsp;|IPv4 address  |
    &emsp;|ipv6    &emsp;|IPv6 address  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
