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

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `server` (String) Network Time Protocol (NTP) server

    &emsp;|Format    &emsp;|Description                                |
    |------------|---------------------------------------------|
    &emsp;|ipv4      &emsp;|IP address of NTP server                   |
    &emsp;|ipv6      &emsp;|IPv6 address of NTP server                 |
    &emsp;|hostname  &emsp;|Fully qualified domain name of NTP server  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
