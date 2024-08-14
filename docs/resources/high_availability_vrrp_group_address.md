---
page_title: "vyos_high_availability_vrrp_group_address Resource - vyos"

subcategory: "High Availability"

description: |- 
  High availability settings⯯Virtual Router Redundancy Protocol settings⯯VRRP group⯯Virtual IP address
---

# vyos_high_availability_vrrp_group_address (Resource)
<center>

High availability settings  
⯯  
Virtual Router Redundancy Protocol settings  
⯯  
VRRP group  
⯯  
**Virtual IP address**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `interface` (String) Interface to use

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `address` (String) Virtual IP address

    |Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    |ipv4net  &emsp;|IPv4 address and prefix length  |
    |ipv6net  &emsp;|IPv6 address and prefix length  |
    |ipv4     &emsp;|IPv4 address                    |
    |ipv6     &emsp;|IPv6 address                    |
- `group` (String) VRRP group


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
