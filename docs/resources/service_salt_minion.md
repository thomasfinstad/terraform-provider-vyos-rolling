---
page_title: "vyos_service_salt_minion Resource - vyos"

subcategory: "Service"

description: |- 
  ~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.
  service⯯Salt Minion
---

# vyos_service_salt_minion (Resource)
<center>

~> This resource is global, having more than one resource of this type will cause configuration drift and possibly conflicts.

*service*  
⯯  
**Salt Minion**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Optional

- `hash` (String) Hash used when discovering file on master server (default: sha256)
- `id_param` (String) Explicitly declare ID for this minion to use (default: hostname)
- `interval` (Number) Interval in minutes between updates (default: 60)

    &emsp;|Format  &emsp;|Description                 |
    |----------|------------------------------|
    &emsp;|1-1440  &emsp;|Update interval in minutes  |
- `master` (List of String) Hostname or IP address of the Salt master server

    &emsp;|Format    &emsp;|Description               |
    |------------|----------------------------|
    &emsp;|ipv4      &emsp;|Salt server IPv4 address  |
    &emsp;|ipv6      &emsp;|Salt server IPv6 address  |
    &emsp;|hostname  &emsp;|Salt server FQDN address  |
- `master_key` (String) URL with signature of master for auth reply verification
- `source_interface` (String) Interface used to establish connection

    &emsp;|Format     &emsp;|Description     |
    |-------------|------------------|
    &emsp;|interface  &emsp;|Interface name  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
