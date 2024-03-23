---
page_title: "vyos_firewall_zone Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  Firewall
  ⯯
  Zone-policy
---

# vyos_firewall_zone (Resource)
<center>

Firewall
⯯
**Zone-policy**


</center>

## Schema

### Required

- `zone_id` (String) Zone-policy

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Zone name    |

### Optional

- `default_action` (String) Default-action for traffic coming into this zone

    &emsp;|Format  &emsp;|Description             |
    |----------|--------------------------|
    &emsp;|drop    &emsp;|Drop silently           |
    &emsp;|reject  &emsp;|Drop and notify source  |
- `default_log` (Boolean) Log packets hitting default-action
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `interface` (List of String) Interface associated with zone

    &emsp;|Format  &emsp;|Description                     |
    |----------|----------------------------------|
    &emsp;|txt     &emsp;|Interface associated with zone  |
    &emsp;|vrf     &emsp;|VRF associated with zone        |
- `intra_zone_filtering` (Attributes) Intra-zone filtering (see [below for nested schema](#nestedatt--intra_zone_filtering))
- `local_zone` (Boolean) Zone to be local-zone
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--intra_zone_filtering&#34;&gt;&lt;/a&gt;
### Nested Schema for `intra_zone_filtering`

Optional:

- `action` (String) Action for intra-zone traffic

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|accept  &emsp;|Accept traffic  |
    &emsp;|drop    &emsp;|Drop silently   |
- `firewall` (Attributes) Use the specified firewall chain (see [below for nested schema](#nestedatt--intra_zone_filtering--firewall))

&lt;a id=&#34;nestedatt--intra_zone_filtering--firewall&#34;&gt;&lt;/a&gt;
### Nested Schema for `intra_zone_filtering.firewall`

Optional:

- `ipv6_name` (String) IPv6 firewall ruleset
- `name` (String) IPv4 firewall ruleset



&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
