---
page_title: "vyos_service_suricata_port_group Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Network IDS, IPS and Security Monitoring⯯Port group name
---

# vyos_service_suricata_port_group (Resource)
<center>

*service*  
⯯  
Network IDS, IPS and Security Monitoring  
⯯  
**Port group name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `group` (List of String) Port group

    |Format  &emsp;|Description                                    |
    |----------|-------------------------------------------------|
    |txt     &emsp;|Port group to match                            |
    |!txt    &emsp;|Exclude the specified port group from matches  |
- `port` (List of String) Port number

    |Format      &emsp;|Description                                                    |
    |--------------|-----------------------------------------------------------------|
    |1-65535     &emsp;|Numeric port to match                                          |
    |!1-65535    &emsp;|Numeric port to exclude from matches                           |
    |start-end   &emsp;|Numbered port range (e.g. 1001-1005) to match                  |
    |!start-end  &emsp;|Numbered port range (e.g. !1001-1005) to exclude from matches  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `port_group` (String) Port group name


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
