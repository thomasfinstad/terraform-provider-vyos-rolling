---
page_title: "vyos_service_monitoring_telegraf_loki Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Monitoring services⯯Telegraf metric collector⯯Output plugin Loki
---

# vyos_service_monitoring_telegraf_loki (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Monitoring services  
⯯  
Telegraf metric collector  
⯯  
**Output plugin Loki**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `metric_name_label` (String) Metric name label

    |Format  &emsp;|Description                       |
    |----------|------------------------------------|
    |txt     &emsp;|Label to use for the metric name  |
- `port` (Number) Port number used by connection

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |1-65535  &emsp;|Numeric IP port  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `url` (String) Remote URL

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |url     &emsp;|Remote HTTP(S) URL  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
