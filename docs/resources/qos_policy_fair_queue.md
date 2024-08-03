---
page_title: "vyos_qos_policy_fair_queue Resource - vyos"

subcategory: "Qos"

description: |- 
  Quality of Service (QoS)⯯Service Policy definitions⯯Stochastic Fairness Queueing
---

# vyos_qos_policy_fair_queue (Resource)
<center>

Quality of Service (QoS)  
⯯  
Service Policy definitions  
⯯  
**Stochastic Fairness Queueing**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `description` (String) Description

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Description  |
- `hash_interval` (Number) Interval in seconds for queue algorithm perturbation

    &emsp;|Format  &emsp;|Description                                                         |
    |----------|----------------------------------------------------------------------|
    &emsp;|0       &emsp;|No perturbation                                                     |
    &emsp;|1-127   &emsp;|Interval in seconds for queue algorithm perturbation (advised: 10)  |
- `queue_limit` (Number) Upper limit of the SFQ

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|1-127   &emsp;|Queue size in packets  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `fair_queue` (String) Stochastic Fairness Queueing

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Policy name  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
