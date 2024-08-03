---
page_title: "vyos_service_conntrack_sync_interface Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Connection tracking synchronization⯯Interface to use for syncing conntrack entries
---

# vyos_service_conntrack_sync_interface (Resource)
<center>

*service*  
⯯  
Connection tracking synchronization  
⯯  
**Interface to use for syncing conntrack entries**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `peer` (String) IP address of the peer to send the UDP conntrack info too. This disable multicast.

    &emsp;|Format  &emsp;|Description                                    |
    |----------|-------------------------------------------------|
    &emsp;|ipv4    &emsp;|IP address to listen for incoming connections  |
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface to use for syncing conntrack entries


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
