---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_isis_spf_delay_ietf Resource - vyos"
subcategory: "protocols"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>protocols</i>

  <br>
  &darr;
  <br>
  Intermediate System to Intermediate System (IS-IS)

  <br>
  &darr;
  <br>
  <b>
  IETF SPF delay algorithm
  </b>
  </div>
---

# vyos_protocols_isis_spf_delay_ietf (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Intermediate System to Intermediate System (IS-IS)

<br>
&darr;
<br>
<b>
IETF SPF delay algorithm
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `holddown` (Number) Time with no received IGP events before considering IGP stable

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Time with no received IGP events before considering IGP stable in ms  |
- `init_delay` (Number) Delay used while in QUIET state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Delay used while in QUIET state (in ms)  |
- `long_delay` (Number) Delay used while in LONG_WAIT

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Delay used while in LONG_WAIT state in ms  |
- `short_delay` (Number) Delay used while in SHORT_WAIT state

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Delay used while in SHORT_WAIT state (in ms)  |
- `time_to_learn` (Number) Maximum duration needed to learn all the events related to a single failure

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-60000  &emsp; |  Maximum duration needed to learn all the events related to a single failure in ms  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).