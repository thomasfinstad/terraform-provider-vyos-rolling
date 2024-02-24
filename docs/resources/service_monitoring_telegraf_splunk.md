---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_monitoring_telegraf_splunk Resource - vyos"
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
  Output plugin Splunk
  </b>
  </div>
---

# vyos_service_monitoring_telegraf_splunk (Resource)

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
Output plugin Splunk
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `url` (String) Remote URL

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  url  &emsp; |  Remote URL  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).