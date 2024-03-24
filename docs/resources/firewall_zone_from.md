---
page_title: "vyos_firewall_zone_from Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall⯯Zone-policy⯯Zone from which to filter traffic
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
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--firewall&#34;&gt;&lt;/a&gt;
### Nested Schema for `firewall`

Optional:

- `ipv6_name` (String) IPv6 firewall ruleset
- `name` (String) IPv4 firewall ruleset


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
