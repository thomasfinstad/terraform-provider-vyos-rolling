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

<!--TOC-->

- [vyos_service_dns_forwarding_authoritative_domain_records_naptr_rule (Resource)](#vyos_service_dns_forwarding_authoritative_domain_records_naptr_rule-resource)
  - [Schema](#schema)
    - [Required](#required)
      - [identifier](#identifier)
    - [Optional](#optional)
      - [lookup_a](#lookup_a)
      - [lookup_srv](#lookup_srv)
      - [order](#order)
      - [preference](#preference)
      - [protocol_specific](#protocol_specific)
      - [regexp](#regexp)
      - [replacement](#replacement)
      - [resolve_uri](#resolve_uri)
      - [service](#service)
      - [timeouts](#timeouts)
    - [Read-Only](#read-only)
      - [id](#id)
    - [Nested Schema for `identifier`](#nested-schema-for-identifier)
    - [Nested Schema for `timeouts`](#nested-schema-for-timeouts)
  - [Import](#import)

<!--TOC-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

#### identifier
- `identifier` (Attributes) (see [below for nested schema](#nestedatt--identifier))

### Optional

#### lookup_a
- `lookup_a` (Boolean) A flag
#### lookup_srv
- `lookup_srv` (Boolean) S flag
#### order
- `order` (Number) Rule order

    |  Format   &emsp;|  Description                                  |
    |-----------|-----------------------------------------------|
    |  0-65535  &emsp;|  Rule order (lower order is evaluated first)  |
#### preference
- `preference` (Number) Rule preference

    |  Format   &emsp;|  Description      |
    |-----------|-------------------|
    |  0-65535  &emsp;|  Rule preference  |
#### protocol_specific
- `protocol_specific` (Boolean) P flag
#### regexp
- `regexp` (String) Regular expression
#### replacement
- `replacement` (String) Replacement DNS name

    |  Format            &emsp;|  Description        |
    |--------------------|---------------------|
    |  name.example.com  &emsp;|  Absolute DNS name  |
#### resolve_uri
- `resolve_uri` (Boolean) U flag
#### service
- `service` (String) Service type
#### timeouts
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

#### id
- `id` (String) Resource ID, full vyos path to the resource with each field separated by dunder (`__`).

<a id="nestedatt--identifier"></a>
### Nested Schema for `identifier`

Required:

- `authoritative_domain` (String) Domain to host authoritative records for

    |  Format  &emsp;|  Description                  |
    |----------|-------------------------------|
    |  txt     &emsp;|  An absolute DNS domain name  |
- `naptr` (String) NAPTR record

    |  Format  &emsp;|  Description                             |
    |----------|------------------------------------------|
    |  txt     &emsp;|  A DNS name relative to the root record  |
    |  @       &emsp;|  Root record                             |
- `rule` (Number) NAPTR rule

    |  Format   &emsp;|  Description  |
    |-----------|---------------|
    |  0-65535  &emsp;|  Rule number  |


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as &#34;30s&#34; or &#34;2h45m&#34;. Valid time units are &#34;s&#34; (seconds), &#34;m&#34; (minutes), &#34;h&#34; (hours).

## Import

Import is supported using the following syntax:

```shell
terraform import vyos_service_dns_forwarding_authoritative_domain_records_naptr_rule.example "service__dns__forwarding__authoritative-domain__<authoritative-domain>__records__naptr__<naptr>__rule__<rule>"
```
