---
page_title: "vyos_firewall_ipv6_output_filter Resource - vyos"

subcategory: "Firewall"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  Firewall⯯IPv6 firewall⯯IPv6 output firewall⯯IPv6 firewall output filter
---

# vyos_firewall_ipv6_output_filter (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

Firewall  
⯯  
IPv6 firewall  
⯯  
IPv6 output firewall  
⯯  
**IPv6 firewall output filter**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `default_action` (String) Default-action for rule-set

    |Format  &emsp;|Description                       |
    |----------|------------------------------------|
    |drop    &emsp;|Drop if no prior rules are hit    |
    |accept  &emsp;|Accept if no prior rules are hit  |
- `default_log` (Boolean) Log packets hitting default-action
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
