---
page_title: "vyos_service_dns_forwarding_authoritative_domain_records_a Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Domain Name System (DNS) related services⯯DNS forwarding⯯Domain to host authoritative records for⯯DNS zone records⯯A record
---

# vyos_service_dns_forwarding_authoritative_domain_records_a (Resource)
<center>

*service*  
⯯  
Domain Name System (DNS) related services  
⯯  
DNS forwarding  
⯯  
Domain to host authoritative records for  
⯯  
DNS zone records  
⯯  
**A record**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `address` (List of String) IPv4 address

    |Format  &emsp;|Description   |
    |----------|----------------|
    |ipv4    &emsp;|IPv4 address  |
- `disable` (Boolean) Disable instance
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `ttl` (Number) Time-to-live (TTL)

    |Format        &emsp;|Description     |
    |----------------|------------------|
    |0-2147483647  &emsp;|TTL in seconds  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `a` (String) A record

    |Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    |txt     &emsp;|A DNS name relative to the root record  |
    |@       &emsp;|Root record                             |
    |any     &emsp;|Wildcard record (any subdomain)         |
- `authoritative_domain` (String) Domain to host authoritative records for

    |Format  &emsp;|Description                  |
    |----------|-------------------------------|
    |txt     &emsp;|An absolute DNS domain name  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
