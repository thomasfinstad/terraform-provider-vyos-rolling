---
page_title: "vyos_policy_route_map Resource - terraform-provider-vyos"
subcategory: "policy"
description: |-
  Routing policy
  ⯯
  IP route-map
---

# vyos_policy_route_map (Resource)
<center>

Routing policy
⯯
**IP route-map**


</center>

## Schema

### Required

- `route_map_id` (String) IP route-map

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Route map name  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).  &emsp;|
