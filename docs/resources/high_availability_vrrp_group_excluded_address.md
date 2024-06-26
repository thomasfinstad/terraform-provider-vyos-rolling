---
page_title: "vyos_high_availability_vrrp_group_excluded_address Resource - vyos"

subcategory: "High Availability"

description: |- 
  High availability settings⯯Virtual Router Redundancy Protocol settings⯯VRRP group⯯Virtual address (If you need additional IPv4 and IPv6 in same group)
---

# vyos_high_availability_vrrp_group_excluded_address (Resource)
<center>

High availability settings  
⯯  
Virtual Router Redundancy Protocol settings  
⯯  
VRRP group  
⯯  
**Virtual address (If you need additional IPv4 and IPv6 in same group)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `excluded_address_id` (String) Virtual address (If you need additional IPv4 and IPv6 in same group)

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4net  &emsp;|IPv4 address and prefix length  |
    &emsp;|ipv6net  &emsp;|IPv6 address and prefix length  |
    &emsp;|ipv4     &emsp;|IPv4 address                    |
    &emsp;|ipv6     &emsp;|IPv6 address                    |
- `group_id` (String) VRRP group

### Optional

- `interface` (String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
