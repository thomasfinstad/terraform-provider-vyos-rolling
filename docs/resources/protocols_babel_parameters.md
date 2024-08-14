---
page_title: "vyos_protocols_babel_parameters Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Babel Routing Protocol⯯Babel-specific parameters
---

# vyos_protocols_babel_parameters (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Babel Routing Protocol  
⯯  
**Babel-specific parameters**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `diversity` (Boolean) Enable diversity-aware routing
- `diversity_factor` (Number) Multiplicative factor used for diversity routing

    |Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    |1-256   &emsp;|Multiplicative factor, in units of 1/256  |
- `resend_delay` (Number) Time before resending a message

    |Format     &emsp;|Description   |
    |-------------|----------------|
    |20-655340  &emsp;|Milliseconds  |
- `smoothing_half_life` (Number) Smoothing half-life

    |Format   &emsp;|Description  |
    |-----------|---------------|
    |0-65534  &emsp;|Seconds      |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
