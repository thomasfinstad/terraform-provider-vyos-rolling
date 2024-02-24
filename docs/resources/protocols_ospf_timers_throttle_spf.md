---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_protocols_ospf_timers_throttle_spf Resource - vyos"
subcategory: "protocols"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>protocols</i>

  <br>
  &darr;
  <br>
  Open Shortest Path First (OSPF)

  <br>
  &darr;
  <br>
  Adjust routing timers

  <br>
  &darr;
  <br>
  Throttling adaptive timers

  <br>
  &darr;
  <br>
  <b>
  OSPF SPF timers
  </b>
  </div>
---

# vyos_protocols_ospf_timers_throttle_spf (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>protocols</i>

<br>
&darr;
<br>
Open Shortest Path First (OSPF)

<br>
&darr;
<br>
Adjust routing timers

<br>
&darr;
<br>
Throttling adaptive timers

<br>
&darr;
<br>
<b>
OSPF SPF timers
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `delay` (Number) Delay from the first change received to SPF calculation

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-600000  &emsp; |  Delay in milliseconds  |
- `initial_holdtime` (Number) Initial hold time between consecutive SPF calculations

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-600000  &emsp; |  Initial hold time in milliseconds  |
- `max_holdtime` (Number) Maximum hold time

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 0-600000  &emsp; |  Max hold time in milliseconds  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).