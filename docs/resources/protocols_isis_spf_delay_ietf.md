---
page_title: "vyos_protocols_isis_spf_delay_ietf Resource - vyos"

subcategory: "Protocols"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  protocols⯯Intermediate System to Intermediate System (IS-IS)⯯IETF SPF delay algorithm
---

# vyos_protocols_isis_spf_delay_ietf (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*protocols*  
⯯  
Intermediate System to Intermediate System (IS-IS)  
⯯  
**IETF SPF delay algorithm**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `holddown` (Number) Time with no received IGP events before considering IGP stable

    &emsp;|Format   &emsp;|Description                                                           |
    |-----------|------------------------------------------------------------------------|
    &emsp;|0-60000  &emsp;|Time with no received IGP events before considering IGP stable in ms  |
- `init_delay` (Number) Delay used while in QUIET state

    &emsp;|Format   &emsp;|Description                              |
    |-----------|-------------------------------------------|
    &emsp;|0-60000  &emsp;|Delay used while in QUIET state (in ms)  |
- `long_delay` (Number) Delay used while in LONG_WAIT

    &emsp;|Format   &emsp;|Description                                |
    |-----------|---------------------------------------------|
    &emsp;|0-60000  &emsp;|Delay used while in LONG_WAIT state in ms  |
- `short_delay` (Number) Delay used while in SHORT_WAIT state

    &emsp;|Format   &emsp;|Description                                   |
    |-----------|------------------------------------------------|
    &emsp;|0-60000  &emsp;|Delay used while in SHORT_WAIT state (in ms)  |
- `time_to_learn` (Number) Maximum duration needed to learn all the events related to a single failure

    &emsp;|Format   &emsp;|Description                                                                        |
    |-----------|-------------------------------------------------------------------------------------|
    &emsp;|0-60000  &emsp;|Maximum duration needed to learn all the events related to a single failure in ms  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  