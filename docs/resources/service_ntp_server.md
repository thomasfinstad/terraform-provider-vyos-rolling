---
page_title: "vyos_service_ntp_server Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Network Time Protocol (NTP) configuration⯯Network Time Protocol (NTP) server
---

# vyos_service_ntp_server (Resource)
<center>

*service*  
⯯  
Network Time Protocol (NTP) configuration  
⯯  
**Network Time Protocol (NTP) server**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `noselect` (Boolean) Marks the server as unused
- `nts` (Boolean) Enable Network Time Security (NTS) for the server
- `pool` (Boolean) Associate with a number of remote servers
- `prefer` (Boolean) Marks the server as preferred
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `server` (String) Network Time Protocol (NTP) server

    |Format    &emsp;|Description                                |
    |------------|---------------------------------------------|
    |ipv4      &emsp;|IP address of NTP server                   |
    |ipv6      &emsp;|IPv6 address of NTP server                 |
    |hostname  &emsp;|Fully qualified domain name of NTP server  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
