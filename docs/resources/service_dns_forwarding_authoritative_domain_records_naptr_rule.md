---
page_title: "vyos_service_dns_forwarding_authoritative_domain_records_naptr_rule Resource - vyos"

subcategory: "Service"

description: |- 
  service⯯Domain Name System (DNS) related services⯯DNS forwarding⯯Domain to host authoritative records for⯯DNS zone records⯯NAPTR record⯯NAPTR rule
---

# vyos_service_dns_forwarding_authoritative_domain_records_naptr_rule (Resource)
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
NAPTR record  
⯯  
**NAPTR rule**


</center>

-> This provider is in "early access", please see the current status at: https://github.com/thomasfinstad/terraform-provider-vyos-rolling/milestones?state=open

## Schema

### Required

- `identifier` (Attributes Map) (see [below for nested schema](#nestedatt--identifier))

### Optional

- `lookup_a` (Boolean) A flag
- `lookup_srv` (Boolean) S flag
- `order` (Number) Rule order

    &emsp;|Format   &emsp;|Description                                  |
    |-----------|-----------------------------------------------|
    &emsp;|0-65535  &emsp;|Rule order (lower order is evaluated first)  |
- `preference` (Number) Rule preference

    &emsp;|Format   &emsp;|Description      |
    |-----------|-------------------|
    &emsp;|0-65535  &emsp;|Rule preference  |
- `protocol_specific` (Boolean) P flag
- `regexp` (String) Regular expression
- `replacement` (String) Replacement DNS name

    &emsp;|Format            &emsp;|Description        |
    |--------------------|---------------------|
    &emsp;|name.example.com  &emsp;|Absolute DNS name  |
- `resolve_uri` (Boolean) U flag
- `service` (String) Service type
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

&lt;a id=&#34;nestedatt--identifier&#34;&gt;&lt;/a&gt;
### Nested Schema for `identifier`

Required:

- `authoritative_domain` (String) Domain to host authoritative records for

    &emsp;|Format  &emsp;|Description                  |
    |----------|-------------------------------|
    &emsp;|txt     &emsp;|An absolute DNS domain name  |
- `naptr` (String) NAPTR record

    &emsp;|Format  &emsp;|Description                             |
    |----------|------------------------------------------|
    &emsp;|txt     &emsp;|A DNS name relative to the root record  |
    &emsp;|@       &emsp;|Root record                             |
- `rule` (Number) NAPTR rule

    &emsp;|Format   &emsp;|Description  |
    |-----------|---------------|
    &emsp;|0-65535  &emsp;|Rule number  |


&lt;a id=&#34;nestedatt--timeouts&#34;&gt;&lt;/a&gt;
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).  
