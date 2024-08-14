---
page_title: "vyos_service_dhcp_relay Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Host Configuration Protocol (DHCP) relay agent
---

# vyos_service_dhcp_relay (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
**Host Configuration Protocol (DHCP) relay agent**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `disable` (Boolean) Disable instance
- `interface` (List of String) Interface to use

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |
- `listen_interface` (List of String) Interface for DHCP Relay Agent to listen for requests

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |
- `server` (List of String) DHCP server address

    |Format  &emsp;|Description               |
    |----------|----------------------------|
    |ipv4    &emsp;|DHCP server IPv4 address  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `upstream_interface` (List of String) Interface for DHCP Relay Agent forward requests out

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
