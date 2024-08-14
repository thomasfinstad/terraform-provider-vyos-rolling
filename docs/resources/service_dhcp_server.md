---
page_title: "vyos_service_dhcp_server Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Dynamic Host Configuration Protocol (DHCP) for DHCP server
---

# vyos_service_dhcp_server (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
**Dynamic Host Configuration Protocol (DHCP) for DHCP server**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `disable` (Boolean) Disable instance
- `dynamic_dns_update` (Boolean) Dynamically update Domain Name System (RFC4702)
- `hostfile_update` (Boolean) Updating /etc/hosts file (per client lease)
- `listen_address` (List of String) Local IPv4 addresses to listen on

    |Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    |ipv4    &emsp;|IPv4 address to listen for incoming connections  |
- `listen_interface` (List of String) Interface to listen on

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
