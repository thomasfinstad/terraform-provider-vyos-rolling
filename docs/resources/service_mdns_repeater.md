---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_mdns_repeater Resource - vyos"
subcategory: "service"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>service</i>

  <br>
  &darr;
  <br>
  Multicast DNS (mDNS) parameters

  <br>
  &darr;
  <br>
  <b>
  mDNS repeater configuration
  </b>
  </div>
---

# vyos_service_mdns_repeater (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
Multicast DNS (mDNS) parameters

<br>
&darr;
<br>
<b>
mDNS repeater configuration
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `disable` (Boolean) Disable instance
- `interface` (List of String) Interface to use

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Interface name  |
- `vrrp_disable` (Boolean) Disables mDNS repeater on VRRP interfaces not in MASTER state

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).