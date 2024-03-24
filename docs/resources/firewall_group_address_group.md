---
page_title: "vyos_firewall_group_address_group Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall⯯Firewall group⯯Firewall address-group
---

# vyos_firewall_group_address_group (Resource)
<center>

Firewall  
⯯  
Firewall group  
⯯  
**Firewall address-group**


</center>

## Schema

### Required

- `address_group_id` (String) Firewall address-group

### Optional

- `address` (List of String) Address-group member

    &emsp;|Format     &emsp;|Description                                     |
    |-------------|--------------------------------------------------|
    &emsp;|ipv4       &emsp;|IPv4 address to match                           |
    &emsp;|ipv4range  &emsp;|IPv4 range to match (e.g. 10.0.0.1-10.0.0.200)  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `include` (List of String) Include another address-group
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
