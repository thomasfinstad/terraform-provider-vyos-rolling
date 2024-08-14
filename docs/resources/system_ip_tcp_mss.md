---
page_title: "vyos_system_ip_tcp_mss Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯IPv4 Settings⯯IPv4 TCP parameters⯯IPv4 TCP MSS probing options
---

# vyos_system_ip_tcp_mss (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
IPv4 Settings  
⯯  
IPv4 TCP parameters  
⯯  
**IPv4 TCP MSS probing options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `base` (Number) Base MSS to start probing from (applicable to &#34;probing force&#34;)

    |Format   &emsp;|Description                                 |
    |-----------|----------------------------------------------|
    |48-1460  &emsp;|Base MSS value for probing (default: 1024)  |
- `floor` (Number) Minimum MSS to stop probing at (default: 48)

    |Format   &emsp;|Description                 |
    |-----------|------------------------------|
    |48-1460  &emsp;|Minimum MSS value to probe  |
- `probing` (String) Attempt to lower the MSS if TCP connections fail to establish

    |Format              &emsp;|Description                                                  |
    |----------------------|---------------------------------------------------------------|
    |on-icmp-black-hole  &emsp;|Attempt TCP MSS probing when an ICMP black hole is detected  |
    |force               &emsp;|Attempt TCP MSS probing by default                           |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
