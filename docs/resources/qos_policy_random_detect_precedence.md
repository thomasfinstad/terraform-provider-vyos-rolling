---
page_title: "vyos_qos_policy_random_detect_precedence Resource - vyos"

subcategory: "Qos"

description: |- 
  Quality of Service (QoS)⯯Service Policy definitions⯯Weighted Random Early Detect policy⯯IP precedence
---

# vyos_qos_policy_random_detect_precedence (Resource)
<center>

Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
Weighted Random Early Detect policy  
⯯  
**IP precedence**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `average_packet` (Number) Average packet size (bytes)

    &emsp;|Format    &emsp;|Description                   |
    |------------|--------------------------------|
    &emsp;|16-10240  &emsp;|Average packet size in bytes  |
- `mark_probability` (Number) Mark probability for random detection

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|u32     &emsp;|Numeric value (1/N)  |
- `maximum_threshold` (Number) Maximum threshold for random detection

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|0-4096  &emsp;|Maximum threshold in packets  |
- `minimum_threshold` (Number) Minimum threshold for random detection

    &emsp;|Format  &emsp;|Description                   |
    |----------|--------------------------------|
    &emsp;|0-4096  &emsp;|Minimum threshold in packets  |
- `queue_limit` (Number) Maximum queue size

    &emsp;|Format        &emsp;|Description            |
    |----------------|-------------------------|
    &emsp;|1-4294967295  &emsp;|Queue size in packets  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `precedence` (Number) IP precedence

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|0-7     &emsp;|IP precedence value  |
- `random_detect` (String) Weighted Random Early Detect policy

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
