---
page_title: "vyos_firewall_group_port_group Resource - vyos"

subcategory: "Firewall"

description: |- 
  Firewall⯯Firewall group⯯Firewall port-group
---

# vyos_firewall_group_port_group (Resource)
<center>

Firewall  
⯯  
Firewall group  
⯯  
**Firewall port-group**


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
- `include` (List of String) Include another port-group
- `port` (List of String) Port-group member

    |Format     &emsp;|Description                                         |
    |-------------|------------------------------------------------------|
    |txt        &emsp;|Named port (any name in /etc/services, e.g., http)  |
    |1-65535    &emsp;|Numbered port                                       |
    |start-end  &emsp;|Numbered port range (e.g. 1001-1050)                |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `port_group` (String) Firewall port-group


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
