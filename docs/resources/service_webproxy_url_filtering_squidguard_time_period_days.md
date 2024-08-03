---
page_title: "vyos_service_webproxy_url_filtering_squidguard_time_period_days Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Webproxy service settings⯯URL filtering settings⯯URL filtering via squidGuard redirector⯯Time period name⯯Time-period days
---

# vyos_service_webproxy_url_filtering_squidguard_time_period_days (Resource)
<center>

*service*  
⯯  
Webproxy service settings  
⯯  
URL filtering settings  
⯯  
URL filtering via squidGuard redirector  
⯯  
Time period name  
⯯  
**Time-period days**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `time` (String) Time for time-period

    &emsp;|Format           &emsp;|Description              |
    |-------------------|---------------------------|
    &emsp;|&lt;hh:mm - hh:mm&gt;  &emsp;|Time range in 24hr time  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `days` (String) Time-period days

    &emsp;|Format    &emsp;|Description            |
    |------------|-------------------------|
    &emsp;|Sun       &emsp;|Sunday                 |
    &emsp;|Mon       &emsp;|Monday                 |
    &emsp;|Tue       &emsp;|Tuesday                |
    &emsp;|Wed       &emsp;|Wednesday              |
    &emsp;|Thu       &emsp;|Thursday               |
    &emsp;|Fri       &emsp;|Friday                 |
    &emsp;|Sat       &emsp;|Saturday               |
    &emsp;|weekdays  &emsp;|Monday through Friday  |
    &emsp;|weekend   &emsp;|Saturday and Sunday    |
    &emsp;|all       &emsp;|All days of the week   |
- `time_period` (String) Time period name


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
