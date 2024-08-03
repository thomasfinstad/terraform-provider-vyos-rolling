---
page_title: "vyos_service_router_advert_interface Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯IPv6 Router Advertisements (RAs) service⯯Interface to send RA on
---

# vyos_service_router_advert_interface (Resource)
<center>

*service*  
⯯  
IPv6 Router Advertisements (RAs) service  
⯯  
**Interface to send RA on**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `default_lifetime` (String) Lifetime associated with the default router in units of seconds

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|4-9000  &emsp;|Router Lifetime in seconds  |
    &emsp;|0       &emsp;|Not a default router        |
- `default_preference` (String) Preference associated with the default router,

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|low     &emsp;|Default router has low preference     |
    &emsp;|medium  &emsp;|Default router has medium preference  |
    &emsp;|high    &emsp;|Default router has high preference    |
- `dnssl` (List of String) DNS search list
- `hop_limit` (Number) Set Hop Count field of the IP header for outgoing packets

    &emsp;|Format  &emsp;|Description                                              |
    |----------|-----------------------------------------------------------|
    &emsp;|0       &emsp;|Unspecified (by this router)                             |
    &emsp;|1-255   &emsp;|Value should represent current diameter of the Internet  |
- `interval` (Attributes) Set interval between unsolicited multicast RAs (see [below for nested schema](#nestedatt--interval))
- `link_mtu` (Number) Link MTU value placed in RAs, exluded in RAs if unset

    &emsp;|Format     &emsp;|Description            |
    |-------------|-------------------------|
    &emsp;|1280-9000  &emsp;|Link MTU value in RAs  |
- `managed_flag` (Boolean) Hosts use the administered (stateful) protocol for address autoconfiguration in addition to any addresses autoconfigured using SLAAC
- `name_server` (List of String) Domain Name Servers (DNS) addresses

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|ipv6    &emsp;|Domain Name Server (DNS) IPv6 address  |
- `name_server_lifetime` (Number) Maximum duration how long the RDNSS entries are used

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|0       &emsp;|Name-servers should no longer be used  |
    &emsp;|1-7200  &emsp;|Maximum interval in seconds            |
- `no_send_advert` (Boolean) Do not send router adverts
- `other_config_flag` (Boolean) Hosts use the administered (stateful) protocol for autoconfiguration of other (non-address) information
- `reachable_time` (Number) Time, in milliseconds, that a node assumes a neighbor is reachable after having received a reachability confirmation

    &emsp;|Format     &emsp;|Description                                    |
    |-------------|-------------------------------------------------|
    &emsp;|0          &emsp;|Reachable Time unspecified by this router      |
    &emsp;|1-3600000  &emsp;|Reachable Time value in RAs (in milliseconds)  |
- `retrans_timer` (Number) Time in milliseconds between retransmitted Neighbor Solicitation messages

    &emsp;|Format        &emsp;|Description                                                                  |
    |----------------|-------------------------------------------------------------------------------|
    &emsp;|0             &emsp;|Time, in milliseconds, between retransmitted Neighbor Solicitation messages  |
    &emsp;|1-4294967295  &emsp;|Minimum interval in milliseconds                                             |
- `source_address` (List of String) Use IPv6 address as source address. Useful with VRRP.

    &emsp;|Format  &emsp;|Description                                                      |
    |----------|-------------------------------------------------------------------|
    &emsp;|ipv6    &emsp;|IPv6 address to be advertized (must be configured on interface)  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `interface` (String) Interface to send RA on


&lt;a id=&#34;nestedatt--interval&#34;&gt;&lt;/a&gt;
### Nested Schema for `interval`

Optional:

- `max` (Number) Maximum interval between unsolicited multicast RAs

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|4-1800  &emsp;|Maximum interval in seconds  |
- `min` (Number) Minimum interval between unsolicited multicast RAs

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|3-1350  &emsp;|Minimum interval in seconds  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
