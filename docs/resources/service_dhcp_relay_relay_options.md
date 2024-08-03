---
page_title: "vyos_service_dhcp_relay_relay_options Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Host Configuration Protocol (DHCP) relay agent⯯Relay options
---

# vyos_service_dhcp_relay_relay_options (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Host Configuration Protocol (DHCP) relay agent  
⯯  
**Relay options**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `hop_count` (Number) Policy to discard packets that have reached specified hop-count

    &emsp;|Format  &emsp;|Description  |
    |----------|---------------|
    &emsp;|1-255   &emsp;|Hop count    |
- `max_size` (Number) Maximum packet size to send to a DHCPv4/BOOTP server

    &emsp;|Format   &emsp;|Description          |
    |-----------|-----------------------|
    &emsp;|64-1400  &emsp;|Maximum packet size  |
- `relay_agents_packets` (String) Policy to handle incoming DHCPv4 packets which already contain relay agent options

    &emsp;|Format   &emsp;|Description                                                  |
    |-----------|---------------------------------------------------------------|
    &emsp;|append   &emsp;|append own relay options to packet                           |
    &emsp;|replace  &emsp;|replace existing agent option field                          |
    &emsp;|forward  &emsp;|forward packet unchanged                                     |
    &emsp;|discard  &emsp;|discard packet (default action if giaddr not set in packet)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
