---
page_title: "vyos_firewall_group_port_group Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Firewall group
  ⯯
  Firewall port-group
---

# vyos_firewall_group_port_group (Resource)
<center>

Firewall
⯯
Firewall group
⯯
**Firewall port-group**


</center>

## Schema

### Required

- `port_group_id` (String) Firewall port-group

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `include` (List of String) Include another port-group
- `port` (List of String) Port-group member

    &emsp;|Format     &emsp;|Description                                         |
    |-------------|------------------------------------------------------|
    &emsp;|txt        &emsp;|Named port (any name in /etc/services, e.g., http)  |
    &emsp;|1-65535    &emsp;|Numbered port                                       |
    &emsp;|start-end  &emsp;|Numbered port range (e.g. 1001-1050)                |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
