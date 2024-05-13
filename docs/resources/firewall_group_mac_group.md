---
page_title: "vyos_firewall_group_mac_group Resource - vyos"

subcategory: "Firewall"

description: |- 
  Firewall⯯Firewall group⯯Firewall mac-group
---

# vyos_firewall_group_mac_group (Resource)
<center>

Firewall  
⯯  
Firewall group  
⯯  
**Firewall mac-group**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `mac_group_id` (String) Firewall mac-group

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `include` (List of String) Include another mac-group
- `mac_address` (List of String) Mac-group member

    &emsp;|Format   &emsp;|Description           |
    |-----------|------------------------|
    &emsp;|macaddr  &emsp;|MAC address to match  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
