---
page_title: "vyos_service_dns_forwarding_domain_name_server Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Domain Name System (DNS) related services⯯DNS forwarding⯯Domain to forward to a custom DNS server⯯Domain Name Servers (DNS) addresses to forward queries to
---

# vyos_service_dns_forwarding_domain_name_server (Resource)
<center>

*service*  
⯯  
Domain Name System (DNS) related services  
⯯  
DNS forwarding  
⯯  
Domain to forward to a custom DNS server  
⯯  
**Domain Name Servers (DNS) addresses to forward queries to**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `port` (Number) Port number used by connection

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|1-65535  &emsp;|Numeric IP port  |
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `domain` (String) Domain to forward to a custom DNS server

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|txt     &emsp;|An absolute DNS domain name  |
- `name_server` (String) Domain Name Servers (DNS) addresses to forward queries to

    &emsp;|Format  &emsp;|Description                            |
    |----------|-----------------------------------------|
    &emsp;|ipv4    &emsp;|Domain Name Server (DNS) IPv4 address  |
    &emsp;|ipv6    &emsp;|Domain Name Server (DNS) IPv6 address  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
