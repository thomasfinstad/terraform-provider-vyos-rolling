---
page_title: "vyos_qos_interface Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)⯯Interface to apply QoS policy
---

# vyos_qos_interface (Resource)
<center>

Quality of Service (QoS)  
⯯  
**Interface to apply QoS policy**


</center>

## Schema

### Required

- `interface_id` (String) Interface to apply QoS policy

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |

### Optional

- `egress` (String) Interface egress traffic policy

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|QoS policy to use  |
- `ingress` (String) Interface ingress traffic policy

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|QoS policy to use  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  &emsp;|
