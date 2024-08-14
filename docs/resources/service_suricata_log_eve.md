---
page_title: "vyos_service_suricata_log_eve Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Network IDS, IPS and Security Monitoring⯯Suricata log outputs⯯Extensible Event Format (EVE)
---

# vyos_service_suricata_log_eve (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
Network IDS, IPS and Security Monitoring  
⯯  
Suricata log outputs  
⯯  
**Extensible Event Format (EVE)**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `filename` (String) Log file

    |Format    &emsp;|Description                                  |
    |------------|-----------------------------------------------|
    |filename  &emsp;|File name in default Suricata log directory  |
    |/path     &emsp;|Absolute file path                           |
- `filetype` (String) EVE logging destination

    |Format   &emsp;|Description      |
    |-----------|-------------------|
    |regular  &emsp;|Log to filename  |
    |syslog   &emsp;|Log to syslog    |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `type` (List of String) Log types

    |Format                             &emsp;|Description                                                                                                                                                                                                                             |
    |-------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    |alert                              &emsp;|Record events for rule matches                                                                                                                                                                                                          |
    |anomaly                            &emsp;|Record unexpected conditions such as truncated packets, packets with invalid IP/UDP/TCP length values, and other events that render the packet invalid for further processing or describe unexpected behavior on an established stream  |
    |drop                               &emsp;|Record events for dropped packets                                                                                                                                                                                                       |
    |file                               &emsp;|Record file details (e.g., MD5) for files extracted from application protocols (e.g., HTTP)                                                                                                                                             |
    |application (http, dns, tls, ...)  &emsp;|Record application-level transactions                                                                                                                                                                                                   |
    |flow                               &emsp;|Record bi-directional flows                                                                                                                                                                                                             |
    |netflow                            &emsp;|Record uni-directional flows                                                                                                                                                                                                            |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
