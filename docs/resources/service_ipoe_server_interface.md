---
page_title: "vyos_service_ipoe_server_interface Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Internet Protocol over Ethernet (IPoE) Server⯯Interface to listen dhcp or unclassified packets
---

# vyos_service_ipoe_server_interface (Resource)
<center>

*service*  
⯯  
Internet Protocol over Ethernet (IPoE) Server  
⯯  
**Interface to listen dhcp or unclassified packets**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `client_subnet` (String) Client address pool

    |Format   &emsp;|Description                     |
    |-----------|----------------------------------|
    |ipv4net  &emsp;|IPv4 address and prefix length  |
- `external_dhcp` (Attributes) DHCP requests will be forwarded (see [below for nested schema](#nestedatt--external_dhcp))
- `mode` (String) Client connectivity mode

    |Format  &emsp;|Description                                 |
    |----------|----------------------------------------------|
    |l2      &emsp;|Client located on same interface as server  |
    |l3      &emsp;|Client located behind a router              |
- `network` (String) Enables clients to share the same network or each client has its own vlan

    |Format  &emsp;|Description                              |
    |----------|-------------------------------------------|
    |shared  &emsp;|Multiple clients share the same network  |
    |vlan    &emsp;|One VLAN per client                      |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vlan` (List of String) VLAN monitor for automatic creation of VLAN interfaces

    |Format     &emsp;|Description                                      |
    |-------------|---------------------------------------------------|
    |1-4094     &emsp;|VLAN for automatic creation                      |
    |start-end  &emsp;|VLAN range for automatic creation (e.g. 1-4094)  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface to listen dhcp or unclassified packets


<a id="nestedatt--external_dhcp"></a>
### Nested Schema for `external_dhcp`

Optional:

- `dhcp_relay` (String) DHCP Server the request will be redirected to.

    |Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    |ipv4    &emsp;|IPv4 address of the DHCP Server  |
- `giaddr` (String) Relay Agent IPv4 Address

    |Format  &emsp;|Description         |
    |----------|----------------------|
    |ipv4    &emsp;|Gateway IP address  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
