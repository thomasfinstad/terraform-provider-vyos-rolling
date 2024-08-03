---
page_title: "vyos_load_balancing_wan_interface_health_test Resource - vyos"

subcategory: "Load"

description: |- 
  Configure load-balancing⯯Configure Wide Area Network (WAN) load-balancing⯯Interface name⯯Rule number
---

# vyos_load_balancing_wan_interface_health_test (Resource)
<center>

Configure load-balancing  
⯯  
Configure Wide Area Network (WAN) load-balancing  
⯯  
Interface name  
⯯  
**Rule number**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `resp_time` (Number) Ping response time (seconds)

    &emsp;|Format  &emsp;|Description              |
    |----------|---------------------------|
    &emsp;|1-30    &emsp;|Response time (seconds)  |
- `target` (String) Health target address

    &emsp;|Format  &emsp;|Description            |
    |----------|-------------------------|
    &emsp;|ipv4    &emsp;|Health target address  |
- `test_script` (String) Path to user-defined script

    &emsp;|Format  &emsp;|Description                |
    |----------|-----------------------------|
    &emsp;|txt     &emsp;|Script in /config/scripts  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `ttl_limit` (Number) TTL limit (hop count)

    &emsp;|Format  &emsp;|Description     |
    |----------|------------------|
    &emsp;|1-254   &emsp;|Number of hops  |
- `type` (String) WLB test type

    &emsp;|Format        &emsp;|Description                         |
    |----------------|--------------------------------------|
    &emsp;|ping          &emsp;|Test with ICMP echo response        |
    &emsp;|ttl           &emsp;|Test with UDP TTL expired response  |
    &emsp;|user-defined  &emsp;|User-defined test script            |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface_health` (String) Interface name
- `test` (Number) Rule number

    &emsp;|Format        &emsp;|Description  |
    |----------------|---------------|
    &emsp;|0-4294967295  &emsp;|Rule number  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
