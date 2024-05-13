---
page_title: "vyos_firewall_ipv4_name Resource - vyos"

subcategory: "Firewall"

description: |- 
  Firewall⯯IPv4 firewall⯯IPv4 custom firewall
---

# vyos_firewall_ipv4_name (Resource)
<center>

Firewall  
⯯  
IPv4 firewall  
⯯  
**IPv4 custom firewall**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `name_id` (String) IPv4 custom firewall

### Optional

- `default_action` (String) Default-action for rule-set

    &emsp;|Format    &emsp;|Description                                                                    |
    |------------|---------------------------------------------------------------------------------|
    &emsp;|drop      &emsp;|Drop if no prior rules are hit                                                 |
    &emsp;|jump      &emsp;|Jump to another chain if no prior rules are hit                                |
    &emsp;|reject    &emsp;|Drop and notify source if no prior rules are hit                               |
    &emsp;|return    &emsp;|Return from the current chain and continue at the next rule of the last chain  |
    &emsp;|accept    &emsp;|Accept if no prior rules are hit                                               |
    &emsp;|continue  &emsp;|Continue parsing next rule                                                     |
- `default_jump_target` (String) Set jump target. Action jump must be defined in default-action to use this setting
- `default_log` (Boolean) Log packets hitting default-action
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

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
