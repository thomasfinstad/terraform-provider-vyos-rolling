---
page_title: "vyos_firewall_zone Resource - vyos"

subcategory: "Firewall"

description: |- 
  Firewall⯯Zone-policy
---

# vyos_firewall_zone (Resource)
<center>

Firewall  
⯯  
**Zone-policy**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `default_action` (String) Default-action for traffic coming into this zone

    |Format  &emsp;|Description             |
    |----------|--------------------------|
    |drop    &emsp;|Drop silently           |
    |reject  &emsp;|Drop and notify source  |
- `default_log` (Boolean) Log packets hitting default-action
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `interface` (List of String) Interface associated with zone

    |Format  &emsp;|Description                     |
    |----------|----------------------------------|
    |txt     &emsp;|Interface associated with zone  |
    |vrf     &emsp;|VRF associated with zone        |
- `intra_zone_filtering` (Attributes) Intra-zone filtering (see [below for nested schema](#nestedatt--intra_zone_filtering))
- `local_zone` (Boolean) Zone to be local-zone
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `zone` (String) Zone-policy

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Zone name    |


<a id="nestedatt--intra_zone_filtering"></a>
### Nested Schema for `intra_zone_filtering`

Optional:

- `action` (String) Action for intra-zone traffic

    |Format  &emsp;|Description     |
    |----------|------------------|
    |accept  &emsp;|Accept traffic  |
    |drop    &emsp;|Drop silently   |
- `firewall` (Attributes) Use the specified firewall chain (see [below for nested schema](#nestedatt--intra_zone_filtering--firewall))

<a id="nestedatt--intra_zone_filtering--firewall"></a>
### Nested Schema for `intra_zone_filtering.firewall`

Optional:

- `ipv6_name` (String) IPv6 firewall ruleset
- `name` (String) IPv4 firewall ruleset



<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
