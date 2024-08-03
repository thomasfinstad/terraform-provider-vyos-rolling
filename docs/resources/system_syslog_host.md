---
page_title: "vyos_system_syslog_host Resource - vyos"

subcategory: "System"

description: |- 
  system⯯System logging⯯Logging to remote host
---

# vyos_system_syslog_host (Resource)
<center>

*system*  
⯯  
System logging  
⯯  
**Logging to remote host**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `format` (Attributes) Logging format (see [below for nested schema](#nestedatt--format))
- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `protocol` (String) Protocol to be used (TCP/UDP)

    &emsp;|Format  &emsp;|Description          |
    |----------|-----------------------|
    &emsp;|udp     &emsp;|Listen protocol UDP  |
    &emsp;|tcp     &emsp;|Listen protocol TCP  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `host` (String) Logging to remote host

    &emsp;|Format    &emsp;|Description                        |
    |------------|-------------------------------------|
    &emsp;|ipv4      &emsp;|Remote syslog server IPv4 address  |
    &emsp;|ipv6      &emsp;|Remote syslog server IPv6 address  |
    &emsp;|hostname  &emsp;|Remote syslog server FQDN          |


&lt;a id=&#34;nestedatt--format&#34;&gt;&lt;/a&gt;
### Nested Schema for `format`

Optional:

- `octet_counted` (Boolean) Allows for the transmission of all characters inside a syslog message


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
