---
page_title: "vyos_service_monitoring_telegraf_prometheus_client Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Monitoring services⯯Telegraf metric collector⯯Output plugin Prometheus client
---

# vyos_service_monitoring_telegraf_prometheus_client (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Monitoring services  
⯯  
Telegraf metric collector  
⯯  
**Output plugin Prometheus client**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `allow_from` (List of String) Networks allowed to query this server

    &emsp;|Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    &emsp;|ipv4net  &emsp;|IP address and prefix length    |
    &emsp;|ipv6net  &emsp;|IPv6 address and prefix length  |
- `listen_address` (String) Local IP addresses to listen on

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address to listen for incoming connections  |
    &emsp;|ipv6    &emsp;|IPv6 address to listen for incoming connections  |
- `metric_version` (Number) Metric version control mapping from Telegraf to Prometheus format

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|1-2     &emsp;|Metric version (default: 2)  |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
