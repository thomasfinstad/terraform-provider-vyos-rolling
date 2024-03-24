---
page_title: "vyos_firewall_group_dynamic_group_ipv6_address_group Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall⯯Firewall group⯯Firewall dynamic group⯯Firewall dynamic IPv6 address group
---

# vyos_firewall_group_dynamic_group_ipv6_address_group (Resource)
<center>

Firewall  
⯯  
Firewall group  
⯯  
Firewall dynamic group  
⯯  
**Firewall dynamic IPv6 address group**


</center>

## Schema

### Required

- `ipv6_address_group_id` (String) Firewall dynamic IPv6 address group

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
