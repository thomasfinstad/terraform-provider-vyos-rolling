---
page_title: "vyos_qos_policy_shaper_hfsc_class Resource - terraform-provider-vyos"
subcategory: "qos"
description: |-
  Quality of Service (QoS)⯯Service Policy definitions⯯Hierarchical Fair Service Curve's policy⯯Class ID
---

# vyos_qos_policy_shaper_hfsc_class (Resource)
<center>

Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
Hierarchical Fair Service Curve's policy  
⯯  
**Class ID**


</center>

## Schema

### Required

- `class_id` (Number) Class ID

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|1-4095  &emsp;|Class Identifier  |
- `shaper_hfsc_id` (String) Hierarchical Fair Service Curve&#39;s policy

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `linkshare` (Attributes) Linkshare class settings (see [below for nested schema](#nestedatt--linkshare))
- `realtime` (Attributes) Realtime class settings (see [below for nested schema](#nestedatt--realtime))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `upperlimit` (Attributes) Upperlimit class settings (see [below for nested schema](#nestedatt--upperlimit))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--linkshare&#34;&gt;&lt;/a&gt;
### Nested Schema for `linkshare`

Optional:

- `d` (String) Service curve delay

    &emsp;|Format    &emsp;|Description           |
    |------------|------------------------|
    &emsp;|&lt;number&gt;  &emsp;|Time in milliseconds  |
- `m1` (String) Linkshare m1 parameter for class traffic

    &emsp;|Format        &emsp;|Description                                              |
    |----------------|-----------------------------------------------------------|
    &emsp;|&lt;number&gt;      &emsp;|Rate in kbit (kilobit per second)                        |
    &emsp;|&lt;number&gt;%%    &emsp;|Percentage of overall rate                               |
    &emsp;|&lt;number&gt;bit   &emsp;|bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    &emsp;|&lt;number&gt;ibit  &emsp;|kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    &emsp;|&lt;number&gt;ibps  &emsp;|kibps(1024&amp;8), mibps(1024^2&amp;8), gibps, tibps - Byte/sec  |
    &emsp;|&lt;number&gt;bps   &emsp;|bps(8),kbps(8&amp;10^3),mbps(8&amp;10^6), gbps, tbps - Byte/sec  |
- `m2` (String) Linkshare m2 parameter for class traffic

    &emsp;|Format        &emsp;|Description                                              |
    |----------------|-----------------------------------------------------------|
    &emsp;|&lt;number&gt;      &emsp;|Rate in kbit (kilobit per second)                        |
    &emsp;|&lt;number&gt;%%    &emsp;|Percentage of overall rate                               |
    &emsp;|&lt;number&gt;bit   &emsp;|bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    &emsp;|&lt;number&gt;ibit  &emsp;|kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    &emsp;|&lt;number&gt;ibps  &emsp;|kibps(1024&amp;8), mibps(1024^2&amp;8), gibps, tibps - Byte/sec  |
    &emsp;|&lt;number&gt;bps   &emsp;|bps(8),kbps(8&amp;10^3),mbps(8&amp;10^6), gbps, tbps - Byte/sec  |


&lt;a id=&#34;nestedatt--realtime&#34;&gt;&lt;/a&gt;
### Nested Schema for `realtime`

Optional:

- `d` (String) Service curve delay

    &emsp;|Format    &emsp;|Description           |
    |------------|------------------------|
    &emsp;|&lt;number&gt;  &emsp;|Time in milliseconds  |
- `m1` (String) Linkshare m1 parameter for class traffic

    &emsp;|Format        &emsp;|Description                                              |
    |----------------|-----------------------------------------------------------|
    &emsp;|&lt;number&gt;      &emsp;|Rate in kbit (kilobit per second)                        |
    &emsp;|&lt;number&gt;%%    &emsp;|Percentage of overall rate                               |
    &emsp;|&lt;number&gt;bit   &emsp;|bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    &emsp;|&lt;number&gt;ibit  &emsp;|kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    &emsp;|&lt;number&gt;ibps  &emsp;|kibps(1024&amp;8), mibps(1024^2&amp;8), gibps, tibps - Byte/sec  |
    &emsp;|&lt;number&gt;bps   &emsp;|bps(8),kbps(8&amp;10^3),mbps(8&amp;10^6), gbps, tbps - Byte/sec  |
- `m2` (String) Linkshare m2 parameter for class traffic

    &emsp;|Format        &emsp;|Description                                              |
    |----------------|-----------------------------------------------------------|
    &emsp;|&lt;number&gt;      &emsp;|Rate in kbit (kilobit per second)                        |
    &emsp;|&lt;number&gt;%%    &emsp;|Percentage of overall rate                               |
    &emsp;|&lt;number&gt;bit   &emsp;|bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    &emsp;|&lt;number&gt;ibit  &emsp;|kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    &emsp;|&lt;number&gt;ibps  &emsp;|kibps(1024&amp;8), mibps(1024^2&amp;8), gibps, tibps - Byte/sec  |
    &emsp;|&lt;number&gt;bps   &emsp;|bps(8),kbps(8&amp;10^3),mbps(8&amp;10^6), gbps, tbps - Byte/sec  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).


&lt;a id=&#34;nestedatt--upperlimit&#34;&gt;&lt;/a&gt;
### Nested Schema for `upperlimit`

Optional:

- `d` (String) Service curve delay

    &emsp;|Format    &emsp;|Description           |
    |------------|------------------------|
    &emsp;|&lt;number&gt;  &emsp;|Time in milliseconds  |
- `m1` (String) Linkshare m1 parameter for class traffic

    &emsp;|Format        &emsp;|Description                                              |
    |----------------|-----------------------------------------------------------|
    &emsp;|&lt;number&gt;      &emsp;|Rate in kbit (kilobit per second)                        |
    &emsp;|&lt;number&gt;%%    &emsp;|Percentage of overall rate                               |
    &emsp;|&lt;number&gt;bit   &emsp;|bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    &emsp;|&lt;number&gt;ibit  &emsp;|kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    &emsp;|&lt;number&gt;ibps  &emsp;|kibps(1024&amp;8), mibps(1024^2&amp;8), gibps, tibps - Byte/sec  |
    &emsp;|&lt;number&gt;bps   &emsp;|bps(8),kbps(8&amp;10^3),mbps(8&amp;10^6), gbps, tbps - Byte/sec  |
- `m2` (String) Linkshare m2 parameter for class traffic

    &emsp;|Format        &emsp;|Description                                              |
    |----------------|-----------------------------------------------------------|
    &emsp;|&lt;number&gt;      &emsp;|Rate in kbit (kilobit per second)                        |
    &emsp;|&lt;number&gt;%%    &emsp;|Percentage of overall rate                               |
    &emsp;|&lt;number&gt;bit   &emsp;|bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    &emsp;|&lt;number&gt;ibit  &emsp;|kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    &emsp;|&lt;number&gt;ibps  &emsp;|kibps(1024&amp;8), mibps(1024^2&amp;8), gibps, tibps - Byte/sec  |
    &emsp;|&lt;number&gt;bps   &emsp;|bps(8),kbps(8&amp;10^3),mbps(8&amp;10^6), gbps, tbps - Byte/sec  |  &emsp;|
