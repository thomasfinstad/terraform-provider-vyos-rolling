---
page_title: "vyos_service_monitoring_telegraf Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Monitoring services⯯Telegraf metric collector
---

# vyos_service_monitoring_telegraf (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Monitoring services  
⯯  
**Telegraf metric collector**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `source` (List of String) Source parameters for monitoring

    |Format                &emsp;|Description                                          |
    |------------------------|-------------------------------------------------------|
    |all                   &emsp;|All parameters                                       |
    |hardware-utilization  &emsp;|Hardware-utilization parameters (CPU, disk, memory)  |
    |logs                  &emsp;|Logs parameters                                      |
    |network               &emsp;|Network parameters (net, netstat, nftables)          |
    |system                &emsp;|System parameters (system, processes, interrupts)    |
    |telegraf              &emsp;|Telegraf internal statistics                         |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF instance name

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
