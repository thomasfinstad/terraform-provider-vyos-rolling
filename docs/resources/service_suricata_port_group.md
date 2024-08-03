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

    &emsp;|Format  &emsp;|Description                                    |
    |----------|-------------------------------------------------|
    &emsp;|txt     &emsp;|Port group to match                            |
    &emsp;|!txt    &emsp;|Exclude the specified port group from matches  |
- `port` (List of String) Port number

    &emsp;|Format      &emsp;|Description                                                    |
    |--------------|-----------------------------------------------------------------|
    &emsp;|1-65535     &emsp;|Numeric port to match                                          |
    &emsp;|!1-65535    &emsp;|Numeric port to exclude from matches                           |
    &emsp;|start-end   &emsp;|Numbered port range (e.g. 1001-1005) to match                  |
    &emsp;|!start-end  &emsp;|Numbered port range (e.g. !1001-1005) to exclude from matches  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `port_group` (String) Port group name


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
