---
page_title: "vyos_service_conntrack_sync Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Connection tracking synchronization
---

# vyos_service_conntrack_sync (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
**Connection tracking synchronization**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `accept_protocol` (List of String) Protocols for which local conntrack entries will be synced

    &emsp;|Format  &emsp;|Description                                          |
    |----------|-------------------------------------------------------|
    &emsp;|tcp     &emsp;|Sync Transmission Control Protocol entries           |
    &emsp;|udp     &emsp;|Sync User Datagram Protocol entries                  |
    &emsp;|icmp    &emsp;|Sync Internet Control Message Protocol entries       |
    &emsp;|icmp6   &emsp;|Sync IPv6 Internet Control Message Protocol entries  |
    &emsp;|sctp    &emsp;|Sync Stream Control Transmission Protocol entries    |
    &emsp;|dccp    &emsp;|Sync Datagram Congestion Control Protocol entries    |
- `disable_external_cache` (Boolean) Directly injects the flow-states into the in-kernel Connection Tracking System of the backup firewall.
- `disable_syslog` (Boolean) Disable connection logging via Syslog
- `event_listen_queue_size` (Number) Queue size for local conntrack events

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|u32     &emsp;|Queue size in MB  |
- `expect_sync` (List of String) Protocol for which expect entries need to be synchronized
- `ignore_address` (List of String) IP addresses for which local conntrack entries will not be synced

    &emsp;|Format   &emsp;|Description             |
    |-----------|--------------------------|
    &emsp;|ipv4     &emsp;|IPv4 address to ignore  |
    &emsp;|ipv4net  &emsp;|IPv4 prefix to ignore   |
    &emsp;|ipv6     &emsp;|IPv6 address to ignore  |
    &emsp;|ipv6net  &emsp;|IPv6 prefix to ignore   |
- `listen_address` (List of String) Local IPv4 addresses to listen on

    &emsp;|Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    &emsp;|ipv4    &emsp;|IPv4 address to listen for incoming connections  |
- `mcast_group` (String) Multicast group to use for syncing conntrack entries
- `startup_resync` (Boolean) Order conntrackd to request a complete conntrack table resync against the other node at startup
- `sync_queue_size` (Number) Queue size for syncing conntrack entries

    &emsp;|Format  &emsp;|Description       |
    |----------|--------------------|
    &emsp;|u32     &emsp;|Queue size in MB  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
