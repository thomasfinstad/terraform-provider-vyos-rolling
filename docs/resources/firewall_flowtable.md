---
page_title: "vyos_firewall_flowtable Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall⯯Flowtable
---

# vyos_firewall_flowtable (Resource)
<center>

Firewall
⯯
**Flowtable**


</center>

## Schema

### Required

- `flowtable_id` (String) Flowtable

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `interface` (List of String) Interfaces to use this flowtable
- `offload` (String) Offloading method

    &emsp;|Format    &emsp;|Description       |
    |------------|--------------------|
    &emsp;|hardware  &emsp;|Hardware offload  |
    &emsp;|software  &emsp;|Software offload  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
