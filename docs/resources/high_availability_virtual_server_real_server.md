---
page_title: "vyos_high_availability_virtual_server_real_server Resource - vyos"

subcategory: "High Availability"

description: |- 
  High availability settings⯯Load-balancing virtual server alias⯯Real server address
---

# vyos_high_availability_virtual_server_real_server (Resource)
<center>

High availability settings  
⯯  
Load-balancing virtual server alias  
⯯  
**Real server address**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `connection_timeout` (Number) Server connection timeout

    |Format   &emsp;|Description                          |
    |-----------|---------------------------------------|
    |1-86400  &emsp;|Connection timeout to remote server  |
- `health_check` (Attributes) Health check script (see [below for nested schema](#nestedatt--health_check))
- `port` (Number) Port number used by connection

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |0-65535  &emsp;|Numeric IP port  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `real_server` (String) Real server address
- `virtual_server` (String) Load-balancing virtual server alias


<a id="nestedatt--health_check"></a>
### Nested Schema for `health_check`

Optional:

- `script` (String) Health check script file


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
