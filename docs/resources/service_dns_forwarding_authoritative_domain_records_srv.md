---
page_title: "vyos_service_dns_forwarding_authoritative_domain_records_srv Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Domain Name System (DNS) related services⯯DNS forwarding⯯Domain to host authoritative records for⯯DNS zone records⯯SRV record
---

# vyos_service_dns_forwarding_authoritative_domain_records_srv (Resource)
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
**SRV record**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `disable` (Boolean) Disable instance
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))
- `ttl` (Number) Time-to-live (TTL)

    &emsp;|Format        &emsp;|Description     |
    |----------------|------------------|
    &emsp;|0-2147483647  &emsp;|TTL in seconds  |

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `authoritative_domain` (String) Domain to host authoritative records for

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|txt     &emsp;|An absolute DNS domain name  |
- `srv` (String) SRV record

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|txt     &emsp;|A DNS name relative to the root record  |
    &emsp;|@       &emsp;|Root record                             |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
