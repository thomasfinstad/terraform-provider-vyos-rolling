---
page_title: "vyos_firewall_bridge_name Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Bridge firewall
  ⯯
  Bridge custom firewall
---

# vyos_firewall_bridge_name (Resource)
<center>

Firewall
⯯
Bridge firewall
⯯
**Bridge custom firewall**


</center>

## Schema

### Required

- `name_id` (String) Bridge custom firewall

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

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
