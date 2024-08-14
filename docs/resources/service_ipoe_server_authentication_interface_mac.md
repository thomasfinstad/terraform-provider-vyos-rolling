---
page_title: "vyos_service_ipoe_server_authentication_interface_mac Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Internet Protocol over Ethernet (IPoE) Server⯯Client authentication methods⯯Network interface for client MAC addresses⯯Media Access Control (MAC) address
---

# vyos_service_ipoe_server_authentication_interface_mac (Resource)
<center>

*service*  
⯯  
Internet Protocol over Ethernet (IPoE) Server  
⯯  
Client authentication methods  
⯯  
Network interface for client MAC addresses  
⯯  
**Media Access Control (MAC) address**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `rate_limit` (Attributes) Upload/Download speed limits (see [below for nested schema](#nestedatt--rate_limit))
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vlan` (Number) VLAN monitor for automatic creation of VLAN interfaces

    |Format  &emsp;|Description     |
    |----------|------------------|
    |1-4094  &emsp;|Client VLAN id  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Network interface for client MAC addresses
- `mac` (String) Media Access Control (MAC) address

    |Format   &emsp;|Description             |
    |-----------|--------------------------|
    |macaddr  &emsp;|Hardware (MAC) address  |


<a id="nestedatt--rate_limit"></a>
### Nested Schema for `rate_limit`

Optional:

- `download` (String) Download bandwidth limit in kbits/sec
- `upload` (String) Upload bandwidth limit in kbits/sec


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
