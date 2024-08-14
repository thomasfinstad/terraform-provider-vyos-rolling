---
page_title: "vyos_service_snmp_v3_group Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Simple Network Management Protocol (SNMP)⯯Simple Network Management Protocol (SNMP) v3⯯Specifies the group with name groupname
---

# vyos_service_snmp_v3_group (Resource)
<center>

*service*  
⯯  
Simple Network Management Protocol (SNMP)  
⯯  
Simple Network Management Protocol (SNMP) v3  
⯯  
**Specifies the group with name groupname**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `mode` (String) Define access permission

    |Format  &emsp;|Description  |
    |----------|---------------|
    |ro      &emsp;|Read-Only    |
    |rw      &emsp;|read write   |
- `seclevel` (String) Security levels

    |Format  &emsp;|Description                                                  |
    |----------|---------------------------------------------------------------|
    |noauth  &emsp;|Messages not authenticated and not encrypted (noAuthNoPriv)  |
    |auth    &emsp;|Messages are authenticated but not encrypted (authNoPriv)    |
    |priv    &emsp;|Messages are authenticated and encrypted (authPriv)          |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `view` (String) Defines the name of view

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `group` (String) Specifies the group with name groupname


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
