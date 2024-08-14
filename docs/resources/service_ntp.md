---
page_title: "vyos_service_ntp Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Network Time Protocol (NTP) configuration
---

# vyos_service_ntp (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
**Network Time Protocol (NTP) configuration**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `interface` (String) Interface to use

    |Format  &emsp;|Description     |
    |----------|------------------|
    |txt     &emsp;|Interface name  |
- `leap_second` (String) Leap second behavior

    |Format    &emsp;|Description                                                                  |
    |------------|-------------------------------------------------------------------------------|
    |ignore    &emsp;|No correction is applied to the clock for the leap second                    |
    |smear     &emsp;|Correct served time slowly be slewing instead of stepping                    |
    |system    &emsp;|Kernel steps the system clock forward or backward                            |
    |timezone  &emsp;|Use UTC timezone database to determine when will the next leap second occur  |
- `listen_address` (List of String) Local IP addresses to listen on

    |Format  &emsp;|Description                                      |
    |----------|---------------------------------------------------|
    |ipv4    &emsp;|IPv4 address to listen for incoming connections  |
    |ipv6    &emsp;|IPv6 address to listen for incoming connections  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `vrf` (String) VRF instance name

    |Format  &emsp;|Description        |
    |----------|---------------------|
    |txt     &emsp;|VRF instance name  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
