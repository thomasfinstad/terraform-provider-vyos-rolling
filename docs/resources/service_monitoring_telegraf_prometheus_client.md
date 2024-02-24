---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_monitoring_telegraf_prometheus_client Resource - vyos"
subcategory: "service"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>service</i>

  <br>
  &darr;
  <br>
  Monitoring services

  <br>
  &darr;
  <br>
  Telegraf metric collector

  <br>
  &darr;
  <br>
  <b>
  Output plugin Prometheus client
  </b>
  </div>
---

# vyos_service_monitoring_telegraf_prometheus_client (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
Monitoring services

<br>
&darr;
<br>
Telegraf metric collector

<br>
&darr;
<br>
<b>
Output plugin Prometheus client
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `allow_from` (List of String) Networks allowed to query this server

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4net  &emsp; |  IP address and prefix length  |
    |  ipv6net  &emsp; |  IPv6 address and prefix length  |
- `listen_address` (String) Local IP addresses to listen on

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ipv4  &emsp; |  IPv4 address to listen for incoming connections  |
    |  ipv6  &emsp; |  IPv6 address to listen for incoming connections  |
- `metric_version` (Number) Metric version control mapping from Telegraf to Prometheus format

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-2  &emsp; |  Metric version (default: 2)  |
- `port` (Number) Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).