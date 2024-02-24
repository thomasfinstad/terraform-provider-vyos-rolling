---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vyos_service_lldp_snmp Resource - vyos"
subcategory: "service"
description: |-
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  <div style="text-align: center">
  <i>service</i>

  <br>
  &darr;
  <br>
  LLDP settings

  <br>
  &darr;
  <br>
  <b>
  SNMP parameters for LLDP
  </b>
  </div>
---

# vyos_service_lldp_snmp (Resource)

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

<div style="text-align: center">
<i>service</i>

<br>
&darr;
<br>
LLDP settings

<br>
&darr;
<br>
<b>
SNMP parameters for LLDP
</b>
</div>



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `enable` (Boolean) Enable SNMP queries of the LLDP database

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field seperated by dunder (`__`).