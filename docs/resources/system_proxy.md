---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_system_proxy Resource - vyos"
subcategory: "system"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>system</i>

  <br>
  &darr;
  <br>
  <b>
  Sets a proxy for system wide use
  </b>
  </div>
---

# vyos_system_proxy (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>system</i>

<br>
&darr;
<br>
<b>
Sets a proxy for system wide use
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `password` (String) Password used for authentication

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Password  |
- `port` (Number) Port number used by connection

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  number: 1-65535  &emsp; |  Numeric IP port  |
- `url` (String) Proxy URL
- `username` (String) Username used for authentication

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Username  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).