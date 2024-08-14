---
page_title: "vyos_interfaces_input Resource - vyos"

subcategory: "Interfaces"

description: |- 
  interfaces⯯Input Functional Block (IFB) interface name
---

# vyos_interfaces_input (Resource)
<center>

*interfaces*  
⯯  
**Input Functional Block (IFB) interface name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `description` (String) Description

    |Format  &emsp;|Description  |
    |----------|---------------|
    |txt     &emsp;|Description  |
- `disable` (Boolean) Administratively disable interface
- `redirect` (String) Redirect incoming packet to destination

    |Format  &emsp;|Description                 |
    |----------|------------------------------|
    |txt     &emsp;|Destination interface name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `input` (String) Input Functional Block (IFB) interface name

    |Format  &emsp;|Description           |
    |----------|------------------------|
    |ifbN    &emsp;|Input interface name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
