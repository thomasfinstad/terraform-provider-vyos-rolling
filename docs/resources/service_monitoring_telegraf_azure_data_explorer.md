---
page_title: "vyos_service_monitoring_telegraf_azure_data_explorer Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Monitoring services⯯Telegraf metric collector⯯Output plugin Azure Data Explorer
---

# vyos_service_monitoring_telegraf_azure_data_explorer (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Monitoring services  
⯯  
Telegraf metric collector  
⯯  
**Output plugin Azure Data Explorer**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `database` (String) Remote database name

    &emsp;|Format  &emsp;|Description           |
    |----------|------------------------|
    &emsp;|txt     &emsp;|Remote database name  |
- `group_metrics` (String) Type of metrics grouping when push to Azure Data Explorer

    &emsp;|Format            &emsp;|Description                                        |
    |--------------------|-----------------------------------------------------|
    &emsp;|single-table      &emsp;|Metrics stores in one table                        |
    &emsp;|table-per-metric  &emsp;|One table per gorups of metric by the metric name  |
- `table` (String) Name of the single table [Only if set group-metrics single-table]

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|txt     &emsp;|Table name   |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `url` (String) Remote URL

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|url     &emsp;|Remote HTTP(S) URL  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
