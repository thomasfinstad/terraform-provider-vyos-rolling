---
page_title: "vyos_firewall_zone_from Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Zone-policy
  ⯯
  Zone from which to filter traffic
---

# vyos_firewall_zone_from (Resource)
<center>

Firewall
⯯
Zone-policy
⯯
**Zone from which to filter traffic**


</center>

## Schema

### Required

- `from_id` (String) Zone from which to filter traffic
- `zone_id` (String) Zone-policy

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Zone name    |

### Optional

- `firewall` (Attributes) Firewall options (see [below for nested schema](#nestedatt--firewall))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--firewall&#34;&gt;&lt;/a&gt;
### Nested Schema for `firewall`

Optional:

- `ipv6_name` (String) IPv6 firewall ruleset
- `name` (String) IPv4 firewall ruleset  &emsp;|
