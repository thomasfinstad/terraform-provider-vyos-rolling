---
page_title: "vyos_service_suricata_address_group Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Network IDS, IPS and Security Monitoring⯯Address group name
---

# vyos_service_suricata_address_group (Resource)
<center>

*service*  
⯯  
Network IDS, IPS and Security Monitoring  
⯯  
**Address group name**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (List of String) IP address or subnet

    &emsp;|Format    &emsp;|Description                                      |
    |------------|---------------------------------------------------|
    &emsp;|ipv4      &emsp;|IPv4 address to match                            |
    &emsp;|ipv6      &emsp;|IPv6 address to match                            |
    &emsp;|ipv4net   &emsp;|IPv4 prefix to match                             |
    &emsp;|ipv6net   &emsp;|IPv6 prefix to match                             |
    &emsp;|!ipv4     &emsp;|Exclude the specified IPv4 address from matches  |
    &emsp;|!ipv6     &emsp;|Exclude the specified IPv6 address from matches  |
    &emsp;|!ipv4net  &emsp;|Exclude the specified IPv6 prefix from matches   |
    &emsp;|!ipv6net  &emsp;|Exclude the specified IPv6 prefix from matches   |
- `group` (List of String) Address group

    &emsp;|Format  &emsp;|Description                                       |
    |----------|----------------------------------------------------|
    &emsp;|txt     &emsp;|Address group to match                            |
    &emsp;|!txt    &emsp;|Exclude the specified address group from matches  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `address_group` (String) Address group name


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
