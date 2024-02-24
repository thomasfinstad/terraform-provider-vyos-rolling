---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_https_certificates Resource - vyos"
subcategory: "service"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>service</i>

  <br>
  &darr;
  <br>
  HTTPS configuration

  <br>
  &darr;
  <br>
  <b>
  TLS certificates
  </b>
  </div>
---

# vyos_service_https_certificates (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
HTTPS configuration

<br>
&darr;
<br>
<b>
TLS certificates
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ca_certificate` (String) Certificate Authority in PKI configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of CA in PKI configuration  |
- `certificate` (String) Certificate in PKI configuration

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  txt  &emsp; |  Name of certificate in PKI configuration  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).