---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_webproxy_authentication Resource - vyos"
subcategory: "service"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>service</i>

  <br>
  &darr;
  <br>
  Webproxy service settings

  <br>
  &darr;
  <br>
  <b>
  Proxy Authentication Settings
  </b>
  </div>
---

# vyos_service_webproxy_authentication (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
Webproxy service settings

<br>
&darr;
<br>
<b>
Proxy Authentication Settings
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `children` (String) Number of authentication helper processes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  n  &emsp; |  Number of authentication helper processes  |
- `credentials_ttl` (String) Authenticated session time to live in minutes

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  n  &emsp; |  Authenticated session timeout  |
- `method` (String) Authentication Method

    |  Format &emsp; | Description  |
    |----------|---------------|
    |  ldap  &emsp; |  Lightweight Directory Access Protocol  |
- `realm` (String) Name of authentication realm (e.g. "My Company proxy server")

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).