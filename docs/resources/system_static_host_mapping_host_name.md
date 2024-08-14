---
page_title: "vyos_system_static_host_mapping_host_name Resource - vyos"

subcategory: "System"

description: |- 
  system⯯Map host names to addresses⯯Host name for static address mapping
---

# vyos_system_static_host_mapping_host_name (Resource)
<center>

*system*  
⯯  
Map host names to addresses  
⯯  
**Host name for static address mapping**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `alias` (List of String) Alias for this address
- `inet` (List of String) IP Address

    |Format  &emsp;|Description   |
    |----------|----------------|
    |ipv4    &emsp;|IPv4 address  |
    |ipv6    &emsp;|IPv6 address  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `host_name` (String) Host name for static address mapping


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
