---
page_title: "vyos_system_sflow Resource - vyos"

subcategory: "System"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  system⯯sFlow settings
---

# vyos_system_sflow (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*system*  
⯯  
**sFlow settings**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `agent_address` (String) sFlow agent IPv4 or IPv6 address

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|ipv4    &emsp;|sFlow IPv4 agent address  |
    &emsp;|ipv6    &emsp;|sFlow IPv6 agent address  |
- `agent_interface` (String) IP address associated with this interface

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `drop_monitor_limit` (Number) Export headers of dropped by kernel packets

    &emsp;|Format   &emsp;|Description                                                               |
    |-----------|----------------------------------------------------------------------------|
    &emsp;|1-65535  &emsp;|Maximum rate limit of N drops per second send out in the sFlow datagrams  |
- `interface` (List of String) Interface to use

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|txt     &emsp;|Interface name  |
- `polling` (Number) Schedule counter-polling in seconds

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-600   &emsp;|Polling rate in seconds  |
- `sampling_rate` (Number) sFlow sampling-rate

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|1-65535  &emsp;|Sampling rate (1 in N packets)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF instance name

    &emsp;|Format  &emsp;|Description        |
    |----------|---------------------|
    &emsp;|txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
