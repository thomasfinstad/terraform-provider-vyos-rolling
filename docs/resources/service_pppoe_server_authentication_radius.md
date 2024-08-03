---
page_title: "vyos_service_pppoe_server_authentication_radius Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Point to Point over Ethernet (PPPoE) Server⯯Authentication for remote access PPPoE Server⯯RADIUS based user authentication
---

# vyos_service_pppoe_server_authentication_radius (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Point to Point over Ethernet (PPPoE) Server  
⯯  
Authentication for remote access PPPoE Server  
⯯  
**RADIUS based user authentication**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `accounting_interim_interval` (Number) Interval in seconds to send accounting information

    &emsp;|Format  &emsp;|Description                                         |
    |----------|------------------------------------------------------|
    &emsp;|1-3600  &emsp;|Interval in seconds to send accounting information  |
- `acct_interim_jitter` (Number) Maximum jitter value in seconds to be applied to accounting information interval

    &emsp;|Format  &emsp;|Description                      |
    |----------|-----------------------------------|
    &emsp;|1-60    &emsp;|Maximum jitter value in seconds  |
- `acct_timeout` (Number) Timeout for Interim-Update packets, terminate session afterwards

    &emsp;|Format  &emsp;|Description                           |
    |----------|----------------------------------------|
    &emsp;|0-60    &emsp;|Timeout in seconds, 0 to keep active  |
- `called_sid_format` (String) Format of Called-Station-Id attribute

    &emsp;|Format      &emsp;|Description                                                                                            |
    |--------------|---------------------------------------------------------------------------------------------------------|
    &emsp;|ifname      &emsp;|NAS-Port-Id - should contain root interface name (NAS-Port-Id=eth1)                                    |
    &emsp;|ifname:mac  &emsp;|NAS-Port-Id - should contain root interface name and mac address (NAS-Port-Id=eth1:00:00:00:00:00:00)  |
- `max_try` (Number) Number of tries to send Access-Request/Accounting-Request queries

    &emsp;|Format  &emsp;|Description    |
    |----------|-----------------|
    &emsp;|1-20    &emsp;|Maximum tries  |
- `nas_identifier` (String) NAS-Identifier attribute sent to RADIUS
- `nas_ip_address` (String) NAS-IP-Address attribute sent to RADIUS

    &emsp;|Format  &emsp;|Description               |
    |----------|----------------------------|
    &emsp;|ipv4    &emsp;|NAS-IP-Address attribute  |
- `preallocate_vif` (Boolean) Enable attribute NAS-Port-Id in Access-Request
- `source_address` (String) IPv4 source address used to initiate connection

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|ipv4    &emsp;|IPv4 source address  |
- `timeout` (Number) Timeout in seconds to wait response from RADIUS server

    &emsp;|Format  &emsp;|Description         |
    |----------|----------------------|
    &emsp;|1-60    &emsp;|Timeout in seconds  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
