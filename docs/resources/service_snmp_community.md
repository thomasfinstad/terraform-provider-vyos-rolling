---
page_title: "vyos_service_snmp_community Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Simple Network Management Protocol (SNMP)⯯Community name
---

# vyos_service_snmp_community (Resource)
<center>

*service*  
⯯  
Simple Network Management Protocol (SNMP)  
⯯  
**Community name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `authorization` (String) Authorization type

    |Format  &emsp;|Description  |
    |----------|---------------|
    |ro      &emsp;|Read-Only    |
    |rw      &emsp;|Read-Write   |
- `client` (List of String) IP address of SNMP client allowed to contact system
- `network` (List of String) Subnet of SNMP client(s) allowed to contact system

    |Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    |ipv4net  &emsp;|IP address and prefix length    |
    |ipv6net  &emsp;|IPv6 address and prefix length  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `community` (String) Community name


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
