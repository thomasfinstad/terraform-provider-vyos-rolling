---
page_title: "vyos_service_dhcp_server_high_availability Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Dynamic Host Configuration Protocol (DHCP) for DHCP server⯯DHCP high availability configuration
---

# vyos_service_dhcp_server_high_availability (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Dynamic Host Configuration Protocol (DHCP) for DHCP server  
⯯  
**DHCP high availability configuration**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `ca_certificate` (String) Certificate Authority in PKI configuration

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|txt     &emsp;|Name of CA in PKI configuration  |
- `certificate` (String) Certificate in PKI configuration

    &emsp;|Format  &emsp;|Description                               |
    |----------|--------------------------------------------|
    &emsp;|txt     &emsp;|Name of certificate in PKI configuration  |
- `mode` (String) Configure high availability mode

    &emsp;|Format          &emsp;|Description                                |
    |------------------|---------------------------------------------|
    &emsp;|active-active   &emsp;|Both server attend DHCP requests           |
    &emsp;|active-passive  &emsp;|Only primary server attends DHCP requests  |
- `name` (String) Peer name used to identify connection
- `remote` (String) IPv4 remote address used for connection

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address of high availability peer  |
- `source_address` (String) IPv4 source address used to initiate connection

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|ipv4    &emsp;|IPv4 source address  |
- `status` (String) High availability hierarchy

    &emsp;|Format     &emsp;|Description                                     |
    |-------------|--------------------------------------------------|
    &emsp;|primary    &emsp;|Configure this server to be the primary node    |
    &emsp;|secondary  &emsp;|Configure this server to be the secondary node  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
