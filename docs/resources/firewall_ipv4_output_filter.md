---
page_title: "vyos_firewall_ipv4_output_filter Resource - terraform-provider-vyos"
subcategory: "firewall"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Firewall
  ⯯IPv4 firewall⯯IPv4 output firewall⯯IPv4 firewall output filter
---

# vyos_firewall_ipv4_output_filter (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

	Firewall
⯯
IPv4 firewall
⯯
IPv4 output firewall
⯯
**IPv4 firewall output filter**


</center>

## Schema

### Optional

- `default_action` (String) Default-action for rule-set

    &emsp;|Format  &emsp;|Description                       |
    |----------|------------------------------------|
    &emsp;|drop    &emsp;|Drop if no prior rules are hit    |
    &emsp;|accept  &emsp;|Accept if no prior rules are hit  |
- `default_log` (Boolean) Log packets hitting default-action
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
