---
page_title: "vyos_qos_policy_shaper_class Resource - vyos"

subcategory: "Qos"

description: |- 
  Quality of Service (QoS)⯯Service Policy definitions⯯Traffic shaping based policy (Hierarchy Token Bucket)⯯Class ID
---

# vyos_qos_policy_shaper_class (Resource)
<center>

Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
Traffic shaping based policy (Hierarchy Token Bucket)  
⯯  
**Class ID**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `class_id` (Number) Class ID

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|2-4095  &emsp;|Class Identifier  |
- `shaper_id` (String) Traffic shaping based policy (Hierarchy Token Bucket)

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |

### Optional

- `average_packet` (Number) Average packet size (bytes)

    &emsp;|Format    &emsp;|Description                   |
    |------------|--------------------------------|
    &emsp;|16-10240  &emsp;|Average packet size in bytes  |
- `bandwidth` (String) Available bandwidth for this policy

    &emsp;|Format        &emsp;|Description                         |
    |----------------|--------------------------------------|
    &emsp;|auto          &emsp;|Bandwidth matches interface speed   |
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
- `ceiling` (String) Bandwidth limit for this class

    &emsp;|Format        &emsp;|Description                                              |
    |----------------|-----------------------------------------------------------|
    &emsp;|&lt;number&gt;      &emsp;|Rate in kbit (kilobit per second)                        |
    &emsp;|&lt;number&gt;%%    &emsp;|Percentage of overall rate                               |
    &emsp;|&lt;number&gt;bit   &emsp;|bit(1), kbit(10^3), mbit(10^6), gbit, tbit               |
    &emsp;|&lt;number&gt;ibit  &emsp;|kibit(1024), mibit(1024^2), gibit(1024^3), tbit(1024^4)  |
    &emsp;|&lt;number&gt;ibps  &emsp;|kibps(1024&amp;8), mibps(1024^2&amp;8), gibps, tibps - Byte/sec  |
    &emsp;|&lt;number&gt;bps   &emsp;|bps(8),kbps(8&amp;10^3),mbps(8&amp;10^6), gbps, tbps - Byte/sec  |
- `codel_quantum` (Number) Deficit in the fair queuing algorithm

    &emsp;|Format     &emsp;|Description                        |
    |-------------|-------------------------------------|
    &emsp;|0-1048576  &emsp;|Number of bytes used as &#39;deficit&#39;  |
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `flows` (Number) Number of flows into which the incoming packets are classified

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65536  &emsp;|Number of flows  |
- `interval` (Number) Interval used to measure the delay

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|u32     &emsp;|Interval in milliseconds  |
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
- `priority` (Number) Priority for rule evaluation

    &emsp;|Format  &emsp;|Description                         |
    |----------|--------------------------------------|
    &emsp;|0-20    &emsp;|Priority for match rule evaluation  |
- `queue_limit` (Number) Maximum queue size

    &emsp;|Format        &emsp;|Description            |
    |----------------|-------------------------|
    &emsp;|1-4294967295  &emsp;|Queue size in packets  |
- `queue_type` (String) Queue type for default traffic

    &emsp;|Format         &emsp;|Description                   |
    |-----------------|--------------------------------|
    &emsp;|drop-tail      &emsp;|First-In-First-Out (FIFO)     |
    &emsp;|fair-queue     &emsp;|Stochastic Fair Queue (SFQ)   |
    &emsp;|fq-codel       &emsp;|Fair Queue Codel              |
    &emsp;|priority       &emsp;|Priority queuing              |
    &emsp;|random-detect  &emsp;|Random Early Detection (RED)  |
- `set_dscp` (String) Change the Differentiated Services (DiffServ) field in the IP header

    &emsp;|Format          &emsp;|Description                        |
    |------------------|-------------------------------------|
    &emsp;|0-63            &emsp;|Priority order for bandwidth pool  |
    &emsp;|default         &emsp;|match DSCP (000000)                |
    &emsp;|reliability     &emsp;|match DSCP (000001)                |
    &emsp;|throughput      &emsp;|match DSCP (000010)                |
    &emsp;|lowdelay        &emsp;|match DSCP (000100)                |
    &emsp;|priority        &emsp;|match DSCP (001000)                |
    &emsp;|immediate       &emsp;|match DSCP (010000)                |
    &emsp;|flash           &emsp;|match DSCP (011000)                |
    &emsp;|flash-override  &emsp;|match DSCP (100000)                |
    &emsp;|critical        &emsp;|match DSCP (101000)                |
    &emsp;|internet        &emsp;|match DSCP (110000)                |
    &emsp;|network         &emsp;|match DSCP (111000)                |
    &emsp;|AF11            &emsp;|High-throughput data               |
    &emsp;|AF12            &emsp;|High-throughput data               |
    &emsp;|AF13            &emsp;|High-throughput data               |
    &emsp;|AF21            &emsp;|Low-latency data                   |
    &emsp;|AF22            &emsp;|Low-latency data                   |
    &emsp;|AF23            &emsp;|Low-latency data                   |
    &emsp;|AF31            &emsp;|Multimedia streaming               |
    &emsp;|AF32            &emsp;|Multimedia streaming               |
    &emsp;|AF33            &emsp;|Multimedia streaming               |
    &emsp;|AF41            &emsp;|Multimedia conferencing            |
    &emsp;|AF42            &emsp;|Multimedia conferencing            |
    &emsp;|AF43            &emsp;|Multimedia conferencing            |
    &emsp;|CS1             &emsp;|Low-priority data                  |
    &emsp;|CS2             &emsp;|OAM                                |
    &emsp;|CS3             &emsp;|Broadcast video                    |
    &emsp;|CS4             &emsp;|Real-time interactive              |
    &emsp;|CS5             &emsp;|Signaling                          |
    &emsp;|CS6             &emsp;|Network control                    |
    &emsp;|CS7             &emsp;|N/A                                |
    &emsp;|EF              &emsp;|Expedited Forwarding               |
- `target` (Number) Acceptable minimum standing/persistent queue delay

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|u32     &emsp;|Queue delay in milliseconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
