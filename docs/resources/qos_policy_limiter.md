---
page_title: "vyos_qos_policy_limiter Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)
  ⯯
  Service Policy definitions
  ⯯
  Traffic input limiting policy
---

# vyos_qos_policy_limiter (Resource)
<center>

Quality of Service (QoS)
⯯
Service Policy definitions
⯯
**Traffic input limiting policy**


</center>

## Schema

### Required

- `limiter_id` (String) Traffic input limiting policy

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |

### Optional

- `default` (Attributes) Default policy (see [below for nested schema](#nestedatt--default))
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).

&lt;a id=&#34;nestedatt--default&#34;&gt;&lt;/a&gt;
### Nested Schema for `default`

Optional:

- `bandwidth` (String) Available bandwidth for this policy

    &emsp;|Format        &emsp;|Description                         |
    |----------------|--------------------------------------|
    &emsp;|&lt;number&gt;      &emsp;|Bits per second                     |
    &emsp;|&lt;number&gt;bit   &emsp;|Bits per second                     |
    &emsp;|&lt;number&gt;kbit  &emsp;|Kilobits per second                 |
    &emsp;|&lt;number&gt;mbit  &emsp;|Megabits per second                 |
    &emsp;|&lt;number&gt;gbit  &emsp;|Gigabits per second                 |
    &emsp;|&lt;number&gt;tbit  &emsp;|Terabits per second                 |
    &emsp;|&lt;number&gt;%%    &emsp;|Percentage of interface link speed  |
- `burst` (String) Burst size for this class

    &emsp;|Format            &emsp;|Description                             |
    |--------------------|------------------------------------------|
    &emsp;|&lt;number&gt;          &emsp;|Bytes                                   |
    &emsp;|&lt;number&gt;&lt;suffix&gt;  &emsp;|Bytes with scaling suffix (kb, mb, gb)  |
- `exceed` (String) Default action for packets exceeding the limiter

    &emsp;|Format      &emsp;|Description                                                                                                                   |
    |--------------|--------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|continue    &emsp;|Do not do anything, just continue with the next action in line                                                                |
    &emsp;|drop        &emsp;|Drop the packet immediately                                                                                                   |
    &emsp;|ok          &emsp;|Accept the packet                                                                                                             |
    &emsp;|reclassify  &emsp;|Treat the packet as non-matching to the filter this action is attached to and continue with the next filter in line (if any)  |
    &emsp;|pipe        &emsp;|Pass the packet to the next action in line                                                                                    |
- `mtu` (Number) MTU size for this class

    &emsp;|Format     &emsp;|Description  |
    |-------------|---------------|
    &emsp;|256-65535  &emsp;|Bytes        |
- `not_exceed` (String) Default action for packets not exceeding the limiter

    &emsp;|Format      &emsp;|Description                                                                                                                   |
    |--------------|--------------------------------------------------------------------------------------------------------------------------------|
    &emsp;|continue    &emsp;|Do not do anything, just continue with the next action in line                                                                |
    &emsp;|drop        &emsp;|Drop the packet immediately                                                                                                   |
    &emsp;|ok          &emsp;|Accept the packet                                                                                                             |
    &emsp;|reclassify  &emsp;|Treat the packet as non-matching to the filter this action is attached to and continue with the next filter in line (if any)  |
    &emsp;|pipe        &emsp;|Pass the packet to the next action in line                                                                                    |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
