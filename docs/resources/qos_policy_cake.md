---
page_title: "vyos_qos_policy_cake Resource - vyos"

subcategory: "Qos"

description: |- 
  Quality of Service (QoS)⯯Service Policy definitions⯯Common Applications Kept Enhanced (CAKE)
---

# vyos_qos_policy_cake (Resource)
<center>

Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
**Common Applications Kept Enhanced (CAKE)**


</center>

## Schema

### Required

- `cake_id` (String) Common Applications Kept Enhanced (CAKE)

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |

### Optional

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
- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `flow_isolation` (Attributes) Flow isolation settings (see [below for nested schema](#nestedatt--flow_isolation))
- `rtt` (Number) Round-Trip-Time for Active Queue Management (AQM)

    &emsp;|Format     &emsp;|Description  |
    |-------------|---------------|
    &emsp;|1-3600000  &emsp;|RTT in ms    |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--flow_isolation&#34;&gt;&lt;/a&gt;
### Nested Schema for `flow_isolation`

Optional:

- `blind` (Boolean) Disables flow isolation, all traffic passes through a single queue
- `dst_host` (Boolean) Flows are defined only by destination address
- `dual_dst_host` (Boolean) Flows are defined by the 5-tuple, fairness is applied first over destination addresses, then over individual flows
- `dual_src_host` (Boolean) Flows are defined by the 5-tuple, fairness is applied first over source addresses, then over individual flows
- `flow` (Boolean) Flows are defined by the entire 5-tuple
- `host` (Boolean) Flows are defined by source-destination host pairs
- `nat` (Boolean) Perform NAT lookup before applying flow-isolation rules
- `src_host` (Boolean) Flows are defined only by source address
- `triple_isolate` (Boolean) Flows are defined by the 5-tuple, fairness is applied over source and destination addresses and also over individual flows (default)


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
