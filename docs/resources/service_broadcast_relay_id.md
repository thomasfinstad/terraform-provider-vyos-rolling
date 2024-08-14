---
page_title: "vyos_service_broadcast_relay_id Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯UDP broadcast relay service⯯Unique ID for each UDP port to forward
---

# vyos_service_broadcast_relay_id (Resource)
<center>

*service*  
⯯  
UDP broadcast relay service  
⯯  
**Unique ID for each UDP port to forward**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (String) Set source IP of forwarded packets, otherwise original senders address is used

    |Format  &emsp;|Description                                    |
    |----------|-------------------------------------------------|
    |ipv4    &emsp;|Optional source address for forwarded packets  |
- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `disable` (Boolean) Disable instance
- `interface` (List of String) Interface to use

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |
- `port` (Number) Port number used by connection

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |1-65535  &emsp;|Numeric IP port  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `id` (Number) Unique ID for each UDP port to forward

    |Format  &emsp;|Description                  |
    |----------|-------------------------------|
    |1-99    &emsp;|Broadcast relay instance ID  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
